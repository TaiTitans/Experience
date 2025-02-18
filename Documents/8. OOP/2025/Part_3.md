
---
## **1. Constructor và Destructor là gì?**

- **Constructor** là một phương thức đặc biệt trong một lớp, được gọi tự động khi một đối tượng được tạo ra. Nó được sử dụng để **khởi tạo (initialize)** các giá trị cho đối tượng.
- **Destructor** là một phương thức đặc biệt trong một lớp, được gọi tự động khi một đối tượng bị hủy. Nó được dùng để **giải phóng tài nguyên** như bộ nhớ, kết nối file hoặc kết nối database.

📌 **Lưu ý:**

- **Java không có Destructor** vì Java có **Garbage Collector** lo việc dọn dẹp bộ nhớ. Trong C++ thì Destructor được sử dụng.

## **2. Cách hoạt động của Constructor và Destructor**

### **Constructor hoạt động như thế nào?**

- Khi một đối tượng được tạo bằng `new`, constructor sẽ chạy đầu tiên.
- Constructor có thể **gán giá trị mặc định** hoặc **chạy logic khởi tạo**.
- Nếu không có constructor, Java sẽ tự động tạo một **constructor mặc định không tham số**.

📌 **Ví dụ về Constructor trong Java:**
```java
class Xe {
    String ten;
    
    // Constructor
    Xe(String ten) {
        this.ten = ten;
        System.out.println("Xe được tạo: " + ten);
    }
}

public class Main {
    public static void main(String[] args) {
        Xe toyota = new Xe("Toyota"); // Constructor tự động được gọi
    }
}
```

## **3. Quy tắc và loại Constructor**

### **Quy tắc của Constructor**

1. **Tên constructor phải trùng với tên lớp.**
2. **Không có kiểu trả về (không dùng `void`, `int`, etc.).**
3. **Có thể có tham số hoặc không.**
4. **Có thể có nhiều constructor trong cùng một lớp (Constructor Overloading).**

### **Các loại Constructor trong Java**

|**Loại**|**Định nghĩa**|
|---|---|
|**Default Constructor**|Không có tham số, dùng để khởi tạo giá trị mặc định.|
|**Parameterized Constructor**|Có tham số, dùng để khởi tạo giá trị cụ thể.|
|**Copy Constructor**|(Không có trong Java, chỉ có trong C++), tạo bản sao của một đối tượng.|
|**Private Constructor**|Giới hạn việc tạo object, thường dùng trong Singleton Pattern.|
## **4. Mục đích của Constructor và Destructor**

|**Constructor**|**Destructor (Chỉ trong C++)**|
|---|---|
|**Dùng để khởi tạo giá trị** khi tạo đối tượng.|**Dùng để giải phóng tài nguyên** khi đối tượng bị hủy.|
|Chạy **tự động** khi `new` được gọi.|Chạy **tự động** khi đối tượng ra khỏi phạm vi.|
|Có thể **quá tải (overload)** constructor.|Không thể overload destructor.|
|Java có constructor.|Java không có destructor, C++ có.|

## **5. Sự khác biệt giữa Constructor và Method**

|**Constructor**|**Method**|
|---|---|
|Có **tên giống với class**.|Có thể đặt bất kỳ tên nào.|
|**Không có kiểu trả về**.|Có thể có kiểu trả về (`void`, `int`, etc.).|
|Chạy **tự động** khi tạo object.|Chạy **khi được gọi**.|
|Dùng để **khởi tạo đối tượng**.|Dùng để **thực hiện hành động**.|

## **6. Constructor có thể làm gì ngoài khởi tạo không?**

Có! Constructor có thể:

1. **Tạo kết nối Database.**
2. **Mở tệp tin.**
3. **Chạy luồng nền (Background Thread).**
4. **Đếm số lượng object tạo ra.**

## **Có Constructor Class không?**

Không có **Constructor Class** nhưng có thể tạo **Class chỉ chứa Constructor**, ví dụ **Singleton Class**.

📌 **Singleton với Private Constructor:**
```java
class Singleton {
    private static Singleton instance;

    private Singleton() {} // Private constructor

    public static Singleton getInstance() {
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }
}

public class Main {
    public static void main(String[] args) {
        Singleton obj1 = Singleton.getInstance();
        Singleton obj2 = Singleton.getInstance();

        System.out.println(obj1 == obj2); // Output: true
    }
}
```

