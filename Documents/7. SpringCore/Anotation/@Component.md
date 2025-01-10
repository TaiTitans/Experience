
---

`@Component` là một annotation trong Spring Framework được sử dụng để đánh dấu một lớp Java như là một bean, cho phép Spring tự động phát hiện và quản lý nó trong Spring Application Context.

### 1. **Chức năng của `@Component`:**

- **Đánh dấu một lớp là Spring Bean**: Khi một lớp được đánh dấu bằng `@Component`, Spring sẽ quét và tự động đăng ký lớp này như là một bean trong container của Spring.
- **Component Scanning**: Spring sẽ tìm kiếm các lớp được đánh dấu bằng `@Component` (hoặc các annotation con của nó) trong các package được chỉ định và thêm chúng vào context của ứng dụng.
### 2. **Cách sử dụng `@Component`:**

**Ví dụ cơ bản:**
```java
import org.springframework.stereotype.Component;

@Component
public class MyComponent {
    public void performTask() {
        System.out.println("Task performed.");
    }
}
```
- `@Component`: Đánh dấu lớp `MyComponent` là một bean.
- Bean này sẽ được Spring quản lý và có thể được tự động inject vào các thành phần khác của ứng dụng.
### 3. **Lifecycle của `@Component`:**

- Khi ứng dụng Spring khởi động, Spring sẽ quét tất cả các package được chỉ định trong `@ComponentScan` để tìm các lớp đánh dấu bằng `@Component` hoặc các annotation con.
- Sau khi tìm thấy, Spring sẽ khởi tạo các bean này và quản lý chúng trong Application Context, cho phép chúng được inject vào các thành phần khác thông qua Dependency Injection.

### 4. **Lợi ích của `@Component`:**

- **Tự động phát hiện và quản lý bean**: Giúp giảm bớt sự cần thiết phải khai báo bean thủ công trong cấu hình.
- **Tái sử dụng mã**: Tạo và quản lý các bean dễ dàng, tái sử dụng trong nhiều phần của ứng dụng.
- **Tích hợp tốt với Dependency Injection**: Dễ dàng inject các bean khác vào nhau thông qua `@Autowired` hoặc constructor injection.

`@Component` là một phần không thể thiếu trong Spring Framework, giúp quản lý các bean một cách hiệu quả và tích hợp chặt chẽ với các cơ chế khác của Spring như Dependency Injection và AOP.