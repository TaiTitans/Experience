
---
	JavaNetwork đề cập đến việc viết các chương trình thực thi trên nhiều thiết bị (máy tính), trong đó tất cả các thiết bị đều được kết nối với nhau bằng mạng.

Ưu điểm:
- Creating server-client applications
- Implementing networking protocols
- Implement socket programming
- Creating web services

## Package Used in Networking

Gói java.net của API J2SE chứa một tập hợp các lớp và giao diện cung cấp chi tiết giao tiếp cấp thấp.

Gói java.net cung cấp hỗ trợ cho hai giao thức mạng phổ biến -

- TCP − TCP là viết tắt của Giao thức điều khiển truyền dẫn, cho phép liên lạc đáng tin cậy giữa hai ứng dụng. TCP thường được sử dụng qua Giao thức Internet, được gọi là TCP/IP.

- UDP − UDP là viết tắt của Giao thức gói dữ liệu người dùng, một giao thức không có kết nối cho phép các gói dữ liệu được truyền giữa các ứng dụng.

## Socket Programming in Java Networking

Sockets cung cấp cơ chế giao tiếp giữa hai máy tính sử dụng TCP. Một chương trình máy khách tạo một socket ở đầu giao tiếp của nó và cố gắng kết nối ổ cắm đó với máy chủ.

	Khi kết nối được thực hiện, máy chủ sẽ tạo một đối tượng socket ở đầu giao tiếp của nó. Máy khách và máy chủ hiện có thể giao tiếp bằng cách ghi và đọc từ socket.

Lớp java.net.Socket đại diện cho một socket và lớp java.net.ServerSocket cung cấp một cơ chế để chương trình máy chủ lắng nghe các máy khách và thiết lập kết nối với chúng.

---
Các bước sau đây xảy ra khi thiết lập kết nối TCP giữa hai máy tính bằng cách sử dụng socket ->
- Máy chủ khởi tạo một đối tượng ServerSocket, biểu thị giao tiếp số cổng nào sẽ diễn ra trên đó.
- Máy chủ gọi phương thức Accept() của lớp ServerSocket. Phương thức này đợi cho đến khi máy khách kết nối với máy chủ trên cổng đã cho.
- Sau khi máy chủ đang chờ, máy khách sẽ khởi tạo một đối tượng Socket, chỉ định tên máy chủ và số cổng để kết nối.
- Hàm tạo của lớp Socket cố gắng kết nối máy khách với máy chủ được chỉ định và số cổng. Nếu giao tiếp được thiết lập, máy khách hiện có đối tượng Socket có khả năng giao tiếp với máy chủ.
- Về phía máy chủ, phương thức Accept() trả về một tham chiếu đến một socket mới trên máy chủ được kết nối với socket của máy khách.

=> Sau khi kết nối được thiết lập, giao tiếp có thể diễn ra bằng cách sử dụng luồng I/O. Mỗi socket có cả OutputStream và InputStream. OutputStream của máy khách được kết nối với InputStream của máy chủ và InputStream của máy khách được kết nối với OutputStream của máy chủ.

TCP là giao thức truyền thông hai chiều, do đó dữ liệu có thể được gửi qua cả hai luồng cùng một lúc.

