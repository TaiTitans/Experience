
---

Trong Spring Framework, `@Controller`, `@RestController`, và `@ResponseBody` là các annotation dùng để định nghĩa các lớp và phương thức điều khiển trong ứng dụng web. Dưới đây là sự khác biệt và cách sử dụng của từng annotation:

### 1. **`@Controller`**:

- Được sử dụng để đánh dấu một lớp như là một Spring MVC Controller, chịu trách nhiệm xử lý các yêu cầu HTTP và trả về một ModelAndView hoặc một tên view để hiển thị giao diện cho người dùng.
- **Thích hợp cho ứng dụng web trả về các trang HTML**.
```java
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class MyController {

    @GetMapping("/welcome")
    public String welcome(Model model) {
        model.addAttribute("message", "Welcome to Spring MVC!");
        return "welcome"; // Trả về tên của view (welcome.html hoặc welcome.jsp)
    }
}

```
### 2. **`@RestController`**:

- Kết hợp `@Controller` và `@ResponseBody`. Tất cả các phương thức trong lớp được đánh dấu bằng `@RestController` sẽ trả về dữ liệu trực tiếp dưới dạng JSON hoặc XML (hoặc định dạng khác) thay vì trả về một tên view.
- **Thích hợp cho việc xây dựng các RESTful API**.
```java
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class MyRestController {

    @GetMapping("/api/message")
    public String getMessage() {
        return "Hello, World!"; // Trả về dữ liệu dạng chuỗi, không phải tên view
    }
}
```
### 3. **`@ResponseBody`**:

- Dùng để chỉ ra rằng dữ liệu trả về từ phương thức sẽ được viết trực tiếp vào body của HTTP response, thường dưới dạng JSON hoặc XML, thay vì trả về tên của một view.
- Có thể được sử dụng trên các phương thức của lớp `@Controller` để trả về dữ liệu trực tiếp, mà không cần kết hợp với view.

### **Khi nào nên sử dụng?**

- **`@Controller`**: Khi bạn xây dựng một ứng dụng web cần trả về các trang HTML hoặc giao diện người dùng.
- **`@RestController`**: Khi bạn xây dựng RESTful APIs hoặc dịch vụ web trả về dữ liệu JSON/XML.
- **`@ResponseBody`**: Khi bạn chỉ cần trả về dữ liệu trực tiếp từ một hoặc vài phương thức cụ thể trong một lớp `@Controller`.

