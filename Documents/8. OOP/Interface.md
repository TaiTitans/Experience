
---

**Interface** (giao diện) là một khái niệm quan trọng trong lập trình hướng đối tượng (OOP) Java, đóng vai trò như một bản thiết kế cho các lớp. Nó **không thể tạo đối tượng trực tiếp** nhưng **xác định các phương thức trừu tượng** mà các lớp triển khai (implement) phải thực thi.


**Đặc điểm của Interface:**

- **Chỉ chứa các phương thức trừu tượng:** Interface không có phương thức cụ thể (đã được định nghĩa), chỉ có các phương thức trừu tượng (abstract) mà các lớp triển khai phải định nghĩa.
- **Có thể khai báo các hằng số:** Interface có thể khai báo các hằng số public static final, giống như lớp.
- **Có thể kế thừa từ các interface khác:** Interface có thể kế thừa từ các interface khác để mở rộng chức năng.
- **Một lớp có thể triển khai nhiều interface:** Một lớp có thể triển khai nhiều interface cùng lúc, miễn là nó thực thi tất cả các phương thức trừu tượng của các interface đó.

**Lợi ích của Interface:**

- **Tăng tính trừu tượng:** Interface giúp tập trung vào hành vi chung của các lớp, thay vì chi tiết cụ thể của từng lớp.
- **Tăng tính linh hoạt:** Interface cho phép các lớp có thể thực hiện các hành vi khác nhau theo cách thống nhất.
- **Tăng tính tái sử dụng:** Interface cho phép tái sử dụng mã code cho các lớp khác nhau.
- **Thúc đẩy tính đa hình:** Interface giúp thúc đẩy đa hình thời gian chạy, cho phép xử lý các đối tượng khác nhau theo hành vi cụ thể của chúng.


### Cách sử dụng Interface:

**Khai báo Interface:** Sử dụng từ khóa `interface` để khai báo interface, sau đó là tên interface và các phương thức trừu tượng.

```Java
interface HinhHoc {
  void tinhDienTich();
  void chuVi();
}
```

**Triển khai Interface:** Sử dụng từ khóa `implements` trong lớp để triển khai interface. Lớp phải định nghĩa tất cả các phương thức trừu tượng của interface.

```Java
class HinhVuong implements HinhHoc {
  private int canh;

  public HinhVuong(int canh) {
    this.canh = canh;
  }

  @Override
  public void tinhDienTich() {
    System.out.println("Tính diện tích hình vuông: " + canh * canh);
  }

  @Override
  public void chuVi() {
    System.out.println("Tính chu vi hình vuông: " + 4 * canh);
  }
}

```
---
## Interface tĩnh và Interface động.

**Interface tĩnh (Static Interface):**

- Là một interface được **khai báo trực tiếp trong một lớp** thay vì được khai báo riêng biệt như interface thông thường.
- Interface tĩnh **không thể được kế thừa bởi các lớp khác**.
- Interface tĩnh **chỉ có thể được sử dụng bởi các thành viên tĩnh** của lớp chứa nó.
VD:
```Java
class MyClass {
  public static interface StaticInterface {
    void doSomething();
  }

  public static void main(String[] args) {
    StaticInterface staticInterface = new StaticInterface() {
      @Override
      public void doSomething() {
        System.out.println("Doing something from static interface");
      }
    };

    staticInterface.doSomething();
  }
}

```

**Interface động (Dynamic Interface):**

- Là một interface được **khai báo riêng biệt** như interface thông thường.
- Interface động **có thể được kế thừa bởi các lớp khác**.
- Interface động **có thể được sử dụng bởi cả thành viên tĩnh và phi tĩnh** của lớp.

```Java
interface DynamicInterface {
  void doSomething();
}

class MyClass implements DynamicInterface {
  @Override
  public void doSomething() {
    System.out.println("Doing something from dynamic interface");
  }

  public static void main(String[] args) {
    DynamicInterface dynamicInterface = new MyClass();
    dynamicInterface.doSomething();
  }
}

```

---

**MORE**:

## Top 5 câu hỏi về Interface khi đi phỏng vấn lập trình Java

1. **Interface là gì và nó khác gì so với Abstract class?**
    
    - Interface chỉ chứa các phương thức trừu tượng và hằng số, không có phương thức cụ thể. Abstract class có thể chứa cả phương thức trừu tượng và phương thức cụ thể.
    - Interface không thể tạo đối tượng trực tiếp, abstract class có thể tạo đối tượng trực tiếp.
    - Một lớp có thể triển khai nhiều interface nhưng chỉ kế thừa từ một abstract class.
2. **Nêu ra các lợi ích khi sử dụng Interface trong lập trình Java.**
    
    - Tăng tính trừu tượng, tập trung vào hành vi chung.
    - Tăng tính linh hoạt, cho phép các lớp thực hiện hành vi khác nhau theo cách thống nhất.
    - Tăng tính tái sử dụng mã code cho các lớp khác nhau.
    - Thúc đẩy tính đa hình thời gian chạy.
3. **Giải thích cách thức khai báo và triển khai Interface trong Java.**
    
    - **Khai báo Interface:** Sử dụng từ khóa `interface` sau đó là tên interface và các phương thức trừu tượng.
    - **Triển khai Interface:** Sử dụng từ khóa `implements` trong lớp để triển khai interface. Lớp phải định nghĩa tất cả các phương thức trừu tượng của interface.
4. **Nêu ví dụ về cách sử dụng Interface trong thực tế.**
    
    - Thiết kế giao diện người dùng (GUI): Interface được sử dụng để tạo các widget có thể hiển thị và tương tác với người dùng theo nhiều cách khác nhau.
    - Xử lý dữ liệu: Interface được sử dụng để xử lý các loại dữ liệu khác nhau theo cùng một cách.
    - Lập trình mạng: Interface được sử dụng để giao tiếp với các loại mạng khác nhau theo cùng một cách.
5. **Ngoài ra, bạn có thể được hỏi về các chủ đề liên quan đến Interface như:**
    
    - Giải thích sự khác biệt giữa Interface tĩnh (static) và Interface động (dynamic).
    - Giải thích cách sử dụng Interface làm tham số cho phương thức và lớp.
    - Giải thích cách sử dụng Interface để ẩn đi sự phức tạp của việc triển khai.

