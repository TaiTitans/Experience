
---
Chương này tập trung vào **bảo toàn dữ liệu (memory consistency), đồng bộ hóa (synchronization), và mô hình bộ nhớ của Java (Java Memory Model - JMM)**. Đây là nền tảng quan trọng giúp Java hỗ trợ **lập trình đa luồng (multithreading)** an toàn và hiệu quả.

`LƯU Ý: SẼ CẦN MỘT BÀI VIẾT KHÁC ĐÀO SÂU HƠN VỀ CÁC CHỦ ĐỀ THREADS, PROCESS, ĐA LUỒNg, ĐỒNG BỘ VÀ BẤT ĐỒNG BỘ HAY LOCKING. NHỮNG CHỦ ĐỀ NÀY RẤT PHỨC TẠP KHÔNG THỂ NÓI ĐƠN GIẢN TRONG VÀI CÂU.`
## **1️⃣ Các khái niệm quan trọng trong Java Memory Model (JMM)**

### **🔹 Biến dùng chung (Shared Variables)**

Khi nhiều luồng cùng truy cập một biến, cần đảm bảo dữ liệu được **đọc/ghi một cách an toàn**. Nếu không đồng bộ hóa, một luồng có thể đọc dữ liệu chưa cập nhật từ bộ nhớ đệm của CPU, dẫn đến lỗi.

📌 **Ví dụ lỗi khi không đồng bộ hóa biến dùng chung:**
```java
class SharedResource {
    static int count = 0; // Biến dùng chung không đồng bộ

    static void increment() {
        count++; // Không đảm bảo an toàn trong môi trường đa luồng
    }
}
```
### **🔹 Happens-Before (Quan Hệ Xảy Ra Trước)**

JMM quy định thứ tự thực thi trong đa luồng bằng **happens-before**. Nếu một hành động **happens-before** hành động khác, giá trị từ hành động trước **phải** được nhìn thấy bởi hành động sau.

📌 **Các quy tắc happens-before quan trọng:**

- **Khởi tạo đối tượng** (`new`) xảy ra trước bất kỳ truy cập nào vào đối tượng đó.
- **Gọi `start()` trên một thread** xảy ra trước khi thread đó thực thi.
- **Gọi `join()` trên một thread** xảy ra trước khi `join()` kết thúc.
- **Ghi vào biến `volatile`** xảy ra trước khi bất kỳ luồng nào đọc nó.
## **2️⃣ Từ khóa `volatile`**

`volatile` đảm bảo mọi luồng đều thấy giá trị **mới nhất** của biến.

📌 **Ví dụ không dùng `volatile` (có thể gặp lỗi):**
```java
class Shared {
    static boolean flag = false;

    static void waitForFlag() {
        while (!flag) {
            // Nếu CPU cache biến `flag`, vòng lặp có thể bị kẹt vô hạn
        }
        System.out.println("Flag changed!");
    }
}
```
📌 **Giải pháp với `volatile`:**
```java
class Shared {
    static volatile boolean flag = false; // Đảm bảo luôn đọc giá trị mới nhất

    static void waitForFlag() {
        while (!flag) {
            // Giờ sẽ thoát khi flag được cập nhật
        }
        System.out.println("Flag changed!");
    }
}
```
💡 **Lưu ý:** `volatile` **không** thay thế được `synchronized` khi nhiều luồng cần cập nhật biến cùng lúc!

## **3️⃣ Đồng bộ hóa với `synchronized`**

`volatile` chỉ giúp **đọc/ghi an toàn**, nhưng không đảm bảo **cập nhật dữ liệu đồng thời**. Khi nhiều luồng cùng thay đổi một biến, ta cần **`synchronized`**.

📌 **Ví dụ lỗi khi không đồng bộ hóa:**
```java
class Counter {
    private int count = 0;

    void increment() {
        count++; // Không an toàn khi nhiều luồng cùng truy cập
    }

    int getCount() {
        return count;
    }
}
```
📌 **Cách sửa với `synchronized`:**
```java
class Counter {
    private int count = 0;

    synchronized void increment() { // Đảm bảo chỉ có 1 luồng chạy tại 1 thời điểm
        count++;
    }

    synchronized int getCount() {
        return count;
    }
}
```
## **4️⃣ Đồng bộ hóa bằng `Lock` (ReentrantLock)**

`Lock` giúp kiểm soát đồng bộ hóa linh hoạt hơn `synchronized`.

📌 **Ví dụ sử dụng `ReentrantLock`:**
```java
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

class Counter {
    private int count = 0;
    private final Lock lock = new ReentrantLock();

    void increment() {
        lock.lock(); // Lấy khóa
        try {
            count++;
        } finally {
            lock.unlock(); // Giải phóng khóa
        }
    }

    int getCount() {
        return count;
    }
}
```
💡 **Ưu điểm của `Lock` so với `synchronized`:**

- **Thử lấy khóa (`tryLock()`)** mà không bị chặn.
- **Giảm độ trễ** nếu khóa đã được giữ bởi luồng khác.
- **Hỗ trợ nhiều loại khóa** như **fair lock** (ưu tiên luồng chờ lâu hơn).

## **Biến `final` trong môi trường đa luồng**

Một biến `final` nếu được gán giá trị **trong constructor** và không bị thay đổi, Java đảm bảo giá trị đó **luôn thấy đúng** trong các luồng khác mà không cần `volatile`.

## **6️⃣ Thread Safety với `ThreadLocal`**

Nếu mỗi luồng cần một **bản sao riêng** của biến thay vì dùng chung, `ThreadLocal` là giải pháp tối ưu.

📌 **Ví dụ sử dụng `ThreadLocal`:**
```java
class Example {
    static ThreadLocal<Integer> threadLocalValue = ThreadLocal.withInitial(() -> 0);

    void increment() {
        threadLocalValue.set(threadLocalValue.get() + 1);
    }

    int getValue() {
        return threadLocalValue.get();
    }
}
```
💡 **Lợi ích của `ThreadLocal`:**

- Mỗi luồng có **bản sao riêng** của biến → **tránh lỗi xung đột**.
- **Hiệu suất cao hơn** vì không cần `synchronized`.

# **📌 7. Tổng kết**

✅ **Java Memory Model (JMM)** đảm bảo dữ liệu giữa các luồng được nhất quán.  
✅ **Happens-before** xác định thứ tự thực thi giữa các luồng.  
✅ **`volatile`** đảm bảo **đọc/ghi an toàn**, nhưng không tránh được **cập nhật xung đột**.  
✅ **`synchronized`** hoặc **`Lock`** giúp bảo vệ dữ liệu khi **nhiều luồng cùng thay đổi**.  
✅ **`ThreadLocal`** giúp mỗi luồng có bản sao riêng, tránh xung đột.
