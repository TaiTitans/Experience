
---
**Hãy cùng tìm hiểu đối tượng bất biến (immutable object) là gì.**  
Một đối tượng được gọi là bất biến nếu nó không thể thay đổi trạng thái sau khi được tạo ra. Trạng thái không thể thay đổi có nghĩa là các biến thành viên trong đối tượng không thể bị thay đổi, bao gồm: giá trị của các kiểu dữ liệu cơ bản không thể thay đổi, các biến kiểu tham chiếu không thể trỏ đến các đối tượng khác, và trạng thái của các đối tượng được trỏ đến bởi kiểu tham chiếu cũng không thể thay đổi.

**Hãy xem xét mã nguồn của lớp String trong Java 8:**

```java
public final class String
    implements java.io.Serializable, Comparable<String>, CharSequence {
    /** Mảng giá trị được sử dụng để lưu trữ ký tự. */
    private final char value[];

    /** Bộ nhớ đệm cho mã băm của chuỗi */
    private int hash; // Mặc định là 0

```
Từ mã nguồn, có thể thấy đối tượng String thực chất là một chuỗi ký tự được lưu trữ trong mảng value.

- Mảng value được đánh dấu bằng từ khóa final, và giá trị của biến được sửa đổi bởi final không thể thay đổi. Do đó, value không thể trỏ đến các đối tượng khác.
- Tất cả các trường bên trong lớp String đều là private (riêng tư). Hơn nữa, String không cung cấp bất kỳ phương thức nào để thay đổi trạng thái bên trong, vì vậy mảng value không thể bị thay đổi.  
    Vậy nên, String là bất biến.

**Vậy tại sao String được thiết kế là bất biến?**  
Có một số lý do cho điều này:

1. **An toàn luồng (Thread-safe)**: Vì chuỗi bất biến, cùng một thể hiện chuỗi có thể được chia sẻ giữa nhiều luồng mà không cần lo lắng về xung đột, do đó nó vốn đã an toàn luồng.
2. **Hỗ trợ ánh xạ băm và bộ nhớ đệm**: Vì giá trị băm của String thường được sử dụng (ví dụ: làm khóa trong Map), tính bất biến đảm bảo giá trị băm không thay đổi và không cần tính toán lại.
3. **Lý do bảo mật**: Địa chỉ mạng, URL, đường dẫn tệp và mật khẩu thường được lưu trữ dưới dạng String. Nếu String không bất biến, nó có thể gây ra nhiều rủi ro bảo mật. Ví dụ, nếu bạn lưu mật khẩu dưới dạng String, nó sẽ tồn tại trong bộ nhớ cho đến khi bộ thu gom rác (garbage collector) dọn dẹp. Nếu String không bất biến, mật khẩu có thể bị thay đổi, dẫn đến rủi ro bảo mật.
4. **Tối ưu hóa vùng nhớ hằng chuỗi (String Constant Pool)**: Sau khi đối tượng String được tạo, nó sẽ được lưu vào vùng nhớ hằng chuỗi. Lần tiếp theo cần tạo một đối tượng giống hệt, tham chiếu từ bộ nhớ đệm sẽ được trả về trực tiếp, giúp tiết kiệm tài nguyên.

**Vì String bất biến, nó có nhiều phương thức như substring, replace, replaceAll. Những phương thức này dường như thay đổi đối tượng String? Giải thích thế nào?**  
Thực tế thì không. Mỗi khi chúng ta gọi một phương thức như replace, một đối tượng mới sẽ được tạo ra trong bộ nhớ heap. Mảng value của đối tượng mới sau đó trỏ đến một đối tượng khác, chứ không phải thay đổi đối tượng ban đầu.