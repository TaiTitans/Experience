
---
Chương này mô tả cách mã nguồn Java được biên dịch thành bytecode và cách JVM thực thi bytecode đó.

## **1️⃣ Quá trình biên dịch Java**

Quá trình biên dịch trong Java diễn ra như sau:

1️⃣ **Viết mã nguồn** (`.java`)  
2️⃣ **Biên dịch** (`javac`) → tạo bytecode (`.class`)  
3️⃣ **JVM thực thi bytecode** bằng **Interpreter** hoặc **JIT Compiler**

Ví dụ:
```java
public class Example {
    public static void main(String[] args) {
        int x = 5;
        int y = 10;
        int sum = x + y;
        System.out.println("Sum: " + sum);
    }
}
```
Mã trên sẽ được biên dịch thành **bytecode**, có thể xem bằng lệnh:
```sh
javap -c Example.class
```
Kết quả bytecode (giản lược):
```makefile
0: iconst_5
1: istore_1
2: iconst_10
3: istore_2
4: iload_1
5: iload_2
6: iadd
7: istore_3
8: getstatic java/lang/System.out : Ljava/io/PrintStream;
9: iload_3
10: invokevirtual java/io/PrintStream.println
```
- `iconst_5` → Đẩy giá trị `5` lên Stack
- `istore_1` → Lưu vào biến `x`
- `iload_1`, `iload_2` → Đọc giá trị của `x`, `y`
- `iadd` → Cộng hai số
- `invokevirtual` → Gọi `println()`

## **2️⃣ Tại sao dùng bytecode?**

- **Độc lập nền tảng**: JVM có thể chạy bytecode trên mọi hệ điều hành.
- **Tối ưu hóa**: JIT Compiler dịch bytecode thành mã máy khi cần.
- **Bảo mật**: JVM kiểm tra bytecode trước khi thực thi.

---
### SÂU HƠN NỮA ĐI ANH!!! OK EM

### **Quá trình biên dịch mã nguồn Java thành bytecode** 🚀

Quá trình biên dịch Java gồm nhiều bước để chuyển đổi mã nguồn (`.java`) thành bytecode (`.class`), sau đó JVM thực thi bytecode này. Chúng ta sẽ đi sâu vào từng giai đoạn.
## **1️⃣ Các giai đoạn của trình biên dịch Java (`javac`)**

Khi bạn chạy lệnh:
`javac Example.java

trình biên dịch (`javac`) thực hiện các bước sau:
### **🔹 Giai đoạn 1: Phân tích mã nguồn (Lexical Analysis & Parsing)**

Trình biên dịch **đọc mã nguồn** và tách thành các **token** (từ khóa, toán tử, biến, phương thức, v.v.).

- **Lexical Analysis**: Chia mã nguồn thành token (sử dụng bộ phân tích từ vựng - Lexer).
- **Parsing**: Xây dựng cây cú pháp trừu tượng (AST - Abstract Syntax Tree) để kiểm tra cú pháp.
Ví dụ:
```java
int sum = a + b;
```
Lexer sẽ tách thành:
```java
[int] [sum] [=] [a] [+] [b] [;]
```
Parser xây dựng cây AST:
```java
Assignment
 ├── Variable: sum
 ├── Expression: a + b
```
### **🔹 Giai đoạn 2: Kiểm tra ngữ nghĩa (Semantic Analysis)**

- **Kiểm tra kiểu dữ liệu** (`Type Checking`): Đảm bảo biến `sum` có cùng kiểu dữ liệu với `a + b`.
- **Kiểm tra phạm vi** (`Scope Checking`): Kiểm tra `a, b` có được khai báo trước không.
- **Kiểm tra tham chiếu**: Kiểm tra `sum` có thể gán giá trị không.

Nếu có lỗi, trình biên dịch sẽ báo lỗi:
`error: cannot find symbol`
### **🔹 Giai đoạn 3: Sinh mã trung gian (Intermediate Representation - IR)**

Sau khi kiểm tra xong, trình biên dịch tạo một **đại diện trung gian** của chương trình, thường dưới dạng **tập lệnh 3 địa chỉ (Three-Address Code - TAC)** hoặc **cây IR**.

Ví dụ, mã Java:
```java
int sum = a + b;
```
có thể chuyển thành:
```java
T1 = load a
T2 = load b
T3 = T1 + T2
sum = T3
```
IR này giúp tối ưu hóa dễ dàng hơn trước khi tạo bytecode.
### **🔹 Giai đoạn 4: Sinh mã bytecode (Bytecode Generation)**

Sau khi có IR, trình biên dịch chuyển đổi nó thành **bytecode**, là tập hợp các **lệnh máy ảo** cho JVM.  
Ví dụ:
```java
public class Example {
    public static void main(String[] args) {
        int x = 5;
        int y = 10;
        int sum = x + y;
        System.out.println(sum);
    }
}
```
Bytecode (`javap -c Example.class`):
```java
0: iconst_5      // Push 5 lên Stack
1: istore_1      // Lưu vào biến x
2: iconst_10     // Push 10 lên Stack
3: istore_2      // Lưu vào biến y
4: iload_1       // Load x từ bộ nhớ
5: iload_2       // Load y từ bộ nhớ
6: iadd          // Cộng x + y
7: istore_3      // Lưu kết quả vào sum
8: getstatic java/lang/System.out
9: iload_3       // Load sum lên Stack
10: invokevirtual java/io/PrintStream.println
```
## **2️⃣ JVM xử lý bytecode như thế nào?**

Khi chạy lệnh:
`java Example`

JVM thực hiện 2 bước:
### **🔹 Bước 1: Kiểm tra mã bytecode (Bytecode Verification)**

Trước khi chạy, JVM kiểm tra:

- Không có lệnh bytecode bất hợp lệ.
- Tất cả biến đều được khởi tạo trước khi sử dụng.
- Không truy cập ngoài vùng nhớ.

Nếu phát hiện lỗi, JVM báo `java.lang.VerifyError`.

### **🔹 Bước 2: Thực thi mã bytecode**

Có 2 cách JVM thực thi bytecode:

🔸 **Interpreter** (Bộ thông dịch):

- Đọc từng lệnh bytecode rồi thực thi ngay.
- Nhanh trong khởi động nhưng chậm về lâu dài.

🔸 **JIT Compiler (Just-In-Time Compilation)**

- Chuyển bytecode thành mã máy ngay khi cần.
- **HotSpot Compiler** tối ưu mã hay dùng để tăng tốc.
- Lưu mã đã biên dịch để dùng lại, giảm thời gian thực thi.

Ví dụ, nếu phương thức `sum()` được gọi nhiều lần, JIT sẽ biên dịch trực tiếp sang mã máy để chạy nhanh hơn.

## **3️⃣ Kết luận**

✔️ **Biên dịch Java** gồm: Lexing, Parsing, Semantic Analysis, IR Generation, Bytecode Generation.  
✔️ **Bytecode** giúp Java **độc lập nền tảng** và **dễ tối ưu hóa**.  
✔️ **JVM** thực thi bytecode bằng **Interpreter** và **JIT Compiler** để cân bằng tốc độ và hiệu suất.