

---
	Database Normalization (Chuẩn hóa Cơ sở Dữ liệu) là một quy trình tổ chức dữ liệu trong cơ sở dữ liệu nhằm giảm thiểu sự dư thừa và đảm bảo tính toàn vẹn dữ liệu. Mục tiêu của việc chuẩn hóa là phân chia cơ sở dữ liệu thành các bảng nhỏ hơn và xác định mối quan hệ giữa chúng để tránh sự trùng lặp và đảm bảo rằng mỗi bảng chỉ chứa thông tin liên quan đến một chủ đề cụ thể.

### Các Dạng Chuẩn Hóa

Có nhiều cấp độ chuẩn hóa, được gọi là các dạng chuẩn (Normal Forms). Dưới đây là các dạng chuẩn hóa phổ biến nhất:

#### 1. First Normal Form (1NF)

- **Quy tắc:** Mỗi cột trong bảng chỉ chứa một giá trị nguyên tử (không chia nhỏ được).

#### 2. Second Normal Form (2NF)

- **Quy tắc:** Đạt 1NF và mỗi cột không phải khóa phải phụ thuộc hoàn toàn vào khóa chính (không phụ thuộc vào một phần của khóa chính).
#### 3. Third Normal Form (3NF)

- **Quy tắc:** Đạt 2NF và không có phụ thuộc bắc cầu (transitive dependency). Tức là, không có cột không phải khóa nào phụ thuộc vào một cột không phải khóa khác.


### Các Dạng Chuẩn Hóa Cao Hơn

Ngoài các dạng chuẩn phổ biến như 1NF, 2NF, và 3NF, còn có các dạng chuẩn hóa cao hơn như Boyce-Codd Normal Form (BCNF), Fourth Normal Form (4NF), và Fifth Normal Form (5NF). Những dạng chuẩn này thường áp dụng cho các cơ sở dữ liệu phức tạp hơn và yêu cầu mức độ chuẩn hóa cao hơn để tránh các vấn đề tiềm ẩn.

### Tại Sao Cần Chuẩn Hóa Cơ Sở Dữ Liệu?

- **Giảm thiểu dư thừa:** Giảm sự trùng lặp dữ liệu giúp tiết kiệm không gian lưu trữ và giảm nguy cơ lỗi khi cập nhật dữ liệu.
- **Cải thiện tính toàn vẹn dữ liệu:** Đảm bảo rằng dữ liệu được lưu trữ một cách nhất quán và chính xác.
- **Tăng cường hiệu suất:** Cải thiện hiệu suất truy vấn bằng cách tổ chức dữ liệu một cách hợp lý và dễ dàng truy cập.

