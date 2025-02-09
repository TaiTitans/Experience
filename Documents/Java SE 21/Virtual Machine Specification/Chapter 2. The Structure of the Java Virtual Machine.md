
---
Chương này mô tả cấu trúc tổng thể của **Java Virtual Machine (JVM)**, bao gồm các thành phần chính và cách nó thực thi chương trình Java.

## **1️⃣ Các thành phần chính của JVM**

JVM bao gồm các thành phần cốt lõi sau:

### **📌 1.1 Class Loader Subsystem (Hệ thống nạp lớp)**

- **Nhiệm vụ:** Tải và liên kết các tệp `.class` (bytecode) vào bộ nhớ.
- **Gồm các giai đoạn chính:**
    - **Loading (Tải lớp)**: Đọc bytecode từ tệp `.class` hoặc từ nguồn khác.
    - **Linking (Liên kết lớp)**:
        - **Verification (Xác minh)**: Kiểm tra xem bytecode có hợp lệ không.
        - **Preparation (Chuẩn bị)**: Cấp phát bộ nhớ cho biến tĩnh.
        - **Resolution (Giải quyết tham chiếu)**: Chuyển tham chiếu từ ký hiệu thành địa chỉ thực tế.
    - **Initialization (Khởi tạo lớp)**: Chạy các khối `static {}` và gán giá trị ban đầu.

### **📌 1.2 Runtime Data Areas (Khu vực dữ liệu trong thời gian chạy)**

JVM quản lý bộ nhớ với các vùng sau:

- **📌 Method Area (Vùng lưu trữ metadata của lớp)**
    
    - Lưu trữ thông tin về lớp, bao gồm tên lớp, tên phương thức, kiểu dữ liệu, và bytecode.
    - Chứa bảng hằng số (Constant Pool).
    - Chia sẻ giữa tất cả các luồng (threads).
- **📌 Heap (Bộ nhớ heap)**
    
    - Dùng để cấp phát đối tượng (objects) và biến instance.
    - Được quản lý bởi Garbage Collector (GC).
- **📌 Java Stack (Ngăn xếp Java)**
    
    - Mỗi luồng có một Stack riêng.
    - Mỗi Stack chứa nhiều **Stack Frames**, tương ứng với mỗi phương thức đang thực thi.
- **📌 PC Register (Thanh ghi bộ đếm chương trình)**
    
    - Lưu địa chỉ lệnh đang thực thi trong mỗi luồng.
- **📌 Native Method Stack (Ngăn xếp phương thức gốc)**
    
    - Chứa dữ liệu và lệnh thực thi của **phương thức native** (viết bằng C/C++).
### **📌 1.3 Execution Engine (Bộ máy thực thi)**

- **Bộ phận chính:**
    - **Interpreter**: Đọc và thực thi từng lệnh bytecode.
    - **Just-In-Time (JIT) Compiler**: Dịch bytecode thành mã máy để tăng hiệu suất.
    - **Garbage Collector (GC)**: Giải phóng bộ nhớ của đối tượng không còn sử dụng.
### **📌 1.4 Native Interface (Giao diện Native)**

- **Java Native Interface (JNI):** Cho phép Java gọi các thư viện native (C/C++).

## **2️⃣ Quy trình thực thi một chương trình Java trong JVM**

1️⃣ **Viết mã nguồn**: Bạn viết mã Java (`.java`).  
2️⃣ **Biên dịch**: Trình biên dịch `javac` dịch mã nguồn thành bytecode (`.class`).  
3️⃣ **Nạp lớp**: Class Loader nạp bytecode vào bộ nhớ.  
4️⃣ **Liên kết và khởi tạo**: JVM liên kết và khởi tạo lớp.  
5️⃣ **Thực thi**: Execution Engine thực thi bytecode.

## **3️⃣ Tóm tắt**

- JVM có các thành phần chính: **Class Loader, Runtime Data Areas, Execution Engine, Native Interface**.
- Quản lý bộ nhớ thông qua **Heap, Stack, Method Area, PC Register, Native Method Stack**.
- Dùng **Interpreter và JIT Compiler** để thực thi chương trình.