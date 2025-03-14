
---
**Sao chép nông (Shallow Copy):**  
Trong sao chép nông, kiểu tham chiếu của đối tượng được sao chép và đối tượng gốc cùng trỏ đến một đối tượng duy nhất.  
Trong ví dụ sau, có một đối tượng Person bên trong đối tượng Cat. Sau khi gọi phương thức clone, đối tượng Person của bản sao và đối tượng gốc cùng tham chiếu đến một đối tượng, đây chính là sao chép nông.
```java
public class Cat implements Cloneable {
    private String name;
    private Person owner;

    @Override
    protected Object clone() throws CloneNotSupportedException {
        return super.clone();
    }

    public static void main(String[] args) throws CloneNotSupportedException {
        Cat c = new Cat();
        Person p = new Person(18, "程序员大彬");
        c.owner = p;

        Cat cloneCat = (Cat) c.clone();
        p.setName("大彬");
        System.out.println(cloneCat.owner.getName());
    }
    // Kết quả đầu ra
    // 大彬
}
```

**Giải thích:** Vì đây là sao chép nông, cloneCat.owner và c.owner cùng trỏ đến cùng một đối tượng Person. Khi thay đổi tên của p thành "大彬", tên trong cloneCat.owner cũng thay đổi theo.

**Sao chép sâu (Deep Copy):**  
Trong sao chép sâu, kiểu tham chiếu của đối tượng được sao chép và đối tượng gốc trỏ đến các đối tượng khác nhau.  
Trong ví dụ sau, trong hàm clone, không chỉ gọi super.clone() mà còn gọi phương thức clone của đối tượng Person (giả sử Person cũng triển khai giao diện Cloneable và ghi đè phương thức clone). Nhờ đó, một bản sao sâu được thực hiện. Kết quả cho thấy giá trị của đối tượng sao chép không bị ảnh hưởng bởi đối tượng gốc.
```java
public class Cat implements Cloneable {
    private String name;
    private Person owner;

    @Override
    protected Object clone() throws CloneNotSupportedException {
        Cat c = null;
        c = (Cat) super.clone();
        c.owner = (Person) owner.clone(); // Sao chép đối tượng Person
        return c;
    }

    public static void main(String[] args) throws CloneNotSupportedException {
        Cat c = new Cat();
        Person p = new Person(18, "程序员大彬");
        c.owner = p;

        Cat cloneCat = (Cat) c.clone();
        p.setName("大彬");
        System.out.println(cloneCat.owner.getName());
    }
    // Kết quả đầu ra
    // 程序员大彬
}
```

**Giải thích:** Vì đây là sao chép sâu, cloneCat.owner là một bản sao độc lập của c.owner. Khi thay đổi tên của p thành "大彬", tên trong cloneCat.owner vẫn giữ nguyên giá trị ban đầu "程序员大彬".