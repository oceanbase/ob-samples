/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.apache.hertzbeat.manager.config;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.usthe.sureness.util.JsonWebTokenUtil;
import jakarta.annotation.PostConstruct;
import jakarta.annotation.Resource;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.security.SecureRandom;
import java.util.Base64;
import java.util.HexFormat;
import java.util.Locale;
import java.util.Set;
import java.util.TimeZone;
import org.apache.hertzbeat.common.constants.CommonConstants;
import org.apache.hertzbeat.common.entity.manager.GeneralConfig;
import org.apache.hertzbeat.manager.dao.GeneralConfigDao;
import org.apache.hertzbeat.manager.pojo.dto.SystemConfig;
import org.apache.hertzbeat.manager.pojo.dto.SystemSecret;
import org.apache.hertzbeat.manager.pojo.dto.TemplateConfig;
import org.apache.hertzbeat.manager.service.AppService;
import org.apache.hertzbeat.manager.service.impl.SystemGeneralConfigServiceImpl;
import org.apache.hertzbeat.manager.service.impl.SystemSecretServiceImpl;
import org.apache.hertzbeat.manager.service.impl.TemplateConfigServiceImpl;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.DependsOn;
import org.springframework.context.annotation.Lazy;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;
import org.springframework.util.StringUtils;

/**
 * Common CommandLineRunner class
 */
@Component
@DependsOn("securityManager")
@Lazy(false)
@Order(value = Ordered.HIGHEST_PRECEDENCE + 2)
public class CommonCommandLineRunner implements CommandLineRunner {
    
    private static final Integer LANG_REGION_LENGTH = 2;
    
    private static final int JWT_SECRET_BYTES = 64;

    // Fingerprints of the public application and Helm defaults, for upgrade migration only.
    private static final Set<String> LEGACY_JWT_SECRET_HASHES = Set.of(
            "936f0ccb54f71860f92489142d579df73b6883c5e8a2d717bc27a89a4275996d",
            "0ad2a9aaa343676fd4d0dd09029054bedcecdd5506af5cad2fdce5087a656428");
    
    @Value("${sureness.jwt.secret:}")
    private String currentJwtSecret;
    
    @Resource
    private SystemGeneralConfigServiceImpl systemGeneralConfigService;
    
    @Resource
    private SystemSecretServiceImpl systemSecretService;
    
    @Resource
    private TemplateConfigServiceImpl templateConfigService;
    
    @Resource
    private AppService appService;
    
    @Resource
    protected GeneralConfigDao generalConfigDao;
    
    @Resource
    protected ObjectMapper objectMapper;

    @Override
    public void run(String... args) throws Exception {
        // for system config
        SystemConfig systemConfig = systemGeneralConfigService.getConfig();
        if (systemConfig != null) {
            if (systemConfig.getTimeZoneId() != null) {
                TimeZone.setDefault(TimeZone.getTimeZone(systemConfig.getTimeZoneId()));
            }
            if (systemConfig.getLocale() != null) {
                String[] arr = systemConfig.getLocale().split(CommonConstants.LOCALE_SEPARATOR);
                if (arr.length == LANG_REGION_LENGTH) {
                    String language = arr[0];
                    String country = arr[1];
                    Locale.setDefault(new Locale(language, country));   
                }
            }
        } else {
            // init system config data
            systemConfig = SystemConfig.builder().timeZoneId(TimeZone.getDefault().getID())
                                   .locale(Locale.getDefault().getLanguage() + CommonConstants.LOCALE_SEPARATOR 
                                                   + Locale.getDefault().getCountry())
                                   .build();
            String contentJson = objectMapper.writeValueAsString(systemConfig);
            GeneralConfig generalConfig2Save = GeneralConfig.builder()
                                                       .type(systemGeneralConfigService.type())
                                                       .content(contentJson)
                                                       .build();
            generalConfigDao.save(generalConfig2Save);
        }
        // for template config, flush the template config in db to memory
        TemplateConfig templateConfig = templateConfigService.getConfig();
        appService.updateCustomTemplateConfig(templateConfig);
    }

    // Override the authentication library's default before the HTTP server can start.
    @PostConstruct
    void initializeJwtSecret() throws NoSuchAlgorithmException {
        if (requiresGeneratedSecret(currentJwtSecret)) {
            SystemSecret systemSecret = systemSecretService.getConfig();
            if (systemSecret == null || requiresGeneratedSecret(systemSecret.getJwtSecret())) {
                byte[] randomBytes = new byte[JWT_SECRET_BYTES];
                new SecureRandom().nextBytes(randomBytes);
                currentJwtSecret = Base64.getEncoder().encodeToString(randomBytes);
                systemSecret = SystemSecret.builder().jwtSecret(currentJwtSecret).build();
                systemSecretService.saveConfig(systemSecret);
            } else {
                currentJwtSecret = systemSecret.getJwtSecret();
            }
        }
        // else use the user custom jwt secret
        // set the jwt secret token in util
        JsonWebTokenUtil.setDefaultSecretKey(currentJwtSecret);
    }

    private static boolean requiresGeneratedSecret(String secret) throws NoSuchAlgorithmException {
        if (!StringUtils.hasText(secret)) {
            return true;
        }
        byte[] digest = MessageDigest.getInstance("SHA-256").digest(secret.getBytes(StandardCharsets.UTF_8));
        return LEGACY_JWT_SECRET_HASHES.contains(HexFormat.of().formatHex(digest));
    }
}
