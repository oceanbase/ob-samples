package com.oceanbase.samples.springboot.entity;

import lombok.Data;

import javax.persistence.*;
import java.io.Serializable;

/*
 * CREATE TABLE TEST(id integer, name varchar2(50))
 *
 */
@Data
@Entity
@Table( name = "staff" )
public class StaffEntity implements Serializable {

    private static final long serialVersionUID = -6578740021873269176L;

    @Id
    @GeneratedValue(strategy=GenerationType.IDENTITY)
    // @GeneratedValue(strategy=GenerationType.AUTO) //oracle 没有自增策略，添加该注解可以自动生成一个序列，提供自增主键，若数据库已有相关序列，可以忽 //略该注解。
    @Column(name = "id")
    private Integer testId;

    @Column( name = "name" )
    private String testName;
}
