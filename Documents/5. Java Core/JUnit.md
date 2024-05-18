
---
JUnit là một framework kiểm thử đơn vị (unit testing) cho ngôn ngữ lập trình Java. Nó cho phép các lập trình viên viết và chạy các bài kiểm thử tự động để kiểm tra tính đúng đắn của mã nguồn.

### Các khái niệm cơ bản trong JUnit

1. **Test Case**: Một bài kiểm thử cụ thể để kiểm tra một phần nhỏ của mã nguồn.
2. **Test Suite**: Một tập hợp các bài kiểm thử.
3. **Assertion**: Phương thức để kiểm tra các điều kiện mong muốn (ví dụ: kiểm tra xem một giá trị có đúng không).
4. **Annotations**: Các chú thích đặc biệt để xác định các phương thức kiểm thử và thiết lập trước và sau khi kiểm thử.
----

### Cài đặt JUnit

Để sử dụng JUnit, bạn cần thêm thư viện JUnit vào dự án của mình.


---

### Cấu trúc cơ bản của một bài kiểm thử trong JUnit

Dưới đây là một ví dụ đơn giản về cách viết và chạy một bài kiểm thử bằng JUnit.

#### Viết mã cần kiểm thử

Giả sử bạn có một lớp `Calculator` với phương thức `add` để cộng hai số:

```Java
public class Calculator {
    public int add(int a, int b) {
        return a + b;
    }
}

```

#### Viết bài kiểm thử

Tạo một lớp kiểm thử để kiểm tra phương thức `add` của lớp `Calculator`:

```Java
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class CalculatorTest {

    @Test
    public void testAdd() {
        Calculator calculator = new Calculator();
        int result = calculator.add(2, 3);
        assertEquals(5, result, "2 + 3 should equal 5");
    }
}

```

#### Giải thích mã nguồn kiểm thử

- `@Test`: Chú thích này đánh dấu phương thức `testAdd` là một bài kiểm thử.
- `assertEquals(expected, actual, message)`: Phương thức này kiểm tra xem giá trị thực tế (`actual`) có bằng giá trị mong đợi (`expected`) không. Nếu không, thông báo lỗi (`message`) sẽ được hiển thị.

### Các chú thích (Annotations) thường dùng trong JUnit

- `@Test`: Đánh dấu một phương thức là một bài kiểm thử.
- `@BeforeEach`: Phương thức được chạy trước mỗi bài kiểm thử.
- `@AfterEach`: Phương thức được chạy sau mỗi bài kiểm thử.
- `@BeforeAll`: Phương thức được chạy một lần trước tất cả các bài kiểm thử trong lớp.
- `@AfterAll`: Phương thức được chạy một lần sau tất cả các bài kiểm thử trong lớp.
- `@Disabled`: Bỏ qua bài kiểm thử.

### Các phương thức Assertion phổ biến

- `assertEquals(expected, actual)`: Kiểm tra xem hai giá trị có bằng nhau không.
- `assertTrue(condition)`: Kiểm tra xem điều kiện có đúng không.
- `assertFalse(condition)`: Kiểm tra xem điều kiện có sai không.
- `assertNull(object)`: Kiểm tra xem đối tượng có null không.
- `assertNotNull(object)`: Kiểm tra xem đối tượng không null.


### Cấu trúc Test:

```bash
my-project/
├── src/
│   ├── main/
│   │   ├── java/
│   │   │   └── com/
│   │   │       └── example/
│   │   │           └── MyMainClass.java
│   │   └── resources/
│   │       └── application.properties
│   └── test/
│       ├── java/
│       │   └── com/
│       │       └── example/
│       │           └── MyMainClassTest.java
│       └── resources/
│           └── test-data.json
├── pom.xml
└── README.md

```