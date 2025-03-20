
---
Hãy đi sâu vào **Cryptographic Key Management (Quản lý khóa mã hóa)** trong bối cảnh OAuth2 và JWT, cùng với cách áp dụng nó trong Spring Security. Tôi sẽ giải thích cách quản lý khóa bí mật (Secret Key) và cặp khóa công khai/riêng tư (Public/Private Key), các nguyên tắc quan trọng, và cách triển khai thực tế.

---

### **1. Quản lý khóa bí mật trong OAuth2/JWT**

#### **a. Các loại khóa**
- **Secret Key (Khóa bí mật đối xứng)**:
  - Dùng trong các thuật toán như HMAC (VD: HS256 trong JWT).
  - Là một chuỗi bí mật duy nhất, được dùng để cả ký (signing) và xác minh (verification) token.
  - Ví dụ: `"my-256-bit-secret"`.
- **Public/Private Key (Khóa công khai/riêng tư bất đối xứng)**:
  - Dùng trong các thuật toán như RSA (VD: RS256 trong JWT).
  - **Private Key**: Dùng để ký token, giữ bí mật.
  - **Public Key**: Dùng để xác minh token, có thể công khai.
  - Ứng dụng phổ biến trong OAuth2/OpenID Connect.

#### **b. Vai trò trong OAuth2/JWT**
- **OAuth2**:
  - **Client Secret**: Khóa bí mật đối xứng giữa Client và Authorization Server để xác thực Client (VD: trong Authorization Code Grant).
  - **Access Token/Refresh Token**: Thường là JWT, được ký bằng Secret Key hoặc Private Key bởi Authorization Server.
- **JWT**:
  - **Signing**: Secret Key (HS256) hoặc Private Key (RS256) tạo chữ ký để đảm bảo tính toàn vẹn.
  - **Verification**: Secret Key hoặc Public Key xác minh token trước khi chấp nhận.

#### **c. Nguyên tắc quản lý khóa**
1. **Tính bí mật**:
   - Secret Key và Private Key phải được bảo vệ tuyệt đối, không để lộ trong mã nguồn hoặc log.
2. **Tính duy nhất**:
   - Mỗi ứng dụng/service nên có khóa riêng để hạn chế rủi ro nếu một khóa bị xâm phạm.
3. **Luân chuyển (Rotation)**:
   - Thay đổi khóa định kỳ để giảm nguy cơ bị tấn công lâu dài.
4. **Phân phối an toàn**:
   - Truyền khóa qua kênh mã hóa (HTTPS, SSH), tránh lưu trữ plaintext.
5. **Quyền tối thiểu**:
   - Chỉ cấp quyền truy cập khóa cho các thành phần cần thiết (VD: Authorization Server ký token, Resource Server xác minh).

---

### **2. Quản lý khóa bí mật (Secret Key)**

#### **a. Cách lưu trữ**
- **Trong mã nguồn (không khuyến khích)**:
  ```java
  String secretKey = "my-256-bit-secret"; // Dễ bị lộ, chỉ dùng cho test
  ```
- **Environment Variables**:
  - Lưu trong biến môi trường của hệ điều hành hoặc container:
    ```bash
    export JWT_SECRET=my-256-bit-secret
    ```
  - Truy xuất trong Spring Boot:
    ```java
    @Value("${jwt.secret}")
    private String secretKey;
    ```
- **Spring Vault hoặc Secret Management Tools**:
  - Dùng HashiCorp Vault, AWS Secrets Manager, hoặc Azure Key Vault để lưu trữ và truy xuất khóa an toàn.
  - Ví dụ với Spring Vault:
  ```properties
    spring.cloud.vault.host=vault-server
    spring.cloud.vault.token=your-vault-token
    spring.cloud.vault.secret.path=secret/my-app```

    ```java
    @Value("${jwt.secret}")
    private String secretKey;
    ```
#### **b. Tạo và ký JWT với Secret Key**
- **Ví dụ (HS256)**:
  ```java
  import io.jsonwebtoken.Jwts;
  import io.jsonwebtoken.SignatureAlgorithm;

  @Service
  public class JwtService {
      @Value("${jwt.secret}")
      private String secretKey;

      public String generateToken(String username) {
          return Jwts.builder()
              .setSubject(username)
              .setExpiration(new Date(System.currentTimeMillis() + 15 * 60 * 1000)) // 15 phút
              .signWith(SignatureAlgorithm.HS256, secretKey.getBytes())
              .compact();
      }
  }
  ```

#### **c. Xác minh JWT với Secret Key**
- **Ví dụ**:
  ```java
  public Authentication verifyToken(String token) {
      try {
          Claims claims = Jwts.parser()
              .setSigningKey(secretKey.getBytes())
              .parseClaimsJws(token)
              .getBody();
          return new UsernamePasswordAuthenticationToken(claims.getSubject(), null, 
              Arrays.asList(new SimpleGrantedAuthority("ROLE_USER")));
      } catch (Exception e) {
          throw new BadCredentialsException("Invalid token");
      }
  }
  ```

#### **d. Luân chuyển Secret Key**
- **Cách làm**:
  1. Tạo Secret Key mới.
  2. Cập nhật Secret Key trong hệ thống (environment variable, Vault).
  3. Hỗ trợ cả khóa cũ và mới trong giai đoạn chuyển tiếp:
     ```java
     private String newSecretKey = "new-256-bit-secret";
     private String oldSecretKey = "my-256-bit-secret";

     public Authentication verifyToken(String token) {
         try {
             return tryVerify(token, newSecretKey);
         } catch (Exception e) {
             return tryVerify(token, oldSecretKey); // Thử khóa cũ
         }
     }

     private Authentication tryVerify(String token, String key) {
         Claims claims = Jwts.parser()
             .setSigningKey(key.getBytes())
             .parseClaimsJws(token)
             .getBody();
         return new UsernamePasswordAuthenticationToken(claims.getSubject(), null, 
             Arrays.asList(new SimpleGrantedAuthority("ROLE_USER")));
     }
     ```
  4. Sau khi token cũ hết hạn, bỏ hỗ trợ khóa cũ.

---

### **3. Quản lý cặp khóa công khai/riêng tư (Public/Private Key)**

#### **a. Cách tạo**
- **Dùng OpenSSL**:
  ```bash
  # Tạo private key
  openssl genrsa -out private.pem 2048
  # Tạo public key từ private key
  openssl rsa -in private.pem -pubout -out public.pem
  ```
- **Dùng Java**:
  ```java
  KeyPairGenerator generator = KeyPairGenerator.getInstance("RSA");
  generator.initialize(2048);
  KeyPair pair = generator.generateKeyPair();
  PrivateKey privateKey = pair.getPrivate();
  PublicKey publicKey = pair.getPublic();
  ```

#### **b. Lưu trữ**
- **Private Key**:
  - Lưu trong Vault hoặc file bảo mật (chỉ Authorization Server truy cập).
  - Ví dụ: `/etc/secrets/private.pem` với quyền `600`.
- **Public Key**:
  - Công khai qua endpoint JWKS (JSON Web Key Set) của Authorization Server:
    ```
    GET https://auth-server/.well-known/jwks.json
    ```
    ```json
    {
      "keys": [
        {
          "kty": "RSA",
          "kid": "key-1",
          "n": "base64-encoded-modulus",
          "e": "AQAB"
        }
      ]
    }
    ```

#### **c. Ký JWT với Private Key**
- **Ví dụ (RS256)**:
  ```java
  import io.jsonwebtoken.Jwts;

  @Service
  public class JwtService {
      @Value("${jwt.private-key}")
      private String privateKeyPem;

      public String generateToken(String username) throws Exception {
          PrivateKey privateKey = loadPrivateKey(privateKeyPem);
          return Jwts.builder()
              .setSubject(username)
              .setExpiration(new Date(System.currentTimeMillis() + 15 * 60 * 1000))
              .signWith(privateKey, SignatureAlgorithm.RS256)
              .compact();
      }

      private PrivateKey loadPrivateKey(String pem) throws Exception {
          // Logic đọc file PEM và chuyển thành PrivateKey
          String key = pem.replace("-----BEGIN PRIVATE KEY-----", "")
                         .replace("-----END PRIVATE KEY-----", "")
                         .replaceAll("\\s", "");
          byte[] decoded = Base64.getDecoder().decode(key);
          return KeyFactory.getInstance("RSA").generatePrivate(new PKCS8EncodedKeySpec(decoded));
      }
  }
  ```

#### **d. Xác minh JWT với Public Key**
- **Ví dụ**:
  ```java
  @Bean
  public JwtDecoder jwtDecoder() {
      return NimbusJwtDecoder.withJwkSetUri("https://auth-server/.well-known/jwks.json").build();
  }
  ```
  - Spring tự động tải public key từ JWKS và xác minh token.

#### **e. Luân chuyển Public/Private Key**
- **Cách làm**:
  1. Tạo cặp khóa mới.
  2. Cập nhật Private Key trên Authorization Server và thêm Public Key mới vào JWKS.
  3. Giữ Public Key cũ trong JWKS một thời gian để hỗ trợ token cũ.
  4. Khi token cũ hết hạn, xóa Public Key cũ khỏi JWKS.

---

### **4. Trong Spring Security**
- **OAuth2 Client Secret**:

```properties
  spring.security.oauth2.client.registration.my-client.client-secret=${CLIENT_SECRET}
```


- **JWT Secret Key**:
  ```java
  @Bean
  public JwtDecoder jwtDecoder(@Value("${jwt.secret}") String secret) {
      return NimbusJwtDecoder.withSecretKey(secret.getBytes()).build();
  }
  ```
- **JWT Public Key (JWKS)**:
  ```java
  @Bean
  public JwtDecoder jwtDecoder() {
      return NimbusJwtDecoder.withJwkSetUri("https://auth-server/.well-known/jwks.json").build();
  }
  ```

---

### **Tóm tắt**
- **Secret Key**:
  - Dùng HS256, lưu trong environment variable hoặc Vault, hỗ trợ luân chuyển.
- **Public/Private Key**:
  - Dùng RS256, Private Key lưu an toàn, Public Key công khai qua JWKS, hỗ trợ luân chuyển.
- **Spring Security**: Tích hợp quản lý khóa qua properties, Vault, hoặc JWKS endpoint.

