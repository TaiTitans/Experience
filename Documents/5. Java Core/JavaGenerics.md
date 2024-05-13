
---
	Các phương thức Java Generic và các lớp chung cho phép người lập trình chỉ định, với một khai báo phương thức duy nhất, một tập hợp các phương thức liên quan hoặc với một khai báo lớp duy nhất, một tập hợp các kiểu liên quan tương ứng.

=> Generics trong Java cho phép bạn tạo ra các lớp, giao diện và phương thức mà hoạt động trên một loại dữ liệu cụ thể mà không cần chỉ định loại dữ liệu đó cho đến khi mã được sử dụng. Điều này giúp đảm bảo tính an toàn về kiểu dữ liệu và cho phép bạn viết mã linh hoạt và có thể tái sử dụng hơn.

Dưới đây là một số điểm quan trọng về generics trong Java:

1. **An Toàn Kiểu Dữ Liệu**: Generics đảm bảo bạn làm việc với một loại dữ liệu cụ thể, giảm nguy cơ các lỗi thời gian chạy như ClassCastException. Ví dụ, một `List<String>` chỉ có thể chứa chuỗi và cố gắng thêm bất cứ điều gì khác sẽ dẫn đến lỗi biên dịch.
    
2. **Tính Tái Sử Dụng Mã**: Với generics, bạn có thể tạo ra các lớp, giao diện và phương thức có thể hoạt động với bất kỳ loại dữ liệu nào, giúp mã của bạn trở nên tổng quát và có thể tái sử dụng hơn. Ví dụ, một lớp generic `Pair<T, U>` có thể chứa một cặp các đối tượng thuộc hai loại bất kỳ.
    
3. **Kiểm Tra Thời Gian Biên Dịch**: Generics cung cấp kiểm tra thời gian biên dịch, điều này có nghĩa là các lỗi kiểu dữ liệu được phát hiện trong quá trình biên dịch thay vì thời gian chạy. Điều này giúp phát hiện vấn đề sớm trong quá trình phát triển.
    
4. **Sự Gợi Ý Kiểu Dữ Liệu**: Sự gợi ý kiểu dữ liệu của Java cho phép bạn bỏ qua các tham số kiểu dữ liệu của một phương thức hay hàm tạo generic khi nó có thể được suy luận từ ngữ cảnh. Điều này cải thiện khả năng đọc mã và giảm sự dài dòng.
    
5. **Wildcard**: Generics hỗ trợ các wildcard (`?`) cho phép bạn làm việc với các loại dữ liệu không xác định hoặc chỉ định một số ràng buộc cụ thể. Ví dụ, `List<?>` đại diện cho một danh sách các đối tượng của một loại không xác định, trong khi `List<? extends Number>` đại diện cho một danh sách các đối tượng là các lớp con của `Number`.
    
6. **Phương Thức Generic**: Ngoài các lớp và giao diện generic, Java cũng cho phép bạn tạo ra các phương thức generic trong các lớp không generic. Điều này hữu ích khi bạn muốn một phương thức làm việc với các loại dữ liệu khác nhau tùy thuộc vào đối số được truyền vào.

- T – Type
- E – Element
- K – Key
- N – Number
- V – Value

```
public class Box<T> {
    private T value;

    public void setValue(T value) {
        this.value = value;
    }

    public T getValue() {
        return value;
    }

    public static void main(String[] args) {
        Box<Integer> intBox = new Box<>();
        intBox.setValue(10);
        System.out.println("Value: " + intBox.getValue());

        Box<String> strBox = new Box<>();
        strBox.setValue("Xin chào, Generics!");
        System.out.println("Value: " + strBox.getValue());
    }
}
```


Trong ví dụ này, `Box<T>` là một lớp generic có thể chứa các giá trị của bất kỳ loại `T` nào. Loại dữ liệu được chỉ định khi tạo các thể hiện của `Box`, như `Box<Integer>` hoặc `Box<String>`.