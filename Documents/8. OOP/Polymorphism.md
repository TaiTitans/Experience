

---

	The word polymorphism means having many forms. In simple words, we can define Java Polymorphism as the ability of a message to be displayed in more than one form.

VN:
Đa hình trong thế giới thực:
	Một người cùng một lúc có thể có những đặc điểm khác nhau. Giống như một người đàn ông vừa là người cha, người chồng, vừa là người làm công. Vì vậy, cùng một người sở hữu những hành vi khác nhau trong những tình huống khác nhau. Điều này được gọi là đa hình.

Polymorphism is considered one of the important features of Object-Oriented Programming. Polymorphism allows us to perform a single action in different ways. In other words, polymorphism allows you to define one interface and have multiple implementations. The word “poly” means many and “morphs” means forms, So it means many forms.

## ***Types of Java Polymorphism***

In Java Polymorphism is mainly divided into two types: 

- Compile-time Polymorphism
- Runtime Polymorphism

## Compile-Time Polymorphism in Java

It is also known as static polymorphism. This type of polymorphism is achieved by function overloading or operator overloading.

VN: 
Nó còn được gọi là đa hình tĩnh. Kiểu đa hình này đạt được bằng cách nạp chồng hàm hoặc nạp chồng toán tử.

**Note:*** But Java doesn’t support the Operator Overloading.

![](http://media.geeksforgeeks.org/wp-content/uploads/OverridingVsOverloading.png)

### ***Method Overloading***

Khi có nhiều hàm có cùng tên nhưng khác tham số thì các hàm này được cho là bị quá tải. Các hàm có thể bị quá tải do thay đổi số lượng đối số hoặc/và thay đổi loại đối số.

---
## [Runtime Polymorphism in Java](https://www.geeksforgeeks.org/dynamic-method-dispatch-runtime-polymorphism-java/)

**Đa hình thời gian chạy** là khả năng quyết định phương thức thực sự sẽ được gọi vào **lúc chạy** (runtime) chứ không phải lúc biên dịch (compile time). Điều này đạt được thông qua **ghi đè phương thức (method overriding)**.


VD: 
```Java
class Animal {
  public void speak() {
    System.out.println("Animal sound");
  }
}

class Dog extends Animal {
  @Override
  public void speak() {
    System.out.println("Woof!");
  }
}

class Cat extends Animal {
  @Override
  public void speak() {
    System.out.println("Meow!");
  }
}

```


**Tại sao runtime polymorphism lại hữu ích?**

- **Linh hoạt hơn:** Cho phép bạn xử lý các đối tượng khác nhau theo hành vi cụ thể của chúng.
- **Tái sử dụng mã tốt hơn:** Bạn có thể viết phương thức chung trong lớp cha và ghi đè nó trong các lớp con để thực hiện hành vi cụ thể.
- **Mã trừu tượng hơn:** Tập trung vào hành vi chung (speak()) thay vì chi tiết cụ thể (Woof! hoặc Meow!).

---
Tóm lại:
Đa hình trong OOP bao gồm 2 phương thức: Nạp chồng (overloading) và Ghi đè (Overriding).

VD Overloading:
```Java
class HinhHoc {
  void tinhDienTich() {
    System.out.println("Tính diện tích hình học mặc định");
  }

  void tinhDienTich(int canh) {
    System.out.println("Tính diện tích hình vuông: " + canh * canh);
  }

  void tinhDienTich(int chieuDai, int chieuRong) {
    System.out.println("Tính diện tích hình chữ nhật: " + chieuDai * chieuRong);
  }
}

```

VD Overriding:

```Java
class HinhVuong extends HinhHoc {
  @Override
  void tinhDienTich(int canh) {
    System.out.println("Tính diện tích hình vuông: " + canh * canh);
  }
}

```

**Lợi ích của đa hình:**

- **Tăng tính linh hoạt:** Đa hình cho phép viết mã linh hoạt hơn, dễ dàng thích ứng với các thay đổi về yêu cầu.
- **Tăng tính tái sử dụng:** Đa hình cho phép tái sử dụng mã code cho các lớp con khác nhau, giúp tiết kiệm thời gian và công sức.
- **Tăng tính trừu tượng:** Đa hình giúp tập trung vào hành vi chung của các đối tượng, thay vì chi tiết cụ thể của từng loại đối tượng, giúp mã code dễ hiểu và bảo trì hơn.

