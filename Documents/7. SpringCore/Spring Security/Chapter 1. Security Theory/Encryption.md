

---
Hãy đi sâu vào **Mã hóa và bảo mật dữ liệu**, một phần quan trọng trong bảo mật ứng dụng và Spring Security. Tôi sẽ giải thích từng mục bạn liệt kê: mã hóa đối xứng và bất đối xứng, các thuật toán phổ biến (AES, RSA, SHA, BCrypt), và chữ ký số (Digital Signatures) cùng ứng dụng trong JWT.

---

### **1. Symmetric vs Asymmetric Encryption (Mã hóa đối xứng và bất đối xứng)**

#### **a. Symmetric Encryption (Mã hóa đối xứng)**
- **Định nghĩa**:
  - Sử dụng cùng một khóa (key) để cả mã hóa và giải mã dữ liệu.
  - Ví dụ: Bạn mã hóa "Hello" thành "XyZk" bằng khóa "123", và dùng chính "123" để giải mã ngược lại.
- **Ưu điểm**:
  - Nhanh, hiệu quả với dữ liệu lớn.
  - Dễ triển khai khi hai bên đã chia sẻ khóa an toàn.
- **Nhược điểm**:
  - Vấn đề chia sẻ khóa: Nếu khóa bị lộ, dữ liệu không còn an toàn.
- **Ứng dụng**:
  - Mã hóa dữ liệu nhạy cảm trong cơ sở dữ liệu.
  - Bảo vệ giao tiếp trong HTTPS (sau khi trao đổi khóa qua asymmetric encryption).
- **Trong Spring Security**:
  - Dùng để mã hóa nội dung session hoặc dữ liệu tạm thời, nhưng ít phổ biến hơn asymmetric trong bảo mật xác thực.

#### **b. Asymmetric Encryption (Mã hóa bất đối xứng)**
- **Định nghĩa**:
  - Sử dụng cặp khóa: **public key** để mã hóa và **private key** để giải mã (hoặc ngược lại).
  - Ví dụ: Bạn gửi "Hello" mã hóa bằng public key của tôi, chỉ tôi (với private key) mới giải mã được.
- **Ưu điểm**:
  - An toàn hơn vì không cần chia sẻ khóa bí mật.
  - Hỗ trợ xác thực và chữ ký số.
- **Nhược điểm**:
  - Chậm hơn symmetric encryption, không phù hợp với dữ liệu lớn.
- **Ứng dụng**:
  - Trao đổi khóa trong HTTPS (SSL/TLS).
  - Xác thực và ký JWT.
- **Trong Spring Security**:
  - Được dùng trong OAuth2 hoặc JWT khi xác minh chữ ký số (dựa trên RSA).

#### **So sánh**:
| Tiêu chí            | Symmetric Encryption       | Asymmetric Encryption      |
|---------------------|----------------------------|----------------------------|
| **Khóa**            | 1 khóa chung              | Cặp khóa (public/private)  |
| **Tốc độ**          | Nhanh                     | Chậm                       |
| **Ứng dụng**        | Mã hóa dữ liệu lớn        | Trao đổi khóa, chữ ký số   |
| **Ví dụ thuật toán**| AES, DES                  | RSA, ECC                   |

---

### **2. Các thuật toán: AES, RSA, SHA, BCrypt**

#### **a. AES (Advanced Encryption Standard)**
- **Loại**: Symmetric Encryption.
- **Mô tả**:
  - Thuật toán mã hóa khối, sử dụng khóa 128, 192, hoặc 256 bit.
  - Được dùng rộng rãi trong SSL/TLS, mã hóa file, và bảo mật dữ liệu.
- **Trong Spring Security**:
  - Không trực tiếp tích hợp, nhưng có thể dùng để mã hóa dữ liệu nhạy cảm trong ứng dụng (ví dụ: thông tin trong session).
- **Ví dụ** (dùng Java `javax.crypto`):
  ```java
  SecretKey key = new SecretKeySpec("mysecretkey12345".getBytes(), "AES");
  Cipher cipher = Cipher.getInstance("AES");
  cipher.init(Cipher.ENCRYPT_MODE, key);
  byte[] encrypted = cipher.doFinal("Hello".getBytes());
  ```

#### **b. RSA (Rivest-Shamir-Adleman)**
- **Loại**: Asymmetric Encryption.
- **Mô tả**:
  - Dựa trên toán học số lớn (nhân và phân tích số nguyên tố).
  - Khóa public/private có độ dài phổ biến 2048 bit.
- **Ứng dụng**:
  - Mã hóa khóa đối xứng trong HTTPS.
  - Ký và xác minh chữ ký số trong JWT.
- **Trong Spring Security**:
  - Dùng trong OAuth2 Resource Server để xác minh JWT ký bằng RSA.
- **Ví dụ** (dùng thư viện `java.security`):
  ```java
  KeyPairGenerator generator = KeyPairGenerator.getInstance("RSA");
  generator.initialize(2048);
  KeyPair pair = generator.generateKeyPair();
  PublicKey publicKey = pair.getPublic();
  PrivateKey privateKey = pair.getPrivate();
  ```

#### **c. SHA (Secure Hash Algorithm)**
- **Loại**: Hashing (không mã hóa, chỉ tạo hàm băm).
- **Mô tả**:
  - Tạo chuỗi băm cố định từ dữ liệu đầu vào (SHA-1, SHA-256, SHA-512).
  - Không thể đảo ngược (one-way function).
- **Ứng dụng**:
  - Kiểm tra tính toàn vẹn dữ liệu.
  - Kết hợp với chữ ký số.
- **Trong Spring Security**:
  - Không dùng trực tiếp để mã hóa mật khẩu (vì thiếu salt), nhưng có thể dùng trong custom logic.

#### **d. BCrypt**
- **Loại**: Hashing (dành riêng cho mật khẩu).
- **Mô tả**:
  - Dựa trên Blowfish cipher, tự động thêm salt, điều chỉnh độ khó (work factor).
  - Ví dụ: "password" → `$2a$10$N9qo8uLOickgx2ZMRZoMye...`.
- **Ứng dụng**:
  - Lưu trữ mật khẩu an toàn trong cơ sở dữ liệu.
- **Trong Spring Security**:
  - Được tích hợp qua `BCryptPasswordEncoder`:
    ```java
    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }
    ```

---

### **3. Digital Signatures và ứng dụng trong JWT**
- **Digital Signatures là gì?**
  - Chữ ký số là cách sử dụng asymmetric encryption để xác minh tính toàn vẹn và nguồn gốc của dữ liệu.
  - Quy trình:
    1. Người gửi băm (hash) dữ liệu bằng SHA.
    2. Mã hóa hash bằng private key → tạo chữ ký.
    3. Người nhận giải mã chữ ký bằng public key, so sánh với hash của dữ liệu để xác minh.
- **Ứng dụng trong JWT**:
  - JWT (JSON Web Token) gồm 3 phần: **Header**, **Payload**, **Signature**.
    - **Header**: Thông tin thuật toán (VD: "HS256" hoặc "RS256").
    - **Payload**: Dữ liệu (VD: user ID, roles).
    - **Signature**: Chữ ký số để đảm bảo JWT không bị giả mạo.
  - Với RSA:
    - Server ký JWT bằng private key.
    - Client hoặc Resource Server xác minh bằng public key.
- **Ví dụ trong Spring Security**:
  - Cấu hình Resource Server để xác minh JWT ký bằng RSA:
    ```java
    @Configuration
    @EnableWebSecurity
    public class SecurityConfig extends WebSecurityConfigurerAdapter {
        @Override
        protected void configure(HttpSecurity http) throws Exception {
            http.authorizeRequests()
                .anyRequest().authenticated()
                .and()
                .oauth2ResourceServer()
                .jwt()
                .jwtAuthenticationConverter(jwtAuthenticationConverter());
        }

        @Bean
        public JwtDecoder jwtDecoder() {
            return NimbusJwtDecoder.withPublicKey(rsaPublicKey()).build();
        }

        private RSAPublicKey rsaPublicKey() throws Exception {
            // Load public key từ file hoặc config
            String publicKeyPEM = "-----BEGIN PUBLIC KEY-----\n...";
            return (RSAPublicKey) KeyFactory.getInstance("RSA")
                .generatePublic(new X509EncodedKeySpec(Base64.getDecoder().decode(publicKeyPEM)));
        }
    }
    ```
  - JWT được ký bằng private key ở Authorization Server (như Keycloak), và Spring Security xác minh bằng public key.

---

### **Tóm tắt**
- **Symmetric vs Asymmetric**:
  - Symmetric (AES): Nhanh, dùng 1 khóa.
  - Asymmetric (RSA): Chậm, dùng cặp khóa, phù hợp cho chữ ký và trao đổi khóa.
- **Thuật toán**:
  - AES: Mã hóa dữ liệu lớn.
  - RSA: Chữ ký và xác thực.
  - SHA: Hàm băm.
  - BCrypt: Mã hóa mật khẩu.
- **Digital Signatures trong JWT**:
  - Đảm bảo tính toàn vẹn và nguồn gốc của token, tích hợp chặt chẽ với Spring Security qua OAuth2/JWT.

