
---
## **1️⃣ Tổng quan**

Mọi chương trình Java đều sử dụng **biến (variables)** để lưu trữ **giá trị (values)**. Các giá trị này thuộc về một **kiểu dữ liệu (type)** cụ thể, có thể là kiểu cơ bản (**primitive**) hoặc kiểu tham chiếu (**reference**).

## **2️⃣ Kiểu Dữ Liệu (Types) (JLS 4.1)**

Java là một **ngôn ngữ kiểu tĩnh (statically typed language)**, nghĩa là mỗi biến phải có một kiểu cụ thể **tại thời điểm biên dịch (compile-time)**.

📌 **Java có hai loại kiểu dữ liệu chính:**

- **Kiểu nguyên thủy (Primitive types)**: `int`, `double`, `char`, `boolean`, v.v.
- **Kiểu tham chiếu (Reference types)**: Class, Interface, Array, Enum.

## **3️⃣ Giá trị (Values) (JLS 4.2)**

📌 Mỗi **kiểu dữ liệu** có **giá trị hợp lệ** mà nó có thể lưu trữ.

### **🔹 3.1 - Giá trị của Kiểu Nguyên Thủy**

|Kiểu|Kích thước|Giá trị tối thiểu|Giá trị tối đa|Ví dụ|
|---|---|---|---|---|
|`byte`|8-bit|-128|127|`byte b = 100;`|
|`short`|16-bit|-32,768|32,767|`short s = 30000;`|
|`int`|32-bit|-2^31|2^31 - 1|`int i = 100000;`|
|`long`|64-bit|-2^63|2^63 - 1|`long l = 9999999999L;`|
|`float`|32-bit|~1.4E-45|~3.4E+38|`float f = 3.14f;`|
|`double`|64-bit|~4.9E-324|~1.8E+308|`double d = 2.71828;`|
|`char`|16-bit|`\u0000`|`\uFFFF`|`char c = 'A';`|
|`boolean`|1-bit|`true` hoặc `false`|-|`boolean isJavaFun = true;`|

📌 **Ví dụ:**
```java
public class PrimitiveExample {
    public static void main(String[] args) {
        int age = 25;
        double pi = 3.14159;
        boolean isJavaFun = true;
        char letter = 'J';

        System.out.println("Age: " + age);
        System.out.println("Pi: " + pi);
        System.out.println("Java is fun: " + isJavaFun);
        System.out.println("First letter: " + letter);
    }
}
```
### **🔹 3.2 - Giá trị của Kiểu Tham Chiếu (Reference Types)**

Kiểu tham chiếu lưu trữ địa chỉ của đối tượng trong bộ nhớ (heap).  
📌 Các loại kiểu tham chiếu:

- **Class**: `String`, `Integer`, `Double`, v.v.
- **Interface**: `List`, `Map`, v.v.
- **Array**: `int[]`, `String[]`
- **Null**: `null` là một giá trị đặc biệt đại diện cho **không có đối tượng nào cả**.

📌 **Ví dụ:**
```java
String name = "Java";
Integer number = 100;
int[] numbers = {1, 2, 3};
```

## **4️⃣ Biến (Variables) (JLS 4.12)**

**Biến** là một vùng nhớ lưu trữ giá trị và có một kiểu dữ liệu xác định.

### **🔹 4.1 - Các Loại Biến trong Java**

|Loại biến|Phạm vi|Khởi tạo mặc định|Ví dụ|
|---|---|---|---|
|**Biến cục bộ (Local variable)**|Trong phương thức hoặc khối lệnh|Không bắt buộc|`int x = 10;`|
|**Biến instance (Instance variable)**|Thuộc đối tượng|Có|`String name;`|
|**Biến static (Static variable)**|Thuộc lớp, chung cho tất cả đối tượng|Có|`static int count;`|

📌 **Ví dụ về các loại biến:**
```java
public class VariableExample {
    static int staticVar = 10; // Biến static

    int instanceVar = 20; // Biến instance

    public void method() {
        int localVar = 30; // Biến cục bộ
        System.out.println("Local: " + localVar);
    }

    public static void main(String[] args) {
        VariableExample obj = new VariableExample();
        obj.method();
        System.out.println("Instance: " + obj.instanceVar);
        System.out.println("Static: " + staticVar);
    }
}
```
## **5️⃣ Kiểu Tự Động Suy Luận (`var`) (JLS 4.10.2)**

Java 10 giới thiệu `var` để suy luận kiểu tự động dựa trên giá trị gán.

📌 **Ví dụ về `var`:**
```java
var message = "Hello, Java"; // Java tự suy luận kiểu là String
var number = 42; // Suy luận kiểu là int
System.out.println(message + " " + number);
```
💡 **Lưu ý:**

- `var` **chỉ được dùng trong phạm vi cục bộ**, không thể làm biến instance hoặc tham số phương thức.
- Java vẫn **kiểm tra kiểu tại compile-time**.
## **6️⃣ Phép Gán và Chuyển Đổi Kiểu (JLS 4.2.1, 4.2.2, 4.2.3)**

📌 **Chuyển đổi kiểu dữ liệu có 2 loại:**

1. **Chuyển đổi tường minh (Explicit Casting)**: Khi chuyển kiểu **hẹp hơn**.
2. **Chuyển đổi ngầm định (Implicit Casting)**: Khi chuyển kiểu **rộng hơn**.

### **🔹 Chuyển đổi ngầm định**

✔ Từ `int` → `long`, `float` → `double`
```java
int num = 10;
double d = num; // OK
```
### **🔹 Chuyển đổi tường minh**

✔ Từ `double` → `int` (mất dữ liệu)
```java
double pi = 3.14;
int intPi = (int) pi; // 3
```
📌 **Ví dụ về chuyển đổi kiểu**
```java
public class CastingExample {
    public static void main(String[] args) {
        int intValue = 100;
        double doubleValue = intValue; // Chuyển đổi ngầm định

        double pi = 3.14;
        int intPi = (int) pi; // Chuyển đổi tường minh

        System.out.println("Double from int: " + doubleValue);
        System.out.println("Int from double: " + intPi);
    }
}
```
## **7️⃣ Tổng Kết**

📌 **Bạn đã học được:**  
✅ Kiểu dữ liệu: Nguyên thủy và tham chiếu  
✅ Giá trị hợp lệ của từng kiểu dữ liệu  
✅ Các loại biến trong Java  
✅ Phép gán và chuyển đổi kiểu dữ liệu  
✅ Sử dụng `var` để suy luận kiểu

---
# BONUS

# **🔥 1. Chuyển đổi ngầm định (Implicit Casting)**

🔹 **Là quá trình chuyển đổi kiểu dữ liệu tự động từ kiểu nhỏ hơn sang kiểu lớn hơn** mà **không cần sự can thiệp của lập trình viên**.  
🔹 **Không bị mất dữ liệu** vì kiểu lớn hơn có thể chứa toàn bộ giá trị của kiểu nhỏ hơn.

💡 **Các quy tắc chuyển đổi ngầm định:**

- `byte` → `short` → `int` → `long` → `float` → `double`
- `char` → `int`, `long`, `float`, `double`

📌 **Ví dụ về chuyển đổi ngầm định:**
```java
public class ImplicitCasting {
    public static void main(String[] args) {
        int intVal = 100;
        double doubleVal = intVal; // Chuyển đổi từ int sang double

        char charVal = 'A'; // 'A' có giá trị ASCII là 65
        int asciiValue = charVal; // Chuyển đổi char thành int

        System.out.println("Double value: " + doubleVal); // 100.0
        System.out.println("ASCII value of 'A': " + asciiValue); // 65
    }
}
```
# **🔥 2. Chuyển đổi tường minh (Explicit Casting)**

🔹 **Là quá trình chuyển đổi từ kiểu dữ liệu lớn hơn sang kiểu dữ liệu nhỏ hơn**.  
🔹 Vì có **nguy cơ mất dữ liệu**, Java **bắt buộc** lập trình viên phải **tường minh** khi thực hiện chuyển đổi bằng cách sử dụng **cú pháp (kiểu mong muốn)**.

💡 **Các quy tắc chuyển đổi tường minh:**

- `double` → `float` → `long` → `int` → `short` → `byte`
- `float` → `int` (mất phần thập phân)
- `double` → `int` (mất phần thập phân)
- `long` → `int` (mất phần vượt quá phạm vi của `int`)

📌 **Ví dụ về chuyển đổi tường minh:**
```java
public class ExplicitCasting {
    public static void main(String[] args) {
        double pi = 3.14159;
        int intPi = (int) pi; // Chuyển từ double sang int (mất phần thập phân)

        long largeNumber = 100000L;
        int smallNumber = (int) largeNumber; // Chuyển từ long sang int

        System.out.println("Int value of Pi: " + intPi); // 3
        System.out.println("Converted long to int: " + smallNumber); // 100000
    }
}
```
# **📌 3. So sánh Chuyển đổi Ngầm định và Tường minh**

| **Thuộc tính**       | **Chuyển đổi Ngầm định**                  | **Chuyển đổi Tường minh**                 |
| -------------------- | ----------------------------------------- | ----------------------------------------- |
| **Hướng chuyển đổi** | Từ kiểu **nhỏ hơn** sang kiểu **lớn hơn** | Từ kiểu **lớn hơn** sang kiểu **nhỏ hơn** |
| **Cách sử dụng**     | Java tự động thực hiện                    | Lập trình viên **phải chỉ định rõ ràng**  |
| **Mất dữ liệu**      | Không                                     | Có thể bị mất dữ liệu                     |
| **Ví dụ**            | `int → double`                            | `double → int`                            |
# **🔥 4. Chuyển đổi giữa các kiểu dữ liệu không tương thích**

⛔ Java **không hỗ trợ** chuyển đổi trực tiếp giữa các kiểu không tương thích (ví dụ: `boolean` ↔ `int`, `String` ↔ `int`).  
📌 Nếu muốn chuyển đổi, ta cần dùng phương thức hỗ trợ như:

- `Integer.parseInt(String)`
- `Double.parseDouble(String)`
- `String.valueOf(int)`

📌 **Ví dụ:**
```java
public class TypeConversion {
    public static void main(String[] args) {
        String str = "123";
        int number = Integer.parseInt(str); // Chuyển String thành int
        System.out.println("Converted number: " + number);

        int num = 456;
        String strNum = String.valueOf(num); // Chuyển int thành String
        System.out.println("Converted String: " + strNum);
    }
}
```
