
---

**CORS** (Cross-Origin Resource Sharing) là một cơ chế bảo mật cho phép một trang web chạy ở một origin (domain) có thể yêu cầu tài nguyên từ một origin khác. Đây là một chính sách bảo mật của trình duyệt nhằm ngăn chặn các ứng dụng web thực hiện các yêu cầu không hợp lệ từ các domain khác nhau.

### Origin là gì?

- **Origin** bao gồm `protocol (HTTP/HTTPS)`, `domain`, và `port`. Ví dụ: `https://example.com:8080`.

### Tại sao cần CORS?

- Các trình duyệt thường áp dụng chính sách **Same-Origin Policy**, ngăn chặn các yêu cầu giữa các domain khác nhau để bảo vệ dữ liệu người dùng khỏi các tấn công như **CSRF (Cross-Site Request Forgery)**.
- Tuy nhiên, trong nhiều ứng dụng hiện đại, việc cần truy cập tài nguyên từ các origin khác là phổ biến, đặc biệt trong các ứng dụng API.

### Cách Hoạt Động của CORS

1. **Preflight Request (Tùy chọn)**
    
    - Trước khi gửi một yêu cầu thực sự (đặc biệt là các yêu cầu phức tạp như POST, PUT, DELETE), trình duyệt sẽ gửi một yêu cầu `OPTIONS` đến server để kiểm tra xem server có cho phép yêu cầu từ origin khác hay không.
    - Yêu cầu này kiểm tra các thông tin như:
        - Phương thức HTTP được phép (`Access-Control-Allow-Methods`).
        - Các header được phép (`Access-Control-Allow-Headers`).
        - Origin của yêu cầu (`Access-Control-Allow-Origin`).
2. **Simple Request**
    
    - Với các yêu cầu đơn giản (GET, POST với một số header cơ bản), trình duyệt có thể bỏ qua preflight và gửi yêu cầu trực tiếp. Server phản hồi sẽ bao gồm thông tin CORS.
3. **Response Headers**
    
    - Nếu server cho phép, nó sẽ trả về các header CORS tương ứng:
        - `Access-Control-Allow-Origin`: Chỉ định origin nào được phép (có thể là một domain cụ thể hoặc `*` để cho phép tất cả).
        - `Access-Control-Allow-Methods`: Liệt kê các phương thức HTTP được phép.
        - `Access-Control-Allow-Headers`: Liệt kê các header được phép.
        - `Access-Control-Allow-Credentials`: Cho phép gửi thông tin xác thực (cookies, HTTP authentication).

