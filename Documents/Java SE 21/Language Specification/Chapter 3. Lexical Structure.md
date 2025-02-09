
---
## **1️⃣ Tổng quan**

Ngôn ngữ Java sử dụng **từ vựng (lexical elements)** để xây dựng mã nguồn. Các thành phần cơ bản gồm:

- **Mã Unicode**: Java sử dụng Unicode để biểu diễn ký tự.
- **Lexical Tokens**: Bao gồm từ khóa, định danh (tên biến, tên hàm), ký tự đặc biệt, toán tử, số, chuỗi.
- **Dấu cách trắng và chú thích**: Dùng để làm mã nguồn dễ đọc hơn.

## **2️⃣ Mã Unicode trong Java (JLS 3.1)**

Java sử dụng bộ mã **Unicode** (UTF-16) để biểu diễn ký tự, cho phép dùng nhiều ngôn ngữ khác nhau.

📌 **Ví dụ về Unicode trong Java**

```java
public class UnicodeExample {
    public static void main(String[] args) {
        char a = 'A'; // Ký tự ASCII
        char b = '\u03A9'; // Ký tự Unicode (Ω - Omega)
        System.out.println("Character A: " + a);
        System.out.println("Unicode Omega: " + b);
    }
}
```
📌 **Kết quả**
```
Character A: A
Unicode Omega: Ω
```

💡 **Lưu ý:** Bạn có thể dùng **`\uXXXX`** để viết ký tự Unicode.

## **3️⃣ Lexical Tokens (JLS 3.5 - 3.10)**

### 🔹 **3.1 - Nhóm Token trong Java**

Java chia token thành 5 nhóm chính:

1. **Từ khóa (Keywords)**: `if`, `else`, `class`, `public`, `static`, `void`, v.v.
2. **Định danh (Identifiers)**: Tên biến, tên hàm, tên class.
3. **Toán tử (Operators)**: `+`, `-`, `*`, `/`, `==`, `!=`, `&&`, `||`, v.v.
4. **Literals (Hằng số)**: Chuỗi `"Hello"`, số `123`, boolean `true`.
5. **Dấu phân cách (Separators)**: `{}`, `()`, `;`, `,`, `.`.

📌 **Ví dụ về Token trong Java**
```java
public class TokenExample {
    public static void main(String[] args) {
        int age = 25; // "int" là từ khóa, "age" là định danh, "25" là literal
        if (age >= 18) { // ">=" là toán tử, "{}" là dấu phân cách
            System.out.println("You are an adult.");
        }
    }
}
```

### 🔹 **3.2 - Nhận diện Định danh (Identifiers) (JLS 3.8)**

**Định danh** là tên dùng để đặt cho biến, phương thức, class, interface.  
**Quy tắc đặt tên định danh:** ✅ Bắt đầu bằng chữ cái, `_` hoặc `$`.  
✅ Không được là từ khóa của Java.  
✅ Phân biệt chữ hoa và chữ thường.

📌 **Ví dụ hợp lệ**
```java
int age;
String $name;
double _salary;
```
🚫 **Không hợp lệ**
```java
int 2number;  // ❌ Bắt đầu bằng số
String class; // ❌ Trùng từ khóa
```
### 🔹 **3.3 - Hằng số Literals (JLS 3.10)**

📌 **Hằng số trong Java có nhiều kiểu**:

- **Số nguyên (Integer)**: `10`, `0xFF` (hex), `0b1010` (binary).
- **Số thực (Floating-point)**: `3.14`, `2.5e3` (2500.0).
- **Chuỗi (String)**: `"Hello, Java!"`
- **Boolean**: `true`, `false`
- **Null**: `null`

📌 **Ví dụ**
```java
int decimal = 10;    // Số thập phân
int hex = 0xA;       // Số hệ 16 (hex)
int binary = 0b1010; // Số nhị phân
double pi = 3.14;    
boolean isJavaFun = true;
String greeting = "Hello, Java!";
```
## **4️⃣ Dấu cách trắng và chú thích (JLS 3.6 - 3.7)**

🔹 **Dấu cách trắng (Whitespace)**  
Java bỏ qua dấu cách trắng (`space`, `tab`, `newline`), trừ khi nó nằm trong chuỗi.

🔹 **Chú thích trong Java**

- **Chú thích một dòng**: `// Đây là comment`
- **Chú thích nhiều dòng**:
```java
/*
  Đây là comment nhiều dòng
*/
```
**Chú thích Javadoc**: Dùng để tạo tài liệu API
```java
/**
 * Đây là comment Javadoc
 */
```
📌 **Ví dụ**
```java
public class CommentExample {
    public static void main(String[] args) {
        // In ra màn hình
        System.out.println("Hello, Java!"); 
    }
}
```
## **5️⃣ Tổng kết**

📌 **Chương này giúp bạn hiểu về:**  
✅ Unicode trong Java  
✅ Token: Từ khóa, định danh, toán tử, hằng số  
✅ Quy tắc đặt tên biến, kiểu dữ liệu  
✅ Dấu cách trắng và chú thích
