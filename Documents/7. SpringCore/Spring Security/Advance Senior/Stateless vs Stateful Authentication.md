
---

Hãy đi sâu vào **Stateless vs Stateful Authentication** trong bối cảnh bảo mật ứng dụng và cách Spring Security hỗ trợ chúng. Tôi sẽ giải thích sự khác biệt giữa session-based (stateful) và token-based (stateless), cùng với ưu và nhược điểm của chúng trong hệ thống phân tán.

---

### **1. Sự khác biệt giữa Session-based và Token-based**

#### **a. Session-based Authentication (Stateful)**
- **Mô tả**:
  - Dựa trên việc lưu trữ trạng thái (state) của người dùng trên server, thường qua **session**.
  - Khi người dùng đăng nhập, server tạo một **Session ID** (ví dụ: `JSESSIONID`), lưu thông tin xác thực trong session store (như bộ nhớ hoặc cơ sở dữ liệu), và gửi Session ID về client qua cookie.
- **Cách hoạt động**:
  1. Người dùng đăng nhập → Server tạo session và lưu `Authentication` trong `HttpSession`.
  2. Client gửi Session ID trong mỗi request (qua cookie).
  3. Server kiểm tra Session ID để lấy thông tin xác thực từ session store.
- **Trong Spring Security**:
  - Mặc định sử dụng session thông qua `SecurityContextPersistenceFilter`.
  - Ví dụ cấu hình:
    ```java
    @Configuration
    @EnableWebSecurity
    public class SecurityConfig extends WebSecurityConfigurerAdapter {
        @Override
        protected void configure(HttpSecurity http) throws Exception {
            http.authorizeRequests()
                .anyRequest().authenticated()
                .and()
                .formLogin()
                .and()
                .sessionManagement()
                .sessionCreationPolicy(SessionCreationPolicy.IF_REQUIRED); // Stateful
        }
    }
    ```

#### **b. Token-based Authentication (Stateless)**
- **Mô tả**:
  - Không lưu trạng thái trên server. Thay vào đó, thông tin xác thực được mã hóa trong một **token** (thường là JWT) và gửi qua mỗi request từ client.
  - Server chỉ cần xác minh token mà không cần tra cứu session store.
- **Cách hoạt động**:
  1. Người dùng đăng nhập → Server tạo JWT chứa thông tin (username, roles, expiration) và gửi về client.
  2. Client gửi JWT trong header (VD: `Authorization: Bearer <token>`) với mỗi request.
  3. Server xác minh JWT (kiểm tra chữ ký, expiration) để xác thực.
- **Trong Spring Security**:
  - Dùng `SessionCreationPolicy.STATELESS` và custom filter:
    ```java
    @Configuration
    @EnableWebSecurity
    public class SecurityConfig extends WebSecurityConfigurerAdapter {
        @Override
        protected void configure(HttpSecurity http) throws Exception {
            http.sessionManagement()
                .sessionCreationPolicy(SessionCreationPolicy.STATELESS)
                .and()
                .addFilterBefore(new JwtAuthenticationFilter(), UsernamePasswordAuthenticationFilter.class)
                .authorizeRequests()
                .anyRequest().authenticated();
        }
    }
    ```
  - `JwtAuthenticationFilter` xử lý và xác minh JWT:
    ```java
    public class JwtAuthenticationFilter extends OncePerRequestFilter {
        @Override
        protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain)
                throws ServletException, IOException {
            String token = request.getHeader("Authorization");
            if (token != null && token.startsWith("Bearer ")) {
                token = token.substring(7);
                // Xác minh token (dùng thư viện như jjwt)
                Authentication auth = new UsernamePasswordAuthenticationToken("user", null, 
                    Arrays.asList(new SimpleGrantedAuthority("ROLE_USER")));
                SecurityContextHolder.getContext().setAuthentication(auth);
            }
            chain.doFilter(request, response);
        }
    }
    ```

#### **So sánh chi tiết**
| Tiêu chí              | Session-based (Stateful)          | Token-based (Stateless)         |
|-----------------------|-----------------------------------|---------------------------------|
| **Lưu trữ trạng thái**| Server (session store)           | Client (trong token)           |
| **Cách truyền**       | Cookie (Session ID)              | Header (JWT)                   |
| **Xác thực**          | Tra cứu session store            | Xác minh token (chữ ký)        |
| **Tính phụ thuộc**    | Phụ thuộc server                 | Độc lập với server             |
| **Ví dụ công nghệ**   | HttpSession, Spring Session      | JWT, OAuth2 Access Token       |

---

### **2. Ưu nhược điểm trong hệ thống phân tán**

#### **a. Session-based (Stateful)**
- **Ưu điểm**:
  1. **Kiểm soát tập trung**: Server có thể dễ dàng hủy session (logout) hoặc giới hạn số session (concurrent sessions).
     - Ví dụ: `http.sessionManagement().maximumSessions(1)`.
  2. **Dễ triển khai trong hệ thống đơn giản**: Không cần logic phức tạp phía client.
  3. **Tích hợp tốt với web truyền thống**: Cookie tự động xử lý bởi trình duyệt.
- **Nhược điểm trong hệ thống phân tán**:
  1. **Đồng bộ session**: 
     - Trong hệ thống phân tán (nhiều server), session phải được đồng bộ qua các node (dùng Redis, database), gây phức tạp và giảm hiệu suất.
     - Ví dụ: Dùng Spring Session với Redis:
       ```java
       @EnableRedisHttpSession
       public class SessionConfig {
           @Bean
           public RedisConnectionFactory redisConnectionFactory() {
               return new JedisConnectionFactory();
           }
       }
       ```
  2. **Tải cân bằng (load balancing)**: 
     - Cần sticky sessions (gắn session với một server) hoặc lưu trữ tập trung, làm tăng chi phí vận hành.
  3. **Hiệu suất**: Tra cứu session store có thể gây độ trễ trong hệ thống lớn.

#### **b. Token-based (Stateless)**
- **Ưu điểm trong hệ thống phân tán**:
  1. **Không cần lưu trữ server-side**: 
     - Mỗi server chỉ cần xác minh token (dùng public key với RS256 hoặc secret với HS256), không cần đồng bộ session.
  2. **Dễ mở rộng**: 
     - Phù hợp với microservices, API Gateway, hoặc hệ thống phân tán vì không phụ thuộc vào trạng thái server.
  3. **Hiệu suất cao**: 
     - Xác minh token nhanh hơn tra cứu session store, đặc biệt trong hệ thống lớn.
  4. **Tính di động**: 
     - Token có thể dùng trên nhiều client (web, mobile) mà không cần thay đổi logic server.
- **Nhược điểm**:
  1. **Không thể hủy tức thì**: 
     - Vì không lưu trạng thái, server không thể vô hiệu hóa token trước khi hết hạn. Cần blacklist (danh sách đen) hoặc giảm thời gian sống của token.
  2. **Kích thước token**: 
     - JWT có thể lớn (do chứa nhiều claim), tăng dung lượng header trong mỗi request.
  3. **Bảo mật khóa**: 
     - Với HS256, secret key phải được chia sẻ an toàn; với RS256, cần quản lý cặp khóa public/private.

#### **So sánh trong hệ thống phân tán**
| Tiêu chí               | Session-based (Stateful)       | Token-based (Stateless)       |
|-----------------------|--------------------------------|-------------------------------|
| **Đồng bộ trạng thái**| Cần (Redis, DB)               | Không cần                    |
| **Hiệu suất**         | Thấp hơn (tra cứu session)    | Cao hơn (chỉ xác minh)       |
| **Mở rộng**           | Khó (sticky session hoặc DB)  | Dễ (mỗi node độc lập)        |
| **Hủy token/session** | Dễ (xóa session)             | Khó (cần blacklist)          |
| **Ứng dụng phù hợp**  | Web truyền thống             | Microservices, API, mobile   |

---

### **3. Ứng dụng thực tế trong Spring Security**
- **Stateful (Session-based)**:
  - Phù hợp với ứng dụng web monolothic, nơi session được quản lý tập trung.
  - Ví dụ: Ứng dụng quản lý nội bộ với form login và logout dễ kiểm soát.
- **Stateless (Token-based)**:
  - Phù hợp với hệ thống phân tán, microservices, hoặc API RESTful.
  - Ví dụ: Xây dựng API cho mobile app, dùng JWT:
    ```java
    @RestController
    public class AuthController {
        @PostMapping("/login")
        public String login(@RequestBody UserCredentials creds) {
            // Giả lập xác thực
            return Jwts.builder()
                .setSubject(creds.getUsername())
                .signWith(SignatureAlgorithm.HS256, "secret".getBytes())
                .compact();
        }
    }
    ```

---

### **Tóm tắt**
- **Session-based (Stateful)**:
  - Lưu trạng thái trên server, dùng session ID.
  - Ưu: Dễ kiểm soát, nhược: Khó mở rộng trong hệ thống phân tán.
- **Token-based (Stateless)**:
  - Lưu thông tin trong token, không cần server-side state.
  - Ưu: Dễ mở rộng, nhược: Khó hủy token tức thì.
