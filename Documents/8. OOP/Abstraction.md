
---

**Abstraction in Java*** is the process in which we only show essential details/functionality to the user. The non-essential implementation details are not displayed to the user.


## What is Abstraction in Java?

In Java, abstraction is achieved by [***interfaces****](https://www.geeksforgeeks.org/interfaces-in-java/) and [****abstract classes****](https://www.geeksforgeeks.org/abstract-classes-in-java/). We can achieve 100% abstraction using interfaces.

Data Abstraction may also be defined as the process of identifying only the required characteristics of an object ignoring the irrelevant details. The properties and behaviours of an object differentiate it from other objects of similar type and also help in classifying/grouping the objects.

----
## ***Java Abstract classes and Java Abstract methods***

1. An abstract class is a class that is declared with an [abstract keyword.](https://www.geeksforgeeks.org/abstract-keyword-in-java/)
2. An abstract method is a method that is declared without implementation.
3. An abstract class may or may not have all abstract methods. Some of them can be concrete methods
4. A method-defined abstract must always be redefined in the subclass, thus making [overriding](https://www.geeksforgeeks.org/overriding-in-java/) compulsory or making the subclass itself abstract.
5. Any class that contains one or more abstract methods must also be declared with an abstract keyword.
6. There can be no object of an abstract class. That is, an abstract class can not be directly instantiated with the [__new operator__](https://www.geeksforgeeks.org/new-operator-java/).
7. An abstract class can have parameterized constructors and the default constructor is always present in an abstract class.

---
VN:
**Abstraction** (trừu tượng hóa) là một khái niệm quan trọng trong lập trình hướng đối tượng (OOP), giúp đơn giản hóa chương trình bằng cách tập trung vào các tính năng thiết yếu và che giấu đi những chi tiết phức tạp bên trong. Trong Java, abstraction được thực hiện thông qua hai công cụ chính: **lớp trừu tượng** và **giao diện** (interface).

**Lớp trừu tượng (Abstract class)**:

- Là lớp có chứa ít nhất một **phương thức trừu tượng** (abstract method).
- Phương thức trừu tượng không có phần thân thực thi, chỉ khai báo tên và kiểu trả về.
- Lớp trừu tượng không thể được khởi tạo trực tiếp.
- Các lớp con kế thừa từ lớp trừu tượng phải **triển khai** (implement) tất cả các phương thức trừu tượng của lớp cha.

```Java
abstract class HinhHoc {
  abstract float tinhDienTich();
}

class HinhVuong extends HinhHoc {
  float canh;

  public HinhVuong(float canh) {
    this.canh = canh;
  }

  @Override
  float tinhDienTich() {
    return canh * canh;
  }
}

```

**Giao diện (Interface):**

- Tương tự như lớp trừu tượng, nhưng chỉ chứa các phương thức trừu tượng, không có thuộc tính hay phương thức cụ thể.
- Lớp con phải triển khai tất cả các phương thức trừu tượng được khai báo trong giao diện.
- Một lớp có thể kế thừa từ nhiều giao diện.
```Java
interface HinhVe {
  void ve();
}

class HinhTron implements HinhVe {
  float banKinh;

  public HinhTron(float banKinh) {
    this.banKinh = banKinh;
  }

  @Override
  public void ve() {
    System.out.println("Vẽ hình tròn với bán kính: " + banKinh);
  }
}

```

---

**Sử dụng lớp trừu tượng khi:**

- Bạn muốn cung cấp **phương thức triển khai mặc định** cho một số phương thức trừu tượng. Các lớp con có thể sử dụng hoặc ghi đè lên các phương thức triển khai này.
- Bạn muốn **ẩn đi các chi tiết triển khai** của một số thành phần bên trong lớp. Các lớp con chỉ có thể truy cập các thành viên được khai báo public của lớp cha.
- Bạn muốn **thiết lập một cấu trúc phân cấp** với các lớp con kế thừa từ cùng một lớp cha.

**Sử dụng giao diện khi:**

- Bạn muốn **định nghĩa một tập hợp hành vi** mà các lớp khác phải tuân theo. Giao diện không cung cấp triển khai cụ thể, mà chỉ khai báo các phương thức trừu tượng.
- Bạn muốn **tạo ra sự linh hoạt** cho các lớp con trong việc triển khai các phương thức theo cách riêng của chúng.
- Bạn muốn **kết hợp các hành vi từ nhiều nguồn khác nhau** vào một lớp. Một lớp có thể triển khai nhiều giao diện.

**Ví dụ:**

- Giả sử bạn đang thiết kế một hệ thống quản lý nhân viên. Bạn có thể sử dụng lớp trừu tượng `NhanVien` để cung cấp các phương thức chung như `getTen()`, `getLuong()`, v.v. Các lớp con như `NhanVienHanhChinh`, `NhanVienKyThuat` sẽ kế thừa từ `NhanVien` và triển khai các phương thức cụ thể cho từng loại nhân viên.
- Giả sử bạn đang thiết kế một trò chơi đồ họa. Bạn có thể sử dụng giao diện `VeHinh` để định nghĩa phương thức `ve()`. Các lớp như `HinhTron`, `HinhChuNhat` có thể triển khai giao diện này để hiển thị hình dạng của chúng trên màn hình.

## Tóm lại:

Lựa chọn sử dụng lớp trừu tượng hay giao diện phụ thuộc vào nhu cầu cụ thể của bạn. Lớp trừu tượng phù hợp khi bạn muốn cung cấp triển khai mặc định và thiết lập cấu trúc phân cấp, trong khi giao diện phù hợp khi bạn muốn định nghĩa hành vi chung và tạo sự linh hoạt cho các lớp con.

### **Khi nào chọn cái nào?**

|**Tiêu chí**|**Abstract class**|**Interface**|
|---|---|---|
|**Kế thừa logic chung**|Dùng khi cần mã logic tái sử dụng|Không phù hợp cho logic dùng chung (trước Java 8)|
|**Đa kế thừa**|Không hỗ trợ|Hỗ trợ|
|**Mối quan hệ**|"is-a" (là một loại gì đó)|"can-do" (có thể làm gì đó)|
|**Ưu tiên thiết kế**|Dùng để tạo cấu trúc cơ sở|Dùng để định nghĩa hành vi|
