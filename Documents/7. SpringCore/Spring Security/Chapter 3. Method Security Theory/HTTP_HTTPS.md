
---
Hãy đi sâu vào **HTTP và HTTPS** trong bối cảnh bảo mật ứng dụng và Spring Security. Tôi sẽ giải thích cách HTTPS sử dụng SSL/TLS để mã hóa dữ liệu, cùng với các HTTP Header liên quan đến bảo mật như `X-Frame-Options` và `Content-Security-Policy`.

---

### **1. Cách HTTPS sử dụng SSL/TLS để mã hóa dữ liệu**

#### **a. HTTP vs HTTPS**
- **HTTP (HyperText Transfer Protocol)**:
  - Là giao thức truyền dữ liệu không mã hóa (plaintext), dễ bị nghe lén (man-in-the-middle attack) hoặc giả mạo.
- **HTTPS (HTTP Secure)**:
  - Là HTTP chạy trên tầng mã hóa SSL/TLS, đảm bảo dữ liệu được mã hóa, toàn vẹn, và xác thực nguồn gốc.

#### **b. SSL/TLS là gì?**
- **SSL (Secure Sockets Layer)**: Giao thức mã hóa cũ, hiện đã lỗi thời.
- **TLS (Transport Layer Security)**: Phiên bản cải tiến của SSL, tiêu chuẩn hiện nay (TLS 1.2, TLS 1.3).
- **Vai trò**:
  - **Mã hóa**: Bảo vệ dữ liệu giữa client và server.
  - **Xác thực**: Đảm bảo server là đáng tin cậy (qua chứng chỉ SSL).
  - **Toàn vẹn**: Ngăn dữ liệu bị thay đổi trong quá trình truyền.

#### **c. Cách HTTPS sử dụng SSL/TLS**
Quy trình thiết lập kết nối HTTPS (TLS Handshake) diễn ra như sau:
1. **Client Hello**:
   - Client (trình duyệt) gửi thông tin: phiên bản TLS hỗ trợ, danh sách cipher suites (thuật toán mã hóa), và một số ngẫu nhiên (client random).
2. **Server Hello**:
   - Server trả lời: phiên bản TLS được chọn, cipher suite, số ngẫu nhiên (server random), và chứng chỉ SSL (chứa public key).
3. **Xác minh chứng chỉ**:
   - Client kiểm tra chứng chỉ SSL của server (do Certificate Authority - CA cấp) để đảm bảo server là thật.
4. **Trao đổi khóa**:
   - Client tạo một **pre-master secret**, mã hóa bằng public key của server (lấy từ chứng chỉ), và gửi cho server.
   - Server giải mã pre-master secret bằng private key.
5. **Tạo khóa phiên (session key)**:
   - Cả client và server dùng pre-master secret, client random, server random để tạo ra một **session key** (khóa đối xứng).
6. **Mã hóa dữ liệu**:
   - Từ đây, mọi dữ liệu (request/response) được mã hóa bằng session key, thường dùng thuật toán như AES.
   - TLS cũng thêm kiểm tra toàn vẹn (HMAC) để phát hiện nếu dữ liệu bị thay đổi.

#### **d. Thuật toán liên quan**
- **Asymmetric Encryption**: RSA hoặc ECC để trao đổi khóa ban đầu.
- **Symmetric Encryption**: AES để mã hóa dữ liệu sau khi thiết lập session key.
- **Hashing**: SHA-256/SHA-384 để kiểm tra toàn vẹn.

#### **e. Trong Spring Security**
- Spring Security không trực tiếp triển khai SSL/TLS (đây là trách nhiệm của web server như Tomcat, Jetty), nhưng có thể ép buộc HTTPS:
  ```java
  @Configuration
  @EnableWebSecurity
  public class SecurityConfig extends WebSecurityConfigurerAdapter {
      @Override
      protected void configure(HttpSecurity http) throws Exception {
          http.requiresChannel()
              .anyRequest().requiresSecure(); // Ép buộc HTTPS cho mọi request
      }
  }
  ```
- Cấu hình server (ví dụ: Tomcat) để dùng HTTPS:
  - Thêm chứng chỉ SSL vào file `application.properties`:
    ```
    server.port=8443
    server.ssl.key-store=classpath:keystore.jks
    server.ssl.key-store-password=secret
    server.ssl.key-alias=mykey
    ```

---

### **2. HTTP Headers liên quan đến bảo mật**
HTTP Headers là các trường trong request/response giúp tăng cường bảo mật. Dưới đây là hai header quan trọng và cách dùng trong Spring Security:

#### **a. X-Frame-Options**
- **Mô tả**:
  - Ngăn trình duyệt nhúng trang web trong `<frame>`, `<iframe>`, hoặc `<object>`, nhằm chống tấn công **Clickjacking** (kẻ tấn công lừa người dùng nhấp vào nút ẩn).
- **Giá trị**:
  - `DENY`: Không cho phép nhúng dưới bất kỳ hình thức nào.
  - `SAMEORIGIN`: Chỉ cho phép nhúng từ cùng nguồn (same origin).
  - `ALLOW-FROM uri`: Cho phép từ nguồn cụ thể (ít dùng, không được hỗ trợ rộng rãi).
- **Ví dụ tấn công Clickjacking**:
  - Kẻ tấn công nhúng trang ngân hàng trong iframe vô hình, đặt nút giả để lừa người dùng chuyển tiền.
- **Trong Spring Security**:
  - Mặc định bật với giá trị `DENY`. Có thể tùy chỉnh:
    ```java
    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http.headers()
            .frameOptions()
            .sameOrigin(); // Chỉ cho phép nhúng từ cùng domain
    }
    ```
  - Hoặc tắt nếu không cần:
    ```java
    http.headers().frameOptions().disable();
    ```

#### **b. Content-Security-Policy (CSP)**
- **Mô tả**:
  - Kiểm soát nguồn tài nguyên (script, style, image, v.v.) mà trang web được phép tải, nhằm giảm nguy cơ tấn công **XSS (Cross-Site Scripting)**.
- **Giá trị**:
  - `default-src`: Nguồn mặc định cho tất cả tài nguyên.
  - `script-src`: Nguồn cho script.
  - `style-src`: Nguồn cho CSS.
  - Ví dụ: `Content-Security-Policy: default-src 'self'; script-src 'self' https://trusted.com`
    - Chỉ cho phép tải tài nguyên từ cùng domain (`'self'`) và script từ `https://trusted.com`.
- **Ví dụ tấn công XSS**:
  - Kẻ tấn công chèn `<script src="evil.com/malware.js">`. CSP chặn nếu nguồn không được phép.
- **Trong Spring Security**:
  - Cấu hình CSP:
    ```java
    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http.headers()
            .contentSecurityPolicy("default-src 'self'; script-src 'self' https://trusted.com");
    }
    ```
  - Kiểm tra header trong response:
    ```
    Content-Security-Policy: default-src 'self'; script-src 'self' https://trusted.com
    ```

#### **Các header bảo mật khác**
- **`Strict-Transport-Security (HSTS)`**:
  - Ép trình duyệt luôn dùng HTTPS trong tương lai.
  - Ví dụ: `Strict-Transport-Security: max-age=31536000; includeSubDomains`
  - Cấu hình:
    ```java
    http.headers().httpStrictTransportSecurity().includeSubDomains(true).maxAgeInSeconds(31536000);
    ```
- **`X-Content-Type-Options`**:
  - Ngăn trình duyệt đoán loại MIME (`nosniff`).
  - Cấu hình mặc định: `X-Content-Type-Options: nosniff`.
- **`X-XSS-Protection`**:
  - Kích hoạt bộ lọc XSS trong trình duyệt (dù đã lỗi thời ở browser mới).
  - Mặc định: `X-XSS-Protection: 1; mode=block`.

#### **Cấu hình tất cả header trong Spring Security**
```java
@Override
protected void configure(HttpSecurity http) throws Exception {
    http.headers()
        .frameOptions().sameOrigin()
        .contentSecurityPolicy("default-src 'self'")
        .httpStrictTransportSecurity().includeSubDomains(true).maxAgeInSeconds(31536000)
        .xssProtection().block(true)
        .contentTypeOptions();
}
```

---

### **Tóm tắt**
- **HTTPS với SSL/TLS**:
  - Mã hóa dữ liệu qua TLS Handshake: trao đổi khóa bất đối xứng (RSA) → mã hóa đối xứng (AES).
  - Spring Security hỗ trợ ép buộc HTTPS qua `requiresChannel`.
- **HTTP Headers bảo mật**:
  - `X-Frame-Options`: Chống Clickjacking.
  - `Content-Security-Policy`: Chống XSS.
  - Cấu hình dễ dàng qua `HttpSecurity.headers()`.
