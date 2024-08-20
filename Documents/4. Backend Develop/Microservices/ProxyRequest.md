
---
Trong lập trình và phát triển phần mềm, "proxy request" là thuật ngữ dùng để chỉ một yêu cầu (request) được gửi thông qua một máy chủ proxy thay vì trực tiếp từ máy khách (client) đến máy chủ đích. Proxy request thường được sử dụng để cải thiện bảo mật, ẩn danh, hoặc điều phối luồng dữ liệu.

### Các mục đích chính của proxy request:

1. **Ẩn danh và bảo mật**: Khi một yêu cầu được gửi qua một máy chủ proxy, địa chỉ IP thật của máy khách có thể được ẩn đi, giúp bảo vệ danh tính và thông tin của người dùng.
    
2. **Lọc và kiểm soát truy cập**: Proxy có thể được sử dụng để kiểm soát truy cập tới các tài nguyên trên mạng, chặn hoặc cho phép truy cập dựa trên các chính sách bảo mật.
    
3. **Cải thiện hiệu suất**: Proxy có thể lưu trữ các nội dung tạm thời (cache) để giảm tải cho máy chủ đích và cung cấp nội dung nhanh hơn cho người dùng.
    
4. **Bypass giới hạn địa lý**: Proxy có thể giúp người dùng truy cập nội dung bị giới hạn bởi vị trí địa lý.
    

### Cách thức hoạt động:

Khi một máy khách gửi một proxy request, yêu cầu này được chuyển đến máy chủ proxy trước. Sau đó, máy chủ proxy sẽ gửi yêu cầu này đến máy chủ đích, nhận phản hồi và gửi lại phản hồi này về máy khách ban đầu.

Ví dụ: Trong một ứng dụng web, có thể sử dụng proxy để gửi các yêu cầu API thông qua một máy chủ trung gian, giúp bảo vệ API backend  hoặc xử lý các vấn đề liên quan đến CORS (Cross-Origin Resource Sharing).