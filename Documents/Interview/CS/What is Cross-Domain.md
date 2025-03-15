
---
### **Hiểu đơn giản về vấn đề cross-domain (liên miền) trong web**

#### **1. Cross-domain là gì?**

Cross-domain (liên miền) xảy ra khi một trang web **yêu cầu tài nguyên từ một miền khác** (domain khác). Do chính sách cùng nguồn (Same-Origin Policy), những truy cập như vậy **thường bị chặn** để đảm bảo an toàn.

Tuy nhiên, trong nhiều trường hợp, **các yêu cầu liên miền là cần thiết**, đặc biệt là khi **frontend và backend được triển khai trên các miền khác nhau**. Điều này dẫn đến **vấn đề cross-domain** khi giao tiếp giữa client và server.

---

#### **2. Chính sách cùng nguồn (Same-Origin Policy) là gì?**

Chính sách cùng nguồn yêu cầu **ba thành phần sau phải giống nhau** để có thể truy cập tài nguyên mà không bị chặn:

- **Giao thức (Protocol):** `http://` hoặc `https://`
- **Tên miền (Domain):** `example.com`, `api.example.com`
- **Cổng (Port):** `:80`, `:443`, `:8080`, v.v.

Ngay cả khi hai miền cùng trỏ đến một **địa chỉ IP**, chúng vẫn bị coi là khác nguồn nếu có sự khác biệt trong ba yếu tố trên.

Ví dụ:

- `https://example.com` và `https://api.example.com` **khác nguồn** vì khác tên miền.
- `http://example.com` và `https://example.com` **khác nguồn** vì khác giao thức.
- `https://example.com:8080` và `https://example.com` **khác nguồn** vì khác cổng.

---

#### **3. Những hạn chế của chính sách cùng nguồn**

Chính sách này hạn chế các thao tác sau để đảm bảo bảo mật:

1. **Không thể đọc dữ liệu từ Cookie, LocalStorage và IndexDB** của miền khác.
2. **Không thể truy cập đối tượng DOM và JavaScript** từ miền khác.
3. **Không thể gửi AJAX request** đến một miền khác mà không có cơ chế hỗ trợ (như CORS).

---

#### **4. Tại sao cần có chính sách cùng nguồn?**

Chính sách này được thiết kế để **bảo vệ người dùng khỏi các cuộc tấn công bảo mật**.

Ví dụ:

- Bạn vừa **đăng nhập vào tài khoản ngân hàng** và kiểm tra số dư.
- Sau đó, bạn **truy cập một trang web độc hại** mà không hề hay biết.
- Nếu không có chính sách cùng nguồn, trang web này có thể **gửi yêu cầu đến ngân hàng của bạn**, lấy **cookie đăng nhập**, và đánh cắp thông tin tài khoản của bạn.

👉 **Nhờ có chính sách cùng nguồn, các trang web độc hại không thể truy cập dữ liệu quan trọng của bạn trên các trang web khác.**

---

📌 **Tóm tắt:**

- **Cross-domain xảy ra khi một trang web yêu cầu tài nguyên từ một miền khác.**
- **Chính sách cùng nguồn ngăn chặn các truy cập liên miền trái phép để đảm bảo bảo mật.**
- **Các hạn chế bao gồm không thể truy cập cookie, DOM, JavaScript hoặc gửi AJAX request đến miền khác.**
- **Chính sách này giúp bảo vệ dữ liệu người dùng khỏi các cuộc tấn công đánh cắp thông tin.**

---
## How to solve cross-domain problems?

### **Các phương pháp giải quyết vấn đề cross-domain (liên miền)**

Có một số cách để xử lý vấn đề cross-domain, bao gồm:

---

### **1. CORS (Cross-Origin Resource Sharing) – Chia sẻ tài nguyên liên miền**

CORS là một **cơ chế được quy định bởi trình duyệt**, giúp các trang web có thể truy cập tài nguyên từ một miền khác bằng cách thiết lập các **HTTP header** trên server.

Ví dụ: Giả sử **trang A** muốn truy cập dữ liệu từ **server B**. Nếu **server B khai báo** rằng **tên miền của trang A được phép truy cập**, thì trình duyệt sẽ cho phép request từ A đến B.

**Cách triển khai CORS trong Spring Boot:**

- Nếu đang sử dụng **Spring Boot**, bạn có thể thêm **annotation** `@CrossOrigin(origins = "*")` vào class Controller để cho phép truy cập liên miền.
- Annotation này cũng có thể được áp dụng **ở cấp phương thức** hoặc **cấp toàn bộ ứng dụng** để xử lý tất cả các API.
- **Lưu ý:** `@CrossOrigin` chỉ được hỗ trợ từ **Spring MVC 4.2 trở lên**.

---

### **2. Sử dụng Nginx làm reverse proxy để xử lý cross-domain**

Cách tiếp cận này dựa trên nguyên tắc:

- **Chính sách cùng nguồn chỉ áp dụng trên trình duyệt**, chứ **không phải là một phần của giao thức HTTP**.
- Khi server gọi HTTP API, nó không thực thi JavaScript nên không bị ảnh hưởng bởi chính sách cùng nguồn.

#### **Cách cấu hình reverse proxy với Nginx để xử lý cross-domain:**

1. **Cấu hình một proxy server** (cùng domain với trang chính nhưng khác port).
2. **Reverse proxy** truy cập đến server đích.
3. **Chỉnh sửa thông tin cookie**, giúp trình duyệt có thể ghi cookie vào domain hiện tại, hỗ trợ đăng nhập liên miền.

**Ví dụ cấu hình Nginx:**
```nginx
server {
    listen       81;
    server_name  www.domain1.com;

    location / {
        proxy_pass   http://www.domain2.com:8080;  # Reverse proxy đến domain2
        proxy_cookie_domain www.domain2.com www.domain1.com; # Chỉnh sửa cookie domain
        
        index  index.html index.htm;
        add_header Access-Control-Allow-Origin http://www.domain1.com;
    }
}
```

👉 **Lợi ích:**

- Frontend chỉ cần gửi request đến `http://www.domain1.com:81/*`, thay vì trực tiếp gọi `www.domain2.com:8080`.
- Nginx sẽ tự động **chuyển tiếp request** đến domain2 mà không vi phạm chính sách cùng nguồn.

---

### **3. Cross-domain bằng JSONP**

JSONP là một phương pháp **truyền thống** để thực hiện cross-domain mà **không cần CORS**.

**Cách hoạt động:**

- Các trình duyệt cho phép tải tài nguyên tĩnh (**JS, CSS, hình ảnh**) từ miền khác.
- Dựa vào nguyên tắc này, ta có thể tạo **thẻ `<script>` động** để gửi request đến server khác.

**Ví dụ:**
```html
<script src="http://api.example.com/data?callback=myCallback"></script>
```

**Server response:**
```js
myCallback({ "message": "Hello from another domain!" });
```

👉 **Lợi ích:**

- Không cần cấu hình server với CORS.
- Dễ dàng sử dụng với các API **chỉ hỗ trợ JSONP**.

**Hạn chế:**

- **Chỉ hỗ trợ HTTP GET**, không thể gửi dữ liệu qua POST, PUT, DELETE.
- Kém bảo mật hơn so với CORS do có nguy cơ bị tấn công XSS.

📌 **Kết luận:**

- Nếu **có quyền kiểm soát server**, sử dụng **CORS** là lựa chọn tốt nhất.
- Nếu **muốn proxy request từ một miền khác**, sử dụng **Nginx reverse proxy**.
- Nếu **chỉ cần GET request nhanh**, có thể sử dụng **JSONP** nhưng không được bảo mật.