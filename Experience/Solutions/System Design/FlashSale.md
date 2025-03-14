
---
### **Thiết Kế Hệ Thống Flash Sale**

Black Friday sắp đến, đồng nghĩa với việc chúng ta cần xây dựng một hệ thống có **độ đồng thời cực cao, tính sẵn sàng cao và phản hồi nhanh**. Để đạt được điều này, cần phải xem xét **từ frontend đến backend**.

Dưới đây là **các nguyên tắc thiết kế quan trọng** để đảm bảo hệ thống hoạt động mượt mà:

---

### **Nguyên tắc thiết kế hệ thống Flash Sale**

1. **Less is more (Ít hơn là tốt hơn)**
    
    - Giảm số lượng **thành phần trên trang web** để tải trang nhanh hơn.
    - **Giảm truy vấn dữ liệu** đến cơ sở dữ liệu để tránh quá tải.
    - **Giảm số lượng request đến backend** để hạn chế áp lực lên hệ thống.
    - **Giảm sự phụ thuộc giữa các hệ thống** để tránh lỗi lan truyền.
2. **Short critical path (Rút ngắn đường đi quan trọng)**
    
    - Giảm **số lần gọi giữa các service**, tránh tình trạng chờ đợi lâu.
    - Nếu có thể, **gộp nhiều logic vào một service duy nhất** để xử lý nhanh hơn.
3. **Async (Bất đồng bộ)**
    
    - Sử dụng **message queue (hàng đợi tin nhắn)** như Kafka, RabbitMQ để xử lý lượng đơn hàng cực lớn mà không làm sập hệ thống.
    - Giảm độ trễ bằng cách xử lý các tác vụ không quan trọng sau khi xác nhận đơn hàng (VD: gửi email thông báo sau).
4. **Isolation (Cô lập để tối ưu hóa hiệu suất)**
    
    - **Cô lập nội dung tĩnh** (hình ảnh, CSS, JS) để giảm tải cho server chính.
    - **Cô lập dữ liệu động** (giỏ hàng, tồn kho) để tăng tốc độ xử lý.
    - **Cô lập các sản phẩm hiếm** bằng cách sử dụng database riêng, tránh gây ảnh hưởng đến hệ thống chính.
5. **Overselling is bad (Bán vượt mức là điều tồi tệ)**
    
    - Kiểm soát chặt chẽ việc **giảm số lượng tồn kho** để tránh bán quá số lượng thực có.
    - Sử dụng **cơ chế khóa hàng tồn kho** ở mức database hoặc Redis để đảm bảo tính nhất quán.
    - Có thể dùng **hàng đợi (queue) để xếp hàng người mua**, chỉ xử lý đơn hàng khi chắc chắn có hàng.
6. **User experience is important (Trải nghiệm người dùng rất quan trọng)**
    
    - **Tránh trường hợp người dùng đặt hàng thành công nhưng sau đó bị hủy vì hết hàng**. Điều này tạo ra trải nghiệm rất tệ.
    - Xây dựng **cơ chế xác nhận tồn kho trước khi chấp nhận đơn hàng**.
    - Cập nhật trạng thái hàng tồn theo **thời gian thực** để tránh việc người dùng đặt hàng mà thực tế sản phẩm đã hết.

---

### **Tóm tắt**

- **Tối ưu frontend** để giảm tải request.
- **Tăng tốc backend** bằng cách rút ngắn luồng xử lý quan trọng.
- **Dùng hệ thống hàng đợi** để xử lý đơn hàng đồng thời lớn.
- **Cô lập dữ liệu** để tăng hiệu suất và bảo vệ hệ thống chính.
- **Kiểm soát chặt chẽ hàng tồn kho** để tránh bán quá mức.
- **Nâng cao trải nghiệm người dùng** bằng cách cập nhật trạng thái hàng chính xác.