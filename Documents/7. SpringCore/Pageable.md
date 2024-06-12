
---

Điều này có lẽ ai cũng biết, đó là đa phần chúng ta không lấy toàn bộ dữ liệu từ Database lên, mà chỉ lấy một số lượng nhất định, và chia nó thành nhiều trang.

Tôi chia dữ liệu thành nhiều trang khác nhau, với mỗi trang thì sẽ lấy ra các bài viết cần thiết.

Tạo sao cần dùng phân trang? vì nó giúp tiết kiệm băng thông và tăng tốc xử lý, ngoài ra nó cũng giảm thiểu việc hiển thị các thông tin thừa không cần thiết.

Chúng ta sẽ tìm hiểu các làm việc này bằng `JPA Pageable`.
