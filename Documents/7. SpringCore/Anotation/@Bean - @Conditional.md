
---
Dưới đây là giải thích chi tiết về các annotation `@Bean` và `@Conditional` trong Spring Framework:
### 1. **`@Bean`**

`@Bean` là một annotation trong Spring, được sử dụng để định nghĩa một bean trong một class cấu hình (configuration class).

- **Mục đích**:
    - Tạo và quản lý các bean trong ApplicationContext. Các bean này có thể được sử dụng ở các phần khác của ứng dụng thông qua Dependency Injection.
    - Giúp cấu hình các bean mà không cần sử dụng XML.
```java
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AppConfig {

    @Bean
    public MyService myService() {
        return new MyService();
    }
}
```

- **Giải thích**:
    
    - `@Configuration`: Đánh dấu `AppConfig` là một class cấu hình.
    - `@Bean`: Định nghĩa một bean có tên là `myService` và trả về một instance của `MyService`.
- **Lợi ích**:
    
    - Dễ dàng cấu hình và quản lý bean trong code Java.
    - Tách biệt rõ ràng cấu hình và logic nghiệp vụ.

---
### 2. **`@Conditional`**

`@Conditional` là một annotation trong Spring, được sử dụng để định nghĩa điều kiện mà theo đó một bean hoặc một cấu hình sẽ được kích hoạt (hoặc không).

- **Mục đích**:
    - Cho phép bạn điều khiển việc tạo bean dựa trên các điều kiện nhất định, chẳng hạn như môi trường, cấu hình hệ thống, hoặc các trạng thái cụ thể khác.

**Cách sử dụng `@Conditional`:**
```java
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Conditional;

@Configuration
public class AppConfig {

    @Bean
    @Conditional(MyCondition.class)
    public MyService myService() {
        return new MyService();
    }
}
```
Tạo một điều kiện tùy chỉnh `MyCondition`:
```java
import org.springframework.context.annotation.Condition;
import org.springframework.context.annotation.ConditionContext;
import org.springframework.core.type.AnnotatedTypeMetadata;

public class MyCondition implements Condition {

    @Override
    public boolean matches(ConditionContext context, AnnotatedTypeMetadata metadata) {
        // Kiểm tra điều kiện
        return context.getEnvironment().getProperty("my.feature.enabled", Boolean.class, false);
    }
}
```
- **Giải thích**:
    
    - `@Conditional(MyCondition.class)`: Chỉ tạo bean `myService` nếu điều kiện `MyCondition` trả về `true`.
    - `matches` method trong `MyCondition`: Kiểm tra xem thuộc tính `my.feature.enabled` có được kích hoạt hay không.
- **Lợi ích**:
    
    - Cung cấp sự linh hoạt cao trong việc tạo bean theo điều kiện.
    - Giúp tối ưu hóa ứng dụng bằng cách chỉ tạo các bean khi cần thiết.
### **Tóm tắt sự khác biệt và sử dụng:**

|Annotation|Mục đích|Sử dụng cho|
|---|---|---|
|`@Bean`|Định nghĩa một bean trong class cấu hình.|Tạo và quản lý các bean một cách rõ ràng.|
|`@Conditional`|Điều kiện hóa việc tạo bean hoặc cấu hình dựa trên trạng thái.|Kích hoạt bean hoặc cấu hình theo điều kiện cụ thể.|

Các annotation này giúp Spring quản lý các bean một cách linh hoạt và hiệu quả, đồng thời cho phép cấu hình động dựa trên các điều kiện cụ thể của ứng dụng.