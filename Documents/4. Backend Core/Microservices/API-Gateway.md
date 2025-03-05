
---

Một hệ thống microservices trung bình sẽ có một vài cho tới hàng trăm services khác nhau, nếu như client giao tiếp trực tiếp với các services này thì sơ đồ giao tiếp giữa client và hệ thống của chúng ta sẽ trông như một nồi cám lợn.

Chính vì cái nồi cám lợn trên cho nên mới xuất hiện một giải pháp đó chính là API Gateway (tạm dịch là cổng kết nối API) đóng vai trò là một cổng trung gian giữa client và hệ thống microservices đằng sau.

---
API Gateway có thể coi là một cổng trung gian, nó là cổng vào duy nhất tới hệ thống microservices của chúng ta, api gateway sẽ nhận các requests từ phía client, chỉnh sửa, xác thực và điều hướng chúng đến các API cụ thể trên các services phía sau. Khi này sơ đồ hệ thống của chúng ta sẽ trông như này.

![](https://images.viblo.asia/7c5add32-77c0-43c4-bb76-f87ec53badaa.png)

Ngoài nhiệm vụ chính là proxy request thì một hệ thống API Gateway thường sẽ đảm nhận luôn vài vai trò khác như bảo mật API, monitoring, analytics số lượng requests cũng như tình trạng hệ thống phía sau.

# Lợi ích của việc sử dụng API Gateway

- ### Che dấu được cấu trúc của hệ thống microservices với bên ngoài

- ### Phần code phía frontend sẽ gọn gàng hơn

- ### Dễ dàng theo dõi và quản lý traffic.

- ### Requests caching và cân bằng tải.

- ### Thêm một lớp bảo mật nữa cho hệ thống.

- ### Thay thế authentication services

# Nhược điểm khi sử dụng API gateway

### Tăng thời gian response

### Thêm tác nhân gây lỗi

### Có thể gây nghẽn cổ chai

### Tốn thêm tiền



