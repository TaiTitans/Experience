
---
# **📌 Phần 1: `synchronized` – Cơ chế đồng bộ cơ bản**

## **1️⃣ `synchronized` là gì?**

- `synchronized` đảm bảo rằng **chỉ một thread** có thể truy cập vào một đoạn code cụ thể tại một thời điểm.
- Java cung cấp ba cách sử dụng `synchronized`:
    1. **Synchronized method**
    2. **Synchronized block**
    3. **Synchronized static method**

2️⃣ Ví dụ: `synchronized` method
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

public class SynchronizedExample {
    public static void main(String[] args) throws InterruptedException {
        Counter counter = new Counter();

        Runnable task = () -> {
            for (int i = 0; i < 1000; i++) {
                counter.increment();
            }
        };

        Thread t1 = new Thread(task);
        Thread t2 = new Thread(task);

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("Final count: " + counter.getCount()); // Kết quả luôn là 2000
    }
}
```
👉 **Giải thích:** `synchronized` đảm bảo chỉ một thread có thể gọi `increment()` cùng một lúc.

## **3️⃣ `synchronized` block – Đồng bộ hóa chỉ một phần code**
```java
class Counter {
    private int count = 0;

    public void increment() {
        synchronized (this) { // Chỉ đồng bộ hóa khối này
            count++;
        }
    }

    public int getCount() {
        return count;
    }
}
```
👉 **Ưu điểm:** Hạn chế phạm vi đồng bộ hóa giúp **tăng hiệu suất** vì không khóa toàn bộ phương thức.

4️⃣ `synchronized` trên static method
```java
class StaticCounter {
    private static int count = 0;

    public static synchronized void increment() { // Khóa ở cấp class
        count++;
    }

    public static int getCount() {
        return count;
    }
}
```
👉 **Static method synchronized khóa ở cấp `Class`, thay vì instance**

# **📌 Phần 2: Locks – Kiểm soát đồng bộ nâng cao**

## **1️⃣ `Lock` vs `synchronized`**

|`synchronized`|`Lock` (ReentrantLock)|
|---|---|
|Dễ sử dụng|Cung cấp kiểm soát chi tiết hơn|
|Tự động giải phóng|Phải `unlock()` thủ công|
|Không thể kiểm tra nếu một thread có lock|Có thể kiểm tra trạng thái lock|
2️⃣ Ví dụ: `ReentrantLock`
```java
import java.util.concurrent.locks.ReentrantLock;

class Counter {
    private int count = 0;
    private final ReentrantLock lock = new ReentrantLock();

    public void increment() {
        lock.lock();
        try {
            count++;
        } finally {
            lock.unlock(); // Luôn giải phóng lock
        }
    }

    public int getCount() {
        return count;
    }
}

public class LockExample {
    public static void main(String[] args) throws InterruptedException {
        Counter counter = new Counter();
        Runnable task = () -> {
            for (int i = 0; i < 1000; i++) {
                counter.increment();
            }
        };

        Thread t1 = new Thread(task);
        Thread t2 = new Thread(task);

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("Final count: " + counter.getCount());
    }
}
```
👉 **Giải thích:** `ReentrantLock` giúp kiểm soát lock chi tiết hơn `synchronized`, nhưng bạn phải `unlock()` thủ công.

3️⃣ `tryLock()` – Tránh deadlock
```java
if (lock.tryLock()) { 
    try {
        // Thực hiện công việc
    } finally {
        lock.unlock();
    }
} else {
    System.out.println("Không thể lấy lock, thử lại sau");
}
```
👉 **Tránh chờ vô hạn nếu lock không khả dụng.**

# **📌 Phần 3: Atomic Variables – Giải pháp tối ưu cho biến dùng chung**

## **1️⃣ `AtomicInteger` – Thay thế `synchronized`**

Nếu chỉ cần cập nhật giá trị đơn giản (tăng, giảm, cập nhật), `AtomicInteger` nhanh hơn **`synchronized` hoặc `Lock`**.
```java
import java.util.concurrent.atomic.AtomicInteger;

class AtomicCounter {
    private AtomicInteger count = new AtomicInteger(0);

    public void increment() {
        count.incrementAndGet(); // Thread-safe
    }

    public int getCount() {
        return count.get();
    }
}

public class AtomicExample {
    public static void main(String[] args) throws InterruptedException {
        AtomicCounter counter = new AtomicCounter();
        Runnable task = () -> {
            for (int i = 0; i < 1000; i++) {
                counter.increment();
            }
        };

        Thread t1 = new Thread(task);
        Thread t2 = new Thread(task);

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("Final count: " + counter.getCount()); // Luôn đúng
    }
}
```

👉 **Nhanh hơn `synchronized` vì không cần lock, dùng CPU cache & compare-and-swap (CAS).**

# **📌 Phần 4: ForkJoinPool – Xử lý song song mạnh mẽ**

## **1️⃣ ForkJoinPool là gì?**

- **ForkJoinPool** là **ThreadPool tối ưu hóa cho tác vụ đệ quy (divide and conquer)**.
- Dùng `ForkJoinTask` để **chia nhỏ công việc thành nhiều task con** và xử lý song song.
```java
import java.util.concurrent.RecursiveTask;
import java.util.concurrent.ForkJoinPool;

class SumTask extends RecursiveTask<Integer> {
    private int start, end;
    private static final int THRESHOLD = 5;

    public SumTask(int start, int end) {
        this.start = start;
        this.end = end;
    }

    @Override
    protected Integer compute() {
        if (end - start <= THRESHOLD) {
            int sum = 0;
            for (int i = start; i <= end; i++) sum += i;
            return sum;
        } else {
            int mid = (start + end) / 2;
            SumTask leftTask = new SumTask(start, mid);
            SumTask rightTask = new SumTask(mid + 1, end);

            leftTask.fork();  // Chạy song song
            return rightTask.compute() + leftTask.join(); // Hợp nhất kết quả
        }
    }
}

public class ForkJoinExample {
    public static void main(String[] args) {
        ForkJoinPool pool = new ForkJoinPool();
        SumTask task = new SumTask(1, 20);
        int result = pool.invoke(task);
        System.out.println("Sum: " + result);
    }
}
```
👉 **`ForkJoinPool` tự động chia nhỏ công việc, tận dụng CPU đa luồng hiệu quả hơn.**

# **🚀 Tổng Kết**

### ✅ **Bạn đã học được gì?**

🔹 **`synchronized`**: Đồng bộ hóa cơ bản, dễ dùng nhưng có thể làm giảm hiệu suất.  
🔹 **`Lock` (ReentrantLock)**: Kiểm soát tốt hơn, tránh deadlock với `tryLock()`.  
🔹 **`AtomicInteger`**: Cách tối ưu để cập nhật biến đơn giản mà không cần lock.  
🔹 **`ForkJoinPool`**: Mô hình xử lý song song mạnh mẽ với thuật toán chia để trị.


## **1️⃣ Bảng So Sánh Tổng Quan**

|**Tiêu chí**|**synchronized**|**ReentrantLock**|**AtomicInteger**|**ForkJoinPool**|
|---|---|---|---|---|
|**Loại cơ chế**|Đồng bộ hóa|Khóa linh hoạt|Biến nguyên tử|Xử lý song song|
|**Cơ chế hoạt động**|Chặn thread khác truy cập|Chặn thread khác, hỗ trợ thử lock|Sử dụng Compare-And-Swap (CAS)|Chia nhỏ task và thực thi song song|
|**Hiệu suất**|Thấp khi có nhiều thread tranh chấp|Tốt hơn `synchronized` nhờ kiểm soát chi tiết|Tốt nhất cho biến đơn giản|Tối ưu cho công việc lớn cần song song|
|**Tính năng nâng cao**|Không có|Hỗ trợ `tryLock()`, `lockInterruptibly()`|Hỗ trợ update không cần lock|Hỗ trợ tự động chia nhỏ task|
|**Xử lý deadlock**|Dễ bị deadlock nếu không cẩn thận|`tryLock()` giúp tránh deadlock|Không có deadlock|Không bị deadlock do chia task tự động|
|**Độ phức tạp**|Đơn giản|Trung bình|Rất đơn giản|Phức tạp hơn|
|**Ứng dụng phù hợp**|Bảo vệ toàn bộ method hoặc block|Cần kiểm soát chi tiết hơn về lock|Khi chỉ cần cập nhật biến đơn giản|Khi xử lý dữ liệu lớn, công việc đệ quy|
## **2️⃣ Khi Nào Nên Sử Dụng Cơ Chế Nào?**

🔹 **`synchronized`** 👉 Dùng khi cần đơn giản và hiệu suất không phải là vấn đề lớn.  
🔹 **`ReentrantLock`** 👉 Dùng khi cần kiểm soát chi tiết hơn, như thử lock (`tryLock()`) hoặc hỗ trợ nhiều điều kiện chờ (`Condition`).  
🔹 **`AtomicInteger`** 👉 Dùng khi chỉ cần tăng/giảm biến số nguyên mà không cần lock, tối ưu hiệu suất.  
🔹 **`ForkJoinPool`** 👉 Dùng khi xử lý công việc lớn có thể chia nhỏ, như thuật toán đệ quy hoặc xử lý song song dữ liệu lớn.