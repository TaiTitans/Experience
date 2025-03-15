
---
### **Những yếu tố cần xem xét khi thiết kế API**

1. **Xác nhận giao thức mạng**
    
    - Xác định giao thức sử dụng: **HTTP/HTTPS** hay một giao thức riêng tùy chỉnh.
    - Nếu dùng **HTTPS**, cần đảm bảo chứng chỉ bảo mật hợp lệ để mã hóa dữ liệu.
2. **Thống nhất định dạng dữ liệu và phản hồi**
    
    - Xác định kiểu dữ liệu truyền tải, phổ biến là **application/json**.
    - Nếu giao tiếp giữa các hệ thống sử dụng ngôn ngữ khác nhau (**weakly typed vs strongly typed**), cần kiểm tra kỹ **việc chuyển đổi dữ liệu** để tránh lỗi.
3. **Đảm bảo bảo mật API**
    
    - Xác định phương thức xác thực:
        - **Token-based authentication (JWT, OAuth2, API Key, v.v.)**.
        - **Xác thực bằng chứng chỉ số (Certificate Authentication)** nếu cần độ bảo mật cao.
4. **Cơ chế retry khi API thất bại**
    
    - Xác định có cần **cơ chế retry tự động** không khi xảy ra lỗi mạng hoặc lỗi tạm thời của hệ thống.
    - Nếu có, cần quy định **số lần thử lại (retry count)** và **thời gian giữa các lần thử (backoff strategy)** để tránh quá tải hệ thống.
    - Đảm bảo **tính nhất quán cuối cùng (eventual consistency)** khi truyền dữ liệu.
5. **Ghi log đầy đủ để dễ debug**
    
    - Ghi lại **tham số đầu vào và đầu ra** của API.
    - Lưu lại **giá trị tham số sau khi được phân tích** để hỗ trợ xử lý lỗi.
    - Đảm bảo log chứa **thời gian thực thi, mã trạng thái HTTP, và nội dung phản hồi** để dễ dàng kiểm tra khi có sự cố.

➡ **Tóm lại**, khi thiết kế API, cần xem xét cả **bảo mật, tính ổn định, tính nhất quán và khả năng debug** để đảm bảo API hoạt động trơn tru và dễ bảo trì.