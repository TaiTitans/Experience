
---
Trong Java, nếu một biến thành viên được khai báo với từ khóa `transient`, giá trị của nó **sẽ không được bảo toàn khi đối tượng được tuần tự hóa**.

Điều này có nghĩa là **các biến được đánh dấu `transient` sẽ bị bỏ qua trong quá trình tuần tự hóa**, và khi đối tượng được giải tuần tự hóa, **giá trị của chúng sẽ được đặt về giá trị mặc định của kiểu dữ liệu tương ứng**:

- `int`, `long`, `double`, `float` → `0`
- `boolean` → `false`
- `char` → `'\u0000'`
- `Object`, `String` → `null`

Ví dụ:
```java
import java.io.*;

class User implements Serializable {
    private static final long serialVersionUID = 1L;
    
    private String username;
    private transient String password; // Không được tuần tự hóa

    public User(String username, String password) {
        this.username = username;
        this.password = password;
    }

    @Override
    public String toString() {
        return "User{username='" + username + "', password='" + password + "'}";
    }
}

public class TransientExample {
    public static void main(String[] args) throws IOException, ClassNotFoundException {
        User user = new User("john_doe", "123456");

        // Serialize object
        ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream("user.ser"));
        oos.writeObject(user);
        oos.close();

        // Deserialize object
        ObjectInputStream ois = new ObjectInputStream(new FileInputStream("user.ser"));
        User deserializedUser = (User) ois.readObject();
        ois.close();

        System.out.println(deserializedUser); 
        // Output: User{username='john_doe', password='null'}
    }
}
```

Trong ví dụ trên, sau khi giải tuần tự hóa, biến `password` có giá trị `null` vì nó được đánh dấu là `transient`.