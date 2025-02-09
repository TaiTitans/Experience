
---
*Thông thạo Java và OOP đồng nghĩa các ngôn ngữ lập trình khác không phải là vấn đề.*

# **📌 1. Tổng quan về Execution trong Java**

Quá trình thực thi một chương trình Java bao gồm các bước:

1. **Khởi động (Startup)**
2. **Tạo và thực thi luồng chính (Main Thread Execution)**
3. **Tạo đối tượng và gọi phương thức**
4. **Xử lý ngoại lệ**
5. **Dừng chương trình (Shutdown)**

# **📌 2. Khởi động chương trình Java (Program Startup)**

Java bắt đầu thực thi bằng cách gọi phương thức `main` trong class được chỉ định.

### **🔹 Cấu trúc của phương thức `main`**
```java
public class MainExample {
    public static void main(String[] args) {
        System.out.println("Hello, Java!");
    }
}
```
📌 **Quy tắc của phương thức `main`:**  
✅ `public`: Để JVM có thể gọi được từ bên ngoài.  
✅ `static`: Không cần tạo đối tượng để gọi.  
✅ `void`: Không trả về giá trị nào.  
✅ `String[] args`: Nhận tham số dòng lệnh.
# **📌 3. Luồng chính và tạo luồng mới**

Sau khi khởi động, JVM tạo **Main Thread**, đây là luồng chính của chương trình.

📌 **Ví dụ về Main Thread:**
```java
public class MainThreadExample {
    public static void main(String[] args) {
        System.out.println("Thread hiện tại: " + Thread.currentThread().getName());
    }
}
```
✅ **Main Thread có thể tạo thêm luồng mới:**
```java
public class MultiThreadExample {
    public static void main(String[] args) {
        Thread thread = new Thread(() -> System.out.println("Thread mới: " + Thread.currentThread().getName()));
        thread.start();
    }
}
```
# **📌 4. Tạo đối tượng và gọi phương thức**

### **🔹 Quy trình khởi tạo đối tượng**

1. **Cấp phát bộ nhớ**
2. **Gọi constructor**
3. **Gán giá trị mặc định cho biến**

📌 **Ví dụ khởi tạo đối tượng:**
```java
class Car {
    String model = "Toyota";

    Car() {
        System.out.println("Constructor được gọi!");
    }
}

public class ObjectExample {
    public static void main(String[] args) {
        Car car = new Car();
        System.out.println("Mẫu xe: " + car.model);
    }
}
```
# **📌 5. Xử lý ngoại lệ trong quá trình thực thi**

Nếu xảy ra ngoại lệ mà không được xử lý, JVM sẽ dừng chương trình.

📌 **Ví dụ:**
```java
public class ExceptionExample {
    public static void main(String[] args) {
        int result = 5 / 0; // ArithmeticException
    }
}
```
✅ **Cách xử lý ngoại lệ:**
```java
public class HandleExceptionExample {
    public static void main(String[] args) {
        try {
            int result = 5 / 0;
        } catch (ArithmeticException e) {
            System.out.println("Lỗi: Chia cho 0!");
        }
    }
}
```
# **📌 6. Dừng chương trình (Program Termination)**

Chương trình có thể dừng tự nhiên hoặc ép dừng bằng **`System.exit(int)`**.

📌 **Ví dụ:**
```java
public class ExitExample {
    public static void main(String[] args) {
        System.out.println("Trước khi thoát");
        System.exit(0); // Thoát chương trình
        System.out.println("Dòng này sẽ không chạy");
    }
}
```
✅ **Code sau `System.exit()` sẽ không được thực thi**.
# **📌 7. Tổng kết**

✅ Java thực thi chương trình bằng **Main Thread**.  
✅ JVM gọi phương thức `main` để bắt đầu chương trình.  
✅ Chương trình có thể tạo nhiều luồng con.  
✅ Khi gặp ngoại lệ không xử lý, JVM sẽ dừng chương trình.  
✅ Có thể dùng `System.exit(0)` để kết thúc chương trình thủ công.