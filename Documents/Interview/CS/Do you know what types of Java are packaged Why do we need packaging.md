
---

**Java là một ngôn ngữ lập trình hướng đối tượng, và có nhiều trường hợp bạn cần sử dụng các đối tượng thay vì các kiểu dữ liệu cơ bản.** Ví dụ, trong các lớp tập hợp (collection), chúng ta không thể thêm trực tiếp các kiểu int, double, v.v. vào đó. Lý do là vì container của tập hợp yêu cầu các phần tử phải thuộc kiểu Object.

Để khiến các kiểu cơ bản cũng có đặc điểm của một đối tượng, các kiểu bao bọc (wrapper types) đã được tạo ra. Điều này tương đương với việc "bao bọc" một kiểu cơ bản để nó sở hữu các thuộc tính của một đối tượng, đồng thời bổ sung thêm các thuộc tính và phương thức, làm phong phú hơn các thao tác trên kiểu cơ bản đó.

|Original type|Type of packaging|
|---|---|
|boolean|Boolean|
|byte|Byte|
|char|Character|
|float|Float|
|int|Integer|
|long|Long|
|short|Short|
|double|Double|

**Đóng gói (Boxing)**: Chuyển đổi kiểu cơ bản thành kiểu bao bọc.  
**Mở gói (Unboxing)**: Chuyển đổi kiểu bao bọc thành kiểu cơ bản.

Trình biên dịch sẽ tự động thực hiện đóng gói hoặc mở gói giữa các kiểu cơ bản và kiểu bao bọc của chúng trong các tình huống sau:

- Thực hiện phép gán (đóng gói hoặc mở gói).
- Thực hiện các phép toán hỗn hợp như cộng, trừ, nhân, chia (mở gói).
- Thực hiện so sánh với các toán tử >, <, == (mở gói).
- Gọi phương thức equals để so sánh (đóng gói).
- Khi thêm dữ liệu cơ bản vào các lớp tập hợp như ArrayList hoặc HashMap (đóng gói).

