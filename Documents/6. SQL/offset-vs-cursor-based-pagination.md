
---
Phân trang là kỹ thuật chia nhỏ dữ liệu lớn thành các phần nhỏ hơn để tránh tình trạng trình duyệt bị treo hoặc trải nghiệm người dùng kém. Có hai phương pháp phân trang phổ biến: phân trang dựa trên offset và phân trang dựa trên cursor.

**Phân trang dựa trên Offset:**

Phương pháp này sử dụng các tham số `LIMIT` và `OFFSET` trong truy vấn SQL để xác định số lượng bản ghi trả về và vị trí bắt đầu. Ví dụ:
```sql
SELECT * FROM table_name ORDER BY id LIMIT 50 OFFSET 200;
```
Truy vấn này sẽ trả về 50 bản ghi bắt đầu từ bản ghi thứ 201.

_Ưu điểm:_

- Dễ triển khai và hiểu.
- Phù hợp với tập dữ liệu không thay đổi nhiều.
- Cho phép người dùng chuyển đến các trang cụ thể dễ dàng.

_Nhược điểm:_

- Hiệu suất giảm khi kích thước dữ liệu tăng, do cơ sở dữ liệu phải duyệt qua nhiều bản ghi để xác định vị trí bắt đầu.
- Kết quả không nhất quán nếu dữ liệu thay đổi trong quá trình phân trang.

**Phân trang dựa trên Cursor:**

Phương pháp này sử dụng một con trỏ (cursor) để đánh dấu vị trí trong tập dữ liệu, thường dựa trên một trường sắp xếp liên tục như timestamp hoặc khóa chính. Mỗi phản hồi chứa một con trỏ để truy vấn tiếp theo. Ví dụ:
```sql
SELECT * FROM table_name WHERE id > last_id ORDER BY id LIMIT 50;
```
Trong đó, `last_id` là giá trị của bản ghi cuối cùng từ truy vấn trước đó.

_Ưu điểm:_

- Hiệu suất tốt hơn với tập dữ liệu lớn, vì không cần duyệt qua các bản ghi trước đó.
- Kết quả nhất quán ngay cả khi dữ liệu thay đổi thường xuyên.

_Nhược điểm:_

- Triển khai phức tạp hơn.
- Không hỗ trợ việc chuyển đến các trang cụ thể, chỉ có thể chuyển tiếp hoặc lùi một bước.

**Lựa chọn phương pháp phù hợp:**

- Nếu ứng dụng của bạn có tập dữ liệu nhỏ hoặc ít thay đổi, phân trang dựa trên offset có thể là lựa chọn phù hợp do tính đơn giản.
- Nếu bạn xử lý tập dữ liệu lớn hoặc dữ liệu thay đổi thường xuyên, phân trang dựa trên cursor sẽ hiệu quả hơn về hiệu suất và tính nhất quán.

Cả hai phương pháp đều yêu cầu sắp xếp dữ liệu một cách nhất quán để đảm bảo kết quả chính xác. Với phân trang dựa trên cursor, cần có một trường có thể sắp xếp liên tục, như timestamp hoặc khóa chính, để làm con trỏ.

