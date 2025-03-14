
---
### TreeSet là gì?

TreeSet là một lớp trong Java thuộc gói java.util, triển khai giao diện Set và dựa trên cấu trúc dữ liệu **cây đỏ-đen (Red-Black Tree)**. Nó được sử dụng để lưu trữ một tập hợp các phần tử không trùng lặp (unique) và tự động sắp xếp chúng theo thứ tự tăng dần (hoặc theo một thứ tự tùy chỉnh nếu bạn chỉ định).

### Đặc điểm chính của TreeSet

1. **Không trùng lặp**: Giống như mọi Set, TreeSet không cho phép các phần tử trùng lặp. Nếu bạn cố thêm một phần tử đã tồn tại, nó sẽ bị bỏ qua.
2. **Sắp xếp tự động**: Các phần tử trong TreeSet luôn được sắp xếp theo thứ tự tự nhiên (natural ordering) hoặc theo một Comparator mà bạn cung cấp.
3. **Hiệu suất**: Vì sử dụng cây đỏ-đen, các thao tác như thêm, xóa, và tìm kiếm có độ phức tạp trung bình là **O(log n)**.
4. **Không cho phép null** (tùy phiên bản): Từ Java 7 trở đi, TreeSet không cho phép thêm null nếu không có Comparator rõ ràng, vì nó cần so sánh các phần tử.

### Khi nào nên dùng TreeSet?

- Khi bạn cần một tập hợp không trùng lặp và luôn được sắp xếp.
- Khi bạn muốn truy cập nhanh các phần tử nhỏ nhất/lớn nhất hoặc kiểm tra thứ tự.

### Ví dụ cơ bản

Dưới đây là một ví dụ đơn giản để minh họa cách sử dụng TreeSet:
```java
import java.util.TreeSet;

public class TreeSetExample {
    public static void main(String[] args) {
        // Tạo một TreeSet
        TreeSet<Integer> numbers = new TreeSet<>();

        // Thêm phần tử
        numbers.add(5);
        numbers.add(2);
        numbers.add(8);
        numbers.add(1);
        numbers.add(5); // Phần tử trùng lặp, sẽ bị bỏ qua

        // In ra TreeSet (đã được sắp xếp)
        System.out.println("TreeSet: " + numbers); // Kết quả: [1, 2, 5, 8]

        // Một số thao tác hữu ích
        System.out.println("Phần tử nhỏ nhất: " + numbers.first()); // 1
        System.out.println("Phần tử lớn nhất: " + numbers.last());  // 8
        System.out.println("Có chứa 2 không? " + numbers.contains(2)); // true

        // Xóa phần tử
        numbers.remove(5);
        System.out.println("Sau khi xóa 5: " + numbers); // [1, 2, 8]
    }
}
```

### Các phương thức quan trọng của TreeSet

- add(E e): Thêm một phần tử vào TreeSet.
- remove(Object o): Xóa một phần tử.
- first(): Trả về phần tử nhỏ nhất.
- last(): Trả về phần tử lớn nhất.
- ceiling(E e): Trả về phần tử nhỏ nhất >= e.
- floor(E e): Trả về phần tử lớn nhất <= e.
- subSet(E from, E to): Trả về một tập con từ from đến to.
- pollFirst(): Lấy và loại bỏ phần tử nhỏ nhất.

### So sánh với HashSet

- HashSet: Không sắp xếp, nhanh hơn (O(1) cho thêm/xóa/tìm), cho phép null.
- TreeSet: Sắp xếp, chậm hơn (O(log n)), không cho phép null (trong hầu hết trường hợp).

### Lưu ý

- Các phần tử trong TreeSet phải triển khai giao diện Comparable (hoặc bạn phải cung cấp Comparator), nếu không sẽ ClassCastException.
- Không phù hợp nếu bạn cần truy cập ngẫu nhiên theo chỉ số (dùng ArrayList thay thế).