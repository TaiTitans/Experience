
---

### **1. Định nghĩa: Quá trình xác minh danh tính của người dùng hoặc hệ thống**
- **Authentication là gì?**
  Authentication là bước xác định "bạn là ai" trong một hệ thống. Nó đảm bảo rằng người dùng (hoặc hệ thống khác) đang truy cập là hợp lệ, thông qua việc kiểm tra thông tin đăng nhập (credentials) như username/password, token, hoặc các yếu tố khác.
- **Trong Spring Security**:
  - Spring Security quản lý xác thực thông qua `AuthenticationManager`, một thành phần cốt lõi chịu trách nhiệm xử lý và xác minh thông tin đăng nhập.
  - Khi xác thực thành công, thông tin người dùng được lưu trong `SecurityContext`, cho phép hệ thống biết ai đang thực hiện yêu cầu.

---

### **2. Các phương pháp xác thực**
Dưới đây là ba phương pháp phổ biến và cách Spring Security hỗ trợ chúng:

#### **a. Username/Password**
- **Ý nghĩa**: Đây là cách truyền thống nhất, yêu cầu người dùng cung cấp tên đăng nhập và mật khẩu.
- **Cách hoạt động trong Spring Security**:
  - Sử dụng `UsernamePasswordAuthenticationFilter` để xử lý form đăng nhập.
  - Cấu hình một `UserDetailsService` để lấy thông tin người dùng từ cơ sở dữ liệu hoặc in-memory.
- **Ví dụ cơ bản**:
  ```java
  @Configuration
  @EnableWebSecurity
  public class SecurityConfig extends WebSecurityConfigurerAdapter {
      @Override
      protected void configure(AuthenticationManagerBuilder auth) throws Exception {
          auth.inMemoryAuthentication()
              .withUser("admin")
              .password("{noop}password") // {noop} để không mã hóa, thực tế dùng BCrypt
              .roles("ADMIN");
      }

      @Override
      protected void configure(HttpSecurity http) throws Exception {
          http.authorizeRequests()
              .antMatchers("/admin/**").hasRole("ADMIN")
              .and()
              .formLogin(); // Form đăng nhập mặc định
      }
  }
  ```
  - Ở đây, Spring Security tạo một form đăng nhập và kiểm tra username/password.

#### **b. Token-based (JWT - JSON Web Token)**
- **Ý nghĩa**: Thay vì lưu session trên server, hệ thống sử dụng token (JWT) để xác thực. Token được gửi trong header của mỗi request (thường là `Authorization: Bearer <token>`).
- **Cách hoạt động trong Spring Security**:
  - Sử dụng một custom filter (thường là `JwtAuthenticationFilter`) để kiểm tra và xác minh token.
  - Không dựa vào session, phù hợp với ứng dụng stateless (như REST API).
- **Ví dụ cơ bản**:
  ```java
  public class JwtAuthenticationFilter extends OncePerRequestFilter {
      @Override
      protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain)
              throws ServletException, IOException {
          String header = request.getHeader("Authorization");
          if (header != null && header.startsWith("Bearer ")) {
              String token = header.substring(7);
              // Xác minh token (dùng thư viện như jjwt)
              UsernamePasswordAuthenticationToken auth = new UsernamePasswordAuthenticationToken("user", null, Arrays.asList(new SimpleGrantedAuthority("ROLE_USER")));
              SecurityContextHolder.getContext().setAuthentication(auth);
          }
          chain.doFilter(request, response);
      }
  }
  ```
  - Trong thực tế, bạn cần tích hợp thư viện JWT (như `io.jsonwebtoken`) để tạo và xác minh token.

#### **c. Multi-Factor Authentication (MFA)**
- **Ý nghĩa**: Yêu cầu nhiều yếu tố để xác thực, ví dụ: mật khẩu + mã OTP (One-Time Password) gửi qua SMS hoặc ứng dụng như Google Authenticator.
- **Cách hoạt động trong Spring Security**:
  - Tùy chỉnh `AuthenticationProvider` để kiểm tra thêm yếu tố thứ hai sau khi xác thực username/password.
  - Thêm bước xác minh OTP vào quy trình.
- **Ví dụ cơ bản**:
  - Bạn có thể viết một `CustomAuthenticationProvider`:
  ```java
  @Component
  public class CustomAuthenticationProvider implements AuthenticationProvider {
      @Override
      public Authentication authenticate(Authentication authentication) throws AuthenticationException {
          String username = authentication.getName();
          String password = authentication.getCredentials().toString();
          // Kiểm tra username/password
          if (validUser(username, password)) {
              // Kiểm tra OTP (giả sử từ request hoặc DB)
              String otp = getOtpFromRequest();
              if (verifyOtp(username, otp)) {
                  return new UsernamePasswordAuthenticationToken(username, password, Arrays.asList(new SimpleGrantedAuthority("ROLE_USER")));
              }
          }
          throw new BadCredentialsException("Xác thực thất bại");
      }

      @Override
      public boolean supports(Class<?> authentication) {
          return UsernamePasswordAuthenticationToken.class.isAssignableFrom(authentication);
      }
  }
  ```
  - Sau đó, thêm provider này vào cấu hình Spring Security.

---

### **3. Lý thuyết liên quan: Hashing, Salt trong mã hóa mật khẩu**
- **Hashing là gì?**
  - Hashing biến mật khẩu thành một chuỗi cố định không thể đảo ngược (one-way function). Ví dụ: "password123" → "a1b2c3...".
  - Mục đích: Tránh lưu mật khẩu dạng plaintext trong cơ sở dữ liệu.
- **Salt là gì?**
  - Salt là một chuỗi ngẫu nhiên thêm vào mật khẩu trước khi hash, để đảm bảo cùng một mật khẩu không tạo ra cùng một hash (phòng chống tấn công Rainbow Table).
  - Ví dụ: "password123" + salt "xyz" → hash khác với "password123" + salt "abc".
- **Trong Spring Security**:
  - Spring Security sử dụng `PasswordEncoder` để mã hóa mật khẩu. Các triển khai phổ biến:
    - `BCryptPasswordEncoder`: Sử dụng BCrypt, tự động thêm salt.
    - `Pbkdf2PasswordEncoder`, `Argon2PasswordEncoder`: Các lựa chọn khác mạnh mẽ hơn.
- **Ví dụ sử dụng BCrypt**:
  ```java
  @Bean
  public PasswordEncoder passwordEncoder() {
      return new BCryptPasswordEncoder();
  }

  @Override
  protected void configure(AuthenticationManagerBuilder auth) throws Exception {
      auth.userDetailsService(customUserDetailsService)
          .passwordEncoder(passwordEncoder());
  }
  ```
  - Khi lưu user vào DB, bạn mã hóa mật khẩu:
  ```java
  String rawPassword = "password123";
  String encodedPassword = passwordEncoder().encode(rawPassword); // Lưu encodedPassword vào DB
  ```
  - Khi xác thực, Spring tự động so sánh mật khẩu người dùng nhập với bản mã hóa.

---

### **Tóm tắt**
- **Username/Password**: Dễ triển khai, phù hợp với ứng dụng web truyền thống.
- **JWT**: Phù hợp với API và hệ thống stateless.
- **MFA**: Nâng cao bảo mật, cần tùy chỉnh nhiều hơn.
- **Hashing/Salt**: Đảm bảo an toàn mật khẩu, BCrypt là lựa chọn mặc định tốt.
