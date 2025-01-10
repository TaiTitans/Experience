
---

`@Service` là một annotation trong Spring Framework được sử dụng để đánh dấu một lớp Java như là một service trong tầng service của ứng dụng. Annotation này giúp Spring quản lý và phát hiện lớp như một bean để sử dụng trong ứng dụng.

###  **Mục đích của `@Service`:**

- **Định nghĩa tầng Service**: Tầng này chứa logic nghiệp vụ (business logic) và thường là nơi thực hiện các quy tắc nghiệp vụ của ứng dụng.
- **Quản lý bean**: Spring tự động phát hiện và quản lý các lớp được đánh dấu với `@Service` như một bean trong Application Context.

### **Lợi ích của `@Service`:**

- **Rõ ràng trong phân lớp kiến trúc**: Phân biệt rõ tầng service với các tầng khác như controller hay repository.
- **Quản lý bean tự động**: Không cần cấu hình thủ công, Spring tự động quản lý lifecycle của bean.
- **Dễ dàng kiểm thử**: Service lớp dễ dàng kiểm thử độc lập, do logic nghiệp vụ thường được tách biệt khỏi các thành phần khác.

### **Khi nào sử dụng `@Service`:**

- Khi bạn viết các lớp chứa logic nghiệp vụ của ứng dụng, đặc biệt là những lớp không phụ thuộc vào giao diện người dùng (UI) hoặc cơ sở dữ liệu (DB).
- Lớp service thường chứa các phương thức gọi đến repository để lấy dữ liệu, xử lý dữ liệu đó theo logic nghiệp vụ, và sau đó trả về kết quả cho controller hoặc các thành phần khác.