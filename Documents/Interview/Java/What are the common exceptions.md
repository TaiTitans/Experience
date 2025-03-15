
---
### **Các RuntimeException phổ biến:**

- **ClassCastException**: Ngoại lệ chuyển đổi kiểu – xảy ra khi cố gắng ép kiểu không hợp lệ.
- **IndexOutOfBoundsException**: Ngoại lệ vượt chỉ số – xảy ra khi truy cập ngoài giới hạn của mảng hoặc danh sách.
- **NullPointerException**: Ngoại lệ con trỏ null – xảy ra khi truy cập vào một đối tượng hoặc phương thức của một tham chiếu null.
- **ArrayStoreException**: Ngoại lệ lưu trữ mảng – xảy ra khi cố gắng lưu trữ một kiểu dữ liệu không tương thích vào mảng.
- **NumberFormatException**: Ngoại lệ định dạng số – xảy ra khi chuỗi không thể chuyển đổi thành số (ví dụ: Integer.parseInt("abc")).
- **ArithmeticException**: Ngoại lệ toán học – xảy ra khi thực hiện phép toán không hợp lệ (ví dụ: chia cho 0).

### **Checked Exception (Ngoại lệ cần kiểm tra):**

- **NoSuchFieldException**: Ngoại lệ phản xạ – xảy ra khi không tìm thấy trường (field) tương ứng trong một lớp khi sử dụng reflection.
- **ClassNotFoundException**: Ngoại lệ không tìm thấy lớp – xảy ra khi không tìm thấy định nghĩa của một lớp (ví dụ: khi gọi Class.forName()).
- **IllegalAccessException**: Ngoại lệ quyền truy cập – xảy ra khi vi phạm quyền bảo mật, chẳng hạn như gọi một phương thức private thông qua reflection.
  
  
  ### Giải thích bổ sung:

- **RuntimeException**: Là các ngoại lệ không bắt buộc phải xử lý (unchecked exceptions), thường do lỗi lập trình (programming errors). Chương trình có thể tiếp tục chạy nếu không xử lý, nhưng thường gây crash.
- **Checked Exception**: Là các ngoại lệ bắt buộc phải xử lý (thông qua try-catch hoặc khai báo throws), thường liên quan đến các tình huống ngoài tầm kiểm soát của lập trình viên (như lỗi hệ thống hoặc tài nguyên).