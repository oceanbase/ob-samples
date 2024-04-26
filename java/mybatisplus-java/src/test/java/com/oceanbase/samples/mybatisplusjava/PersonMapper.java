package com.oceanbase.samples.mybatisplusjava;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface PersonMapper extends BaseMapper<Person> {
    // 这里可以定义其他复杂的 SQL 操作
}
