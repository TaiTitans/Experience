
---
**Error (Lỗi):**

- Là những vấn đề nghiêm trọng mà JVM không thể giải quyết, chẳng hạn như tràn ngăn xếp (stack overflow), tràn bộ nhớ (memory overflow), v.v.
- Đây là những lỗi mà chương trình không thể xử lý được.

**Ví dụ:**

- **StackOverflowError**: Lỗi tràn ngăn xếp, thường xảy ra khi đệ quy vô hạn.
- **OOMException (OutOfMemoryError)**: Lỗi hết bộ nhớ, xảy ra khi JVM không còn đủ bộ nhớ để cấp phát cho chương trình.

**Các vấn đề khác (thường là Exception):**

- Là những vấn đề thông thường do lỗi lập trình hoặc các yếu tố bên ngoài ngẫu nhiên gây ra. Những vấn đề này có thể được xử lý trong mã nguồn.
- Ví dụ: ngoại lệ con trỏ null (NullPointerException), chỉ số mảng vượt quá giới hạn (ArrayIndexOutOfBoundsException), v.v.

---
### Giải thích bổ sung:

- **Error**: Đại diện cho các sự cố nghiêm trọng ở cấp độ hệ thống (system-level issues), thường không thể phục hồi trong chương trình. Lập trình viên hiếm khi xử lý chúng mà thay vào đó cần sửa cấu hình hệ thống hoặc mã nguồn (ví dụ: tăng kích thước stack hoặc heap).
- **Exception**: Đại diện cho các lỗi có thể kiểm soát được (program-level issues), thường được xử lý bằng try-catch để chương trình tiếp tục chạy mà không crash.

**Lưu ý**: Trong nội dung gốc, OOMException không phải là tên chính xác; đúng ra là OutOfMemoryError. Tôi giữ nguyên cách viết trong bản dịch để trung thành với nguyên văn, nhưng cần lưu ý đây là một lỗi nhỏ trong tài liệu gốc.