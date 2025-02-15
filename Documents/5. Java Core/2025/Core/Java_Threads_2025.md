
---
## Java Threads - Tổng Quan

### 1. **Thread là gì?**

- Trong Java, **thread** là đơn vị thực thi nhỏ nhất trong một chương trình.
- Java hỗ trợ **đa luồng (multithreading)**, giúp chương trình chạy song song nhiều tác vụ thay vì chạy tuần tự.

### 2. **Cách tạo một Thread trong Java**

Có hai cách chính để tạo một thread trong Java:

1. **Kế thừa `Thread` class**
2. **Triển khai `Runnable` interface**

#### **Cách 1: Kế thừa Thread Class**
```java
class MyThread extends Thread {
    @Override
    public void run() {
        System.out.println("Thread is running...");
    }
}

public class Main {
    public static void main(String[] args) {
        MyThread thread = new MyThread();
        thread.start();  // Bắt đầu luồng
    }
}
```
💡 **Lưu ý**: Không gọi `thread.run()` trực tiếp, vì như vậy nó sẽ chạy trên main thread chứ không tạo luồng mới.

Cách 2: Triển khai Runnable Interface (Tốt hơn)
```java
class MyRunnable implements Runnable {
    @Override
    public void run() {
        System.out.println("Runnable thread is running...");
    }
}

public class Main {
    public static void main(String[] args) {
        Thread thread = new Thread(new MyRunnable());
        thread.start();
    }
}
```
✔ **Ưu điểm của Runnable**:

- Java không hỗ trợ đa kế thừa, nên dùng Runnable giúp bạn có thể kế thừa class khác.
- Tách biệt rõ ràng giữa logic và luồng.

### 3. **Các trạng thái của Thread (Thread Lifecycle)**

Thread trong Java có 5 trạng thái chính:

1. **NEW** – Thread được tạo nhưng chưa chạy (`new Thread()`).
2. **RUNNABLE** – Thread đã sẵn sàng chạy, chờ CPU cấp phát (`start()`).
3. **BLOCKED** – Thread bị chặn do một luồng khác giữ lock.
4. **WAITING** – Thread đợi vô thời hạn đến khi có signal (`wait()`).
5. **TIMED_WAITING** – Thread đợi trong khoảng thời gian nhất định (`sleep(time)`, `join(time)`).
6. **TERMINATED** – Thread kết thúc sau khi thực thi xong.

🔹 **Ví dụ minh họa các trạng thái**:
```java
class MyThread extends Thread {
    @Override
    public void run() {
        try {
            Thread.sleep(2000); // Chuyển sang TIMED_WAITING
            System.out.println("Thread is running...");
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        Thread thread = new MyThread();
        System.out.println("State: " + thread.getState()); // NEW

        thread.start();
        System.out.println("State: " + thread.getState()); // RUNNABLE

        Thread.sleep(500);
        System.out.println("State: " + thread.getState()); // TIMED_WAITING

        thread.join();
        System.out.println("State: " + thread.getState()); // TERMINATED
    }
}
```
### 4. **Thread.sleep() vs Thread.yield() vs Thread.join()**

|**Phương thức**|**Ý nghĩa**|
|---|---|
|`Thread.sleep(ms)`|Tạm dừng thread trong thời gian ms (mili giây).|
|`Thread.yield()`|Gợi ý để CPU chuyển quyền cho thread khác (không đảm bảo).|
|`Thread.join()`|Chờ một thread khác hoàn thành trước khi tiếp tục.|

🔹 **Ví dụ `join()`** – Đảm bảo thread chính chờ thread con chạy xong:
```java
class MyThread extends Thread {
    @Override
    public void run() {
        try {
            Thread.sleep(2000);
            System.out.println("Thread finished!");
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        Thread thread = new MyThread();
        thread.start();
        thread.join(); // Chờ thread chạy xong
        System.out.println("Main thread tiếp tục sau khi thread con hoàn thành.");
    }
}
```
### 5. **Thread Priority (Độ ưu tiên của Thread)**

Java có mức độ ưu tiên từ `1` (thấp nhất) đến `10` (cao nhất), mặc định là `5`.
```java
Thread thread1 = new Thread(() -> System.out.println("Thread 1"));
Thread thread2 = new Thread(() -> System.out.println("Thread 2"));

thread1.setPriority(Thread.MIN_PRIORITY); // 1
thread2.setPriority(Thread.MAX_PRIORITY); // 10

thread1.start();
thread2.start();
```
💡 **Lưu ý**:

- Không đảm bảo thread có priority cao sẽ chạy trước (phụ thuộc vào JVM và hệ điều hành).


1. **Thread Interruption (Dừng Thread đúng cách)**
```java
class MyThread extends Thread {
    @Override
    public void run() {
        while (!Thread.currentThread().isInterrupted()) {
            System.out.println("Thread is running...");
            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                System.out.println("Thread bị interrupt!");
                return; // Thoát khỏi vòng lặp
            }
        }
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        MyThread thread = new MyThread();
        thread.start();
        Thread.sleep(3000);
        thread.interrupt(); // Gửi tín hiệu dừng thread
    }
}
```
✔ **Tốt hơn `stop()` vì không gây mất trạng thái của thread**.

### 7. **Thread Synchronization (Đồng bộ hóa luồng)**

Nếu nhiều thread cùng truy cập vào một biến, có thể xảy ra race condition. Giải pháp: **synchronized**.

```java
class Counter {
    private int count = 0;

    public synchronized void increment() {
        count++;
    }

    public int getCount() {
        return count;
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        Counter counter = new Counter();

        Thread t1 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) counter.increment();
        });

        Thread t2 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) counter.increment();
        });

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("Final count: " + counter.getCount()); // 2000
    }
}
```
✔ **Lưu ý**: Nếu không có `synchronized`, kết quả có thể nhỏ hơn 2000 do race condition.
## **Tóm Tắt**

| Chủ đề              | Nội dung                                                        |
| ------------------- | --------------------------------------------------------------- |
| **Cách tạo Thread** | `Thread` class hoặc `Runnable` interface (nên dùng Runnable).   |
| **Thread States**   | NEW → RUNNABLE → BLOCKED → WAITING → TIMED_WAITING → TERMINATED |
| **Quản lý Thread**  | `sleep()`, `yield()`, `join()`, `interrupt()`                   |
| **Đồng bộ hóa**     | Dùng `synchronized` để tránh race condition.                    |
| **Ưu tiên Thread**  | `setPriority()`, nhưng không đảm bảo thực thi đúng thứ tự.      |