
---
Java NIO (New Input/Output) là một API được giới thiệu trong Java để cung cấp các cơ chế I/O hiệu quả và linh hoạt hơn so với các phương thức I/O truyền thống. Nó bao gồm các khái niệm chính như **Buffer**, **Charset**, **Channel**, và **Selectable Channel**, giúp cải thiện hiệu suất và khả năng quản lý I/O trong các ứng dụng Java.

**Buffer:**

Buffer là các container cho dữ liệu, được sử dụng để lưu trữ và thao tác với dữ liệu trong quá trình I/O. Mỗi loại buffer được thiết kế cho một kiểu dữ liệu nguyên thủy cụ thể, chẳng hạn như `ByteBuffer` cho byte, `CharBuffer` cho ký tự, và `IntBuffer` cho số nguyên. Buffer hoạt động như một vùng đệm trung gian giữa nguồn dữ liệu và đích, cho phép đọc và ghi dữ liệu một cách hiệu quả.

**Charset:**

Charset đại diện cho các ánh xạ giữa byte và các ký tự Unicode. Nó bao gồm các bộ mã hóa (encoder) và giải mã (decoder) để chuyển đổi giữa byte và ký tự. Ví dụ, `UTF-8` và `US-ASCII` là các charset phổ biến được sử dụng để mã hóa văn bản.

**Channel:**

Channel đại diện cho một kết nối đến một thực thể có khả năng thực hiện các hoạt động I/O, chẳng hạn như tệp, socket mạng, hoặc thiết bị phần cứng. Không giống như các luồng (stream) truyền thống, channel có thể không đồng bộ và hỗ trợ các hoạt động I/O không chặn (non-blocking I/O). Các loại channel bao gồm `FileChannel` cho tệp, `SocketChannel` cho kết nối TCP, và `DatagramChannel` cho kết nối UDP.

**Selectable Channel:**

Selectable Channel là một loại channel có thể được multiplexed, nghĩa là chúng có thể xử lý nhiều hoạt động I/O trong một channel duy nhất. Điều này được thực hiện thông qua việc sử dụng `Selector`, cho phép một luồng đơn lẻ quản lý nhiều kết nối mà không cần tạo ra nhiều luồng, cải thiện hiệu suất và khả năng mở rộng của ứng dụng.

Việc sử dụng Java NIO cho phép các nhà phát triển xây dựng các ứng dụng có hiệu suất cao hơn, đặc biệt là trong các trường hợp yêu cầu xử lý I/O đồng thời hoặc khối lượng lớn dữ liệu. Nó cung cấp các công cụ mạnh mẽ để quản lý I/O một cách linh hoạt và hiệu quả.