

---

Eureka Server là một máy chủ đăng ký dịch vụ trong hệ thống Microservices. Nó đảm nhiệm việc đặt tên cho mỗi microservice. Tại sao cần đặt tên? Bởi vì khi có nhiều microservices được triển khai và hoạt động trên nhiều instance khác nhau, không cần phải mã hóa địa chỉ IP cứng của mỗi service, thay vào đó, chúng ta có thể sử dụng tên service đã đăng ký trên Eureka Server để tìm kiếm và truy cập các dịch vụ này. Điều này giúp cho việc quản lý và mở rộng các dịch vụ một cách dễ dàng và hiệu quả hơn.

Vì vậy, mỗi dịch vụ đăng ký với Eureka và gửi yêu cầu ping tới Eureka server để thông báo rằng nó đang hoạt động.

Nếu Eureka server không nhận được bất kỳ thông báo nào từ một dịch vụ, dịch vụ đó sẽ bị hủy đăng ký tự động từ Eureka server.


---

