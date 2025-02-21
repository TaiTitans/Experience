
---
Trong Spring, **Bean Proxying** là một cơ chế sử dụng Proxy (có thể là **JDK Dynamic Proxy** hoặc **CGLIB**) để thay thế Bean gốc bằng một Bean được wrap lại, nhằm thêm các tính năng bổ sung như **AOP, Transaction Management, Security, Lazy Initialization, v.v.**

## **1. Vì sao Spring sử dụng Bean Proxying?**

Spring sử dụng Bean Proxying để:

- **Hỗ trợ AOP (Aspect-Oriented Programming):** Giúp chèn logic trước, sau hoặc xung quanh method mà không cần sửa code gốc.
- **Hỗ trợ @Transactional:** Giúp Spring quản lý giao dịch tự động.
- **Tạo Singleton Proxy:** Cho phép Bean singleton có thể chứa dependencies có scope khác nhau (ví dụ: Prototype).
- **Hỗ trợ Lazy Initialization:** Trì hoãn khởi tạo Bean đến khi nó được sử dụng lần đầu tiên.

## **2. Các Loại Proxy trong Spring**

Spring có hai cách để tạo Proxy:

1. **JDK Dynamic Proxy (Interface-based Proxying)**
2. **CGLIB Proxy (Class-based Proxying)**

### **2.1. JDK Dynamic Proxy**

- Chỉ hoạt động khi Bean implement **interface**.
- Sử dụng `java.lang.reflect.Proxy`.
- Tạo Proxy chỉ cho các **method trong interface**, không áp dụng cho method của class cụ thể.

Ví dụ:
```java
public interface MyService {
    void doSomething();
}

@Service
@Transactional
public class MyServiceImpl implements MyService {
    public void doSomething() {
        System.out.println("Doing something...");
    }
}
```

- Khi Spring thấy `@Transactional`, nó tạo một Proxy cho `MyServiceImpl` bằng **JDK Dynamic Proxy**.
- Khi bạn gọi `doSomething()`, thực chất là gọi thông qua Proxy, chứ không phải trực tiếp vào `MyServiceImpl`.

### **2.2. CGLIB Proxy**

- Được sử dụng khi Bean **không implement interface** hoặc bạn cấu hình Spring để dùng CGLIB.
- Sử dụng **subclassing** thay vì **interface**.
- Dựa trên thư viện **CGLIB (Code Generation Library)** để tạo subclass của Bean và override method.
- Hoạt động với cả **interface method và non-interface method**.

Ví dụ:
```java
@Service
@Transactional
public class MyService {
    public void doSomething() {
        System.out.println("Doing something...");
    }
}
```
Vì `MyService` không implement interface nào, Spring sẽ dùng **CGLIB Proxy** để tạo subclass của `MyService` và override `doSomething()`.

💡 **Lưu ý:** CGLIB Proxy không hoạt động với các **final methods** vì không thể override chúng.


## **Cấu Hình Spring Để Dùng Proxy**

Mặc định, Spring chọn Proxy phù hợp dựa vào cấu trúc Bean. Tuy nhiên, bạn có thể ép buộc Spring sử dụng kiểu Proxy mong muốn.

## **Proxying trong Một Số Tính Năng Quan Trọng của Spring**

### **AOP (Aspect-Oriented Programming)**

Spring AOP sử dụng Proxy để chèn logic vào trước/sau method.

### **Transaction Management (@Transactional)**

Spring sử dụng Proxy để quấn các method có `@Transactional`, giúp tự động bắt đầu và commit transaction.

### **Lazy Initialization**

Spring dùng Proxy để trì hoãn khởi tạo Bean cho đến khi nó thực sự được sử dụng.

## **Hạn Chế của Proxy trong Spring**

1. **Không thể proxy method `final`** → Vì Proxy cần override method để can thiệp logic.
2. **Không proxy được `private` method** → Proxy chỉ hoạt động với `public` và `protected` method.
3. **JDK Proxy chỉ hoạt động với Interface** → Nếu không có Interface, phải dùng CGLIB.
4. **Self-invocation không hoạt động với Proxy** → Khi một method trong cùng một class gọi method khác trong chính nó, Spring không thể áp dụng Proxy.