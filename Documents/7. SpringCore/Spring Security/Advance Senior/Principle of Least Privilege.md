
---
Hãy đi sâu vào **Principle of Least Privilege (Nguyên tắc cấp quyền tối thiểu)** trong bối cảnh bảo mật ứng dụng và cách áp dụng nó trong Spring Security. Tôi sẽ giải thích nguyên tắc này là gì và cách triển khai cụ thể trong Spring Security để đảm bảo an toàn.

---

### **1. Nguyên tắc cấp quyền tối thiểu (Principle of Least Privilege)**

#### **a. Định nghĩa**
- **Principle of Least Privilege (PoLP)**:
  - Là một nguyên tắc bảo mật yêu cầu mỗi người dùng, quy trình, hoặc hệ thống chỉ được cấp quyền truy cập tối thiểu cần thiết để thực hiện nhiệm vụ của họ, không hơn không kém.
  - Mục tiêu: Giảm thiểu rủi ro nếu tài khoản bị xâm phạm, hạn chế phạm vi thiệt hại.
- **Ví dụ**:
  - Một nhân viên kế toán chỉ nên có quyền đọc/sửa dữ liệu tài chính, không được phép truy cập mã nguồn hệ thống.
  - Một ứng dụng chỉ nên có quyền đọc cơ sở dữ liệu, không cần quyền xóa.

#### **b. Tại sao quan trọng?**
- **Giảm bề mặt tấn công**: Nếu kẻ tấn công chiếm được tài khoản, chúng chỉ có thể làm những gì tài khoản đó được phép, không hơn.
- **Hạn chế lỗi người dùng**: Ngăn người dùng vô tình thực hiện hành động nguy hiểm (VD: xóa dữ liệu hệ thống).
- **Tuân thủ quy định**: Nhiều tiêu chuẩn bảo mật (như GDPR, PCI DSS) yêu cầu áp dụng PoLP.

#### **c. Nguyên tắc trong thực tế**
- **Cấp quyền cụ thể**: Chỉ cho phép truy cập tài nguyên cần thiết (endpoint, file, database).
- **Thời gian giới hạn**: Quyền chỉ tồn tại trong thời gian cần thiết (VD: token hết hạn nhanh).
- **Phân vai trò**: Gán vai trò (role) với quyền hạn rõ ràng, tránh cấp quyền "toàn năng" (admin) không cần thiết.

---
