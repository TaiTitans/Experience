
---
Chương 5 của **Java Virtual Machine Specification (JVM Spec) SE 21** mô tả **quá trình tải lớp (Class Loading) và liên kết (Linking) trong JVM**. Đây là một trong những bước quan trọng trong quá trình thực thi chương trình Java.

# 🔥 **1. Quá trình thực thi một class trong JVM**

JVM thực thi một class bằng cách trải qua **5 giai đoạn chính**:

1️⃣ **Loading (Tải lớp)**  
2️⃣ **Linking (Liên kết)**

- Verification (Kiểm tra)
- Preparation (Chuẩn bị)
- Resolution (Giải quyết tham chiếu)  
    3️⃣ **Initialization (Khởi tạo)**  
    4️⃣ **Execution (Thực thi - chạy bytecode)**  
    5️⃣ **Unloading (Giải phóng bộ nhớ, nếu cần)**
# 🔹 **2. Loading - Tải lớp**

Khi một class được sử dụng, JVM cần **nạp (load)** nó vào bộ nhớ từ `.class` file.

📌 **Các Class Loader quan trọng**:

|Class Loader|Mô tả|
|---|---|
|**Bootstrap ClassLoader**|Tải các class cốt lõi của Java (rt.jar, java.base)|
|**Extension ClassLoader**|Tải các thư viện mở rộng từ `ext/`|
|**Application ClassLoader**|Tải các class từ đường dẫn ứng dụng (`classpath`)|
|**Custom ClassLoader**|ClassLoader do lập trình viên tự định nghĩa|

📌 **Ví dụ:**  
Khi ta chạy:
```java
public class Main {
    public static void main(String[] args) {
        System.out.println("Hello, JVM!");
    }
}
```
JVM thực hiện: 1️⃣ **Bootstrap ClassLoader** tải `java.lang.Object`, `java.lang.String`, v.v.  
2️⃣ **Application ClassLoader** tải `Main.class`.

Bạn có thể kiểm tra class loader bằng:
```java
System.out.println(Main.class.getClassLoader());  // Application ClassLoader
System.out.println(String.class.getClassLoader());  // null (Bootstrap ClassLoader)
```
# 🔹 **3. Linking - Liên kết**

Sau khi một class được tải, JVM sẽ liên kết nó. Quá trình này có **3 bước**:

## 🔸 (1) Verification (Kiểm tra) 🔍

- JVM kiểm tra **tính hợp lệ của bytecode** để đảm bảo nó không có lỗi hoặc mã độc.
- Nếu có lỗi, JVM báo `java.lang.VerifyError`.

📌 **Các kiểm tra chính**: ✅ Kiểm tra định dạng `.class` file hợp lệ (magic number, version, v.v.)  
✅ Kiểm tra kiểu dữ liệu hợp lệ (ví dụ: không thể ép kiểu `int` thành `String`)  
✅ Kiểm tra truy cập hợp lệ (private, protected, public)
## 🔸 (2) Preparation (Chuẩn bị) 🏗️

- JVM cấp phát bộ nhớ cho **biến static**, gán giá trị mặc định (`0`, `null`, `false`, v.v.).
- **Chưa gọi đến giá trị khởi tạo do lập trình viên gán!**
```java
class A {
    static int x = 10;
}
```
💡 Ở bước **Preparation**, `x` được cấp phát bộ nhớ và gán **mặc định là `0`**.  
💡 Sau đó, ở bước **Initialization**, `x` mới được gán giá trị `10`.

## 🔸 (3) Resolution (Giải quyết tham chiếu) 🔗

- JVM thay thế các tham chiếu **biểu tượng (symbolic reference)** trong constant pool bằng **địa chỉ bộ nhớ thực tế**.
- Điều này áp dụng cho:
    - Class & Interface (chuyển `"java/lang/String"` → `String.class`)
    - Field & Method (chuyển `#3 (println)` → `System.out.println`)

📌 **Ví dụ**:
`System.out.println("Hello");`

1️⃣ **Constant Pool** chứa `MethodRef #3 (println)`  
2️⃣ **Resolution** sẽ ánh xạ `#3` thành **địa chỉ thực tế của `System.out.println` trong RAM**.

# 🔹 **4. Initialization - Khởi tạo** 🚀

Đây là lúc **JVM thực thi code khởi tạo class**:

- Gán giá trị cho các biến `static`
- Gọi **static block**
- Gọi **hàm khởi tạo (constructor) nếu là instance**

📌 **Ví dụ**:
```java
class A {
    static int x = 10;
    static { 
        System.out.println("Static block chạy!"); 
        x += 5;
    }
}
```
💡 **Output khi class được dùng lần đầu tiên**:
```
Static block chạy!
```
💡 Lúc này, `x = 15`.

# 🔹 **5. Execution - Thực thi bytecode** 🏃

Sau khi class được khởi tạo, JVM chạy **main method**:
```java
public static void main(String[] args) {
    System.out.println("Hello, JVM!");
}
```
JVM thực thi từng lệnh bytecode bằng **Java Interpreter hoặc JIT Compiler**.

Bạn có thể xem bytecode bằng:
```
javap -c Main.class
```
Ví dụ output:
```java
0: getstatic     #2 (System.out)
3: ldc           #3 ("Hello, JVM!")
5: invokevirtual #4 (println)
```
# 🔹 **6. Unloading - Giải phóng bộ nhớ**

JVM có thể **xóa class khỏi bộ nhớ** khi:

- ClassLoader chứa nó bị **GC thu hồi**.
- Nó là một **dynamically loaded class** mà không còn reference.

Tuy nhiên, **Bootstrap ClassLoader** thường không unload class.

📌 **Ví dụ unload class:**
```java
ClassLoader myLoader = new MyClassLoader();
Class<?> dynamicClass = myLoader.loadClass("MyClass");
dynamicClass = null;
System.gc(); // JVM có thể unload MyClass nếu không còn tham chiếu
```
# 🚀 **Tóm tắt**

| Giai đoạn          | Mô tả                                                  |
| ------------------ | ------------------------------------------------------ |
| **Loading**        | ClassLoader tải `.class` vào bộ nhớ                    |
| **Verification**   | JVM kiểm tra tính hợp lệ của bytecode                  |
| **Preparation**    | Cấp phát bộ nhớ và gán giá trị mặc định                |
| **Resolution**     | Biến đổi tham chiếu **biểu tượng** thành địa chỉ thật  |
| **Initialization** | Chạy **static block** và gán giá trị cho static fields |
| **Execution**      | Chạy bytecode bằng JVM Interpreter / JIT Compiler      |
| **Unloading**      | JVM giải phóng class nếu cần                           |