
---

**Có bốn đặc trưng chính của lập trình hướng đối tượng: đóng gói, kế thừa, đa hình và trừu tượng.**

1. **Đóng gói (Encapsulation)**: Là việc ẩn thông tin của lớp bên trong lớp đó, không cho phép các chương trình bên ngoài truy cập trực tiếp, mà thực hiện thao tác và truy cập vào thông tin ẩn thông qua các phương thức của lớp. Đóng gói tốt giúp giảm sự phụ thuộc (coupling) giữa các thành phần.
2. **Kế thừa (Inheritance)**: Là quá trình tạo ra một lớp mới từ một lớp đã tồn tại. Lớp mới kế thừa các thuộc tính và hành vi của lớp cha, đồng thời có thể mở rộng thêm các khả năng mới, giúp tăng đáng kể khả năng tái sử dụng và bảo trì của chương trình. Trong Java, kế thừa là đơn kế thừa, nghĩa là một lớp con chỉ có một lớp cha duy nhất.
3. **Đa hình (Polymorphism)**: Là khả năng một hành vi có thể biểu hiện dưới nhiều hình thức khác nhau. Đa hình cho phép thay đổi mã được liên kết với chương trình trong thời gian chạy mà không cần sửa đổi mã nguồn. Để thực hiện đa hình, cần ba yếu tố: kế thừa, ghi đè (override), và tham chiếu của lớp cha đến đối tượng của lớp con.
    - **Đa hình tĩnh (Static Polymorphism)**: Cùng một phương thức có các danh sách tham số khác nhau thông qua cơ chế nạp chồng (overloading), và có thể xử lý khác nhau dựa trên các tham số khác nhau.
    - **Đa hình động (Dynamic Polymorphism)**: Là việc ghi đè phương thức của lớp cha trong lớp con. Trong thời gian chạy, loại thực tế của đối tượng được tham chiếu sẽ được xác định, và phương thức tương ứng sẽ được gọi dựa trên loại thực tế đó.
4. **Trừu tượng (Abstraction)**: Là việc trừu tượng hóa các sự vật khách quan trong mã lập trình.