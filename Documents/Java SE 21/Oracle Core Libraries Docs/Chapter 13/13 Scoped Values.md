
---
Trong Java SE 21, **Scoped Values** là một tính năng xem trước (preview feature) cho phép chia sẻ an toàn và hiệu quả các giá trị giữa các phương thức mà không cần sử dụng tham số phương thức. Điều này giúp đơn giản hóa việc truyền dữ liệu trong các chuỗi lời gọi phương thức phức tạp và cải thiện khả năng bảo trì mã nguồn.

[Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/scoped-values.html?utm_source=chatgpt.com)

**Lưu ý:** Vì Scoped Values là một tính năng xem trước, thiết kế, đặc tả và triển khai của nó có thể thay đổi trong các phiên bản Java SE tương lai. Để biên dịch và chạy mã chứa các tính năng xem trước, bạn cần chỉ định các tùy chọn dòng lệnh bổ sung.

[Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/scoped-values.html?utm_source=chatgpt.com)

**Để biết thêm thông tin chi tiết:**

- Tham khảo lớp `ScopedValue` trong đặc tả API của Java SE.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/lang/ScopedValue.html?utm_source=chatgpt.com)
    
- Xem JEP 446 để có thông tin nền tảng về Scoped Values.
    

Việc sử dụng Scoped Values có thể cải thiện cấu trúc và khả năng bảo trì của mã nguồn, đặc biệt trong các ứng dụng phức tạp yêu cầu chia sẻ dữ liệu giữa nhiều phương thức và luồng.