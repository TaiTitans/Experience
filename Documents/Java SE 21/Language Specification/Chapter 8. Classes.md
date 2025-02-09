
---
`Lưu ý: MUỐN ĐÀO SÂU OOP THÌ TỰ MÀ ĐÀO ĐÂY LÀ JAVA CORE OK!!!`
### **🔥 1. Giới thiệu về Lớp (Classes)**

Trong Java, **lớp (class)** là một khuôn mẫu (blueprint) để tạo ra các đối tượng (objects). Một lớp có thể chứa:

- **Biến (fields)**: dùng để lưu trữ trạng thái của đối tượng.
- **Phương thức (methods)**: dùng để định nghĩa hành vi của đối tượng.
- **Khởi tạo (constructors)**: dùng để khởi tạo đối tượng.
- **Lớp lồng nhau (nested classes)**: gồm inner class, static nested class, local class và anonymous class.
### **🔥 2. Cú pháp khai báo Lớp**

Cú pháp:
```java
<modifier> class <TênLớp> {
    // Các biến (fields)
    // Các phương thức (methods)
    // Các khối khởi tạo (initializer blocks)
    // Các lớp con (nested classes)
}
```
📌 **Ví dụ cơ bản:**
```java
public class Animal {
    String name; // Thuộc tính

    // Constructor
    public Animal(String name) {
        this.name = name;
    }

    // Phương thức
    public void makeSound() {
        System.out.println(name + " is making a sound.");
    }
}
```
### **🔥 3. Thành phần của một Lớp**

Một lớp có thể chứa các thành phần sau:

| Thành phần                | Mô tả                            | Ví dụ                                   |
| ------------------------- | -------------------------------- | --------------------------------------- |
| **Biến (fields)**         | Lưu trữ trạng thái của đối tượng | `private int age;`                      |
| **Hằng số (constants)**   | Biến có giá trị không thay đổi   | `public static final double PI = 3.14;` |
| **Phương thức (methods)** | Định nghĩa hành vi của đối tượng | `public void eat() {}`                  |
| **Constructor**           | Khởi tạo đối tượng               | `public Animal(String name) {}`         |
| **Khối khởi tạo**         | Chạy khi tạo đối tượng           | `static {}` hoặc `{}`                   |
| **Lớp lồng nhau**         | Lớp được khai báo trong lớp khác | `static class InnerClass {}`            |
### **🔥 4. Các loại Lớp trong Java**

#### **4.1. Lớp thông thường (Regular Class)**

Lớp thông thường là lớp có thể tạo đối tượng trực tiếp.

📌 **Ví dụ:**
```java
public class Dog {
    String breed;

    public Dog(String breed) {
        this.breed = breed;
    }

    public void bark() {
        System.out.println("Woof!");
    }
}
```

#### **4.2. Lớp trừu tượng (Abstract Class)**

Lớp trừu tượng là lớp không thể tạo đối tượng trực tiếp và có thể chứa **phương thức trừu tượng**.

📌 **Ví dụ:**
```java
abstract class Animal {
    abstract void makeSound(); // Phương thức trừu tượng
}
```
#### **4.3. Lớp lồng nhau (Nested Class)**

Lớp lồng nhau là lớp nằm bên trong một lớp khác.
```Java
class OuterClass {
    class InnerClass {
        void show() {
            System.out.println("Hello from InnerClass");
        }
    }
}
```
#### **4.4. Lớp vô danh (Anonymous Class)**

Lớp vô danh là lớp không có tên, được tạo ngay trong lúc khởi tạo một đối tượng.

📌 **Ví dụ:**
```java
interface Animal {
    void makeSound();
}

public class Main {
    public static void main(String[] args) {
        Animal dog = new Animal() {
            public void makeSound() {
                System.out.println("Woof!");
            }
        };
        dog.makeSound();
    }
}
```
### **🔥 5. Kế thừa (Inheritance)**

Lớp con có thể kế thừa từ lớp cha bằng từ khóa `extends`.

📌 **Ví dụ:**
```java
class Animal {
    void makeSound() {
        System.out.println("Animal makes a sound");
    }
}

class Dog extends Animal {
    void bark() {
        System.out.println("Woof!");
    }
}
```
### **🔥 6. Từ khóa `super` và `this`**

- `super`: Dùng để gọi constructor hoặc phương thức của lớp cha.
- `this`: Dùng để tham chiếu đến chính đối tượng hiện tại.

📌 **Ví dụ về `super và this`:**
```java
class Animal {
    Animal() {
        System.out.println("Animal constructor");
    }
}

class Dog extends Animal {
    Dog() {
        super(); // Gọi constructor của Animal
        System.out.println("Dog constructor");
    }
}
```
```java
class Dog {
    String breed;

    Dog(String breed) {
        this.breed = breed; // Tham chiếu đến biến instance
    }
}

```
### **🔥 7. Từ khóa `final` trong Lớp**

- **`final class`**: Không cho phép lớp khác kế thừa.
- **`final method`**: Không cho phép lớp con ghi đè.
- **`final variable`**: Không thể thay đổi giá trị.

📌 **Ví dụ:**
```java
final class Animal {} // Không thể bị kế thừa

class Dog {
    final void bark() {
        System.out.println("Woof!");
    }
}
```
### **🔥 8. Tóm tắt**

✅ Một lớp trong Java có thể chứa các **biến, phương thức, constructor, khối khởi tạo, lớp lồng nhau**.  
✅ Java hỗ trợ **lớp trừu tượng, lớp lồng nhau, lớp vô danh**.  
✅ **Từ khóa `extends`** để kế thừa, **`super` và `this`** để tham chiếu.  
✅ **`final`** được dùng để ngăn kế thừa và ghi đè.

