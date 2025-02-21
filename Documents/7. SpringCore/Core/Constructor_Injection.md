

---
**Constructor Injection** là một phương pháp **Dependency Injection (DI)** trong Spring, trong đó các dependencies được inject vào Bean thông qua **constructor**. Đây là cách **ưu tiên** trong Spring vì giúp đảm bảo **tính bất biến (immutability)** của Bean và **dễ dàng kiểm thử (testability)**.

## **1. Vì Sao Nên Dùng Constructor Injection?**

### ✅ **Ưu điểm**

- **Bất biến**: Các dependencies chỉ được set một lần khi khởi tạo object, không thể thay đổi sau đó.
- **Dễ dàng kiểm thử (Unit Testing)**: Không cần Spring Context, có thể dùng **Mockito** để mock dependencies.
- **Giảm phụ thuộc vào Spring**: Không cần dùng `@Autowired` (Spring Boot từ v4.3 trở lên có thể bỏ `@Autowired` nếu chỉ có một constructor).
- **Dễ phát hiện lỗi thiếu dependency**: Nếu quên inject dependency, Spring sẽ báo lỗi **ngay khi ứng dụng khởi động** thay vì lỗi runtime.

### ❌ **Nhược điểm**

- Khi một class có quá nhiều dependencies, constructor sẽ dài, gây khó đọc.
- Có thể khó sử dụng khi dependency có nhiều constructor (phải chọn constructor phù hợp).

## **2. Cách Dùng Constructor Injection trong Spring**

Có thể thực hiện Constructor Injection bằng **Java Configuration**, **Annotations**, hoặc **XML Configuration**.

Constructor Injection với Annotation (`@Autowired`)
Constructor Injection với `@Bean` trong Java Configuration
Constructor Injection với Nhiều Dependency 
Khi có nhiều Bean cùng loại, dùng `@Qualifier` để chỉ định chính xác Bean cần inject.

## **So Sánh Constructor Injection vs Field Injection vs Setter Injection**

|Feature|**Constructor Injection**|**Field Injection (@Autowired trên biến)**|**Setter Injection**|
|---|---|---|---|
|**Khả năng kiểm thử (Testability)**|✅ Dễ test với Mockito|❌ Cần Reflection hoặc Spring Context|✅ Dễ test|
|**Bất biến (Immutability)**|✅ Giữ nguyên dependencies|❌ Có thể thay đổi runtime|❌ Có thể thay đổi runtime|
|**Dễ phát hiện lỗi thiếu dependency**|✅ Lỗi tại thời điểm khởi động|❌ Có thể gây lỗi runtime|❌ Có thể gây lỗi runtime|
|**Khả năng mở rộng**|✅ Dễ dàng mở rộng với constructor chaining|❌ Không tốt khi mở rộng|✅ Linh hoạt|
|**Sử dụng trong Spring Boot**|🔥 Mặc định được ưu tiên|🟡 Có thể dùng nhưng không khuyến nghị|🟡 Có thể dùng nhưng không khuyến nghị|

✅ **Tóm lại:**

- **Constructor Injection** là **cách tốt nhất** khi **dependencies là bắt buộc**.
- **Field Injection** có thể dùng trong trường hợp **đơn giản hoặc với @Lazy**.
- **Setter Injection** thích hợp khi dependency **có thể thay đổi sau khi tạo Bean**.

## **Constructor Injection và Circular Dependency**

Một vấn đề thường gặp khi dùng Constructor Injection là **Circular Dependency (vòng lặp phụ thuộc)**.
```java
@Component
public class A {
    private final B b;
    
    @Autowired
    public A(B b) {
        this.b = b;
    }
}
```

```java
@Component
public class B {
    private final A a;

    @Autowired
    public B(A a) {
        this.a = a;
    }
}
```

⛔ **Lỗi: Bean A phụ thuộc vào B, và B lại phụ thuộc vào A ⇒ Vòng lặp!**

### **Cách giải quyết Circular Dependency**

1. **Dùng `@Lazy` trên một trong các Bean**
2. Dùng Setter Injection thay vì Constructor Injection

## **Kết Luận**

🔹 **Constructor Injection là cách tốt nhất để quản lý dependencies trong Spring**.  
🔹 **Dễ dàng kiểm thử, bảo đảm tính bất biến, và phát hiện lỗi thiếu dependency sớm**.  
🔹 **Có thể gây Circular Dependency, cần xử lý bằng `@Lazy` hoặc Setter Injection**.