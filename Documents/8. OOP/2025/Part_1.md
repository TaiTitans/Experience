
---
## **1. Class là gì?**

Lớp (**Class**) là một bản thiết kế (blueprint) để tạo ra các đối tượng (**Objects**). Một lớp định nghĩa các **thuộc tính (attributes)** và **phương thức (methods)** mà đối tượng của lớp đó sẽ có.

Ví dụ về một lớp **XeHơi** trong Java:
```java
public class XeHoi {
    String hangXe; // Thuộc tính
    int namSanXuat;

    void chay() { // Phương thức
        System.out.println("Xe đang chạy...");
    }
}
```
Lớp này chỉ là một **mô tả**, nó không chiếm bộ nhớ cho đến khi chúng ta tạo ra một đối tượng.

---
## **2. Object là gì?**

Đối tượng (**Object**) là một thể hiện cụ thể của một lớp. Khi một đối tượng được tạo ra, nó chiếm không gian bộ nhớ và có giá trị riêng của nó.

Ví dụ tạo một đối tượng từ lớp `XeHoi`:
```java
public class Main {
    public static void main(String[] args) {
        XeHoi toyota = new XeHoi(); // Tạo đối tượng từ lớp XeHoi
        toyota.hangXe = "Toyota";
        toyota.namSanXuat = 2023;
        toyota.chay(); // Gọi phương thức chay()
    }
}
```
📌 **Kết quả:**
```java
Xe đang chạy...
```

---
## **3. Sự khác biệt giữa Class và Object**

|**Tiêu chí**|**Class (Lớp)**|**Object (Đối tượng)**|
|---|---|---|
|**Bản chất**|Bản thiết kế (Blueprint)|Thể hiện cụ thể của class|
|**Chiếm bộ nhớ**|Không|Có|
|**Cách sử dụng**|Định nghĩa các thuộc tính và phương thức|Gọi các phương thức và lưu dữ liệu cụ thể|
|**Số lượng**|Chỉ có một bản thiết kế|Có thể có nhiều đối tượng được tạo từ một lớp|

---
## **4. Sử dụng Class và Object trong và ngoài lớp chính (Main Class)**

### **📌 Bên trong lớp chính (Main Class)**
```java
public class Main {
    static class XeHoi { // Lớp lồng nhau trong Main
        String hangXe;
        void chay() {
            System.out.println("Xe " + hangXe + " đang chạy...");
        }
    }

    public static void main(String[] args) {
        XeHoi toyota = new XeHoi(); // Sử dụng class bên trong Main
        toyota.hangXe = "Toyota";
        toyota.chay();
    }
}
```
Kết quả:
`Xe Toyota đang chạy...`

📌 Ngoài lớp chính (Main Class)
```java
// Lớp XeHoi nằm riêng biệt
public class XeHoi {
    String hangXe;
    void chay() {
        System.out.println("Xe " + hangXe + " đang chạy...");
    }
}

// Lớp Main
public class Main {
    public static void main(String[] args) {
        XeHoi toyota = new XeHoi(); // Tạo đối tượng từ lớp khác
        toyota.hangXe = "Toyota";
        toyota.chay();
    }
}
```
📌 **Kết quả giống nhau**, nhưng tổ chức mã tốt hơn, dễ bảo trì hơn.

---
## **5. Class và Object có thể tồn tại riêng lẻ không?**

- **Lớp (Class) không thể hoạt động nếu không có đối tượng** vì nó chỉ là một bản thiết kế.
- **Đối tượng không thể tồn tại nếu không có lớp** vì đối tượng được tạo ra từ một lớp.

⏩ **Tóm lại:** **Lớp và đối tượng luôn đi đôi với nhau.** Một lớp có thể tồn tại mà không cần tạo đối tượng ngay, nhưng đối tượng **bắt buộc** phải có lớp.

---
## **6. Ví dụ thực tế về Class và Object**

**Ví dụ:** Lớp `XeHoi` là bản thiết kế, các xe cụ thể là đối tượng:

|**Class (Lớp)**|**Objects (Đối tượng)**|
|---|---|
|`XeHoi`|`Toyota Camry`, `Honda Civic`, `Ford Ranger`|
|`Nguoi`|`Nguyễn Văn A`, `Trần Thị B`|
|`Laptop`|`MacBook Pro`, `Dell XPS`, `HP Pavilion`|

---
## **7. Các phạm vi truy cập (Access Modifiers)**

Access Modifiers kiểm soát phạm vi truy cập của thuộc tính và phương thức trong lớp.

|**Modifier**|**Phạm vi truy cập**|
|---|---|
|`public`|Có thể truy cập từ bất kỳ đâu|
|`private`|Chỉ có thể truy cập trong cùng một lớp|
|`protected`|Truy cập trong cùng package và lớp con|
|_(mặc định - default)_|Chỉ có thể truy cập trong cùng package|
```java
public class XeHoi {
    private String bienSo; // Chỉ có thể truy cập trong class này
    public String hangXe; // Có thể truy cập từ bất kỳ đâu

    public void setBienSo(String bienSo) {
        this.bienSo = bienSo; // Sử dụng getter/setter để truy cập private
    }

    public String getBienSo() {
        return bienSo;
    }
}
```

```java
public class Main {
    public static void main(String[] args) {
        XeHoi xe = new XeHoi();
        xe.hangXe = "Toyota"; // OK vì public
        xe.setBienSo("51A-12345"); // OK vì dùng setter
        System.out.println("Biển số: " + xe.getBienSo()); // OK
    }
}
```

📌 **Kết quả:**
`Biển số: 51A-12345`

**Lưu ý:** Không thể truy cập `bienSo` trực tiếp vì nó là `private`.

---
## **8. Hàm thành viên (Member Functions - Inner & Outer Class Function)**

Hàm thành viên là các phương thức bên trong một lớp, chia làm hai loại chính:

### **1️⃣ Hàm thành viên của lớp bên trong (Inner Class Function)**

Lớp bên trong (`Inner Class`) là lớp được khai báo bên trong một lớp khác.

```java
class XeHoi {
    class DongCo { // Lớp bên trong
        void chay() {
            System.out.println("Động cơ đang chạy...");
        }
    }
}
public class Main {
    public static void main(String[] args) {
        XeHoi xe = new XeHoi();
        XeHoi.DongCo dongCo = xe.new DongCo(); // Truy cập lớp bên trong
        dongCo.chay();
    }
}
```

`Động cơ đang chạy...`

2️⃣ Hàm thành viên của lớp bên ngoài (Outer Class Function)

```java
class XeHoi {
    void chay() { // Phương thức trong lớp chính
        System.out.println("Xe hơi đang chạy...");
    }
}
public class Main {
    public static void main(String[] args) {
        XeHoi toyota = new XeHoi();
        toyota.chay();
    }
}
```
⏩ **Tóm lại:**

- **Outer Class Functions** là các phương thức bình thường trong lớp chính.
- **Inner Class Functions** là các phương thức nằm trong lớp con bên trong một lớp khác.
