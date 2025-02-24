
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

---
HTTP (HyperText Transfer Protocol) có nhiều phiên bản, mỗi phiên bản cải tiến so với phiên bản trước về hiệu suất, bảo mật và tính linh hoạt. Dưới đây là các phiên bản chính của HTTP:

### **1. HTTP/0.9 (1991) - Phiên bản đầu tiên**

- Chỉ hỗ trợ **GET** request.
- Chỉ truyền dữ liệu dạng **HTML thuần** (không hỗ trợ hình ảnh, CSS, JS).
- Không có tiêu đề (header) trong request và response.
- Kết nối sẽ bị đóng ngay sau khi server gửi phản hồi.

---

### **2. HTTP/1.0 (1996 - RFC 1945)**

- **Bổ sung phương thức mới**: Hỗ trợ `POST`, `HEAD`.
- **Thêm header**: Request và response có thể chứa metadata (ví dụ: `Content-Type`, `Content-Length`).
- **Hỗ trợ định dạng dữ liệu khác** ngoài HTML (như hình ảnh, video, JSON).
- **Nhược điểm**: Mỗi request đều mở một kết nối TCP mới, gây lãng phí tài nguyên và giảm hiệu suất.

---

### **3. HTTP/1.1 (1997 - RFC 2068, cập nhật RFC 2616)**

- **Giữ kết nối lâu hơn** (Persistent Connection): Cho phép sử dụng cùng một kết nối TCP để gửi nhiều request (Keep-Alive).
- **Thêm các phương thức mới**: `PUT`, `DELETE`, `OPTIONS`, `TRACE`.
- **Hỗ trợ chunked transfer encoding**: Giúp truyền dữ liệu lớn mà không cần biết trước kích thước.
- **Pipeline Requests**: Cho phép gửi nhiều request cùng lúc mà không cần chờ response của request trước.
- **Cải thiện cache**: Thêm nhiều cơ chế kiểm soát bộ nhớ đệm với các header như `ETag`, `Cache-Control`.

---

### **4. HTTP/2 (2015 - RFC 7540)**

- **Nén header** bằng HPACK: Giảm kích thước request và response.
- **Hỗ trợ multiplexing**: Cho phép gửi nhiều request đồng thời trên một kết nối TCP duy nhất.
- **Server Push**: Server có thể gửi tài nguyên cho client mà không cần client yêu cầu (hữu ích cho việc tải trước CSS, JS).
- **Binary Protocol**: HTTP/2 sử dụng binary thay vì text như HTTP/1.1, giúp parsing nhanh hơn.
- **Cải thiện tốc độ**: Giảm độ trễ, tăng hiệu suất tải trang.

---

### **5. HTTP/3 (2022 - RFC 9114)**

- **Chuyển từ TCP sang QUIC**: QUIC là giao thức truyền tải mới chạy trên UDP, giúp giảm độ trễ.
- **Tích hợp TLS 1.3**: Cải thiện bảo mật và tốc độ handshake.
- **Hạn chế vấn đề head-of-line blocking**: Nhờ QUIC, dữ liệu không bị chặn nếu một gói tin bị mất.
- **Cải thiện hiệu suất trên mạng không ổn định**: Phù hợp hơn với thiết bị di động, mạng Wi-Fi, 4G, 5G.

---

### **So sánh nhanh các phiên bản HTTP**

|**Phiên bản**|**Kết nối TCP**|**Multiplexing**|**Compression**|**Bảo mật**|
|---|---|---|---|---|
|HTTP/0.9|Mở/kết thúc mỗi request|❌|❌|❌|
|HTTP/1.0|Mở/kết thúc mỗi request|❌|❌|❌|
|HTTP/1.1|Keep-Alive|Pipeline (không hiệu quả)|❌|✅ (TLS 1.2)|
|HTTP/2|Keep-Alive|✅|✅ (HPACK)|✅ (TLS 1.2, 1.3)|
|HTTP/3|QUIC (UDP)|✅|✅ (HPACK)|✅ (TLS 1.3)|

**Kết luận:**

- **HTTP/1.1** vẫn phổ biến nhưng không tối ưu về hiệu suất.
- **HTTP/2** cải thiện tốc độ tải trang nhờ multiplexing và nén header.
- **HTTP/3** khắc phục nhược điểm của HTTP/2 khi gặp mạng kém ổn định bằng cách sử dụng QUIC.

Hiện nay, **HTTP/2 và HTTP/3** đang được khuyến khích sử dụng để tối ưu hiệu suất website và ứng dụng. 🚀