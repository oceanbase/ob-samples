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

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.clearInvocations;
import static org.mockito.Mockito.doAnswer;
import static org.mockito.Mockito.doThrow;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.mockStatic;
import static org.mockito.Mockito.never;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoInteractions;
import static org.mockito.Mockito.when;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.usthe.sureness.util.JsonWebTokenUtil;
import io.jsonwebtoken.JwtException;
import java.util.Base64;
import java.util.List;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicReference;
import java.util.stream.Stream;
import org.apache.hertzbeat.manager.dao.GeneralConfigDao;
import org.apache.hertzbeat.manager.pojo.dto.SystemSecret;
import org.apache.hertzbeat.manager.service.AppService;
import org.apache.hertzbeat.manager.service.impl.SystemGeneralConfigServiceImpl;
import org.apache.hertzbeat.manager.service.impl.SystemSecretServiceImpl;
import org.apache.hertzbeat.manager.service.impl.TemplateConfigServiceImpl;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;
import org.junit.jupiter.params.provider.ValueSource;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockedStatic;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.boot.LazyInitializationBeanFactoryPostProcessor;
import org.springframework.context.SmartLifecycle;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.test.util.ReflectionTestUtils;

/**
 * Tests JWT key selection, migration and persistence during startup.
 */
@ExtendWith(MockitoExtension.class)
class CommonCommandLineRunnerTest {

    // Public, compromised defaults retained only as migration regression fixtures.
    private static final String LEGACY_APPLICATION_SECRET = "CyaFv0bwq2Eik0jdrKUtsA6bx3sDJeFV643R "
            + "LnfKefTjsIfJLBa2YkhEqEGtcHDTNe4CU6+9 "
            + "8tVt4bisXQ13rbN0oxhUZR73M6EByXIO+SV5 "
            + "dKhaX0csgOCTlCxq20yhmUea6H6JIpSE2Rwp";

    private static final String LEGACY_HELM_SECRET = LEGACY_APPLICATION_SECRET.replace("FV643R", "FV143R");

    private static final String CUSTOM_SECRET = Base64.getEncoder().encodeToString(new byte[64]);

    @InjectMocks
    private CommonCommandLineRunner runner;

    @Mock
    private SystemGeneralConfigServiceImpl systemGeneralConfigService;

    @Mock
    private SystemSecretServiceImpl systemSecretService;

    @Mock
    private TemplateConfigServiceImpl templateConfigService;

    @Mock
    private AppService appService;

    static Stream<String> automaticSecrets() {
        return Stream.of(null, "", " \t", LEGACY_APPLICATION_SECRET, LEGACY_HELM_SECRET);
    }

    @ParameterizedTest
    @MethodSource("automaticSecrets")
    void generatesAndPersistsKeyForMissingOrLegacyConfiguration(String configuredSecret) throws Exception {
        ReflectionTestUtils.setField(runner, "currentJwtSecret", configuredSecret);
        try (MockedStatic<JsonWebTokenUtil> jwt = mockStatic(JsonWebTokenUtil.class)) {
            runner.initializeJwtSecret();
            ArgumentCaptor<SystemSecret> saved = ArgumentCaptor.forClass(SystemSecret.class);
            verify(systemSecretService).saveConfig(saved.capture());
            String key = saved.getValue().getJwtSecret();
            assertEquals(64, Base64.getDecoder().decode(key).length);
            assertNotEquals(LEGACY_APPLICATION_SECRET, key);
            assertNotEquals(LEGACY_HELM_SECRET, key);
            jwt.verify(() -> JsonWebTokenUtil.setDefaultSecretKey(key));
        }
    }

    @ParameterizedTest
    @MethodSource("automaticSecrets")
    void reusesPersistedKeyForMissingOrLegacyConfiguration(String configuredSecret) throws Exception {
        ReflectionTestUtils.setField(runner, "currentJwtSecret", configuredSecret);
        when(systemSecretService.getConfig()).thenReturn(SystemSecret.builder().jwtSecret(CUSTOM_SECRET).build());
        try (MockedStatic<JsonWebTokenUtil> jwt = mockStatic(JsonWebTokenUtil.class)) {
            runner.initializeJwtSecret();
            verify(systemSecretService, never()).saveConfig(any());
            jwt.verify(() -> JsonWebTokenUtil.setDefaultSecretKey(CUSTOM_SECRET));
        }
    }

    @ParameterizedTest
    @MethodSource("automaticSecrets")
    void replacesMissingOrLegacyPersistedKey(String persistedSecret) throws Exception {
        ReflectionTestUtils.setField(runner, "currentJwtSecret", "");
        when(systemSecretService.getConfig()).thenReturn(SystemSecret.builder().jwtSecret(persistedSecret).build());
        try (MockedStatic<JsonWebTokenUtil> jwt = mockStatic(JsonWebTokenUtil.class)) {
            runner.initializeJwtSecret();
            ArgumentCaptor<SystemSecret> saved = ArgumentCaptor.forClass(SystemSecret.class);
            verify(systemSecretService).saveConfig(saved.capture());
            String key = saved.getValue().getJwtSecret();
            assertEquals(64, Base64.getDecoder().decode(key).length);
            jwt.verify(() -> JsonWebTokenUtil.setDefaultSecretKey(key));
        }
    }

    @Test
    void preservesExplicitCustomKey() throws Exception {
        ReflectionTestUtils.setField(runner, "currentJwtSecret", CUSTOM_SECRET);
        try (MockedStatic<JsonWebTokenUtil> jwt = mockStatic(JsonWebTokenUtil.class)) {
            runner.initializeJwtSecret();
            verifyNoInteractions(systemSecretService);
            jwt.verify(() -> JsonWebTokenUtil.setDefaultSecretKey(CUSTOM_SECRET));
        }
    }

    @Test
    void generatedKeySignsTokensAndSurvivesRestart() throws Exception {
        ReflectionTestUtils.setField(runner, "currentJwtSecret", "");
        runner.initializeJwtSecret();
        ArgumentCaptor<SystemSecret> saved = ArgumentCaptor.forClass(SystemSecret.class);
        verify(systemSecretService).saveConfig(saved.capture());
        String token = JsonWebTokenUtil.issueJwt("regression-user", 60_000L, List.of("admin"));

        ReflectionTestUtils.setField(runner, "currentJwtSecret", "");
        when(systemSecretService.getConfig()).thenReturn(saved.getValue());
        clearInvocations(systemSecretService);
        runner.initializeJwtSecret();

        verify(systemSecretService, never()).saveConfig(any());
        assertEquals("regression-user", JsonWebTokenUtil.parseJwt(token).getSubject());
    }

    @Test
    void doesNotInstallKeyWhenPersistenceFails() {
        ReflectionTestUtils.setField(runner, "currentJwtSecret", "");
        doThrow(new IllegalStateException("storage unavailable")).when(systemSecretService).saveConfig(any());
        try (MockedStatic<JsonWebTokenUtil> jwt = mockStatic(JsonWebTokenUtil.class)) {
            assertThrows(IllegalStateException.class, () -> runner.initializeJwtSecret());
            jwt.verifyNoInteractions();
        }
    }

    @ParameterizedTest
    @ValueSource(booleans = {false, true})
    void initializesKeyBeforeLifecycleStarts(boolean lazyInitialization) {
        AtomicBoolean securityManagerInitialized = new AtomicBoolean();
        AtomicReference<String> legacyToken = new AtomicReference<>();
        AtomicBoolean persisted = new AtomicBoolean();
        AtomicBoolean lifecycleStarted = new AtomicBoolean();
        doAnswer(invocation -> {
            assertTrue(securityManagerInitialized.get(), "Starter must initialize before JWT key selection");
            persisted.set(true);
            return null;
        }).when(systemSecretService).saveConfig(any());
        try (AnnotationConfigApplicationContext context = new AnnotationConfigApplicationContext()) {
            if (lazyInitialization) {
                context.addBeanFactoryPostProcessor(new LazyInitializationBeanFactoryPostProcessor());
            }
            context.registerBean(CommonCommandLineRunner.class);
            // The starter installs its configured key while constructing this bean.
            context.registerBean("securityManager", Object.class, () -> {
                JsonWebTokenUtil.setDefaultSecretKey(LEGACY_APPLICATION_SECRET);
                legacyToken.set(JsonWebTokenUtil.issueJwt("legacy-test-user", 60_000L, List.of("admin")));
                securityManagerInitialized.set(true);
                return new Object();
            });
            context.registerBean(SystemGeneralConfigServiceImpl.class, () -> systemGeneralConfigService);
            context.registerBean(SystemSecretServiceImpl.class, () -> systemSecretService);
            context.registerBean(TemplateConfigServiceImpl.class, () -> templateConfigService);
            context.registerBean(AppService.class, () -> appService);
            context.registerBean(GeneralConfigDao.class, () -> mock(GeneralConfigDao.class));
            context.registerBean(ObjectMapper.class, () -> new ObjectMapper());
            context.registerBean("httpLifecycle", SmartLifecycle.class, () -> new SmartLifecycle() {
                @Override
                public void start() {
                    assertTrue(persisted.get(), "JWT key must be persisted before HTTP lifecycle starts");
                    assertThrows(JwtException.class, () -> JsonWebTokenUtil.parseJwt(legacyToken.get()));
                    lifecycleStarted.set(true);
                }

                @Override
                public void stop() {
                    lifecycleStarted.set(false);
                }

                @Override
                public boolean isRunning() {
                    return lifecycleStarted.get();
                }
            });
            context.refresh();
            assertTrue(lifecycleStarted.get());
        }
    }
}
