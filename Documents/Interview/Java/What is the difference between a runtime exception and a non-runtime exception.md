
---
**Unchecked Exception (Ngoại lệ không kiểm tra):**

- Bao gồm RuntimeException và Error.
- Tất cả các ngoại lệ khác ngoài hai loại này được gọi là **Checked Exception (Ngoại lệ cần kiểm tra)**.

**RuntimeException:**

- Được gây ra bởi lỗi trong chương trình (program error).
- Chương trình nên được sửa chữa để tránh các ngoại lệ này xảy ra.

**Checked Exception:**

- Là ngoại lệ do một môi trường cụ thể gây ra (ví dụ: tệp cần đọc không tồn tại, tệp rỗng, hoặc ngoại lệ SQL bất thường).
- Phải được xử lý (bằng cách bắt - catch, hoặc ném - throw), nếu không mã sẽ không vượt qua quá trình biên dịch.

---
### Giải thích bổ sung:

1. **Unchecked Exception**:
    - Bao gồm:
        - **RuntimeException**: Ví dụ NullPointerException, ArrayIndexOutOfBoundsException – thường do lỗi lập trình viên (như truy cập null hoặc vượt chỉ số mảng).
        - **Error**: Ví dụ StackOverflowError, OutOfMemoryError – lỗi hệ thống nghiêm trọng.
    - Không bắt buộc xử lý trong mã, chương trình vẫn biên dịch được, nhưng có thể crash khi chạy nếu không được sửa.
2. **Checked Exception**:
    - Ví dụ: FileNotFoundException, SQLException.
    - Liên quan đến các yếu tố ngoài tầm kiểm soát của lập trình viên (như tệp hoặc cơ sở dữ liệu).
    - Trình biên dịch yêu cầu xử lý (dùng try-catch hoặc khai báo throws), nếu không sẽ báo lỗi biên dịch.