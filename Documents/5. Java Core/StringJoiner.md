
---
**StringJoiner là một API mới trong Java 8, dựa trên triển khai của StringBuilder, được sử dụng để nối các chuỗi bằng các dấu phân cách.**  
StringJoiner có hai hàm khởi tạo:

- Hàm khởi tạo đầu tiên yêu cầu truyền vào dấu phân cách (delimiter), tiền tố (prefix) và hậu tố (suffix).
- Hàm khởi tạo thứ hai chỉ yêu cầu truyền vào dấu phân cách (tiền tố và hậu tố mặc định là chuỗi rỗng).

```java
StringJoiner(CharSequence delimiter, CharSequence prefix, CharSequence suffix)
StringJoiner(CharSequence delimiter)
```

Trong một số tình huống nối chuỗi, việc sử dụng StringBuffer hoặc StringBuilder có thể trở nên phức tạp và rườm rà.

```java
List<Integer> values = Arrays.asList(1, 3, 5);
StringBuilder sb = new StringBuilder("(");

for (int i = 0; i < values.size(); i++) {
    sb.append(values.get(i));
    if (i != values.size() - 1) {
        sb.append(",");
    }
}

sb.append(")");
```

Đoạn mã trên nối các phần tử của danh sách thành chuỗi dạng (1,3,5), nhưng cần kiểm tra điều kiện để tránh thêm dấu phẩy thừa, khiến mã trở nên dài dòng.

Khi sử dụng StringJoiner để nối các phần tử của danh sách, mã trở nên ngắn gọn và rõ ràng hơn:
```java
List<Integer> values = Arrays.asList(1, 3, 5);
StringJoiner sj = new StringJoiner(",", "(", ")");

for (Integer value : values) {
    sj.add(value.toString());
}
```

Kết quả cũng là (1,3,5), nhưng cách viết đơn giản hơn nhiều.

Ngoài ra, giống như phương thức Collectors.joining(",") thường dùng trong Java, lớp bên dưới cũng được triển khai thông qua StringJoiner. Mã nguồn như sau:

```java
public static Collector<CharSequence, ?, String> joining(
    CharSequence delimiter, CharSequence prefix, CharSequence suffix) {
    return new CollectorImpl<>(
            () -> new StringJoiner(delimiter, prefix, suffix),
            StringJoiner::add, StringJoiner::merge,
            StringJoiner::toString, CH_NOID);
}
```

Điều này cho thấy StringJoiner là nền tảng cho việc nối chuỗi trong các API cấp cao như Collectors.

- indexOf(): returns the index of the specified character.
- charAt(): Returns the character at the specified index.
- replace(): String replacement.
- trim(): removes the blanks at both ends of the string.
- split(): splits the string, returns an array of split strings.
- getBytes(): Returns an array of byte types for strings.
- length(): Returns the length of the string.
- toLowerCase(): Converts the string to lowercase letters.
- toUpperCase(): converts the string to uppercase characters.
- substring(): Truncates a string.
- equals(): String comparison
