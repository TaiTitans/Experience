
---
**Trong Java, các lớp không hỗ trợ đa kế thừa (multiple inheritance).**

- Tuy nhiên, các giao diện (API) hỗ trợ đa kế thừa.
- Mục đích của giao diện là mở rộng chức năng của một đối tượng. Nếu một giao diện con kế thừa nhiều giao diện cha, giao diện con sẽ mở rộng nhiều chức năng. Khi một lớp triển khai giao diện đó, nó cũng sẽ được mở rộng nhiều chức năng.

**Lý do Java không hỗ trợ đa kế thừa cho các lớp:**

- Vì lý do bảo mật: Nếu một lớp con kế thừa từ nhiều lớp cha mà các lớp cha này có cùng phương thức hoặc thuộc tính, lớp con sẽ không biết nên kế thừa từ lớp nào. Điều này dẫn đến xung đột (diamond problem).
- Java cung cấp **giao diện** và **lớp bên trong (inner class)** để đạt được chức năng tương tự đa kế thừa, bù đắp cho những hạn chế của đơn kế thừa.

### Giải thích bổ sung:

1. **Đa kế thừa và vấn đề kim cương (Diamond Problem)**:
    - Trong các ngôn ngữ như C++, nếu lớp D kế thừa từ hai lớp B và C (mà cả B và C đều kế thừa từ A và có cùng phương thức), sẽ gây ra sự mơ hồ khi D gọi phương thức đó. Java tránh vấn đề này bằng cách không cho phép đa kế thừa lớp.
    - Ví dụ: Nếu B và C đều có phương thức doSomething(), D không biết nên dùng phiên bản nào.
2. **Giải pháp của Java**:
    - **Giao diện**: Một lớp có thể triển khai nhiều giao diện (ví dụ: class MyClass implements Interface1, Interface2), cho phép mở rộng nhiều hành vi mà không xung đột cấu trúc. Từ Java 8, giao diện còn hỗ trợ phương thức mặc định (default methods) để cung cấp triển khai.
    - **Lớp bên trong**: Có thể sử dụng để mô phỏng một số khía cạnh của đa kế thừa bằng cách nhúng các lớp khác bên trong.