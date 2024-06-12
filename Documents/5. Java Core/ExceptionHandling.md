
---

	Exception Handling trong Java là một cơ chế mạnh mẽ để xử lý các lỗi runtime một cách hiệu quả. Cơ chế này giúp chương trình không bị dừng đột ngột và cung cấp các thông báo lỗi cụ thể để dễ dàng xác định và sửa lỗi.

### Các loại Exception

Java chia các exception thành hai loại chính:

1. **Checked Exception**: Đây là các lỗi mà trình biên dịch Java bắt buộc bạn phải xử lý. Nếu không, chương trình sẽ không biên dịch được. Ví dụ: `IOException`, `SQLException`.
2. **Unchecked Exception**: Đây là các lỗi mà trình biên dịch Java không bắt buộc bạn phải xử lý. Chúng thường là các lỗi về lập trình, như lỗi logic hoặc lỗi thời gian chạy. Ví dụ: `ArithmeticException`, `NullPointerException`, `ArrayIndexOutOfBoundsException`.

Ngoài ra còn có `Error` là các lỗi mà thông thường bạn không xử lý được, ví dụ như `OutOfMemoryError`.

---
### Cấu trúc cơ bản của Exception Handling

Java cung cấp bốn từ khóa chính để xử lý exception: `try`, `catch`, `finally`, và `throw`.

#### Try và Catch

- **`try`**: Bao bọc đoạn mã có thể gây ra exception.
- **`catch`**: Được sử dụng để xử lý exception xảy ra trong khối `try`.

```Java
try {
    int division = 10 / 0; // Sẽ ném ra ArithmeticException
} catch (ArithmeticException e) {
    System.out.println("Cannot divide by zero: " + e.getMessage());
}

```
#### Finally

- **`finally`**: Khối mã này luôn được thực thi dù exception có xảy ra hay không. Thường dùng để giải phóng tài nguyên như đóng file, đóng kết nối cơ sở dữ liệu, v.v.

```Java
try {
    int division = 10 / 0;
} catch (ArithmeticException e) {
    System.out.println("Cannot divide by zero: " + e.getMessage());
} finally {
    System.out.println("This block is always executed.");
}

```


#### Throw và Throws

- **`throw`**: Từ khóa này được sử dụng để ném một exception rõ ràng từ một khối mã hoặc phương thức.
- **`throws`**: Được sử dụng trong khai báo phương thức để chỉ ra rằng phương thức có thể ném ra các exception nào.

```Java
public void checkAge(int age) throws Exception {
    if (age < 18) {
        throw new Exception("Age is not valid to vote");
    }
}

try {
    checkAge(15);
} catch (Exception e) {
    System.out.println(e.getMessage());
}

```

