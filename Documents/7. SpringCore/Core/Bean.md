
---

	Trong Spring Framework, **Bean** là các đối tượng được quản lý bởi IoC (Inversion of Control) Container. Những đối tượng này là phần cốt lõi của ứng dụng và có thể là bất kỳ đối tượng nào do người dùng định nghĩa. Bean được khởi tạo, lắp ráp và quản lý bởi Spring IoC Container.

### Các thành phần chính của Bean

1. **Configuration Metadata**: Thông tin cấu hình cho các bean. Metadata này có thể được định nghĩa bằng các tệp XML, các chú thích (annotations) trong mã nguồn Java, hoặc các lớp Java Configuration.
    
2. **Scope**: Định nghĩa phạm vi của bean, tức là cách bean được khởi tạo và chia sẻ trong ứng dụng. Các phạm vi phổ biến bao gồm:
    
    - **singleton** (mặc định): Chỉ có một instance của bean được tạo và chia sẻ trong toàn bộ Spring IoC Container.
    - **prototype**: Mỗi lần yêu cầu bean sẽ tạo ra một instance mới.
    - **request**: Một instance được tạo ra cho mỗi yêu cầu HTTP (chỉ dùng trong ứng dụng web).
    - **session**: Một instance được tạo ra cho mỗi phiên làm việc HTTP (chỉ dùng trong ứng dụng web).
    - **globalSession**: Một instance được tạo ra cho mỗi phiên làm việc toàn cục (hiếm khi được dùng, chỉ dùng trong ứng dụng portlet).
3. **Lifecycle Callbacks**: Các phương thức callback cho phép bean thực hiện các hành động trong các giai đoạn cụ thể của vòng đời của nó (khởi tạo và hủy).
    
4. **Dependencies**: Các phụ thuộc của bean, tức là các đối tượng khác mà bean cần để hoạt động. Các phụ thuộc này được Spring IoC Container tự động tiêm vào bean thông qua Dependency Injection.

---
# **Vòng đời của Bean**

Vòng đời của một Bean trong Spring gồm các giai đoạn:

1. **Tạo instance** (Instantiation)
2. **Tiêm phụ thuộc (Dependency Injection)**
3. **Gọi phương thức `@PostConstruct` hoặc `InitializingBean#afterPropertiesSet()`**
4. **Bean sẵn sàng để sử dụng**
5. **Gọi phương thức `@PreDestroy` hoặc `DisposableBean#destroy()` khi Bean bị hủy**

# **Cơ chế Lazy Initialization**

Mặc định, Spring sẽ khởi tạo tất cả các Bean ngay khi ứng dụng khởi động (**Eager Initialization**). Tuy nhiên, bạn có thể trì hoãn việc khởi tạo bằng cách dùng `@Lazy`.
```java
@Component
@Lazy
public class LazyBean {
    public LazyBean() {
        System.out.println("LazyBean created!");
    }
}
```
Spring sẽ chỉ khởi tạo Bean này khi nó thực sự được sử dụng.

# **Bean Proxying và AOP**

Spring sử dụng **CGLIB** hoặc **JDK Dynamic Proxy** để tạo Bean proxy cho các tính năng như **Transactional**, **Security**, và **AOP**.

Ví dụ, nếu bạn sử dụng `@Transactional`, Spring sẽ tạo một proxy của Bean để xử lý transaction:
```java
@Service
@Transactional
public class MyService {
    public void doWork() {
        System.out.println("Doing work...");
    }
}
```
Spring sẽ thay thế `MyService` bằng một proxy để quản lý giao dịch tự động.

# **Tóm tắt**

- **Bean là các object được Spring IoC Container quản lý**.
- **Có nhiều cách để định nghĩa Bean**: XML, Annotation (`@Component`), hoặc `@Bean` trong `@Configuration`.
- **Spring hỗ trợ nhiều phạm vi Bean**: `singleton`, `prototype`, `request`, `session`, v.v.
- **Bean có vòng đời gồm các giai đoạn: Khởi tạo → Tiêm phụ thuộc → Sử dụng → Hủy**.
- **Hỗ trợ Lazy Initialization để trì hoãn việc tạo Bean**.
- **Spring sử dụng Proxying để hỗ trợ AOP và Transaction Management**.