
---
## 📌 **1️⃣ Kế Thừa (Inheritance) là gì?**

- **Kế thừa (Inheritance)** là một tính năng của lập trình hướng đối tượng (OOP) cho phép một lớp **mới kế thừa thuộc tính và phương thức** từ một lớp khác.
- Lớp con **(subclass/child class)** có thể mở rộng hoặc ghi đè phương thức của lớp cha **(superclass/parent class)**.

📌 **Tính chất quan trọng:**  
✅ **Tái sử dụng mã nguồn** (Reusability)  
✅ **Mở rộng chức năng mà không sửa đổi lớp cũ**  
✅ **Dễ dàng bảo trì và nâng cấp**

🔹 **Ví dụ đơn giản về kế thừa trong Java**:
```JAVA
// Lớp cha (Superclass)
class Animal {
    void eat() {
        System.out.println("Động vật đang ăn...");
    }
}

// Lớp con (Subclass)
class Dog extends Animal {
    void bark() {
        System.out.println("Chó sủa: Gâu Gâu!");
    }
}

public class Main {
    public static void main(String[] args) {
        Dog myDog = new Dog();
        myDog.eat();  // Gọi phương thức từ lớp cha
        myDog.bark(); // Gọi phương thức của lớp con
    }
}

```

```
Động vật đang ăn...
Chó sủa: Gâu Gâu!
```

✅ **Lớp `Dog` có thể sử dụng `eat()` từ `Animal` mà không cần định nghĩa lại.**

## **2️⃣ Các loại kế thừa trong Java**

📌 **Java hỗ trợ các loại kế thừa sau:**

|Loại Kế Thừa|Mô Tả|
|---|---|
|**Single Inheritance**|Một lớp kế thừa từ một lớp khác|
|**Multilevel Inheritance**|Một lớp kế thừa từ một lớp kế thừa khác|
|**Hierarchical Inheritance**|Một lớp cha có nhiều lớp con|
|**Hybrid Inheritance**|Kết hợp nhiều loại kế thừa|
|**Multiple Inheritance (Không hỗ trợ với `class`)**|Một lớp kế thừa từ nhiều lớp cha (_Java không hỗ trợ với `class`, chỉ hỗ trợ với `interface`_)|

### **🔹 1. Single Inheritance (Kế thừa đơn)**
```java
class Animal {
    void sound() {
        System.out.println("Động vật phát ra âm thanh.");
    }
}

class Dog extends Animal {
    void bark() {
        System.out.println("Chó sủa!");
    }
}
```

🔹 2. Multilevel Inheritance (Kế thừa đa cấp)
```java
class Animal {
    void eat() {
        System.out.println("Động vật đang ăn...");
    }
}

class Mammal extends Animal {
    void walk() {
        System.out.println("Thú đi bằng chân.");
    }
}

class Dog extends Mammal {
    void bark() {
        System.out.println("Chó sủa!");
    }
}
```
📌 **Lớp `Dog` kế thừa `Mammal`, và `Mammal` kế thừa `Animal`, nên `Dog` có thể gọi cả `eat()`, `walk()` và `bark()`.**


🔹 3. Hierarchical Inheritance (Kế thừa phân cấp)
```java
class Animal {
    void sound() {
        System.out.println("Động vật phát ra âm thanh.");
    }
}

class Dog extends Animal {
    void bark() {
        System.out.println("Chó sủa!");
    }
}

class Cat extends Animal {
    void meow() {
        System.out.println("Mèo kêu: Meo Meo!");
    }
}
```
📌 **`Dog` và `Cat` đều kế thừa `Animal`, nhưng có phương thức riêng biệt.**

### **🔹 4. Hybrid Inheritance (Kế thừa lai)**

📌 **Không thể thực hiện bằng `class`, nhưng có thể dùng `interface`.**
```java
interface Animal {
    void sound();
}

interface Mammal {
    void walk();
}

// Lớp này kế thừa từ 2 interface
class Dog implements Animal, Mammal {
    public void sound() {
        System.out.println("Chó sủa!");
    }

    public void walk() {
        System.out.println("Chó chạy bằng 4 chân.");
    }
}
```
📌 **Java hỗ trợ `Multiple Inheritance` qua `interface`, nhưng không hỗ trợ với `class` do vấn đề `Diamond Problem`.**

## **3️⃣ Tái sử dụng mã nguồn (Reusability)**

- **Lợi ích lớn nhất của kế thừa là tái sử dụng mã nguồn.**
- **Giảm trùng lặp code và giúp bảo trì dễ dàng hơn.**
- **Tuy nhiên, không phải lúc nào cũng nên dùng kế thừa nếu có thể dùng `composition`.**

## **4️⃣ OOP có thể tồn tại mà không có kế thừa không?**

✅ **Có thể!**  
📌 **Kế thừa không phải là bắt buộc trong OOP.**

---
# 🔥 **Hiểu Sâu về Interface và Implement trong Java**

## 📌 **1️⃣ Interface là gì?**

### **Khái niệm**

- **Interface** trong Java là một **hợp đồng** (contract) quy định rằng một lớp phải triển khai (implement) các phương thức mà nó khai báo.
- Tất cả các phương thức trong interface **mặc định là `public abstract`**.
- Interface giúp **định nghĩa hành vi chung** mà nhiều lớp có thể triển khai.
Cú pháp
```JAVA
interface Animal {
    void makeSound(); // Mặc định là public abstract
}
```

## 📌 **2️⃣ Implement là gì?**

- Khi một lớp **`implements`** một interface, nó phải triển khai tất cả các phương thức của interface đó.
- Một lớp có thể `implement` nhiều interface cùng lúc (hỗ trợ **đa kế thừa** gián tiếp).

🔹 **Ví dụ về interface và implement**:
```java
// Định nghĩa Interface
interface Animal {
    void makeSound();
}

// Lớp Dog triển khai (implement) interface Animal
class Dog implements Animal {
    public void makeSound() {
        System.out.println("Chó sủa: Gâu Gâu!");
    }
}

// Lớp Cat cũng triển khai interface Animal
class Cat implements Animal {
    public void makeSound() {
        System.out.println("Mèo kêu: Meo Meo!");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal myDog = new Dog();
        myDog.makeSound(); // Output: Chó sủa: Gâu Gâu!

        Animal myCat = new Cat();
        myCat.makeSound(); // Output: Mèo kêu: Meo Meo!
    }
}
```
📌 **Lớp `Dog` và `Cat` triển khai (`implements`) interface `Animal`, nhưng mỗi lớp có cách triển khai riêng.**



## 📌 **3️⃣ Interface vs Abstract Class – Khi nào dùng gì?**

|**Đặc điểm**|**Interface**|**Abstract Class**|
|---|---|---|
|**Mục đích**|Định nghĩa hành vi chung (contract)|Tạo lớp cha có thể chứa logic chung|
|**Từ khóa**|`interface`|`abstract class`|
|**Phương thức**|Chỉ có `public abstract` (Java 7 trở xuống)|Có thể có cả `abstract` và `concrete` methods|
|**Thuộc tính (fields)**|Chỉ có hằng số (`public static final`)|Có cả biến instance và biến tĩnh|
|**Đa kế thừa**|Hỗ trợ đa kế thừa (`implements` nhiều interface)|Không hỗ trợ đa kế thừa (`extends` chỉ 1 class)|
|**Constructor**|Không có|Có thể có|

📌 **Khi nào dùng `interface`?**

- Khi **các lớp không có quan hệ cha-con**, nhưng cần **chung một hành vi**.
- Khi **cần đa kế thừa** (vì một lớp có thể `implements` nhiều interface).

📌 **Khi nào dùng `abstract class`?**

- Khi **các lớp có quan hệ cha-con**, và **có logic chung** có thể kế thừa.

🔹 **Ví dụ về Abstract Class và Interface kết hợp**:
```java
// Interface
interface Swimmable {
    void swim();
}

// Abstract Class
abstract class Animal {
    abstract void makeSound();
}

// Lớp Dog kế thừa Animal và implement Swimmable
class Dog extends Animal implements Swimmable {
    public void makeSound() {
        System.out.println("Chó sủa: Gâu Gâu!");
    }
    
    public void swim() {
        System.out.println("Chó có thể bơi!");
    }
}
```

📌 **Ở đây `Dog` vừa kế thừa từ `Animal` vừa implement từ `Swimmable`.**

## 📌 **4️⃣ Một lớp có thể implements nhiều interface không?**

✅ **Có!**

- Java hỗ trợ **đa kế thừa với interface** (một lớp có thể `implements` nhiều interface).


## 📌 **5️⃣ Default và Static Methods trong Interface (Java 8+)**

- **Trước Java 8**, tất cả phương thức trong interface phải là **abstract**.
- **Java 8+** cho phép có **default method** và **static method** trong interface.

```java
interface Vehicle {
    void start();
    
    // Default method có sẵn trong interface (không bắt buộc phải override)
    default void honk() {
        System.out.println("Bíp bíp! Đây là còi mặc định.");
    }
}

class Car implements Vehicle {
    public void start() {
        System.out.println("Xe khởi động!");
    }
}

public class Main {
    public static void main(String[] args) {
        Car myCar = new Car();
        myCar.start();
        myCar.honk(); // Gọi method mặc định từ interface
    }
}
```
📌 **Nếu một lớp `implements` interface nhưng không override `default method`, thì vẫn có thể dùng nó.**

🔹 **Ví dụ về `static` method trong interface**:
```java
interface MathUtils {
    static int add(int a, int b) {
        return a + b;
    }
}

public class Main {
    public static void main(String[] args) {
        int sum = MathUtils.add(5, 10);
        System.out.println("Tổng là: " + sum);
    }
}
```
📌 **Interface có thể chứa `static` method giống như class bình thường.**

## 📌 **6️⃣ Interface Kế Thừa Interface**

✅ **Một interface có thể kế thừa (`extends`) interface khác.**
```java
interface Animal {
    void eat();
}

interface Bird extends Animal {
    void fly();
}

class Sparrow implements Bird {
    public void eat() {
        System.out.println("Chim sẻ ăn hạt.");
    }
    
    public void fly() {
        System.out.println("Chim sẻ bay.");
    }
}
```
📌 **`Bird` kế thừa `Animal`, nên `Sparrow` phải triển khai cả `eat()` và `fly()`.**

## 📌 **7️⃣ Một lớp có thể vừa extends một class vừa implements interface không?**

✅ **Có!**

- Java cho phép một lớp **kế thừa (`extends`) một class và đồng thời `implements` nhiều interface**.

🔹 **Ví dụ**:
```java
class Animal {
    void eat() {
        System.out.println("Động vật đang ăn...");
    }
}

interface Runnable {
    void run();
}

class Dog extends Animal implements Runnable {
    public void run() {
        System.out.println("Chó chạy rất nhanh!");
    }
}
```
📌 **Lớp `Dog` vừa kế thừa `Animal` vừa `implements` `Runnable`.**


## **8️⃣ Tổng Kết**

🔹 **Interface giúp định nghĩa hành vi chung, hỗ trợ đa kế thừa**  
🔹 **Lớp `implements` interface phải override tất cả các phương thức**  
🔹 **Java 8+ hỗ trợ `default` và `static` methods trong interface**  
🔹 **Một lớp có thể vừa `extends` class vừa `implements` nhiều interface**

