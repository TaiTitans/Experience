
---
Dưới đây là giải thích chi tiết về các annotation `@Lazy`, `@Primary`, `@Qualifier`, và `@DependsOn` trong Spring Framework:

### 1. **`@Lazy`**

`@Lazy` là một annotation được sử dụng để trì hoãn việc khởi tạo bean cho đến khi nó được thực sự cần thiết.

- **Mục đích**:
    - Giúp cải thiện thời gian khởi động ứng dụng bằng cách chỉ tạo các bean khi chúng được sử dụng.
    - Tiết kiệm tài nguyên bằng cách tránh khởi tạo các bean không cần thiết ngay lập tức.

**Cách sử dụng**:
```java
import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;

@Component
@Lazy
public class MyLazyBean {
    public MyLazyBean() {
        System.out.println("MyLazyBean is initialized");
    }
}

```
- **Giải thích**:
    
    - `@Lazy`: Bean `MyLazyBean` chỉ được khởi tạo khi nó thực sự được gọi, không phải khi ứng dụng khởi động.
- **Lợi ích**:
    
    - Giúp cải thiện hiệu suất khởi động ứng dụng.
    - Hữu ích trong các ứng dụng lớn hoặc khi cần tối ưu hóa tài nguyên.
### 2. **`@Primary`**

`@Primary` được sử dụng để đánh dấu một bean là ưu tiên khi có nhiều bean cùng loại được định nghĩa.

- **Mục đích**:
    - Định rõ bean nào sẽ được sử dụng mặc định khi có nhiều bean cùng kiểu hoặc cùng loại trong ApplicationContext.

**Cách sử dụng**:
```java
import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Component;

@Component
@Primary
public class PrimaryBean implements MyInterface {
    // Implementation
}

@Component
public class SecondaryBean implements MyInterface {
    // Implementation
}
```
- **Giải thích**:
    
    - `@Primary`: Khi Spring cần một bean của `MyInterface`, `PrimaryBean` sẽ được sử dụng mặc định.
- **Lợi ích**:
    
    - Tránh phải sử dụng `@Qualifier` hoặc các phương pháp khác để chỉ định rõ ràng bean nào sẽ được sử dụng
### 3. **`@Qualifier`**

`@Qualifier` được sử dụng để chỉ định rõ ràng bean nào sẽ được tiêm khi có nhiều bean cùng loại.

- **Mục đích**:
    - Giải quyết vấn đề khi có nhiều bean cùng loại trong ApplicationContext và bạn cần xác định rõ bean nào sẽ được sử dụng.

**Cách sử dụng**:
```java
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

@Service
public class MyService {

    private final MyInterface myInterface;

    @Autowired
    public MyService(@Qualifier("secondaryBean") MyInterface myInterface) {
        this.myInterface = myInterface;
    }
}
```
- **Giải thích**:
    
    - `@Qualifier("secondaryBean")`: Xác định rõ ràng rằng bean `secondaryBean` sẽ được tiêm vào `myInterface`.
- **Lợi ích**:
    
    - Giải quyết xung đột giữa nhiều bean cùng loại.
    - Cung cấp sự linh hoạt trong việc lựa chọn bean cụ thể.

### 4. **`@DependsOn`**

`@DependsOn` được sử dụng để chỉ định rằng một bean phụ thuộc vào một hoặc nhiều bean khác và các bean đó cần được khởi tạo trước.

- **Mục đích**:
    - Đảm bảo thứ tự khởi tạo các bean khi có sự phụ thuộc rõ ràng giữa chúng.

**Cách sử dụng**:
```java
import org.springframework.context.annotation.DependsOn;
import org.springframework.stereotype.Component;

@Component
@DependsOn({"beanA", "beanB"})
public class MyBean {
    public MyBean() {
        System.out.println("MyBean is initialized after BeanA and BeanB");
    }
}
```
- **Giải thích**:
    
    - `@DependsOn({"beanA", "beanB"})`: Đảm bảo rằng `beanA` và `beanB` được khởi tạo trước khi `MyBean` được khởi tạo.
- **Lợi ích**:
    
    - Quản lý thứ tự khởi tạo của các bean khi có các phụ thuộc phức tạp.
    - Giảm thiểu các lỗi tiềm ẩn do khởi tạo bean không đúng thứ tự.
### **Tóm tắt sự khác biệt và sử dụng:**

| Annotation   | Mục đích                                                            | Sử dụng cho                                           |
| ------------ | ------------------------------------------------------------------- | ----------------------------------------------------- |
| `@Lazy`      | Trì hoãn khởi tạo bean cho đến khi cần thiết.                       | Tối ưu hóa thời gian khởi động và tài nguyên.         |
| `@Primary`   | Đánh dấu bean ưu tiên khi có nhiều bean cùng loại.                  | Định rõ bean mặc định để sử dụng.                     |
| `@Qualifier` | Chỉ định rõ ràng bean nào sẽ được tiêm khi có nhiều bean cùng loại. | Giải quyết xung đột khi có nhiều bean cùng loại.      |
| `@DependsOn` | Định rõ thứ tự khởi tạo giữa các bean có phụ thuộc nhau.            | Đảm bảo khởi tạo đúng thứ tự giữa các bean phụ thuộc. |