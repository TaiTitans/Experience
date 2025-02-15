
---
Trong lập trình đa luồng (Java Concurrency), việc tạo và quản lý **threads** một cách hiệu quả là rất quan trọng. Ở cấp độ **Senior**, bạn cần hiểu cách **Thread Pool**, **ExecutorService**, và **CompletableFuture** hoạt động để tối ưu hiệu suất ứng dụng.

# **📌 Phần 1: Thread Pool & ExecutorService**

## **1️⃣ Thread Pool là gì?**

Thread Pool là một nhóm các **worker threads** được quản lý để thực thi các tác vụ một cách hiệu quả, thay vì tạo một thread mới mỗi lần.

📌 **Lợi ích của Thread Pool:**

- **Giảm overhead** của việc tạo thread mới liên tục.
- **Tái sử dụng threads** giúp tăng hiệu suất.
- **Kiểm soát số lượng thread hoạt động** tránh quá tải CPU.

## **2️⃣ ExecutorService – Cách tạo Thread Pool**

`ExecutorService` là API chính trong Java để quản lý **Thread Pool**. Nó cung cấp nhiều loại Thread Pool khác nhau:

|Loại Thread Pool|Đặc điểm|
|---|---|
|`newFixedThreadPool(n)`|Pool có `n` threads cố định.|
|`newCachedThreadPool()`|Pool có số threads linh hoạt, tạo thêm nếu cần.|
|`newSingleThreadExecutor()`|Chỉ có một thread, chạy tuần tự.|
|`newScheduledThreadPool(n)`|Pool có thể lên lịch chạy task sau một khoảng thời gian.|
## **3️⃣ Ví dụ về Thread Pool**

📌 **Tạo một Fixed Thread Pool với 3 threads:**

```java
import java.util.concurrent.*;

public class ThreadPoolExample {
    public static void main(String[] args) {
        ExecutorService executor = Executors.newFixedThreadPool(3);

        for (int i = 1; i <= 5; i++) {
            int taskId = i;
            executor.submit(() -> {
                System.out.println("Task " + taskId + " is running on " + Thread.currentThread().getName());
                try {
                    Thread.sleep(2000); // Giả lập công việc
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            });
        }

        executor.shutdown(); // Không nhận thêm task mới, nhưng đợi các task hiện tại hoàn thành
    }
}
```
👉 **Giải thích:**

- Pool có **3 threads** nhưng có **5 tasks**, nên 3 task đầu tiên sẽ chạy, 2 task còn lại phải **đợi** đến khi có thread trống.
- **Không cần tự tạo và quản lý threads**, chỉ cần submit task vào ExecutorService.
## **4️⃣ So sánh Fixed Thread Pool & Cached Thread Pool**

📌 **Ví dụ: Cached Thread Pool tạo nhiều threads hơn nếu cần**
```java
ExecutorService executor = Executors.newCachedThreadPool();

for (int i = 1; i <= 10; i++) {
    int taskId = i;
    executor.submit(() -> {
        System.out.println("Task " + taskId + " running on " + Thread.currentThread().getName());
    });
}

executor.shutdown();
```
👉 Nếu có sẵn thread, nó sẽ tái sử dụng; nếu không, nó sẽ **tạo thread mới**, có thể gây tốn tài nguyên nếu không kiểm soát tốt.

# **📌 Phần 2: CompletableFuture – Lập trình Bất Đồng Bộ**

## **1️⃣ CompletableFuture là gì?**

- Là API trong Java 8 giúp lập trình bất đồng bộ dễ dàng hơn.
- Cho phép **chaining** các bước xử lý mà không cần callback phức tạp.
- Có thể kết hợp nhiều tasks với nhau (combine, compose).

---

## **2️⃣ Tạo CompletableFuture cơ bản**

📌 **Chạy một task bất đồng bộ:**
```java
import java.util.concurrent.CompletableFuture;

public class CompletableFutureExample {
    public static void main(String[] args) {
        CompletableFuture<Void> future = CompletableFuture.runAsync(() -> {
            System.out.println("Running in: " + Thread.currentThread().getName());
        });

        future.join(); // Đợi task hoàn thành
    }
}
```
👉 **`runAsync`** chạy trên một thread từ **ForkJoinPool** (thread pool mặc định của Java).


## **3️⃣ Chaining Tasks (thenApply, thenAccept)**

📌 **Ví dụ về chaining các bước xử lý:**
```java
import java.util.concurrent.CompletableFuture;

public class CompletableFutureChain {
    public static void main(String[] args) {
        CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
            System.out.println("Fetching data...");
            return "Data from API";
        }).thenApply(data -> {
            return data.toUpperCase();
        }).thenAccept(result -> {
            System.out.println("Processed Result: " + result);
        });

        future.join();
    }
}
```
👉 **Giải thích:**

1. `supplyAsync` chạy một tác vụ trả về **String**.
2. `thenApply` chuyển đổi dữ liệu (**chaining**).
3. `thenAccept` nhận kết quả cuối cùng và in ra.

## **4️⃣ Kết hợp nhiều CompletableFuture**

📌 **Ví dụ: Kết hợp hai tasks chạy song song**
```java
import java.util.concurrent.CompletableFuture;

public class CombineCompletableFuture {
    public static void main(String[] args) {
        CompletableFuture<String> task1 = CompletableFuture.supplyAsync(() -> {
            return "Hello";
        });

        CompletableFuture<String> task2 = CompletableFuture.supplyAsync(() -> {
            return "World";
        });

        CompletableFuture<String> combined = task1.thenCombine(task2, (res1, res2) -> res1 + " " + res2);

        System.out.println(combined.join()); // Kết quả: "Hello World"
    }
}
```
👉 **`thenCombine`** giúp kết hợp kết quả của hai `CompletableFuture`.

## **5️⃣ Handling Exceptions**

📌 **Ví dụ: Bắt lỗi với `exceptionally`**
```java
import java.util.concurrent.CompletableFuture;

public class CompletableFutureErrorHandling {
    public static void main(String[] args) {
        CompletableFuture<Integer> future = CompletableFuture.supplyAsync(() -> {
            if (true) throw new RuntimeException("Something went wrong!");
            return 10;
        }).exceptionally(ex -> {
            System.out.println("Caught Exception: " + ex.getMessage());
            return 0; // Trả về giá trị mặc định khi lỗi
        });

        System.out.println("Result: " + future.join());
    }
}
```
👉 **`exceptionally`** giúp xử lý lỗi và trả về giá trị mặc định khi có lỗi.
# **🚀 Tổng kết**

### ✅ **Bạn đã học được gì?**

🔹 **Thread Pool giúp quản lý threads hiệu quả** bằng `ExecutorService`.  
🔹 **Fixed Thread Pool vs Cached Thread Pool** có cách hoạt động khác nhau.  
🔹 **CompletableFuture giúp lập trình bất đồng bộ đơn giản hơn** so với `Future`.  
🔹 **Có thể chaining, kết hợp nhiều tasks, và xử lý lỗi dễ dàng** với `CompletableFuture`.