
---

## 🔥 **1️⃣ Dynamic Binding (Ràng buộc động)**

🔹 **Dynamic Binding** là quá trình quyết định phương thức nào sẽ được gọi tại **runtime (thời gian chạy)** thay vì **compile-time (thời gian biên dịch)**. Điều này xảy ra khi bạn sử dụng **method overriding**.

📌 **Ví dụ:**
```java
class Parent {
    void show() {
        System.out.println("Parent class");
    }
}

class Child extends Parent {
    @Override
    void show() {
        System.out.println("Child class");
    }
}

public class Main {
    public static void main(String[] args) {
        Parent obj = new Child(); // Upcasting
        obj.show();  // Output: Child class (ràng buộc động)
    }
}
```
**📌 Giải thích:**

- Biến `obj` có kiểu `Parent`, nhưng vì nó tham chiếu đến một **đối tượng của `Child`**, nên phương thức `show()` của `Child` được gọi thay vì `Parent`.
- Quyết định này chỉ có thể được thực hiện **khi chạy chương trình** → **Dynamic Binding**.

🔹 **Khi nào dùng?**  
✔ Khi muốn **đa hình động (Runtime Polymorphism)**.  
✔ Khi muốn viết mã dễ mở rộng mà không cần sửa đổi nhiều.


---
## 🔥 **2️⃣ Message Passing (Truyền thông điệp giữa các đối tượng)**

🔹 **Message Passing** trong OOP là cách **các đối tượng giao tiếp với nhau** thông qua việc gọi phương thức (method invocation).  
🔹 Mỗi lần một đối tượng gọi một phương thức của đối tượng khác, nó đang **truyền thông điệp**.

📌 **Ví dụ thực tế:**  
🚗 Một chiếc xe hơi (`Car`) có thể nhận thông điệp từ người lái (`Driver`) để thực hiện hành động như **chạy, dừng, tăng tốc**.

📌 **Ví dụ trong Java:**
```java
class Car {
    void start() {
        System.out.println("Xe đang khởi động...");
    }

    void accelerate() {
        System.out.println("Xe đang tăng tốc...");
    }
}

public class Main {
    public static void main(String[] args) {
        Car myCar = new Car();
        myCar.start();      // Gửi thông điệp "start" đến đối tượng Car
        myCar.accelerate(); // Gửi thông điệp "accelerate" đến đối tượng Car
    }
}
```
🔹 **Khi nào dùng?**  
✔ Khi các đối tượng cần **tương tác với nhau**.  
✔ Khi muốn tăng **tính module hóa và mở rộng của chương trình**.

---
## 🔥 **3️⃣ Object Cloning (Nhân bản đối tượng)**

🔹 **Object Cloning** là quá trình tạo một **bản sao (copy)** của một đối tượng.  
🔹 Java hỗ trợ nhân bản đối tượng bằng cách sử dụng **`clone()`** từ interface `Cloneable`.

📌 **Ví dụ:**
```java
class Person implements Cloneable {
    String name;

    Person(String name) {
        this.name = name;
    }

    @Override
    protected Object clone() throws CloneNotSupportedException {
        return super.clone();
    }
}

public class Main {
    public static void main(String[] args) {
        try {
            Person p1 = new Person("Alice");
            Person p2 = (Person) p1.clone();  // Tạo bản sao p1

            System.out.println(p1.name);  // Output: Alice
            System.out.println(p2.name);  // Output: Alice
        } catch (CloneNotSupportedException e) {
            e.printStackTrace();
        }
    }
}
```

🔹 **Có hai loại cloning:**  
✔ **Shallow Copy** (Sao chép nông): Copy tham chiếu đến đối tượng con (không copy nội dung).  
✔ **Deep Copy** (Sao chép sâu): Copy toàn bộ nội dung, bao gồm cả đối tượng con.

🔹 **Khi nào dùng?**  
✔ Khi cần tạo **nhiều bản sao của đối tượng mà không thay đổi giá trị gốc**.  
✔ Khi muốn sao chép đối tượng mà không cần khởi tạo lại.

## 🎯 **Tóm tắt**

|Khái niệm|Mô tả|Khi nào dùng?|
|---|---|---|
|**Dynamic Binding**|Quyết định phương thức nào sẽ gọi tại runtime (hỗ trợ overriding)|Khi cần đa hình động (Runtime Polymorphism)|
|**Message Passing**|Các đối tượng giao tiếp với nhau thông qua phương thức|Khi các đối tượng cần tương tác với nhau|
|**Object Cloning**|Tạo một bản sao của đối tượng|Khi cần sao chép nhanh đối tượng mà không khởi tạo lại|
|**Wrapper Class**|Biến kiểu nguyên thủy thành Object (`Integer`, `Double`, ...)|Khi cần làm việc với Collections hoặc APIs yêu cầu Object|
