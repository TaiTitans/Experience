
---
**Hãy bắt đầu với khái niệm về đối tượng.**  
Đối tượng là một thể hiện (instance) được tạo ra từ một lớp, đại diện cho một cá thể cụ thể trong một nhóm các sự vật thuộc lớp đó. Đối tượng có nhiều thuộc tính khác nhau và sở hữu một số hành vi cụ thể. Từ góc nhìn của máy tính, đối tượng là một khối bộ nhớ trong RAM, bao gói một số dữ liệu, tức là các thuộc tính được định nghĩa trong lớp. Vì vậy, đối tượng được sử dụng để đóng gói dữ liệu.

**Mảng trong Java** có một số đặc điểm cơ bản giống như các đối tượng khác trong Java. Chẳng hạn, bạn có thể đóng gói một số dữ liệu, truy cập các thuộc tính, và cũng có thể gọi các phương thức. Do đó, có thể nói rằng mảng là một đối tượng.

Chúng ta cũng có thể xác minh rằng mảng là một đối tượng thông qua mã code. Ví dụ, đoạn mã sau đây sẽ xuất ra java.lang.Object:
```java
Class clz = int[].class;
System.out.println(clz.getSuperclass().getName());
```
Từ đây, có thể thấy rằng lớp cha của lớp mảng là lớp Object. Do đó, ta có thể suy ra rằng mảng là một đối tượng.

