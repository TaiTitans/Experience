
---
Hãy đi sâu vào **OpenID Connect** trong bối cảnh bảo mật ứng dụng và cách nó liên quan đến Spring Security. Tôi sẽ giải thích OpenID Connect là gì, sự khác biệt với OAuth2, vai trò của ID Token, và cách xác minh nó.

---

### **1. OpenID Connect là gì và khác biệt với OAuth2**

#### **a. OpenID Connect là gì?**
- **Định nghĩa**:
  - OpenID Connect (OIDC) là một giao thức xác thực (authentication protocol) được xây dựng trên nền OAuth2. Nó bổ sung khả năng xác minh danh tính người dùng (identity) bên cạnh việc cấp quyền truy cập tài nguyên (authorization) mà OAuth2 cung cấp.
- **Mục đích**:
  - Cho phép ứng dụng (Client) không chỉ truy cập tài nguyên mà còn biết "người dùng là ai" thông qua thông tin danh tính đáng tin cậy.
- **Ví dụ**:
  - Khi bạn đăng nhập vào một ứng dụng bằng "Login with Google", OIDC xác minh bạn là ai và cung cấp thông tin như tên, email.

#### **b. Khác biệt với OAuth2**
- **OAuth2**:
  - Là giao thức **phân quyền (authorization)**.
  - Tập trung vào việc cấp quyền truy cập tài nguyên (access resources) thông qua Access Token.
  - Không cung cấp thông tin danh tính người dùng một cách tiêu chuẩn.
  - Ví dụ: OAuth2 cho phép ứng dụng đọc Google Drive của bạn, nhưng không nói bạn là ai.
- **OpenID Connect**:
  - Là giao thức **xác thực (authentication)** dựa trên OAuth2.
  - Bổ sung **ID Token**, một token chứa thông tin danh tính người dùng (identity).
  - Dùng các endpoint và luồng của OAuth2, nhưng mở rộng với tính năng xác thực.
- **So sánh chi tiết**:

| Tiêu chí             | OAuth2                          | OpenID Connect                  |
|----------------------|---------------------------------|---------------------------------|
| **Mục đích**         | Phân quyền (Authorization)     | Xác thực (Authentication)       |
| **Token chính**      | Access Token, Refresh Token    | Access Token, ID Token          |
| **Thông tin danh tính** | Không tiêu chuẩn            | Cung cấp qua ID Token           |
| **Ứng dụng**         | Truy cập API, tài nguyên       | Đăng nhập SSO, xác minh danh tính |
| **Giao thức cơ sở**  | -                              | Xây trên OAuth2                 |

- **Kết hợp**: OIDC dùng OAuth2 làm nền tảng, nên bạn có thể vừa xác thực (ID Token) vừa phân quyền (Access Token) trong cùng một luồng.

---

### **2. ID Token và cách xác minh**

#### **a. ID Token là gì?**
- **Định nghĩa**:
  - ID Token là một **JSON Web Token (JWT)** chứa thông tin danh tính của người dùng, được cấp bởi Authorization Server trong luồng OIDC.
- **Cấu trúc**:
  - Gồm 3 phần: **Header**, **Payload**, **Signature** (giống JWT).
  - **Header**: Thông tin thuật toán ký (VD: "RS256").
  - **Payload**: Chứa các claim (thuộc tính) về người dùng:
    - `iss` (issuer): Định danh của Authorization Server.
    - `sub` (subject): ID duy nhất của người dùng.
    - `aud` (audience): Định danh của Client nhận token.
    - `exp` (expiration): Thời gian hết hạn.
    - `iat` (issued at): Thời gian phát hành.
    - Các claim tùy chọn: `name`, `email`, `picture`, v.v.
  - **Signature**: Chữ ký số để xác minh tính toàn vẹn.
- **Ví dụ ID Token (đã giải mã)**:
  ```json
  {
    "iss": "https://accounts.google.com",
    "sub": "1234567890",
    "aud": "your-client-id",
    "exp": 1698765432,
    "iat": 1698761832,
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```

#### **b. Cách xác minh ID Token**
Xác minh ID Token là bước quan trọng để đảm bảo token hợp lệ và không bị giả mạo. Quy trình bao gồm:

1. **Lấy khóa công khai (Public Key)**:
   - Authorization Server cung cấp endpoint **JWKS (JSON Web Key Set)** (VD: `https://accounts.google.com/.well-known/jwks.json`) chứa các public key để xác minh chữ ký.
   - Client tải JWKS để lấy key tương ứng với `kid` (key ID) trong header của ID Token.

2. **Kiểm tra chữ ký (Signature)**:
   - Dùng public key để giải mã phần Signature và so sánh với hash của Header + Payload.
   - Nếu khớp, token không bị thay đổi.

3. **Kiểm tra các claim**:
   - **`iss`**: Phải khớp với issuer mong đợi (VD: "https://accounts.google.com").
   - **`aud`**: Phải chứa `client_id` của Client.
   - **`exp`**: Token chưa hết hạn (so với thời gian hiện tại).
   - **`iat`**: Thời gian phát hành hợp lý.
   - **`nonce`** (nếu có): Khớp với giá trị Client gửi trong request ban đầu, chống tấn công replay.

4. **Kết quả**:
   - Nếu tất cả kiểm tra đều qua, ID Token được coi là hợp lệ, và thông tin danh tính có thể tin cậy.

#### **c. Trong Spring Security**
- Spring Security hỗ trợ OIDC qua cấu hình OAuth2 Client:
  ```java
  @Configuration
  @EnableWebSecurity
  public class SecurityConfig extends WebSecurityConfigurerAdapter {
      @Override
      protected void configure(HttpSecurity http) throws Exception {
          http.authorizeRequests()
              .anyRequest().authenticated()
              .and()
              .oauth2Login(); // Kích hoạt OIDC với Authorization Code flow
      }
  }
  ```
- **Cấu hình trong `application.properties`**:
  ```
  spring.security.oauth2.client.registration.google.client-id=your-client-id
  spring.security.oauth2.client.registration.google.client-secret=your-client-secret
  spring.security.oauth2.client.registration.google.scope=openid,email,profile
  spring.security.oauth2.client.provider.google.issuer-uri=https://accounts.google.com
  ```
  - Scope `openid` kích hoạt OIDC, yêu cầu ID Token bên cạnh Access Token.

- **Truy xuất ID Token**:
  ```java
  @GetMapping("/user")
  public String getUserInfo(OAuth2AuthenticationToken token) {
      OAuth2User user = token.getPrincipal();
      Map<String, Object> attributes = user.getAttributes(); // Lấy thông tin từ ID Token
      return "Hello, " + attributes.get("name");
  }
  ```

- **Xác minh tự động**:
  - Spring Security dùng `NimbusJwtDecoder` để xác minh ID Token:
    - Tự động tải JWKS từ `issuer-uri`.
    - Kiểm tra chữ ký và các claim cơ bản (`iss`, `aud`, `exp`).

- **Tùy chỉnh xác minh**:
  ```java
  @Bean
  public JwtDecoder jwtDecoder() {
      return NimbusJwtDecoder.withJwkSetUri("https://accounts.google.com/.well-known/jwks.json").build();
  }
  ```

---

### **Tóm tắt**
- **OpenID Connect**:
  - Là lớp xác thực trên OAuth2, cung cấp thông tin danh tính qua ID Token.
  - Khác OAuth2 ở mục đích (authentication vs authorization) và sự hiện diện của ID Token.
- **ID Token**:
  - JWT chứa thông tin người dùng (claim).
  - Xác minh bằng public key từ JWKS, kiểm tra chữ ký và claim.
- **Spring Security**: Hỗ trợ OIDC qua `oauth2Login`, tự động xử lý ID Token.
