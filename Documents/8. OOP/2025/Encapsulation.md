
---
# 🔥 **Encapsulation trong Java – Hiểu Sâu Từ Cơ Bản Đến Nâng Cao**

## 📌 **1️⃣ Encapsulation là gì?**

**Encapsulation (Đóng gói)** là một nguyên tắc của lập trình hướng đối tượng (OOP), trong đó:  
✅ **Dữ liệu (biến) của một lớp được bảo vệ** khỏi truy cập trực tiếp từ bên ngoài.  
✅ **Các phương thức (`getter` và `setter`) được cung cấp** để truy cập dữ liệu một cách an toàn.  
✅ **Hạn chế truy cập trực tiếp vào các thuộc tính**, giúp **kiểm soát chặt chẽ dữ liệu**.

### **Cú pháp cơ bản của Encapsulation**:
```java
class Person {
    private String name;  // Biến private chỉ truy cập trong lớp này

    // Getter - Lấy giá trị name
    public String getName() {
        return name;
    }

    // Setter - Gán giá trị cho name
    public void setName(String newName) {
        this.name = newName;
    }
}
```
📌 **Biến `name` là `private`, nhưng có thể truy cập gián tiếp qua `getName()` và `setName()`**.

## 📌 **2️⃣ Tại sao cần Encapsulation?**

🔹 **1. Bảo mật dữ liệu**: Ngăn chặn thay đổi dữ liệu trực tiếp từ bên ngoài lớp.  
🔹 **2. Kiểm soát dữ liệu**: Chỉ cho phép thay đổi thông qua các phương thức (`setter`) có kiểm tra hợp lệ.  
🔹 **3. Dễ bảo trì & mở rộng**: Nếu cần thay đổi logic, chỉ cần sửa trong lớp mà không ảnh hưởng đến mã bên ngoài.  
🔹 **4. Tăng tính module hóa**: Dữ liệu và hành vi được nhóm lại, giúp mã dễ hiểu và dễ quản lý.

## 📌 **3️⃣ Làm thế nào để đạt được Encapsulation trong Java?**

✅ **Bước 1:** Đặt biến thành `private`.  
✅ **Bước 2:** Cung cấp các phương thức `getter` và `setter` để truy cập biến.  
✅ **Bước 3:** Sử dụng các phương thức `getter` và `setter` để thay đổi hoặc lấy giá trị của biến.

🔹 **Ví dụ hoàn chỉnh về Encapsulation**:
```java
class BankAccount {
    private double balance;  // Số dư tài khoản (private)

    public BankAccount(double initialBalance) {
        if (initialBalance > 0) {
            this.balance = initialBalance;
        } else {
            this.balance = 0;
        }
    }

    // Getter - Lấy số dư
    public double get
```

## 🔥 **Hiểu Sâu Về Encapsulation**

### ✅ **1️⃣ Encapsulation Không Chỉ Là `private` + `getter/setter`**

Nhiều lập trình viên nghĩ rằng chỉ cần:
```java
class Person {
    private String name;

    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
}
```
📌 **Là đã áp dụng Encapsulation? Sai!**

❌ **Vấn đề 1**: Dữ liệu vẫn có thể bị thay đổi một cách không hợp lệ.
```java
Person p = new Person();
p.setName(null);  // Đặt giá trị không hợp lệ
```
❌ **Vấn đề 2**: Nếu tất cả biến đều có `getter/setter`, thì chúng ta lại vô tình "lộ" dữ liệu giống như khi dùng `public`!

💡 **Encapsulation đúng nghĩa** là **ẩn dữ liệu** (data hiding) + **kiểm soát quyền truy cập** (controlled access).
### ✅ **2️⃣ Cách Làm Đúng – Kiểm Soát Dữ Liệu**

Chúng ta phải **kiểm soát việc gán giá trị** trong `setter`.
```java
class Person {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        if (name == null || name.trim().isEmpty()) {
            throw new IllegalArgumentException("Tên không được để trống!");
        }
        this.name = name;
    }
}
```
📌 **Giờ đây, dữ liệu luôn hợp lệ!**


