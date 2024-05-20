
---

Spring Data JPA là một phần của dự án Spring Data, cung cấp một cách dễ dàng để truy cập và thao tác với cơ sở dữ liệu quan hệ bằng cách sử dụng Java Persistence API (JPA). Spring Data JPA giúp bạn dễ dàng tạo các lớp repository để thực hiện các thao tác CRUD (Create, Read, Update, Delete) và các truy vấn phức tạp mà không cần phải viết nhiều mã boilerplate.


---
Dưới đây là hướng dẫn chi tiết để cài đặt và sử dụng Spring Data JPA trong một dự án sử dụng Gradle.

### Cấu hình Gradle

Trước tiên, bạn cần thêm các dependencies cần thiết vào file `build.gradle` của dự án. Spring Boot sẽ giúp bạn thiết lập mọi thứ dễ dàng hơn.

```build.gradle
plugins {
    id 'org.springframework.boot' version '2.7.3'
    id 'io.spring.dependency-management' version '1.0.11.RELEASE'
    id 'java'
}

group = 'com.example'
version = '0.0.1-SNAPSHOT'
sourceCompatibility = '11'

repositories {
    mavenCentral()
}

dependencies {
    implementation 'org.springframework.boot:spring-boot-starter-data-jpa'
    implementation 'org.springframework.boot:spring-boot-starter-web'
    runtimeOnly 'com.h2database:h2' // Sử dụng H2 cho mục đích học tập, bạn có thể thay bằng MySQL hoặc PostgreSQL
    testImplementation 'org.springframework.boot:spring-boot-starter-test'
}

test {
    useJUnitPlatform()
}

```

### Cấu hình cơ sở dữ liệu

Spring Boot sẽ tự động cấu hình kết nối cơ sở dữ liệu cho bạn dựa trên các thiết lập trong file `application.properties` hoặc `application.yml`.

```
spring.datasource.url=jdbc:h2:mem:testdb
spring.datasource.driverClassName=org.h2.Driver
spring.datasource.username=sa
spring.datasource.password=password
spring.jpa.database-platform=org.hibernate.dialect.H2Dialect
spring.h2.console.enabled=true

```

### Tạo Entity

Entity là một lớp Java đại diện cho một bảng trong cơ sở dữ liệu.

`src/main/java/com/example/demo/entity/Student.java`


```Java
package com.example.demo.entity;

import javax.persistence.*;

@Entity
@Table(name = "students")
public class Student {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "name", nullable = false)
    private String name;

    @Column(name = "email", nullable = false, unique = true)
    private String email;

    // Getters và setters

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}

```

### Tạo Repository

Repository là một interface cung cấp các phương thức để thao tác với cơ sở dữ liệu.

`src/main/java/com/example/demo/repository/StudentRepository.java`

```Java
package com.example.demo.repository;

import com.example.demo.entity.Student;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface StudentRepository extends JpaRepository<Student, Long> {
    // Bạn có thể thêm các phương thức truy vấn tùy chỉnh tại đây
}

```

### Tạo Service

Service chứa logic nghiệp vụ và tương tác với repository.

`src/main/java/com/example/demo/service/StudentService.java`

```Java
package com.example.demo.service;

import com.example.demo.entity.Student;
import com.example.demo.repository.StudentRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class StudentService {

    @Autowired
    private StudentRepository studentRepository;

    public List<Student> getAllStudents() {
        return studentRepository.findAll();
    }

    public Optional<Student> getStudentById(Long id) {
        return studentRepository.findById(id);
    }

    public Student createStudent(Student student) {
        return studentRepository.save(student);
    }

    public void deleteStudent(Long id) {
        studentRepository.deleteById(id);
    }
}

```

### Tạo Controller

Controller xử lý các request từ người dùng và trả về response.

`src/main/java/com/example/demo/controller/StudentController.java`


```Java
package com.example.demo.controller;

import com.example.demo.entity.Student;
import com.example.demo.service.StudentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/students")
public class StudentController {

    @Autowired
    private StudentService studentService;

    @GetMapping
    public List<Student> getAllStudents() {
        return studentService.getAllStudents();
    }

    @GetMapping("/{id}")
    public ResponseEntity<Student> getStudentById(@PathVariable Long id) {
        return studentService.getStudentById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }

    @PostMapping
    public Student createStudent(@RequestBody Student student) {
        return studentService.createStudent(student);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteStudent(@PathVariable Long id) {
        studentService.deleteStudent(id);
        return ResponseEntity.noContent().build();
    }
}

```

### Chạy ứng dụng

Sau khi đã cấu hình và viết mã xong, bạn có thể chạy ứng dụng bằng cách sử dụng Gradle.

