

---

## **1. Tổng quan về Java I/O**

- **Khái niệm:**
    - Java I/O cung cấp các lớp và phương thức để thực hiện nhập (Input) và xuất (Output) dữ liệu.
    - Giao diện chính: `java.io` và `java.nio`.
- **Hai loại stream:**
    - **Byte Stream:** Làm việc với dữ liệu dạng byte (`InputStream`, `OutputStream`).
    - **Character Stream:** Làm việc với dữ liệu dạng ký tự (`Reader`, `Writer`).
## **2. Byte Stream**

### **a. InputStream và các lớp con**

- Đọc dữ liệu dạng byte (8-bit).
- Các lớp quan trọng:
    - **FileInputStream:** Đọc dữ liệu từ file.
    - **BufferedInputStream:** Cải thiện hiệu năng bằng cách thêm bộ đệm.
    - **DataInputStream:** Đọc dữ liệu nguyên thủy (int, float, double, ...).

```java
import java.io.FileInputStream;
import java.io.IOException;

public class ByteStreamExample {
    public static void main(String[] args) {
        try (FileInputStream fis = new FileInputStream("example.txt")) {
            int data;
            while ((data = fis.read()) != -1) {
                System.out.print((char) data);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

```

### **b. OutputStream và các lớp con**

- Ghi dữ liệu dạng byte.
- Các lớp quan trọng:
    - **FileOutputStream:** Ghi dữ liệu vào file.
    - **BufferedOutputStream:** Cải thiện hiệu năng bằng cách thêm bộ đệm.
    - **DataOutputStream:** Ghi dữ liệu nguyên thủy.

```java
import java.io.FileOutputStream;
import java.io.IOException;

public class ByteStreamExample {
    public static void main(String[] args) {
        String data = "Hello, Java I/O!";
        try (FileOutputStream fos = new FileOutputStream("example.txt")) {
            fos.write(data.getBytes());
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

```

## **3. Character Stream**

### **a. Reader và các lớp con**

- Đọc dữ liệu dạng ký tự (16-bit Unicode).
- Các lớp quan trọng:
    - **FileReader:** Đọc tệp ký tự.
    - **BufferedReader:** Hiệu năng cao hơn bằng cách thêm bộ đệm.

```java
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class CharacterStreamExample {
    public static void main(String[] args) {
        try (BufferedReader br = new BufferedReader(new FileReader("example.txt"))) {
            String line;
            while ((line = br.readLine()) != null) {
                System.out.println(line);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

### **b. Writer và các lớp con**

- Ghi dữ liệu dạng ký tự.
- Các lớp quan trọng:
    - **FileWriter:** Ghi tệp ký tự.
    - **BufferedWriter:** Hiệu năng cao hơn nhờ bộ đệm.

```java
import java.io.BufferedWriter;
import java.io.FileWriter;
import java.io.IOException;

public class CharacterStreamExample {
    public static void main(String[] args) {
        String data = "Hello, Java Character Stream!";
        try (BufferedWriter bw = new BufferedWriter(new FileWriter("example.txt"))) {
            bw.write(data);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

```

## **4. Serialization (Tuần tự hóa dữ liệu)**

- **Khái niệm:**
    - Tuần tự hóa là quá trình chuyển đổi một đối tượng thành byte stream để lưu trữ hoặc truyền tải.
    - Các đối tượng phải implement interface `Serializable`.

```java
import java.io.*;

class Person implements Serializable {
    private static final long serialVersionUID = 1L;
    String name;
    int age;

    Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
}

public class SerializationExample {
    public static void main(String[] args) {
        Person person = new Person("John Doe", 30);

        // Serialize
        try (ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream("person.ser"))) {
            oos.writeObject(person);
        } catch (IOException e) {
            e.printStackTrace();
        }

        // Deserialize
        try (ObjectInputStream ois = new ObjectInputStream(new FileInputStream("person.ser"))) {
            Person deserializedPerson = (Person) ois.readObject();
            System.out.println("Name: " + deserializedPerson.name + ", Age: " + deserializedPerson.age);
        } catch (IOException | ClassNotFoundException e) {
            e.printStackTrace();
        }
    }
}

```

## **5. Java NIO (New I/O)**

- **Khái niệm:**
    - Làm việc không đồng bộ (Non-blocking) và hiệu quả cao hơn.
    - Các khái niệm chính:
        - **Buffer:** Lưu trữ dữ liệu.
        - **Channel:** Giao tiếp với thiết bị.
        - **Selector:** Đa luồng I/O.

```java
import java.nio.file.*;
import java.nio.charset.StandardCharsets;
import java.io.IOException;

public class NIOExample {
    public static void main(String[] args) {
        Path path = Paths.get("example.txt");
        try {
            String content = Files.readString(path, StandardCharsets.UTF_8);
            System.out.println(content);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

## **6. Tổng kết**

### **Chọn Byte Stream hay Character Stream?**

- Sử dụng **Byte Stream** khi làm việc với dữ liệu nhị phân (ảnh, video, v.v.).
- Sử dụng **Character Stream** khi làm việc với dữ liệu văn bản.

### **Tối ưu hóa:**

- Sử dụng **Buffered** stream để tăng hiệu năng.
- Sử dụng **try-with-resources** để tự động đóng tài nguyên.