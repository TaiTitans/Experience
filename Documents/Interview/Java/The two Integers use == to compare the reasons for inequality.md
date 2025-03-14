
---
```java
Integer a = 100;
Integer b = 100;
System.out.println(a == b); // true

Integer c = 200;
Integer d = 200;
System.out.println(c == d); // false
```

### Giải thích:

Trong Java, khi bạn sử dụng == để so sánh hai biến kiểu Integer (một kiểu bao bọc), bạn đang so sánh **tham chiếu** (reference) của chúng, tức là kiểm tra xem hai biến có trỏ đến cùng một đối tượng trong bộ nhớ hay không, chứ không phải so sánh giá trị thực tế của chúng. Tuy nhiên, kết quả lại khác nhau giữa hai trường hợp trên do cơ chế **Integer caching** (bộ nhớ đệm của Integer) trong Java.

#### 1. Integer Caching (Bộ nhớ đệm Integer):

- Java có một cơ chế tối ưu hóa gọi là **Integer caching** cho các giá trị Integer trong khoảng từ **-128 đến 127**.
- Khi bạn tạo một biến Integer trong phạm vi này bằng cách sử dụng autoboxing (tự động đóng gói từ int sang Integer, như Integer a = 100), Java không tạo một đối tượng Integer mới mà sử dụng lại các đối tượng đã được lưu trong bộ nhớ đệm. Điều này giúp tiết kiệm bộ nhớ và tăng hiệu suất.
- Do đó, tất cả các biến Integer trong khoảng -128 đến 127 được gán giá trị bằng autoboxing sẽ trỏ đến **cùng một đối tượng** trong bộ nhớ.

#### 2. Trường hợp a == b:

- Integer a = 100; và Integer b = 100; đều nằm trong phạm vi bộ nhớ đệm (-128 đến 127).
- Khi trình biên dịch thực hiện autoboxing, cả a và b đều trỏ đến cùng một đối tượng Integer có giá trị 100 trong bộ nhớ đệm.
- Vì == so sánh tham chiếu, và a với b cùng trỏ đến một đối tượng, nên a == b trả về **true**.

#### 3. Trường hợp c == d:

- Integer c = 200; và Integer d = 200; nằm ngoài phạm vi bộ nhớ đệm (-128 đến 127).
- Khi giá trị vượt quá phạm vi này, Java không sử dụng bộ nhớ đệm mà tạo ra **hai đối tượng Integer riêng biệt** trong bộ nhớ heap cho c và d, mặc dù chúng có cùng giá trị 200.
- Vì == so sánh tham chiếu, và c với d trỏ đến hai đối tượng khác nhau trong bộ nhớ, nên c == d trả về **false**.

#### 4. Minh họa:

- Với a và b: Cả hai cùng trỏ đến một đối tượng Integer@100 trong bộ nhớ đệm → a == b là true.
- Với c và d: c trỏ đến Integer@200 và d trỏ đến Integer@200 (hai địa chỉ khác nhau) → c == d là false.

#### 5. Lưu ý:

- Nếu bạn muốn so sánh **giá trị** của Integer thay vì tham chiếu, hãy sử dụng phương thức equals():
```java
System.out.println(c.equals(d)); // true
```

- Kết quả sẽ là true vì equals() so sánh giá trị thực tế (200 = 200), không phải tham chiếu.
- Nếu bạn tạo Integer bằng từ khóa new (ví dụ: Integer c = new Integer(200);), bộ nhớ đệm sẽ không được sử dụng, và mỗi lần đều tạo một đối tượng mới.