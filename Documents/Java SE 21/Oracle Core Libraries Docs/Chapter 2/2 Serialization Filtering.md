
---
Cơ chế **Serialization Filtering** trong Java được thiết kế để tăng cường bảo mật và độ tin cậy bằng cách kiểm tra các luồng dữ liệu được tuần tự hóa (serialized) trước khi chúng được giải tuần tự hóa (deserialized). Điều này giúp ngăn chặn các lỗ hổng bảo mật tiềm ẩn liên quan đến quá trình deserialization.

**Mục tiêu của Serialization Filtering:**

- **Giới hạn các lớp có thể được deserialized:** Chỉ cho phép deserialization các lớp phù hợp với ngữ cảnh cụ thể, ngăn chặn việc tải các lớp không mong muốn hoặc độc hại.
    
- **Cung cấp các chỉ số cho bộ lọc về kích thước và độ phức tạp của đồ thị đối tượng:** Giúp xác minh các hành vi bình thường của đồ thị đối tượng trong quá trình deserialization.
    
- **Cho phép các đối tượng được export qua RMI xác thực các lớp dự kiến trong các lời gọi:** Đảm bảo rằng chỉ các lớp mong đợi mới được deserialized trong các cuộc gọi RMI.
    

**Các loại bộ lọc (Filters):**

1. **JVM-wide filter:** Áp dụng cho mọi quá trình deserialization trong JVM. Cách bộ lọc này xác thực các lớp trong một quá trình deserialization cụ thể phụ thuộc vào cách nó được kết hợp với các bộ lọc khác.
    
2. **Stream-specific filter:** Xác thực các lớp từ một `ObjectInputStream` cụ thể.
    

**Cách triển khai bộ lọc:**

- **Pattern-based filter thông qua thuộc tính `jdk.serialFilter`:** Bộ lọc dựa trên mẫu bao gồm một chuỗi các mẫu có thể chấp nhận hoặc từ chối tên của các lớp, gói hoặc mô-đun cụ thể. Nó có thể đặt giới hạn về kích thước mảng, độ sâu của đồ thị, tổng số tham chiếu và kích thước luồng. Một trường hợp sử dụng điển hình là thêm các lớp đã được xác định là có thể gây nguy hiểm cho runtime Java vào danh sách từ chối. Việc chỉ định bộ lọc dựa trên mẫu với thuộc tính `jdk.serialFilter` không yêu cầu thay đổi ứng dụng.
    
- **Custom hoặc pattern-based stream-specific filter với API `ObjectInputFilter`:** Bạn có thể triển khai một bộ lọc bằng API `ObjectInputFilter`, sau đó đặt nó trên một `ObjectInputStream`. Bạn có thể tạo một bộ lọc dựa trên mẫu bằng cách gọi phương thức `Config.createFilter(String)` của API `ObjectInputFilter`.
    

**Lưu ý:** Mặc định, một bộ lọc serialization không được kích hoạt hoặc cấu hình. Quá trình lọc serialization sẽ không diễn ra trừ khi bạn đã chỉ định bộ lọc trong một thuộc tính hệ thống hoặc thuộc tính bảo mật, hoặc đặt nó với lớp `ObjectInputFilter`.

**Thực tiễn tốt nhất để ngăn chặn các lỗ hổng deserialization:**

- **Không deserialization dữ liệu không tin cậy:** Tránh deserialization dữ liệu từ các nguồn không đáng tin cậy để giảm nguy cơ bảo mật.
    
- **Sử dụng SSL để mã hóa và xác thực các kết nối giữa các ứng dụng:** Đảm bảo rằng dữ liệu được truyền tải một cách an toàn và không bị can thiệp.
    
- **Xác thực giá trị của các trường trước khi gán:** Ví dụ, kiểm tra các bất biến của đối tượng bằng cách sử dụng phương thức `readObject`.
    

**Bộ lọc tích hợp cho RMI:**

Java cung cấp các bộ lọc tích hợp cho RMI. Tuy nhiên, bạn nên sử dụng các bộ lọc này như điểm khởi đầu và cấu hình thêm danh sách từ chối hoặc mở rộng danh sách cho phép để tăng cường bảo vệ cho ứng dụng sử dụng RMI. Xem thêm tại [Built-in Filters](https://docs.oracle.com/en/java/javase/23/core/built-filters.html).

Để biết thêm thông tin về các chiến lược này và các chiến lược khác, hãy tham khảo phần "Serialization and Deserialization" trong [Secure Coding Guidelines for Java SE](https://docs.oracle.com/javase/8/docs/technotes/guides/serialization/filters/serialization-filtering.html).

Việc triển khai cơ chế Serialization Filtering đúng cách giúp bảo vệ ứng dụng Java khỏi các lỗ hổng bảo mật liên quan đến deserialization và đảm bảo rằng chỉ các đối tượng hợp lệ mới được deserialized.