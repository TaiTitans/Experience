
---


Hãy đi sâu vào **Authorization (Phân quyền)** trong Spring Security. Tôi sẽ giải thích chi tiết định nghĩa, các mô hình phân quyền chính (RBAC và ABAC), và cách chúng được triển khai trong Spring Security.

---

### **1. Định nghĩa: Quyết định những gì người dùng được phép làm sau khi xác thực**
- **Authorization là gì?**
  - Sau khi xác thực (authentication) xác minh "bạn là ai", authorization quyết định "bạn được làm gì". Đây là quá trình kiểm soát quyền truy cập vào tài nguyên hoặc chức năng dựa trên danh tính, vai trò, hoặc thuộc tính của người dùng.
  - Ví dụ: Một người dùng có thể đăng nhập (authenticated), nhưng chỉ người dùng có vai trò "ADMIN" mới được phép xóa dữ liệu.
- **Trong Spring Security**:
  - Spring Security quản lý phân quyền thông qua các bộ lọc (filters) và các biểu thức (expressions) trong cấu hình `HttpSecurity` hoặc annotation như `@PreAuthorize`.
  - Quy trình: Sau khi xác thực, hệ thống kiểm tra quyền của người dùng trước khi cho phép truy cập tài nguyên.

---

### **2. Mô hình phân quyền**
Spring Security hỗ trợ nhiều mô hình phân quyền, nhưng hai mô hình phổ biến nhất là **RBAC** và **ABAC**. Hãy xem xét từng mô hình:

#### **a. Role-Based Access Control (RBAC)**
- **Ý nghĩa**:
  - Phân quyền dựa trên vai trò (roles) được gán cho người dùng. Mỗi vai trò có một tập hợp quyền (permissions) cụ thể.
  - Ví dụ: Vai trò "USER" chỉ được đọc dữ liệu, vai trò "ADMIN" được đọc và xóa dữ liệu.
- **Ưu điểm**:
  - Đơn giản, dễ quản lý trong hệ thống có số lượng vai trò cố định.
  - Phù hợp với ứng dụng nhỏ đến trung bình.
- **Nhược điểm**:
  - Thiếu linh hoạt khi cần phân quyền chi tiết hoặc theo ngữ cảnh.
- **Cách triển khai trong Spring Security**:
  - Sử dụng `GrantedAuthority` để đại diện cho vai trò (thường là chuỗi như "ROLE_USER", "ROLE_ADMIN").
  - Cấu hình phân quyền trong `HttpSecurity`:
    ```java
    @Configuration
    @EnableWebSecurity
    public class SecurityConfig extends WebSecurityConfigurerAdapter {
        @Override
        protected void configure(HttpSecurity http) throws Exception {
            http.authorizeRequests()
                .antMatchers("/user/**").hasRole("USER") // Chỉ USER truy cập
                .antMatchers("/admin/**").hasRole("ADMIN") // Chỉ ADMIN truy cập
                .anyRequest().authenticated()
                .and()
                .formLogin();
        }

        @Override
        protected void configure(AuthenticationManagerBuilder auth) throws Exception {
            auth.inMemoryAuthentication()
                .withUser("user").password("{noop}password").roles("USER")
                .and()
                .withUser("admin").password("{noop}password").roles("ADMIN");
        }
    }
    ```
  - Sử dụng annotation tại phương thức:
    ```java
    @PreAuthorize("hasRole('ADMIN')")
    public void deleteData() {
        // Chỉ ADMIN được gọi
    }
    ```

#### **b. Attribute-Based Access Control (ABAC)**
- **Ý nghĩa**:
  - Phân quyền dựa trên các thuộc tính (attributes) của người dùng, tài nguyên, hoặc ngữ cảnh, thay vì chỉ dựa vào vai trò.
  - Ví dụ: Một người dùng chỉ được chỉnh sửa tài liệu nếu thuộc tính "department = HR" và "time = working hours".
- **Ưu điểm**:
  - Linh hoạt hơn RBAC, phù hợp với hệ thống phức tạp hoặc phân quyền động.
- **Nhược điểm**:
  - Phức tạp hơn để triển khai và quản lý.
- **Cách triển khai trong Spring Security**:
  - Sử dụng biểu thức SpEL (Spring Expression Language) trong `@PreAuthorize` hoặc tùy chỉnh `PermissionEvaluator`.
  - Ví dụ với `@PreAuthorize`:
    ```java
    @PreAuthorize("hasAuthority('ROLE_USER') and #userId == authentication.principal.id")
    public void updateUserProfile(Long userId) {
        // Chỉ cho phép người dùng chỉnh sửa profile của chính họ
    }
    ```
  - Tùy chỉnh `PermissionEvaluator` để kiểm tra thuộc tính:
    ```java
    @Component
    public class CustomPermissionEvaluator implements PermissionEvaluator {
        @Override
        public boolean hasPermission(Authentication auth, Object targetDomainObject, Object permission) {
            if (targetDomainObject instanceof Document) {
                Document doc = (Document) targetDomainObject;
                User user = (User) auth.getPrincipal();
                // Kiểm tra thuộc tính: chỉ cho phép chỉnh sửa nếu user thuộc department của document
                return doc.getDepartment().equals(user.getDepartment());
            }
            return false;
        }

        @Override
        public boolean hasPermission(Authentication auth, Serializable targetId, String targetType, Object permission) {
            return false; // Không dùng trong ví dụ này
        }
    }
    ```
    - Đăng ký `PermissionEvaluator` trong cấu hình:
    ```java
    @Configuration
    @EnableGlobalMethodSecurity(prePostEnabled = true)
    public class MethodSecurityConfig extends GlobalMethodSecurityConfiguration {
        @Autowired
        private CustomPermissionEvaluator permissionEvaluator;

        @Override
        protected MethodSecurityExpressionHandler createExpressionHandler() {
            DefaultMethodSecurityExpressionHandler handler = new DefaultMethodSecurityExpressionHandler();
            handler.setPermissionEvaluator(permissionEvaluator);
            return handler;
        }
    }
    ```
    - Sử dụng trong annotation:
    ```java
    @PreAuthorize("hasPermission(#document, 'edit')")
    public void editDocument(Document document) {
        // Chỉ cho phép chỉnh sửa nếu PermissionEvaluator trả về true
    }
    ```

---

### **So sánh RBAC và ABAC trong Spring Security**
| Tiêu chí            | RBAC                          | ABAC                          |
|---------------------|-------------------------------|-------------------------------|
| **Cơ sở phân quyền**| Vai trò (roles)              | Thuộc tính (attributes)       |
| **Độ phức tạp**     | Đơn giản                    | Phức tạp hơn                  |
| **Tính linh hoạt**  | Thấp                        | Cao                           |
| **Ví dụ sử dụng**   | Quản lý user/admin          | Phân quyền theo ngữ cảnh (department, time) |
| **Cấu hình**        | `hasRole()`, `hasAuthority()` | `@PreAuthorize` với SpEL hoặc custom logic |

---

### **Thực hành trong Spring Security**
1. **RBAC**:
   - Thử cấu hình một ứng dụng với 2 vai trò (USER, ADMIN) và giới hạn truy cập vào các endpoint khác nhau.
   - Dùng `@Secured` hoặc `@PreAuthorize` để bảo vệ method.
2. **ABAC**:
   - Tạo một ứng dụng kiểm tra quyền dựa trên thuộc tính (ví dụ: chỉ cho phép chỉnh sửa nếu user sở hữu tài nguyên).
   - Thử viết một `PermissionEvaluator` tùy chỉnh.

---

### **Tóm tắt**
- **RBAC**: Dễ triển khai, dựa trên vai trò, phù hợp với hệ thống đơn giản.
- **ABAC**: Linh hoạt, dựa trên thuộc tính, cần tùy chỉnh nhiều hơn nhưng mạnh mẽ trong các kịch bản phức tạp.