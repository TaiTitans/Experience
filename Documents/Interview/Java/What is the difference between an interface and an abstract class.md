
---
### **1. Sự khác biệt về ngữ pháp**

- **Lớp trừu tượng (Abstract Class)** có thể chứa các triển khai phương thức, trong khi các phương thức trong **giao diện (Interface)** chỉ có thể là phương thức trừu tượng (từ Java 8 trở đi, giao diện có thể có triển khai mặc định - default implementation).
- Các biến thành viên trong lớp trừu tượng có thể thuộc nhiều kiểu khác nhau, còn các biến thành viên trong giao diện chỉ có thể là kiểu public static final (hằng số công khai).
- Giao diện không thể chứa khối tĩnh (static block) hoặc phương thức tĩnh (trước Java 8), trong khi lớp trừu tượng có thể chứa cả khối tĩnh và phương thức tĩnh (từ Java 8, giao diện cũng có thể có phương thức tĩnh).
- Một lớp chỉ có thể kế thừa **một lớp trừu tượng**, nhưng một lớp có thể triển khai **nhiều giao diện**.

### **2. Sự khác biệt ở cấp độ thiết kế**

- **Mức độ trừu tượng khác nhau**:
    - Lớp trừu tượng trừu tượng hóa toàn bộ lớp, bao gồm cả thuộc tính và hành vi.
    - Giao diện chỉ trừu tượng hóa hành vi của lớp.
- Kế thừa lớp trừu tượng thể hiện mối quan hệ kiểu **"là một" (is-a)**, trong khi triển khai giao diện thể hiện mối quan hệ kiểu **"có thể làm" (can-do)**.
    - Nếu một lớp kế thừa một lớp trừu tượng, lớp con phải là một loại của lớp trừu tượng đó (ví dụ: một lớp con phải thuộc về loại của lớp cha).
    - Triển khai giao diện thì không cần mối quan hệ này, mà chỉ cần xem lớp có khả năng thực hiện hành vi đó hay không (ví dụ: một con chim có thể bay hay không).
- Các lớp kế thừa từ một lớp trừu tượng thường có đặc điểm tương tự nhau, trong khi các lớp triển khai cùng một giao diện có thể hoàn toàn khác nhau về bản chất.
Ví dụ về cửa và chuông báo động:
```java
class AlarmDoor extends Door implements Alarm {
    // Mã nguồn
}

class BMWCar extends Car implements Alarm {
    // Mã nguồn
}
```

- AlarmDoor vừa là một loại cửa (Door) vừa có khả năng báo động (Alarm).
- BMWCar vừa là một loại xe (Car) vừa có khả năng báo động (Alarm).
- AlarmDoor và BMWCar không liên quan trực tiếp, nhưng đều triển khai giao diện Alarm, thể hiện khả năng chung.

### Giải thích bổ sung:

- **Lớp trừu tượng**: Thích hợp khi bạn muốn chia sẻ mã nguồn và định nghĩa một cấu trúc chung cho các lớp con (ví dụ: Door có thuộc tính height và phương thức open()).
- **Giao diện**: Thích hợp khi bạn chỉ muốn định nghĩa hành vi mà không quan tâm đến chi tiết triển khai hoặc mối quan hệ "là một" (ví dụ: Alarm chỉ quan tâm đến khả năng ring()).
- Từ Java 8, giao diện trở nên linh hoạt hơn với default và static methods, nhưng vẫn khác biệt về mục đích thiết kế so với lớp trừu tượng.