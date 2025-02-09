
---
`Focus on improving, not proving.`
# **📌 1. Tổng quan về Exception Handling**

Trong Java, **ngoại lệ (Exception)** là một sự kiện xảy ra trong quá trình thực thi chương trình làm gián đoạn luồng thực thi bình thường. Java cung cấp cơ chế để **bắt (catch), xử lý (handle) và khôi phục (recover)** từ lỗi thông qua hệ thống **try-catch-finally**.

### **🔹 Có hai loại lỗi chính trong Java:**
|**Loại lỗi**|**Mô tả**|
|---|---|
|**Error**|Những lỗi nghiêm trọng không thể khắc phục được như OutOfMemoryError, StackOverflowError.|
|**Exception**|Những lỗi có thể xử lý được như NullPointerException, IOException, ArithmeticException.|

📌 **Ví dụ về lỗi không thể xử lý (`Error`):**
```java
public class TestError {
    public static void main(String[] args) {
        main(args); // Gọi đệ quy vô hạn → StackOverflowError
    }
}
```
📌 **Ví dụ về lỗi có thể xử lý (`Exception`):**
```java
public class TestException {
    public static void main(String[] args) {
        int a = 5 / 0; // ArithmeticException
    }
}
```
# **📌 2. Exception Hierarchy (Phân cấp ngoại lệ trong Java)**

Tất cả các ngoại lệ trong Java đều kế thừa từ `Throwable`.
```lua
            Throwable
           /         \
        Error      Exception
                    /      \
   RuntimeException  IOException (và các loại khác)
```
✅ **`Error`**: Lỗi nghiêm trọng (không bắt được).  
✅ **`Exception`**: Lỗi có thể xử lý được.  
✅ **`RuntimeException`**: Ngoại lệ xảy ra trong runtime (Unchecked Exception).
### **🔹 Checked vs Unchecked Exceptions**

|**Loại**|**Mô tả**|**Ví dụ**|
|---|---|---|
|**Checked Exception**|Bắt buộc phải xử lý (compile-time)|IOException, SQLException|
|**Unchecked Exception**|Không bắt buộc xử lý (runtime)|NullPointerException, ArithmeticException|

📌 **Ví dụ về Checked Exception (phải xử lý):**
```java
import java.io.File;
import java.io.FileReader;
import java.io.IOException;

public class CheckedExceptionExample {
    public static void main(String[] args) {
        try {
            FileReader fr = new FileReader(new File("file.txt"));
        } catch (IOException e) {
            System.out.println("File not found!");
        }
    }
}
```
📌 **Ví dụ về Unchecked Exception (không bắt buộc xử lý):**
```java
public class UncheckedExceptionExample {
    public static void main(String[] args) {
        int[] arr = new int[3];
        System.out.println(arr[5]); // ArrayIndexOutOfBoundsException
    }
}
```
# **📌 3. Cơ chế Try-Catch-Finally**

Java cung cấp 3 khối chính để xử lý ngoại lệ:

### **🔹 try-catch**

Cú pháp:
```java
try {
    // Code có thể gây lỗi
} catch (ExceptionType e) {
    // Xử lý lỗi
}
```
📌 **Ví dụ:**
```java
public class TryCatchExample {
    public static void main(String[] args) {
        try {
            int a = 5 / 0;
        } catch (ArithmeticException e) {
            System.out.println("Lỗi chia cho 0");
        }
    }
}
```
### **🔹 finally**

`finally` luôn chạy dù có ngoại lệ hay không.

📌 **Ví dụ:**
```java
public class FinallyExample {
    public static void main(String[] args) {
        try {
            int a = 5 / 0;
        } catch (ArithmeticException e) {
            System.out.println("Lỗi chia cho 0");
        } finally {
            System.out.println("Luôn chạy finally");
        }
    }
}
```
### **🔹 try-with-resources**

Dùng cho `AutoCloseable` như `FileReader`, `BufferedReader`.

📌 **Ví dụ:**
```java
import java.io.FileReader;
import java.io.IOException;

public class TryWithResourcesExample {
    public static void main(String[] args) {
        try (FileReader fr = new FileReader("file.txt")) {
            // Đọc file
        } catch (IOException e) {
            System.out.println("Lỗi đọc file");
        }
    }
}
```
# **📌 4. Tạo Custom Exception**

Chúng ta có thể tạo ngoại lệ riêng bằng cách kế thừa `Exception` hoặc `RuntimeException`.

📌 **Ví dụ:**
```java
class MyException extends Exception {
    public MyException(String message) {
        super(message);
    }
}

public class CustomExceptionExample {
    public static void main(String[] args) throws MyException {
        throw new MyException("Lỗi tùy chỉnh!");
    }
}
```
# **📌 5. Throw vs Throws
|**Từ khóa**|**Mô tả**|
|---|---|
|`throw`|Ném một ngoại lệ cụ thể trong code|
|`throws`|Khai báo phương thức có thể ném ngoại lệ|

📌 **Ví dụ về `throw`:**
```java
public class ThrowExample {
    public static void main(String[] args) {
        throw new ArithmeticException("Lỗi chia cho 0");
    }
}
```
📌 **Ví dụ về `throws`:**
```java
public class ThrowsExample {
    public static void method() throws IOException {
        throw new IOException("Lỗi đọc file");
    }
}
```
# **📌 6. Chặn Exception trong kế thừa**

Nếu một phương thức `throws` một ngoại lệ, phương thức override có thể:

- **Không khai báo ngoại lệ**
- **Khai báo cùng ngoại lệ**
- **Khai báo ngoại lệ con**

📌 **Ví dụ hợp lệ:**
```java
class Parent {
    void show() throws IOException {}
}

class Child extends Parent {
    void show() throws FileNotFoundException {} // OK (FileNotFoundException là con của IOException)
}
```
📌 **Ví dụ không hợp lệ:**
```java
class Parent {
    void show() throws IOException {}
}

class Child extends Parent {
    void show() throws Exception {} // LỖI: Exception rộng hơn IOException
}
```
# **📌 7. Lưu ý quan trọng**

✅ **Luôn dùng `try-with-resources` để tránh rò rỉ tài nguyên**.  
✅ **Không nên catch Exception chung chung (`catch (Exception e)`) trừ khi thật cần thiết**.  
✅ **Chỉ bắt những ngoại lệ mà bạn có thể xử lý**.  
✅ **Ghi log ngoại lệ để dễ debug**.

📌 **Ví dụ log lỗi đúng cách:**
```java
import java.util.logging.Level;
import java.util.logging.Logger;

public class LoggingExample {
    private static final Logger LOGGER = Logger.getLogger(LoggingExample.class.getName());

    public static void main(String[] args) {
        try {
            int a = 5 / 0;
        } catch (ArithmeticException e) {
            LOGGER.log(Level.SEVERE, "Lỗi chia cho 0", e);
        }
    }
}
```
# **📌 8. Tổng kết**

✅ **Exception là lỗi có thể xử lý, Error là lỗi không thể khắc phục**.  
✅ **Có Checked Exception (phải xử lý) và Unchecked Exception (không bắt buộc xử lý)**.  
✅ **Cơ chế xử lý: try-catch-finally, throw, throws**.  
✅ **Có thể tạo Custom Exception**.  
✅ **Dùng try-with-resources để tránh rò rỉ tài nguyên**.