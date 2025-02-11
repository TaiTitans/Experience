
---
API **Preferences** trong Java cung cấp một cách tiêu chuẩn để các ứng dụng lưu trữ và truy xuất dữ liệu cấu hình và tùy chọn của người dùng. Điều này cho phép các ứng dụng điều chỉnh hành vi dựa trên nhu cầu của từng người dùng và môi trường cụ thể. Gói `java.util.prefs` chứa các lớp và giao diện cần thiết cho việc quản lý dữ liệu này.

**Cấu trúc của Preferences API:**

- **Cây nút (node) ưu tiên:** Dữ liệu ưu tiên được tổ chức dưới dạng cây, với hai cây riêng biệt: một cho ưu tiên của người dùng và một cho ưu tiên của hệ thống. Mỗi nút trong cây đại diện cho một tập hợp các cặp khóa-giá trị.

**Các tính năng chính:**

- **Lưu trữ bất đồng bộ:** Các phương thức thay đổi dữ liệu ưu tiên có thể hoạt động bất đồng bộ, tức là chúng có thể trả về ngay lập tức và các thay đổi sẽ được truyền đến bộ lưu trữ nền theo thời gian. Phương thức `flush` có thể được sử dụng để buộc các thay đổi được ghi ngay lập tức.
    
- **An toàn cho luồng (thread-safe):** Các phương thức trong lớp `Preferences` có thể được gọi đồng thời bởi nhiều luồng trong một JVM mà không cần đồng bộ hóa bên ngoài. Kết quả sẽ tương đương với một thực thi tuần tự nào đó. Tuy nhiên, nếu lớp này được sử dụng đồng thời bởi nhiều JVM lưu trữ dữ liệu ưu tiên trong cùng một bộ lưu trữ nền, dữ liệu sẽ không bị hỏng, nhưng không có đảm bảo nào khác về tính nhất quán của dữ liệu ưu tiên.
    

**So sánh với các cơ chế khác:**

Trước khi có Preferences API, các nhà phát triển thường lưu trữ dữ liệu cấu hình trong các tệp properties thông qua API `java.util.Properties`. Mặc dù cách tiếp cận này đơn giản, nhưng nó thiếu tính linh hoạt và không hỗ trợ lưu trữ nền độc lập. Preferences API cung cấp một giải pháp linh hoạt hơn, cho phép lưu trữ dữ liệu trong các kho lưu trữ khác nhau tùy thuộc vào triển khai.
**Sử dụng cơ bản:**

Để lưu trữ và truy xuất một giá trị ưu tiên, bạn có thể sử dụng lớp `Preferences` như sau:
```java
import java.util.prefs.Preferences;

public class PreferenceExample {
    public static void main(String[] args) {
        // Lấy đối tượng Preferences cho lớp hiện tại
        Preferences prefs = Preferences.userNodeForPackage(PreferenceExample.class);

        // Lưu trữ một giá trị
        prefs.put("username", "john_doe");

        // Truy xuất một giá trị với giá trị mặc định nếu khóa không tồn tại
        String username = prefs.get("username", "default_user");

        System.out.println("Username: " + username);
    }
}
```
Trong ví dụ trên, chúng ta lưu trữ một giá trị chuỗi với khóa "username" và sau đó truy xuất nó. Nếu khóa không tồn tại, giá trị mặc định "default_user" sẽ được trả về.

Để biết thêm chi tiết và hướng dẫn sử dụng, bạn có thể tham khảo tài liệu chính thức của Oracle về [Preferences API](https://docs.oracle.com/en/java/javase/21/core/preferences-api1.html).