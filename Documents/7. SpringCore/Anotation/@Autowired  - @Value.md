
---
Dưới đây là giải thích chi tiết về các annotation `@Autowired` và `@Value` trong Spring Framework:

### 1. **`@Autowired`**

`@Autowired` là một annotation trong Spring được sử dụng để tự động tiêm (inject) các dependency vào các bean khác.

- **Mục đích**:
    - Giúp Spring tự động tìm và tiêm các bean cần thiết vào một đối tượng mà không cần cấu hình thủ công.
    - Hỗ trợ tiêm các dependency theo kiểu constructor, setter, hoặc field.

**Cách sử dụng**:
```java
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class MyService {

    private final MyRepository myRepository;

    @Autowired
    public MyService(MyRepository myRepository) {
        this.myRepository = myRepository;
    }

    // Hoặc sử dụng field injection
    // @Autowired
    // private MyRepository myRepository;

    // Hoặc sử dụng setter injection
    // @Autowired
    // public void setMyRepository(MyRepository myRepository) {
    //     this.myRepository = myRepository;
    // }
}
```
- **Giải thích**:
    
    - `@Autowired` trên constructor: Spring tự động tìm và tiêm `MyRepository` vào `MyService`.
    - `@Autowired` trên field hoặc setter: Cung cấp các cách khác để tiêm dependency.
- **Lợi ích**:
    
    - Giảm thiểu cấu hình thủ công, tự động quản lý dependency.
    - Hỗ trợ tính linh hoạt trong việc lựa chọn cách tiêm dependency.

### 2. **`@Value`**

`@Value` là một annotation trong Spring được sử dụng để tiêm giá trị từ file cấu hình (`application.properties` hoặc `application.yml`) hoặc từ một biểu thức khác vào các bean.

- **Mục đích**:
    - Giúp tiêm giá trị từ các nguồn bên ngoài vào các thuộc tính của bean.
    - Hỗ trợ việc đọc cấu hình dễ dàng từ file cấu hình hoặc các biến môi trường.

**Cách sử dụng**:
```java
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class MyComponent {

    @Value("${my.property}")
    private String myProperty;

    @Value("${my.number:42}")
    private int myNumber;

    @Value("#{systemProperties['user.name']}")
    private String systemUserName;

    public void printProperties() {
        System.out.println("myProperty: " + myProperty);
        System.out.println("myNumber: " + myNumber);
        System.out.println("System User Name: " + systemUserName);
    }
}
```
- **Giải thích**:
    
    - `@Value("${my.property}")`: Tiêm giá trị của `my.property` từ `application.properties`.
    - `@Value("${my.number:42}")`: Tiêm giá trị của `my.number`, nếu không có thì sử dụng giá trị mặc định là `42`.
    - `@Value("#{systemProperties['user.name']}")`: Tiêm giá trị từ thuộc tính hệ thống `user.name`.
- **Lợi ích**:
    
    - Dễ dàng quản lý và tiêm giá trị từ các nguồn bên ngoài.
    - Hỗ trợ tiêm các giá trị mặc định và biểu thức phức tạp.