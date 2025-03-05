
---
## **1. Giới thiệu về ExecutorService**

`ExecutorService` là một interface trong Java (`java.util.concurrent.ExecutorService`) được sử dụng để quản lý luồng (`Thread`) một cách hiệu quả. Thay vì tạo và quản lý luồng thủ công, `ExecutorService` cung cấp một cơ chế mạnh mẽ giúp thực thi các tác vụ (tasks) trong một thread pool.

## **2. Tại sao cần dùng ExecutorService?**

🔹 **Quản lý luồng tốt hơn:** Không cần tạo và hủy luồng thủ công.  
🔹 **Tối ưu hiệu suất:** Giảm chi phí tạo mới thread liên tục.  
🔹 **Hỗ trợ đa dạng kiểu tác vụ:** Runnable, Callable, Future, CompletableFuture.  
🔹 **Kiểm soát tài nguyên tốt hơn:** Giới hạn số lượng thread hoạt động.

## **3. Cách sử dụng ExecutorService**

### **3.1. Tạo ExecutorService**

Có nhiều cách tạo `ExecutorService`, tùy vào số lượng luồng mong muốn:
```java
import java.util.concurrent.*;

public class ExecutorServiceExample {
    public static void main(String[] args) {
        // 1. Executor với số luồng cố định
        ExecutorService fixedThreadPool = Executors.newFixedThreadPool(5);

        // 2. Executor với số luồng động (cached)
        ExecutorService cachedThreadPool = Executors.newCachedThreadPool();

        // 3. Executor với một luồng duy nhất (single thread)
        ExecutorService singleThreadExecutor = Executors.newSingleThreadExecutor();

        // 4. Scheduled Executor (Chạy task theo lịch trình)
        ScheduledExecutorService scheduledExecutor = Executors.newScheduledThreadPool(2);

        // Đừng quên shutdown sau khi sử dụng xong!
    }
}
```
📌 **Lưu ý:**

- `newFixedThreadPool(n)`: Tạo pool với `n` thread, tái sử dụng thread cũ sau khi hoàn thành task.
- `newCachedThreadPool()`: Tạo số lượng thread tùy ý, tốt cho task ngắn và nhanh.
- `newSingleThreadExecutor()`: Chạy từng task tuần tự.
- `newScheduledThreadPool(n)`: Dùng để chạy task theo lịch trình.

## **4. Gửi task vào ExecutorService**

Có nhiều cách để gửi task:

### **4.1. submit() với Runnable**
```java
ExecutorService executor = Executors.newFixedThreadPool(3);

Runnable task = () -> {
    System.out.println(Thread.currentThread().getName() + " is running...");
};

executor.submit(task);
executor.shutdown();
```

🔹 `Runnable` không trả về kết quả.

### 4.2. submit() với Callable (có giá trị trả về)
```java
ExecutorService executor = Executors.newFixedThreadPool(3);

Callable<String> task = () -> {
    Thread.sleep(1000);
    return "Task completed by " + Thread.currentThread().getName();
};

Future<String> future = executor.submit(task);

try {
    System.out.println(future.get()); // Blocking đến khi có kết quả
} catch (InterruptedException | ExecutionException e) {
    e.printStackTrace();
}

executor.shutdown();
```
🔹 `Callable` trả về kết quả, sử dụng `Future` để lấy giá trị.

## **5. Quản lý ExecutorService**

Sau khi sử dụng, **luôn gọi `shutdown()`** để giải phóng tài nguyên.
```java
ExecutorService executor = Executors.newFixedThreadPool(2);

executor.submit(() -> System.out.println("Task running"));

executor.shutdown(); // Dừng nhận task mới, chờ task hiện tại hoàn thành

try {
    if (!executor.awaitTermination(5, TimeUnit.SECONDS)) {
        executor.shutdownNow(); // Hủy các task đang chạy
    }
} catch (InterruptedException e) {
    executor.shutdownNow();
}
```

🔹 **`shutdown()`**: Không nhận task mới, chờ task hiện tại hoàn thành.  
🔹 **`shutdownNow()`**: Cố gắng dừng ngay lập tức, nhưng task đang chạy vẫn có thể hoàn thành.  
🔹 **`awaitTermination()`**: Đợi trong thời gian xác định trước khi dừng.


### 6. ScheduledExecutorService – Chạy task theo lịch
```java
import java.util.concurrent.*;

public class ScheduledExecutorExample {
    public static void main(String[] args) {
        ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(2);

        // Chạy task sau 3 giây
        scheduler.schedule(() -> System.out.println("Task executed after delay"), 3, TimeUnit.SECONDS);

        // Chạy task mỗi 2 giây, bắt đầu sau 1 giây
        scheduler.scheduleAtFixedRate(() -> System.out.println("Fixed rate task"), 1, 2, TimeUnit.SECONDS);

        // Chạy task sau mỗi 2 giây kể từ khi task trước kết thúc
        scheduler.scheduleWithFixedDelay(() -> System.out.println("Fixed delay task"), 1, 2, TimeUnit.SECONDS);

        try {
            Thread.sleep(10_000); // Đợi để xem kết quả
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        scheduler.shutdown();
    }
}
```

🔹 **`schedule()`**: Chạy một lần sau một khoảng thời gian.  
🔹 **`scheduleAtFixedRate()`**: Chạy task định kỳ, kể từ khi bắt đầu.  
🔹 **`scheduleWithFixedDelay()`**: Chạy task sau một khoảng thời gian từ khi task trước kết thúc.

## **7. Khi nào sử dụng loại Executor nào?**

|Loại Executor|Khi nào sử dụng?|
|---|---|
|`newFixedThreadPool(n)`|Khi số lượng task ổn định, tránh quá tải CPU.|
|`newCachedThreadPool()`|Khi có nhiều task ngắn, số lượng không cố định.|
|`newSingleThreadExecutor()`|Khi cần chạy tuần tự (ví dụ: ghi log).|
|`newScheduledThreadPool(n)`|Khi cần lên lịch chạy task định kỳ.|
## **8. So sánh ExecutorService vs CompletableFuture**

|Đặc điểm|ExecutorService|CompletableFuture|
|---|---|---|
|Quản lý Thread|✅|❌ (Dùng ForkJoinPool)|
|Blocking|✅|❌ (Non-blocking)|
|Chaining (thenApply, thenCombine)|❌|✅|
|Xử lý ngoại lệ linh hoạt|❌|✅|
|Dễ đọc code hơn|❌|✅|

📌 **Khi nào dùng gì?**

- **ExecutorService**: Khi cần kiểm soát thread, kết hợp với Runnable/Callable.
- **CompletableFuture**: Khi cần async chaining, xử lý ngoại lệ tốt hơn.

## **9. Best Practices khi dùng ExecutorService**

✅ **Sử dụng `newFixedThreadPool(n)` nếu số lượng task ổn định**.  
✅ **Dùng `shutdown()` sau khi hoàn thành công việc**.  
✅ **Dùng `awaitTermination()` để đợi hoàn tất**.  
✅ **Không dùng `newCachedThreadPool()` nếu không kiểm soát được số lượng task**.  
✅ **Dùng `scheduleWithFixedDelay()` thay vì `scheduleAtFixedRate()` để tránh lỗi khi task chạy lâu hơn dự kiến**.  
✅ **Dùng `CompletableFuture` khi có thể để tận dụng async programming**.

---

## **10. Kết luận**

- `ExecutorService` giúp quản lý thread pool, tăng hiệu suất.
- Có nhiều loại `ExecutorService`, chọn loại phù hợp với bài toán.
- **Luôn shutdown() sau khi dùng** để tránh memory leak.
- Nếu cần async chaining, nên dùng **CompletableFuture** thay vì `Future`.