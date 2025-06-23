
---
## **1. MyBatis là gì?**

**MyBatis** là một framework **ORM (Object Relational Mapping)** nhẹ trong Java, giúp bạn kết nối và làm việc với cơ sở dữ liệu. Không giống như Hibernate, MyBatis không tự động ánh xạ các đối tượng với bảng. Thay vào đó, bạn cần định nghĩa các câu lệnh SQL trong file XML hoặc annotation.

**Ưu điểm của MyBatis**:

- Tùy chỉnh SQL dễ dàng.
    
- Hỗ trợ ánh xạ thủ công giữa bảng và đối tượng.
    
- Tối ưu hóa hiệu năng với caching (bộ nhớ đệm).
    
- Hỗ trợ stored procedures.


---
## **2. Cài đặt MyBatis**

### **2.1. Thêm dependency vào Maven hoặc Gradle**

#### **Với Maven**
```xml
<dependency>
    <groupId>org.mybatis</groupId>
    <artifactId>mybatis</artifactId>
    <version>3.5.13</version>
</dependency>
<dependency>
    <groupId>org.mybatis.spring.boot</groupId>
    <artifactId>mybatis-spring-boot-starter</artifactId>
    <version>3.0.1</version>
</dependency>
```

Với Gradle
```gradle
implementation 'org.mybatis:mybatis:3.5.13'
implementation 'org.mybatis.spring.boot:mybatis-spring-boot-starter:3.0.1'
```

2.2. Cấu hình `application.properties`
```
spring.datasource.url=jdbc:mysql://localhost:3306/my_database
spring.datasource.username=root
spring.datasource.password=your_password
spring.datasource.driver-class-name=com.mysql.cj.jdbc.Driver

# Cấu hình MyBatis
mybatis.config-location=classpath:mybatis-config.xml
mybatis.mapper-locations=classpath:mapper/*.xml
```

## **3. Cấu trúc thư mục MyBatis**

Cấu trúc thư mục điển hình khi sử dụng MyBatis:
```
src/main/java
  └── com.example.myapp
       ├── entity
       │    └── User.java
       ├── mapper
       │    ├── UserMapper.java
       │    └── xml
       │         └── UserMapper.xml
       └── service
            └── UserService.java
src/main/resources
  ├── application.properties
  ├── mybatis-config.xml
  ├── mapper
       └── UserMapper.xml
```

## **4. Làm quen với MyBatis**

### **4.1. Tạo entity**

Tạo một class đại diện cho bảng trong cơ sở dữ liệu:
```java
package com.example.myapp.entity;

public class User {
    private int id;
    private String username;
    private String email;

    // Getters và Setters
    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}
```

### **4.2. Tạo Mapper Interface**

Mapper interface định nghĩa các phương thức ánh xạ với SQL.
```java
package com.example.myapp.mapper;

import com.example.myapp.entity.User;
import org.apache.ibatis.annotations.*;

import java.util.List;

@Mapper
public interface UserMapper {
    @Select("SELECT * FROM users WHERE id = #{id}")
    User findById(int id);

    @Insert("INSERT INTO users(username, email) VALUES(#{username}, #{email})")
    @Options(useGeneratedKeys = true, keyProperty = "id")
    void insertUser(User user);

    @Update("UPDATE users SET email = #{email} WHERE id = #{id}")
    void updateUser(User user);

    @Delete("DELETE FROM users WHERE id = #{id}")
    void deleteUser(int id);

    @Select("SELECT * FROM users")
    List<User> findAll();
}
```

### **4.3. Tạo file XML Mapper**

Ngoài annotation, MyBatis hỗ trợ định nghĩa SQL trong file XML để dễ bảo trì hơn. Ví dụ:
```java
<!-- src/main/resources/mapper/UserMapper.xml -->
<mapper namespace="com.example.myapp.mapper.UserMapper">
    <select id="findById" resultType="com.example.myapp.entity.User">
        SELECT * FROM users WHERE id = #{id}
    </select>

    <insert id="insertUser" useGeneratedKeys="true" keyProperty="id">
        INSERT INTO users (username, email)
        VALUES (#{username}, #{email})
    </insert>

    <update id="updateUser">
        UPDATE users
        SET email = #{email}
        WHERE id = #{id}
    </update>

    <delete id="deleteUser">
        DELETE FROM users
        WHERE id = #{id}
    </delete>

    <select id="findAll" resultType="com.example.myapp.entity.User">
        SELECT * FROM users
    </select>
</mapper>
```

4.4. Cấu hình `mybatis-config.xml`
```java
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
    PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
    "http://mybatis.org/dtd/mybatis-3-config.dtd">
<configuration>
    <settings>
        <setting name="mapUnderscoreToCamelCase" value="true" />
    </settings>
</configuration>
```

## **5. Tích hợp MyBatis với Spring Boot**

### **5.1. Sử dụng Mapper**

Service sử dụng mapper để thực hiện các thao tác:
```java
package com.example.myapp.service;

import com.example.myapp.entity.User;
import com.example.myapp.mapper.UserMapper;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserService {
    private final UserMapper userMapper;

    public UserService(UserMapper userMapper) {
        this.userMapper = userMapper;
    }

    public List<User> getAllUsers() {
        return userMapper.findAll();
    }

    public User getUserById(int id) {
        return userMapper.findById(id);
    }

    public void createUser(User user) {
        userMapper.insertUser(user);
    }

    public void updateUser(User user) {
        userMapper.updateUser(user);
    }

    public void deleteUser(int id) {
        userMapper.deleteUser(id);
    }
}

```


## **6. Các tính năng nâng cao**

### **6.1. Ánh xạ phức tạp với kết quả tùy chỉnh**

Khi ánh xạ các bảng quan hệ phức tạp, MyBatis hỗ trợ **resultMap**:
```xml
<resultMap id="userWithDetails" type="com.example.myapp.entity.User">
    <id property="id" column="id"/>
    <result property="username" column="username"/>
    <result property="email" column="email"/>
</resultMap>

<select id="findByIdWithDetails" resultMap="userWithDetails">
    SELECT * FROM users WHERE id = #{id}
</select>
```

### **6.2. Caching (Bộ nhớ đệm)**

**MyBatis** hỗ trợ caching ở cấp 1 (session) và cấp 2 (global). Để bật caching cấp 2:
```xml
<cache />
```

### **6.3. Dynamic SQL**

Sử dụng **if**, **choose**, **where**... để tạo câu lệnh SQL động:
```xml
<select id="findUsersByConditions" resultType="com.example.myapp.entity.User">
    SELECT * FROM users
    <where>
        <if test="username != null">
            AND username = #{username}
        </if>
        <if test="email != null">
            AND email = #{email}
        </if>
    </where>
</select>
```

### **6.4. Plugin**

MyBatis cho phép tạo các plugin để can thiệp vào quy trình xử lý SQL (VD: logging, phân trang...).


