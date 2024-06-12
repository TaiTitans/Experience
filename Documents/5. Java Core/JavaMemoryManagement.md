
------
## Tại sao cần biết về Java Memory Management

-> Java cung cấp một trình quản lý bộ nhớ và thu gôm rác tự động. Tuy nhiên là một lập trình viên ta cần biết về quản lý bộ nhớ là cần thiết vì nó sẽ mang lại lợi ích cho người lập trình ví dụ hiệu suất cao tối ưu chương trình.

----
Trong một ngôn ngữ lập trình việc quản lý bộ nhớ luôn cần thiết và quan trọng. Tuy nhiên thì java không giống NNLT khác JVM cụ thể là Garbage Collector có vai trò cấp phát bộ nhớ tự động.

---
- Java memory management bao gồm:
	- JVM Memory Structure
	- Working of Garbage Collector

### Java memory structure

JVM (Java Virtual Machine) định nghĩa các khu vực dữ liệu thời gian chạy khác nhau được sử dụng trong quá trình thực thi chương trình. Một số khu vực được tạo ra bởi JVM trong khi một số khác được tạo ra bởi các luồng (threads) được sử dụng trong chương trình. Tuy nhiên, khu vực bộ nhớ được tạo ra bởi JVM chỉ bị phá hủy khi JVM thoát. Các khu vực dữ liệu của luồng được tạo ra trong quá trình khởi tạo và bị phá hủy khi luồng thoát.

![[Pasted image 20240512172955.png]]



**Heap:**

- Là vùng dữ liệu thời gian chạy được chia sẻ và lưu trữ đối tượng trong thực tế trong bộ nhớ. Được khởi tạo khi khởi động máy ảo.
- Bộ nhớ này được phân bổ cho tất cả các thể hiện của lớp và mảng. Heap có thể có kích thước cố định hoặc động tuỳ thuộc vào cấu hình.
- JVM cung cấp quyển kiểm soát của người dùng để khởi tạo hoặc thay đổi kích thước của heap tuỳ theo yêu cầu. Khi một từ khoá mới được sử dụng, đối tượng được gán vào một khoảng trắng trong heap, nhưng tham chiếu của nó vẫn tồn tại trong ngăn xếp.
- Chỉ tồn tại 1 heap trong một tiến trình JVM đang chạy.

Ví dụ:
```
Scanner sc = new Scanner(System.in);
```
Câu lệnh trên tạo ra đối tượng của lớp Scanner được phân bổ vào heap trong khi tham chiếu 'sc' được đẩy vào ngăn xếp.

**Method Area:**

- Là một phần logic của vùng heap và được tạo khi khởi động máy ảo.
- Bộ nhớ này được phân bổ cho các cấu trúc lớp, dữ liệu phương thức và dữ liệu trường hàm tạo cũng như cho các giao diện hoặc phương thức đặc biệt được sử dụng trong lớp. Heap có thể có kích thước cố định hoặc động tùy thuộc vào cấu hình của hệ thống.
- Có thể có kích thước cố định hoặc mở rộng theo yêu cầu tính toán. Không cần phải liền kề.

**JVM Stacks:**
- Ngăn xếp được tạo cùng lúc khi một luồng được tạo và được sử dụng để lưu trữ dữ liệu cũng như một phần kết quả cần thiết trong khi trả về giá trị cho phương thức và thực hiện liên kết động.
- Ngăn xếp có thể có kích thước cố định hoặc động. Kích thước của ngăn xếp có thể được chọn độc lập khi nó được tạo.
- Bộ nhớ cho ngăn xếp không cần phải liền kề nhau.

**Native method Stacks:**
Còn được gọi là ngăn xếp C, ngăn xếp phương thức gốc không được viết bằng ngôn ngữ Java. Bộ nhớ này được phân bổ cho mỗi luồng khi nó được tạo. Và nó có thể có tính chất cố định hoặc động.

**Program counter (PC) registers:**

Mỗi luồng JVM thực hiện nhiệm vụ của một phương thức cụ thể đều có một thanh ghi bộ đếm chương trình được liên kết với nó. Phương thức không gốc có một PC lưu trữ địa chỉ của lệnh JVM có sẵn trong khi ở phương thức gốc, giá trị của bộ đếm chương trình không được xác định. Thanh ghi PC có khả năng lưu trữ địa chỉ trả về hoặc con trỏ gốc trên một số nền tảng cụ thể.

### Working of a Garbage Collector

JVM kích hoạt quá trình này và theo quy trình thu gom rác của JVM, quá trình này có thể được thực hiện hoặc tạm dừng. Điều này giảm bớt gánh nặng cho người lập trình bởi vì nó tự động thực hiện việc cấp phát hoặc giải phóng bộ nhớ.

Quá trình thu gom rác khiến cho các tiến trình hoặc luồng khác bị tạm dừng và do đó có tính chất tốn kém. Vấn đề này không được chấp nhận đối với khách hàng nhưng có thể được loại bỏ bằng cách áp dụng một số thuật toán dựa trên bộ thu gom rác. Quá trình này được gọi là "điều chỉnh bộ thu gom rác" và quan trọng để cải thiện hiệu suất của chương trình.

Một giải pháp khác là bộ thu gom rác thế hệ, thêm trường "tuổi" vào các đối tượng được gán bộ nhớ. Khi có nhiều đối tượng được tạo ra, danh sách rác tăng lên, từ đó tăng thời gian thu gom rác. Dựa trên số chu kỳ đồng hồ mà các đối tượng đã tồn tại, các đối tượng được nhóm lại và được gán một 'tuổi' tương ứng. Điều này giúp phân phối công việc thu gom rác.

	Hiểu cách chương trình và dữ liệu của nó được lưu trữ hoặc tổ chức là rất quan trọng vì nó giúp khi lập trình viên muốn viết mã tối ưu về mặt tài nguyên và tiêu thụ tài nguyên. Nó cũng giúp phát hiện rò rỉ bộ nhớ hoặc không nhất quán và giúp sửa lỗi liên quan đến bộ nhớ. Tuy nhiên, khái niệm quản lý bộ nhớ rất rộng lớn và do đó, người ta phải cố gắng học nó càng nhiều càng tốt để cải thiện kiến thức về chúng.

