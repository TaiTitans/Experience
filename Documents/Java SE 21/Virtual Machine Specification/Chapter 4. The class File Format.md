
---
Chương 4 của **Java Virtual Machine Specification (JVM Spec) SE 21** mô tả định dạng của file **`.class`**, hay còn gọi là **Class File Format**. Đây là định dạng mà mã Java được biên dịch thành trước khi JVM thực thi.

## 🔹 1. Tổng quan về Class File Format

Một file `.class` chứa **bytecode**, là tập hợp các lệnh được JVM hiểu và thực thi. File này có cấu trúc **nhị phân**, không thể đọc trực tiếp như mã Java.

Mọi file `.class` đều có cấu trúc theo thứ tự sau:

| Thành phần        | Kích thước | Chức năng                                                |
| ----------------- | ---------- | -------------------------------------------------------- |
| **Magic Number**  | 4 byte     | Xác định đây là file `.class` hợp lệ                     |
| **Version**       | 4 byte     | Phiên bản của JVM hỗ trợ file này                        |
| **Constant Pool** | Biến đổi   | Bảng hằng số chứa thông tin tên, kiểu, giá trị, v.v.     |
| **Access Flags**  | 2 byte     | Quyền truy cập của class (public, final, abstract, v.v.) |
| **This Class**    | 2 byte     | Chỉ mục của class hiện tại trong constant pool           |
| **Super Class**   | 2 byte     | Chỉ mục của class cha trong constant pool                |
| **Interfaces**    | Biến đổi   | Danh sách các interface mà class này implement           |
| **Fields**        | Biến đổi   | Danh sách các thuộc tính (fields) của class              |
| **Methods**       | Biến đổi   | Danh sách các phương thức (methods) của class            |
| **Attributes**    | Biến đổi   | Các thông tin bổ sung như annotations, generic, v.v.     |
## 🔹 2. Chi tiết từng thành phần

### 🟢 **Magic Number**

- Mỗi file `.class` bắt đầu bằng **4 byte cố định** có giá trị **0xCAFEBABE**.
- Mục đích: Giúp JVM nhận diện đây là file `.class` hợp lệ.
- Nếu file không bắt đầu bằng `CAFEBABE`, JVM sẽ báo lỗi.

📌 **Ví dụ** (hexadecimal format của file `.class`):
```
CA FE BA BE 00 00 00 37 ...
```
(`00 37` là phiên bản của JVM – tương đương Java 21)
### 🟢 **Version (Major & Minor Version)**

- **4 byte** tiếp theo xác định phiên bản của JVM có thể chạy file này.
- **Minor version (2 byte) + Major version (2 byte)**.
- Major version của các phiên bản Java thông dụng:

|Java Version|Major Version|
|---|---|
|Java 8|52|
|Java 11|55|
|Java 17|61|
|Java 21|65|

📌 **Ví dụ**: `00 00 00 37` (0x0037 = 55) → file này được biên dịch bằng **Java 21**. (45->65)

### 🟢 **Constant Pool**

- Đây là phần lớn nhất trong file `.class`, chứa **hằng số** (tên class, tên method, chuỗi, số nguyên, v.v.).
- Dữ liệu được lưu theo **index** (chỉ mục), các thành phần khác trong file `.class` sẽ tham chiếu đến đây.
- Có nhiều loại entry, mỗi loại có **1 byte tag** đầu tiên để phân loại.

📌 **Ví dụ các entry trong constant pool:**
```java
CONSTANT_Class      # Tên class
CONSTANT_Fieldref   # Tham chiếu field
CONSTANT_Methodref  # Tham chiếu method
CONSTANT_String     # Lưu chuỗi
CONSTANT_Integer    # Lưu số nguyên
```
🔥 **Ví dụ minh họa** Giả sử chúng ta có class:
```java
public class Hello {
    public static void main(String[] args) {
        System.out.println("Hello, JVM!");
    }
}
```
📌 Một phần **constant pool** của file `.class` sẽ trông như thế này:
```java
#1 = Class          #2         // Hello
#2 = Utf8           Hello
#3 = Methodref      #4.#5      // System.out.println
#4 = Class          #6         // java/lang/System
#5 = NameAndType    #7:#8      // println:(Ljava/lang/String;)V
#6 = Utf8           java/lang/System
#7 = Utf8           println
#8 = Utf8           (Ljava/lang/String;)V
```
- `#1` là class `Hello`
- `#3` là method `System.out.println`
- `#8` là kiểu dữ liệu `void (String)`

### 🟢 **Access Flags**

- **2 byte** xác định **modifier** của class:
    - `0x0001` → `public`
    - `0x0010` → `final`
    - `0x0020` → `super` (sử dụng `invokespecial` cho super class)
    - `0x0200` → `interface`
    - `0x0400` → `abstract`
    - `0x1000` → `synthetic` (được tạo bởi trình biên dịch)
📌 **Ví dụ**:
```java
Access Flags: 0x0021 (public, super)
```
Nghĩa là class này có `public` và `super`.
### 🟢 **This Class & Super Class**

- **This Class (2 byte)**: Chỉ mục của class hiện tại trong constant pool.
- **Super Class (2 byte)**: Chỉ mục của class cha trong constant pool (`Object` nếu không kế thừa gì).

📌 **Ví dụ**:
```java
This Class: #1 (Hello)
Super Class: #2 (java/lang/Object)
```
### 🟢 **Interfaces**

- Nếu class implement interface, danh sách các interface được lưu ở đây.

📌 **Ví dụ**:
```java
public class MyClass implements Runnable {
    @Override
    public void run() {}
}
```
📌 Trong file `.class`:
```java
Interfaces: 1
Interface[0]: #3 (java/lang/Runnable)
```
### 🟢 **Fields**

- Danh sách các thuộc tính (biến instance/static).
- Mỗi field có:
    - Access flags (private, public, static, final)
    - Name index (trỏ vào constant pool)
    - Descriptor index (kiểu dữ liệu)
    - Attribute list (annotations, generic, v.v.)
📌 **Ví dụ**:
```java
private int count;
```
Trong file `.class`:
```java
Field:
  Name: #5 (count)
  Descriptor: #6 (I)  // I = int
  Flags: 0x0002 (private)
```
### 🟢 **Methods**

- Danh sách các phương thức (methods).
- Cấu trúc tương tự như fields nhưng có thêm **bytecode** trong phần Attributes.

📌 **Ví dụ**:
```java
public void print() { System.out.println("Hello"); }
```
Trong file `.class`:
```java
Method:
  Name: #9 (print)
  Descriptor: #10 ()V  // ()V = void print()
  Flags: 0x0001 (public)
  Code:
    aload_0
    invokevirtual #11 (java/lang/System.out.println)
    return
```
### 🟢 **Attributes**

- Chứa thông tin bổ sung:
    - `Code` (chứa bytecode)
    - `LineNumberTable` (map giữa dòng code Java và bytecode)
    - `SourceFile` (tên file `.java`)
    - `Annotations`, `Generic`, v.v.
## 🔥 Kết luận

File `.class` chứa **bytecode**, và JVM đọc nó theo cấu trúc cụ thể. Hiểu Class File Format giúp bạn:

- Debug lỗi class loading.
- Viết trình phân tích file `.class`.
- Hiểu cách JVM thực thi Java.