
---


Tính đóng gói cho phép che giấu thông tin và những tính chất xử lý bên trong của đối tượng. Các đối tượng khác không thể tác động trực tiếp đến dữ liệu bên trong và làm thay đổi trạng thái của đối tượng mà bắt buộc phải thông qua các phương thức công khai do đối tượng đó cung cấp.

Tính chất này giúp tăng tính bảo mật cho đối tượng và tránh tình trạng dữ liệu bị hư hỏng ngoài ý muốn.


---
	In Java, encapsulation is achieved by declaring the instance variables of a class as private, which means they can only be accessed within the class. To allow outside access to the instance variables, public methods called getters and setters are defined, which are used to retrieve and modify the values of the instance variables, respectively. By using getters and setters, the class can enforce its own data validation rules and ensure that its internal state remains consistent.

VN:
	Nói dễ hiểu tính đóng gói khai báo class một cách riêng tư và chỉ có thể truy cập bên trong nó. Nếu muốn truy cập từ bên ngoài thì cần khai báo các method được gọi như setter và getter nhưng cần đảm bảo rằng trạng thái bên trong vẫn nhất quán.

![](https://media.geeksforgeeks.org/wp-content/uploads/Encapsulation.jpg)

---
Below is the example with Java Encapsulation:

```Java
class Person{
private String name;
private int age;
public String getName(){return name;}
public void setName(String name){
this.name = name;
}
public int getAge(){return age;}
public void setAge(int age){
this.age=age;
}
}

public class Main{
public static void main(String[] args){
Person person = new Person();
person.setName("Titans");
person.setAge(22);
}
}
```


	****Encapsulation**** is defined as the wrapping up of data under a single unit. It is the mechanism that binds together code and the data it manipulates. Another way to think about encapsulation is, that it is a protective shield that prevents the data from being accessed by the code outside this shield.

VN:
	Đóng gọi được định nghĩa là việc gói dữ liệu dưới một đơn vị duy nhất. Một cách nghĩ dễ hiểu là một lá chắn bảo vệ ngăn chặn dữ liệu bị truy cập bởi bên ngoài tấm lá chắn này.

- Technically in encapsulation, the variables or data of a class is hidden from any other class and can be accessed only through any member function of its own class in which it is declared.
-  Đồng thời, lớp vẫn được tiết lộ ra bên ngoài để người dùng cuối có thể sử dụng mà không cần biết chi tiết cài đặt bên trong, điều này liên quan đến khái niệm **abstraction** (trừu tượng). Vì vậy, encapsulation thường được gọi là sự kết hợp giữa data-hiding và abstraction.
- Encapsulation can be achieved by Declaring all the variables in the class as private and writing public methods in the class to set and get the values of variables.
## ***Advantages of Encapsulation***

- **Data Hiding***
- **Increased Flexibility**
- **Reusability**: We can make the variables of the class read-only or write-only depending on our requirements. If we wish to make the variables read-only then we have to omit the setter methods like setName(), setAge(), etc. from the above program or if we wish to make the variables write-only then we have to omit the get methods like getName(), getAge(), etc. from the above program.
- **Reusability***
- ***Testing code is easy***
- ***Freedom to programmer in implementing the details of the system***

### Disadvantages of Encapsulation in Java

- Can lead to increased complexity, especially if not used properly.
- Can make it more difficult to understand how the system works.
- May limit the flexibility of the implementation.