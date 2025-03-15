
---
**Triển khai giao diện Cloneable và ghi đè phương thức clone().**

- Phương thức này thực hiện **sao chép nông (shallow copy)**, tức là nếu thuộc tính trong lớp có kiểu tham chiếu tự định nghĩa (custom reference type), chỉ tham chiếu được sao chép, còn đối tượng mà tham chiếu trỏ tới không được sao chép.
- Nếu lớp của các thuộc tính trong đối tượng cũng triển khai giao diện Cloneable, thì khi đối tượng được sao chép bằng clone(), các thuộc tính đó cũng được sao chép, tức là thực hiện **sao chép sâu (deep copy)**.

**Kết hợp với tuần tự hóa (serialization) để thực hiện sao chép sâu.**

- Sao chép đối tượng có thể được thực hiện thông qua lớp công cụ và trong lớp:
    - org.apache.commons.BeanUtils.PropertyUtils

### Giải thích bổ sung:

1. **Cloneable và clone()**:
    - Giao diện Cloneable là một marker interface (không có phương thức), dùng để đánh dấu rằng lớp hỗ trợ sao chép.
    - Phương thức clone() mặc định từ Object chỉ sao chép nông (copy tham chiếu của các thuộc tính). Để thực hiện sao chép sâu, cần ghi đè clone() và sao chép các đối tượng con nếu chúng cũng triển khai Cloneable.

- **Sao chép sâu qua tuần tự hóa**:
    - Một cách khác để sao chép sâu là tuần tự hóa đối tượng thành byte stream (dùng ObjectOutputStream) rồi giải tuần tự hóa (dùng ObjectInputStream) để tạo bản sao hoàn toàn độc lập.
- **BeanUtils và PropertyUtils**:
    - Thư viện Apache Commons BeanUtils cung cấp các công cụ để sao chép thuộc tính giữa các đối tượng. PropertyUtils có thể được dùng để sao chép sâu bằng cách sao chép từng thuộc tính, nhưng cần đảm bảo các thuộc tính tham chiếu cũng được sao chép đúng cách nếu cần deep copy.
