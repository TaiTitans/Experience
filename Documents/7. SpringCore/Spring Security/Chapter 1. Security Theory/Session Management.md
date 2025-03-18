
---
Hãy đi sâu vào **Session Management** trong Spring Security. Tôi sẽ giải thích cách hoạt động của session trong ứng dụng web, hai mối đe dọa phổ biến là **Session Fixation** và **Session Hijacking**, cùng với cách Spring Security giúp phòng chống chúng.

---

### **1. Cách hoạt động của session trong ứng dụng web**
- **Session là gì?**
  - Session là cơ chế để lưu trữ thông tin trạng thái của người dùng giữa các request trong ứng dụng web. Vì HTTP là stateless (không nhớ trạng thái), session giúp duy trì thông tin như trạng thái đăng nhập, giỏ hàng, v.v.
  - Thông thường, session được quản lý qua **Session ID**, một chuỗi duy nhất được gửi qua cookie (thường là `JSESSIONID`) giữa client và server.
- **Cách hoạt động cơ bản**:
  1. Người dùng đăng nhập → Server tạo một session và gán một Session ID.
  2. Session ID được gửi về client qua cookie.
  3. Mỗi request sau đó, client gửi Session ID để server nhận diện người dùng.
  4. Server truy xuất dữ liệu session từ bộ nhớ (hoặc cơ sở dữ liệu nếu dùng session store phân tán).
- **Trong Spring Security**:
  - Khi người dùng xác thực thành công (qua form login, ví dụ), Spring Security tự động tạo session và lưu thông tin xác thực trong `SecurityContext`.
  - `SecurityContextHolder` sử dụng `HttpSession` để lưu trữ thông tin này giữa các request.
  - Ví dụ cấu hình form login cơ bản:
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
                .sessionManagement() // Quản lý session
                .sessionCreationPolicy(SessionCreationPolicy.IF_REQUIRED); // Tạo session khi cần
        }
    }
    ```

- **Các tùy chọn SessionCreationPolicy trong Spring Security**:
  - `ALWAYS`: Luôn tạo session cho mỗi request.
  - `IF_REQUIRED`: Chỉ tạo session khi cần (mặc định).
  - `NEVER`: Không tạo session, nhưng vẫn dùng nếu đã có.
  - `STATELESS`: Không dùng session (phù hợp với JWT).

---

### **2. Session Fixation**
- **Session Fixation là gì?**
  - Đây là một loại tấn công mà kẻ tấn công cố định (fix) Session ID của người dùng trước khi họ đăng nhập, sau đó sử dụng Session ID đó để truy cập tài khoản của nạn nhân sau khi đăng nhập thành công.
  - Quy trình tấn công:
    1. Kẻ tấn công lấy một Session ID hợp lệ (ví dụ: bằng cách truy cập trang web).
    2. Gửi Session ID này cho nạn nhân (thường qua URL hoặc cookie bị ép buộc).
    3. Nạn nhân đăng nhập bằng Session ID đó.
    4. Kẻ tấn công dùng cùng Session ID để giả danh nạn nhân.
- **Cách phòng chống trong Spring Security**:
  - Spring Security cung cấp cơ chế thay đổi Session ID sau khi đăng nhập để ngăn chặn Session Fixation.
  - Cấu hình mặc định:
    ```java
    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http.sessionManagement()
            .sessionFixation().migrateSession(); // Thay đổi Session ID sau đăng nhập
    }
    ```
  - Các tùy chọn `sessionFixation()`:
    - `migrateSession`: Tạo Session ID mới sau đăng nhập, sao chép dữ liệu từ session cũ.
    - `newSession`: Tạo Session ID mới, không giữ dữ liệu cũ.
    - `changeSessionId`: Chỉ thay đổi Session ID mà không tạo session mới (khuyến nghị từ Servlet 3.1+).
    - `none`: Không làm gì (không an toàn).
- **Thực hành**:
  - Để tối ưu, dùng `changeSessionId` nếu server hỗ trợ Servlet 3.1+:
    ```java
    http.sessionManagement()
        .sessionFixation().changeSessionId();
    ```

---

### **3. Session Hijacking**
- **Session Hijacking là gì?**
  - Đây là tấn công khi kẻ tấn công đánh cắp Session ID của người dùng để giả danh họ. Session ID có thể bị đánh cắp qua:
    - Nghe lén mạng (nếu không dùng HTTPS).
    - XSS (kẻ tấn công chèn script để lấy cookie).
    - Đánh cắp cookie qua các kênh không an toàn.
  - Ví dụ: Nếu `JSESSIONID` bị lộ qua mạng không mã hóa, kẻ tấn công có thể dùng nó để truy cập hệ thống.
- **Cách phòng chống trong Spring Security**:
  1. **Sử dụng HTTPS**:
     - Ép buộc tất cả request qua HTTPS để mã hóa Session ID:
       ```java
       http.requiresChannel()
           .anyRequest().requiresSecure(); // Chỉ chấp nhận HTTPS
       ```
  2. **Cấu hình cookie an toàn**:
     - Đặt cookie `HttpOnly` và `Secure` để ngăn XSS và truyền qua kênh không an toàn:
       ```java
       http.sessionManagement()
           .sessionAuthenticationStrategy(sessionAuthenticationStrategy());

       @Bean
       public SessionAuthenticationStrategy sessionAuthenticationStrategy() {
           return new CompositeSessionAuthenticationStrategy(
               Arrays.asList(
                   new ChangeSessionIdAuthenticationStrategy(),
                   new RegisterSessionAuthenticationStrategy(new SessionRegistryImpl())
               )
           );
       }

       // Cấu hình trong application.properties
       // server.servlet.session.cookie.http-only=true
       // server.servlet.session.cookie.secure=true
       ```
  3. **Giới hạn thời gian session**:
     - Đặt timeout để session hết hạn nếu không hoạt động:
       ```java
       http.sessionManagement()
           .maximumSessions(1) // Giới hạn 1 session mỗi user
           .and()
           .sessionCreationPolicy(SessionCreationPolicy.IF_REQUIRED)
           .invalidSessionUrl("/login"); // Chuyển hướng nếu session hết hạn
       ```
       - Hoặc trong `application.properties`:
         ```
         server.servlet.session.timeout=15m
         ```
  4. **Xử lý concurrent sessions**:
     - Ngăn người dùng đăng nhập nhiều session cùng lúc (phòng trường hợp kẻ tấn công dùng Session ID bị đánh cắp):
       ```java
       http.sessionManagement()
           .maximumSessions(1) // Chỉ cho phép 1 session
           .maxSessionsPreventsLogin(true) // Chặn đăng nhập mới nếu đã có session
           .expiredUrl("/login?expired"); // Chuyển hướng khi session hết hạn
       ```

---

### **Tóm tắt**
- **Cách hoạt động của session**: Session lưu trữ trạng thái qua Session ID, Spring Security tích hợp với `HttpSession` để quản lý xác thực.
- **Session Fixation**: Ngăn chặn bằng cách thay đổi Session ID sau đăng nhập (`changeSessionId` là lựa chọn tối ưu).
- **Session Hijacking**: Bảo vệ bằng HTTPS, cookie `HttpOnly`/`Secure`, giới hạn session, và quản lý concurrent sessions.

---
