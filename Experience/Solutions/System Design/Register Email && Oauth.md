
---
1. **Đăng ký bằng email với xác thực OTP**:
    
    - Khi người dùng đăng ký bằng email, hãy lấy thông tin từ form đăng ký như `username`, `password`, `email`.
    - Tạo một mã OTP ngẫu nhiên và lưu nó vào database hoặc lưu trữ tạm thời (ví dụ: trong bộ nhớ cache).
    - Gửi mã OTP đến email của người dùng bằng cách sử dụng dịch vụ email (ví dụ: SMTP).
    - Yêu cầu người dùng nhập mã OTP để xác thực email.
    - Sau khi xác thực thành công, tạo một `User` entity mới với các thông tin như `username`, `password`, `email`, `oauth_provider = null`, `oauth_id = null`.
    - Thiết lập vai trò (role) mặc định cho người dùng mới, ví dụ: `ROLE_CUSTOMER`.
    - Lưu `User` entity mới vào database.
2. **Đăng ký bằng OAuth (ví dụ: Gmail)**:
    
    - Khi người dùng chọn tùy chọn "Đăng ký bằng Gmail", chuyển hướng họ đến trang đăng nhập của Google.
    - Sau khi người dùng đăng nhập và cấp quyền, nhận mã token từ Google.
    - Sử dụng mã token để lấy thông tin về người dùng, như email, tên, ảnh đại diện, v.v.
    - Kiểm tra xem người dùng đã có tài khoản trong hệ thống chưa, bằng cách tìm kiếm người dùng có `oauth_provider = 'google'` và `oauth_id = [id người dùng từ Google]`.
    - Nếu người dùng chưa có tài khoản, tạo một `User` entity mới với các thông tin như `username = [tên người dùng từ Google]`, `password = null`, `email = [email người dùng từ Google]`, `oauth_provider = 'google'`, `oauth_id = [id người dùng từ Google]`.
    - Thiết lập vai trò (role) mặc định cho người dùng mới, ví dụ: `ROLE_CUSTOMER`.
    - Lưu `User` entity mới vào database.
    - Nếu người dùng đã có tài khoản, hãy đăng nhập họ vào hệ thống.