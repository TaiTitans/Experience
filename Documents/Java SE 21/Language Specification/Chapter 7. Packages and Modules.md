
---
### **🔥 1. Giới thiệu về Khai báo trong Java**

Trong Java, **khai báo (declaration)** được sử dụng để giới thiệu các **biến, phương thức, lớp, giao diện, package, và module** vào phạm vi của chương trình.

Khi khai báo, bạn có thể chỉ định **tên, kiểu dữ liệu, phạm vi truy cập, và các thuộc tính (modifiers)** như `static`, `final`, `abstract`, v.v.
### **🔥 2. Các loại khai báo trong Java**

|Loại Khai Báo|Ví Dụ|
|---|---|
|**Biến (Variable)**|`int age = 25;`|
|**Phương thức (Method)**|`public void sayHello() {}`|
|**Lớp (Class)**|`public class Person {}`|
|**Giao diện (Interface)**|`interface Flyable {}`|
|**Gói (Package)**|`package com.example;`|
|**Module** (Java 9+)|`module my.module {}`|
### **🔥 3. Khai báo Biến (Variable Declarations)**

Biến là thành phần cơ bản trong Java để lưu trữ dữ liệu.

#### **3.1. Cú pháp khai báo biến**
`<modifier> <kiểu dữ liệu> <tên biến> = <giá trị>;`
📌 **Ví dụ:**
```java
int age = 25;
double price = 19.99;
String name = "Alice";
```
#### **3.2. Các loại biến trong Java**

Có 3 loại biến chính:

|Loại biến|Định nghĩa|Ví dụ|
|---|---|---|
|**Biến cục bộ (Local Variable)**|Khai báo trong một phương thức hoặc khối mã. Chỉ có thể truy cập trong phạm vi đó.|`void test() { int x = 10; }`|
|**Biến thực thể (Instance Variable)**|Thuộc về một đối tượng, có thể được truy cập thông qua `this`.|`class Person { String name; }`|
|**Biến lớp (Class Variable - static)**|Được khai báo với `static`, thuộc về lớp thay vì từng đối tượng.|`static int count;`|

#### **3.3. Khai báo biến với Modifiers**

|Modifier|Ý nghĩa|Ví dụ|
|---|---|---|
|`final`|Không thể thay đổi giá trị sau khi gán lần đầu|`final int MAX = 100;`|
|`static`|Thuộc về lớp thay vì đối tượng|`static int counter = 0;`|
|`volatile`|Đảm bảo giá trị luôn được cập nhật giữa các thread|`volatile boolean running;`|

📌 **Ví dụ về `final`, `static` và `volatile`:**
```java
class Example {
    static int count = 0;  // Biến lớp (static)
    final int MAX_VALUE = 100;  // Biến không đổi (final)
    volatile boolean isRunning; // Biến có thể thay đổi giữa các thread
}
```
### **🔥 4. Khai báo Phương thức (Method Declarations)**

Phương thức là khối mã có thể gọi để thực thi một chức năng nào đó.

#### **4.1. Cú pháp khai báo phương thức**

```
<modifier> <kiểu dữ liệu trả về> <tên phương thức>(<tham số>) { 
    // Thân phương thức
}
```
📌 **Ví dụ:**
```java
public void sayHello() {
    System.out.println("Hello!");
}
```
#### **4.2. Modifier của phương thức**

| Modifier       | Ý nghĩa                                 | Ví dụ                           |
| -------------- | --------------------------------------- | ------------------------------- |
| `final`        | Không thể override trong lớp con        | `final void print() {}`         |
| `static`       | Gọi mà không cần tạo đối tượng          | `static void hello() {}`        |
| `abstract`     | Phương thức không có thân, chỉ khai báo | `abstract void run();`          |
| `synchronized` | Chỉ cho phép một thread truy cập        | `synchronized void update() {}` |
📌 **Ví dụ về `final`, `static`, và `abstract`:**
```java
abstract class Animal {
    abstract void makeSound(); // Phương thức trừu tượng
}

class Dog extends Animal {
    @Override
    void makeSound() {
        System.out.println("Woof!");
    }
}
```
### **🔥 5. Khai báo Lớp và Giao diện (Class & Interface Declarations)**

#### **5.1. Khai báo Lớp (Class)**

Cú pháp:
```java
<modifier> class <TênLớp> {
    // Thành phần của lớp
}
```
📌 **Ví dụ:**
```java
public class Person {
    String name;
    
    public void sayHello() {
        System.out.println("Hello, " + name);
    }
}
```
#### **5.2. Modifier của lớp**

|Modifier|Ý nghĩa|Ví dụ|
|---|---|---|
|`final`|Không thể bị kế thừa|`final class MathUtils {}`|
|`abstract`|Không thể tạo đối tượng từ lớp này|`abstract class Animal {}`|

📌 **Ví dụ về `final` và `abstract`:**
```java
final class Utility {
    public static void printMessage() {
        System.out.println("Hello");
    }
}
```
### **🔥 6. Khai báo Gói (Package Declarations)**

Package giúp tổ chức các lớp và giao diện theo nhóm.

Cú pháp:
```java
package com.example;
```
📌 **Ví dụ:**
```java
package com.mekongocop.service;
public class OrderService { }
```
### **🔥 7. Khai báo Module (Java 9+)**

Module giúp tổ chức mã nguồn thành các đơn vị độc lập.

Cú pháp:
```java
module <TênModule> {
    requires <TênModuleKhác>;
    exports <TênPackage>;
}
```
📌 **Ví dụ:**
```java
module com.mekongocop {
    exports com.mekongocop.service;
}
```
### **🔥 8. Tóm tắt & Kết luận**

✅ **Khai báo trong Java** bao gồm biến, phương thức, lớp, giao diện, package và module.  
✅ **Modifiers** như `static`, `final`, `abstract` giúp kiểm soát hành vi của các thành phần.  
✅ **Phạm vi khai báo (Scope)** quyết định nơi một thành phần có thể được sử dụng.  
✅ **Module (Java 9+)** giúp tổ chức mã nguồn lớn theo mô-đun.