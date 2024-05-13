
---

Dưới đây là một số thuật ngữ và khái niệm quan trọng liên quan đến Spring Core, được giải thích một cách dễ hiểu và đầy đủ bằng tiếng Việt:

1. **Beans** (Đối Tượng Bean):
    
    - Trong Spring, một "bean" là một đối tượng Java được quản lý bởi container của Spring. Beans thường được định nghĩa bằng cách sử dụng siêu dữ liệu cấu hình, có thể được chỉ định trong XML, Java annotations, hoặc mã Java.
2. **Inversion of Control (IoC)** (Nguyên Tắc Đảo Ngược Kiểm Soát):
    
    - Một trong những nguyên tắc chính của Spring là Inversion of Control (IoC), có nghĩa là container của Spring đảm nhận việc quản lý vòng đời của các beans và tiêm các phụ thuộc vào chúng.
3. **Dependency Injection (DI)** (Tiêm Phụ Thuộc):
    
    - Spring sử dụng Dependency Injection (DI) để quản lý các phụ thuộc giữa các beans. Trong DI, các phụ thuộc của một đối tượng được cung cấp bởi container, thay vì đối tượng tạo hoặc tra cứu phụ thuộc của mình.
4. **Container** (Container):
    
    - Container của Spring là phần cốt lõi của Framework Spring, tạo và quản lý các beans và phụ thuộc của chúng.
5. **ApplicationContext** (ApplicationContext):
    
    - ApplicationContext là một triển khai của container Spring. Nó chịu trách nhiệm tải và quản lý siêu dữ liệu cấu hình và tạo các beans được định nghĩa trong siêu dữ liệu đó.
6. **Aspect-Oriented Programming (AOP)** (Lập Trình Hướng Mặt Định):
    
    - Spring hỗ trợ Aspect-Oriented Programming (AOP), cho phép bạn tách các quan tâm chéo, như ghi nhật ký hoặc bảo mật, ra khỏi logic kinh doanh của ứng dụng.
7. **Events** (Sự Kiện):
    
    - Spring cung cấp một mô hình sự kiện cho phép các beans gửi và nhận sự kiện. Điều này được sử dụng để tách các beans khỏi nhau, làm cho ứng dụng trở nên lỏng lẻo hơn.
8. **ApplicationEvent và listener** (Sự kiện Ứng Dụng và Người Nghe):
    
    - Spring hỗ trợ mô hình xuất bản và đăng ký để xử lý sự kiện, ApplicationEvent định nghĩa đối tượng sự kiện và người nghe là một lớp thực hiện giao diện ApplicationListener, lắng nghe sự kiện cụ thể và thực hiện hành động cần thiết.
9. **Data Access** (Truy Cập Dữ Liệu):
    
    - Spring cung cấp một sự trừu tượng cao và nhất quán cho việc truy cập dữ liệu sử dụng các framework khác nhau như JDBC, Hibernate, JPA.
10. **Transactions** (Giao Dịch):
    
    - Spring cung cấp một cách linh hoạt, nhất quán và dễ dàng để quản lý giao dịch một cách khai báo với các công nghệ cơ sở như JPA, JDBC và Hibernate.
11. **Task Execution and Scheduling** (Thực Hiện và Lập Lịch Công Việc):
    
    - Spring cung cấp TaskExecutor và TaskScheduler, cung cấp một cách thuận tiện để chạy các nhiệm vụ song song, theo lịch hoặc bất đồng bộ.