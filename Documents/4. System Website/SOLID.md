# SOLID - Là gì ? Tại sao cần SOLID để code tốt hơn ?

## SOLID là gì ?

Giải thích:

> Phần mềm được xem là tốt khi khi nó có kiến trúc tốt. Kiến trúc phần mềm tương tự như móng nhà, móng yếu nhà sẽ không vững. Để viết được phần mềm tốt bạn phải học rất nhiều, điều đầu tiên bạn cần biết là **SOLID**.

**SOLID** là viết tắt của 5 chữ cái đầu trong 5 nguyên tắc thiết kế hướng đối tượng. Giúp cho lập trình viên viết ra những đoạn code dễ đọc, dễ hiểu, dễ maintain. Nó được đưa ra bởi [Robert C. Martin](http://www.goodreads.com/author/show/45372.Robert_C_Martin) và Michael Feathers. 5 nguyên tắc đó bao gồm:

- **S**ingle responsibility priciple (SRP)
- **O**pen/Closed principle (OCP)
- **L**iskov substitution principe (LSP)
- **I**nterface segregation principle (ISP)
- **D**ependency inversion principle (DIP)

## SOLID ra đời như thế nào ?

Trong OOP bao gồm:

- **Tính trừu tượng** (abstraction): Tạo ra các lớp trừu tượng mô hình hoá các đối tượng trong thế giới thực.
- **Tính đóng gói** (Encapsulation): Các thực thể của lớp trừu tượng có các giá trị thuộc tính riêng biệt.
- **Tính kế thừa** (Inheritance): Các đối tượng có thể dễ dàng kế thừa và mở rộng lẫn nhau.
- **Tính đa hình** (Polymorphism): Có thể thực hiện một hành động đơn theo nhiều cách thức khác nhau tuỳ theo loại đối tượng cụ thể đang được gọi.

Vậy tóm lại SOLID liên quan gì đến OOP ?

-> Hầu hết lập trình viên đều đã biết các tính chất này của OOP, nhưng cách thức để phối hợp các tính chất này với nhau để tăng hiệu quả của ứng dụng thì không phải ai cũng nắm được. Một trong những chỉ dẫn để giúp chúng ta sử dụng được OOP hiệu quả hơn đó là nguyên tắc **SOLID**.

------------

### Single responsibility

> Mỗi lớp chỉ nên chịu trách nhiệm về một nhiệm vụ cụ thể.

Có ý nghĩa là một class chỉ nên giữ một trách nhiệm duy nhất. Một class có quá nhiều chức năng sẽ trở nên cồng kềnh và trở nên khó đọc, khó maintain. Mà đối với ngành IT việc requirement thay đổi, cần thêm sửa chức năng là rất bình thường, nên việc code trong sáng, dễ đọc dễ hiểu là rất cần thiết.

### Open/Closed principle

> Không được sửa đổi một Class có sẵn, nhưng có thể mở rộng bằng kế thừa.

Theo nguyên lý này, mỗi khi ta muốn thêm chức năng cho chương trình, chúng ta nên viết class mới mở rộng class cũ (bằng cách kế thừa hoặc sở hữu class cũ) chứ không nên sửa đổi class cũ. Việc này dẫn đến tình trạng phát sinh nhiều class, nhưng chúng ta sẽ không cần phải test lại các class cũ nữa, mà chỉ tập trung vào test các class mới, nơi chứa các chức năng mới.

### Liskov substitution principle

> Các đối tượng (instance) kiểu class con có thể thay thế các đối tượng kiểu class cha mà không gây lỗi.

Nguyên tắc này nói về quan hệ kế thừa giữa các lớp trong lập trình hướng đối tượng và quy định các điều kiện mà các lớp con cần tuân thủ để có thể thay thế cho lớp cha mà không làm thay đổi tính đúng đắn của chương trình.

### Interface segregation principle

> Thay vì dùng 1 interface lớn, ta nên tách thành nhiều interface nhỏ, với nhiều mục đích cụ thể.

Hãy tưởng tượng chúng ta có 1 interface lớn, khoảng 100 methods. Việc implements sẽ rất vất vả vì các class impliment interface này sẽ bắt buộc phải phải thực thi toàn bộ các method của interface. Ngoài ra còn có thể dư thừa vì 1 class không cần dùng hết 100 method. Khi tách interface ra thành nhiều interface nhỏ, gồm các method liên quan tới nhau, việc implement và quản lý sẽ dễ hơn.

### Dependency inversion principle

> ```
> 1.Các module cấp cao không nên phụ thuộc vào các modules cấp thấp. Cả 2 nên phụ thuộc vào abstraction.
> 2.Interface (abstraction) không nên phụ thuộc vào chi tiết, mà ngược lại (Các class giao tiếp với nhau thông qua interface (abstraction), không phải thông qua implementation.)
> ```

Có thể hiểu nguyên lí này như sau: những thành phần trong 1 chương trình chỉ nên phụ thuộc vào những cái trừu tượng (abstraction). Những thành phần trừu tượng không nên phụ thuộc vào các thành phần mang tính cụ thể mà nên ngược lại.

Những cái trừu tượng (abstraction) là những cái ít thay đổi và biến động, nó tập hợp những đặc tính chung nhất của những cái cụ thể. Những cái cụ thể dù khác nhau thế nào đi nữa đều tuân theo các quy tắc chung mà cái trừu tượng đã định ra. Việc phụ thuộc vào cái trừu tượng sẽ giúp chương trình linh động và thích ứng tốt với các sự thay đổi diễn ra liên tục.

## Tổng kế

SOLID Giúp lập trình viên:

1. Rõ ràng dễ hiểu
2. Dễ thay đổi
3. Tái sử dụng

-----

[Giới thiệu về quy tắc SOLID: Kĩ năng cần có của một dev trình cao (youtube.com)](https://www.youtube.com/watch?v=pqiIdCFyfkU)