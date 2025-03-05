
---
**Clean Architecture** là một kiến trúc ứng dụng rất nổi tiếng dựa trên nguyên lý loại bỏ sự lệ thuộc giữa các đối tượng cũng như các layer trong ứng dụng. Nguyên lý này kế thừa và phát triển dựa trên **Dependency Inversion** - nguyên lý nổi tiếng trong **SOLID**.
Trong kiến trúc **Clean Architecture** bao gồm 4 layer được đại diện thông qua các vòng tròn đồng tâm. Các vòng tròn ở trong sẽ không hề biết gì về các vòng tròn bên ngoài. Nguyên tắc "hướng tâm" này được minh hoạ như sau:

![](https://statics.cdn.200lab.io/2022/05/clean-architecture.png)

---
Từ trong ra ngoài **Clean Architecture** sẽ bao gồm: **Entities**, **Use Cases**, **Interface Adapters** và **Frameworks & Drivers**.

_**Entities**_: là khái niệm dùng để mô tả các Business Logic. Đây là layer quan trọng nhất, là nơi bạn thực hiện giải quyết các vấn đề - mục đích khi xây dựng app. Layer này không chứa bất kỳ một framework nào, bạn có thể chạy nó mà không cần emulator. Nó giúp bạn dễ dàng test, maintain và develop phần business logic.

_**Use case**_ : chứa các rule liên quan trực tiếp tới ứng dụng cục bộ (application-specific business rules).

_**Interface Adapter**_ : tập hợp các adapter phục vụ quá trình tương tác với các công nghệ.

_**Framework and Drivers**_ : chứa các công cụ về cơ sở dữ liệu và các framework, thông thường bạn sẽ không phải lập trình nhiều ở tầng này. Tuy nhiên cần chắc chắn về mức ưu tiên sử dụng các công cụ này trong project.

---
## Spring Boot - Clean Architecture

### **1. Tầng Core (Enterprise Business Rules):**

Đây là tầng chứa logic nghiệp vụ cốt lõi của ứng dụng. Tầng này độc lập với các chi tiết triển khai cụ thể như cơ sở dữ liệu, giao diện người dùng, hay các dịch vụ bên ngoài.

- **Entities**: Các lớp biểu diễn các đối tượng nghiệp vụ, chứa các logic cơ bản, như các thuộc tính và phương thức.
- **Use Cases (Application Business Rules)**: Đây là nơi các quy tắc nghiệp vụ cấp ứng dụng được định nghĩa. Các lớp trong tầng này sẽ tương tác với Entities để thực hiện các nghiệp vụ cụ thể.

### **2. Tầng Interface Adapters (Interface Adapters):**

Tầng này đóng vai trò là bộ chuyển đổi giữa tầng Core và các chi tiết cụ thể của hệ thống như cơ sở dữ liệu, giao diện người dùng, hoặc các API bên ngoài.

- **Controllers**: Xử lý các yêu cầu HTTP, chuyển đổi dữ liệu đầu vào thành các đối tượng nghiệp vụ và gọi các Use Cases để thực hiện nghiệp vụ.
- **Presenters/View Models**: Chuyển đổi dữ liệu từ tầng Core để đưa ra kết quả cho người dùng.
- **Gateways/Repositories**: Triển khai các giao diện để tương tác với cơ sở dữ liệu hoặc các dịch vụ bên ngoài.

### **3. Tầng Frameworks and Drivers (Frameworks and Drivers):**

Tầng này chứa các chi tiết cụ thể về framework và công nghệ được sử dụng. Đây là nơi bạn triển khai các chi tiết như cấu hình Spring Boot, truy cập cơ sở dữ liệu, và kết nối mạng.

- **Configuration**: Cấu hình Spring Boot, cấu hình kết nối cơ sở dữ liệu, bảo mật, và các cấu hình khác.
- **Adapters**: Các bộ điều hợp để kết nối các chi tiết bên ngoài như cơ sở dữ liệu, các API bên ngoài, hoặc các dịch vụ khác.

### **4. Tầng External (Infrastructure):**

Bao gồm cơ sở dữ liệu, message brokers, hoặc bất kỳ hệ thống bên ngoài nào mà ứng dụng sẽ tương tác.


---

**Ví dụ cấu trúc thư mục của một dự án Spring Boot sử dụng Clean Architecture:**

```
src
└── main
    ├── java
    │   └── com.example.cleanarchitecture
    │       ├── core
    │       │   ├── entities
    │       │   │   ├── User.java
    │       │   │   └── Product.java
    │       │   └── usecase
    │       │       ├── CreateUserUseCase.java
    │       │       ├── GetUserUseCase.java
    │       │       └── CreateOrderUseCase.java
    │       ├── adapters
    │       │   ├── controllers
    │       │   │   ├── UserController.java
    │       │   │   └── ProductController.java
    │       │   ├── presenters/DTOs
    │       │   │   └── UserPresenter.java
    │       │   └── gateways
    │       │       ├── UserRepository.java
    │       │       └── ProductRepository.java
    │       ├── infrastructure
    │       │   ├── database
    │       │   │   ├── JpaUserRepository.java
    │       │   │   └── JpaProductRepository.java
    │       │   ├── api
    │       │   │   ├── ExternalPaymentService.java
    │       │   │   └── ExternalShippingService.java
    │       │   └── configuration
    │       │       └── AppConfig.java
    │       └── application
    │           └── CleanArchitectureApplication.java
    └── resources
        └── application.yml

```

