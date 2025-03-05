
---
1. Định nghĩa JWT:

- JWT là một chuẩn mở (RFC 7519) định nghĩa một cách compact và tự chứa thông tin để truyền đạt an toàn giữa các bên dưới dạng một đối tượng JSON.
- Thông tin trong JWT có thể được xác minh và tin cậy vì nó được ký số.
- JWT có thể được ký bằng một bí mật (với thuật toán HMAC) hoặc một cặp khóa công khai/riêng tư sử dụng RSA hoặc ECDSA.

2. Khi nào nên sử dụng JWT?

- Ủy quyền: Đây là trường hợp phổ biến nhất khi sử dụng JWT. Sau khi người dùng đăng nhập, mỗi yêu cầu tiếp theo sẽ bao gồm JWT, cho phép người dùng truy cập các tuyến đường, dịch vụ và tài nguyên được phép.
- Trao đổi thông tin: JWT là một cách an toàn để truyền thông tin giữa các bên. Vì JWT có thể được ký, bạn có thể đảm bảo người gửi là ai họ nói.

3. Cấu trúc JWT:

- JWT bao gồm 3 phần: Header, Payload và Signature, được phân tách bằng dấu chấm.
- Header: Chứa thông tin về loại token (JWT) và thuật toán ký.
- Payload: Chứa các claim (tuyên bố) về thực thể (thường là người dùng) và dữ liệu bổ sung.
- Signature: Được tạo bằng cách sử dụng Header, Payload và một bí mật, để xác minh thông điệp không bị thay đổi.

4. Cách hoạt động của JWT:

- Khi người dùng đăng nhập thành công, một JWT sẽ được trả về.
- Khi người dùng muốn truy cập một tài nguyên được bảo vệ, JWT sẽ được gửi trong tiêu đề Authorization, sử dụng sơ đồ Bearer.
- Các tuyến đường được bảo vệ sẽ kiểm tra tính hợp lệ của JWT trong tiêu đề Authorization, nếu hợp lệ, người dùng sẽ được phép truy cập.

5. Tại sao nên sử dụng JWT?

- JWT nhỏ gọn hơn SAML, dễ sử dụng hơn XML.
- JWT có thể sử dụng chữ ký khóa công khai/riêng tư, an toàn hơn SWT.
- JWT dễ sử dụng hơn SAML do JSON dễ ánh xạ sang đối tượng hơn XML.
- JWT được sử dụng rộng rãi trên Internet, đặc biệt là trên các nền tảng di động.

6. Các loại claim trong payload của JWT:
    
    - Registered claims: Những claim được định nghĩa sẵn, ví dụ như iss (issuer), exp (expiration time), sub (subject), aud (audience).
    - Public claims: Những claim tùy chỉnh do người dùng định nghĩa, nhưng cần tránh xung đột.
    - Private claims: Những claim tùy chỉnh được định nghĩa giữa các bên sử dụng JWT.