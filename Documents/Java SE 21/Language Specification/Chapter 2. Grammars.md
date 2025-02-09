
---
Tài liệu JLS nói rằng:

> "A Java program consists of one or more compilation units (§7.3) that are compiled jointly (§7.3). Each compilation unit automatically imports every type declared in the predefined package java.lang (§7.4.3)."

### **1️⃣ Chương trình Java bao gồm gì?**

- Một chương trình Java bao gồm **một hoặc nhiều đơn vị biên dịch (compilation units)**.
- **Tất cả các compilation units đều được biên dịch cùng nhau.**
- **Mặc định, Java tự động import toàn bộ các class trong package `java.lang`.**

📌 **Ví dụ về compilation unit (`Main.java`)**

Ví dụ một chương trình đơn giản:
```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, Java!");
    }
}
```
👉 **Mỗi chương trình Java bắt đầu bằng một class chứa phương thức `main`.**

**📌 Compilation Unit là gì?**

- Mỗi file `.java` được gọi là một **Compilation Unit**.
- Nó có thể chứa **class, interface, enum, record**, v.v.

📌 Biên dịch Compilation Unit:
```
javac Main.java  # Biên dịch thành Main.class
java Main        # Chạy chương trình
```
### **2️⃣ Java tự động import `java.lang`**

Khi bạn viết chương trình Java, bạn có thể sử dụng các class như `String`, `System`, `Math` mà không cần import. Đó là vì Java **tự động import toàn bộ package `java.lang`**.

📌 **Ví dụ: Sử dụng class từ `java.lang`**
```java
public class AutoImport {
    public static void main(String[] args) {
        System.out.println(Math.sqrt(16)); // Không cần import Math
    }
}
```
Vì `Math` thuộc `java.lang`, ta không cần `import java.lang.Math;`.
### **3️⃣ Nhiều Compilation Units cùng biên dịch**

Trong một chương trình lớn, bạn có thể có **nhiều file `.java`**, tất cả sẽ được biên dịch cùng lúc.
📌 **Biên dịch và chạy chương trình**:
```
javac Animal.java Dog.java Main.java
java Main
```
👉 **Tất cả compilation units đều phải được biên dịch trước khi chạy.**


