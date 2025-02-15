
---
Khi làm việc với dữ liệu kiểu text trong Java cung cấp 3 class String, StringBuffer và StringBuilder. Cơ bản về 3 class này như sau: String là không thể thay đổi (immutable), và không cho phép có class con. StringBuffer, StringBuilder có thể thay đổi (mutable) StringBuilder và StringBuffer là giống nhau, nó chỉ khác biệt tình huống sử dụng có liên quan tới đa luồng (Multi Thread). => về tốc độ xử lý StringBuilder là tốt nhất, sau đó StringBuffer và cuối cùng mới là String.


**String** Trong java, String là một class đặc biệt, nguyên nhân là nó được sử dụng một cách thường xuyên trong một chương trình, vì vậy đòi hỏi nó phải có hiệu suất và sự mềm dẻo. Đó là lý do tại sao String có tính đối tượng và vừa có tính nguyên thủy (primitive). Ví Dụ về tính nguyên thủy: Chúng ta hoàn toàn có thể khai báo: String s1 = "Hello word" Bạn cũng có thể sử dụng toán tử + để nối 2 string, toán tử này vốn quen thuộc và sử dụng cho các kiểu dữ liệu nguyên thủy như int, float, double. Tính đối tượng: Vì String là một class, vì vậy nó có thể được tạo ra thông qua toán tử new. `String object = new String("Hello World");` Các đối tượng String được lưu trữ trên Heap, yêu cầu quản lý bộ nhớ phức tạp và tốn không gian lưu trữ. Hai đối tượng String có nội dung giống nhau lưu trữ trên 2 vùng bộ nhớ khác nhau của Heap. Phương thức Equals vs == Phương thức equals() sử dụng để so sánh 2 đối tượng, với String nó có ý nghĩa là so sánh nội dung của 2 string. Đối với các kiểu tham chiếu (reference) toán tử == có ý nghĩa là so sánh địa chỉ vùng bộ nhớ lưu trữ của đối tượng.

```Java
String s1 = "Hello";
String s2 = "Hello"; 
String s3 = s1;  
String s4 = new String("Hello");  
String s5 = new String("Hello"); 
 
s1 == s1;         // true
s1 == s2;         // true
s1 == s3;         // true
s1 == s4;         // false
s4 == s5;         // false
 
s1.equals(s3);    // true, cùng nội dung
s1.equals(s4);    // true, cùng nội dung
s4.equals(s5);    // true, cùng nội dung.
```

**StringBuffer vs StringBuilder** StringBuilder và StringBuffer là khá giống nhau, điều khác biệt là tất cả các phương thức của **StringBuffer đã được đồng bộ, nó thích hợp khi bạn làm việc với ứng dụng đa luồng, nhiều luồng có thể truy cập vào một đối tượng StringBuffer cùng lúc.** Trong khi đó **StringBuilder có các phương thức tương tự nhưng không được đồng bộ, vì vậy mà hiệu suất của nó cao hơn, bạn nên sử dụng StringBuilder trong ứng dụng đơn luồng,** hoặc sử dụng như một biến địa phương trong một phương thức.
**String Buffer**
```Java
public class BufferString {
    public static void main(String[] args) {
        StringBuffer buffer = new StringBuffer("Hello");
        buffer.append("Word");
        System.out.println(buffer);
    }
}
```
**String Builder**
```Java
public class BuiderString {
    public static void main(String[] args) {
        StringBuffer buider = new StringBuffer("Hello");
        buffer.append("Word");
        System.out.println(buider);
    }
}
```


---
Hiểu sâu về `String`, `StringBuilder`, và `StringBuffer` trong Java như một **Senior Java Developer** đòi hỏi bạn phải nắm rõ:

- **Cấu trúc bộ nhớ và cách hoạt động của từng loại**
- **Khác biệt trong tính bất biến (immutability)**
- **Hiệu suất và cách sử dụng tối ưu**
- **Thread-safety và các tình huống sử dụng phù hợp**
- **Cách JVM xử lý String trong bộ nhớ**
- **Cách tối ưu hóa và tránh memory leak**

Hãy đi vào từng khía cạnh một cách chi tiết.

## 1. `String` – Immutable và cơ chế Interning

### 📌 Đặc điểm quan trọng:

- `String` là **immutable** (bất biến), nghĩa là một khi đã tạo ra thì không thể thay đổi nội dung.
- Khi thay đổi `String`, một đối tượng mới sẽ được tạo trong **heap memory** hoặc **String Pool**.
- JVM có **String Interning**, giúp tái sử dụng các chuỗi giống nhau trong **String Pool**.

### 🛠 Cách hoạt động của String Pool:
```java
String s1 = "hello";  // Tạo trong String Pool
String s2 = "hello";  // Không tạo mới, dùng lại s1
String s3 = new String("hello"); // Tạo mới trong heap, không vào String Pool

System.out.println(s1 == s2); // true (cùng tham chiếu trong Pool)
System.out.println(s1 == s3); // false (khác vùng nhớ)
```
🔥 **Lưu ý:** Nếu muốn đưa `s3` vào String Pool, dùng `intern()`
```java
String s4 = s3.intern();
System.out.println(s1 == s4); // true
```
### 📉 Hạn chế của `String`:

- Việc thay đổi giá trị dẫn đến **tạo đối tượng mới**, gây tốn bộ nhớ và hiệu suất kém khi xử lý nhiều thay đổi.
- Ví dụ tệ khi dùng `String`:
```java
String result = "";
for (int i = 0; i < 10000; i++) {
    result += i; // Mỗi lần lặp tạo một String mới!
}
```
💡 **Giải pháp:** Dùng `StringBuilder` hoặc `StringBuffer`.

## 2. `StringBuilder` – Mutable và Hiệu suất cao

### 📌 Đặc điểm quan trọng:

- `StringBuilder` là **mutable**, tức là có thể thay đổi nội dung mà không tạo đối tượng mới.
- Sử dụng **internal character array**, giúp tránh lãng phí bộ nhớ khi thay đổi chuỗi.
- **Không thread-safe** nhưng nhanh hơn `StringBuffer`.

### 🛠 Cách sử dụng:
```java
StringBuilder sb = new StringBuilder("Hello");
sb.append(" World"); // Thay đổi nội dung mà không tạo object mới
System.out.println(sb); // Hello World
```
### 📊 Hiệu suất vượt trội so với `String`:
```java
long startTime = System.nanoTime();
StringBuilder sb = new StringBuilder();
for (int i = 0; i < 10000; i++) {
    sb.append(i);
}
long endTime = System.nanoTime();
System.out.println("Execution Time: " + (endTime - startTime));
```
🔥 **Lưu ý:** `StringBuilder` nhanh hơn nhưng **không thread-safe**.

## 3. `StringBuffer` – Thread-safe nhưng chậm hơn

### 📌 Đặc điểm quan trọng:

- Giống `StringBuilder` nhưng hỗ trợ **thread-safety** bằng cách dùng **synchronized**.
- Thích hợp cho môi trường **multi-threaded**.

### 🛠 Cách sử dụng:
```java
StringBuffer sb = new StringBuffer("Hello");
sb.append(" World");
System.out.println(sb); // Hello World
```
💡 **So sánh tốc độ `StringBuffer` vs `StringBuilder`:**
```java
long startTime = System.nanoTime();
StringBuffer sb = new StringBuffer();
for (int i = 0; i < 10000; i++) {
    sb.append(i);
}
long endTime = System.nanoTime();
System.out.println("StringBuffer Time: " + (endTime - startTime));
```
🔥 **Kết luận:** Nếu không cần đồng bộ (synchronization), hãy **dùng `StringBuilder` thay vì `StringBuffer`**.

## 4. So sánh chi tiết

|Đặc điểm|String|StringBuilder|StringBuffer|
|---|---|---|---|
|**Immutable?**|✅ Có|❌ Không|❌ Không|
|**Thread-safe?**|✅ Có (Immutable)|❌ Không|✅ Có (Synchronized)|
|**Hiệu suất**|🚫 Chậm|⚡ Nhanh|🐢 Chậm hơn SB|
|**Dùng trong?**|Hằng số, ít thay đổi|Xử lý chuỗi nhanh|Đa luồng|
## 5. Cách JVM xử lý String trong bộ nhớ

### 📌 Heap vs Stack vs String Pool

- **Heap Memory**: Chứa tất cả các object trong runtime.
- **Stack Memory**: Chứa biến local, reference đến object trên heap.
- **String Pool (PermGen/Metaspace)**: Chứa các chuỗi immutable được JVM tối ưu.
```java
String a = "Hello";       // Nằm trong String Pool
String b = new String("Hello"); // Tạo object mới trên Heap
```
💡 **String Pool giúp tiết kiệm bộ nhớ nhưng nếu quá nhiều chuỗi, có thể gây OutOfMemoryError.**

## 6. Khi nào dùng gì?

✅ **Dùng `String` khi:**

- Chuỗi không thay đổi nhiều.
- Chuỗi cần sử dụng nhiều lần để tối ưu bộ nhớ (`intern()`).

✅ **Dùng `StringBuilder` khi:**

- Cần thay đổi chuỗi nhiều lần mà không cần thread-safety.
- Xử lý văn bản lớn (concat, insert, replace).

✅ **Dùng `StringBuffer` khi:**

- Đang làm việc trong môi trường **đa luồng**.

