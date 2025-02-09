
---
Chương 16 của **Java Language Specification (JLS)** là về **Definite Assignment (Gán giá trị xác định)**


## **1️⃣ Gán giá trị xác định là gì?**

Trong Java, trước khi sử dụng một biến cục bộ (local variable), trình biên dịch **bắt buộc** biến đó phải được gán giá trị một cách rõ ràng (**definite assignment**). Điều này giúp tránh lỗi sử dụng **biến chưa khởi tạo**.

📌 **Ví dụ hợp lệ:**
```java
void test() {
    int x;  
    x = 10;  // Xác định giá trị trước khi dùng
    System.out.println(x); // ✅ Hợp lệ
}
```
📌 **Ví dụ lỗi (biến chưa được gán giá trị):**
```java
void test() {
    int x;
    System.out.println(x); // ❌ Lỗi: Variable 'x' might not have been initialized
}
```
## **2️⃣ Gán giá trị xác định với biến instance và static**

- **Biến instance (thuộc về đối tượng) và biến static (thuộc về lớp)** **luôn** được gán giá trị mặc định.
- **Biến cục bộ (local variable)** **không có giá trị mặc định** → Bắt buộc phải gán giá trị trước khi dùng.

📌 **Ví dụ hợp lệ với biến instance & static:**
```java
class Example {
    static int staticVar; // Mặc định là 0
    int instanceVar; // Mặc định là 0

    void display() {
        System.out.println(staticVar); // ✅ In ra 0
        System.out.println(instanceVar); // ✅ In ra 0
    }
}
```
📌 **Ví dụ lỗi với biến cục bộ:**
```java
void test() {
    int x; 
    System.out.println(x); // ❌ Lỗi: Variable 'x' might not have been initialized
}
```
## **3️⃣ Gán giá trị xác định trong các biểu thức điều kiện**

### **`if-else`**

Trình biên dịch Java kiểm tra tất cả các nhánh của `if-else` để đảm bảo biến được gán trước khi dùng.

📌 **Ví dụ hợp lệ:**
```java
void test(boolean condition) {
    int x;
    if (condition) {
        x = 10; 
    } else {
        x = 20; 
    }
    System.out.println(x); // ✅ Hợp lệ vì x luôn được gán
}
```
📌 **Ví dụ lỗi (một nhánh không gán giá trị):**
```java
void test(boolean condition) {
    int x;
    if (condition) {
        x = 10;  
    }
    System.out.println(x); // ❌ Lỗi nếu condition == false, x chưa được gán
}
```
💡 **Cách khắc phục:**  
Gán giá trị mặc định ban đầu.
```java
void test(boolean condition) {
    int x = 0; // ✅ Luôn được gán
    if (condition) {
        x = 10; 
    }
    System.out.println(x); 
}
```
## **4️⃣ Gán giá trị xác định trong vòng lặp**

### **Vòng lặp `while` và `for`**

Nếu biến được gán trong một vòng lặp, Java phải đảm bảo **vòng lặp luôn chạy ít nhất một lần** trước khi biến được sử dụng.

📌 **Ví dụ lỗi (có thể không chạy `while`):**
```java
void test(boolean condition) {
    int x;
    while (condition) { // Nếu condition == false, x không được gán
        x = 10;
    }
    System.out.println(x); // ❌ Lỗi: x có thể chưa được gán
}
```
💡 **Cách khắc phục:**  
Gán giá trị mặc định hoặc kiểm soát logic vòng lặp.
```java
void test(boolean condition) {
    int x = 0; // ✅ Luôn có giá trị
    while (condition) {
        x = 10;
    }
    System.out.println(x);
}
```
## **5️⃣ Gán giá trị xác định với `final`**

Khi một biến cục bộ được khai báo là `final`, nó phải được gán một lần duy nhất **trước khi sử dụng**.

📌 **Ví dụ hợp lệ:**
```java
void test(boolean condition) {
    final int x;
    if (condition) {
        x = 10;
    } else {
        x = 20;
    }
    System.out.println(x); // ✅ Hợp lệ vì x luôn được gán
}
```
# **📌 6. Tổng kết**

✅ Java **bắt buộc biến cục bộ phải được gán giá trị trước khi sử dụng** để tránh lỗi runtime.  
✅ Biến **instance và static luôn có giá trị mặc định**, nhưng **biến cục bộ thì không**.  
✅ Trình biên dịch Java kiểm tra **tất cả các nhánh `if-else` và vòng lặp** để đảm bảo biến được gán trước khi dùng.  
✅ Biến `final` phải được gán **một lần duy nhất** và **trước khi sử dụng**.

