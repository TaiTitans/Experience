# Java I/O như một Senior Java Developer

## 1. Java I/O Overview
### 🔹 Java có hai hệ thống I/O chính:
1️⃣ **Java I/O (java.io)** – **Cổ điển, blocking I/O, dùng Streams**  
2️⃣ **Java NIO (java.nio)** – **Non-blocking I/O, dùng Buffers & Channels**  

![Java IO vs NIO]()  
(*IO dùng Streams, NIO dùng Buffers & Channels.*)

---

## 2. Java I/O – Streams API (java.io)
### 📌 **Phân loại Streams**
1️⃣ **Byte Streams** – Đọc ghi dữ liệu nhị phân (`InputStream`, `OutputStream`)  
2️⃣ **Character Streams** – Đọc ghi văn bản (`Reader`, `Writer`)  
3️⃣ **Buffered Streams** – Tăng tốc I/O (`BufferedReader`, `BufferedWriter`)  
4️⃣ **Object Streams** – Serialize/Deserialize object (`ObjectInputStream`, `ObjectOutputStream`)  

### 🚀 Byte Streams (Xử lý dữ liệu nhị phân)
```java
FileInputStream fis = new FileInputStream("input.txt");
FileOutputStream fos = new FileOutputStream("output.txt");

int data;
while ((data = fis.read()) != -1) { 
    fos.write(data);
}
fis.close();
fos.close();
```

### 🚀 Buffered Streams (Tăng hiệu suất đọc ghi)
```java
BufferedReader reader = new BufferedReader(new FileReader("input.txt"));
BufferedWriter writer = new BufferedWriter(new FileWriter("output.txt"));

String line;
while ((line = reader.readLine()) != null) {
    writer.write(line);
    writer.newLine();
}
reader.close();
writer.close();
```

### 🚀 Object Streams (Serialization & Deserialization)
```java
class Person implements Serializable {
    private static final long serialVersionUID = 1L;
    String name;
    int age;

    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
}

// Ghi object vào file
ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream("person.dat"));
oos.writeObject(new Person("John", 30));
oos.close();

// Đọc object từ file
ObjectInputStream ois = new ObjectInputStream(new FileInputStream("person.dat"));
Person p = (Person) ois.readObject();
ois.close();
System.out.println(p.name + " - " + p.age);
```

---

## 3. Java NIO (Non-blocking I/O)
### 🔹 Thành phần chính của Java NIO:
| Thành phần | Mô tả |
|------------|-------|
| **Buffer** | Lưu trữ dữ liệu trước khi xử lý |
| **Channel** | Đọc/ghi dữ liệu từ Buffer |
| **Selector** | Theo dõi nhiều Channels cùng lúc |
| **FileChannel** | Đọc/Ghi file hiệu suất cao |

### 🚀 FileChannel (Đọc ghi file nhanh hơn)
```java
RandomAccessFile file = new RandomAccessFile("data.txt", "r");
FileChannel channel = file.getChannel();

ByteBuffer buffer = ByteBuffer.allocate(1024);
channel.read(buffer);
buffer.flip();

while (buffer.hasRemaining()) {
    System.out.print((char) buffer.get());
}
channel.close();
file.close();
```

### 🚀 Memory-mapped File (Tăng tốc xử lý file lớn)
```java
RandomAccessFile file = new RandomAccessFile("largefile.txt", "rw");
FileChannel channel = file.getChannel();

// Ánh xạ file vào bộ nhớ
MappedByteBuffer buffer = channel.map(FileChannel.MapMode.READ_WRITE, 0, channel.size());

while (buffer.hasRemaining()) {
    System.out.print((char) buffer.get());
}
channel.close();
file.close();
```

---

## 4. Java File API (java.nio.file)
```java
Path path = Paths.get("data.txt");

// Đọc toàn bộ file
List<String> lines = Files.readAllLines(path);
lines.forEach(System.out::println);

// Ghi file
Files.write(path, "Hello Java NIO".getBytes(), StandardOpenOption.APPEND);
```

---

## 5. So sánh Java IO vs Java NIO

| Tiêu chí | Java I/O (Blocking) | Java NIO (Non-blocking) |
|----------|---------------------|-------------------------|
| **Cách hoạt động** | Streams | Buffers |
| **Tốc độ** | Chậm hơn | Nhanh hơn |
| **Blocking?** | Có | Không |
| **Ứng dụng** | Đọc file nhỏ | Server, xử lý file lớn |

---

## 6. Best Practices cho Java I/O
✅ **Dùng `BufferedReader/BufferedWriter` thay vì `FileReader/FileWriter`**  
✅ **Dùng `StringBuilder` thay vì `String` khi xử lý nhiều chuỗi**  
✅ **Dùng `NIO FileChannel` hoặc `Memory-mapped file` cho file lớn**  
✅ **Dùng `try-with-resources` để tự động đóng Streams**  
```java
try (BufferedReader br = new BufferedReader(new FileReader("input.txt"))) {
    System.out.println(br.readLine());
} catch (IOException e) {
    e.printStackTrace();
}
```

---

## 7. Kết luận
- **Java I/O:** Dễ dùng nhưng chậm, phù hợp xử lý file nhỏ.
- **Java NIO:** Non-blocking, nhanh, phù hợp xử lý file lớn, server, WebSockets.
- **Tối ưu bằng Buffer, Channels, Memory-mapped File.**
---

