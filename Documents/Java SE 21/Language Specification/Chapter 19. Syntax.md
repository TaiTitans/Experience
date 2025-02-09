
---
*Bắt đầu rồi kết thúc, rồi lại bắt đầu mới.

Chương này trong **Java Language Specification (JLS)** mô tả chi tiết về **cú pháp của Java**, bao gồm quy tắc viết mã nguồn Java theo định dạng chính thức bằng **BNF (Backus-Naur Form)**.

## **1️⃣ BNF là gì?**

BNF (**Backus-Naur Form**) là cách mô tả cú pháp của ngôn ngữ lập trình bằng một tập hợp quy tắc **ngữ pháp phi ngữ cảnh** (_context-free grammar_).

Ví dụ về quy tắc BNF đơn giản:
```
<expression> ::= <term> | <term> "+" <expression>
<term> ::= <factor> | <factor> "*" <term>
<factor> ::= <number> | "(" <expression> ")"
```
Quy tắc này mô tả biểu thức toán học có phép cộng (`+`), nhân (`*`), và dấu ngoặc `()`.

Trong Java, các thành phần như **biến, kiểu dữ liệu, biểu thức, khối lệnh** đều có thể được biểu diễn bằng BNF.

## **2️⃣ Cú pháp tổng quan của Java**

Cú pháp Java được mô tả qua nhiều phần, bao gồm:

- **Compilation Unit (Đơn vị biên dịch)**
- **Declarations (Khai báo biến, lớp, phương thức, v.v.)**
- **Statements (Câu lệnh như `if`, `while`, `switch`, `return`, v.v.)**
- **Expressions (Biểu thức toán học, logic, v.v.)**

Ví dụ, quy tắc khai báo biến có thể được viết theo dạng BNF như sau:
```java
VariableDeclaration ::= Type VariableDeclarator
VariableDeclarator ::= Identifier [ "=" Expression ]
```
Điều này có nghĩa là một biến phải có một **kiểu dữ liệu**, một **tên**, và **có thể có giá trị khởi tạo**.

📌 **Ví dụ hợp lệ:**
```java
int x;       // Hợp lệ theo quy tắc trên
int y = 10;  // Hợp lệ, có giá trị khởi tạo
```
## **3️⃣ Cấu trúc của một chương trình Java**

Một chương trình Java gồm các thành phần sau:
```java
CompilationUnit ::= [PackageDeclaration] {ImportDeclaration} {TypeDeclaration}
```
- `PackageDeclaration`: Khai báo package (nếu có).
- `ImportDeclaration`: Khai báo các thư viện cần dùng.
- `TypeDeclaration`: Khai báo lớp hoặc interface.

📌 **Ví dụ:**
```java
package mypackage;  // PackageDeclaration
import java.util.List;  // ImportDeclaration

public class Example {  // TypeDeclaration
    public static void main(String[] args) {
        System.out.println("Hello, Java!");
    }
}
```
## **4️⃣ Quy tắc về từ khóa, tên biến, và toán tử**

### **🔹 Từ khóa**

Java có một danh sách từ khóa không thể dùng làm tên biến, bao gồm:
```java
abstract, continue, for, new, switch, assert, default, goto, package, synchronized, 
boolean, do, if, private, this, break, double, implements, protected, throw, 
byte, else, import, public, throws, case, enum, instanceof, return, transient, 
catch, extends, int, short, try, char, final, interface, static, void, 
class, finally, long, strictfp, volatile, const, float, native, super, while
```
### **🔹 Quy tắc đặt tên biến**

Tên biến phải bắt đầu bằng **chữ cái, `_` hoặc `$`**, không được bắt đầu bằng số, và không chứa khoảng trắng.
### **🔹 Toán tử**

Java hỗ trợ các loại toán tử như:

- **Toán tử số học**: `+`, `-`, `*`, `/`, `%`
- **Toán tử so sánh**: `==`, `!=`, `<`, `>`, `<=`, `>=`
- **Toán tử logic**: `&&`, `||`, `!`
- **Toán tử gán**: `=`, `+=`, `-=`, `*=`, `/=`
- **Toán tử bitwise**: `&`, `|`, `^`, `~`, `<<`, `>>`
- **Toán tử điều kiện**: `?:`
## **5️⃣ Cú pháp khối lệnh và điều khiển luồng**

Java sử dụng `{}` để nhóm các câu lệnh thành khối lệnh.
## **6️⃣ Biểu thức (Expressions)**

Java có nhiều loại biểu thức như:

- **Biểu thức số học:** `a + b * c`
- **Biểu thức logic:** `x > 0 && y < 10`
- **Biểu thức gán:** `x = y + 2`
- **Biểu thức điều kiện:** `result = (a > b) ? "A lớn hơn" : "B lớn hơn";`

---

## **7️⃣ Tổng kết**

✅ **Java sử dụng BNF để mô tả cú pháp chính xác.**  
✅ **Cấu trúc chương trình gồm package, import, class, method.**  
✅ **Tuân theo quy tắc đặt tên biến, từ khóa, toán tử, khối lệnh.**  
✅ **Hiểu được cú pháp giúp viết code chính xác hơn.**
