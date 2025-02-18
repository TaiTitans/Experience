
---
`Đa hình`

## **1. Đa Hình là gì?**

Đa hình (**Polymorphism**) là một **tính chất trong OOP** cho phép cùng một phương thức hoặc toán tử có thể hoạt động theo nhiều cách khác nhau tùy vào **kiểu dữ liệu** hoặc **đối tượng** gọi nó.

📌 **Lợi ích chính của đa hình:**

- **Tái sử dụng mã nguồn** hiệu quả
- **Tăng tính linh hoạt** của chương trình
- **Dễ mở rộng và bảo trì hơn**

📌 **Ví dụ thực tế:**

- Một **người** có thể vừa là **sinh viên, nhân viên hoặc phụ huynh**, nhưng vẫn là **cùng một con người**.

---

## **2. Các loại Đa Hình**

Có **hai loại chính** của đa hình trong OOP:

1. **Đa hình tĩnh (Compile-time Polymorphism)**
    
    - Thực hiện tại **thời điểm biên dịch**
    - Được triển khai thông qua **nạp chồng hàm (Function Overloading)** và **nạp chồng toán tử (Operator Overloading)**
2. **Đa hình động (Runtime Polymorphism)**
    
    - Xảy ra tại **thời điểm chạy chương trình**
    - Được triển khai thông qua **ghi đè phương thức (Method Overriding)**

## **3. Nạp chồng phương thức và toán tử**

### **Nạp chồng phương thức (Function Overloading)**

- Cùng tên phương thức nhưng khác **số lượng hoặc kiểu dữ liệu tham số**.

🔹 **Ví dụ:**
```java
class Demo {
    void show(int a) {
        System.out.println("Số nguyên: " + a);
    }

    void show(double a) {
        System.out.println("Số thực: " + a);
    }
}

public class Main {
    public static void main(String[] args) {
        Demo obj = new Demo();
        obj.show(5);       // Gọi hàm với tham số int
        obj.show(5.5);     // Gọi hàm với tham số double
    }
}
```

📌 **Lưu ý:**  
🔹 **Có thể nạp chồng phương thức** khi chỉ thay đổi kiểu hoặc số lượng tham số.  
🔹 **Không thể nạp chồng phương thức** chỉ bằng cách thay đổi kiểu trả về.

## **4. Ghi đè phương thức (Method Overriding)**

- Khi một **lớp con** thay đổi hành vi của phương thức từ **lớp cha**.

🔹 **Ví dụ:**
```java
class Animal {
    void sound() {
        System.out.println("Động vật phát ra âm thanh");
    }
}

class Dog extends Animal {
    void sound() {
        System.out.println("Chó sủa: Gâu Gâu!");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal a = new Dog();
        a.sound();  // Gọi phương thức từ lớp con (Dog)
    }
}
```
📌 **Lưu ý:**  
🔹 Khi ghi đè, phương thức phải có cùng **tên, kiểu trả về và danh sách tham số**.  
🔹 **Ghi đè phương thức chỉ xảy ra với kế thừa (inheritance)**.

## **5. Hàm ảo (Virtual Function) và Liên kết động (Dynamic Binding)**

- **Virtual Function** là các phương thức có thể được ghi đè trong lớp con.
- **Dynamic Binding** giúp quyết định phương thức nào được gọi tại **thời điểm chạy** thay vì **thời điểm biên dịch**

📌 **Lưu ý:**  
🔹 Nếu không dùng **Dynamic Binding**, chương trình sẽ gọi phương thức của lớp cha thay vì lớp con.

## **6. Lớp trừu tượng (Abstract Class) và Phương thức ảo thuần túy (Pure Virtual Function)**

### **Lớp trừu tượng (Abstract Class)**

- Là lớp **không thể khởi tạo trực tiếp**.
- Chứa phương thức **trừu tượng** (không có phần thân).

🔹 **Ví dụ:**
```java
abstract class Vehicle {
    abstract void start();  // Phương thức trừu tượng
}

class Car extends Vehicle {
    void start() {
        System.out.println("Xe hơi khởi động bằng chìa khóa");
    }
}

public class Main {
    public static void main(String[] args) {
        Vehicle myCar = new Car();
        myCar.start();
    }
}
```
📌 **Lưu ý:**  
🔹 **Lớp trừu tượng có thể có cả phương thức bình thường và phương thức trừu tượng**.  
🔹 **Lớp con bắt buộc phải triển khai các phương thức trừu tượng**.

### **Phương thức ảo thuần túy (Pure Virtual Function)**

- Trong Java, mọi phương thức trừu tượng đều là **pure virtual function** (giống như trong C++).


## **7. Toán tử `instanceof`**

- Kiểm tra xem một đối tượng có thuộc về một lớp cụ thể hay không.

🔹 **Ví dụ:**
```java
class Animal { }

class Dog extends Animal { }

public class Main {
    public static void main(String[] args) {
        Dog d = new Dog();
        System.out.println(d instanceof Dog);     // true
        System.out.println(d instanceof Animal);  // true
    }
}
```
📌 **Lưu ý:**  
🔹 Nếu `d` không phải là thể hiện của lớp `Dog`, nó sẽ trả về `false`.

## **8. So sánh Đa hình tại thời gian biên dịch vs Đa hình tại thời gian chạy**

|Đặc điểm|Đa hình tĩnh (Compile-time)|Đa hình động (Runtime)|
|---|---|---|
|**Cách triển khai**|Nạp chồng phương thức/toán tử|Ghi đè phương thức|
|**Thời điểm quyết định**|Lúc biên dịch|Lúc chạy chương trình|
|**Hiệu suất**|Nhanh hơn|Chậm hơn|
|**Dễ sử dụng**|Dễ triển khai hơn|Yêu cầu dùng kế thừa|

---

## **Tổng kết**

✅ **Đa hình giúp mã nguồn linh hoạt, dễ mở rộng và tái sử dụng hơn.**  
✅ **Đa hình có hai loại chính: Compile-time và Runtime.**  
✅ **Method Overloading giúp nạp chồng phương thức với tham số khác nhau.**  
✅ **Method Overriding giúp lớp con thay đổi hành vi của lớp cha.**  
✅ **Abstract Class và Pure Virtual Function giúp đảm bảo tính trừu tượng.**  
✅ **Virtual Function và Dynamic Binding cho phép quyết định phương thức tại thời gian chạy.**

---

## 📌 **1️⃣ `Abstract` là gì?**

- `Abstract` là một **tính chất của trừu tượng hóa** trong lập trình hướng đối tượng (OOP).
- **Mục đích:** Định nghĩa một khuôn mẫu (template) cho các lớp con kế thừa, nhưng **không thể tạo đối tượng từ lớp `abstract` trực tiếp**.

## **2️⃣ Lớp Trừu Tượng (`abstract class`)**

📌 **Lớp trừu tượng là lớp có thể chứa các phương thức trừu tượng (`abstract methods`).**  
📌 **Các lớp con bắt buộc phải override (ghi đè) các phương thức trừu tượng này.**

### **🔹 Ví dụ về Lớp Trừu Tượng**
```java
// Định nghĩa lớp trừu tượng
abstract class Animal {
    // Phương thức trừu tượng (không có thân hàm)
    abstract void makeSound();

    // Phương thức bình thường (có thân hàm)
    void sleep() {
        System.out.println("Động vật đang ngủ...");
    }
}

// Lớp Dog kế thừa từ Animal
class Dog extends Animal {
    @Override
    void makeSound() {
        System.out.println("Chó sủa: Gâu Gâu!");
    }
}

public class Main {
    public static void main(String[] args) {
        Dog myDog = new Dog();
        myDog.makeSound(); // Gọi phương thức đã ghi đè
        myDog.sleep(); // Gọi phương thức có sẵn trong Animal
    }
}

```

```
Chó sủa: Gâu Gâu!
Động vật đang ngủ...
```

✅ **Lớp `Dog` buộc phải override `makeSound()`, nhưng vẫn có thể sử dụng `sleep()`.**

## **3️⃣ Tại sao dùng `abstract class`?**

|**Đặc điểm**|**Có Lớp Trừu Tượng**|**Không Có Lớp Trừu Tượng**|
|---|---|---|
|**Tính linh hoạt**|✅ Dễ mở rộng|❌ Dễ bị trùng lặp mã nguồn|
|**Bắt buộc triển khai**|✅ Lớp con phải override phương thức abstract|❌ Không có ràng buộc|
|**Tính kế thừa**|✅ Tạo khuôn mẫu chung|❌ Dễ bị sai sót khi mở rộng|
|**Tạo đối tượng**|❌ Không thể tạo trực tiếp|✅ Tạo trực tiếp|

📌 **Dùng `abstract` khi cần đảm bảo tất cả lớp con có các phương thức chung nhưng có cách triển khai riêng biệt.**

## **4️⃣ Phương thức Trừu Tượng (`abstract method`)**

📌 **Là phương thức không có phần thân và chỉ định nghĩa trong lớp `abstract`.**  
📌 **Lớp con phải override nó.**

🔹 **Ví dụ:**
```java
abstract class Vehicle {
    abstract void start();  // Phương thức trừu tượng

    void stop() {
        System.out.println("Dừng xe");
    }
}

class Car extends Vehicle {
    @Override
    void start() {
        System.out.println("Xe hơi khởi động bằng chìa khóa");
    }
}
```
📌 **Không thể tạo `Vehicle vehicle = new Vehicle();` vì `Vehicle` là `abstract`.**

## **5️⃣ `abstract class` vs. `interface`**

|**Đặc điểm**|**Abstract Class**|**Interface**|
|---|---|---|
|**Mục đích**|Tạo khuôn mẫu cho các lớp con|Xác định hành vi chung cho nhiều lớp|
|**Hỗ trợ phương thức có thân?**|✅ Có thể có|❌ Không có (trước Java 8)|
|**Hỗ trợ biến instance?**|✅ Có|❌ Không|
|**Hỗ trợ đa kế thừa?**|❌ Không hỗ trợ|✅ Có thể implement nhiều interface|
|**Tốc độ thực thi**|⏳ Nhanh hơn|🐢 Chậm hơn do phải tìm phương thức|

📌 **Khi nào dùng `abstract class`?**  
✅ Khi có **thuộc tính chung** và **một số phương thức có thân**.

📌 **Khi nào dùng `interface`?**  
✅ Khi **muốn lớp có thể thực hiện nhiều hành vi khác nhau** mà không bị giới hạn bởi kế thừa đơn.

---

## **6️⃣ Lớp Trừu Tượng Có Constructor Không?**

❓ **Constructor có thể tồn tại trong lớp trừu tượng không?**  
✅ **Có!**  
📌 **Mặc dù không thể tạo đối tượng từ `abstract class`, nhưng constructor giúp khởi tạo thuộc tính cho lớp con.**

🔹 **Ví dụ:**
```java
abstract class Animal {
    String name;

    // Constructor trong abstract class
    Animal(String name) {
        this.name = name;
    }

    abstract void makeSound();
}

class Dog extends Animal {
    Dog(String name) {
        super(name);  // Gọi constructor của lớp cha
    }

    @Override
    void makeSound() {
        System.out.println(name + " sủa: Gâu Gâu!");
    }
}

public class Main {
    public static void main(String[] args) {
        Dog myDog = new Dog("Buddy");
        myDog.makeSound();
    }
}
```

✅ **Lớp con `Dog` kế thừa constructor từ `Animal` và có thể sử dụng nó.**

## **7️⃣ Lớp Trừu Tượng Có Biến `static` Không?**

✅ **Có thể có biến `static`, và có thể truy cập trực tiếp mà không cần tạo đối tượng.**

🔹 **Ví dụ:**
```java
abstract class Vehicle {
    static int totalVehicles = 0;

    Vehicle() {
        totalVehicles++;
    }

    static void showTotal() {
        System.out.println("Tổng số phương tiện: " + totalVehicles);
    }
}

class Car extends Vehicle {}

public class Main {
    public static void main(String[] args) {
        new Car();
        new Car();
        Vehicle.showTotal();
    }
}
```
✅ **`static` giúp lưu trạng thái chung cho tất cả các đối tượng kế thừa từ `Vehicle`.**


## **8️⃣ Tổng kết**

### ✅ **`abstract class` giúp tạo khuôn mẫu cho các lớp con.**

### ✅ **Phương thức `abstract` không có thân và bắt buộc lớp con phải override.**

### ✅ **Không thể tạo đối tượng từ `abstract class`, nhưng có thể có constructor và biến `static`.**

### ✅ **Khác với `interface`, `abstract class` có thể chứa cả phương thức có thân.**

### ✅ **Giúp tăng khả năng tái sử dụng mã nguồn và giảm trùng lặp.**

