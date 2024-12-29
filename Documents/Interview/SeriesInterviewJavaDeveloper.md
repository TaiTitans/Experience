
---

- [Tại sao trong Java chỉ có truyền theo giá trị?]
**Truyền tham trị là gì?**

- Khi bạn truyền một giá trị vào một phương thức, Java sẽ tạo **bản sao** của giá trị đó và sử dụng bản sao này trong phương thức.
- Bất kỳ thay đổi nào đối với giá trị bên trong phương thức sẽ không ảnh hưởng đến giá trị gốc bên ngoài.

- Java được thiết kế để đơn giản, dễ hiểu và nhất quán.
- Truyền **tham trị** giúp giảm thiểu lỗi khó theo dõi liên quan đến tham chiếu và bộ nhớ.
- Nó giúp đảm bảo rằng các biến bên ngoài phương thức không bị thay đổi trừ khi thay đổi được thực hiện rõ ràng thông qua trạng thái của đối tượng.
---
- [Giải thích về Seri hóa Java]
	**Serialization** là quá trình chuyển đổi một đối tượng Java thành một chuỗi byte để lưu trữ hoặc truyền qua mạng. Quá trình ngược lại, từ chuỗi byte quay trở lại đối tượng, gọi là **Deserialization**.

- [Giải thích về Generic và Wildcard]
	**Generic** và **Wildcard** là hai khái niệm quan trọng trong Java, giúp bạn viết mã linh hoạt, an toàn và tái sử dụng trong các tình huống xử lý kiểu dữ liệu khác nhau.

**Generic** cho phép bạn định nghĩa kiểu dữ liệu một cách tổng quát, tránh việc lặp lại mã và giảm thiểu lỗi liên quan đến kiểu dữ liệu. Generic được giới thiệu trong Java 5.
```java
List<String> list = new ArrayList<>();
list.add("Hello");
// list.add(123); // Lỗi biên dịch

String value = list.get(0); // Không cần ép kiểu

```

Ưu điểm:

- An toàn kiểu (type safety): Chỉ cho phép một loại dữ liệu xác định.
- Giảm lỗi runtime liên quan đến kiểu.
- Không cần ép kiểu thủ công.

**Wildcard** (`?`) là một ký hiệu đại diện cho một kiểu dữ liệu không xác định. Nó giúp bạn làm việc với các kiểu Generic khi không biết chính xác kiểu dữ liệu tại thời điểm biên dịch.

Các loại Wildcard
	Unbounded Wildcard (`?`) : Dùng khi bạn không quan tâm đến kiểu dữ liệu cụ thể.
	Upper Bounded Wildcard (`? extends T`): Dùng khi bạn muốn làm việc với các kiểu là **T hoặc con của T**.
	Lower Bounded Wildcard (`? super T`): Dùng khi bạn muốn làm việc với các kiểu là **T hoặc cha của T**.

---
- [Giải thích về cơ chế phản chiếu Java Reflection]
	**Reflection** trong Java là một cơ chế mạnh mẽ cho phép chương trình kiểm tra hoặc sửa đổi cấu trúc của nó trong thời gian chạy. Nó nằm trong gói **`java.lang.reflect`** và thường được sử dụng để truy cập thông tin về các lớp, phương thức, thuộc tính, hoặc để tạo đối tượng và gọi phương thức động.

Reflection được hỗ trợ bởi các lớp trong gói `java.lang.reflect`:

1. **`Class`:** Đại diện cho một lớp hoặc interface.
2. **`Field`:** Đại diện cho một trường (biến thành viên) của lớp.
3. **`Method`:** Đại diện cho một phương thức của lớp.
4. **`Constructor`:** Đại diện cho một constructor của lớp.

---
- [Giải thích về mô hình Proxy trong Java]
	**Proxy Pattern** là một mẫu thiết kế cấu trúc (**Structural Design Pattern**) trong đó một lớp thay mặt cho một lớp khác để kiểm soát truy cập đến đối tượng đó. Proxy được sử dụng để cung cấp một "lớp vỏ" xung quanh một đối tượng thực, giúp thêm các chức năng bổ sung như kiểm soát truy cập, caching, logging, hoặc lazy initialization mà không thay đổi bản thân đối tượng thực.

**1. Các loại Proxy**

1. **Remote Proxy**  
    Đại diện cho một đối tượng nằm trên một máy tính khác hoặc hệ thống mạng (ví dụ: sử dụng RMI trong Java).
    
2. **Virtual Proxy**  
    Trì hoãn việc khởi tạo hoặc tải tài nguyên tốn kém cho đến khi thật sự cần thiết.
    
3. **Protection Proxy**  
    Kiểm soát quyền truy cập đối với các phương thức hoặc thuộc tính của đối tượng.
    
4. **Smart Proxy**  
    Cung cấp các hành vi bổ sung, chẳng hạn như logging, đếm số lần truy cập, hoặc cache dữ liệu.

 **2. Cấu trúc Proxy Pattern**

 **Thành phần chính:**

1. **Subject**  
    Một interface hoặc lớp trừu tượng mà cả Proxy và đối tượng thực (Real Subject) đều triển khai.
    
2. **Real Subject**  
    Lớp thực hiện hành vi chính của đối tượng.
    
3. **Proxy**  
    Lớp thay thế đối tượng thực, kiểm soát truy cập và thêm chức năng bổ sung.

---
- [Giải thích về BigDecimal]

	**`BigDecimal`** là một lớp trong Java thuộc gói **`java.math`**, được thiết kế để xử lý các phép tính số học chính xác với số thập phân. Nó thường được sử dụng khi làm việc với các con số yêu cầu độ chính xác cao, như trong ứng dụng tài chính, kế toán, hoặc khoa học.

Lưu ý: BigDecimal là 1 Object không sử dụng như các toán tử bình thương.


---

- [Giải thích về lớp ma thuật Unsafe trong Java]
	**`sun.misc.Unsafe`** là một lớp thuộc gói nội bộ của JDK, cung cấp quyền truy cập **cấp thấp** vào bộ nhớ và các thao tác hệ thống. Nó được gọi là **"lớp ma thuật"** vì cung cấp các tính năng mà các lập trình viên Java thông thường không có quyền truy cập thông qua API công khai.

Lớp `Unsafe` được thiết kế chủ yếu cho các **framework và thư viện Java**, chẳng hạn như:

- **Sun/Oracle JVM internals** (Ví dụ: triển khai `java.util.concurrent`).
- **Thao tác với bộ nhớ** như trong ngôn ngữ C/C++.
- **Tối ưu hóa hiệu suất** cho các tác vụ cụ thể.

**Tại sao `Unsafe` được coi là nguy hiểm?**

- **Truy cập trực tiếp vào bộ nhớ**: Cho phép thao tác trực tiếp vào bộ nhớ ngoài quyền kiểm soát của JVM.
- **Vượt qua cơ chế kiểm tra an toàn**: Có thể phá vỡ bảo mật, như đọc/ghi dữ liệu bất hợp pháp.
- **Không an toàn**: Sai sót trong việc sử dụng `Unsafe` có thể gây **lỗi JVM**, như segmentation faults hoặc crash toàn hệ thống.


---
- [Giải thích về cơ chế SPI trong Java]
	**SPI (Service Provider Interface)** là một cơ chế trong Java cho phép các thư viện hoặc framework mở rộng chức năng bằng cách cho phép các nhà phát triển bên thứ ba cung cấp các **service implementations** (triển khai dịch vụ) mà không cần thay đổi mã nguồn của thư viện.

 **Tổng quan về SPI**

- **Mục tiêu**: SPI được thiết kế để cung cấp một cơ chế mở rộng linh hoạt. Thay vì triển khai cố định trong thư viện, các dịch vụ có thể được tùy chỉnh hoặc mở rộng bởi người dùng.
- **Nguyên lý**: Một **API công khai** định nghĩa cách sử dụng dịch vụ, trong khi **SPI** định nghĩa cách dịch vụ được triển khai.

 **Thành phần chính của SPI**

1. **Interface (hoặc abstract class)**:
    
    - Được cung cấp bởi thư viện hoặc framework.
    - Định nghĩa các phương thức mà service phải triển khai.
2. **Service Implementation**:
    
    - Được cung cấp bởi bên thứ ba hoặc người dùng.
    - Thực hiện các phương thức được định nghĩa bởi interface SPI.
3. **Cơ chế đăng ký dịch vụ**:
    
    - **Tệp cấu hình** (`META-INF/services`): Lưu trữ danh sách các class triển khai.
    - **`ServiceLoader`**: Lớp hỗ trợ Java tích hợp sẵn để tải các triển khai dịch vụ.
---

- [Giải thích về cú pháp đường trong Java]
	**Cú pháp đường (syntactic sugar)** là một thuật ngữ dùng để chỉ những tính năng của ngôn ngữ lập trình giúp mã nguồn dễ đọc, dễ viết hơn mà không thêm bất kỳ chức năng mới nào vào ngôn ngữ. Trong Java, cú pháp đường giúp giảm bớt sự phức tạp trong việc viết mã nhưng vẫn biên dịch về các cấu trúc cơ bản hơn.

Ví dụ: `for-each`, `? :`, `String.format`...

---

### Bộ sưu tập (Collections)

Tổng hợp các điểm kiến thức và câu hỏi phỏng vấn về Collections Java.
Java **Collections Framework** cung cấp một tập hợp các lớp và giao diện để lưu trữ và thao tác với dữ liệu. Dưới đây là các thành phần chính trong Collections:
- **Collections Interfaces**:
    
    - **Collection**: Là giao diện cha của tất cả các lớp Collection khác. Nó không có phương thức riêng, nhưng là cơ sở cho các giao diện khác.
    - **List**: Lưu trữ các phần tử theo thứ tự và cho phép các phần tử trùng lặp. Ví dụ: `ArrayList`, `LinkedList`, `Vector`.
    - **Set**: Lưu trữ các phần tử duy nhất (không có phần tử trùng lặp). Ví dụ: `HashSet`, `LinkedHashSet`, `TreeSet`.
    - **Queue**: Được sử dụng để lưu trữ các phần tử theo thứ tự FIFO (First-In-First-Out). Ví dụ: `LinkedList`, `PriorityQueue`.
    - **Map**: Lưu trữ dữ liệu dưới dạng cặp khóa-giá trị (key-value). Ví dụ: `HashMap`, `TreeMap`, `LinkedHashMap`.
- **Các lớp chính trong Collections**:
    
    - **ArrayList**: Là một triển khai của giao diện `List`, lưu trữ theo dạng mảng động, cho phép truy cập theo chỉ số và hỗ trợ phần tử trùng lặp.
    - **HashSet**: Là một triển khai của `Set`, không lưu trữ phần tử trùng lặp và không đảm bảo thứ tự.
    - **LinkedList**: Cung cấp cả hai tính năng của `List` và `Queue`, hỗ trợ thao tác thêm/xóa phần tử nhanh chóng.
    - **HashMap**: Là một triển khai của `Map`, lưu trữ cặp khóa-giá trị và không đảm bảo thứ tự của phần tử.
    - **TreeMap**: Là một triển khai của `Map`, lưu trữ cặp khóa-giá trị trong một thứ tự sắp xếp theo các khóa.
- **Các phương thức chính của các Collection**:
    
    - **add()**: Thêm một phần tử vào Collection.
    - **remove()**: Xóa một phần tử khỏi Collection.
    - **size()**: Trả về số lượng phần tử trong Collection.
    - **contains()**: Kiểm tra xem một phần tử có tồn tại trong Collection hay không.
    - **clear()**: Xóa tất cả phần tử trong Collection.
    - **isEmpty()**: Kiểm tra xem Collection có rỗng hay không.
    - **iterator()**: Trả về một `Iterator` để lặp qua các phần tử trong Collection.
- **Ưu điểm và nhược điểm**:
    
    - **ArrayList**: Truy cập nhanh theo chỉ số, nhưng chậm khi chèn hoặc xóa phần tử giữa danh sách.
    - **LinkedList**: Hỗ trợ thêm/xóa phần tử nhanh chóng, nhưng chậm trong việc truy cập phần tử theo chỉ số.
    - **HashSet**: Tìm kiếm nhanh, nhưng không duy trì thứ tự các phần tử.
    - **HashMap**: Tìm kiếm nhanh dựa trên khóa, nhưng không duy trì thứ tự của các phần tử.
- **Sắp xếp và tìm kiếm trong Collections**:
    
    - **Sorting**: `Collections.sort(List<T> list)` hoặc sử dụng `Comparator` để tùy chỉnh cách sắp xếp.
    - **Binary Search**: Dùng `Collections.binarySearch(List<T> list, T key)` để tìm kiếm nhị phân trong một danh sách đã sắp xếp.

---

### **II. Các câu hỏi phỏng vấn về Collections trong Java**

#### **1. Kiến thức cơ bản về Collections**

- **Câu hỏi 1**: **Collection trong Java là gì?**
    
    - **Trả lời**: `Collection` là một giao diện trong Java mà tất cả các lớp thuộc bộ sưu tập (Collection Framework) đều phải triển khai. Nó là cha của các giao diện con như `List`, `Set`, `Queue`, và `Map`.
- **Câu hỏi 2**: **Sự khác biệt giữa `List`, `Set` và `Map` trong Java là gì?**
    
    - **Trả lời**:
        - `List`: Duy trì thứ tự phần tử và cho phép các phần tử trùng lặp. Ví dụ: `ArrayList`, `LinkedList`.
        - `Set`: Không duy trì thứ tự phần tử và không cho phép các phần tử trùng lặp. Ví dụ: `HashSet`, `TreeSet`.
        - `Map`: Lưu trữ các phần tử dưới dạng cặp khóa-giá trị (key-value). Ví dụ: `HashMap`, `TreeMap`.
- **Câu hỏi 3**: **Giải thích về `ArrayList` và `LinkedList`.**
    
    - **Trả lời**:
        - `ArrayList`: Lưu trữ các phần tử trong một mảng động, hỗ trợ truy cập phần tử theo chỉ số nhanh chóng nhưng chậm khi chèn hoặc xóa phần tử ở giữa.
        - `LinkedList`: Lưu trữ các phần tử trong một danh sách liên kết, nhanh trong việc chèn/xóa phần tử nhưng truy cập theo chỉ số chậm hơn.
#### **2. Các phương thức và thao tác cơ bản**

- **Câu hỏi 4**: **Phân biệt giữa `add()`, `remove()`, `contains()` trong `Collection`.**
    
    - **Trả lời**:
        - `add()`: Thêm phần tử vào Collection.
        - `remove()`: Xóa phần tử khỏi Collection.
        - `contains()`: Kiểm tra xem phần tử có tồn tại trong Collection hay không.
- **Câu hỏi 5**: **Làm thế nào để lặp qua tất cả phần tử trong một `List` hoặc `Set`?**
    
    - **Trả lời**: Bạn có thể sử dụng vòng lặp `for-each`, `Iterator`, hoặc `Stream` để lặp qua các phần tử trong `List` hoặc `Set`.
#### **3. Sắp xếp và tìm kiếm trong Collections**

- **Câu hỏi 6**: **Làm thế nào để sắp xếp một `List` trong Java?**
    
    - **Trả lời**: Bạn có thể sử dụng `Collections.sort(List<T> list)` để sắp xếp một danh sách theo thứ tự tự nhiên của các phần tử hoặc sử dụng một `Comparator` tùy chỉnh.
- **Câu hỏi 7**: **Giải thích về `HashSet` và `TreeSet`.**
    
    - **Trả lời**:
        - `HashSet`: Không đảm bảo thứ tự của các phần tử và không cho phép phần tử trùng lặp.
        - `TreeSet`: Lưu trữ các phần tử trong thứ tự đã sắp xếp (theo mặc định hoặc theo `Comparator`).
- **Câu hỏi 8**: **Làm thế nào để thực hiện tìm kiếm nhị phân trong một `List`?**
    
    - **Trả lời**: Dùng `Collections.binarySearch(List<T> list, T key)` để tìm kiếm nhị phân trong một danh sách đã được sắp xếp.
#### **4. Các câu hỏi về Map**

- **Câu hỏi 9**: **Sự khác biệt giữa `HashMap` và `TreeMap` là gì?**
    
    - **Trả lời**:
        - `HashMap`: Không đảm bảo thứ tự của các phần tử, sử dụng băm để tìm kiếm khóa.
        - `TreeMap`: Lưu trữ các phần tử theo thứ tự sắp xếp của các khóa (theo mặc định hoặc theo `Comparator`).
- **Câu hỏi 10**: **Giải thích về phương thức `put()` và `get()` trong `Map`.**
    
    - **Trả lời**:
        - `put()`: Thêm cặp khóa-giá trị vào `Map`.
        - `get()`: Lấy giá trị từ `Map` bằng khóa.

#### **5. Các câu hỏi nâng cao**

- **Câu hỏi 11**: **Phân biệt giữa `HashSet` và `LinkedHashSet`.**
    
    - **Trả lời**: `LinkedHashSet` lưu trữ các phần tử trong thứ tự mà chúng được thêm vào, trong khi `HashSet` không đảm bảo thứ tự.
- **Câu hỏi 12**: **Java `Collections.synchronizedList()` là gì?**
    
    - **Trả lời**: `Collections.synchronizedList()` trả về một danh sách đồng bộ hóa (thread-safe) bằng cách bao bọc danh sách ban đầu.



----

### Tổng hợp câu hỏi phỏng vấn về IO (Input/Output) trong Java

Java cung cấp một API mạnh mẽ để làm việc với các luồng dữ liệu thông qua **java.io** và **java.nio**. Các khái niệm cơ bản trong IO của Java bao gồm:
- **Byte Streams**: Được sử dụng để đọc và ghi dữ liệu dưới dạng các byte, thường được sử dụng cho các dữ liệu nhị phân như hình ảnh, âm thanh, v.v. Các lớp chính là:
    
    - `FileInputStream`
    - `FileOutputStream`
    - `BufferedInputStream`
    - `BufferedOutputStream`
- **Character Streams**: Được sử dụng để đọc và ghi dữ liệu dưới dạng các ký tự, chủ yếu là cho văn bản. Các lớp chính là:
    
    - `FileReader`
    - `FileWriter`
    - `BufferedReader`
    - `BufferedWriter`
- **File I/O**: Làm việc với các file trên hệ thống, bao gồm các thao tác như đọc, ghi, tạo, xóa, kiểm tra sự tồn tại của file, v.v.
    
    - `File`
    - `Files`
- **Serializable**: Một giao diện trong Java được sử dụng để đánh dấu các đối tượng có thể được chuyển đổi thành chuỗi byte để lưu trữ hoặc truyền qua mạng.
    
- **NIO (New I/O)**: Java NIO cung cấp một API hiện đại và hiệu suất cao để xử lý I/O, khác với I/O truyền thống (IO Streams). Các lớp chính là:
    
    - `FileChannel`
    - `Path`
    - `Files`
    - `ByteBuffer`

---
### **Tổng hợp câu hỏi phỏng vấn về IO (Input/Output) trong Java**

---

### **I. Kiến thức cơ bản về IO trong Java**

Java cung cấp một API mạnh mẽ để làm việc với các luồng dữ liệu thông qua **java.io** và **java.nio**. Các khái niệm cơ bản trong IO của Java bao gồm:

1. **Byte Streams**: Được sử dụng để đọc và ghi dữ liệu dưới dạng các byte, thường được sử dụng cho các dữ liệu nhị phân như hình ảnh, âm thanh, v.v. Các lớp chính là:
    
    - `FileInputStream`
    - `FileOutputStream`
    - `BufferedInputStream`
    - `BufferedOutputStream`
2. **Character Streams**: Được sử dụng để đọc và ghi dữ liệu dưới dạng các ký tự, chủ yếu là cho văn bản. Các lớp chính là:
    
    - `FileReader`
    - `FileWriter`
    - `BufferedReader`
    - `BufferedWriter`
3. **File I/O**: Làm việc với các file trên hệ thống, bao gồm các thao tác như đọc, ghi, tạo, xóa, kiểm tra sự tồn tại của file, v.v.
    
    - `File`
    - `Files`
4. **Serializable**: Một giao diện trong Java được sử dụng để đánh dấu các đối tượng có thể được chuyển đổi thành chuỗi byte để lưu trữ hoặc truyền qua mạng.
    
5. **NIO (New I/O)**: Java NIO cung cấp một API hiện đại và hiệu suất cao để xử lý I/O, khác với I/O truyền thống (IO Streams). Các lớp chính là:
    
    - `FileChannel`
    - `Path`
    - `Files`
    - `ByteBuffer`

---

### **II. Các câu hỏi phỏng vấn về IO trong Java**

#### **1. Kiến thức cơ bản về IO**

- **Câu hỏi 1**: **IO trong Java là gì và phân biệt giữa Stream và Reader/Writer?**
    
    - **Trả lời**:
        - **Stream** là một khái niệm trong Java I/O dùng để xử lý các luồng byte (dữ liệu nhị phân).
        - **Reader/Writer** là các lớp trong Java để xử lý các luồng ký tự (dữ liệu văn bản).
        - Stream: `InputStream`, `OutputStream`.
        - Reader/Writer: `Reader`, `Writer`.
- **Câu hỏi 2**: **Sự khác biệt giữa `InputStream` và `Reader` trong Java?**
    
    - **Trả lời**: `InputStream` được sử dụng cho việc xử lý dữ liệu nhị phân (byte), trong khi `Reader` được dùng cho dữ liệu văn bản (ký tự). `InputStream` đọc byte, còn `Reader` đọc ký tự.
- **Câu hỏi 3**: **Giải thích sự khác biệt giữa `FileInputStream` và `FileReader`.**
    
    - **Trả lời**:
        - `FileInputStream` là một lớp byte stream, đọc dữ liệu từ file dưới dạng byte.
        - `FileReader` là một lớp character stream, đọc dữ liệu từ file dưới dạng ký tự.
- **Câu hỏi 4**: **Làm thế nào để đọc và ghi file trong Java?**
    
    - **Trả lời**:
        - Đọc file: `FileReader fr = new FileReader("file.txt"); BufferedReader br = new BufferedReader(fr);`
        - Ghi file: `FileWriter fw = new FileWriter("file.txt"); BufferedWriter bw = new BufferedWriter(fw);`

---

#### **2. Các lớp và phương thức trong IO**

- **Câu hỏi 5**: **Giải thích về lớp `BufferedReader` và `BufferedWriter` trong Java.**
    
    - **Trả lời**:
        - `BufferedReader` giúp cải thiện hiệu suất khi đọc dữ liệu bằng cách sử dụng bộ đệm để đọc các ký tự theo khối, giảm số lần truy cập vào ổ đĩa.
        - `BufferedWriter` giúp cải thiện hiệu suất khi ghi dữ liệu bằng cách sử dụng bộ đệm để ghi các ký tự theo khối.
- **Câu hỏi 6**: **Lớp `PrintWriter` có gì đặc biệt?**
    
    - **Trả lời**: `PrintWriter` có thể tự động ghi dữ liệu dạng văn bản và hỗ trợ phương thức `println()` để ghi dữ liệu mà không cần phải gọi `flush()`, rất hữu ích khi làm việc với văn bản.
- **Câu hỏi 7**: **Giải thích về lớp `File` trong Java.**
    
    - **Trả lời**: `File` là lớp trong Java đại diện cho các file và thư mục trong hệ thống file. Nó cung cấp các phương thức để kiểm tra, xóa, tạo, thay đổi tên, di chuyển, và lấy các thuộc tính của file hoặc thư mục.
- **Câu hỏi 8**: **Sự khác biệt giữa `FileInputStream` và `BufferedInputStream` là gì?**
    
    - **Trả lời**: `BufferedInputStream` là một lớp bao bọc (wrapper) xung quanh `FileInputStream` để cải thiện hiệu suất bằng cách đọc dữ liệu thành các khối lớn thay vì từng byte một.

---

#### **3. Đọc và ghi dữ liệu trong IO**

- **Câu hỏi 9**: **Làm thế nào để đọc và ghi các đối tượng trong Java?**
    
    - **Trả lời**:
        - Để đọc và ghi các đối tượng, bạn cần sử dụng `ObjectInputStream` và `ObjectOutputStream`.
- **Câu hỏi 10**: **Giải thích về giao diện `Serializable` trong Java.**
    
    - **Trả lời**: Giao diện `Serializable` được sử dụng để đánh dấu rằng một đối tượng có thể được chuyển đổi thành chuỗi byte và lưu trữ hoặc truyền qua mạng. Các đối tượng phải thực hiện giao diện này để có thể được ghi vào hoặc đọc từ luồng.

---

#### **4. NIO (New Input/Output)**

- **Câu hỏi 11**: **Java NIO là gì và sự khác biệt giữa NIO và IO truyền thống?**
    
    - **Trả lời**: Java NIO (New I/O) là một API mạnh mẽ để xử lý I/O phi đồng bộ (non-blocking) với hiệu suất cao, bao gồm các khái niệm như `Buffer`, `Channel`, `Selector`. Khác với IO truyền thống, NIO cung cấp khả năng làm việc với nhiều kết nối đồng thời mà không bị chặn.
- **Câu hỏi 12**: **Giải thích về `FileChannel` trong NIO.**
    
    - **Trả lời**: `FileChannel` trong NIO cho phép bạn đọc và ghi file theo cách nhanh chóng và phi đồng bộ. Nó cho phép thao tác trực tiếp với hệ thống file thông qua bộ đệm (`ByteBuffer`).
- **Câu hỏi 13**: **Phân biệt giữa `Path` và `File` trong NIO và IO.**
    
    - **Trả lời**:
        - `Path` là một lớp trong NIO đại diện cho đường dẫn tới một file hoặc thư mục. NIO sử dụng `Path` để làm việc với hệ thống file.
        - `File` là lớp truyền thống trong IO đại diện cho một file hoặc thư mục trong hệ thống file.
- **Câu hỏi 14**: **Giải thích về `ByteBuffer` trong NIO.**
    
    - **Trả lời**: `ByteBuffer` là một lớp trong NIO giúp xử lý dữ liệu byte trong bộ nhớ. Nó cung cấp các phương thức để đọc và ghi dữ liệu từ buffer vào các channel, giúp xử lý I/O hiệu quả hơn.

---

#### **5. Thực tiễn và xử lý ngoại lệ**

- **Câu hỏi 15**: **Khi nào bạn sử dụng `finally` trong khối IO?**
    
    - **Trả lời**: `finally` được sử dụng để đảm bảo rằng các tài nguyên như `InputStream`, `OutputStream`, `Reader`, và `Writer` được đóng lại sau khi sử dụng, ngay cả khi có lỗi xảy ra.
- **Câu hỏi 16**: **Sự khác biệt giữa `IOException` và `FileNotFoundException`?**
    
    - **Trả lời**:
        - `IOException`: Là một ngoại lệ chung cho tất cả các lỗi I/O.
        - `FileNotFoundException`: Là một ngoại lệ con của `IOException`, được ném khi một file không thể tìm thấy.

---

### **III. Câu hỏi nâng cao**

- **Câu hỏi 17**: **Làm thế nào để xử lý I/O phi đồng bộ trong Java?**
    
    - **Trả lời**: Bạn có thể sử dụng `java.nio.channels` và `Selector` để thực hiện I/O phi đồng bộ trong Java. Điều này cho phép chương trình xử lý nhiều kết nối cùng lúc mà không bị chặn.
- **Câu hỏi 18**: **Giải thích về `Mmap` trong Java.**
    
    - **Trả lời**: `Mmap` (Memory Mapping) là kỹ thuật sử dụng bộ nhớ ảo để trực tiếp ánh xạ file vào bộ nhớ. Điều này giúp tăng hiệu suất I/O, đặc biệt là khi làm việc với các file lớn.

---
- **Java 8**:
    - Giới thiệu **Lambda Expressions** và **Streams API**, giúp xử lý dữ liệu dễ dàng hơn.
    - **Optional API**: xử lý giá trị `null` an toàn hơn.
    - **Default Methods**: cho phép thêm phương thức mặc định vào interface mà không phá vỡ code cũ.
- **Java 17**:
    - **Sealed Classes** (JEP 409): kiểm soát khả năng kế thừa của các lớp.
    - **Pattern Matching** (JEP 394, JEP 395): hỗ trợ cú pháp gọn hơn khi làm việc với instanceof.
    - **Text Blocks** (ra mắt từ Java 15): giúp viết chuỗi nhiều dòng dễ đọc hơn.
    - **Removal of RMI Activation**: loại bỏ tính năng lỗi thời.
- **Java 21**:
    - **Virtual Threads (Project Loom)**: cải thiện khả năng xử lý đồng thời (concurrency).
    - **String Templates (Preview)**: cú pháp động để xử lý chuỗi.
    - **Structured Concurrency (Preview)**: tổ chức luồng xử lý logic hơn.
    - **Record Patterns** và **Pattern Matching for switch**: tăng cường cú pháp đơn giản hóa xử lý logic.

---
### **1. Lightweight**

Spring là một **lightweight framework**, nghĩa là nó được thiết kế để hoạt động hiệu quả, không tiêu tốn nhiều tài nguyên, và linh hoạt trong cách sử dụng. Các ứng dụng Spring chỉ sử dụng những thành phần cần thiết, tránh lãng phí tài nguyên.

### **2. Inversion of Control (IOC)**

IOC là nguyên tắc mà Spring sử dụng để quản lý các đối tượng trong ứng dụng. Thay vì tạo và quản lý các đối tượng một cách thủ công, Spring Container làm nhiệm vụ đó, dựa trên cấu hình hoặc chú thích (annotations).

### **3. Aspect-Oriented Programming (AOP)**

AOP cho phép bạn tách biệt các **cross-cutting concerns** (như logging, bảo mật, transaction management) khỏi logic chính của ứng dụng bằng cách sử dụng **aspects**.

### **4. Container**

Spring **Container** là phần chịu trách nhiệm quản lý vòng đời của các bean. Nó tạo, khởi tạo, định cấu hình, và quản lý các đối tượng trong ứng dụng.

### **5. Transaction Management**

Spring hỗ trợ quản lý transaction theo lập trình hoặc khai báo (declarative). Bạn có thể thực hiện các thao tác trên cơ sở dữ liệu và để Spring đảm bảo tính toàn vẹn của transaction.


Nếu có bất kỳ ngoại lệ nào xảy ra, toàn bộ transaction sẽ bị rollback.

### **6. JDBC Exception Handling**

Spring cung cấp một lớp ngoại lệ tùy chỉnh như `DataAccessException` để giúp xử lý lỗi liên quan đến JDBC dễ dàng và có ý nghĩa hơn.


---

### **1. Singleton**

**Singleton** là một mẫu thiết kế đảm bảo rằng một lớp chỉ có **một đối tượng duy nhất** và cung cấp một cách để truy cập đối tượng đó.

### **2. Factory + Abstract Factory**

#### **Factory Pattern**

**Factory Pattern** tạo ra đối tượng mà không cần chỉ định lớp cụ thể.

#### **Abstract Factory Pattern**

**Abstract Factory** là một factory của các factory.


### **3. Service Locator**

**Service Locator Pattern** cung cấp một cách để lấy đối tượng (services) mà không cần khởi tạo chúng trực tiếp.

### **4. Façade**

**Façade Pattern** cung cấp một giao diện đơn giản hơn để truy cập vào một hệ thống phức tạp.

### **5. Observer**

**Observer Pattern** cho phép một đối tượng (subject) thông báo cho các đối tượng khác (observers) khi có sự thay đổi trạng thái.

### **6. Builder**

**Builder Pattern** giúp xây dựng các đối tượng phức tạp bằng cách sử dụng một interface hoặc lớp builder.

---
### **Tóm tắt chức năng cho từng Pattern:**

|**Pattern**|**Chức năng**|
|---|---|
|**Singleton**|Logger, Quản lý cấu hình, kết nối cơ sở dữ liệu.|
|**Factory**|Tạo thông báo (Email/SMS), tạo đối tượng động (DAO, Service).|
|**Service Locator**|Tìm kiếm và quản lý service (Email, Báo cáo, Plugin).|
|**Façade**|Tích hợp API, xử lý đơn hàng (Thanh toán, vận chuyển), quản lý tài khoản.|
|**Observer**|Hệ thống thông báo, cập nhật dữ liệu thời gian thực (Chứng khoán, Mạng xã hội).|
|**Builder**|Tạo đối tượng phức tạp (Email, Báo cáo, Cấu hình tài liệu).|


---

### **1. Redis là gì? Nêu một số tính năng chính của Redis.**

Redis (Remote Dictionary Server) là một cơ sở dữ liệu NoSQL in-memory, mã nguồn mở, thường được sử dụng như một bộ nhớ đệm, cơ sở dữ liệu, hoặc message broker.

**Tính năng chính:**

- **In-memory data store:** Tốc độ cực nhanh nhờ dữ liệu lưu trữ trực tiếp trong RAM.
- **Hỗ trợ nhiều cấu trúc dữ liệu:** String, Hash, List, Set, Sorted Set, Stream, GeoSpatial, v.v.
- **Persistence:** Hỗ trợ lưu trữ dữ liệu vào ổ đĩa thông qua cơ chế RDB hoặc AOF.
- **Replication:** Hỗ trợ master-slave replication để tăng tính sẵn sàng.
- **Cluster:** Có khả năng mở rộng theo chiều ngang với Redis Cluster.
- **Pub/Sub:** Cung cấp cơ chế giao tiếp theo mô hình publish/subscribe.

---

### **2. Redis hỗ trợ những kiểu dữ liệu nào?**

Redis hỗ trợ các kiểu dữ liệu sau:

- **String:** Dạng key-value cơ bản, dùng để lưu văn bản hoặc số.
- **Hash:** Lưu trữ các cặp key-value nhỏ bên trong một key lớn, hữu ích cho thông tin người dùng.
- **List:** Danh sách có thứ tự, hỗ trợ các thao tác như push, pop, trim.
- **Set:** Một tập hợp không có thứ tự, không cho phép giá trị trùng lặp.
- **Sorted Set:** Tập hợp có thứ tự với mỗi phần tử được gắn điểm số (score).
- **Stream:** Hỗ trợ lưu trữ dữ liệu theo dòng thời gian, phù hợp cho hệ thống ghi log.
- **GeoSpatial:** Lưu trữ và truy vấn dữ liệu địa lý.


### **3. Sự khác biệt giữa RDB và AOF trong Redis là gì?**

| Đặc điểm       | RDB (Redis Database)                        | AOF (Append-Only File)                            |
| -------------- | ------------------------------------------- | ------------------------------------------------- |
| **Mục đích**   | Snapshot dữ liệu định kỳ.                   | Ghi lại tất cả các thao tác ghi.                  |
| **Hiệu suất**  | Nhanh hơn, ít ảnh hưởng hiệu suất.          | Chậm hơn do ghi liên tục.                         |
| **Độ an toàn** | Có thể mất dữ liệu trong lần snapshot cuối. | Ít mất dữ liệu hơn do ghi gần như thời gian thực. |
| **Dung lượng** | File nhỏ hơn.                               | File lớn hơn, cần tối ưu định kỳ.                 |
| **Sử dụng**    | Phù hợp cho backup.                         | Phù hợp cho đảm bảo an toàn dữ liệu.              |

### **Redis Cluster là gì? Làm thế nào để thiết lập Redis Cluster?**

Redis Cluster là một phương pháp phân mảnh (sharding) để Redis có thể mở rộng theo chiều ngang, cung cấp tính năng phân phối dữ liệu và khả năng chịu lỗi (fault tolerance).

**Cách thiết lập:**

1. Khởi tạo nhiều instance Redis trên các máy chủ khác nhau.
2. Cấu hình file `redis.conf` với các tham số như `cluster-enabled yes` và `cluster-config-file nodes.conf`.
3. Sử dụng lệnh `redis-cli --cluster create` để thiết lập cluster.

### **Redis khác gì so với Memcached?**

| Đặc điểm             | Redis                                   | Memcached                        |
| -------------------- | --------------------------------------- | -------------------------------- |
| **Cấu trúc dữ liệu** | Hỗ trợ nhiều kiểu (String, Hash, List). | Chỉ hỗ trợ String.               |
| **Persistence**      | Có (RDB, AOF).                          | Không có.                        |
| **Replication**      | Hỗ trợ master-slave, cluster.           | Không hỗ trợ.                    |
| **Tốc độ**           | Tương đương nhưng Redis linh hoạt hơn.  | Chỉ tối ưu cho key-value cơ bản. |
| **Ứng dụng**         | Dùng làm cache, message broker, DB.     | Chủ yếu làm cache.               |

### **Redis xử lý cache avalanche như thế nào?**

**Cache avalanche** xảy ra khi một lượng lớn key trong cache hết hạn cùng lúc, dẫn đến tải lớn lên database. Redis xử lý bằng cách:

- **Thiết lập thời gian hết hạn ngẫu nhiên (random TTL):** Tránh việc các key hết hạn cùng lúc.
- **Cache Warm-up:** Làm nóng cache trước khi triển khai hệ thống.
- **Locking:** Sử dụng distributed locks để tránh nhiều request cùng tải dữ liệu từ database.

### **Redis quản lý bộ nhớ như thế nào? Điều gì xảy ra khi bộ nhớ đầy?**

#### **Quản lý bộ nhớ:**

- Redis lưu dữ liệu chủ yếu trong RAM, giúp truy cập rất nhanh.
- Redis sử dụng cơ chế **dynamic memory allocation**, thường thông qua thư viện `jemalloc`.
- Các key không cần thiết hoặc hết hạn sẽ bị xóa tự động nếu TTL được thiết lập.

#### **Khi bộ nhớ đầy:**

Redis sẽ sử dụng **Eviction Policy** để quyết định cách xử lý:

1. **noeviction**: Từ chối ghi thêm dữ liệu mới (mặc định).
2. **allkeys-lru**: Xóa key ít được sử dụng gần đây nhất (Least Recently Used).
3. **volatile-lru**: Xóa key ít được sử dụng nhất, chỉ trong các key có TTL.
4. **allkeys-random**: Xóa ngẫu nhiên bất kỳ key nào.
5. **volatile-random**: Xóa ngẫu nhiên các key có TTL.
6. **volatile-ttl**: Xóa key có TTL sắp hết hạn nhất.

**Cách kiểm soát:**

- Sử dụng lệnh `CONFIG SET maxmemory <size>` để giới hạn bộ nhớ.
- Theo dõi bộ nhớ bằng lệnh `INFO memory`.

### **Redis có phải là thread-safe không?**

Redis **thread-safe**, nhưng hoạt động theo cơ chế **single-threaded** để xử lý các lệnh từ client. Điều này đảm bảo:

1. **Tính tuần tự:** Không có tình trạng race condition giữa các lệnh.
2. **Hiệu suất cao:** Nhờ sử dụng mô hình event-loop không bị overhead do locking.

Mặc dù Redis là single-threaded cho các lệnh chính, nhưng một số tác vụ như snapshot (RDB) hoặc sao chép (AOF) có thể chạy trên các thread riêng biệt. Redis 6+ hỗ trợ **I/O multi-threading** để cải thiện hiệu suất trong các hệ thống có nhiều kết nối đồng thời.

### **Làm thế nào để giảm thiểu cache miss trong Redis?**

#### **Cache miss** xảy ra khi dữ liệu yêu cầu không có trong cache, dẫn đến phải truy vấn nguồn dữ liệu gốc (database), làm giảm hiệu suất.

**Cách giảm thiểu:**

1. **Cache Warm-up (Khởi động trước):**
    
    - Nạp dữ liệu phổ biến vào Redis ngay khi khởi động hệ thống.
    - Ví dụ: Nạp danh sách sản phẩm bán chạy vào cache.
2. **Tối ưu TTL:**
    
    - Thiết lập TTL phù hợp để key không hết hạn quá nhanh.


1. **Chính sách Prefetch:**
    
    - Khi dữ liệu sắp hết hạn, tải lại từ nguồn gốc trước khi bị truy vấn lần tiếp theo.
2. **Read-Through Cache:**
    
    - Sử dụng cơ chế tự động nạp dữ liệu từ database nếu cache miss.
3. **Chống Cache Penetration:**
    
    - Tránh truy vấn key không tồn tại bằng cách lưu trữ placeholder (`null`) cho các key thường bị truy vấn sai.
4. **Sử dụng Cluster:**
    
    - Redis Cluster giúp phân phối tải và giảm áp lực lên một node duy nhất.
5. **Theo dõi và phân tích:**
    
    - Dùng lệnh `MONITOR` hoặc `INFO stats` để kiểm tra các key thường xuyên bị cache miss và tối ưu.


---


