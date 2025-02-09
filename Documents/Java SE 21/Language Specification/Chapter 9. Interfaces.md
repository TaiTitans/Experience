
---
**Quotes:**
*Đừng theo đuổi mà hãy thu hút. Ngồi lựa đào hoặc để phú bà lựa.

## **🔥 1. Giới thiệu về Giao diện (Interfaces)**

- **Giao diện (Interface)** trong Java là một tập hợp các phương thức trừu tượng (abstract methods) và các hằng số (constants).
- Một interface **không thể tạo đối tượng trực tiếp**, nhưng có thể được **các lớp khác triển khai (implement)**.
- Kể từ Java 8, interface có thể chứa:
    - **Phương thức mặc định (default methods)**
    - **Phương thức tĩnh (static methods)**
    - **Phương thức riêng tư (private methods) (từ Java 9)**

📌 **Ví dụ cơ bản về Interface:**
```java
interface Animal {
    void makeSound(); // Phương thức trừu tượng
}

class Dog implements Animal {
    public void makeSound() {
        System.out.println("Woof!");
    }
}
```
## **🔥 2. Cú pháp khai báo Giao diện**

Cú pháp:
```java
<modifier> interface <TênGiaoDiện> {
    // Hằng số (constants)
    // Phương thức trừu tượng
    // Phương thức mặc định
    // Phương thức tĩnh
}
```
📌 **Ví dụ:**
```java
public interface Vehicle {
    int MAX_SPEED = 120; // Hằng số

    void start(); // Phương thức trừu tượng

    default void show() { // Phương thức mặc định
        System.out.println("This is a vehicle");
    }

    static void display() { // Phương thức tĩnh
        System.out.println("Static method in interface");
    }
}
```
## **🔥 3. Triển khai Interface trong Java**

Một lớp có thể triển khai (implement) một hoặc nhiều giao diện bằng từ khóa `implements`.
```java
interface Animal {
    void makeSound();
}

// Lớp Dog triển khai interface Animal
class Dog implements Animal {
    public void makeSound() {
        System.out.println("Woof!");
    }
}

```
## **🔥 4. Một lớp có thể triển khai nhiều Giao diện**

Một lớp có thể triển khai nhiều giao diện cùng lúc.
```java
interface A {
    void methodA();
}

interface B {
    void methodB();
}

// Lớp C triển khai cả hai interface A và B
class C implements A, B {
    public void methodA() {
        System.out.println("Method A");
    }

    public void methodB() {
        System.out.println("Method B");
    }
}

```

## **🔥 5. Phương thức mặc định (Default Methods)**

- **Từ Java 8**, interface có thể chứa phương thức mặc định (`default`) có phần thân (body).
- Các lớp triển khai **không bắt buộc phải override phương thức mặc định**.

📌 **Ví dụ:**
```java
interface Vehicle {
    default void show() {
        System.out.println("This is a vehicle");
    }
}

class Car implements Vehicle {}

public class Main {
    public static void main(String[] args) {
        Car car = new Car();
        car.show(); // Gọi phương thức mặc định
    }
}
```
## **🔥 6. Phương thức tĩnh (Static Methods)**

- **Từ Java 8**, interface có thể chứa phương thức `static`.
- Phương thức này chỉ có thể được gọi thông qua interface, **không thể override trong lớp triển khai**.

📌 **Ví dụ:**
```java
interface MathUtils {
    static int add(int a, int b) {
        return a + b;
    }
}

public class Main {
    public static void main(String[] args) {
        int result = MathUtils.add(5, 10);
        System.out.println(result); // 15
    }
}
```
## **🔥 7. Phương thức riêng tư (Private Methods) (Java 9+)**

- **Từ Java 9**, interface có thể có phương thức `private`.
- Phương thức này **chỉ có thể được gọi từ phương thức mặc định hoặc phương thức tĩnh của chính giao diện đó**.

📌 **Ví dụ:**
```java
interface Logger {
    private void log(String message) {
        System.out.println("Logging: " + message);
    }

    default void info(String message) {
        log("INFO: " + message);
    }
}

class AppLogger implements Logger {}

public class Main {
    public static void main(String[] args) {
        AppLogger logger = new AppLogger();
        logger.info("Application started");
    }
}
```
## **🔥 8. Interface mở rộng (Extending Interfaces)**

- Một interface có thể kế thừa từ một hoặc nhiều giao diện khác bằng từ khóa `extends`.

📌 **Ví dụ:**
```java
interface A {
    void methodA();
}

interface B extends A {
    void methodB();
}

class C implements B {
    public void methodA() {
        System.out.println("Method A");
    }

    public void methodB() {
        System.out.println("Method B");
    }
}
```
## **🔥 9. Mối quan hệ giữa Lớp và Interface**

| Đặc điểm                      | Lớp               | Interface                                          |
| ----------------------------- | ----------------- | -------------------------------------------------- |
| Hỗ trợ đa kế thừa             | ❌ Không           | ✅ Có                                               |
| Chứa phương thức có phần thân | ✅ Có (Java SE 7-) | ✅ Có (`default`, `static`, `private` - Java SE 8+) |
| Có constructor                | ✅ Có              | ❌ Không                                            |
| Có thể chứa biến instance     | ✅ Có              | ❌ Không (Chỉ có `static final`)                    |
## **🔥 10. Từ khóa `instanceof` với Interface**

- Có thể kiểm tra xem một đối tượng có thuộc một interface hay không bằng từ khóa `instanceof`.

📌 **Ví dụ:**
```java
interface Animal {}
class Dog implements Animal {}

public class Main {
    public static void main(String[] args) {
        Dog dog = new Dog();
        System.out.println(dog instanceof Animal); // true
    }
}
```
## **🔥 11. Tóm tắt**

✅ Giao diện trong Java là một tập hợp các **phương thức trừu tượng và hằng số**.  
✅ **Lớp có thể triển khai nhiều giao diện** bằng từ khóa `implements`.  
✅ **Từ Java 8**, giao diện có thể có **phương thức mặc định (`default`) và phương thức tĩnh (`static`)**.  
✅ **Từ Java 9**, giao diện có thể có **phương thức riêng tư (`private`)**.  
✅ **Giao diện có thể kế thừa từ giao diện khác** bằng từ khóa `extends`.

