
---
**Dependency Injection (DI)** là một kỹ thuật trong lập trình, đặc biệt là trong lập trình hướng đối tượng, để quản lý các phụ thuộc của một lớp (class). Thay vì để lớp tự khởi tạo hoặc tìm cách để khởi tạo các phụ thuộc của nó, DI cung cấp cho lớp những phụ thuộc này từ bên ngoài. Điều này giúp làm giảm sự phụ thuộc chặt chẽ giữa các lớp, tăng tính linh hoạt và dễ dàng bảo trì.

### **Hoạt động của Dependency Injection:**

DI hoạt động dựa trên nguyên tắc **Inversion of Control (IoC)**, nơi mà việc kiểm soát khởi tạo và quản lý vòng đời của các phụ thuộc được chuyển từ lớp sử dụng sang một bên ngoài (thường là một container DI).

Các cách thức chính để thực hiện DI:

- **Constructor Injection**: Các phụ thuộc được cung cấp thông qua constructor của lớp.
- **Setter Injection**: Các phụ thuộc được cung cấp thông qua các phương thức setter.
- **Field Injection**: Các phụ thuộc được cung cấp trực tiếp vào các trường (field) của lớp bằng cách sử dụng annotation hoặc các kỹ thuật khác.

---
**Dependency Injection (DI)** trong **Spring Framework** là một phần cốt lõi của Spring, giúp quản lý các phụ thuộc của các đối tượng một cách hiệu quả và dễ dàng. Spring sử dụng **IoC Container** để thực hiện DI, nơi các đối tượng (beans) được quản lý và cung cấp các phụ thuộc cần thiết.

### 1. **Cách Dependency Injection hoạt động trong Spring:**

Spring DI hoạt động thông qua việc cấu hình các beans và các mối quan hệ giữa chúng trong container Spring, sau đó Spring sẽ tự động tiêm các phụ thuộc vào các bean khi cần thiết.

Có ba cách chính để cấu hình DI trong Spring:

- **XML Configuration**: Cấu hình beans và các phụ thuộc trong file XML.
- **Java-based Configuration**: Sử dụng annotation và Java config classes để cấu hình beans.
- **Annotation-based Configuration**: Sử dụng các annotation như `@Autowired`, `@Component`, `@Service`, `@Repository`, và `@Controller` để tự động phát hiện và tiêm các phụ thuộc.

### 2. **Các loại Dependency Injection trong Spring:**

- **Constructor Injection**: Sử dụng constructor để tiêm phụ thuộc.
- **Setter Injection**: Sử dụng setter method để tiêm phụ thuộc.
- **Field Injection**: Sử dụng annotation `@Autowired` trực tiếp trên field.

### 3. **Vì sao cần Dependency Injection trong Spring:**

- **Quản lý phụ thuộc tự động**: Spring IoC container tự động quản lý các phụ thuộc, giảm bớt công việc thủ công và tăng tính chính xác.
- **Dễ dàng kiểm thử**: DI giúp tách biệt việc khởi tạo phụ thuộc, cho phép mock hoặc inject các phụ thuộc thay thế dễ dàng trong quá trình kiểm thử.
- **Giảm sự phụ thuộc cứng**: Giúp tách biệt các module, làm cho ứng dụng linh hoạt hơn và dễ bảo trì.
- **Cấu hình tập trung**: Spring cho phép cấu hình tất cả các phụ thuộc tại một nơi, giúp dễ dàng kiểm soát và thay đổi.

Spring DI là một trong những lý do chính giúp Spring Framework trở nên phổ biến, bởi nó giúp phát triển ứng dụng một cách linh hoạt, dễ kiểm thử, và dễ dàng bảo trì.