
---
**Trách nhiệm đơn nhất của đối tượng** (**Single responsibility of the object**): Đối tượng mà chúng ta thiết kế và tạo ra phải có trách nhiệm rõ ràng. Ví dụ, với danh mục sản phẩm, các thuộc tính và phương thức liên quan trong đó phải liên quan đến sản phẩm, không được chứa nội dung không liên quan như đơn hàng. Lớp có thể là mô-đun, thư viện, tập hợp, không chỉ giới hạn ở lớp.

**Nguyên tắc thay thế (Principle of Substitution)**: Lớp con có thể hoàn toàn thay thế lớp cha và ngược lại. Nguyên tắc này thường được sử dụng khi triển khai giao diện. Vì lớp con có thể thay thế hoàn toàn lớp cơ sở (lớp cha), lớp cha có thể có nhiều lớp con, giúp dễ dàng mở rộng trong các lần mở rộng chương trình sau này mà không cần sửa đổi mã nguồn. Ví dụ, triển khai của giao diện IA là A, nhưng do yêu cầu dự án thay đổi, cần một triển khai mới, ta có thể thay thế trực tiếp tại nơi tiêm container.

**Luật Demeter (Dimmitt's Law)**, còn gọi là nguyên tắc tối thiểu hoặc giảm thiểu liên kết (minimum coupling): Khi thiết kế hoặc phát triển chương trình, cần cố gắng đạt được sự gắn kết cao (high cohesion) và liên kết thấp (low coupling). Khi hai lớp tương tác, một sự phụ thuộc được tạo ra. Luật Demeter cho rằng càng ít sự phụ thuộc như vậy càng tốt. Giống như khi một hàm khởi tạo được tiêm vào một đối tượng cha, khi một đối tượng cần phụ thuộc, nó không quan tâm cách triển khai bên trong mà chỉ cần tiêm triển khai tương ứng vào container. Điều này không chỉ tuân theo nguyên tắc thay thế mà còn hỗ trợ giảm liên kết.

**Nguyên tắc mở-đóng (Open-Closed Principle)**: Mở để mở rộng, đóng để sửa đổi. Khi yêu cầu dự án thay đổi, không sửa đổi mã nguồn gốc mà mở rộng trên nền tảng ban đầu.

**Nguyên tắc đảo ngược phụ thuộc (Dependency Inversion Principle)**: Các mô-đun cấp cao không nên phụ thuộc trực tiếp vào triển khai cụ thể của các mô-đun cấp thấp, mà nên phụ thuộc vào abstraction của chúng. Giao diện và lớp trừu tượng không nên phụ thuộc vào lớp triển khai, mà các lớp triển khai phụ thuộc vào giao diện hoặc lớp trừu tượng.

**Nguyên tắc phân tách giao diện (Interface Segregation Principle)**: Khi một đối tượng tương tác với đối tượng khác, nội dung phụ thuộc nên được giảm thiểu. Nghĩa là khi thiết kế giao diện, nội dung giao diện cần được tối giản trong khi vẫn tuân theo nguyên tắc trách nhiệm đơn nhất của đối tượng.

**Phiên bản ngắn gọn**:

- **Trách nhiệm đơn nhất**: Thiết kế đối tượng cần độc lập, không nên tạo đối tượng chung chung.
- **Nguyên tắc mở-đóng**: Giảm thiểu sửa đổi đối tượng.
- **Thay thế**: Abstraction có thể được thay thế bằng đối tượng cụ thể trong quá trình mở rộng chương trình (giao diện, lớp cha, đối tượng có thể là lớp triển khai, và đối tượng có thể được thay bằng lớp con).
- **Demeter**: Gắn kết cao, liên kết thấp. Tránh phụ thuộc vào chi tiết.
- **Đảo ngược phụ thuộc**: Lập trình hướng abstraction. Nghĩa là tham số truyền vào hoặc giá trị trả về có thể sử dụng kiểu lớp cha hoặc kiểu giao diện. Nói rộng hơn: lập trình dựa trên giao diện, thiết kế khung giao diện trước.
- **Phân tách giao diện**: Kích thước thiết kế giao diện cần vừa phải. Quá lớn dẫn đến ô nhiễm, quá nhỏ gây khó khăn khi gọi.