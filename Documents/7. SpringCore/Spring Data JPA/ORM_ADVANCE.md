
---
Trong JPA, **Cascade & Fetching Strategies** giúp kiểm soát cách các entity liên kết được xử lý khi thực hiện các thao tác CRUD.


## **Cascade Types trong JPA**

Cascade giúp tự động áp dụng các thao tác (persist, merge, remove, refresh, detach) từ **entity cha** sang **entity con**.

### **📌 Các loại Cascade chính**

|Cascade Type|Mô tả|
|---|---|
|`CascadeType.PERSIST`|Khi `persist` entity cha, entity con cũng được `persist`.|
|`CascadeType.MERGE`|Khi `merge` entity cha, entity con cũng được `merge`.|
|`CascadeType.REMOVE`|Khi `remove` entity cha, entity con cũng bị `remove`.|
|`CascadeType.REFRESH`|Khi `refresh` entity cha, entity con cũng được `refresh`.|
|`CascadeType.DETACH`|Khi `detach` entity cha, entity con cũng bị `detach`.|
|`CascadeType.ALL`|Gồm tất cả các loại trên.|
### **📌 Ví dụ: Cascade trong JPA**

Giả sử **User có nhiều Orders**, khi tạo User mới, ta muốn tự động `persist` luôn danh sách Orders.
```java
import jakarta.persistence.*;
import java.util.List;

@Entity
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String username;

    @OneToMany(mappedBy = "user", cascade = CascadeType.ALL, orphanRemoval = true)
    private List<Order> orders;

    // Getters và Setters
}

@Entity
public class Order {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String productName;

    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user;

    // Getters và Setters
}
```

---
## **Fetching Strategies trong JPA**

Fetching Strategy xác định **khi nào** và **cách nào** dữ liệu liên quan được tải từ database.

### **📌 Hai chiến lược Fetch chính**

|Fetch Type|Mô tả|
|---|---|
|`FetchType.EAGER`|Tải dữ liệu ngay lập tức, ngay cả khi không cần dùng đến.|
|`FetchType.LAZY`|Chỉ tải dữ liệu khi cần truy cập.|

### **📌 Ví dụ FetchType**

```java
@Entity
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String username;

    @OneToMany(mappedBy = "user", cascade = CascadeType.ALL, fetch = FetchType.LAZY)
    private List<Order> orders;
}
```

**🚀 Khi nào dùng FetchType LAZY hay EAGER?**  
✅ **LAZY** (Mặc định với `@OneToMany`) – Dùng khi **không phải lúc nào cũng cần dữ liệu con**.  
✅ **EAGER** (Mặc định với `@ManyToOne`) – Dùng khi **luôn cần dữ liệu con** ngay khi load entity cha.

# **Inheritance trong JPA** 🏗️

JPA hỗ trợ kế thừa giúp một entity **có thể mở rộng từ một entity khác**, giúp tái sử dụng code.

## **5.1. Các Chiến Lược Inheritance trong JPA**

JPA cung cấp 3 chiến lược kế thừa chính:

|Chiến lược|Mô tả|
|---|---|
|**SINGLE_TABLE**|Lưu tất cả entity vào một bảng duy nhất (Dùng chung cột, có cột `dtype` để phân loại).|
|**JOINED**|Mỗi subclass có bảng riêng, dùng `JOIN` để lấy dữ liệu.|
|**TABLE_PER_CLASS**|Mỗi subclass có bảng riêng, không có mối quan hệ `JOIN`.|

## **SINGLE_TABLE Strategy**

📌 **Tất cả entity con dùng chung một bảng, phân biệt bằng cột `dtype`**.

### **📌 Cấu trúc Database**
`📄 payment (id, amount, payment_type, cardNumber, paypalEmail)`

📌 Code Implementation
```java
@Entity
@Inheritance(strategy = InheritanceType.SINGLE_TABLE)
@DiscriminatorColumn(name = "payment_type", discriminatorType = DiscriminatorType.STRING)
public abstract class Payment {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private double amount;
}

@Entity
@DiscriminatorValue("CREDIT_CARD")
public class CreditCardPayment extends Payment {
    private String cardNumber;
}

@Entity
@DiscriminatorValue("PAYPAL")
public class PaypalPayment extends Payment {
    private String paypalEmail;
}
```

📌 **Khi lưu dữ liệu**
```java
Payment cardPayment = new CreditCardPayment();
cardPayment.setAmount(100.0);
((CreditCardPayment) cardPayment).setCardNumber("1234-5678-9012");

Payment paypalPayment = new PaypalPayment();
paypalPayment.setAmount(200.0);
((PaypalPayment) paypalPayment).setPaypalEmail("user@example.com");

entityManager.persist(cardPayment);
entityManager.persist(paypalPayment);
```

🔥 **Ưu điểm:** Hiệu suất cao do chỉ cần một bảng, dễ truy vấn.  
⚠️ **Nhược điểm:** Cột NULL nhiều khi subclass không dùng chung tất cả các trường.

## **JOINED Strategy**

📌 **Mỗi entity con có bảng riêng, JOIN để lấy dữ liệu**.

### **📌 Cấu trúc Database**
```
📄 payment (id, amount)
📄 credit_card_payment (id, cardNumber)
📄 paypal_payment (id, paypalEmail)
```
📌 Code Implementation
```java
@Entity
@Inheritance(strategy = InheritanceType.JOINED)
public abstract class Payment {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private double amount;
}

@Entity
public class CreditCardPayment extends Payment {
    private String cardNumber;
}

@Entity
public class PaypalPayment extends Payment {
    private String paypalEmail;
}
```

🔥 **Ưu điểm:** Chuẩn hóa dữ liệu, tránh dư thừa cột NULL.  
⚠️ **Nhược điểm:** Truy vấn chậm hơn do cần `JOIN`.


## **TABLE_PER_CLASS Strategy**

📌 **Mỗi entity con có bảng riêng, không có `JOIN` với entity cha**.

### **📌 Cấu trúc Database**
```
📄 credit_card_payment (id, amount, cardNumber)
📄 paypal_payment (id, amount, paypalEmail)
```
📌 Code Implementation
```java
@Entity
@Inheritance(strategy = InheritanceType.TABLE_PER_CLASS)
public abstract class Payment {
    @Id
    @GeneratedValue(strategy = GenerationType.TABLE) // Cần TABLE ID Generator
    private Long id;
    private double amount;
}

@Entity
public class CreditCardPayment extends Payment {
    private String cardNumber;
}

@Entity
public class PaypalPayment extends Payment {
    private String paypalEmail;
}
```
🔥 **Ưu điểm:** Không có `JOIN`, truy vấn nhanh.  
⚠️ **Nhược điểm:** **Không thể dùng quan hệ `@ManyToOne` trên Payment**.

## **Khi nào dùng chiến lược nào?**

✅ `SINGLE_TABLE` → **Hiệu suất tốt, ít bảng, nhưng dữ liệu có nhiều NULL**.  
✅ `JOINED` → **Chuẩn hóa dữ liệu, nhưng truy vấn JOIN có thể chậm**.  
✅ `TABLE_PER_CLASS` → **Không có JOIN nhưng khó quản lý quan hệ giữa các bảng**.

