
---

OAuth là viết tắt của Open Authorization, là một giao thức mở cho phép ứng dụng (client) truy cập các tài nguyên của người dùng (resource owner) trên một trang web khác (resource server) mà không cần chia sẻ thông tin đăng nhập.

Cách thức hoạt động của OAuth:

1. Người dùng (resource owner) xác thực với trang web (authorization server) và cấp quyền truy cập cho ứng dụng (client).
2. Authorization server trả về một mã thông báo (access token) cho ứng dụng.
3. Ứng dụng sử dụng access token để truy cập các tài nguyên của người dùng trên resource server.

Có 4 vai trò chính trong OAuth:

1. Resource Owner (Chủ sở hữu tài nguyên): Là người dùng cung cấp quyền truy cập cho ứng dụng.
2. Client (Ứng dụng): Là ứng dụng muốn truy cập các tài nguyên của người dùng.
3. Authorization Server (Máy chủ ủy quyền): Là nơi xác thực người dùng và cấp quyền truy cập.
4. Resource Server (Máy chủ tài nguyên): Là nơi lưu trữ các tài nguyên của người dùng.

OAuth có 4 dòng (flow) chính:

1. Authorization Code Flow: Phù hợp cho ứng dụng web, người dùng được chuyển hướng đến authorization server để xác thực.
2. Implicit Flow: Phù hợp cho ứng dụng Single Page App (SPA), access token được trả về trực tiếp.
3. Resource Owner Password Credentials Flow: Người dùng cung cấp tài khoản/mật khẩu trực tiếp cho ứng dụng.
4. Client Credentials Flow: Ứng dụng sử dụng thông tin xác thực của chính nó để lấy access token.

Các lợi ích chính của OAuth:

- Tách biệt quyền truy cập, người dùng không phải chia sẻ thông tin đăng nhập.
- Cung cấp cơ chế ủy quyền linh hoạt, người dùng có thể thu hồi quyền truy cập bất kỳ lúc nào.
- Hỗ trợ nhiều loại ứng dụng khác nhau (web, mobile, desktop).
- Được sử dụng rộng rãi trong các dịch vụ, API và ứng dụng trên internet.
