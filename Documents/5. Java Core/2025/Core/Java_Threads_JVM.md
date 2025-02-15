
---
## **JVM Thread Memory Model**

### **Stack Memory vs Heap Memory**

- **Stack Memory**: Chứa **local variables**, **method call frames**, và riêng cho từng thread.
- **Heap Memory**: Chứa **objects** được chia sẻ giữa các threads, có thể gây ra vấn đề **race condition** nếu không đồng bộ đúng cách.

## **Context Switching & Performance Impact**

- Khi có nhiều threads, CPU cần chuyển đổi giữa chúng gọi là **context switching**.
- Việc này tiêu tốn tài nguyên vì CPU cần lưu trạng thái của thread cũ và tải trạng thái của thread mới.
- Dùng **Thread Pool** giúp giảm context switching không cần thiết.
# **Java Memory Model (JMM) & Synchronization**

## **1️⃣ JMM: Visibility, Atomicity, Ordering**

JMM định nghĩa cách **threads** đọc/ghi dữ liệu vào bộ nhớ.

📌 **Ba vấn đề lớn trong JMM**:

- **Visibility** – Khi một thread thay đổi giá trị, thread khác có thể không thấy.
- **Atomicity** – Một phép toán có thể bị gián đoạn giữa chừng.
- **Ordering** – Compiler có thể thay đổi thứ tự thực thi code gây ra lỗi không mong muốn.
📌 **Ví dụ về Visibility Issue (Dùng `volatile` để fix)**
```java
class VisibilityExample {
    private static volatile boolean running = true;

    public static void main(String[] args) {
        new Thread(() -> {
            while (running) {
                // Do something
            }
            System.out.println("Thread stopped");
        }).start();

        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        running = false;
        System.out.println("Main thread set running = false");
    }
}
```
👉 Nếu không có `volatile`, thread phụ có thể không thấy `running = false`.

# **📌Garbage Collection & Multi-threaded Impact**

## **1️⃣ Các loại GC trong JVM**

1. **Serial GC** – Dùng cho single-threaded application, không phù hợp cho ứng dụng lớn.
2. **Parallel GC** – Chạy nhiều thread GC đồng thời, tăng hiệu suất.
3. **CMS (Concurrent Mark-Sweep)** – Chạy đồng thời với ứng dụng, giảm STW (Stop-The-World).
4. **G1 GC (Garbage-First GC)** – Tự động tối ưu hiệu suất GC, phù hợp cho ứng dụng lớn.
5. **ZGC & Shenandoah GC** – GC thế hệ mới, gần như không có STW.