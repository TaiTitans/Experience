
---
### 1. **Xác định các yêu cầu hệ thống**

- **Loại voucher**: Thiết kế hỗ trợ các loại voucher như voucher toàn sàn, voucher cửa hàng, voucher thương hiệu, và voucher sự kiện.
- **Điều kiện áp dụng**: Voucher có thể được áp dụng với các điều kiện cụ thể như mức chi tiêu tối thiểu, số lần sử dụng tối đa, giới hạn cho người dùng mới, hoặc các nhóm khách hàng cụ thể.
- **Ngày hiệu lực và hết hạn**: Mỗi voucher sẽ có thời gian hiệu lực và hết hạn để quản lý thời gian sử dụng.
- **Hạn mức sử dụng**: Thiết lập số lượng sử dụng tối đa cho mỗi voucher, có thể là tổng số lần trên hệ thống hoặc cho từng người dùng.

### 2. **Thiết kế cơ sở dữ liệu**

- **Bảng Voucher**: Chứa các thông tin cơ bản như mã voucher, loại voucher, giá trị (phần trăm giảm giá hoặc số tiền cố định), điều kiện áp dụng (chi tiêu tối thiểu, sản phẩm hoặc cửa hàng áp dụng), số lượng tối đa, và ngày bắt đầu/kết thúc.
- **Bảng Voucher_User**: Quan hệ giữa voucher và người dùng để lưu các thông tin như đã sử dụng hay chưa, số lần sử dụng, và lịch sử sử dụng.
- **Bảng Order** (Đơn hàng): Để lưu các thông tin về đơn hàng của người dùng, giúp kiểm tra các điều kiện sử dụng voucher, ví dụ, nếu voucher áp dụng cho những đơn hàng có tổng giá trị từ một mức nhất định.

### 3. **Xây dựng các API cho tính năng voucher**

- **API tạo voucher**: API dành cho admin hoặc hệ thống để tạo voucher mới với các điều kiện và thuộc tính đã được thiết lập.
- **API lấy voucher khả dụng**: API cho phép người dùng xem các voucher hiện có, bao gồm các voucher toàn sàn hoặc voucher của các cửa hàng mà họ ghé thăm.
- **API thu thập voucher**: Người dùng sẽ có thể thu thập voucher để lưu trong tài khoản của họ và sử dụng khi đặt hàng.
- **API áp dụng voucher**: Trong bước thanh toán, API này sẽ kiểm tra xem voucher có hợp lệ cho đơn hàng đó không dựa trên các điều kiện áp dụng (mức chi tiêu, sản phẩm, ngày hiệu lực, giới hạn người dùng).
- **API kiểm tra trạng thái**: Để kiểm tra trạng thái của voucher (còn hiệu lực, hết hạn, hoặc đã đạt giới hạn sử dụng).

### 4. **Logic kiểm tra và áp dụng voucher**

- **Xác thực điều kiện**: Khi người dùng áp dụng voucher, backend sẽ kiểm tra từng điều kiện, ví dụ:
    - **Điều kiện về loại sản phẩm hoặc cửa hàng**: Đảm bảo sản phẩm trong giỏ hàng thuộc danh mục hoặc cửa hàng hợp lệ.
    - **Điều kiện về mức chi tiêu**: Tổng giá trị giỏ hàng đáp ứng mức tối thiểu cần thiết.
    - **Hạn mức sử dụng**: Kiểm tra xem voucher có còn lượt sử dụng không hoặc người dùng đã sử dụng voucher này bao nhiêu lần.
- **Tính toán giảm giá**: Sau khi xác thực voucher, hệ thống sẽ tính toán phần giảm giá và cập nhật tổng tiền thanh toán.
- **Lưu lịch sử sử dụng**: Cập nhật số lần sử dụng của voucher trong bảng `Voucher_User` và ghi lại lịch sử để hỗ trợ các phân tích và báo cáo.

### 5. **Xử lý tình huống cạnh tranh khi sử dụng voucher**

- **Đảm bảo tính toàn vẹn của dữ liệu**: Nếu voucher có số lượng giới hạn, cần thực hiện lock (khóa) để tránh trường hợp nhiều người dùng áp dụng voucher đồng thời, dẫn đến vượt quá số lượng cho phép.
- **Sử dụng Redis hoặc các hệ thống lưu trữ tạm thời khác**: Redis có thể được sử dụng để lưu trữ trạng thái của voucher trong bộ nhớ, giúp tăng tốc độ truy xuất và cập nhật trạng thái nhanh chóng.
- **Xử lý bằng hàng đợi (Queue)**: Với các voucher giới hạn trong sự kiện lớn, có thể dùng hàng đợi như Kafka hoặc RabbitMQ để xử lý yêu cầu sử dụng voucher tuần tự, đảm bảo công bằng.

### 6. **Quản lý và theo dõi dữ liệu voucher**

- **Theo dõi và phân tích**: Lưu lịch sử sử dụng voucher để phân tích hiệu quả của từng loại voucher, số lượng người dùng tham gia, mức chi tiêu trung bình.
- **Hệ thống thông báo**: Nếu voucher sắp hết hạn hoặc đạt giới hạn sử dụng, hệ thống sẽ thông báo tới người dùng để khuyến khích sử dụng sớm.

### 7. **Triển khai và tối ưu hóa hiệu suất**

- **Caching**: Sử dụng cache cho các voucher phổ biến, đặc biệt là trong các đợt sale lớn, để giảm tải truy vấn vào cơ sở dữ liệu.
- **Xử lý tải cao**: Đảm bảo hệ thống có khả năng mở rộng (scalable) để đáp ứng lượng lớn người dùng khi có các sự kiện giảm giá lớn.
- **Kiểm thử**: Thực hiện kiểm thử tải (load testing) và kiểm thử điều kiện (stress testing) để đảm bảo voucher hoạt động đúng khi có lượng lớn người dùng.

### 8. **Bảo mật**

- **Mã hóa thông tin voucher**: Đảm bảo các thông tin nhạy cảm (như mã giảm giá) được mã hóa.
- **Xác thực và phân quyền**: Kiểm soát quyền truy cập API theo vai trò, chỉ cho phép người dùng có quyền hợp lệ áp dụng hoặc xem voucher.
- **Chống gian lận**: Theo dõi các hành vi đáng ngờ như sử dụng voucher nhiều lần bất thường hoặc các mẫu gian lận khác.