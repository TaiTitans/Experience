
---
Trong Spring Boot, một **Embedded Web Server** là một web server được tích hợp sẵn trong ứng dụng, cho phép bạn chạy ứng dụng web mà không cần cài đặt và cấu hình một web server riêng biệt. Spring Boot hỗ trợ các web server nhúng phổ biến như Tomcat, Jetty và Undertow.


### Lợi ích của Embedded Web Server trong Spring Boot

1. **Đơn giản hóa quá trình triển khai**:
    
    - Bạn chỉ cần một tệp JAR hoặc WAR duy nhất để triển khai ứng dụng của mình. Tất cả các cấu hình cần thiết cho web server đã được tích hợp sẵn.
2. **Tự động cấu hình**:
    
    - Spring Boot tự động cấu hình các cài đặt cần thiết cho web server, giúp giảm bớt công việc cho các nhà phát triển.
3. **Tính di động**:
    
    - Ứng dụng của bạn có thể chạy trên bất kỳ môi trường nào mà không cần lo lắng về sự khác biệt giữa các web server hoặc các cấu hình cụ thể của chúng.
4. **Tích hợp và thử nghiệm dễ dàng**:
    
    - Việc tích hợp và thử nghiệm trở nên dễ dàng hơn vì bạn không cần phải thiết lập một web server bên ngoài mỗi khi muốn kiểm tra ứng dụng.

---
### Cách hoạt động của Embedded Web Server trong Spring Boot

Khi bạn tạo một ứng dụng Spring Boot, mặc định nó sẽ bao gồm một web server nhúng. Dưới đây là cách Spring Boot xử lý việc này:

1. **Khởi tạo Web Server**:
    
    - Khi ứng dụng Spring Boot khởi động, nó sẽ tự động khởi tạo một instance của web server nhúng.
2. **Cấu hình Web Server**:
    
    - Spring Boot sẽ áp dụng các cấu hình mặc định và bất kỳ cấu hình tùy chỉnh nào mà bạn đã xác định trong `application.properties` hoặc `application.yml`.
3. **Triển khai ứng dụng**:
    
    - Ứng dụng của bạn sẽ được triển khai trên web server nhúng, và bạn có thể truy cập nó thông qua một URL (mặc định là `http://localhost:8080`).

---
