
---
## **1. Transaction Isolation Levels**

Trong Spring Data JPA, chúng ta có thể kiểm soát mức độ cách ly của các giao dịch bằng cách sử dụng các **Transaction Isolation Levels**.

### **🔹 Các Mức Cách Ly trong Transaction**

|Isolation Level|Dirty Read|Non-repeatable Read|Phantom Read|
|---|---|---|---|
|**READ UNCOMMITTED**|✅ Có|✅ Có|✅ Có|
|**READ COMMITTED**|❌ Không|✅ Có|✅ Có|
|**REPEATABLE READ**|❌ Không|❌ Không|✅ Có|
|**SERIALIZABLE**|❌ Không|❌ Không|❌ Không|

### **🔹 Vấn Đề Cần Giải Quyết**

1. **Dirty Read** 🟠
    
    - Một transaction có thể đọc dữ liệu chưa commit từ transaction khác.
    - Có thể dẫn đến trạng thái không nhất quán khi transaction bị rollback.
    - Giải pháp: **READ COMMITTED** trở lên.
2. **Non-repeatable Read** 🟡
    
    - Một transaction đọc cùng một bản ghi nhiều lần nhưng nhận được các giá trị khác nhau do transaction khác đã thay đổi dữ liệu.
    - Giải pháp: **REPEATABLE READ** trở lên.
3. **Phantom Read** 🔴
    
    - Một transaction thực hiện cùng một truy vấn nhiều lần nhưng nhận được tập kết quả khác nhau do transaction khác đã **thêm/xóa** dữ liệu.
    - Giải pháp: **SERIALIZABLE**.


### **🔹 Cấu hình Isolation Level trong Spring Boot**

Spring Data JPA hỗ trợ việc thiết lập mức cách ly bằng `@Transactional(isolation = Isolation.LEVEL)`.

Ví dụ: **Chỉ đọc dữ liệu đã được commit (READ COMMITTED)**:
```java
@Transactional(isolation = Isolation.READ_COMMITTED)
public void processOrder(Long orderId) {
    // Code xử lý đơn hàng
}
```

## **2. Pessimistic vs Optimistic Locking**

Locking giúp **tránh xung đột dữ liệu** trong các ứng dụng nhiều luồng.

### **🔹 Optimistic Locking (Khóa Lạc Quan)**

- Sử dụng `@Version` để kiểm soát xung đột mà không cần khóa cấp database.
- Nếu dữ liệu bị sửa đổi bởi transaction khác, JPA sẽ **ném ngoại lệ** (`OptimisticLockException`).
- **Ưu điểm**: Không khóa dữ liệu, hiệu suất tốt hơn.
- **Nhược điểm**: Nếu có nhiều cập nhật đồng thời, có thể gây lỗi lặp.

**🔹 Cách triển khai Optimistic Locking với `@Version`**
```java
@Entity
public class Product {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String name;
    private Double price;

    @Version
    private Integer version; // Dùng để kiểm soát xung đột
}
```
- Khi một bản ghi được đọc, **`version`** sẽ được lưu lại.
- Khi cập nhật bản ghi, JPA kiểm tra xem `version` có thay đổi không.
- Nếu có thay đổi từ transaction khác, JPA sẽ **ném ngoại lệ**.

🔹 Cách xử lý ngoại lệ Optimistic Locking
```java
try {
    productService.updateProductPrice(productId, newPrice);
} catch (OptimisticLockException e) {
    // Xử lý khi có xung đột cập nhật
    System.out.println("Conflict detected, retrying...");
}
```

### **🔹 Pessimistic Locking (Khóa Bi Quan)**

- Dùng khi dữ liệu **được đọc và ghi thường xuyên**.
- Có 2 loại khóa chính:
    - `PESSIMISTIC_READ`: Chặn ghi nhưng vẫn cho phép đọc.
    - `PESSIMISTIC_WRITE`: Chặn cả đọc và ghi từ các transaction khác.

**🔹 Cách dùng Pessimistic Locking trong Spring Data JPA**
```java
@Lock(LockModeType.PESSIMISTIC_WRITE)
@Query("SELECT p FROM Product p WHERE p.id = :id")
Optional<Product> findByIdForUpdate(@Param("id") Long id);
```

- Khi một transaction lấy dữ liệu, nó sẽ **chặn** các transaction khác **đọc và ghi** vào bản ghi đó.

**🔹 Khi nào nên dùng?** ✅ **Optimistic Locking** → Khi **xung đột hiếm xảy ra**.  
✅ **Pessimistic Locking** → Khi **có nhiều cập nhật đồng thời** vào cùng dữ liệu.

## **3. Deadlock & Database Locking Strategies**

### **🔹 Deadlock là gì?**

- Deadlock xảy ra khi **hai transaction chờ nhau** để giải phóng khóa và không thể tiếp tục.
- VD:
    - Transaction A khóa **Bản ghi 1**, cần **Bản ghi 2**.
    - Transaction B khóa **Bản ghi 2**, cần **Bản ghi 1**.
    - Cả hai bị kẹt 🛑.

### **🔹 Cách tránh Deadlock**

1. **Duy trì thứ tự truy cập dữ liệu**
    
    - Luôn truy vấn dữ liệu theo **cùng một thứ tự** để tránh deadlock.
2. **Giới hạn thời gian khóa**
    
    - Dùng `lock timeout` để tránh chờ vô hạn:
- `@Transactional(timeout = 5)`
**Sử dụng Retry Mechanism**

- Khi bị deadlock, **thử lại giao dịch** sau vài mili giây.

### **Advisory Locks trong PostgreSQL**

- **Advisory Lock** là một dạng khóa mà ứng dụng tự quản lý thay vì database tự động khóa bản ghi.
- Dùng khi **cần kiểm soát truy cập mà không cần khóa row-level**.

```sql
SELECT pg_advisory_lock(12345);  -- Giữ khóa
SELECT pg_advisory_unlock(12345); -- Bỏ khóa
```

Sử dụng cho các tác vụ mà **không muốn JPA tự động khóa dữ liệu**.


## **🚀 Tổng Kết**

|Tính năng|Khi nào dùng?|
|---|---|
|**READ COMMITTED**|Ngăn **dirty read** nhưng cho phép **non-repeatable read**|
|**REPEATABLE READ**|Ngăn **dirty read** và **non-repeatable read** nhưng có **phantom read**|
|**SERIALIZABLE**|Chặn mọi xung đột nhưng có thể gây chậm|
|**Optimistic Locking (`@Version`)**|Khi **xung đột hiếm**, ít ảnh hưởng hiệu suất|
|**Pessimistic Locking (`@Lock`)**|Khi có nhiều transaction ghi đồng thời|
|**Advisory Locks**|Khi muốn kiểm soát thủ công mà không khóa database|