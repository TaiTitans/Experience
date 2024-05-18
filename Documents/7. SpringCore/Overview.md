
---
**Spring Framework** là một khung phát triển mã nguồn mở phổ biến dành cho Java, giúp bạn dễ dàng xây dựng các ứng dụng doanh nghiệp. Nó cung cấp mọi thứ cần thiết để tận dụng sức mạnh của ngôn ngữ Java trong môi trường doanh nghiệp, đồng thời hỗ trợ cả Groovy và Kotlin như những lựa chọn thay thế trên nền tảng JVM (Java Virtual Machine).

**Ưu điểm của Spring Framework:**

- **Linh hoạt:** Spring cho phép bạn tạo nhiều loại kiến trúc khác nhau tùy theo nhu cầu của ứng dụng.
- **Hỗ trợ đa dạng:** Spring hỗ trợ nhiều kịch bản ứng dụng khác nhau, từ các ứng dụng doanh nghiệp lớn chạy trên các máy chủ cũ đến các ứng dụng đóng gói gọn nhẹ chạy trên nền tảng đám mây.
- **Cộng đồng lớn:** Là mã nguồn mở, Spring có cộng đồng người dùng và nhà phát triển sôi động, cung cấp phản hồi liên tục dựa trên các tình huống thực tế đa dạng. Điều này giúp Spring không ngừng cải tiến trong một thời gian dài.

**Những điểm cần lưu ý về Spring:**

- **Yêu cầu Java 17+:** Để sử dụng Spring Framework phiên bản 6.0 trở lên, bạn cần sử dụng Java 17 trở lên.
- **Các loại hình ứng dụng:** Spring có thể dùng để xây dựng:
    - Các ứng dụng doanh nghiệp lớn, chạy trên các máy chủ cũ có chu kỳ nâng cấp hạn chế.
    - Các ứng dụng đóng gói thành một file JAR đơn lẻ, tích hợp sẵn server, thường chạy trên nền tảng đám mây.
    - Các ứng dụng độc lập (chẳng hạn như xử lý hàng loạt hoặc tích hợp dữ liệu) không cần server.
---
## "Spring" trong Lập trình Java

**"Spring"** là một thuật ngữ đa nghĩa trong lập trình Java, có thể ám chỉ:

**1. Spring Framework:**

- Đây là **nền tảng cốt lõi** cung cấp các tính năng thiết yếu cho phát triển ứng dụng doanh nghiệp Java.
- Spring Framework được **chia thành các module** riêng biệt, cho phép ứng dụng lựa chọn sử dụng những module phù hợp với nhu cầu.
- **Các module cốt lõi** bao gồm: mô hình cấu hình và cơ chế tiêm phụ thuộc (dependency injection).
- Ngoài ra, Spring Framework cung cấp **nền tảng hỗ trợ** cho nhiều kiến trúc ứng dụng khác nhau, bao gồm:
    - **Nhắn tin:** Giao tiếp giữa các thành phần ứng dụng.
    - **Quản lý giao dịch dữ liệu:** Đảm bảo tính toàn vẹn dữ liệu trong các thao tác.
    - **Persistence:** Lưu trữ dữ liệu bền vững.
    - **Web:** Phát triển ứng dụng web.
- Spring Framework bao gồm **hai framework web phổ biến:**
    - **Spring MVC:** Framework web dựa trên Servlet.
    - **Spring WebFlux:** Framework web phản ứng (reactive).

**2. Family of Spring:**

- Theo thời gian, dựa trên Spring Framework, nhiều dự án khác được phát triển, tạo thành **"family of Spring"**.
- Khi nói về "Spring" một cách tổng thể, người ta thường đề cập đến **toàn bộ các dự án này**.

**Lưu ý về module:**

- Các JAR (file thư viện) của Spring Framework có thể được triển khai trên **đường dẫn module của JDK 9 ("Jigsaw")**.
- Để sử dụng trong các ứng dụng hỗ trợ Jigsaw, JAR của Spring Framework 5 đi kèm với các mục nhập manifest **"Automatic-Module-Name"** giúp xác định tên module ổn định ở cấp độ ngôn ngữ.

---

Triết lý thiết kế của Spring Framework tập trung vào tính linh hoạt, khả năng tương thích ngược, thiết kế API tốt và chất lượng code cao. Điều này giúp các nhà phát triển xây dựng các ứng dụng Java doanh nghiệp mạnh mẽ, dễ bảo trì và mở rộng.

---
### Các Công nghệ Cốt lõi của Spring Framework:

Ba nội dung chính bao gồm:

1. **IoC Container (Kiểm soát đảo ngược):** Đây là thành phần quan trọng nhất của Spring Framework. Nó chịu trách nhiệm quản lý vòng đời của các object (đối tượng), tiêm các phụ thuộc (dependency injection) và cung cấp cấu hình cho ứng dụng.
    
2. **Lập trình hướng khía cạnh (AOP):** Spring cung cấp framework AOP riêng, dễ hiểu về mặt lý thuyết và đáp ứng được 80% yêu cầu lập trình AOP thông thường trong lập trình doanh nghiệp Java.
    
3. **Tích hợp với AspectJ:** Bên cạnh framework AOP của riêng mình, Spring cũng hỗ trợ tích hợp với AspectJ - một framework AOP phổ biến và trưởng thành trong lĩnh vực doanh nghiệp Java.
    

**Ngoài ra, tài liệu còn đề cập đến:**

- **AOT processing (Xử lý trước thời gian):** Đây là kỹ thuật tối ưu hóa ứng dụng trước khi chạy, thường được sử dụng để triển khai ảnh gốc (native image) bằng GraalVM.