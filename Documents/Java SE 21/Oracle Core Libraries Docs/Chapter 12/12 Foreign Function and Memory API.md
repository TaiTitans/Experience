
---
API **Foreign Function and Memory (FFM)** trong Java SE 21 cho phép các chương trình Java tương tác với mã và dữ liệu bên ngoài môi trường runtime của Java. Điều này giúp các chương trình Java có thể gọi các thư viện native và xử lý dữ liệu native mà không cần sử dụng Java Native Interface (JNI), vốn phức tạp và dễ gây lỗi.

[Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/foreign-function-and-memory-api.html?utm_source=chatgpt.com)

**Các khái niệm chính trong FFM API:**

- **Memory Segment:** Đại diện cho một vùng nhớ liên tục, có thể nằm trong hoặc ngoài heap của Java. Memory Segment cung cấp cách truy cập an toàn và hiệu quả vào bộ nhớ ngoại, cho phép thao tác với dữ liệu không được quản lý bởi JVM.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/lang/foreign/package-summary.html?utm_source=chatgpt.com)
    
- **Memory Layout:** Mô tả cấu trúc của một Memory Segment, bao gồm kích thước, độ lệch (offset), và các ràng buộc về căn chỉnh (alignment). Memory Layout giúp định nghĩa cách dữ liệu được sắp xếp trong bộ nhớ, hỗ trợ truy cập có cấu trúc và an toàn.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/lang/foreign/package-summary.html?utm_source=chatgpt.com)
    
- **Linker:** Cung cấp khả năng liên kết với các hàm ngoại (foreign functions), cho phép gọi các hàm từ thư viện native. Linker chịu trách nhiệm ánh xạ các hàm Java tới các hàm native tương ứng, quản lý các chi tiết liên quan đến ABI (Application Binary Interface).
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/calling-c-library-function-foreign-function-and-memory-api.html?utm_source=chatgpt.com)
    
- **Symbol Lookup:** Cho phép tìm kiếm các hàm hoặc biến trong các thư viện native theo tên, hỗ trợ quá trình liên kết động với các thành phần ngoại.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/calling-c-library-function-foreign-function-and-memory-api.html?utm_source=chatgpt.com)
    

**Sử dụng FFM API:**

Để sử dụng FFM API, bạn cần bật tính năng preview trong quá trình biên dịch và chạy chương trình, vì đây vẫn là một tính năng đang trong giai đoạn xem trước. Điều này có thể được thực hiện bằng cách sử dụng các tùy chọn dòng lệnh bổ sung khi biên dịch và chạy mã chứa các tính năng preview.

[Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/foreign-function-and-memory-api.html?utm_source=chatgpt.com)

**Ví dụ: Gọi hàm `strlen` từ thư viện C:**

Dưới đây là một ví dụ về cách sử dụng FFM API để gọi hàm `strlen` từ thư viện C:
```java
static long invokeStrlen(String s) throws Throwable {
    try (Arena arena = Arena.ofConfined()) {
        // Cấp phát bộ nhớ ngoài heap và sao chép chuỗi Java vào bộ nhớ này
        MemorySegment nativeString = arena.allocateUtf8String(s);
        
        // Liên kết và gọi hàm C strlen
        Linker linker = Linker.nativeLinker();
        SymbolLookup stdLib = linker.defaultLookup();
        MemorySegment strlenAddr = stdLib.find("strlen").get();
        FunctionDescriptor strlenSig = FunctionDescriptor.of(ValueLayout.JAVA_LONG, ValueLayout.ADDRESS);
        MethodHandle strlen = linker.downcallHandle(strlenAddr, strlenSig);
        
        // Gọi hàm C trực tiếp từ Java
        return (long) strlen.invokeExact(nativeString);
    }
}
```
Trong ví dụ này, chúng ta sử dụng `Arena` để quản lý bộ nhớ ngoài heap, `Linker` và `SymbolLookup` để tìm và liên kết với hàm `strlen` trong thư viện C, và `MethodHandle` để gọi hàm này từ Java.

[Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/calling-c-library-function-foreign-function-and-memory-api.html?utm_source=chatgpt.com)

**Lưu ý:**

Việc sử dụng FFM API yêu cầu hiểu biết sâu về cả Java và ngôn ngữ native mà bạn đang tương tác, cũng như cách quản lý bộ nhớ và tài nguyên trong môi trường native. Điều này giúp đảm bảo rằng việc tương tác giữa Java và mã native diễn ra an toàn và hiệu quả.

Để biết thêm chi tiết, bạn có thể tham khảo tài liệu chính thức về [Foreign Function and Memory API](https://docs.oracle.com/en/java/javase/21/core/foreign-function-and-memory-api.html).