
---
Chương này tập trung vào các biểu thức (expressions) trong Java – cách Java đánh giá và thực thi các biểu thức trong chương trình.

# **📌 1. Biểu thức là gì?**

Biểu thức (Expression) trong Java là bất kỳ đoạn mã nào có thể được **đánh giá (evaluate)** để tạo ra một giá trị.

📌 **Ví dụ:**
```java
int x = 10 + 5;  // Biểu thức "10 + 5" tạo ra giá trị 15
boolean result = (x > 10);  // Biểu thức so sánh
```
# **📌 2. Các loại biểu thức trong Java**

## **1️⃣ Biểu thức giá trị (Primary Expressions)**

Đây là những biểu thức đơn giản nhất, bao gồm:

- **Literals (Hằng số)** → `10`, `"Hello"`, `3.14`, `true`
- **Biến (Variable access)** → `x`, `myVar`
- **Phép toán `this` và `super`**
- **Gọi phương thức** → `myObject.method()`

📌 **Ví dụ:**
```java
int x = 10;  // x là một biểu thức giá trị
System.out.println("Hello");  // Gọi phương thức là một biểu thức
```
## **2️⃣ Biểu thức toán tử (Operators and Expressions)**

Gồm các phép toán như:  
✅ **Toán tử số học:** `+`, `-`, `*`, `/`, `%`  
✅ **Toán tử logic:** `&&`, `||`, `!`  
✅ **Toán tử so sánh:** `==`, `!=`, `<`, `>`  
✅ **Toán tử bit:** `&`, `|`, `^`, `~`, `<<`, `>>`  
✅ **Toán tử gán:** `=`, `+=`, `-=`, `*=`, `/=`

📌 **Ví dụ:**
```java
int a = 10, b = 20;
int sum = a + b;  // Toán tử số học
boolean check = (a < b) && (b > 15);  // Toán tử logic và so sánh
```
## **3️⃣ Biểu thức điều kiện (`if-else` và `switch`)**

✅ **Toán tử ba ngôi (`?:`)**  
📌 **Ví dụ:**
```java
int age = 18;
String status = (age >= 18) ? "Adult" : "Minor";
```
✅ **Switch Expression (Java 14+)**  
📌 **Ví dụ:**
```java
String result = switch (age) {
    case 18 -> "Just became an adult";
    case 20 -> "In twenties";
    default -> "Other";
};
```

## **4️⃣ Biểu thức gọi phương thức (Method Invocation Expressions)**

Gọi phương thức cũng là một biểu thức, vì nó trả về một giá trị.

📌 **Ví dụ:**
```java
String s = "Hello".toUpperCase();  // "Hello" là một biểu thức, gọi phương thức cũng là một biểu thức
```
## **5️⃣ Biểu thức Lambda (Lambda Expressions)**

Lambda là biểu thức vô danh dùng trong lập trình hàm (Functional Programming).

📌 **Ví dụ:**
```java
Function<Integer, Integer> square = x -> x * x;
System.out.println(square.apply(5));  // Kết quả: 25
```
# **📌 3. Đánh giá biểu thức trong Java**

✅ Java **tính toán từ trái sang phải** theo **độ ưu tiên toán tử**.  
✅ **Ngắn mạch (Short-circuit evaluation)**:

- `&&` dừng lại nếu vế trái là `false`.
- `||` dừng lại nếu vế trái là `true`.

📌 **Ví dụ:**
```java
boolean result = (5 > 3) || (10 / 0 > 1);  // Không bị lỗi chia cho 0 vì vế trái đã là true
```
# **📌 4. Tổng kết**

✅ **Biểu thức trong Java là nền tảng của mọi phép tính và logic**.  
✅ **Có nhiều loại biểu thức**, từ toán học, logic, điều kiện đến lambda.  
✅ **Java xử lý biểu thức theo quy tắc chặt chẽ**, bao gồm độ ưu tiên và ngắn mạch.