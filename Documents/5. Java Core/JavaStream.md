
---
	Trong Java, Stream là một API được giới thiệu từ phiên bản Java 8 để xử lý dữ liệu theo cách linh hoạt, hiệu quả và dễ đọc. Stream không phải là một cấu trúc dữ liệu để lưu trữ dữ liệu, mà nó là một chuỗi các phương thức cho phép thực hiện các thao tác xử lý dữ liệu như lọc, ánh xạ, giới hạn, sắp xếp, và gom nhóm trên các tập dữ liệu.

Dưới đây là một số điểm quan trọng để hiểu về Stream trong Java:

1. **Lười Biếng (Lazy Evaluation)**: Một tính năng quan trọng của Stream là nó thực hiện các phương thức trên dữ liệu một cách lười biếng. Điều này có nghĩa là các phép biến đổi hoặc thao tác được thực hiện chỉ khi kết quả thực sự cần thiết, giúp tối ưu hóa hiệu suất và tiết kiệm tài nguyên.
    
2. **Luồng Tuần Tự**: Stream trong Java có thể được xem như một luồng tuần tự các phần tử dữ liệu. Bạn có thể thực hiện các phép biến đổi trên Stream để xử lý từng phần tử một cách tuần tự mà không cần quan tâm đến bản chất của dữ liệu.
    
3. **Phương Thức Trung Gian và Kết Thúc**: Stream thực hiện các phép biến đổi thông qua phương thức trung gian như `filter`, `map`, `sorted`, `distinct`, và nhiều phương thức khác. Cuối cùng, các phương thức kết thúc như `collect`, `forEach`, `reduce` được gọi để lấy kết quả cuối cùng sau khi thực hiện các phép biến đổi.
    
4. **Không Thay Đổi Dữ Liệu Gốc**: Một Stream không thay đổi dữ liệu gốc mà nó được tạo ra từ. Thay vào đó, nó tạo ra một chuỗi các bước biến đổi mà bạn có thể áp dụng lên dữ liệu gốc mà không làm thay đổi dữ liệu đó.
    
5. **Parallel Streams**: Java cũng hỗ trợ Parallel Streams, cho phép xử lý các phần tử của Stream một cách song song trên nhiều luồng, tăng hiệu suất xử lý đối với các tập dữ liệu lớn.


Ví dụ :

```
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class StreamExample {
    public static void main(String[] args) {
        List<Integer> numbers = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);

        // Filter even numbers
        List<Integer> evenNumbers = numbers.stream()
                                           .filter(num -> num % 2 == 0)
                                           .collect(Collectors.toList());

        System.out.println("Even numbers: " + evenNumbers);

        // Map and collect
        List<Integer> squares = numbers.stream()
                                       .map(num -> num * num)
                                       .collect(Collectors.toList());

        System.out.println("Squares: " + squares);

        // Reduce to get sum
        int sum = numbers.stream()
                         .reduce(0, Integer::sum);

        System.out.println("Sum: " + sum);
    }
}

```

