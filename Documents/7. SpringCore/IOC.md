
---
	**IoC (Inversion of Control) Container** trong Spring Framework là một phần cốt lõi của Spring, giúp quản lý các đối tượng của ứng dụng, bao gồm việc khởi tạo, cấu hình và quản lý vòng đời của chúng. Spring IoC Container sử dụng nguyên tắc Dependency Injection (DI) để quản lý các đối tượng và phụ thuộc giữa chúng.
## IoC là gì?

IoC(**I**nversion **o**f **C**ontrol): Đảo ngược điều khiển, nó giúp làm thay đổi luồng điều khiển của chương trình một cách linh hoạt.

Thường dùng với **Denpendency Injection**.
## Spring IoC

IoC Container là thành phần thực hiện IoC.

Trong Spring, Spring Container (IoC Container) sẽ tạo các đối tượng, lắp rắp chúng lại với nhau, cấu hình các đối tượng và quản lý vòng đời của chúng từ lúc tạo ra cho đến lúc bị hủy.

Spring container sử dụng DI để quản lý các thành phần, đối tượng để tạo nên 1 ứng dụng. Các thành phần, đối tượng này gọi là Spring Bean.

Để tạo đối tượng, cấu hình, lắp rắp chúng, Spring Container sẽ đọc thông tin từ các file xml và thực thi chúng.

![](https://stackjava.com/wp-content/uploads/2017/12/spring-ioc-container.jpg)
### Các loại IoC Container chính

Spring Framework cung cấp hai loại container chính:

1. **BeanFactory**:
    
    - Đây là container cơ bản và nhẹ nhất, cung cấp các tính năng cấu hình và cơ chế cơ bản để quản lý các đối tượng.
    - Thích hợp cho các ứng dụng với tài nguyên hạn chế hoặc yêu cầu thời gian khởi động nhanh.
    - Các lớp triển khai phổ biến của BeanFactory là `XmlBeanFactory` (đã lỗi thời) và `DefaultListableBeanFactory`.
2. **ApplicationContext**:
    
    - Đây là một phần mở rộng của BeanFactory, cung cấp nhiều tính năng bổ sung cho các ứng dụng doanh nghiệp.
    - ApplicationContext cung cấp các tính năng như tích hợp với AOP (Aspect Oriented Programming), xử lý tài nguyên thông báo (cho quốc tế hóa), xuất bản sự kiện, và các ngữ cảnh cụ thể như WebApplicationContext.
    - Các lớp triển khai phổ biến của ApplicationContext là `ClassPathXmlApplicationContext`, `FileSystemXmlApplicationContext`, và `AnnotationConfigApplicationContext`.

### Các tính năng chính của ApplicationContext

- **Tích hợp AOP**: ApplicationContext dễ dàng tích hợp với các tính năng AOP của Spring, giúp áp dụng các khía cạnh (aspects) như logging, transaction management một cách dễ dàng.
- **Quản lý thông báo và quốc tế hóa (I18N)**: Hỗ trợ xử lý tài nguyên thông báo, giúp ứng dụng dễ dàng hỗ trợ nhiều ngôn ngữ.
- **Xuất bản sự kiện**: ApplicationContext có cơ chế xuất bản và lắng nghe sự kiện, cho phép các đối tượng trong ứng dụng giao tiếp với nhau thông qua các sự kiện tùy chỉnh.
- **Ngữ cảnh cụ thể cho ứng dụng web**: `WebApplicationContext` là một ngữ cảnh con của ApplicationContext, được sử dụng trong các ứng dụng web để tích hợp chặt chẽ với Servlet container.

### Cách làm việc của IoC Container

IoC Container làm việc thông qua việc cấu hình metadata (dữ liệu cấu hình), mà có thể được định nghĩa bằng:

- **Tệp XML**: Bạn có thể định nghĩa các bean và các phụ thuộc của chúng trong một tệp cấu hình XML.
- **Annotations**: Sử dụng các chú thích (annotations) trong mã nguồn Java để định nghĩa các bean và các phụ thuộc.
- **Java Configuration**: Sử dụng các lớp Java với chú thích `@Configuration` và các phương thức được đánh dấu với `@Bean` để cấu hình các bean.

Spring Container (hay ApplicationContext) là thành phần trung tâm của Spring Framework. Nó chịu trách nhiệm quản lý vòng đời của các bean (các đối tượng được Spring quản lý), bao gồm:

- **Khởi tạo** bean.
- **Cấu hình** các phụ thuộc của bean.
- **Quản lý** vòng đời của bean (bao gồm cả việc phá hủy).

### **Bean Definition**

Các bean và sự phụ thuộc của chúng được định nghĩa trong các file cấu hình (XML, Java Config, hoặc Annotation). Spring Container sử dụng các định nghĩa này để biết cách khởi tạo và quản lý các bean.

### **Bean Lifecycle**

Spring Container quản lý vòng đời của các bean theo các giai đoạn sau:

- **Instantiation**: Tạo một instance của bean.
- **Populate Properties**: Tiêm các phụ thuộc vào bean.
- **Post Initialization**: Gọi các callback như `@PostConstruct` hoặc `InitializingBean` interface nếu có.
- **Usage**: Bean được sử dụng trong ứng dụng.
- **Destruction**: Khi ứng dụng kết thúc hoặc bean không còn cần thiết, Spring gọi các phương thức như `@PreDestroy` hoặc `DisposableBean` interface nếu có.

### **Bean Scopes**

Spring cung cấp nhiều loại scope khác nhau cho các bean để xác định cách chúng được tạo ra và sử dụng:

- **Singleton** (mặc định): Chỉ một instance duy nhất cho mỗi bean trong Spring Container.
- **Prototype**: Một instance mới được tạo mỗi lần bean được yêu cầu.
- **Request**: Một instance duy nhất cho mỗi HTTP request (chỉ dùng trong ứng dụng web).
- **Session**: Một instance duy nhất cho mỗi HTTP session (chỉ dùng trong ứng dụng web).