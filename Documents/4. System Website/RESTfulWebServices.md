
---
**RESTful Web Services** (Representational State Transfer) là một cách để cung cấp và truy cập các dịch vụ web theo các nguyên tắc của REST. REST là một phong cách kiến trúc được sử dụng rộng rãi cho việc thiết kế các hệ thống mạng, đặc biệt là các dịch vụ web. RESTful Web Services sử dụng HTTP (HyperText Transfer Protocol) và các phương thức HTTP để thực hiện các thao tác CRUD (Create, Read, Update, Delete) trên các tài nguyên.


---
## Các nguyên tắc cơ bản của REST

1. **Stateless**: Mỗi yêu cầu từ client đến server phải chứa đủ thông tin để server có thể hiểu và xử lý nó hoàn toàn. Server không lưu trữ bất kỳ thông tin trạng thái nào giữa các yêu cầu.
    
2. **Client-Server**: Kiến trúc REST tách biệt client và server. Client chịu trách nhiệm về giao diện người dùng và trải nghiệm người dùng, trong khi server chịu trách nhiệm về lưu trữ dữ liệu và logic kinh doanh.
    
3. **Cacheable**: Các phản hồi từ server phải được đánh dấu có thể cache hoặc không. Khi có thể cache, client có thể lưu trữ phản hồi và sử dụng lại chúng mà không cần yêu cầu lại từ server, giúp tăng hiệu suất.
    
4. **Uniform Interface**: REST định nghĩa một giao diện thống nhất giữa các thành phần, đơn giản hóa và tách biệt kiến trúc tổng thể. Điều này đạt được thông qua các tài nguyên được xác định bởi URI, sử dụng các phương thức HTTP, và trao đổi thông tin qua các định dạng như JSON hoặc XML.
    
5. **Layered System**: Kiến trúc REST có thể được tổ chức thành các lớp để tăng cường khả năng mở rộng và bảo mật. Client không cần biết rằng nó đang làm việc với một server hay một proxy.
    
6. **Code on Demand (optional)**: Server có thể cung cấp mã thực thi cho client để mở rộng chức năng, mặc dù điều này không bắt buộc.
    
---
## Các phương thức HTTP chính trong REST

- **GET**: Lấy thông tin về một tài nguyên.
- **POST**: Tạo mới một tài nguyên.
- **PUT**: Cập nhật một tài nguyên hiện có.
- **DELETE**: Xóa một tài nguyên.


---
### **Stateless trong RESTful API có nghĩa là gì và tại sao nó quan trọng?**

**Gợi ý trả lời:** Stateless có nghĩa là mỗi yêu cầu từ client đến server phải chứa đủ thông tin để server có thể hiểu và xử lý yêu cầu mà không cần lưu trữ trạng thái giữa các yêu cầu. Điều này quan trọng vì:

- **Đơn giản hóa thiết kế server**: Server không cần quản lý trạng thái của client, giảm phức tạp.
- **Tăng khả năng mở rộng**: Dễ dàng phân phối tải giữa các server vì không cần duy trì phiên làm việc.
- **Bảo mật**: Giảm nguy cơ lưu trữ thông tin nhạy cảm trên server.

### **Làm thế nào để bảo mật RESTful API?**

**Gợi ý trả lời:**

- **Sử dụng HTTPS**: Bảo mật truyền thông qua việc mã hóa dữ liệu.
- **Xác thực và Ủy quyền**: Sử dụng các cơ chế xác thực như OAuth, JWT (JSON Web Tokens) để xác định và ủy quyền người dùng.
- **Sử dụng các phương thức HTTP đúng cách**: Đảm bảo chỉ cho phép các phương thức HTTP phù hợp với từng endpoint.
- **Giới hạn tốc độ (Rate Limiting)**: Ngăn chặn các cuộc tấn công DDoS bằng cách giới hạn số lượng yêu cầu từ một client trong một khoảng thời gian nhất định.
- **CORS (Cross-Origin Resource Sharing)**: Kiểm soát truy cập từ các nguồn gốc khác nhau để ngăn chặn các yêu cầu không hợp lệ.
- **Validating input**: Kiểm tra và xác thực dữ liệu đầu vào để ngăn chặn các cuộc tấn công Injection.