
---
Khi chúng ta chỉ chạy các instance Java trong JVM cục bộ, thì không cần phải thực hiện tuần tự hóa (serialization) và giải tuần tự hóa (deserialization). Tuy nhiên, khi cần lưu trữ các đối tượng trong bộ nhớ vào ổ đĩa, cơ sở dữ liệu, khi cần tương tác với trình duyệt, hoặc khi cần triển khai RPC, thì quá trình tuần tự hóa và giải tuần tự hóa là bắt buộc.

Hai trường hợp đầu tiên đặt ra một câu hỏi lớn: Chúng ta dường như không thực hiện tuần tự hóa hay giải tuần tự hóa khi tương tác với trình duyệt hoặc khi lưu trữ đối tượng vào cơ sở dữ liệu, vì chúng ta không sử dụng interface `Serializable`, nhưng tại sao mọi thứ vẫn hoạt động?

Hãy đi đến kết luận:  
**Bất cứ khi nào cần lưu trữ hoặc truyền tải đối tượng qua mạng, chúng ta đều phải thực hiện tuần tự hóa và giải tuần tự hóa.**

### Lý do:

Liệu máy chủ có thực sự không sử dụng interface `Serializable` khi tương tác với trình duyệt? Trên thực tế, định dạng JSON là cách chuyển đổi một đối tượng thành chuỗi (string). Dữ liệu mà máy chủ gửi đến trình duyệt thực chất là một chuỗi, hãy xem mã nguồn của kiểu `String` trong Java:
```java
public final class String
    implements java.io.Serializable, Comparable<String>, CharSequence {
    
    /** Giá trị dùng để lưu trữ ký tự */
    private final char value[];

    /** Bộ nhớ đệm mã băm của chuỗi */
    private int hash; // Mặc định là 0

    /** Dùng serialVersionUID từ JDK 1.0.2 để đảm bảo tương thích */
    private static final long serialVersionUID = -6849794470754667710L;

    ......
}
```

Kiểu `String` trong Java đã triển khai interface `Serializable` và xác định rõ giá trị của `serialVersionUID`.

Tiếp theo, hãy xét trường hợp đối tượng được lưu trữ vào cơ sở dữ liệu. Đây là đoạn mã chèn dữ liệu trong file mapping của MyBatis:
```java
<insert id="insertUser" parameterType="org.tyshawn.bean.User">
    INSERT INTO t_user(name, age) VALUES (#{name}, #{age})
</insert>
```
Thay vì lưu trữ toàn bộ đối tượng vào cơ sở dữ liệu, chúng ta chỉ lưu các thuộc tính của đối tượng. Những thuộc tính này (ví dụ như `Date` và `String`) đều đã triển khai interface `Serializable`, do đó quá trình tuần tự hóa vẫn diễn ra ở cấp độ thuộc tính thay vì toàn bộ đối tượng.


---
## Implementing Serialization and Deserialization Why implement the Serializable interface?

Khi một lớp trong Java triển khai interface `Serializable`, JVM sẽ nhận biết điều này khi lớp được tải, sau đó hỗ trợ chúng ta tự động thực hiện tuần tự hóa (serialization) và giải tuần tự hóa (deserialization) ở tầng thấp khi đối tượng được khởi tạo.

Nếu kiểu đối tượng được ghi không phải là `String`, `Array`, `Enum` và chưa triển khai interface `Serializable`, một ngoại lệ `NotSerializableException` sẽ bị ném ra trong quá trình tuần tự hóa.

Dưới đây là đoạn mã nguồn xử lý trong Java:
```java
// Các trường hợp còn lại
if (obj instanceof String) {
    writeString((String) obj, unshared);
} else if (cl.isArray()) {
    writeArray(obj, desc, unshared);
} else if (obj instanceof Enum) {
    writeEnum((Enum<?>) obj, desc, unshared);
} else if (obj instanceof Serializable) {
    writeOrdinaryObject(obj, desc, unshared);
} else {
    if (extendedDebugInfo) {
        throw new NotSerializableException(
            cl.getName() + "\n" + debugInfoStack.toString());
    } else {
        throw new NotSerializableException(cl.getName());
    }
}
```
Nếu một đối tượng không thuộc các loại `String`, `Array`, `Enum` và cũng không triển khai `Serializable`, chương trình sẽ ném ra ngoại lệ `NotSerializableException`, chỉ rõ tên lớp không thể tuần tự hóa.

---
Why do you need to display the value of the specified serialVersionUID after implementing the Serializable interface?

Nếu `serialVersionUID` không được chỉ định rõ ràng, JVM sẽ tự động tạo một giá trị `serialVersionUID` dựa trên các thuộc tính của lớp trong quá trình tuần tự hóa. Giá trị này sẽ được tuần tự hóa cùng với các thuộc tính của đối tượng trước khi được lưu trữ hoặc truyền qua mạng.

Khi giải tuần tự hóa, JVM sẽ tự động tạo một phiên bản `serialVersionUID` mới dựa trên các thuộc tính hiện tại của lớp, sau đó so sánh với `serialVersionUID` đã được tạo trong quá trình tuần tự hóa. Nếu hai giá trị này giống nhau, quá trình giải tuần tự hóa sẽ thành công. Ngược lại, nếu không trùng khớp, chương trình sẽ ném lỗi.

Nếu chúng ta chỉ định rõ ràng `serialVersionUID`, JVM vẫn sẽ sử dụng nó khi tuần tự hóa và giải tuần tự hóa, đảm bảo rằng giá trị `serialVersionUID` của phiên bản cũ và mới luôn giống nhau, giúp quá trình giải tuần tự hóa diễn ra suôn sẻ ngay cả khi lớp có sự thay đổi.

Trong trường hợp lớp không thay đổi sau khi được viết, việc không chỉ định `serialVersionUID` sẽ không gây ra vấn đề. Tuy nhiên, trong thực tế phát triển, các lớp thường xuyên được cập nhật và thay đổi. Khi đó, nếu không có `serialVersionUID`, đối tượng cũ sẽ không thể được giải tuần tự hóa thành công, gây ra lỗi.

Vì lý do này, trong thực tế phát triển, **chúng ta luôn chỉ định rõ `serialVersionUID`** để đảm bảo tính tương thích giữa các phiên bản của lớp.

---
## Why aren't static properties serialized ?

Vì quá trình tuần tự hóa (serialization) áp dụng cho đối tượng (instance), trong khi các thuộc tính `static` được tải lên cùng với lớp khi lớp được nạp vào bộ nhớ, nên **các thuộc tính `static` không được tuần tự hóa**.

Tuy nhiên, điều này dẫn đến một câu hỏi: **Tại sao `serialVersionUID` là một thuộc tính `static`, nhưng vẫn có vẻ như nó được tuần tự hóa?**

Thực tế, **`serialVersionUID` không hề được tuần tự hóa**. Khi một đối tượng được tuần tự hóa, JVM sẽ tự động tạo một giá trị `serialVersionUID`, sau đó gán giá trị này cho thuộc tính `serialVersionUID` nếu nó đã được khai báo trong lớp.

Nói cách khác, `serialVersionUID` chỉ được sử dụng để kiểm tra tính tương thích của lớp trong quá trình giải tuần tự hóa, nhưng chính nó **không phải là một phần của dữ liệu tuần tự hóa**.