
---
### 1. Định nghĩa JWT:
JWT (JSON Web Token) là một chuẩn mở (RFC 7519) dùng để truyền tải thông tin giữa các bên dưới dạng đối tượng JSON một cách an toàn và compact.

JWT tự chứa thông tin, cho phép bên nhận xác minh danh tính của người gửi cũng như đảm bảo rằng dữ liệu không bị chỉnh sửa trong quá trình truyền tải.

JWT có thể được ký số:

HMAC (symmetric-key): Dùng một bí mật chung để ký và xác minh.

RSA hoặc ECDSA (asymmetric-key): Sử dụng cặp khóa công khai và riêng tư để ký và xác minh.

### 2. Khi nào nên sử dụng JWT?
Ủy quyền (Authorization):
Đây là trường hợp phổ biến nhất khi sử dụng JWT.

Sau khi người dùng đăng nhập thành công, máy chủ sẽ cung cấp một JWT.

Mỗi yêu cầu tiếp theo sẽ gửi JWT này trong tiêu đề (header) để xác minh và cấp quyền truy cập tài nguyên hoặc dịch vụ.

Trao đổi thông tin (Information Exchange):
JWT là một cách an toàn để truyền thông tin giữa các bên.

Vì JWT được ký số, bên nhận có thể xác minh danh tính của người gửi và đảm bảo dữ liệu không bị sửa đổi.

### 3. Cấu trúc JWT:
JWT bao gồm 3 phần chính, được phân tách bằng dấu chấm (.):

Header: Chứa loại token (luôn là JWT) và thuật toán ký (ví dụ: HS256 hoặc RS256).

Payload: Lưu trữ thông tin (claim) về thực thể (thường là người dùng) hoặc dữ liệu khác.

Signature: Được tạo ra từ Header, Payload và một bí mật (hoặc khóa riêng tư).

Ví dụ một JWT:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkphbmUgRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```
Chi tiết các phần:
Header:
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```
Payload:
```json
{
  "sub": "1234567890",
  "name": "Jane Doe",
  "iat": 1516239022
}
```
Signature:
```
HMACSHA256(
  base64UrlEncode(header) + "." + base64UrlEncode(payload),
  secret
)
```
---
### 4. Cách hoạt động của JWT:
Người dùng đăng nhập thành công → máy chủ tạo một JWT và gửi lại cho người dùng.

Người dùng lưu trữ JWT (thường trong localStorage hoặc cookies).

Khi truy cập tài nguyên được bảo vệ, JWT được gửi kèm theo mỗi yêu cầu HTTP (thường trong tiêu đề Authorization với schema Bearer).

Máy chủ kiểm tra JWT:

Giải mã và xác thực chữ ký.

Nếu hợp lệ, người dùng sẽ được cấp quyền truy cập.

Quy trình minh họa:
```
Client -> Login Request -> Server (JWT Created & Returned)
Client -> Request with JWT -> Server (JWT Verified -> Access Granted)
```
### 5. Tại sao nên sử dụng JWT?
Nhỏ gọn và hiệu quả: So với SAML, JWT có kích thước nhỏ hơn nhờ sử dụng JSON thay vì XML.

An toàn hơn: JWT hỗ trợ sử dụng chữ ký số với RSA hoặc ECDSA.

Tính phổ biến cao: JWT là một tiêu chuẩn được áp dụng rộng rãi trong các hệ thống API, microservices, và nền tảng di động.

Dễ dàng sử dụng: JSON dễ dàng ánh xạ (map) vào các đối tượng trong các ngôn ngữ lập trình hơn so với XML.

### 6. Các loại claim trong Payload của JWT:

1. Registered claims (Claim được đăng ký):
Những claim được định nghĩa sẵn trong chuẩn JWT, ví dụ:

iss (issuer): Định danh của bên phát hành token.

exp (expiration): Thời gian hết hạn của token.

sub (subject): Thực thể mà token đại diện.

aud (audience): Người nhận mà token hướng đến.

2. Public claims (Claim công khai):
Những claim tùy chỉnh có thể được định nghĩa bởi người dùng.

Cần đăng ký với IANA để tránh xung đột với các claim khác.

3. Private claims (Claim riêng tư):
Những claim được sử dụng nội bộ giữa các hệ thống cụ thể.

Ví dụ: user_role, permissions.

