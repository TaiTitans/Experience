
---
Trong JPA, mỗi entity có một vòng đời được quản lý bởi **EntityManager** thông qua **Persistence Context**. Hiểu rõ về **lifecycle của entity** sẽ giúp bạn tối ưu hóa hiệu suất, tránh lỗi liên quan đến transaction và cache.

## **2.1. Các Trạng Thái Của Entity**

Một entity trong JPA có **4 trạng thái chính**:

1. **Transient (Tạm thời - Chưa Quản Lý)**
2. **Managed (Được Quản Lý)**
3. **Detached (Bị Tách Rời)**
4. **Removed (Bị Xóa)**

📌 **Entity có thể di chuyển giữa các trạng thái dựa trên các thao tác CRUD và Persistence Context.**

### **Sơ đồ trạng thái Entity trong JPA**
```
+------------+   persist()    +-----------+
|  Transient | -------------> |  Managed  |
+------------+               +-----------+
       |                           |
       | remove()                  | detach()
       |                           |
       v                           v
  +---------+                 +---------+
  | Removed |                 | Detached|
  +---------+                 +---------+
```
## **2.2. Giải Thích Chi Tiết Từng Trạng Thái**

### **1️⃣ Transient (Tạm thời - Chưa được quản lý)**

🔹 Khi một entity mới được tạo, nhưng chưa được lưu vào database, nó thuộc trạng thái **Transient**.  
🔹 JPA không quản lý entity này, nên không thể thực hiện **rollback, update** hoặc **persist** tự động.

📌 **Ví dụ:**
```java
User user = new User();
user.setUsername("john_doe");
user.setPassword("password123");
// user hiện đang ở trạng thái Transient vì chưa được quản lý bởi EntityManager.
```

👉 **Cách đưa entity từ Transient -> Managed:**

- Sử dụng **persist()** hoặc **save()** (nếu dùng Spring Data JPA).

```java
@EntityManager.persist(user);
```
🔥 **Sau khi gọi persist(), entity sẽ được quản lý bởi Persistence Context và chuyển sang trạng thái Managed.**

### **2️⃣ Managed (Được Quản Lý bởi Persistence Context)**

🔹 Khi một entity được lưu vào database hoặc truy vấn từ database, nó thuộc trạng thái **Managed**.  
🔹 Các thay đổi trên entity sẽ tự động được cập nhật vào database mà không cần gọi `save()`.

📌 **Ví dụ:**
```java
User user = entityManager.find(User.class, 1L); // user được lấy từ DB
user.setPassword("newpassword123"); // Cập nhật password
// Không cần gọi save(), JPA sẽ tự động update khi commit transaction
```

📌 **Persistence Context sẽ tự động đồng bộ (flush) các thay đổi vào database khi:**

- Gọi **commit transaction**.
- Gọi `entityManager.flush()`.
- Gọi `entityManager.refresh(entity)`.
- Gọi một truy vấn SQL/JPA khác có liên quan.

👉 **Cách để đưa entity từ Detached -> Managed:**
- Sử dụng `merge()`.
```java
User detachedUser = new User();
detachedUser.setId(1L);
detachedUser.setUsername("john_updated");

User managedUser = entityManager.merge(detachedUser);
```

🔥 **merge() sẽ lấy dữ liệu từ entity Detached và cập nhật vào entity Managed trong Persistence Context.**

### **3️⃣ Detached (Bị Tách Rời)**

🔹 Khi một entity đã từng được quản lý nhưng hiện tại không còn thuộc Persistence Context, nó ở trạng thái **Detached**.  
🔹 Các thay đổi trên entity sẽ không được cập nhật vào database trừ khi `merge()` được gọi.

📌 **Ví dụ:**
```java
User user = entityManager.find(User.class, 1L);
entityManager.detach(user); // user giờ là Detached
user.setPassword("newpassword"); // Không có hiệu lực vì user bị tách khỏi Persistence Context
```

👉 **Cách đưa entity từ Managed -> Detached:**

- Gọi `entityManager.detach(entity)`.
- Đóng EntityManager hoặc kết thúc transaction.

### **4️⃣ Removed (Bị Xóa)**

🔹 Khi một entity bị đánh dấu là **Removed**, nó sẽ bị xóa khỏi database sau khi transaction commit.  
🔹 Tuy nhiên, nó vẫn nằm trong Persistence Context cho đến khi `flush()` hoặc `commit()` xảy ra.

📌 **Ví dụ:**
```java
User user = entityManager.find(User.class, 1L);
entityManager.remove(user); // Đánh dấu user là Removed
// Khi transaction commit, user sẽ bị xóa khỏi database.
```

👉 **Lưu ý:** Nếu gọi `persist(user)` sau khi `remove()`, entity sẽ quay lại trạng thái **Managed**.

## **2.3. Sự Khác Nhau Giữa persist(), merge(), remove(), detach()**

| Phương thức       | Trạng thái ban đầu | Trạng thái sau khi gọi |
| ----------------- | ------------------ | ---------------------- |
| `persist(entity)` | Transient          | Managed                |
| `merge(entity)`   | Detached           | Managed                |
| `remove(entity)`  | Managed            | Removed                |
| `detach(entity)`  | Managed            | Detached               |

## **Khi Nào Cần Hiểu Về Lifecycle Của Entity?**

- **Khi làm việc với Transaction**: Tránh mất dữ liệu khi Persistence Context bị đóng.
- **Khi tối ưu hiệu suất**: Giảm số lượng query bằng cách tận dụng cơ chế Managed và tránh truy vấn không cần thiết.
- **Khi xử lý caching**: Hiểu được khi nào entity còn trong cache giúp tránh lỗi **LazyInitializationException**.

## **Tóm Tắt**

✅ Entity có 4 trạng thái chính: **Transient, Managed, Detached, Removed**.  
✅ **persist()** đưa entity từ **Transient -> Managed**.  
✅ **merge()** đưa entity từ **Detached -> Managed**.  
✅ **remove()** đưa entity từ **Managed -> Removed**.  
✅ **detach()** đưa entity từ **Managed -> Detached**.


---
# **Entity Lifecycle Callbacks trong JPA** 🚀

JPA cung cấp **Entity Lifecycle Callbacks** giúp thực thi các phương thức tự động khi một entity thay đổi trạng thái. Điều này hữu ích để:  
✅ Ghi log các sự kiện quan trọng.  
✅ Kiểm tra hoặc tính toán dữ liệu trước khi lưu vào DB.  
✅ Đồng bộ dữ liệu với hệ thống khác.

## **Các Annotation Lifecycle trong JPA**

JPA hỗ trợ **6 loại callback**, mỗi loại tương ứng với một giai đoạn trong lifecycle của entity:

|Annotation|Thời điểm kích hoạt|
|---|---|
|`@PostLoad`|Sau khi entity được load từ DB|
|`@PrePersist`|Trước khi entity được persist (insert vào DB)|
|`@PostPersist`|Sau khi entity được persist thành công|
|`@PreUpdate`|Trước khi entity được update|
|`@PostUpdate`|Sau khi entity được update thành công|
|`@PreRemove`|Trước khi entity bị xóa khỏi DB|
|`@PostRemove`|Sau khi entity bị xóa khỏi DB|

Ví dụ Thực Tế với Entity Lifecycle Callbacks

```java
import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
@EntityListeners(UserEntityListener.class) // Cách 2: Listener
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String username;
    private String email;

    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;

    // Lifecycle Callbacks
    @PrePersist
    public void prePersist() {
        this.createdAt = LocalDateTime.now();
        System.out.println("💾 User is about to be persisted: " + this.username);
    }

    @PostPersist
    public void postPersist() {
        System.out.println("✅ User persisted: " + this.username);
    }

    @PreUpdate
    public void preUpdate() {
        this.updatedAt = LocalDateTime.now();
        System.out.println("✏️ User is about to be updated: " + this.username);
    }

    @PostUpdate
    public void postUpdate() {
        System.out.println("✅ User updated: " + this.username);
    }

    @PreRemove
    public void preRemove() {
        System.out.println("❌ User is about to be removed: " + this.username);
    }

    @PostRemove
    public void postRemove() {
        System.out.println("✅ User removed: " + this.username);
    }

    @PostLoad
    public void postLoad() {
        System.out.println("🔄 User loaded: " + this.username);
    }
}
```

📌 **Chạy thử:**
```java
User user = new User();
user.setUsername("john_doe");
user.setEmail("john@example.com");

entityManager.persist(user);  // Kích hoạt @PrePersist và @PostPersist
user.setEmail("john_updated@example.com");
entityManager.merge(user);    // Kích hoạt @PreUpdate và @PostUpdate
entityManager.remove(user);   // Kích hoạt @PreRemove và @PostRemove
```

📌 **Output Console:**
```
💾 User is about to be persisted: john_doe
✅ User persisted: john_doe
✏️ User is about to be updated: john_doe
✅ User updated: john_doe
❌ User is about to be removed: john_doe
✅ User removed: john_doe
```
Cách 2: Tạo Listener Class để Xử Lý Callbacks
JPA cho phép tách logic callback ra một class riêng để **giữ cho entity sạch hơn**.
🔥 **Ưu điểm của cách này:**  
✅ **Tách biệt logic callback** ra khỏi entity, giúp entity dễ đọc hơn.  
✅ **Tái sử dụng** listener cho nhiều entity khác.


## **Khi Nào Nên Dùng Entity Lifecycle Callbacks?**

- **Ghi log (Audit Log)**: Ghi nhận sự kiện khi entity được tạo, cập nhật hoặc xóa.
- **Tự động cập nhật timestamps** (`createdAt`, `updatedAt`).
- **Đồng bộ dữ liệu với hệ thống khác**: Gửi event khi entity thay đổi.
- **Validate dữ liệu trước khi insert hoặc update**.