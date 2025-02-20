
---
**Setter Injection** là một phương pháp **Dependency Injection (DI)** trong Spring, trong đó dependencies được inject vào Bean thông qua phương thức setter (`setXxx`).

Mặc dù **Constructor Injection** được ưu tiên trong hầu hết các trường hợp, nhưng **Setter Injection** vẫn hữu ích khi dependency **không bắt buộc** hoặc **có thể thay đổi sau khi Bean được tạo**.


## **1. Vì Sao Nên Dùng Setter Injection?**

### ✅ **Ưu điểm**

- **Linh hoạt**: Có thể thay đổi dependency sau khi Bean đã được tạo.
- **Không gây lỗi Circular Dependency**: Tránh vòng lặp khi hai Bean phụ thuộc lẫn nhau.
- **Dễ sử dụng với các framework như Spring MVC** (các controllers có thể sử dụng Setter Injection để nhận Bean).

### ❌ **Nhược điểm**

- **Không đảm bảo bất biến (Immutability)**: Các dependencies có thể bị thay đổi sau khi Bean được tạo.
- **Không đảm bảo dependency luôn tồn tại**: Nếu quên gọi setter, dependency có thể bị `null`.
- **Khó test hơn Constructor Injection** vì cần gọi setter trong Unit Test.

## **2. Cách Dùng Setter Injection trong Spring**

Có thể thực hiện Setter Injection bằng **Annotations**, **Java Configuration**, hoặc **XML Configuration**.

### **2.1. Setter Injection với `@Autowired`**
🟢 **Lưu ý:** Nếu `UserRepository` chưa được inject, `userRepository` có thể bị `null`.

### **2.2. Setter Injection với `@Bean` trong Java Configuration**

Nếu không dùng `@Component`, có thể khai báo Bean bằng `@Configuration` và inject bằng setter.

🟢 **Ưu điểm**: Linh hoạt, có thể cấu hình dependency lúc runtime.


### **2.3. Setter Injection với XML Configuration (Cũ)
🛑 **Không còn phổ biến**, nhưng vẫn hữu ích trong các dự án cũ.

## **3. Khi Nào Nên Dùng Setter Injection?**

### ✅ Dùng khi:

- **Dependency không bắt buộc** (ví dụ: Logger, Audit Service).
- **Cần thay đổi dependency sau khi Bean được khởi tạo**.
- **Tránh Circular Dependency** khi Constructor Injection gây vòng lặp.

### ❌ Không dùng khi:

- Dependency **luôn bắt buộc** (nên dùng Constructor Injection).
- Cần đảm bảo **immutability** và dễ kiểm thử.

## **4. So Sánh Constructor Injection vs Setter Injection**

|Feature|**Constructor Injection**|**Setter Injection**|
|---|---|---|
|**Khả năng kiểm thử (Testability)**|✅ Dễ test với Mockito|❌ Cần gọi setter trong test|
|**Bất biến (Immutability)**|✅ Bất biến|❌ Có thể thay đổi|
|**Dễ phát hiện lỗi thiếu dependency**|✅ Phát hiện ngay khi chạy ứng dụng|❌ Có thể bị `null` nếu quên gọi setter|
|**Circular Dependency**|❌ Có thể gây lỗi vòng lặp|✅ Không gây lỗi vòng lặp|
|**Khả năng mở rộng**|✅ Tốt cho class lớn|✅ Linh hoạt|

✅ **Tóm lại:**

- Dùng **Constructor Injection** cho **dependencies bắt buộc**.
- Dùng **Setter Injection** khi **dependency có thể thay đổi hoặc không bắt buộc**.

## **6. Kết Luận**

🔹 **Setter Injection linh hoạt hơn nhưng không đảm bảo dependency luôn tồn tại**.  
🔹 **Dùng Setter Injection khi dependency có thể thay đổi hoặc không bắt buộc**.  
🔹 **Constructor Injection vẫn là lựa chọn tốt hơn trong hầu hết các trường hợp**.
