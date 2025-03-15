
---
Generics là một tính năng được giới thiệu từ **JDK 5** giúp sử dụng **tham số kiểu (type parameters)** khi định nghĩa **lớp, interface, và phương thức**. Điều này giúp **tái sử dụng code**, **tăng tính an toàn kiểu dữ liệu**, và **giảm lỗi runtime**.

---

## **1. Lợi ích của Generics**

✅ **Tăng tính an toàn kiểu dữ liệu (Type Safety)**:

- Giúp phát hiện lỗi kiểu dữ liệu **ngay tại thời điểm biên dịch** thay vì runtime.

✅ **Tái sử dụng mã nguồn (Code Reusability)**:

- Cho phép tạo ra **các lớp và phương thức chung** có thể làm việc với nhiều kiểu dữ liệu khác nhau.

✅ **Giảm việc sử dụng ép kiểu (Casting)**:

- Tránh các ép kiểu **tốn kém và dễ gây lỗi** khi sử dụng kiểu `Object`.

---
```java
import java.util.ArrayList;
import java.util.List;

public class GenericExample {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>(); // Chỉ cho phép String
        list.add("Hello");
        // list.add(100);  // Lỗi biên dịch nếu bỏ comment

        String str = list.get(0);  // Không cần ép kiểu
        System.out.println(str);
    }
}
```

