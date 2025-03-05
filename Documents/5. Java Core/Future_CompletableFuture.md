
---
### **Future và CompletableFuture trong Java như một Senior**

Trong Java, **Future** và **CompletableFuture** được sử dụng để xử lý lập trình bất đồng bộ (asynchronous programming), giúp cải thiện hiệu suất bằng cách tránh chờ đợi các tác vụ chạy lâu. Tuy nhiên, **CompletableFuture** là một sự nâng cấp đáng kể so với **Future**. Hãy đi sâu vào từng khái niệm.

## **1. Future trong Java**

### **1.1. Giới thiệu về Future**

`Future<T>` là một interface trong Java (`java.util.concurrent`) được sử dụng để đại diện cho kết quả của một tác vụ bất đồng bộ.

- Được giới thiệu từ Java 5 cùng với `ExecutorService`.
- Cho phép chạy một tác vụ trong nền và truy xuất kết quả sau đó.
- Tuy nhiên, `Future` có nhiều hạn chế:
    - Không thể tự động nhận thông báo khi tác vụ hoàn thành (phải gọi `isDone()` để kiểm tra).
    - Không thể thực hiện chaining các tác vụ (không có `thenApply`, `thenAccept`, v.v.).
    - Không thể xử lý ngoại lệ một cách linh hoạt.
    - Không có phương thức callback khi tác vụ hoàn tất.

```java
import java.util.concurrent.*;

public class FutureExample {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        ExecutorService executor = Executors.newFixedThreadPool(2);
        
        Future<String> future = executor.submit(() -> {
            Thread.sleep(2000);  // Giả lập tác vụ mất thời gian
            return "Hello from Future!";
        });

        // Kiểm tra xem task đã hoàn thành chưa
        while (!future.isDone()) {
            System.out.println("Waiting for result...");
            Thread.sleep(500);
        }

        // Lấy kết quả từ Future (block nếu chưa có kết quả)
        String result = future.get();
        System.out.println("Result: " + result);
        
        executor.shutdown();
    }
}
```

🔹 **Hạn chế:**

- `future.get()` là một cuộc gọi **blocking**, làm chậm luồng chính nếu task chưa hoàn thành.
- Không thể thực hiện chaining hoặc xử lý ngoại lệ một cách dễ dàng.

---
## **2. CompletableFuture - Giải pháp nâng cấp Future**

### **2.1. Giới thiệu CompletableFuture**

Được giới thiệu từ **Java 8**, `CompletableFuture<T>` là một sự mở rộng của `Future` giúp viết code async dễ dàng hơn. Các tính năng nổi bật: ✔ **Hỗ trợ callback** (`thenApply`, `thenAccept`, `whenComplete`...)  
✔ **Xử lý ngoại lệ tốt hơn** (`exceptionally`, `handle`)  
✔ **Chaining nhiều bước** (kết hợp nhiều tác vụ)  
✔ **Hỗ trợ combine nhiều Future với nhau**

2.2. Cách sử dụng CompletableFuture cơ bản
```java
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;

public class CompletableFutureExample {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                throw new IllegalStateException(e);
            }
            return "Hello from CompletableFuture!";
        });

        // Không block, code vẫn tiếp tục chạy
        future.thenAccept(result -> System.out.println("Result: " + result));

        System.out.println("Main thread is not blocked!");
        
        Thread.sleep(3000);  // Chờ đủ lâu để thấy kết quả
    }
}
```

🔹 **Ưu điểm:**

- `thenAccept(result -> ...)` giúp xử lý kết quả ngay khi hoàn tất.
- Không block `main thread`, giúp tăng hiệu suất.

### **2.3. Chaining với thenApply**

Khi cần xử lý dữ liệu sau khi nhận kết quả từ Future:
```java
CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> "Hello")
    .thenApply(result -> result + " World!")
    .thenApply(String::toUpperCase);

System.out.println(future.get()); // HELLO WORLD!
```

🔹 **thenApply()** cho phép chuyển đổi dữ liệu qua từng bước.

2.4. Xử lý ngoại lệ với exceptionally

```java
CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
    if (true) throw new RuntimeException("Something went wrong!");
    return "Success!";
}).exceptionally(ex -> "Recovered from error: " + ex.getMessage());

System.out.println(future.get()); // Recovered from error: Something went wrong!
```

🔹 **exceptionally()** giúp xử lý lỗi và trả về giá trị thay thế.

### **2.5. Kết hợp nhiều CompletableFuture**

#### **2.5.1. Chạy song song với thenCombine**
```java
CompletableFuture<String> future1 = CompletableFuture.supplyAsync(() -> "Hello");
CompletableFuture<String> future2 = CompletableFuture.supplyAsync(() -> " World");

CompletableFuture<String> combinedFuture = future1.thenCombine(future2, (res1, res2) -> res1 + res2);
System.out.println(combinedFuture.get()); // Hello World
```

🔹 **thenCombine()** giúp kết hợp kết quả từ hai Futures.

2.5.2. Chạy song song với allOf
```java
CompletableFuture<String> future1 = CompletableFuture.supplyAsync(() -> "Task 1");
CompletableFuture<String> future2 = CompletableFuture.supplyAsync(() -> "Task 2");
CompletableFuture<String> future3 = CompletableFuture.supplyAsync(() -> "Task 3");

CompletableFuture<Void> allFutures = CompletableFuture.allOf(future1, future2, future3);

allFutures.thenRun(() -> System.out.println("All tasks completed!"));
allFutures.get();
```

🔹 **allOf()** đợi tất cả Futures hoàn thành mà không quan tâm kết quả.

## **3. Khi nào dùng Future và CompletableFuture?**

|Đặc điểm|Future|CompletableFuture|
|---|---|---|
|Dễ sử dụng|✅|✅✅|
|Blocking|❌ (blocking)|✅ (non-blocking)|
|Callback support|❌|✅|
|Chaining operations|❌|✅|
|Xử lý ngoại lệ|❌|✅|
|Combine nhiều tác vụ|❌|✅|

📌 **Tóm lại:**

- **Future** chỉ hữu ích khi chạy các tác vụ đơn giản.
- **CompletableFuture** là lựa chọn tối ưu cho lập trình bất đồng bộ hiện đại.

## **4. Best Practices khi sử dụng CompletableFuture**

✅ **Luôn sử dụng thread pool hợp lý (ForkJoinPool, Executors.newFixedThreadPool).**  
✅ **Tránh dùng .get() nếu không cần thiết để tránh blocking.**  
✅ **Luôn xử lý ngoại lệ bằng exceptionally() hoặc handle().**  
✅ **Dùng thenCombine hoặc allOf để chạy song song các task.**

Ví dụ sử dụng `ExecutorService`:
```java
ExecutorService executor = Executors.newFixedThreadPool(10);
CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> "Hello", executor);
executor.shutdown();
```
## **Kết luận**

- **Future** đã lỗi thời và không linh hoạt.
- **CompletableFuture** cung cấp các API mạnh mẽ cho lập trình bất đồng bộ.
- Sử dụng **thenApply, thenAccept, exceptionally, allOf** giúp tối ưu hiệu suất.
- **Không nên sử dụng get() mà hãy dùng callback**.


