# **Layered Architecture** (Kiến trúc phân lớp)

Mô hình này chia ứng dụng thành các lớp độc lập như:

> **Controller Layer:** Lớp này xử lý các yêu cầu HTTP, chuyển tiếp dữ liệu giữa client và Service Layer.
> **Service Layer:** Lớp này chứa các logic nghiệp vụ chính, giao tiếp với Repository Layer.
> **Repository Layer:** Lớp này quản lý truy cập và thao tác dữ liệu với cơ sở dữ liệu.

**Cấu trúc project base:**

Tạo một dự án Spring Boot mới bằng cách sử dụng Spring Initializr hoặc IDE của bạn.
Tổ chức cấu trúc thư mục dựa trên mô hình Layered Architecture như sau:
**com.example.myapp**

- `controller`
- `service`
- `repository`
- `model`
- `config`
- `exception`
- `util`

**Thiết lập các dependency:**

Thêm các dependency cần thiết như spring-boot-starter-web, spring-boot-starter-data-jpa, mysql-connector-java, lombok, v.v.
**Cấu hình ứng dụng:**

Tạo các class cấu hình như application.properties hoặc application.yml để thiết lập các cài đặt như kết nối cơ sở dữ liệu, logging, etc.
**Viết các lớp controller, service, repository:**

Tạo các lớp controller để xử lý các yêu cầu HTTP RESTful.
Tạo các lớp service để triển khai logic nghiệp vụ.
Tạo các lớp repository để thao tác với cơ sở dữ liệu.
**Thêm các tính năng bổ sung:**

Xử lý exception, validation, authorization/authentication, caching, logging, etc.