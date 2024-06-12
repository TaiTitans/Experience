
---
MVC framework được thiết kế xoay quanh DispatcherServlet - cho phép xử lý tất cả các HTTP request và response. Sơ đồ dưới đây giải thích flow xử lý request của Spring Web MVC DispatcherServlet.


![](https://images.viblo.asia/9fdcb667-825a-42fb-9c78-187ae1a937f4.png)


Đây là chuỗi các sự kiện tương ứng với một yêu cầu HTTP đến DispatcherServlet:

1. Sau khi nhận được một yêu cầu HTTP, DispatcherServlet chỉ định cho HandlerMapping gọi Controller thích hợp.
    
2. Controller sẽ nhận yêu cầu và gọi các Service tương ứng thích hợp dựa trên GET được sử dụng hoặc phương thức POST. Các phương thức service này sẽ thiết lập một nhóm các dữ liệu Model được định nghĩa theo logic business và trả về tên View cho DispatcherServlet.
    
3. Các DispatcherServlet sẽ được các ViewResolver hỗ trợ để chọn được View đã định nghĩa tương ứng với Request.
    
4. Khi View được hoàn thiện, Các DispatcherServlet sẽ chuyển dữ liệu Model tới View và render trên trình duyệt.
    

Tất cả các thành phần nêu trên, ví dụ như. HandlerMapping, Controller và ViewResolver là bộ phận của WebApplicationContext - một mở rộng của ApplicationContext với một số tính năng bổ sung cần thiết cho các ứng dụng web.


---

### 1. Kiến trúc của Spring MVC

- **Model**: Chứa dữ liệu của ứng dụng và logic nghiệp vụ. Model thường bao gồm các POJO (Plain Old Java Objects) và có thể tương tác với cơ sở dữ liệu thông qua DAO hoặc các dịch vụ.
- **View**: Chịu trách nhiệm hiển thị dữ liệu cho người dùng. View có thể là JSP, Thymeleaf, hoặc các công nghệ khác.
- **Controller**: Xử lý các yêu cầu từ người dùng, thao tác dữ liệu thông qua Model, và trả về View để hiển thị kết quả.

### Cấu trúc một ứng dụng Spring MVC

Ứng dụng Spring MVC thường có cấu trúc thư mục như sau:

```bash
src/main/java
    com/example/controller
    com/example/model
    com/example/service
    com/example/repository
src/main/resources
    application.properties
src/main/webapp
    WEB-INF
        views
            home.jsp
pom.xml

```


### 3. Cài đặt và cấu hình Spring MVC


#### Bước 1: Thiết lập dự án Maven
Tạo tệp `pom.xml` và thêm các dependency cần thiết cho Spring MVC.

#### Bước 2: Cấu hình DispatcherServlet
Tạo tệp `web.xml` trong thư mục `src/main/webapp/WEB-INF` để cấu hình DispatcherServlet.

#### Bước 3: Cấu hình Spring MVC

Tạo tệp `dispatcher-servlet.xml` trong thư mục `src/main/webapp/WEB-INF` để cấu hình Spring MVC.

-- Sau đó bắt đầu viết Controller và các thành phần khác. Và chạy server bằng Apache Tomcat.