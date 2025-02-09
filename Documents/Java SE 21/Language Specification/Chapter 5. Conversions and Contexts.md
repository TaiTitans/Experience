
---
### **🔥 1. Giới thiệu về Chuyển đổi và Phát sinh kiểu**

Trong Java, **chuyển đổi kiểu dữ liệu (type conversion)** xảy ra khi giá trị của một kiểu dữ liệu được chuyển đổi thành kiểu khác. Chuyển đổi có thể xảy ra **tường minh (explicit casting)** hoặc **ngầm định (implicit casting)**.

Ngoài ra, Java còn có **các ngữ cảnh phát sinh kiểu (contexts for type inference)**, nơi trình biên dịch có thể suy luận kiểu một cách tự động.

### **🔥 2. Các loại chuyển đổi trong Java**

Chuyển đổi kiểu trong Java được chia thành **13 loại chính**, trong đó quan trọng nhất gồm:

1. **Chuyển đổi mở rộng (Widening Conversion)** ✅
    
    - Chuyển từ **kiểu nhỏ hơn** sang **kiểu lớn hơn** (không mất dữ liệu).
    - Ví dụ: `int → long`, `float → double`.
2. **Chuyển đổi thu hẹp (Narrowing Conversion)** ❌
    
    - Chuyển từ **kiểu lớn hơn** sang **kiểu nhỏ hơn** (có thể mất dữ liệu).
    - Ví dụ: `double → int`, `long → short`.
3. **Chuyển đổi tham chiếu (Reference Conversion)**
    
    - Chuyển đổi giữa các kiểu **lớp cha - lớp con** (theo kế thừa).
    - Ví dụ: `Animal → Dog`, `List → ArrayList`.
4. **Chuyển đổi giá trị Boxing và Unboxing**
    
    - **Boxing**: Chuyển kiểu nguyên thủy thành wrapper (ví dụ: `int → Integer`).
    - **Unboxing**: Chuyển wrapper thành kiểu nguyên thủy (ví dụ: `Integer → int`).
5. **Chuyển đổi không tương thích (Unchecked Conversion)**
    
    - Liên quan đến Generics (chuyển từ `List` sang `List<String>`).
6. **Chuyển đổi phát sinh kiểu (Inference Conversion)**
    
    - Dùng trong lambda, generic methods, `var`.
### **🔥 3. Chi tiết về từng loại chuyển đổi**

#### **3.1. Chuyển đổi mở rộng (Widening Conversion)**

✅ **Được thực hiện tự động**  
✅ **Không mất dữ liệu**  
🔹 **Các quy tắc:**

- `byte → short → int → long → float → double`
- `char → int`, `long`, `float`, `double`

📌 **Ví dụ:**
```java
public class WideningExample {
    public static void main(String[] args) {
        int num = 100;
        double d = num; // Chuyển đổi int sang double (tự động)

        char c = 'A';
        int ascii = c; // Chuyển đổi char thành int

        System.out.println("Double value: " + d); // 100.0
        System.out.println("ASCII of 'A': " + ascii); // 65
    }
}
```

#### **3.2. Chuyển đổi thu hẹp (Narrowing Conversion)**

❌ **Không được thực hiện tự động**  
❌ **Có thể mất dữ liệu hoặc gây lỗi**  
🔹 **Cần ép kiểu tường minh (Explicit Casting)**

📌 **Ví dụ:**
```java
public class NarrowingExample {
    public static void main(String[] args) {
        double d = 9.7;
        int num = (int) d; // Ép kiểu từ double xuống int (mất phần thập phân)

        long bigNum = 100000L;
        short smallNum = (short) bigNum; // Có thể bị mất dữ liệu

        System.out.println("Converted int: " + num); // 9
        System.out.println("Converted short: " + smallNum); // 100000 (hoặc lỗi nếu quá phạm vi)
    }
}
```
#### **3.3. Chuyển đổi tham chiếu (Reference Conversion)**

🔹 **Chỉ áp dụng cho kiểu đối tượng (Objects)**  
🔹 **Dựa trên quan hệ kế thừa hoặc giao diện**

📌 **Ví dụ (Lớp cha → Lớp con):**
```java
class Animal {}
class Dog extends Animal {}

public class ReferenceConversion {
    public static void main(String[] args) {
        Animal a = new Dog(); // Chuyển đổi ngầm định (Upcasting)
        Dog d = (Dog) a; // Ép kiểu tường minh (Downcasting)
    }
}
```
📌 **Ví dụ (Interface):**
```java
interface Vehicle {}
class Car implements Vehicle {}

public class InterfaceConversion {
    public static void main(String[] args) {
        Vehicle v = new Car(); // Upcasting
        Car c = (Car) v; // Downcasting
    }
}
```
#### **3.4. Chuyển đổi giữa kiểu nguyên thủy và Wrapper (Boxing & Unboxing)**

🔹 **Boxing:** `int → Integer` (Tự động)  
🔹 **Unboxing:** `Integer → int` (Tự động)

📌 **Ví dụ:**
```java
public class BoxingExample {
    public static void main(String[] args) {
        Integer obj = 10; // Boxing
        int num = obj; // Unboxing

        System.out.println("Boxed: " + obj);
        System.out.println("Unboxed: " + num);
    }
}
```
#### **3.5. Chuyển đổi không tương thích (Unchecked Conversion)**

⛔ **Khi dùng Generics, Java không tự động ép kiểu**  
📌 **Ví dụ:**
```java
import java.util.*;

public class UncheckedConversion {
    public static void main(String[] args) {
        List rawList = new ArrayList(); // Không có generic type
        rawList.add("Hello");
        List<String> stringList = rawList; // Unchecked conversion (có cảnh báo)
    }
}
```
#### **3.6. Chuyển đổi phát sinh kiểu (Inference Conversion)**

🔹 **Dùng trong Generics, Lambda, `var`**  
📌 **Ví dụ (Suy luận kiểu trong Generics):**
```java
public class TypeInferenceExample {
    public static void main(String[] args) {
        List<String> list = List.of("Java", "Spring");
        var myList = list; // Java tự suy luận kiểu List<String>

        System.out.println(myList);
    }
}
```
📌 **Ví dụ (Lambda Expression):**
```java
import java.util.function.Function;

public class LambdaExample {
    public static void main(String[] args) {
        Function<Integer, String> func = (num) -> "Number: " + num;
        System.out.println(func.apply(10));
    }
}
```
### **🔥 4. Kết luận**

✅ Java có nhiều loại chuyển đổi kiểu dữ liệu, chủ yếu gồm **chuyển đổi mở rộng, thu hẹp, tham chiếu, boxing/unboxing**.  
✅ Một số chuyển đổi xảy ra **tự động**, một số cần **ép kiểu tường minh**.  
✅ Các ngữ cảnh như **Generics, Lambda, `var`** hỗ trợ **phát sinh kiểu tự động**.

