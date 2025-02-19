
---

# 🔥 **Abstraction trong Java – Hiểu Sâu Từ Cơ Bản Đến Nâng Cao**

---

## 📌 **1️⃣ Abstraction là gì?**

**Abstraction (Trừu tượng hóa)** là một nguyên tắc quan trọng trong **Lập trình hướng đối tượng (OOP)**, giúp **ẩn đi các chi tiết không cần thiết** và chỉ hiển thị những gì quan trọng.

**💡 Định nghĩa**:  
🔹 **Chỉ hiển thị hành vi (methods) cần thiết** mà **không tiết lộ chi tiết triển khai** bên trong.  
🔹 **Giúp tập trung vào "cái gì" (What), thay vì "như thế nào" (How)**.

---

## 📌 **2️⃣ Khi nào nên sử dụng Abstraction?**

✅ Khi bạn muốn **giấu đi các chi tiết triển khai phức tạp** và chỉ cung cấp một giao diện chung cho các lớp con.  
✅ Khi bạn muốn **định nghĩa một "khung sườn" chung** cho các lớp liên quan mà chưa cần biết chi tiết từng lớp cụ thể.  
✅ Khi bạn muốn **hạn chế khả năng thay đổi hoặc truy cập trực tiếp vào logic bên trong của một đối tượng**.

Ví dụ:

- **Xe ô tô**: Khi lái xe, bạn chỉ cần biết cách **đạp ga, thắng, đánh lái**, chứ không cần quan tâm cách động cơ hoạt động bên trong.
- **Hệ thống thanh toán**: Bạn chỉ cần gọi phương thức `processPayment()`, chứ không cần biết chi tiết nó kết nối với ngân hàng thế nào.

---

## 📌 **3️⃣ Cách đạt được Abstraction trong Java?**

Có **2 cách chính** để đạt được Abstraction trong Java:

### 🔹 **Cách 1: Sử dụng Abstract Class (`class abstract`)**

✅ **Là một lớp chứa phương thức `abstract` (chưa có triển khai)** và có thể chứa phương thức bình thường.  
✅ **Không thể khởi tạo (instantiate) trực tiếp** mà phải được kế thừa bởi lớp con.  
✅ **Có thể có constructor, biến, và cả phương thức có thân (`concrete methods`)**.

📌 **Ví dụ: Abstract Class trong Java**
```java
abstract class Vehicle {
    abstract void start();  // Phương thức abstract (không có triển khai)
    
    public void stop() {  // Phương thức bình thường (có triển khai)
        System.out.println("Xe đã dừng lại!");
    }
}

class Car extends Vehicle {
    @Override
    void start() {
        System.out.println("Xe hơi khởi động bằng chìa khóa!");
    }
}

class Bike extends Vehicle {
    @Override
    void start() {
        System.out.println("Xe máy khởi động bằng nút bấm!");
    }
}

public class Main {
    public static void main(String[] args) {
        Vehicle car = new Car();
        car.start(); // Gọi phương thức trừu tượng
        car.stop();  // Gọi phương thức cụ thể
    }
}
```
🔹 **Kết quả đầu ra:**
```
Xe hơi khởi động bằng chìa khóa!
Xe đã dừng lại!
```


✅ **Ưu điểm:**

- Giúp tổ chức mã nguồn rõ ràng hơn.
- Tạo một chuẩn chung mà các lớp con phải tuân theo.
- Có thể chứa cả phương thức `abstract` và phương thức `concrete`.

---
### 🔹 **Cách 2: Sử dụng Interface (`interface`)**

✅ **Chỉ chứa phương thức trừu tượng (trước Java 8)**, không có phần thân (implementation).  
✅ **Lớp thực thi (`implements`) interface bắt buộc phải override tất cả các phương thức**.  
✅ **Hỗ trợ đa kế thừa** (một class có thể `implements` nhiều interface).

📌 **Ví dụ: Interface trong Java**
```java
interface Animal {
    void makeSound();  // Phương thức trừu tượng
}

class Dog implements Animal {
    @Override
    public void makeSound() {
        System.out.println("Gâu gâu!");
    }
}

class Cat implements Animal {
    @Override
    public void makeSound() {
        System.out.println("Meo meo!");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal dog = new Dog();
        dog.makeSound();  // Output: Gâu gâu!
        
        Animal cat = new Cat();
        cat.makeSound();  // Output: Meo meo!
    }
}
```

🔹 **Kết quả đầu ra:**
```
Gâu gâu!
Meo meo!
```

✅ **Ưu điểm:**

- **Hỗ trợ đa kế thừa** (Java không hỗ trợ `extends` nhiều class nhưng cho phép `implements` nhiều interface).
- **Tạo một "bản hợp đồng" (contract)** bắt buộc các lớp triển khai phải tuân theo.
- **Dễ dàng mở rộng và thay đổi trong tương lai**.

## 📌 **4️⃣ So sánh Abstraction và các khái niệm khác**

### 🆚 **Abstraction vs. Inheritance**

|Đặc điểm|Abstraction|Inheritance|
|---|---|---|
|**Mục đích**|Che giấu chi tiết không cần thiết|Tái sử dụng mã nguồn từ lớp cha|
|**Dùng khi nào?**|Khi chỉ cần xác định hành vi chung|Khi cần mở rộng một lớp hiện có|
|**Công cụ**|`abstract class`, `interface`|`extends` (Kế thừa lớp cha)|
|**Tính linh hoạt**|Cho phép tạo quy tắc chung|Cho phép chia sẻ dữ liệu và phương thức|

📌 **Hiểu đơn giản:**

- **Abstraction** giúp **tạo một khuôn mẫu** để các lớp con phải tuân theo.
- **Inheritance** giúp **kế thừa các thuộc tính và phương thức** từ lớp cha.

### 🆚 **Abstraction vs. Encapsulation**

|Đặc điểm|Abstraction|Encapsulation|
|---|---|---|
|**Mục đích**|Ẩn chi tiết triển khai, chỉ hiển thị cái cần thiết|Ẩn dữ liệu bên trong và kiểm soát truy cập|
|**Cách thực hiện**|`abstract class`, `interface`|`private`, `protected`, `getter/setter`|
|**Tập trung vào**|Hành vi của đối tượng|Bảo vệ và kiểm soát dữ liệu bên trong|
|**Tính bảo mật**|Không bảo vệ dữ liệu, chỉ che giấu logic|Bảo vệ dữ liệu khỏi truy cập trái phép|

📌 **Hiểu đơn giản:**

- **Abstraction** = **Tập trung vào "cái gì"**, không quan tâm "như thế nào".
- **Encapsulation** = **Bảo vệ dữ liệu**, kiểm soát quyền truy cập.

---

## 📌 **5️⃣ Tổng kết**

🔹 **Abstraction giúp:**  
✅ Ẩn đi chi tiết không quan trọng, chỉ hiển thị thông tin cần thiết.  
✅ Giảm phụ thuộc giữa các phần của chương trình.  
✅ Cung cấp cấu trúc chung mà các lớp con phải tuân theo.

🔹 **Cách đạt được Abstraction:**  
✅ **Abstract Class**: Chứa phương thức `abstract` và `concrete`.  
✅ **Interface**: Chỉ chứa phương thức `abstract`, hỗ trợ đa kế thừa.

🔹 **Phân biệt với các khái niệm khác:**  
✅ **So với Inheritance**: Abstraction tạo quy tắc, Inheritance tái sử dụng mã nguồn.  
✅ **So với Encapsulation**: Abstraction tập trung vào hành vi, Encapsulation bảo vệ dữ liệu.

💡 **Nắm vững Abstraction giúp bạn viết mã dễ bảo trì, mở rộng và tổ chức chương trình tốt hơn!** 🚀


---
# ✅ **Khi nào dùng `abstract class`, khi nào dùng `interface`?**

Trong Java, cả **`abstract class`** và **`interface`** đều được sử dụng để đạt **Abstraction (Trừu tượng hóa)**. Tuy nhiên, việc chọn cái nào phụ thuộc vào **mục đích sử dụng**. Dưới đây là **các trường hợp cụ thể** giúp bạn quyết định nên dùng cái nào.

## 🎯 **1️⃣ Khi nào dùng `abstract class`?**

✅ **Dùng khi bạn cần một lớp cơ sở có thể chứa cả phương thức trừu tượng (`abstract`) và phương thức cụ thể (`concrete`).**  
✅ **Dùng khi bạn muốn cung cấp một số logic mặc định mà các lớp con có thể kế thừa.**  
✅ **Dùng khi có chung một phần thuộc tính (fields) và hành vi giữa các lớp con.**  
✅ **Dùng khi cần bảo toàn dữ liệu chung mà không muốn mọi lớp con tự định nghĩa lại từ đầu.**

📌 **Ví dụ:**  
Một **`abstract class`** có thể chứa **thuộc tính chung**, ví dụ như tất cả động vật đều có **tên và tuổi**, nhưng cách chúng kêu thì khác nhau.

**📌 Khi nào nên dùng `abstract class`?**  
👉 Khi có **thuộc tính chung** (name, age) cần dùng cho tất cả lớp con.  
👉 Khi có một số **hành vi mặc định** (`showInfo()` không bắt buộc override).


---
## 🎯 **2️⃣ Khi nào dùng `interface`?**

✅ **Dùng khi bạn chỉ muốn định nghĩa hành vi mà không cần thuộc tính hay phương thức cụ thể.**  
✅ **Dùng khi bạn muốn hỗ trợ đa kế thừa (vì Java không cho phép `extends` nhiều class nhưng có thể `implements` nhiều interface).**  
✅ **Dùng khi bạn muốn tách biệt hoàn toàn logic giữa các lớp.**

📌 **Ví dụ:**  
Giả sử bạn có nhiều loại động vật khác nhau, nhưng không phải con nào cũng có thể bay hoặc bơi. Khi đó, bạn dùng **`interface`** để chỉ định hành vi mà không ảnh hưởng đến cấu trúc lớp cha.

**📌 Khi nào nên dùng `interface`?**  
👉 Khi bạn cần **đa kế thừa** (`Bird` có thể `implements` cả `Flyable` và `Swimmable` nếu nó biết bơi).  
👉 Khi bạn chỉ quan tâm **hành vi mà không cần lưu trữ dữ liệu chung**.


## 🔥 **3️⃣ So sánh chi tiết `abstract class` vs `interface`**

|Tiêu chí|`abstract class`|`interface`|
|---|---|---|
|**Mục đích**|Dùng để tạo lớp cha có thể chứa cả phương thức `abstract` và `concrete`|Dùng để định nghĩa các hành vi mà lớp phải tuân theo|
|**Chứa phương thức có thân (`concrete method`)?**|✅ Có thể chứa|❌ Không thể chứa (trước Java 8), từ Java 8 có thể có `default` và `static methods`|
|**Chứa biến (`fields`)?**|✅ Có thể chứa (biến `protected` hoặc `private`)|❌ Chỉ có biến `public static final` (hằng số)|
|**Đa kế thừa**|❌ Không hỗ trợ (vì Java không cho phép `extends` nhiều class)|✅ Hỗ trợ (một class có thể `implements` nhiều interface)|
|**Có constructor?**|✅ Có|❌ Không|
|**Khi nào nên dùng?**|Khi có chung trạng thái (biến, thuộc tính) và cần triển khai một số phương thức chung|Khi chỉ cần định nghĩa hành vi, hỗ trợ đa kế thừa|
## 💡 **4️⃣ Khi nào nên chọn cái nào?**

### ✅ **Dùng `abstract class` khi:**

✔ Khi có **thuộc tính chung** giữa các lớp con.  
✔ Khi có một số **hành vi mặc định** mà không muốn tất cả các lớp con phải triển khai lại.  
✔ Khi muốn sử dụng **constructor** trong lớp cha.

📌 **Ví dụ thực tế:**

- **Lớp `abstract` `Vehicle` chứa `maxSpeed` và phương thức `move()`.**
- **Lớp `Car`, `Bike` kế thừa và sử dụng chung phương thức `move()`.**

---

### ✅ **Dùng `interface` khi:**

✔ Khi muốn tạo một **giao diện chung** cho các lớp không có quan hệ cha-con.  
✔ Khi muốn **hỗ trợ đa kế thừa** (vì một class có thể implements nhiều interface).  
✔ Khi muốn **định nghĩa hành vi mà không cần lưu trữ trạng thái chung**.

📌 **Ví dụ thực tế:**

- **`Flyable` và `Swimmable` dùng cho các lớp khác nhau như `Bird`, `Fish`, `Superman` mà không ảnh hưởng đến class gốc.**
- **Java `List`, `Set`, `Map` đều `implements` interface `Collection`, dù chúng có cách hoạt động khác nhau.**

## 🎯 **5️⃣ Tổng kết**

🔹 **Dùng `abstract class` khi** bạn cần một **lớp cha có thuộc tính chung và có thể chứa phương thức triển khai sẵn**.  
🔹 **Dùng `interface` khi** bạn chỉ cần **định nghĩa hành vi** và muốn **hỗ trợ đa kế thừa**.

💡 **Gợi ý:**  
✔ Nếu không biết chọn cái nào, hãy **ưu tiên `interface` nếu không cần thuộc tính chung**.  
✔ Nếu muốn tạo một "bộ khung" có cả thuộc tính và phương thức mặc định, **dùng `abstract class`**.