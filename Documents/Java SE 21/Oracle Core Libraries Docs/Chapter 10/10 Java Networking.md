
---
API **Java Networking** cung cấp các lớp và giao diện để thực hiện các chức năng mạng trong Java, bao gồm:

- **Định địa chỉ (Addressing):** Sử dụng các lớp như `InetAddress` để đại diện cho địa chỉ IP và tên máy chủ.
    
- **Sử dụng URL và URI:** Các lớp như `URL` và `URI` cho phép làm việc với các địa chỉ tài nguyên trên mạng.
    
- **Kết nối thông qua Socket:** Sử dụng các lớp như `Socket` và `ServerSocket` để thiết lập kết nối TCP giữa các máy.
    
- **Bảo mật mạng:** Các lớp trong gói `javax.net` và `javax.net.ssl` cung cấp các chức năng tạo socket bảo mật và quản lý kết nối an toàn.
    

API này bao gồm các gói và mô-đun sau:

- **`java.net`:** Chứa các lớp cho việc triển khai ứng dụng mạng, bao gồm các lớp cho địa chỉ, socket, và các giao thức mạng.
    
- **`java.net.http`:** Cung cấp API cho HTTP Client, hỗ trợ các phiên bản HTTP/1.1 và HTTP/2, cũng như giao diện cho WebSocket.
    
- **`javax.net` và `javax.net.ssl`:** Cung cấp các lớp cho việc tạo socket và socket bảo mật.
    
- **`jdk.httpserver`:** Cung cấp API cho việc xây dựng máy chủ HTTP đơn giản, hữu ích cho mục đích giáo dục và thử nghiệm.
    
- **`jdk.net`:** Cung cấp các tùy chọn socket đặc thù cho nền tảng cho các lớp socket trong `java.net` và `java.nio.channels`.
    

Ngoài ra, bạn có thể sử dụng công cụ `jwebserver` để kiểm thử và gỡ lỗi ứng dụng khách của mình.

Để cấu hình các thuộc tính hệ thống liên quan đến mạng, bạn có thể sử dụng tùy chọn `-D` của lệnh `java`, phương thức `System.setProperty(String, String)`, hoặc chỉ định chúng trong tệp `$JAVA_HOME/conf/net.properties`. Lưu ý rằng bạn chỉ có thể chỉ định các thuộc tính liên quan đến proxy trong tệp này.

Để biết thêm thông tin chi tiết, bạn có thể tham khảo tài liệu chính thức về [Java Networking](https://docs.oracle.com/en/java/javase/21/core/java-networking.html).