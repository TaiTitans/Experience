
---
Chương 14 của **Java Language Specification (JLS)** nói về **Blocks and Statements (Khối lệnh và câu lệnh)**. Đây là phần rất quan trọng vì nó định nghĩa cách Java thực thi từng câu lệnh trong chương trình.

---
# **📌 1. Khối lệnh (Block)**

Một **khối lệnh** (block) trong Java là một nhóm các câu lệnh được đặt trong dấu `{}`.  
📌 **Ví dụ:**
```java
{
    int x = 10;
    System.out.println(x);
}
```
✔ **Mỗi khối lệnh có phạm vi riêng**, các biến khai báo bên trong không thể truy cập từ bên ngoài.

💡 **Ứng dụng phổ biến của khối lệnh:**

- **Trong thân của phương thức:**
```java
public void myMethod() {
    int a = 5;
    System.out.println(a);
}
```
- **Trong vòng lặp hoặc điều kiện:
```java
if (true) {
    System.out.println("Hello");
}
```
- **Trong khai báo cục bộ (local block):
```java
int x = 10;
{
    int y = 5; // Chỉ có phạm vi trong khối này
    System.out.println(y);
}
// System.out.println(y); // ❌ Lỗi vì y không tồn tại ngoài khối
```
# 📌 2. Các loại câu lệnh trong Java

Chương này định nghĩa nhiều loại câu lệnh, bao gồm:
### **1️⃣ Câu lệnh trống (Empty Statement)**

Là một dấu `;` đơn lẻ, không làm gì cả.  
📌 **Ví dụ:**
```java;
;
System.out.println("Hello"); // Câu lệnh này vẫn chạy bình thường
```
💡 **Ứng dụng:** Đôi khi dùng trong vòng lặp hoặc cấu trúc điều kiện.
```java
while (condition)
    ; // Lặp vô tận mà không làm gì
```
### **2️⃣ Câu lệnh biểu thức (Expression Statement)**

Các biểu thức như gán giá trị, gọi phương thức, tăng giảm giá trị.  
📌 **Ví dụ:**
```java
int a = 10;      // Gán giá trị
a++;             // Biểu thức tăng giá trị
System.out.println(a); // Gọi phương thức
```
### **3️⃣ Câu lệnh khai báo (Declaration Statement)**

Khai báo biến hoặc phương thức.  
📌 **Ví dụ:**
```java
int x = 5;
double pi = 3.14;
```
### **4️⃣ Câu lệnh điều kiện (If-Else Statement)**

📌 **Ví dụ:**
```java
if (x > 0) {
    System.out.println("Positive number");
} else {
    System.out.println("Negative number");
}
```
### **5️⃣ Câu lệnh lặp (Loop Statements)**

📌 **Vòng lặp `for`:**
```java
for (int i = 0; i < 5; i++) {
    System.out.println(i);
}
```
📌 **Vòng lặp `while`:**
```java
while (x > 0) {
    System.out.println(x);
    x--;
}
```
### **6️⃣ Câu lệnh `switch-case`**

📌 **Ví dụ:**
```java
int day = 2;
switch (day) {
    case 1:
        System.out.println("Monday");
        break;
    case 2:
        System.out.println("Tuesday");
        break;
    default:
        System.out.println("Other day");
}
```
### **7️⃣ Câu lệnh `break` và `continue`**

📌 **Dừng vòng lặp với `break`:**
```java
for (int i = 0; i < 10; i++) {
    if (i == 5) break;
    System.out.println(i);
}
```
📌 **Bỏ qua phần còn lại của vòng lặp với `continue`:**
```java
for (int i = 0; i < 10; i++) {
    if (i == 5) continue;
    System.out.println(i);
}
```
### **8️⃣ Câu lệnh `return`**

📌 **Trả về giá trị trong phương thức:**
```java
public int getNumber() {
    return 42;
}
```
---
# **📌 1. Patterns là gì?**

**Pattern** là một cách để kiểm tra giá trị của một biến và trích xuất dữ liệu một cách an toàn. Java hỗ trợ **Pattern Matching** trong `instanceof`, `switch`, và Record.

🔹 **Lợi ích của Pattern Matching:**

- Giảm **ép kiểu thủ công** (casting).
- Cải thiện **tính rõ ràng** và **tính an toàn** của mã nguồn.
- Hỗ trợ **Destructuring** (trích xuất dữ liệu).

# **📌 2. Các loại Patterns trong Java 21**

### **1️⃣ Type Patterns (Kiểu mẫu)**

🔹 **Sử dụng trong `instanceof` để kiểm tra kiểu dữ liệu**  
📌 **Trước đây (Java 14 trở về trước):**
```java
Object obj = "Hello";
if (obj instanceof String) {
    String s = (String) obj; // Ép kiểu thủ công
    System.out.println(s.length());
}
```
📌 **Sau khi có Pattern Matching (Java 16+):**
```java
Object obj = "Hello";
if (obj instanceof String s) { // Không cần ép kiểu
    System.out.println(s.length());
}
```
✅ Java tự động gán `obj` vào biến `s` nếu điều kiện đúng.

### **2️⃣ Record Patterns**

🔹 **Dùng để trích xuất dữ liệu từ Record**  
📌 **Ví dụ:**
```java
record Point(int x, int y) {}

void printCoordinates(Object obj) {
    if (obj instanceof Point(int x, int y)) {  // Record Pattern
        System.out.println("X: " + x + ", Y: " + y);
    }
}
```
✅ Java tự động "giải nén" giá trị của Record mà không cần gọi `getX()` hay `getY()`.
### **3️⃣ Deconstruction Patterns (Pattern Tổ Hợp)**

🔹 **Kết hợp nhiều Pattern bên trong nhau**  
📌 **Ví dụ:**
```java
record Circle(Point center, double radius) {}

void printCircleInfo(Object obj) {
    if (obj instanceof Circle(Point(int x, int y), double r)) { // Nested Patterns
        System.out.println("Tâm: (" + x + ", " + y + "), Bán kính: " + r);
    }
}
```
✅ Kết hợp `Record Pattern` và `Type Pattern` để trích xuất sâu hơn.
### **4️⃣ Switch Pattern Matching**

🔹 **Dùng Pattern trong `switch-case`**  
📌 **Ví dụ:**
```java
static void test(Object obj) {
    switch (obj) {
        case Integer i -> System.out.println("Số nguyên: " + i);
        case String s -> System.out.println("Chuỗi: " + s);
        case Point(int x, int y) -> System.out.println("Point: (" + x + ", " + y + ")");
        default -> System.out.println("Không xác định");
    }
}
```
✅ **Không cần `if-else` và ép kiểu thủ công**, giúp code gọn gàng hơn.
# **📌 3. Tổng kết**

✅ **Patterns giúp Java dễ đọc hơn** nhờ giảm ép kiểu thủ công.  
✅ Hỗ trợ trong `instanceof`, `switch`, `Record`, và **deconstruction**.  
✅ **Cải thiện hiệu suất** vì Java thực hiện kiểm tra kiểu tối ưu hơn.