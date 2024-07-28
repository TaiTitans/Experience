
---
# OAuth2 là gì?
  OAuth2 hay Open Authentication thực ra nó là một _phương thức chứng thực_, xây dựng ra để các ứng dụng có thể chia sẻ tài nguyên cho nhau mà không cần chia sẻ thông tin tài khoản (_username_ và _password_). OAuth2 đảm nhận 2 việc **Authentication** (Xác thực người dùng) và **Authorization** (Ủy quyền truy cập vào tài nguyên).

![](https://images.viblo.asia/full/c1d1f627-3e40-43f5-b0f8-7c4c73f4d5e1.png)

Nhìn vào hình ảnh trên, ta có thể thấy giao thức OAuth2 được thực hiện qua 2 bên là **Service API** - cung cấp dịch vụ xác thực và ủy quyền sử dụng tài nguyên và **Application** - ứng dụng bên thứ 3 muốn lấy tài nguyên của Resource Owner. Ngoài ra, còn có một số khái niệm quan trọng như:

- **Resource Owner**: Là những người muốn ủy quyền cho phép **Application** truy cập vào tài khoản của mình. **Application** có thể lấy thông tin nhưng sẽ bị giới hạn bởi scope được cấp phép.
- **Authorization Server**: Có nhiệm vụ kiểm tra và xác thực thông tin đăng nhập của user, sau đó cung cấp một chuỗi gọi là **Access Token** để Client có thể xác thực về sau.
- **Resource Server**: Là nơi lưu trữ tài nguyên của User và bảo mật tài nguyên cho User.

### Luồng hoạt động

1. **Application** gửi yêu cầu ủy quyền truy cập vào **Resource Server**.
2. Nếu **User** ủy quyền cho phép **Application** truy cập vào tài nguyên, thì sẽ nhận được một _"giẩy ủy quyền"_ từ **User**.
3. **Application** gửi thông tin đăng nhập kèm theo _"giẩy ủy quyền"_ tới **Authorization Server**.
4. Nếu thông tin gửi đi từ **Application** là hợp lệ, **Authorization Server** sẽ trả về một đoạn chuỗi gọi là `access_token`.
5. **Application** muốn truy cập tài nguyên của **Resource Server** thì phải gửi yêu cầu kèm theo đó là chuỗi `access_token` vừa nhận được.
6. Nếu `access_token` hợp lệ, **Resource Server** sẽ trả về tài nguyên mà **Application** đang yêu cầu.

### Authorization Grant

      Để lấy được `access_token` thì phải qua bước ủy quyền và xác thực. Loại ủy quyền (Authorization Grant) phụ thuộc vào phương thức mà **Application** sử dụng ủy quyền. Đối với OAuth2, định nghĩa có 4 loại ủy quyền sau:

- **Authorization Code**: Thường sử dụng trong các server-side application.
- **Resource Owner Password Credentials**: Sử dụng với các ứng dụng được tin cậy
- **Client Credentials**: Sử dụng truy cập với các ứng dụng thông qua API
- **Implicit**: Sử dụng trong các Mobile App hoặc Web App


---

Để tạo chức năng đăng nhập OAuth sử dụng Java Spring Boot, bạn có thể thực hiện các bước sau:

1. **Cấu hình Spring Security**:
    
    - Sử dụng thư viện `spring-boot-starter-security` để cấu hình Spring Security trong ứng dụng của bạn.
    - Xác định các cấu hình bảo mật, chẳng hạn như định nghĩa các đường dẫn được bảo vệ, cấu hình các vai trò và quyền truy cập.
2. **Cấu hình OAuth2**:
    
    - Sử dụng thư viện `spring-security-oauth2-autoconfigure` để tích hợp OAuth2 vào ứng dụng.
    - Định cấu hình các nhà cung cấp OAuth, chẳng hạn như Google, Facebook, GitHub, v.v. bao gồm client ID, client secret, và các URL callback.
3. **Xử lý quy trình đăng nhập OAuth**:
    
    - Tạo một endpoint `/login/oauth2/code/{provider}` để xử lý các yêu cầu đăng nhập OAuth từ các nhà cung cấp khác nhau.
    - Sau khi người dùng đăng nhập thành công, lưu trữ thông tin người dùng trong cơ sở dữ liệu, bao gồm `oauth_provider` và `oauth_id`.
    - Nếu người dùng chưa tồn tại trong cơ sở dữ liệu, bạn có thể tạo một người dùng mới.
4. **Xác thực người dùng**:
    
    - Sử dụng `UserDetailsService` để tải thông tin người dùng từ cơ sở dữ liệu.
    - Cập nhật `UserDetails` với vai trò và quyền của người dùng.
5. **Quản lý phiên đăng nhập**:
    
    - Sử dụng `HttpSession` hoặc `CookieHttpSessionStrategy` để lưu trữ thông tin phiên đăng nhập.
    - Đảm bảo rằng người dùng có thể đăng xuất và quản lý phiên đăng nhập của mình.