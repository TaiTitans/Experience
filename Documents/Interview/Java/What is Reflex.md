
---
**Cơ chế phản xạ (Reflection) trong Java** cho phép chương trình **thu thập thông tin động** và **gọi phương thức động** của một đối tượng trong thời gian chạy.

### **Đặc điểm của Reflection trong Java**:

- Cho phép **truy xuất thông tin** của một lớp (class), bao gồm các thuộc tính, phương thức, constructor, modifier, v.v.
- Cho phép **tạo đối tượng** một cách động mà không cần biết lớp cụ thể tại thời điểm biên dịch.
- Cho phép **gọi phương thức của đối tượng** ngay cả khi không biết tên phương thức trước đó.

**Ví dụ về sử dụng Reflection trong Java**:
```java
import java.lang.reflect.Method;
import java.lang.reflect.Field;
import java.lang.reflect.Constructor;

class Person {
    private String name;

    public Person() {}

    public Person(String name) {
        this.name = name;
    }

    private void sayHello() {
        System.out.println("Hello, my name is " + name);
    }
}

public class ReflectionExample {
    public static void main(String[] args) throws Exception {
        // Lấy thông tin lớp Person
        Class<?> clazz = Class.forName("Person");

        // Tạo một đối tượng Person bằng Reflection
        Constructor<?> constructor = clazz.getDeclaredConstructor(String.class);
        Object person = constructor.newInstance("John Doe");

        // Lấy thông tin trường name và thay đổi giá trị của nó
        Field field = clazz.getDeclaredField("name");
        field.setAccessible(true); // Bỏ giới hạn truy cập private
        field.set(person, "Alice");

        // Lấy thông tin phương thức private sayHello() và gọi nó
        Method method = clazz.getDeclaredMethod("sayHello");
        method.setAccessible(true);
        method.invoke(person);
    }
}
```

### **Ứng dụng của Reflection**:

- **Frameworks như Spring, Hibernate** sử dụng Reflection để khởi tạo bean, inject dependencies và ánh xạ dữ liệu.
- **Giao tiếp với cơ sở dữ liệu**: ORM frameworks như Hibernate sử dụng Reflection để ánh xạ bảng database với entity class.
- **Xây dựng công cụ phát triển**: IDE như IntelliJ hoặc Eclipse sử dụng Reflection để cung cấp gợi ý mã nguồn.
- **Lập trình động**: Reflection giúp tải các lớp và phương thức trong runtime, hữu ích cho plugin hoặc hệ thống mở rộng.

⚠ **Lưu ý**: Reflection có thể gây **hiệu suất kém** và **vi phạm tính đóng gói** của OOP, nên cần sử dụng một cách hợp lý.