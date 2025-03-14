
---
**Mục đích chính là tiết kiệm bộ nhớ mà chuỗi String chiếm dụng.**  
Chuỗi String chiếm phần lớn không gian trong bộ nhớ heap của hầu hết các chương trình Java, và phần lớn các chuỗi chỉ chứa các ký tự Latin-1, vốn chỉ cần 1 byte để biểu diễn.

Trước JDK 9, JVM sử dụng mảng char để lưu trữ chuỗi, trong đó mỗi phần tử char chiếm 2 byte. Vì vậy, ngay cả khi chuỗi chỉ cần 1 byte để biểu diễn, nó vẫn phải được cấp phát theo kích thước 2 byte, dẫn đến lãng phí một nửa không gian bộ nhớ.

Từ JDK 9 trở đi, với mỗi chuỗi, hệ thống sẽ kiểm tra trước xem chuỗi đó chỉ chứa các ký tự Latin-1 hay không. Nếu đúng, bộ nhớ sẽ được cấp phát theo chuẩn 1 byte; nếu không, bộ nhớ sẽ được cấp phát theo chuẩn 2 byte. Nhờ vậy, việc sử dụng bộ nhớ được cải thiện, số lần thu gom rác (GC) cũng giảm, từ đó nâng cao hiệu suất.

Tuy nhiên, tập mã Latin-1 hỗ trợ một số lượng ký tự hạn chế và không bao gồm các ký tự tiếng Trung. Vì vậy, đối với chuỗi tiếng Trung, mã hóa UTF-16 (2 byte) được sử dụng. Do đó, trong trường hợp này, không có sự khác biệt giữa việc triển khai bằng byte[] hay char[].