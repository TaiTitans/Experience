# Servers URL-Syntax



- Các thành phần cơ bản của một url

1. Schema

   - Biểu thị các protocol (HTTPS, HTTP) -> Browser sẽ sử dụng để yêu cầu nguồn tài nguyên


2. Domain Name
3. Port

4.Parameters:

  - Là một danh sách bao gồm key/value kết hợp với & symbol.


5. Anchor:
   - Sử dụng để định vị một vị trí cụ thể trong tài liệu HTML + #ID


6. Cách hoạt động của URL:
```
Phân giải DNS: Khi bạn nhập URL vào trình duyệt web, trình duyệt sẽ gửi yêu cầu đến máy chủ DNS (Domain Name System) để tìm địa chỉ IP của máy chủ web tương ứng với tên miền trong URL.

Tạo kết nối: Trình duyệt sử dụng địa chỉ IP nhận được từ bước trên để thiết lập một kết nối TCP/IP đến máy chủ web.

Gửi yêu cầu HTTP: Trình duyệt gửi một yêu cầu HTTP đến máy chủ web, yêu cầu trang web cụ thể được chỉ định trong URL.

Xử lý yêu cầu: Máy chủ web nhận yêu cầu, xác định trang web được yêu cầu và tiến hành xử lý yêu cầu.

Tạo và trả về nội dung: Máy chủ web tạo ra trang web hoặc tải trang từ các tệp đã được lưu trữ, sau đó trả về nội dung của trang web đó dưới dạng các gói dữ liệu HTTP.

Hiển thị trang web: Trình duyệt nhận dữ liệu từ máy chủ web và hiển thị nội dung trang web cho người dùng.

Tải tài nguyên phụ: Nếu trang web chứa các tài nguyên như hình ảnh, tập tin CSS hoặc JavaScript, trình duyệt sẽ tiếp tục gửi yêu cầu để tải các tài nguyên này.

Hiển thị đầy đủ trang web: Khi tất cả các tài nguyên đã được tải xuống và trang web đã được hiển thị hoàn chỉnh, người dùng có thể tương tác với trang web và trình duyệt sẽ xử lý các hoạt động tiếp theo.
```
