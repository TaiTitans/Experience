
---

	Trong Spring Framework, **Bean** là các đối tượng được quản lý bởi IoC (Inversion of Control) Container. Những đối tượng này là phần cốt lõi của ứng dụng và có thể là bất kỳ đối tượng nào do người dùng định nghĩa. Bean được khởi tạo, lắp ráp và quản lý bởi Spring IoC Container.

### Các thành phần chính của Bean

1. **Configuration Metadata**: Thông tin cấu hình cho các bean. Metadata này có thể được định nghĩa bằng các tệp XML, các chú thích (annotations) trong mã nguồn Java, hoặc các lớp Java Configuration.
    
2. **Scope**: Định nghĩa phạm vi của bean, tức là cách bean được khởi tạo và chia sẻ trong ứng dụng. Các phạm vi phổ biến bao gồm:
    
    - **singleton** (mặc định): Chỉ có một instance của bean được tạo và chia sẻ trong toàn bộ Spring IoC Container.
    - **prototype**: Mỗi lần yêu cầu bean sẽ tạo ra một instance mới.
    - **request**: Một instance được tạo ra cho mỗi yêu cầu HTTP (chỉ dùng trong ứng dụng web).
    - **session**: Một instance được tạo ra cho mỗi phiên làm việc HTTP (chỉ dùng trong ứng dụng web).
    - **globalSession**: Một instance được tạo ra cho mỗi phiên làm việc toàn cục (hiếm khi được dùng, chỉ dùng trong ứng dụng portlet).
3. **Lifecycle Callbacks**: Các phương thức callback cho phép bean thực hiện các hành động trong các giai đoạn cụ thể của vòng đời của nó (khởi tạo và hủy).
    
4. **Dependencies**: Các phụ thuộc của bean, tức là các đối tượng khác mà bean cần để hoạt động. Các phụ thuộc này được Spring IoC Container tự động tiêm vào bean thông qua Dependency Injection.