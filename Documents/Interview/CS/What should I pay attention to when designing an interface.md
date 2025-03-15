
---
### **Xác minh và tối ưu hóa API trong thiết kế hệ thống**

1. **Xác minh tham số đầu vào**  
    API cần kiểm tra tính hợp lệ của các tham số, bao gồm việc xác định xem tham số có được phép để trống không và liệu độ dài của tham số có đáp ứng yêu cầu hay không.
    
2. **Đảm bảo khả năng mở rộng của API**  
    Khi thiết kế API, cần cân nhắc khả năng mở rộng để đảm bảo API có thể được **tái sử dụng** và dễ bảo trì. Hãy xem xét cách thiết kế API sao cho linh hoạt và có thể mở rộng trong tương lai.
    
3. **Chuyển đổi từ gọi tuần tự sang gọi song song**  
    Nếu API yêu cầu truy xuất nhiều nguồn dữ liệu khác nhau (ví dụ: thông tin sản phẩm, thông tin khuyến mãi, thông tin người dùng trên trang chủ của một sàn thương mại điện tử), việc thực hiện tuần tự có thể làm tăng thời gian phản hồi. Thay vào đó, sử dụng **gọi song song (parallel calls)** để giảm độ trễ và tăng hiệu suất.
    
4. **Ngăn chặn xử lý trùng lặp**  
    Nếu API thực hiện các thao tác ghi trên cơ sở dữ liệu, cần xem xét **cơ chế chống trùng lặp**. Một cách tiếp cận phổ biến là sử dụng **bảng chống trùng lặp (anti-duplication table)** với một **số serial duy nhất (unique serial number)** làm khóa chính.
    
5. **Ghi log đầy đủ**
    
    - Cần ghi nhận **cả tham số đầu vào và đầu ra** của API để dễ dàng theo dõi và xử lý sự cố.
    - Ghi nhận **thời gian thực thi** của API để đánh giá hiệu suất.
    - Lưu trữ log đầy đủ để hỗ trợ quá trình **debug và phân tích lỗi**.
6. **Đảm bảo tính tương thích khi cập nhật API**  
    Khi chỉnh sửa hoặc nâng cấp API, cần đảm bảo **khả năng tương thích ngược (backward compatibility)** để tránh ảnh hưởng đến các hệ thống đang sử dụng API cũ.
    
7. **Xử lý ngoại lệ một cách an toàn**
    
    - Sử dụng **khối `finally`** để đóng tài nguyên (streams, connections) nhằm tránh rò rỉ bộ nhớ.
    - Ghi log lỗi thay vì sử dụng `e.printStackTrace()` để dễ dàng theo dõi lỗi trong môi trường sản xuất.
    - Không "nuốt" ngoại lệ (swallow exceptions), cần **bắt và xử lý hợp lý** để đảm bảo API hoạt động ổn định.
8. **Xem xét giới hạn tốc độ (Rate Limiting)**  
    Trong các hệ thống có lưu lượng truy cập cao, cần triển khai **cơ chế giới hạn tốc độ (throttling)** để bảo vệ hệ thống khỏi bị quá tải trong các thời điểm lưu lượng tăng đột biến.