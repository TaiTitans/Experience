
---
### **1. `static` Keyword**

Dùng để khai báo một biến, phương thức hoặc khối code mà thuộc về **lớp (class)** chứ không phải đối tượng.
```java
class XeHoi {
    static int soLuongXe = 0; // Biến static (thuộc về class)

    XeHoi() {
        soLuongXe++; // Mỗi lần tạo một đối tượng, giá trị tăng lên
    }

    static void hienThiSoLuongXe() { // Phương thức static
        System.out.println("Số lượng xe: " + soLuongXe);
    }
}

public class Main {
    public static void main(String[] args) {
        XeHoi xe1 = new XeHoi();
        XeHoi xe2 = new XeHoi();
        XeHoi.hienThiSoLuongXe(); // Gọi phương thức static mà không cần tạo object
    }
}
```
📌 **Lưu ý:**

- `static` thuộc về lớp, không phải riêng từng object.
- Phương thức `static` không thể truy cập biến **non-static** của đối tượng.

### **2. `virtual` Keyword** (Không có trong Java, có trong C++)

Trong C++, từ khóa `virtual` dùng để **ghi đè (override)** phương thức trong lập trình đa hình (polymorphism).

📌 **Ví dụ trong C++:**
```java
class Cha {
public:
    virtual void hienThi() { // Khai báo phương thức ảo
        cout << "Phương thức lớp Cha\n";
    }
};

class Con : public Cha {
public:
    void hienThi() override { // Ghi đè phương thức cha
        cout << "Phương thức lớp Con\n";
    }
};

int main() {
    Cha* obj = new Con();
    obj->hienThi(); // Gọi phương thức lớp Con (đa hình)
}
```
📌 **Lưu ý:**

- Java không có từ khóa `virtual`, thay vào đó dùng **`@Override`**.
- Java mặc định mọi phương thức đều có thể ghi đè trừ khi bị khai báo `final`.

### **3. `abstract` Keyword**

Dùng để khai báo **lớp trừu tượng (abstract class)** hoặc **phương thức trừu tượng**.  
Lớp trừu tượng **không thể tạo đối tượng trực tiếp** và thường được dùng làm lớp cha.

📌 **Ví dụ:**
```java
abstract class DongVat {
    abstract void keu(); // Phương thức trừu tượng (không có phần thân)

    void an() { // Phương thức bình thường
        System.out.println("Động vật đang ăn...");
    }
}

class Cho extends DongVat {
    void keu() {
        System.out.println("Gâu gâu!");
    }
}

public class Main {
    public static void main(String[] args) {
        DongVat conCho = new Cho();
        conCho.keu();
        conCho.an();
    }
}
```
📌 **Lưu ý:**

- Nếu một lớp có **ít nhất một phương thức trừu tượng**, thì nó **phải là abstract class**.
- Một lớp kế thừa `abstract class` **phải triển khai tất cả phương thức trừu tượng** hoặc trở thành **abstract class**.

### **4. `final` Keyword**

Dùng để **ngăn chặn thay đổi**, có thể áp dụng cho **biến, phương thức hoặc lớp**.

### **6. `this` Keyword**

Dùng để tham chiếu **đối tượng hiện tại** bên trong lớp.
📌 **Lưu ý:**

- `this` được dùng khi biến cục bộ trùng tên với biến instance.
- Có thể dùng `this` để gọi **constructor khác** trong cùng một lớp.

### **7. `new` Keyword**

Dùng để **tạo đối tượng mới**.
📌 **Lưu ý:**

- `new` cấp phát bộ nhớ trên **heap** cho object.
- Trong C++, `new` còn dùng để cấp phát bộ nhớ động (`int* p = new int;`).

### **9. `super` Keyword**

Dùng để **gọi constructor hoặc phương thức của lớp cha**.

## **Kết luận**

|**Từ khóa**|**Ý nghĩa**|
|---|---|
|`static`|Thành viên thuộc về lớp, không thuộc về object|
|`virtual`|(C++ only) Dùng để ghi đè phương thức cha|
|`abstract`|Khai báo lớp/phương thức trừu tượng|
|`final`|Ngăn chặn thay đổi (biến, phương thức, lớp)|
|`explicit`|(C++ only) Ngăn ép kiểu ngầm định|
|`this`|Tham chiếu đến đối tượng hiện tại|
|`new`|Tạo đối tượng mới|
|`const`|(C++ only) Khai báo biến không thay đổi|
|`super`|Gọi constructor hoặc phương thức lớp cha|

