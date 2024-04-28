# Spring vs. Spring Boot vs. Spring MVC

### Spring Framework

Spring Framework là một framework lập trình Java cho việc phát triển các ứng dụng doanh nghiệp. Nó cung cấp một cách tiếp cận linh hoạt và mạnh mẽ để xây dựng các ứng dụng Java. Spring Framework bao gồm nhiều tính năng như Dependency Injection (DI), Aspect-Oriented Programming (AOP), và một số module như Spring JDBC, Spring MVC, và Spring Security.

### Spring MVC

Spring MVC là một phần của Spring Framework và được sử dụng để xây dựng các ứng dụng web. Nó cung cấp một cấu trúc để phát triển các ứng dụng web theo mô hình MVC (Model-View-Controller). Trong Spring MVC, Controller nhận các yêu cầu từ người dùng, sau đó xử lý chúng và gửi dữ liệu đến View để hiển thị. Spring MVC hỗ trợ nhiều cấu hình và tính năng như các validators, Interceptors, và RESTful web services.

### Spring Boot

Spring Boot là một dự án con của Spring Framework, tập trung vào việc giảm độ phức tạp trong việc phát triển các ứng dụng Spring. Nó cung cấp các cấu hình mặc định và tự động cấu hình một số tính năng, giúp người phát triển tập trung vào việc viết mã chức năng chính của ứng dụng mà không cần quan tâm nhiều đến việc cấu hình. Spring Boot cũng hỗ trợ các tính năng như embedded server (tomcat, jetty), auto-configuration, và các dependency management.

### So sánh

- **Spring Framework**: Là core của mọi thứ, cung cấp các tính năng cơ bản như DI, AOP.
- **Spring MVC**: Dành cho việc xây dựng ứng dụng web theo mô hình MVC.
- **Spring Boot**: Giúp giảm bớt công việc cấu hình và tập trung vào việc viết mã chức năng chính của ứng dụng.

Khi lựa chọn giữa chúng:

- Sử dụng **Spring Framework** khi cần một phạm vi lớn hơn các tính năng của Spring và không cần mức độ tự động hóa cao.
- Sử dụng **Spring MVC** khi xây dựng các ứng dụng web truyền thống theo mô hình MVC và muốn sử dụng các tính năng cụ thể của Spring trong việc phát triển web.
- Sử dụng **Spring Boot** khi muốn nhanh chóng bắt đầu một dự án mới và muốn giảm bớt công việc cấu hình ban đầu.

Việc lựa chọn giữa chúng phụ thuộc vào nhu cầu cụ thể của dự án và mức độ tự động hóa.