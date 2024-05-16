

---

	Java constructors or constructors in Java is a terminology used to construct something in our programs. A constructor in Java is a ****special method**** that is used to initialize objects. The constructor is called when an object of a class is created. It can be used to set initial values for object attributes.

In Java, a Constructor is a block of codes similar to the method. It is called when an instance of the class is created. At the time of calling the constructor, memory for the object is allocated in the memory. It is a special type of method that is used to initialize the object. Every time an object is created using the new() keyword, at least one constructor is called.

Dưới đây là một số điểm quan trọng về constructors trong Java:

1. **Tên và Kiểu Trả Về:** Constructors có cùng tên với tên lớp và không có kiểu trả về (không có từ khóa `void`, cũng không có kiểu dữ liệu).
    
2. **Sử Dụng:** Constructors được gọi khi một đối tượng mới được tạo ra bằng từ khóa `new`. Chúng đảm bảo rằng mọi trạng thái ban đầu của đối tượng được thiết lập đúng cách.
    
3. **Mặc Định và Tùy Chỉnh:** Một lớp có thể có constructor mặc định (không có tham số) và constructor tùy chỉnh (có tham số). Nếu không có constructor nào được khai báo, Java sẽ tạo một constructor mặc định mà không có tham số.
    
4. **Đa Nạp Chồng (Overloading):** Như các phương thức, constructors cũng có thể được đa nạp chồng (overloaded), có nghĩa là một lớp có thể có nhiều constructor với các danh sách đối số khác nhau.
    
5. **Gọi Constructor Khác:** Trong constructor của một lớp, bạn có thể gọi constructor khác của cùng lớp bằng từ khóa `this`. Ví dụ: `this(đối số)`.

VD:

```Java
public class Car {
    String brand;
    int year;

    // Constructor mặc định
    public Car() {
        brand = "Unknown";
        year = 2020;
    }

    // Constructor tùy chỉnh
    public Car(String brand, int year) {
        this.brand = brand;
        this.year = year;
    }

    public static void main(String[] args) {
        // Tạo một đối tượng Car với constructor mặc định
        Car car1 = new Car();
        System.out.println("Car 1: " + car1.brand + " " + car1.year);

        // Tạo một đối tượng Car với constructor tùy chỉnh
        Car car2 = new Car("Toyota", 2022);
        System.out.println("Car 2: " + car2.brand + " " + car2.year);
    }
}

```



[Java Constructors - GeeksforGeeks](https://www.geeksforgeeks.org/constructors-in-java/?ref=lbp)