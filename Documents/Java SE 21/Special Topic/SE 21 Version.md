
---
Java 21 là một phiên bản LTS (Long-Term Support) mới nhất của Java, mang lại nhiều cải tiến so với các phiên bản trước. Dưới đây là những điểm khác biệt chính của Java 21 so với các phiên bản trước:
## 🔥 **Các tính năng mới và cải tiến trong Java 21**

### 1️⃣ **Virtual Threads (Luồng ảo) - JEP 444 (Stable)**

- **Mô tả**: Virtual Threads giúp đơn giản hóa lập trình đa luồng bằng cách tạo ra hàng triệu luồng nhẹ mà không tốn quá nhiều tài nguyên.
- **Lợi ích**:
    - Giảm chi phí tạo và quản lý luồng.
    - Tối ưu hiệu suất cho các ứng dụng sử dụng nhiều I/O (như web servers, microservices).
    - Không cần dùng thread pools phức tạp như trước đây.
- **So sánh với Java 8-17**: Trước đây, các luồng được quản lý bởi OS (Platform Threads), tốn tài nguyên và không thể mở rộng tốt.
### 2️⃣ **Pattern Matching for Switch (Khớp mẫu với `switch`) - JEP 441 (Stable)**

- **Mô tả**: Cải tiến cú pháp `switch`, hỗ trợ `instanceof` và kiểu mẫu (`pattern matching`).
```java
static String test(Object obj) {
    return switch (obj) {
        case Integer i -> "Số nguyên: " + i;
        case String s when s.length() > 5 -> "Chuỗi dài: " + s;
        case null -> "Giá trị null";
        default -> "Không xác định";
    };
}
```
**Lợi ích**:

- Viết mã ngắn gọn hơn.
- Tăng cường kiểm tra kiểu (`type safety`).
- Tránh lỗi `NullPointerException`.

### 3️⃣ **Unnamed Classes and Instance Main Methods - JEP 445 (Preview)**

- **Mô tả**: Viết Java mà không cần khai báo lớp (`class`) rõ ràng, giúp code đơn giản hơn, nhất là cho lập trình viên mới.
- **Ví dụ** (Chạy được trực tiếp mà không cần `public class Main`):
```java
void main() {
    System.out.println("Hello, Java 21!");
}
```
- **So sánh với Java 8-17**: Trước đây bắt buộc phải khai báo `public static void main(String[] args)`.

### 4️⃣ **Scoped Values (Giá trị phạm vi) - JEP 446 (Preview)**

- **Mô tả**: Thay thế `ThreadLocal` để truyền dữ liệu hiệu quả hơn giữa các luồng.
- **Lợi ích**:
    - Giảm chi phí bộ nhớ.
    - Tương thích tốt hơn với Virtual Threads.
- **Ví dụ**:
```java
static final ScopedValue<String> USERNAME = ScopedValue.newInstance();

void process() {
    ScopedValue.where(USERNAME, "Alice").run(() -> {
        System.out.println("Hello " + USERNAME.get());
    });
}
```
### 5️⃣ **Sequenced Collections (Danh sách tuần tự) - JEP 431 (Stable)**

- **Mô tả**: Java 21 thêm interface `SequencedCollection`, `SequencedSet`, `SequencedMap` giúp quản lý danh sách tuần tự dễ hơn.
- **Ví dụ**:
```java
SequencedCollection<String> list = new LinkedList<>();
list.addLast("A");
list.addLast("B");
list.addFirst("C");
System.out.println(list); // [C, A, B]
```
**So sánh với Java 8-17**: Trước đây, không có API chung để truy cập phần tử đầu/cuối danh sách.
### 6️⃣ **Foreign Function & Memory API - JEP 454 (Stable)**

- **Mô tả**: Giao tiếp trực tiếp với bộ nhớ ngoài (`native memory`) mà không cần JNI (`Java Native Interface`).
- **Lợi ích**:
    - Hiệu suất cao hơn.
    - An toàn hơn so với JNI.
- **Ví dụ**:
```java
try (Arena arena = Arena.ofConfined()) {
    MemorySegment segment = arena.allocate(100);
    segment.set(ValueLayout.JAVA_INT, 0, 42);
    int value = segment.get(ValueLayout.JAVA_INT, 0);
    System.out.println(value); // 42
}
```
### 7️⃣ **Record Patterns (Cải tiến `record`) - JEP 440 (Stable)**

- **Mô tả**: Mở rộng `record` để hỗ trợ `pattern matching` tốt hơn.
- **Ví dụ**:
```java
record Point(int x, int y) {}

static void print(Point p) {
    if (p instanceof Point(int a, int b)) {
        System.out.println("X: " + a + ", Y: " + b);
    }
}
```
## ⚡ **So sánh Java 21 với các phiên bản trước**

|Tính năng|Java 8|Java 11|Java 17|Java 21|
|---|---|---|---|---|
|**LTS**|✅|✅|✅|✅|
|**Lambda & Streams**|✅|✅|✅|✅|
|**Records (`record`)**|❌|❌|✅|✅|
|**Sealed Classes**|❌|❌|✅|✅|
|**Virtual Threads**|❌|❌|❌|✅|
|**Pattern Matching for `switch`**|❌|❌|✅ (preview)|✅ (stable)|
|**Foreign Function API**|❌|❌|✅ (preview)|✅ (stable)|
|**Scoped Values**|❌|❌|❌|✅ (preview)|