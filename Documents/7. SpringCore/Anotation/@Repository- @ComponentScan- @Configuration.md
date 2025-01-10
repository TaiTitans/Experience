
---
Dưới đây là mô tả chi tiết về các annotation `@Repository`, `@ComponentScan`, và `@Configuration` trong Spring Framework:

### 1. **`@Repository`**

`@Repository` là một annotation trong Spring, được sử dụng để đánh dấu một lớp như là tầng repository (Data Access Object - DAO).

- **Mục đích**:
    - Dùng để tương tác với cơ sở dữ liệu.
    - Đánh dấu lớp này như một bean của Spring, đồng thời cung cấp khả năng xử lý các ngoại lệ liên quan đến dữ liệu (data access exceptions).
 **Lợi ích**:
- Giúp Spring hiểu rằng lớp này là một thành phần DAO, có thể tự động xử lý các ngoại lệ liên quan đến cơ sở dữ liệu thông qua cơ chế Spring Data Access Exception.

---
### 2. **`@ComponentScan`**

`@ComponentScan` là một annotation được sử dụng để chỉ định các gói (packages) mà Spring nên quét để tìm các bean (các lớp được đánh dấu với `@Component`, `@Service`, `@Repository`, `@Controller`, `@RestController`).

- **Mục đích**:
    - Tự động phát hiện và đăng ký các bean trong ApplicationContext.
**Cách sử dụng**:
```java
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

@Configuration
@ComponentScan(basePackages = "com.example.project")
public class AppConfig {
    // Cấu hình bổ sung nếu cần
}
```

**Lợi ích**:

- Giúp Spring quét và tự động quản lý các bean mà không cần phải khai báo thủ công trong tệp cấu hình XML hoặc Java.
- Hữu ích trong các ứng dụng lớn, nơi có nhiều package và các component phân tán.

---
### 3. **`@Configuration`**

`@Configuration` là một annotation trong Spring, được sử dụng để định nghĩa các class chứa các bean định nghĩa, tương đương với việc khai báo các bean trong file XML truyền thống.

- **Mục đích**:
    - Đánh dấu một lớp như là một lớp cấu hình (configuration class) trong Spring.
    - Chứa các phương thức `@Bean` để khai báo và quản lý các bean trong ApplicationContext.
```java
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AppConfig {

    @Bean
    public MyService myService() {
        return new MyService();
    }

    @Bean
    public MyRepository myRepository() {
        return new MyRepository();
    }
}
```
**Lợi ích**:

- Giúp cấu hình các bean trong Java code một cách rõ ràng và dễ kiểm soát.
- Thay thế cấu hình XML, mang lại sự linh hoạt và dễ bảo trì hơn.