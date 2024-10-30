

---
### Config Service - Spring Cloud Config

> Spring Cloud Config là một dịch vụ của Spring Cloud, cung cấp khả năng quản lý cấu hình trung tâm cho các ứng dụng phân tán. Nó cho phép các ứng dụng lấy cấu hình của chúng từ một vị trí tập trung thay vì phải lưu trữ cấu hình trực tiếp trong mã nguồn hoặc tệp tin. Các ứng dụng có thể truy cập cấu hình của chúng thông qua giao diện REST API, hoặc thông qua các thư viện khách được cung cấp bởi Spring Cloud Config. Sử dụng Spring Cloud Config, ta có thể dễ dàng quản lý và triển khai các ứng dụng phân tán một cách hiệu quả và an toàn hơn.

**Các tính năng của Spring Cloud Config Server:**

- API dựa trên HTTP và tài nguyên cho cấu hình bên ngoài (các cặp tên-giá trị hoặc tương đương nội dung YAML).
- Mã hóa và giải mã các giá trị thuộc tính (symmetric or asymmetric).
- Có thể dễ dàng nhúng vào một ứng dụng Spring Boot bằng cách sử dụng `@EnableConfigServer`.

**Các tính năng của Config Client (cho các ứng dụng Spring):**

- Kết nối đến Config Server và khởi tạo Spring Environment với các nguồn thuộc tính từ xa.
- Mã hóa (Encrypt) và giải mã (decrypt) các giá trị thuộc tính (symmetric or asymmetric).