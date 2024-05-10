package com.oceanbase.samples.entity;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Objects;

/*
 * CREATE TABLE TEST(id integer, name varchar2(50))
 *
 */
@Entity
@Table( name = "test" )
public class TestEntity implements Serializable {

    private static final long serialVersionUID = -6578740021873269176L;

    @Id
    @GeneratedValue(strategy=GenerationType.IDENTITY)
    // @GeneratedValue(strategy=GenerationType.AUTO) //oracle 没有自增策略，添加该注解可以自动生成一个序列，提供自增主键，若数据库已有相关序列，可以忽 //略该注解。
    @Column(name = "id")
    private Integer testId;

    @Column( name = "name" )
    private String testName;



    public TestEntity(){

    }

    public TestEntity(String bauer) {
        this.testName = bauer;
    }


    public Integer getTestId() {
        return testId;
    }

    public void setTestId(Integer testId) {
        this.testId = testId;
    }

    public String getTestName() {
        return testName;
    }

    public void setTestName(String testName) {
        testName = testName;
    }

    @Override
    public String toString() {
        return "TestEntity{" +
            "testId=" + testId +
            ", TestName='" + testName + '\'' +
            '}';
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        TestEntity that = (TestEntity) o;
        return Objects.equals(testId, that.testId) && Objects.equals(testName, that.testName);
    }

    @Override
    public int hashCode() {
        return Objects.hash(testId, testName);
    }
}
