


---

### Giới thiệu

**try-with-resources** là một tính năng mạnh mẽ được giới thiệu trong Java 7, giúp đơn giản hóa việc quản lý tài nguyên, đặc biệt là các đối tượng triển khai giao diện **AutoCloseable**. Tính năng này đảm bảo rằng các tài nguyên như file, kết nối cơ sở dữ liệu, hoặc bất kỳ đối tượng nào thực hiện giao diện AutoCloseable sẽ được đóng tự động ngay cả khi có ngoại lệ xảy ra.

**AutoCloseable** là một giao diện đánh dấu, có nghĩa là một lớp triển khai giao diện này phải cung cấp một phương thức `close()` để đóng tài nguyên.

### Tại sao cần try-with-resources?

- **Tránh rò rỉ tài nguyên:** Nếu không đóng các tài nguyên một cách chính xác, chúng có thể dẫn đến tình trạng rò rỉ bộ nhớ và các vấn đề khác.
- **Mã gọn gàng hơn:** So với cách viết truyền thống sử dụng khối `try-finally`, `try-with-resources` giúp mã trở nên ngắn gọn và dễ đọc hơn.
- **Đảm bảo tài nguyên được đóng:** Ngay cả khi có ngoại lệ xảy ra trong khối `try`, phương thức `close()` của các đối tượng sẽ vẫn được gọi.
  
  
  Cách sử dụng try-with-resources
  
  ```Java
  try (FileInputStream fis = new FileInputStream("input.txt")) {
    // Sử dụng fis để đọc dữ liệu từ file
  } catch (IOException e) {
    // Xử lý ngoại lệ
  }
```

Trong ví dụ trên:

- `FileInputStream` triển khai giao diện `AutoCloseable`.
- Khi khối `try` kết thúc (bằng cách hoàn thành bình thường hoặc do một ngoại lệ), phương thức `close()` của `fis` sẽ được tự động gọi để đóng file.

### Cơ chế hoạt động

1. Các đối tượng được khai báo ngay sau từ khóa `try` phải triển khai giao diện `AutoCloseable`.
2. Khi khối `try` kết thúc:
    - Nếu không có ngoại lệ xảy ra, phương thức `close()` của từng đối tượng sẽ được gọi theo thứ tự ngược lại với thứ tự khai báo.
    - Nếu có ngoại lệ xảy ra, các ngoại lệ sẽ được ném ra. Sau đó, phương thức `close()` của từng đối tượng sẽ được gọi theo thứ tự ngược lại.

### Ưu điểm của try-with-resources

- **Tăng tính đọc của mã:** Mã trở nên rõ ràng hơn, dễ bảo trì.
- **Giảm lỗi:** Giảm thiểu khả năng quên đóng tài nguyên.
- **Tăng hiệu suất:** Đóng tài nguyên kịp thời giúp giải phóng tài nguyên hệ thống.
