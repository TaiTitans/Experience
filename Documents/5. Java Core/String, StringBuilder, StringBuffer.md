
---
Khi làm việc với dữ liệu kiểu text trong Java cung cấp 3 class String, StringBuffer và StringBuilder. Cơ bản về 3 class này như sau: String là không thể thay đổi (immutable), và không cho phép có class con. StringBuffer, StringBuilder có thể thay đổi (mutable) StringBuilder và StringBuffer là giống nhau, nó chỉ khác biệt tình huống sử dụng có liên quan tới đa luồng (Multi Thread). => về tốc độ xử lý StringBuilder là tốt nhất, sau đó StringBuffer và cuối cùng mới là String.


**String** Trong java, String là một class đặc biệt, nguyên nhân là nó được sử dụng một cách thường xuyên trong một chương trình, vì vậy đòi hỏi nó phải có hiệu suất và sự mềm dẻo. Đó là lý do tại sao String có tính đối tượng và vừa có tính nguyên thủy (primitive). Ví Dụ về tính nguyên thủy: Chúng ta hoàn toàn có thể khai báo: String s1 = "Hello word" Bạn cũng có thể sử dụng toán tử + để nối 2 string, toán tử này vốn quen thuộc và sử dụng cho các kiểu dữ liệu nguyên thủy như int, float, double. Tính đối tượng: Vì String là một class, vì vậy nó có thể được tạo ra thông qua toán tử new. `String object = new String("Hello World");` Các đối tượng String được lưu trữ trên Heap, yêu cầu quản lý bộ nhớ phức tạp và tốn không gian lưu trữ. Hai đối tượng String có nội dung giống nhau lưu trữ trên 2 vùng bộ nhớ khác nhau của Heap. Phương thức Equals vs == Phương thức equals() sử dụng để so sánh 2 đối tượng, với String nó có ý nghĩa là so sánh nội dung của 2 string. Đối với các kiểu tham chiếu (reference) toán tử == có ý nghĩa là so sánh địa chỉ vùng bộ nhớ lưu trữ của đối tượng.

```Java
String s1 = "Hello";
String s2 = "Hello"; 
String s3 = s1;  
String s4 = new String("Hello");  
String s5 = new String("Hello"); 
 
s1 == s1;         // true
s1 == s2;         // true
s1 == s3;         // true
s1 == s4;         // false
s4 == s5;         // false
 
s1.equals(s3);    // true, cùng nội dung
s1.equals(s4);    // true, cùng nội dung
s4.equals(s5);    // true, cùng nội dung.
```

**StringBuffer vs StringBuilder** StringBuilder và StringBuffer là khá giống nhau, điều khác biệt là tất cả các phương thức của **StringBuffer đã được đồng bộ, nó thích hợp khi bạn làm việc với ứng dụng đa luồng, nhiều luồng có thể truy cập vào một đối tượng StringBuffer cùng lúc.** Trong khi đó **StringBuilder có các phương thức tương tự nhưng không được đồng bộ, vì vậy mà hiệu suất của nó cao hơn, bạn nên sử dụng StringBuilder trong ứng dụng đơn luồng,** hoặc sử dụng như một biến địa phương trong một phương thức.
**String Buffer**
```Java
public class BufferString {
    public static void main(String[] args) {
        StringBuffer buffer = new StringBuffer("Hello");
        buffer.append("Word");
        System.out.println(buffer);
    }
}
```
**String Builder**
```Java
public class BuiderString {
    public static void main(String[] args) {
        StringBuffer buider = new StringBuffer("Hello");
        buffer.append("Word");
        System.out.println(buider);
    }
}
```

