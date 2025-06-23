
---
## **1. Spring Data JPA Auditing**

Spring Data JPA Auditing giúp **tự động theo dõi các thay đổi trên Entity**, như **người tạo, người sửa, ngày tạo, ngày sửa**.

### **🔹 Kích hoạt Spring Auditing**

Để sử dụng Spring Auditing, bạn cần **kích hoạt** bằng `@EnableJpaAuditing`:
```java
@Configuration
@EnableJpaAuditing
public class JpaConfig {
}
```
### **🔹 Các Annotation Auditing**

Spring hỗ trợ 4 annotation chính:

| Annotation          | Mô tả                              |
| ------------------- | ---------------------------------- |
| `@CreatedBy`        | Lưu lại người tạo record.          |
| `@LastModifiedBy`   | Lưu lại người sửa record gần nhất. |
| `@CreatedDate`      | Lưu ngày tạo record.               |
| `@LastModifiedDate` | Lưu ngày cập nhật gần nhất.        |
🔹 Ví dụ: Entity có Audit
```java
import org.springframework.data.annotation.CreatedBy;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedBy;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
@EntityListeners(AuditingEntityListener.class)
public class Product {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String name;

    @CreatedBy
    private String createdBy;

    @LastModifiedBy
    private String lastModifiedBy;

    @CreatedDate
    private LocalDateTime createdDate;

    @LastModifiedDate
    private LocalDateTime lastModifiedDate;
}
```

### **🔹 Cấu hình `AuditorAware`**

`AuditorAware` giúp lấy thông tin người dùng hiện tại để điền vào `@CreatedBy` & `@LastModifiedBy`.
```java
import org.springframework.data.domain.AuditorAware;
import org.springframework.stereotype.Component;

import java.util.Optional;

@Component
public class AuditorAwareImpl implements AuditorAware<String> {
    @Override
    public Optional<String> getCurrentAuditor() {
        // Lấy username từ Spring Security hoặc bất kỳ nguồn nào khác
        return Optional.of("admin"); // Ví dụ
    }
}
```

🔹 **Khai báo `AuditorAware` trong cấu hình:**
```java
@Configuration
@EnableJpaAuditing(auditorAwareRef = "auditorAwareImpl")
public class JpaConfig {
}
```

## **2. Hibernate Interceptor vs. Spring Auditing**

|Tiêu chí|Hibernate Interceptor|Spring Auditing|
|---|---|---|
|**Mức độ hoạt động**|Toàn bộ Hibernate|Chỉ áp dụng cho Entity|
|**Lắng nghe sự kiện**|Trước/sau INSERT, UPDATE, DELETE|Chỉ hỗ trợ CREATED, MODIFIED|
|**Hiệu suất**|Nặng hơn|Nhẹ hơn|
|**Dùng khi nào?**|Khi cần can thiệp sâu vào Hibernate|Khi chỉ cần tracking audit|

🔹 **Khi nào dùng Hibernate Interceptor?**

- Khi cần **thay đổi dữ liệu toàn cục** trước khi Hibernate thực hiện INSERT/UPDATE.
- Khi cần **ghi log tất cả thao tác trên DB**, không chỉ các trường `@CreatedBy`, `@LastModifiedDate`.
## **3. Application Events & Entity Listeners**

Spring hỗ trợ **2 cách để lắng nghe sự kiện thay đổi Entity**:

1. **Spring Events (`ApplicationEventPublisher`)** → Lắng nghe **sự kiện toàn ứng dụng**.
2. **JPA Entity Listeners (`@EntityListeners`)** → Lắng nghe **thay đổi của từng Entity**.
### **🔹 Spring Events (`ApplicationEventPublisher`)**

Spring Events giúp **phát và lắng nghe sự kiện** mà không bị ràng buộc vào Entity cụ thể.

**🔹 Định nghĩa sự kiện**
```java
public class ProductCreatedEvent {
    private final Product product;

    public ProductCreatedEvent(Product product) {
        this.product = product;
    }

    public Product getProduct() {
        return product;
    }
}
```
🔹 Phát sự kiện khi tạo Product
```java
import org.springframework.context.ApplicationEventPublisher;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class ProductService {
    private final ProductRepository productRepository;
    private final ApplicationEventPublisher eventPublisher;

    public ProductService(ProductRepository productRepository, ApplicationEventPublisher eventPublisher) {
        this.productRepository = productRepository;
        this.eventPublisher = eventPublisher;
    }

    @Transactional
    public Product createProduct(Product product) {
        Product savedProduct = productRepository.save(product);
        eventPublisher.publishEvent(new ProductCreatedEvent(savedProduct));
        return savedProduct;
    }
}
```
🔹 Lắng nghe sự kiện
```java
import org.springframework.context.event.EventListener;
import org.springframework.stereotype.Component;

@Component
public class ProductEventListener {
    @EventListener
    public void handleProductCreated(ProductCreatedEvent event) {
        System.out.println("New product created: " + event.getProduct().getName());
    }
}
```

### **🔹 Entity Listeners (`@EntityListeners`)**

Entity Listener gắn trực tiếp vào Entity và chạy trước/sau các sự kiện của Hibernate.

**🔹 Định nghĩa Listener**
```java
import jakarta.persistence.*;

public class ProductEntityListener {

    @PrePersist
    public void prePersist(Product product) {
        System.out.println("Before saving product: " + product.getName());
    }

    @PostPersist
    public void postPersist(Product product) {
        System.out.println("After saving product: " + product.getName());
    }

    @PreUpdate
    public void preUpdate(Product product) {
        System.out.println("Before updating product: " + product.getName());
    }
}
```

🔹 Áp dụng Listener vào Entity
```java
@Entity
@EntityListeners(ProductEntityListener.class)
public class Product {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    private String name;
}
```

## **4. So sánh @EntityListeners và Spring Events**

| Tiêu chí                | `@EntityListeners`                          | Spring Events (`ApplicationEventPublisher`) |
| ----------------------- | ------------------------------------------- | ------------------------------------------- |
| **Mức độ hoạt động**    | Chỉ với Entity cụ thể                       | Toàn bộ ứng dụng                            |
| **Hỗ trợ sự kiện nào?** | `@PrePersist`, `@PostPersist`, `@PreUpdate` | Tùy chỉnh sự kiện                           |
| **Hiệu suất**           | Cao hơn vì gắn với Hibernate                | Thấp hơn vì chạy trong Spring Context       |
| **Khi nào nên dùng?**   | Khi cần kiểm soát logic ngay trên Entity    | Khi cần xử lý tác vụ khác như gửi email     |
