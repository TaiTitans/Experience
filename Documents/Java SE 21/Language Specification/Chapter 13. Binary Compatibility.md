
---
Chương 13 của **Java Language Specification (JLS)** nói về **Binary Compatibility (Tương thích nhị phân)**. Đây là một phần quan trọng khi nâng cấp mã nguồn mà vẫn giữ được khả năng tương thích với các phiên bản trước.

---
# **📌 1. Tương thích nhị phân là gì?**

Tương thích nhị phân trong Java có nghĩa là:  
✅ **Mã bytecode biên dịch từ phiên bản cũ vẫn chạy trên phiên bản mới** mà không cần biên dịch lại.  
✅ Các thay đổi trong mã nguồn không làm hỏng các chương trình đã biên dịch trước đó.

💡 **Ví dụ về tính tương thích nhị phân:**

- **Thêm phương thức mới vào class cũ** vẫn giữ tính tương thích.
- **Sửa đổi phần thân của một phương thức** không làm hỏng chương trình đã biên dịch trước đó.

# **📌 2. Những thay đổi KHÔNG phá vỡ tính tương thích nhị phân**

Các thay đổi sau vẫn giữ được tính tương thích:
### **✅ Thêm phương thức hoặc constructor mới**

📌 **Ví dụ:**
```java
public class MyClass {
    public void oldMethod() {}
    
    // Thêm phương thức mới
    public void newMethod() {}
}
```
✔ Chương trình biên dịch cũ vẫn chạy được, miễn là không gọi `newMethod()`.
### **✅ Sửa đổi phần thân của phương thức**
📌 **Ví dụ:**
```java
public class MyClass {
    public void printMessage() {
        System.out.println("Phiên bản cũ");
    }
}

// Sau này cập nhật:
public class MyClass {
    public void printMessage() {
        System.out.println("Phiên bản mới");
    }
}
```
✔ Không làm ảnh hưởng đến chương trình cũ.
### **✅ Thêm một class hoặc interface mới**
📌 **Ví dụ:**
```java
// Ban đầu
public class OldClass {}

// Sau đó thêm class mới
public class NewClass {}
```
✔ Không ảnh hưởng đến mã đã biên dịch trước đó.
### **✅ Thêm một field mới vào class**
📌 **Ví dụ:**
```java
public class MyClass {
    public int oldField;
}

// Sau đó thêm:
public class MyClass {
    public int oldField;
    public String newField; // Thêm mới
}
```
✔ Không phá vỡ chương trình cũ.

---
# **📌 3. Những thay đổi PHÁ VỠ tính tương thích nhị phân**

Một số thay đổi có thể khiến chương trình không chạy được nếu không biên dịch lại.
### **❌ Thay đổi kiểu dữ liệu của field**
📌 **Ví dụ:**
```java
public class MyClass {
    public int myField;
}

// Sau đó sửa thành:
public class MyClass {
    public String myField; // ❌ Lỗi tương thích
}
```
🚨 **Lỗi**: Code cũ truy cập `myField` kiểu `int` nhưng code mới lại là `String`.
### **❌ Xóa một phương thức hoặc constructor**
📌 **Ví dụ:**
```java
public class MyClass {
    public void myMethod() {}
}

// Nếu xóa:
public class MyClass {} // ❌ Lỗi
```
🚨 **Lỗi**: Chương trình cũ gọi `myMethod()`, nhưng nó đã bị xóa.
### **❌ Thay đổi kiểu trả về của phương thức**
📌 **Ví dụ:**
```java
public class MyClass {
    public int getValue() { return 42; }
}

// Nếu thay đổi thành:
public class MyClass {
    public String getValue() { return "42"; } // ❌ Lỗi
}
```
🚨 **Lỗi**: Chương trình cũ mong đợi kiểu `int`, nhưng bây giờ là `String`.
### **❌ Di chuyển một class sang package khác**
📌 **Ví dụ:**
```java
package oldpackage;
public class MyClass {}

// Nếu di chuyển sang package khác:
package newpackage; // ❌ Lỗi
public class MyClass {}
```
🚨 **Lỗi**: Các chương trình cũ sử dụng `oldpackage.MyClass` không thể tìm thấy class nữa.
# **📌 4. Tổng kết**

✅ **Tương thích nhị phân giúp code cũ chạy được trên phiên bản mới mà không cần biên dịch lại.**  
✅ Các thay đổi như **thêm phương thức, thêm field, sửa body phương thức** vẫn giữ tính tương thích.  
❌ Các thay đổi như **xóa phương thức, đổi kiểu field, đổi kiểu trả về, di chuyển class** có thể phá vỡ chương trình.