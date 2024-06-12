
---
#### Định nghĩa

**Port** (cổng) trong mạng máy tính là một khái niệm logic dùng để xác định một kênh giao tiếp cụ thể cho truyền dữ liệu giữa các thiết bị mạng. Mỗi cổng được xác định bởi một số nguyên từ 0 đến 65535, giúp phân biệt các dịch vụ mạng khác nhau chạy trên cùng một địa chỉ IP.

#### Chức năng

Port giúp hệ điều hành và các thiết bị mạng biết được ứng dụng hoặc dịch vụ nào sẽ nhận dữ liệu đến từ một kết nối mạng. Mỗi port tương ứng với một dịch vụ hoặc ứng dụng cụ thể, cho phép nhiều dịch vụ khác nhau có thể hoạt động đồng thời trên cùng một địa chỉ IP.

#### Phân loại

1. **Well-Known Ports (Cổng đã biết)**: Từ 0 đến 1023. Được gán cho các dịch vụ và giao thức thông dụng như HTTP (port 80), HTTPS (port 443), FTP (port 21), SSH (port 22), và SMTP (port 25).
2. **Registered Ports (Cổng đã đăng ký)**: Từ 1024 đến 49151. Được IANA (Internet Assigned Numbers Authority) đăng ký cho các dịch vụ và ứng dụng cụ thể.
3. **Dynamic or Private Ports (Cổng động hoặc riêng tư)**: Từ 49152 đến 65535. Thường được sử dụng tạm thời bởi các ứng dụng để giao tiếp.

#### Cách sử dụng

Mỗi kết nối mạng được xác định bởi một cặp địa chỉ IP và port. Ví dụ, một kết nối TCP/IP giữa hai máy tính sẽ có một địa chỉ IP nguồn và port nguồn, cùng với một địa chỉ IP đích và port đích. Điều này giúp định tuyến dữ liệu chính xác đến ứng dụng hoặc dịch vụ cần thiết.

#### Một số cổng phổ biến và ứng dụng của chúng

1. **HTTP (HyperText Transfer Protocol)**: Port 80. Dùng cho truy cập web không bảo mật.
2. **HTTPS (HyperText Transfer Protocol Secure)**: Port 443. Dùng cho truy cập web bảo mật.
3. **FTP (File Transfer Protocol)**: Port 21. Dùng cho truyền tệp.
4. **SMTP (Simple Mail Transfer Protocol)**: Port 25. Dùng cho gửi email.
5. **DNS (Domain Name System)**: Port 53. Dùng cho dịch vụ phân giải tên miền.
6. **SSH (Secure Shell)**: Port 22. Dùng cho truy cập từ xa bảo mật.