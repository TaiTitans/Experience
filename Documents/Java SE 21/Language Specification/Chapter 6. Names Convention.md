
---
### **🔥 1. Giới thiệu về Tên trong Java**

Trong Java, **tên (name)** là cách chúng ta xác định các thực thể trong chương trình như **biến, phương thức, lớp, giao diện, gói (package)**, v.v.  
Tên giúp chương trình có cấu trúc rõ ràng, dễ hiểu và dễ quản lý.

### **🔥 2. Các loại Tên trong Java**

Java định nghĩa nhiều loại tên khác nhau, bao gồm:

1. **Tên đơn (Simple Name)**: Là tên không chứa dấu chấm `.`
    - Ví dụ: `x`, `myMethod`, `Person`
2. **Tên đủ (Qualified Name)**: Là tên có chứa dấu chấm `.` để thể hiện mối quan hệ giữa các thành phần.
    - Ví dụ: `java.util.List`, `mypackage.MyClass`, `System.out`
3. **Tên hoàn toàn đủ (Fully Qualified Name)**: Là tên bao gồm cả gói chứa nó.
    - Ví dụ: `java.util.ArrayList`, `com.example.MyService`

### **🔥 3. Quy tắc đặt Tên trong Java**

#### **3.1. Quy tắc cú pháp**

Tên trong Java phải tuân theo các quy tắc sau:  
✅ **Chỉ chứa chữ cái (`A-Z`, `a-z`), số (`0-9`), dấu gạch dưới (`_`), và dấu `$`**  
✅ **Không bắt đầu bằng số**  
✅ **Không trùng với từ khóa của Java**

📌 **Ví dụ hợp lệ:**
```java
int myVariable;
String $name;
double _price;
```
🚫 **Ví dụ không hợp lệ:**
```java
int 123abc;   // ❌ Lỗi: Không thể bắt đầu bằng số
float float;  // ❌ Lỗi: "float" là từ khóa
char my-name; // ❌ Lỗi: Không thể chứa dấu '-'
```
#### **3.2. Quy tắc đặt tên theo Java Convention**

Mặc dù Java không bắt buộc cách đặt tên, nhưng có những **quy ước (convention)** giúp mã nguồn dễ đọc và dễ bảo trì:

|Thành phần|Quy tắc đặt tên|Ví dụ|
|---|---|---|
|**Biến (variable)**|camelCase (chữ thường, từ thứ 2 viết hoa)|`customerName`, `totalPrice`|
|**Hằng số (constant)**|SNAKE_CASE (chữ hoa, cách bằng `_`)|`MAX_SIZE`, `PI_VALUE`|
|**Lớp (class)**|PascalCase (chữ đầu mỗi từ viết hoa)|`Person`, `OrderService`|
|**Giao diện (interface)**|PascalCase (chữ đầu mỗi từ viết hoa)|`Runnable`, `Serializable`|
|**Phương thức (method)**|camelCase (chữ thường, từ thứ 2 viết hoa)|`getName()`, `calculateTotal()`|
|**Gói (package)**|all lowercase (chữ thường, không dấu `_`)|`com.example.utils`|
|**Enum**|PascalCase cho Enum, SNAKE_CASE cho giá trị|`enum Status { ACTIVE, INACTIVE }`|

📌 **Ví dụ tuân theo convention:**
```java
class CustomerManager {
    private String customerName;  // Biến theo camelCase

    public String getCustomerName() {  // Phương thức theo camelCase
        return customerName;
    }
}
```

### **🔥 4. Phạm vi của Tên (Scope of Names)**

Phạm vi của một tên trong Java xác định nơi nó có thể được sử dụng. Có 4 loại chính:

#### **4.1. Phạm vi biến cục bộ (Local Scope)**

✅ **Chỉ có thể sử dụng trong khối `{}` mà nó được khai báo.**  
📌 **Ví dụ:**
```java
public class Example {
    public void printMessage() {
        String message = "Hello"; // Chỉ tồn tại trong printMessage()
        System.out.println(message);
    }
}
```
#### **4.2. Phạm vi biến thực thể (Instance Scope)**

✅ **Thuộc về đối tượng (object), tồn tại khi đối tượng còn tồn tại.**  
📌 **Ví dụ:**
```java
class Person {
    String name; // Biến thực thể

    void sayHello() {
        System.out.println("Hello, " + name);
    }
}
```
#### **4.3. Phạm vi biến lớp (Class Scope - Static)**

✅ **Gắn liền với lớp, không phụ thuộc vào đối tượng.**  
📌 **Ví dụ:**
```java
class MathUtil {
    static final double PI = 3.14159; // Biến static có phạm vi lớp
}
```
#### **4.4. Phạm vi toàn cục (Global Scope - Package/Imports)**

✅ **Có thể truy cập từ các lớp khác nếu được khai báo `public` hoặc `protected`.**  
📌 **Ví dụ:**
```java
package com.example.utils; // Package-level scope

public class MathUtil {
    public static double square(double x) {
        return x * x;
    }
}
```
### **🔥 5. Xung đột Tên (Name Shadowing & Hiding)**

Java cho phép nhiều tên giống nhau trong các phạm vi khác nhau, nhưng có thể gây lỗi nếu không cẩn thận.

#### **5.1. Shadowing 

🔹 **Biến cục bộ có thể che khuất biến thực thể cùng tên**  
📌 **Ví dụ:**
```java
class ShadowingExample {
    int x = 10; // Biến thực thể

    void printX() {
        int x = 20; // Biến cục bộ che khuất biến thực thể
        System.out.println(x); // In ra 20, không phải 10
    }
}
```
- Dùng `this.x` để lấy giá trị biến instance

#### **5.2. Hiding (Ẩn biến static)**

🔹 **Biến static bị ẩn khi có biến instance cùng tên trong lớp con**  
📌 **Ví dụ:**
```java
class Parent {
    static String message = "Hello from Parent";
}

class Child extends Parent {
    static String message = "Hello from Child"; // Ẩn message của Parent
}
```
### **🔥 6. Nhập Tên từ Package khác (Importing Names)**

Trong Java, chúng ta có thể sử dụng `import` để nhập các lớp từ package khác.

#### **6.1. Nhập từng lớp**
```java
import java.util.List;
import java.util.ArrayList;
```
#### **6.2. Nhập toàn bộ package**
```java
import java.util.*;
```
#### **6.3. Nhập tĩnh (Static Import)
```java
import static java.lang.Math.PI;
import static java.lang.Math.sqrt;
```
### **🔥 7. Kết luận**

✅ **Java có nhiều loại tên khác nhau**, từ đơn giản đến đầy đủ.  
✅ **Quy tắc đặt tên giúp mã nguồn dễ đọc, dễ bảo trì**.  
✅ **Phạm vi của tên quyết định nơi nó có thể được sử dụng**.  
✅ **Xung đột tên có thể xảy ra khi shadowing hoặc hiding biến**.  
✅ **Dùng `import` để tái sử dụng tên từ package khác**.

