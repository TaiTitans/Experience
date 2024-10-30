
---
## 1. Lập trình hướng đối tượng là gì? Các tính chất của lập trình hướng đối tượng?

- Gợi ý:
    - Lập trình hướng đối tượng gồm có 4 tính chất:
        - Tính đóng gói
        - Tính kế thừa
        - Tính trừu tượng
        - Tính đa hình
    - **Để câu trả lời thuyết phục hơn, ta cần trình bày lý do tại sao cần những tính chất này**

- Tham khảo:
    - Lập trình hướng đối tượng (OOP - Object-Oriented Programming) là một phương pháp lập trình mà trong đó, các chương trình được tổ chức xung quanh các đối tượng (objects). Mỗi đối tượng là một thể hiện của một lớp (class), và lớp định nghĩa cách mà các đối tượng thuộc về lớp đó sẽ hoạt động.
    - Các tính chất của lập trình hướng đối tượng:
        - Tính đóng gói (Encapsulation):
            - Đóng gói giúp che dấu thông tin và triển khai chi tiết bên trong của một đối tượng, chỉ cho phép truy cập thông qua các phương thức public.
            - Bảo vệ dữ liệu khỏi sự thay đổi trái phép.
        - Tính kế thừa (Inheritance):
            - Kế thừa cho phép một lớp (class) sử dụng lại các thuộc tính và phương thức của một lớp khác.
            - Giúp giảm sự lặp lại mã và làm tăng tính linh hoạt của mã nguồn.
        - Tính trừu tượng (Abstraction):
            - Trừu tượng giúp ẩn đi các chi tiết phức tạp và chỉ tập trung vào các khía cạnh quan trọng của đối tượng.
            - Giúp giảm độ phức tạp của mã nguồn và tăng cường sự dễ hiểu.
            - Tính trừu tượng được thể hiện thông qua interface hoặc abtract class trong Java.
        - Tính đa hình (Polymorphism):
            - Tính đa hình là một khái niệm quan trọng trong lập trình hướng đối tượng, nơi một đối tượng, một hàm có thể thể hiện nhiều hình thái khác nhau.
            - Tính đa hình có thể thực hiện thông qua kỹ thuật như overloading và overriding, upcasting và downcasting.
            - Tính đa hình giúp phần mềm trở nên linh hoạt, dễ mở rộng


## 2. Phân biệt hai cách so sánh == và equals()?

- Gợi ý:
    - `==`: Là một **toán tử** so sánh tham chiếu của hai biến
    - `equals()`: Là một **hàm dùng** để so sánh giá trị của hai biến

- Tham khảo:
    - Giống nhau: Equals() và “==” trong Java đều được sử dụng để so sánh các đối tượng nhằm kiểm tra sự bằng nhau, nhưng giữa chúng có sự khác biệt
    - Khác nhau:
        - Equals là một phương thức, còn == là một toán tử.
        - Toán tử == để so sánh tham chiếu (so sánh địa chỉ) và phương thức equals() để so sánh nội dung. Nói một cách đơn giản, == kiểm tra xem cả hai đối tượng có trỏ đến cùng một vị trí bộ nhớ hay không trong khi equals() sẽ so sánh các giá trị trong các đối tượng. Bất cứ khi nào chúng ta tạo một đối tượng bằng toán tử new, Java sẽ tạo một vị trí bộ nhớ mới cho đối tượng đó. Vì vậy chúng ta sử dụng toán tử == để kiểm tra vị trí bộ nhớ hoặc địa chỉ của 2 đối tượng có giống nhau hay không.
        - Nếu một lớp không ghi đè (override) phương thức equals() thì theo mặc định, nó sử dụng phương thức equals() của lớp cha gần nhất đã ghi đè phương thức này.


## 3. Trong Java, có mấy loại phạm vi truy cập (access modifiers)? Nêu đặc điểm của từng loại?

- Gợi ý:
    - Có 4 loại phạm vi truy cập bao gồm private, default, protected, public
- Tham khảo:
    - Private: Có phạm vi truy cập ít nhất, các thuộc tính, phương thức được đặt private thì chỉ được truy cập bên trong class đã khai báo nó.
	- Default: Một lớp, thuộc tính, phương thức được đặt default thì chỉ được truy cập bởi các lớp nằm trong cùng package.
	- Protected: Được mở rộng từ default, các thuộc tính, phương thức được đặt là protected thì chỉ được truy cập bởi các lớp nằm trong cùng package và bên trong lớp con.
	- Public: Có phạm truy cập nhiều nhất, các thuộc tính, phương thức được đặt là public thì sẽ được truy cập ở mọi nơi trong project.
## 4. Nêu đặc điểm của các biến, phương thức, lớp có từ khóa final?

- Gợi ý:
    - Biến final: Giá trị không thể thay đổi sau khi khởi tạo giá trị của biến.
    - Phương thức final: Phương thức này sẽ không thể được ghi đè (override) bởi một class con nào đó.
    - Lớp final: Lớp final sẽ không cho phép các lớp khác kế thừa nó.

## 5. Phân biệt giữa override và overload?

- Gợi ý:
    - Overload: Là một phương pháp tạo ra hai hoặc nhiều phương thức có cùng tên giống nhau, nhưng khác tham số.
    - Override: Là một phương pháp triển khai lại logic của một hàm nào đó của lớp cha trong lớp con.

- Tham khảo:
    - Overload (Nạp chồng):
        - Overloading là quá trình tạo ra nhiều phương thức có cùng tên trong một lớp nhưng có các danh sách tham số khác nhau.
        - Các phương thức được overload phải có tên giống nhau nhưng phải có danh sách tham số khác nhau, có thể khác kiểu dữ liệu hoặc số lượng tham số.
        - Overriding xảy ra khi một lớp con cung cấp triển khai mới cho một phương thức đã được định nghĩa trong lớp cha.
        - Các phương thức được ghi đè phải có cùng chữ ký (cùng tên, cùng kiểu trả về, và cùng danh sách tham số) như phương thức trong lớp cha.
        - Mục đích của overriding là cung cấp một triển khai mới cho phương thức đã tồn tại trong lớp cha để thích ứng với yêu cầu cụ thể của lớp con.

## 6. Sự khác nhau StringBuilder và StringBuffer trong Java?

- Gợi ý:
    - Hiệu suất chương trình khi sử dụng
    - Có an toàn trong môi trường đa luồng (thread-safe) không

- Tham khảo:
    - StringBuilder và StringBuffer đều giống nhau và có các phương thức tương tự nhưng điểm khác biệt là StringBuffer được đồng bộ (synchronized) và an toàn trong môi trường đa luồng (thread-safe), có thể có nhiều luồng cùng truy cập vào một đối tượng StringBuffer cùng lúc. StringBuilder không được đồng bộ (non-synchronized ) và nên sử dụng trong đơn luồng hoặc biến địa phương (local variable) trong một hàm (method).
    - Vì StringBuilder không cần đồng bộ nên hiệu suất sẽ nhanh hơn StringBuffer

## 7. Trong Java, interface dùng để làm gì?

- Gợi ý:
    - Tính đa kế thừa
    - Đạt được tính trừu tượng (Abstraction)
    - Đạt được sự phụ thuộc lỏng lẻo (Loose coupling)
    - Java không hỗ trợ đa kế thừa vì có thể xảy ra [Diamond Problem](https://en.wikipedia.org/wiki/Multiple_inheritance?ref=roninhub.com#The_diamond_problem) (Khi class D kế thừa từ 2 class B và C cùng override lại hàm print() của A, thì khi đó class D sẽ không biết phải override lại hàm print() của class nào). Interface sẽ giải quyết được bài toán đa kế thừa. Nếu một class D implement 3 Interface B, C và D có cùng một hàm là print(), thì class A sẽ chỉ có một implement của hàm print(), đây là một đặc điểm của Interface(có khả năng thay đổi hành vi tại thời điểm runtime)
    - Sử dụng Interface sẽ đạt tính trừu tượng vì đặc điểm của Interface chỉ có các public abstract method nên khi một class implement Interface thì phải implement lại toàn bộ các method của Interface đó.
    - Sử dụng Interface để giảm bớt sự phụ thuộc với nhau. Giúp các lớp hoạt động độc lập mà không cần biết triển khai của lớp khác.
## 8. Tại sao Java không hoàn toàn là ngôn ngữ hướng đối tượng (OOP)?

- Gợi ý:
    - Java có các kiểu dữ liệu khác ngoài kiểu dữ liệu object không?
- Tham khảo:
    - Java không phải là ngôn ngữ hướng đối tượng (OOP) 100% vì ngoài các kiểu dữ liệu đối tượng (Object data type) thì Java còn có 8 kiểu dữ liệu nguyên thủy(Primitive data type) là: char, boolean, byte, short, int, long, float, double.

## 9. Wrapper class trong Java là gì? Tại sao chúng ta cần chúng?

- Gợi ý:
    - Làm sao để sử dụng các kiểu dữ liệu nguyên thủy (primitive) trong trường hợp đối tượng cần sử dụng không hỗ trợ như Collection(List, Set, Map…).

- Cần phải có dữ liệu có thể null
- Cần sử dụng các hàm của object

- Tham khảo:
    - Wrapper class trong Java dùng để biểu diễn 8 kiểu đối tượng nguyên thủy(primitive data type) trong Java thành kiểu đối tượng(object data type). Tất cả các lớp Wrapper trong Java là bất biến và final.
    - Lý do chúng ta cần có wrapper là:
        - Cho phép dữ liệu có thể null
        - Có thể sử dụng ở Collections như List, Map… vì Collection chỉ hỗ trợ kiểu dữ liệu đối tượng
        - Để có thể sử dụng được các hàm(method) có tham số đầu vào là kiểu đối tượng
        - Có thể sử dụng được các hàm của đối tượng như là: clone(), equals(), hashCode(), toString()

## 10. Java là tham chiếu (pass-by-reference) hay tham trị (pass-by-value)?

- Gợi ý:
    - Trong method có thể thay đổi reference của object đầu vào trỏ tới object khác không?
- Tham khảo:
    - Java là tham trị (pass-by-value). KHÔNG có tham chiếu (pass-by-reference) trong Java

- Java không có khái niệm con trỏ nên tất cả đều là tham trị (pass-by-value).
- Method nhận 1 bản copy của reference tới object

