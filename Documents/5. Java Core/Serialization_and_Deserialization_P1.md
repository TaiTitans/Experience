
---
	Serialization là quá trình chuyển đổi trạng thái của một đối tượng thành một luồng byte; deserialization là quá trình ngược lại. Nói cách khác, serialization là quá trình chuyển đổi một đối tượng Java thành một luồng (trình tự) tĩnh của byte, mà sau đó chúng ta có thể lưu vào cơ sở dữ liệu hoặc truyền qua mạng.



![](../../Assets/Images/JavaCore/serialize-deserialize-java.png)
**Ưu điểm của Serialization trong java**

	Nó chủ yếu được sử dụng để truyền trạng thái của đối tượng qua mạng 

**java.io.Serializable interface**

**Serializable** là một interface (giao diện) đánh dấu, không có thuộc tính và phương thức bên trong. Nó được sử dụng để “**đánh dấu**” các lớp java để các đối tượng của các lớp này có thể nhận được khả năng nhất định. **Cloneable** và **Remote** cũng là những interface đánh dấu.

Nó phải được implements bởi lớp mà đối tượng của nó bạn muốn persist.

## Deserialization trong java

Deserialization là quá trình tái thiết lại các đối tượng từ trạng thái serialized. Đây là hoạt động ngược lại của serialization.