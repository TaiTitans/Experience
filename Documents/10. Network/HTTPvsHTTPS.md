
---
HTTP (Hypertext Transfer Protocol) và HTTPS (Hypertext Transfer Protocol Secure) là hai giao thức truyền thông được sử dụng để truyền tải dữ liệu qua Internet, với một số khác biệt quan trọng:

### HTTP (Hypertext Transfer Protocol)

1. **Chức năng**: HTTP là giao thức chuẩn để truyền tải tài liệu siêu văn bản như HTML. Nó là nền tảng của việc trao đổi thông tin trên World Wide Web.
2. **Cổng (Port)**: Mặc định sử dụng cổng 80.
3. **Bảo mật**: HTTP không cung cấp bất kỳ bảo mật nào cho dữ liệu truyền tải. Dữ liệu được truyền đi dưới dạng văn bản thuần (plaintext), dễ bị tấn công và đánh cắp.
4. **Sử dụng**: HTTP thường được sử dụng cho các trang web không yêu cầu bảo mật cao như trang thông tin công cộng, blog, diễn đàn,...

### HTTPS (Hypertext Transfer Protocol Secure)

1. **Chức năng**: HTTPS là phiên bản bảo mật của HTTP. Nó sử dụng SSL (Secure Sockets Layer) hoặc TLS (Transport Layer Security) để mã hóa dữ liệu truyền tải.
2. **Cổng (Port)**: Mặc định sử dụng cổng 443.
3. **Bảo mật**: HTTPS mã hóa dữ liệu giữa trình duyệt và máy chủ, ngăn chặn việc dữ liệu bị đọc hoặc sửa đổi bởi bên thứ ba. Nó cũng xác thực máy chủ, đảm bảo rằng người dùng đang giao tiếp với máy chủ chính xác.
4. **Sử dụng**: HTTPS được sử dụng cho các trang web yêu cầu bảo mật như trang thanh toán trực tuyến, giao dịch ngân hàng, trang đăng nhập,...

### So sánh

- **Bảo mật**: HTTPS vượt trội hơn HTTP về bảo mật do mã hóa dữ liệu và xác thực máy chủ.
- **Hiệu suất**: HTTPS có thể chậm hơn HTTP do cần thêm quá trình mã hóa và giải mã dữ liệu, tuy nhiên sự khác biệt này thường rất nhỏ và không đáng kể với hầu hết người dùng.
- **SEO**: Các công cụ tìm kiếm như Google ưu tiên các trang web sử dụng HTTPS, giúp cải thiện thứ hạng SEO của trang web.

### Tóm tắt

- **HTTP**: Nhanh hơn, không bảo mật, phù hợp cho các trang thông tin công cộng.
- **HTTPS**: Bảo mật hơn, mã hóa dữ liệu, yêu cầu cho các giao dịch nhạy cảm và trang web yêu cầu xác thực.