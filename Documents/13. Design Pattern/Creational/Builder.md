	Tách biệt quá trình khởi tạo đối tượng phức tạp khởi các đại diện của nó

---

## Giới thiệu


Builder Pattern là một Creational Design Pattern cho phép xây dựng đối tượng phức tạp bằng cách sử dụng các đối tượng riêng biệt đại diện cho từng bộ phận cấu thành.

Builder Pattern tách rời quá trình khởi tạo đối tượng phức tạp khỏi các đại diện của nó. Điều này cho phép cùng một quá trình xây dựng có thể tạo ra nhiều biểu diễn khác nhau của đối tượng.

Mục đích: Builder Pattern được sử dụng để tách rời quá trình khởi tạo đối tượng phức tạp khỏi các đại diện của nó, giúp đạt được những lợi ích sau:

- Tăng tính linh hoạt trong khởi tạo đối tượng phức tạp
- Dễ dàng thay đổi cách khởi tạo đối tượng.
- Hỗ trợ tạo nhiều biểu diễn khác nhau của đối tượng.
- Đơn giản hóa việc test và debug.

Builder Pattern tách rời quá trình xây dựng đối tượng phức tạp thành nhiều bước riêng biệt. Mỗi bước tập trung vào một khía cạnh của đối tượng.

Các đại diện chỉ đơn giản lưu trữ kết quả, không cần quan tâm đến quá trình tạo ra chúng.


### Đặt vấn đề


Trong phát triển phần mềm, ta thường gặp các đối tượng phức tạp với nhiều thuộc tính và thành phần. Ví dụ một đối tượng House có thể bao gồm các thành phần như phòng khách, phòng ngủ, nhà bếp, cửa ra vào, cửa sổ, hệ thống điện, nước, và nhiều thành phần khác.

```mermaid
classDiagram

  House "1" *-- "n" Room
  Room : -name
  Room : -size

  House "1" *-- "n" Door
  Door : -width
  Door : -height
  Door : -material

  House "1" *-- "1" Kitchen
  Kitchen : -layout

  House "1" *-- "n" Window
  Window : -size
  Window : -position

  House "1" *-- "1" ElectricalSystem
  ElectricalSystem : -wiring

  House "1" *-- "1" PlumbingSystem
  PlumbingSystem : -pipes

  class House{
    +House(rooms, doors, windows, kitchen, electrical, plumbing)
  }

```

- Quá trình khởi tạo phức tạp, dễ gây nhầm lẫn với nhiều tham số truyền vào
- Các thành phần của `House` bị phụ thuộc lẫn nhau, khó thay đổi một phần mà không ảnh hưởng các thành phần khác.
- Khó tạo các biến thể khác nhau của `House` một cách linh hoạt.

Như vậy, việc xây dựng các đối tượng phức tạp cần được thiết kế cẩn thận để tránh các vấn đề trên.

---
### Giải quyết Vấn Đề


Builder Pattern giúp giải quyết vấn đề của việc tạo ra đối tượng phức tạp, như ví dụ sau về việc xây dựng một ngôi nhà. Thay vì tạo ngôi nhà một cách trực tiếp từng phần, chúng ta chia quá trình này thành nhiều bước riêng biệt. Mỗi bước tập trung vào việc xây dựng một khía cạnh cụ thể của ngôi nhà, chẳng hạn như cửa, cửa sổ và nhà bếp.

#### Cách hoạt động

1. **Director (Quản lý)**: Đầu tiên, chúng ta có một người quản lý, được gọi là Director. Quản lý này có nhiệm vụ chỉ đạo quá trình xây dựng ngôi nhà.
    
2. **Các Concrete Builder (Xây dựng cụ thể)**: Sau đó, chúng ta có các xây dựng cụ thể, ví dụ: Xây dựng Cửa, Xây dựng Cửa sổ và Xây dựng Nhà bếp. Mỗi Concrete Builder chịu trách nhiệm cho việc xây dựng một phần cụ thể của ngôi nhà.
    
3. **Tạo ngôi nhà**: Chúng ta kết hợp các phần đã xây dựng từ các Concrete Builder để tạo ra ngôi nhà hoàn chỉnh

#### Ví dụ minh hoạ


Để hiểu rõ hơn, hãy xem ví dụ sau:

```mermaid
graph LR
  A[Director] --> B[Concrete Builder]
  B --> C[Build Door]
  B --> D[Build Window]
  B --> E[Build Kitchen]
  C --> F[House]
  D --> F
  E --> F

```
- Director (Quản lý) gọi Concrete Builder (Xây dựng cụ thể) để bắt đầu xây dựng ngôi nhà.
    
- Concrete Builder (Xây dựng cụ thể) thực hiện công việc của mình, ví dụ: xây dựng cửa, cửa sổ và nhà bếp.
    
- Các phần này được kết hợp lại để tạo thành ngôi nhà hoàn chỉnh.
    

Kết quả là, người dùng ngôi nhà không cần quan tâm đến chi tiết cụ thể của quá trình xây dựng, mà chỉ cần sử dụng ngôi nhà đã hoàn thành một cách dễ dàng.


### Cấu trúc

Builder Pattern có cấu trúc đơn giản, bao gồm các thành phần sau:

```mermaid
classDiagram
Direction TB

Builder <|-- ConcreteBuilder
Director o-- Builder
Product o-- Director

class Director {
+Construct()
}

class Builder {
+BuildPartA()
+BuildPartB() 
}

class ConcreteBuilder {
+BuildPartA()
+BuildPartB()
}

class Product {
-PartA
-PartB
}


```
- **Builder**: Định nghĩa phương thức xây dựng chung.
- **ConcreteBuilder**: Triển khai chi tiết các bước xây dựng cụ thể.
- **Director**: Sử dụng Builder để xây dựng sản phẩm.
- **Product**: Là sản phẩm được tạo ra, chứa các phần do Builder tạo.

Như vậy Builder tách rời quá trình xây dựng phức tạp thành nhiều bước đơn giản, từng bước tập trung vào một khía cạnh riêng lẻ.

---
## Cách triển khai

Builder Pattern có thể được triển khai theo nhiều cách khác nhau. Trong Java, có một số cách triển khai phổ biến như sau:

**Định nghĩa Product**: Chúng ta bắt đầu bằng việc định nghĩa lớp `Product` để lưu trữ thông tin đối tượng cuối cùng:
```Java
class Product {
    private String partA;
    private String partB;

    public void setPartA(String partA) {
        this.partA = partA;
    }

    public void setPartB(String partB) {
        this.partB = partB;
    }

    public void show() {
        System.out.println("Product Parts: " + partA + " and " + partB);
    }
}

```

**Định nghĩa Builder** : Sau đó, chúng ta định nghĩa lớp `Builder` với các phương thức xây dựng, nhưng với Builder, mỗi phương thức xây dựng trả về chính builder, cho phép chúng ta gọi tiếp theo một cách liền mạch:
```Java
class Builder {
    private Product product = new Product();

    public Builder buildPartA(String partA) {
        product.setPartA(partA);
        return this;
    }

    public Builder buildPartB(String partB) {
        product.setPartB(partB);
        return this;
    }

    public Product getResult() {
        return product;
    }
}
```

Bây giờ, chúng ta có thể sử dụng Builder để xây dựng sản phẩm một cách dễ đọc và gần gũi:
```Java
public class Main {
    public static void main(String[] args) {
        // Sử dụng Builder để xây dựng sản phẩm
        Product product = new Builder()
            .buildPartA("Part A")
            .buildPartB("Part B")
            .getResult();
        
        // Hiển thị sản phẩm
        product.show();
    }
}
```

Giải thích

- Chúng ta đã tạo một lớp `Builder` với các phương thức xây dựng trả về chính builder. Điều này cho phép chúng ta gọi tiếp các phương thức một cách liền mạch, tạo một chuỗi dễ đọc để xây dựng sản phẩm.
- Khi sử dụng Builder, các phương thức xây dựng có thể được gọi nối tiếp trên một đối tượng builder duy nhất, giúp giảm bớt sự phức tạp trong mã nguồn và làm cho mã trở nên rõ ràng hơn.
- Cuối cùng, chúng ta gọi `getResult()` để lấy đối tượng Product đã được xây dựng và hiển thị nó.

---
## So sánh

Builder Pattern có thể được so sánh với một số Design Pattern tương tự, bao gồm:

- **Factory Pattern**: Builder Pattern tập trung vào xây dựng một đối tượng phức tạp bằng cách tạo và cấu hình từng phần, trong khi Factory Pattern tập trung vào việc tạo đối tượng duy nhất và trả về nó.
- **Abstract Factory Pattern**: Cả Builder Pattern và Abstract Factory Pattern đều giúp trong việc tạo đối tượng phức tạp, nhưng Abstract Factory tạo ra một tập hợp các đối tượng có liên quan và cung cấp một giao diện chung để tạo chúng, trong khi Builder tập trung vào việc xây dựng một đối tượng duy nhất.
- **Singleton Pattern**: Singleton Pattern chỉ đảm bảo rằng một lớp chỉ có một đối tượng duy nhất và cung cấp một điểm truy cập đến nó. Builder Pattern không liên quan đến việc tạo đối tượng duy nhất mà tập trung vào việc xây dựng đối tượng phức tạp.

## Kết luận

Builder Pattern là một Design Pattern hữu ích trong những trường hợp cần xây dựng các đối tượng phức tạp. Builder Pattern giúp việc xây dựng các đối tượng phức tạp trở nên dễ dàng hơn và ít xảy ra lỗi hơn.

- Nên sử dụng Builder Pattern khi:
    - Đối tượng phức tạp có nhiều thuộc tính hoặc thành phần.
    - Cần xây dựng nhiều phiên bản khác nhau của đối tượng phức tạp.
    - Cần dễ dàng kiểm tra đối tượng phức tạp.
- Không nên sử dụng Builder Pattern khi:
    - Đối tượng phức tạp không có nhiều thuộc tính hoặc thành phần.
    - Chỉ cần xây dựng một phiên bản duy nhất của đối tượng phức tạp.
    - Không cần dễ dàng kiểm tra đối tượng phức tạp.