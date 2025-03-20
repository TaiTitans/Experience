
---
Hãy đi sâu vào **OAuth2** trong bối cảnh bảo mật ứng dụng và cách Spring Security hỗ trợ nó. Tôi sẽ giải thích các vai trò (Resource Owner, Client, Authorization Server, Resource Server), các loại Grant Types (Authorization Code, Implicit, Password, Client Credentials), và vai trò của Refresh Token cùng Access Token.

---

### **1. Các vai trò trong OAuth2**

OAuth2 định nghĩa bốn vai trò chính trong luồng xác thực và phân quyền:

#### **a. Resource Owner**
- **Định nghĩa**: Là người sở hữu tài nguyên (dữ liệu) được bảo vệ, thường là người dùng cuối (end-user).
- **Vai trò**: Cấp quyền cho ứng dụng (Client) truy cập tài nguyên của họ thông qua Authorization Server.
- **Ví dụ**: Bạn (người dùng) cho phép ứng dụng bên thứ ba (như một app ghi chú) truy cập Google Drive của bạn.

#### **b. Client**
- **Định nghĩa**: Là ứng dụng hoặc dịch vụ muốn truy cập tài nguyên của Resource Owner.
- **Vai trò**: Gửi yêu cầu đến Authorization Server để lấy token, sau đó dùng token để truy cập Resource Server.
- **Loại Client**:
  - **Confidential Client**: Có thể giữ bí mật (client secret), như ứng dụng server-side.
  - **Public Client**: Không giữ bí mật, như ứng dụng mobile hoặc SPA (Single Page Application).
- **Ví dụ**: Một ứng dụng web cần đọc email của bạn từ Gmail.

#### **c. Authorization Server**
- **Định nghĩa**: Server chịu trách nhiệm xác thực Resource Owner và cấp token (Access Token, Refresh Token) cho Client.
- **Vai trò**: Xác minh danh tính người dùng, cấp quyền, và phát hành token.
- **Ví dụ**: Google, Keycloak, Okta là các Authorization Server phổ biến.

#### **d. Resource Server**
- **Định nghĩa**: Server lưu trữ tài nguyên được bảo vệ (API, dữ liệu) và chỉ chấp nhận truy cập với token hợp lệ.
- **Vai trò**: Xác minh Access Token từ Client trước khi trả về tài nguyên.
- **Ví dụ**: API của Gmail (nơi lưu email của bạn).

#### **Luồng cơ bản**:
```
Resource Owner → Client → Authorization Server → Client (với token) → Resource Server
```

---

### **2. Grant Types trong OAuth2**
Grant Type là cách Client lấy Access Token từ Authorization Server. Dưới đây là bốn loại chính:

#### **a. Authorization Code**
- **Mô tả**: 
  - Client (thường là ứng dụng web) yêu cầu Resource Owner cấp quyền thông qua Authorization Server, nhận mã code, sau đó đổi code lấy Access Token.
- **Luồng**:
  1. Client chuyển hướng người dùng đến Authorization Server (với `client_id`, `redirect_uri`).
  2. Người dùng đăng nhập và đồng ý cấp quyền.
  3. Authorization Server trả về mã code qua `redirect_uri`.
  4. Client gửi code + `client_secret` đến Authorization Server để lấy Access Token.
- **Ưu điểm**: An toàn vì token không được gửi trực tiếp qua trình duyệt.
- **Ứng dụng**: Web app server-side.
- **Trong Spring Security**:
  - Cấu hình Client:
    ```java
    @Configuration
    public class SecurityConfig extends WebSecurityConfigurerAdapter {
        @Override
        protected void configure(HttpSecurity http) throws Exception {
            http.oauth2Login() // Hỗ trợ OAuth2 Authorization Code flow
                .and()
                .authorizeRequests()
                .anyRequest().authenticated();
        }
    }
    ```
    - `application.properties`:
      ```
      spring.security.oauth2.client.registration.google.client-id=your-client-id
      spring.security.oauth2.client.registration.google.client-secret=your-client-secret
      spring.security.oauth2.client.registration.google.redirect-uri={baseUrl}/login/oauth2/code/google
      ```

#### **b. Implicit**
- **Mô tả**: 
  - Client (thường là SPA hoặc ứng dụng không có backend) nhận Access Token trực tiếp từ Authorization Server mà không qua bước code.
- **Luồng**:
  1. Client chuyển hướng người dùng đến Authorization Server.
  2. Sau khi người dùng đồng ý, token được gửi về qua `redirect_uri` (trong fragment URL: `#access_token=...`).
- **Ưu điểm**: Đơn giản, không cần backend.
- **Nhược điểm**: Token lộ qua trình duyệt, kém an toàn hơn.
- **Ứng dụng**: SPA, ứng dụng không giữ được secret.
- **Trong Spring Security**: Ít được dùng trực tiếp, thường thay bằng Authorization Code với PKCE (xem bên dưới).

#### **c. Password (Resource Owner Password Credentials)**
- **Mô tả**: 
  - Client lấy username/password trực tiếp từ Resource Owner và gửi đến Authorization Server để nhận Access Token.
- **Luồng**:
  1. Client gửi `username`, `password`, `client_id`, `client_secret` đến Authorization Server.
  2. Authorization Server trả về Access Token.
- **Ưu điểm**: Đơn giản, không cần redirect.
- **Nhược điểm**: Yêu cầu Client đáng tin cậy tuyệt đối (ít dùng trong ứng dụng hiện đại).
- **Ứng dụng**: Ứng dụng nội bộ hoặc legacy.
- **Trong Spring Security**:
  - Dùng với Resource Server hoặc custom logic, ít phổ biến trong Spring Boot.

#### **d. Client Credentials**
- **Mô tả**: 
  - Client dùng thông tin của chính nó (`client_id`, `client_secret`) để lấy Access Token, không liên quan đến Resource Owner.
- **Luồng**:
  1. Client gửi `client_id`, `client_secret` đến Authorization Server.
  2. Nhận Access Token để truy cập tài nguyên.
- **Ưu điểm**: Phù hợp cho giao tiếp server-to-server.
- **Ứng dụng**: Microservices, API không cần người dùng.
- **Trong Spring Security**:
  - Cấu hình Resource Server:
    ```java
    @Configuration
    public class SecurityConfig extends WebSecurityConfigurerAdapter {
        @Override
        protected void configure(HttpSecurity http) throws Exception {
            http.oauth2ResourceServer()
                .jwt(); // Xác minh token từ Client Credentials
        }
    }
    ```

#### **PKCE (Proof Key for Code Exchange)**:
- Là phần mở rộng của Authorization Code, tăng cường bảo mật cho Public Client (như mobile app) bằng cách thêm `code_verifier` và `code_challenge`.

---

### **3. Refresh Token và Access Token**

#### **a. Access Token**
- **Định nghĩa**: 
  - Là chuỗi (thường là JWT) được Client dùng để truy cập Resource Server.
- **Đặc điểm**:
  - Có thời hạn ngắn (thường 5 phút - 1 giờ).
  - Mang thông tin quyền (scope) và danh tính (nếu là JWT).
- **Ví dụ**:
  ```
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```

#### **b. Refresh Token**
- **Định nghĩa**: 
  - Là chuỗi được dùng để lấy Access Token mới khi token cũ hết hạn, mà không cần người dùng đăng nhập lại.
- **Đặc điểm**:
  - Thời hạn dài hơn Access Token (ngày, tuần, hoặc tháng).
  - Chỉ gửi đến Authorization Server, không dùng trực tiếp với Resource Server.
- **Luồng**:
  1. Client gửi Refresh Token đến Authorization Server (endpoint `/token`).
  2. Nhận Access Token mới (và đôi khi Refresh Token mới nếu được cấu hình).
- **Ví dụ**:

```http
  POST /oauth/token
  grant_type=refresh_token
  refresh_token=abc123...
  client_id=your-client-id
  client_secret=your-client-secret
```
  
#### **Trong Spring Security**
- **Client lấy token**:
  - Spring Boot hỗ trợ thông qua `OAuth2AuthorizedClientService`.
- **Resource Server xác minh**:
  ```java
  @Configuration
  public class SecurityConfig extends WebSecurityConfigurerAdapter {
      @Override
      protected void configure(HttpSecurity http) throws Exception {
          http.oauth2ResourceServer()
              .jwt()
              .jwtAuthenticationConverter(jwtAuthenticationConverter());
      }
  }
  ```
- **Refresh Token**:
  - Spring Security không tự động xử lý refresh, cần gọi endpoint Authorization Server thủ công hoặc dùng thư viện như `RestTemplate`.

---

### **Tóm tắt**
- **Vai trò**:
  - Resource Owner: Người dùng sở hữu tài nguyên.
  - Client: Ứng dụng truy cập tài nguyên.
  - Authorization Server: Cấp token.
  - Resource Server: Bảo vệ tài nguyên.
- **Grant Types**:
  - Authorization Code: An toàn cho web app.
  - Implicit: Đơn giản cho SPA (ít dùng).
  - Password: Legacy, ít an toàn.
  - Client Credentials: Server-to-server.
- **Refresh Token và Access Token**:
  - Access Token: Ngắn hạn, truy cập tài nguyên.
  - Refresh Token: Dài hạn, làm mới Access Token.
