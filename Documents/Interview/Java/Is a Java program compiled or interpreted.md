
---
**Ngôn ngữ biên dịch**  
Trước khi chương trình chạy, mã nguồn sẽ được biên dịch thành mã máy (machine code) dưới dạng tệp nhị phân có thể chạy được thông qua trình biên dịch (compiler). Khi chương trình được thực thi sau đó, nó không cần phải biên dịch lại.

- **Ưu điểm**: Trình biên dịch thường có quá trình tiền biên dịch để tối ưu hóa mã. Vì quá trình biên dịch chỉ thực hiện một lần và không cần biên dịch lại trong thời gian chạy, hiệu suất thực thi của chương trình thuộc ngôn ngữ biên dịch rất cao và có thể chạy độc lập mà không phụ thuộc vào môi trường cục bộ.
- **Nhược điểm**: Nếu cần chỉnh sửa sau khi biên dịch, bạn phải biên dịch lại toàn bộ mô-đun. Trong quá trình biên dịch, mã máy được tạo ra dựa trên môi trường hoạt động tương ứng, dẫn đến vấn đề về tính di động giữa các hệ điều hành khác nhau. Do đó, cần biên dịch các tệp thực thi khác nhau tùy theo môi trường hệ điều hành.
- **Tóm tắt**: Tốc độ thực thi nhanh, hiệu quả cao; phụ thuộc vào trình biên dịch và khả năng tương thích đa nền tảng thấp.
- **Ngôn ngữ tiêu biểu**: C, C++, Pascal, Objective-C, Swift.

**Ngôn ngữ thông dịch**

- **Định nghĩa**: Mã nguồn của ngôn ngữ thông dịch không được dịch trực tiếp thành mã máy, mà trước tiên được dịch sang mã trung gian (intermediate code). Sau đó, trình thông dịch (interpreter) sẽ giải thích và chạy mã trung gian này. Mã nguồn được dịch sang mã máy trong thời gian chạy, từng câu lệnh được dịch và thực thi lần lượt cho đến khi kết thúc.
- **Ưu điểm**: Tính tương thích nền tảng tốt, có thể chạy trên bất kỳ môi trường nào miễn là có cài đặt trình thông dịch (ví dụ: máy ảo). Linh hoạt, bạn có thể chỉnh sửa mã trực tiếp và triển khai nhanh mà không cần thời gian ngừng hoạt động hay bảo trì.
- **Nhược điểm**: Mỗi lần chạy đều phải giải thích lại, do đó hiệu suất không bằng ngôn ngữ biên dịch.
- **Tóm tắt**: Ngôn ngữ thông dịch chạy chậm và hiệu suất thấp; phụ thuộc vào trình thông dịch nhưng có khả năng tương thích đa nền tảng tốt.
- **Ngôn ngữ tiêu biểu**: JavaScript, Python, Erlang, PHP, Perl, Ruby.

**Đối với ngôn ngữ Java**: Mã nguồn của Java sẽ được biên dịch thành bytecode thông qua javac, sau đó được chuyển đổi thành mã máy thông qua JVM (Máy ảo Java). Nói cách khác, Java kết hợp cả quá trình chạy thông dịch và chạy biên dịch, vì vậy nó có thể được gọi là ngôn ngữ lai (hybrid) hoặc bán biên dịch (semi-compiled).