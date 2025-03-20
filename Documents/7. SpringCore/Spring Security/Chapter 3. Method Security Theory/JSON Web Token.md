
---
Hãy đi sâu vào **JSON Web Token (JWT)** trong bối cảnh bảo mật ứng dụng và cách nó được sử dụng trong Spring Security. Tôi sẽ giải thích cấu trúc của JWT (Header, Payload, Signature) và lý thuyết về Signing (ký) cùng Verification (xác minh).

---

### **1. Cấu trúc: Header, Payload, Signature**

JWT là một chuỗi được mã hóa, bao gồm ba phần chính, phân tách bằng dấu chấm (`.`): **Header**, **Payload**, và **Signature**. Chuỗi hoàn chỉnh trông như sau:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

#### **a. Header**
- **Mô tả**: 
  - Phần đầu tiên của JWT, chứa thông tin về loại token và thuật toán ký.
- **Cấu trúc**: 
  - Một đối tượng JSON, được mã hóa Base64 URL.
- **Các trường phổ biến**:
  - `alg`: Thuật toán ký (VD: "HS256" - HMAC SHA-256, "RS256" - RSA SHA-256).
  - `typ`: Loại token (thường là "JWT").
- **Ví dụ**:
  ```json
  {
    "alg": "HS256",
    "typ": "JWT"
  }
  ```
  - Sau khi mã hóa Base64 URL: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9`.

#### **b. Payload**
- **Mô tả**: 
  - Phần thứ hai, chứa dữ liệu thực tế (claim) mà token mang theo, như thông tin người dùng hoặc quyền.
- **Cấu trúc**: 
  - Một đối tượng JSON, mã hóa Base64 URL.
- **Các loại claim**:
  1. **Registered Claims** (chuẩn):
     - `iss` (issuer): Nguồn phát hành token.
     - `sub` (subject): Chủ thể (VD: ID người dùng).
     - `aud` (audience): Đối tượng nhận token.
     - `exp` (expiration): Thời gian hết hạn.
     - `iat` (issued at): Thời gian phát hành.
  2. **Public Claims**: Claim công khai, tự định nghĩa (VD: `role`).
  3. **Private Claims**: Claim riêng giữa các bên (VD: `department`).
- **Ví dụ**:
  ```json
  {
    "sub": "1234567890",
    "name": "John Doe",
    "iat": 1516239022,
    "role": "USER"
  }
  ```
  - Sau khi mã hóa Base64 URL: `eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ`.

#### **c. Signature**
- **Mô tả**: 
  - Phần thứ ba, đảm bảo tính toàn vẹn và nguồn gốc của token bằng cách ký Header và Payload.
- **Cách tạo**:
  - Dùng thuật toán trong Header (`alg`) để ký chuỗi `Base64URL(Header).Base64URL(Payload)` với một khóa bí mật (secret key).
  - Kết quả là chuỗi mã hóa Base64 URL.
- **Ví dụ (HS256)**:
  - Input: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ`
  - Secret: `your-256-bit-secret`
  - Ký bằng HMAC SHA-256 → `SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

#### **Kết hợp**:
- JWT hoàn chỉnh: `Header.Payload.Signature`.
- Dễ dàng giải mã Header và Payload bằng Base64, nhưng không thể giả mạo Signature nếu không có khóa bí mật.

---

### **2. Lý thuyết về Signing và Verification**

#### **a. Signing (Ký JWT)**
- **Mục đích**:
  - Đảm bảo token không bị thay đổi (tính toàn vẹn) và đến từ nguồn đáng tin cậy (tính xác thực).
- **Quy trình**:
  1. **Tạo Header và Payload**:
     - Xây dựng JSON cho Header (chọn thuật toán) và Payload (thêm claim).
  2. **Mã hóa Base64 URL**:
     - Chuyển Header và Payload thành chuỗi Base64 URL.
  3. **Ký dữ liệu**:
     - Kết hợp `Base64URL(Header).Base64URL(Payload)` và ký bằng khóa:
       - **HS256 (HMAC SHA-256)**: Dùng khóa bí mật đối xứng (symmetric key).
       - **RS256 (RSA SHA-256)**: Dùng private key (asymmetric key).
  4. **Tạo JWT**:
     - Nối Header, Payload, và Signature bằng dấu chấm.
- **Ví dụ (HS256)**:
  ```java
  import io.jsonwebtoken.Jwts;
  import io.jsonwebtoken.SignatureAlgorithm;

  String secret = "your-256-bit-secret";
  String jwt = Jwts.builder()
      .setSubject("1234567890")
      .setIssuedAt(new Date())
      .signWith(SignatureAlgorithm.HS256, secret.getBytes())
      .compact();
  ```

#### **b. Verification (Xác minh JWT)**
- **Mục đích**:
  - Kiểm tra xem JWT có hợp lệ, không bị giả mạo, và còn hiệu lực hay không.
- **Quy trình**:
  1. **Tách JWT**:
     - Chia chuỗi thành Header, Payload, Signature.
  2. **Kiểm tra chữ ký**:
     - Dùng khóa (secret key với HS256, public key với RS256) để tái tạo Signature từ `Base64URL(Header).Base64URL(Payload)`.
     - So sánh với Signature trong JWT:
       - Nếu khớp → Token không bị thay đổi.
       - Nếu không khớp → Token giả mạo, ném lỗi.
  3. **Kiểm tra claim**:
     - `exp`: Token chưa hết hạn.
     - `iss`, `aud`: Khớp với giá trị mong đợi.
     - Các claim khác (nếu cần): `sub`, `role`, v.v.
- **Ví dụ (HS256)**:
  ```java
  import io.jsonwebtoken.Claims;
  import io.jsonwebtoken.Jwts;

  String secret = "your-256-bit-secret";
  try {
      Claims claims = Jwts.parser()
          .setSigningKey(secret.getBytes())
          .parseClaimsJws(jwt)
          .getBody();
      String subject = claims.getSubject(); // "1234567890"
  } catch (Exception e) {
      // Token không hợp lệ (hết hạn, giả mạo, v.v.)
  }
  ```

#### **c. HS256 vs RS256**
- **HS256 (Symmetric)**:
  - Dùng cùng một khóa bí mật để ký và xác minh.
  - Ưu điểm: Nhanh, đơn giản.
  - Nhược điểm: Cần chia sẻ khóa an toàn giữa các bên.
- **RS256 (Asymmetric)**:
  - Dùng private key để ký, public key để xác minh.
  - Ưu điểm: Public key có thể công khai, phù hợp với hệ thống phân tán (như OAuth2/OpenID Connect).
  - Nhược điểm: Chậm hơn HS256.
- **Ứng dụng**:
  - HS256: Hệ thống nội bộ.
  - RS256: OAuth2, OpenID Connect (JWKS endpoint cung cấp public key).

#### **d. Trong Spring Security**
- **Resource Server xác minh JWT**:
  ```java
  @Configuration
  @EnableWebSecurity
  public class SecurityConfig extends WebSecurityConfigurerAdapter {
      @Override
      protected void configure(HttpSecurity http) throws Exception {
          http.oauth2ResourceServer()
              .jwt()
              .decoder(jwtDecoder());
      }

      @Bean
      public JwtDecoder jwtDecoder() {
          return NimbusJwtDecoder.withSecretKey("your-256-bit-secret".getBytes()).build(); // HS256
          // Hoặc dùng RS256 với JWKS: NimbusJwtDecoder.withJwkSetUri("https://auth-server/.well-known/jwks.json").build();
      }
  }
  ```
- **Truy xuất thông tin**:
  ```java
  @GetMapping("/user")
  public String getUser(JwtAuthenticationToken token) {
      String subject = token.getToken().getSubject();
      return "User: " + subject;
  }
  ```

---

### **Tóm tắt**
- **Cấu trúc JWT**:
  - **Header**: Thuật toán và loại token.
  - **Payload**: Claim (dữ liệu người dùng).
  - **Signature**: Chữ ký đảm bảo toàn vẹn.
- **Signing**: 
  - Ký Header + Payload bằng HS256 (symmetric) hoặc RS256 (asymmetric).
- **Verification**: 
  - Kiểm tra chữ ký và claim để xác minh token.
