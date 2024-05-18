# Spring Initializr

> Spring Initializr là một công cụ dựa trên **web-based tool** được cung cấp bởi Pivotal Web Service.Với sự giúp đỡ của **Spring Initializr**,chúng ta có thể dễ dàng tạo ra cấu trúc của **Spring Boot Project**. Nó cung cấp API có thể mở rộng để tạo các dự án dựa trên JVM.

Nó cũng cung cấp nhiều tùy chọn khác nhau cho dự án được thể hiện dưới dạng mô hình siêu dữ liệu.

Mô hình siêu dữ liệu cho phép chúng tôi định cấu hình danh sách các phụ thuộc được hỗ trợ bởi các phiên bản nền tảng và JVM, v.v. Nó phục vụ siêu dữ liệu của nó ở dạng nổi tiếng cung cấp hỗ trợ cần thiết cho khách hàng bên thứ ba.

## Spring Initializr Modules

- **initializr-actuator:** Là module tuỳ chọn cung cấp thông tin và dữ liệu thông kê về việc tạo một project.

- **initializr-bom:** Trong mô-đun này, BOM là viết tắt của Bill Of Materials. Trong Spring Boot, BOM là một loại POM đặc biệt được sử dụng để kiểm soát các phiên bản phụ thuộc của dự án. Nó cung cấp một vị trí trung tâm để xác định và cập nhật các phiên bản đó. Nó cung cấp sự linh hoạt để thêm phần phụ thuộc vào mô-đun của chúng tôi mà không phải lo lắng về các phiên bản.

  Bên ngoài thế giới phần mềm, BOM là danh sách các bộ phận, hạng mục, cụm lắp ráp và các vật liệu khác cần thiết để tạo ra sản phẩm. Nó giải thích những gì, làm thế nào và ở đâu để thu thập các tài liệu cần thiết

- **initializr-docs**: Nó cung cấp tài liệu.

- **initializr-generator:** Nó là một thư viện tạo dự án cốt lõi.

- **initializr-generator-spring: ** là một công cụ được sử dụng để tạo ra các dự án Spring Boot khởi đầu nhanh chóng và dễ dàng. Nó cung cấp một giao diện người dùng web cho phép bạn tùy chỉnh và cấu hình các yếu tố khác nhau của dự án Spring Boot, chẳng hạn như ngôn ngữ lập trình, phiên bản Spring Boot, các phụ thuộc và các module cần thiết.

- **initializr-generator-test:** Nó cung cấp cơ sở hạ tầng thử nghiệm để tạo dự án.

- **initializr-metadata:**  Nó cung cấp cơ sở hạ tầng siêu dữ liệu cho các khía cạnh khác nhau của dự án.

- **initializr-service-example:** Nó cung cấp các trường hợp tùy chỉnh.

- **initializr-version-resolver:** Đây là mô-đun tùy chọn để trích xuất số phiên bản từ POM tùy ý.

- **initializr-web**: Nó cung cấp điểm cuối web cho khách hàng bên thứ ba.