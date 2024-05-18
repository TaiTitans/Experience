
---

Concurrency trong Java là một khái niệm quan trọng cho phép thực thi đồng thời nhiều tác vụ trong một chương trình. Nó giúp cải thiện hiệu suất và khả năng đáp ứng của ứng dụng bằng cách tận dụng tối đa tài nguyên hệ thống, đặc biệt là các CPU đa lõi.

	Concurrency là khả năng chạy nhiều phần của một chương trình đồng thời, có thể là trên cùng một CPU thông qua switching context hoặc trên nhiều CPU trong một hệ thống đa lõi.

### Các khái niệm cơ bản trong Concurrency

- **Thread**: Một đơn vị thực thi nhỏ nhất của một chương trình. Mỗi thread trong Java là một thể hiện của lớp `Thread` hoặc một đối tượng của một lớp thực hiện giao diện `Runnable`.
- **Synchronization**: Đảm bảo rằng hai hoặc nhiều thread không truy cập đồng thời vào dữ liệu chia sẻ để tránh các vấn đề về tính toàn vẹn của dữ liệu.
- **Deadlock**: Xảy ra khi hai hoặc nhiều thread chờ đợi lẫn nhau để giải phóng tài nguyên, dẫn đến tình trạng không thread nào có thể tiếp tục hoạt động.
- **Race Condition**: Xảy ra khi hai hoặc nhiều thread truy cập đồng thời vào dữ liệu chia sẻ và thứ tự thực hiện ảnh hưởng đến kết quả cuối cùng.

---
Có hai cách để tạo và chạy một thread trong Java: kế thừa từ lớp `Thread` hoặc thực hiện giao diện `Runnable`.

#### Cách 1: Kế thừa từ lớp `Thread`

```Java
class MyThread extends Thread {
    public void run() {
        System.out.println("Thread đang chạy...");
    }
    
    public static void main(String[] args) {
        MyThread thread = new MyThread();
        thread.start();  // Bắt đầu thread
    }
}

```


#### Cách 2: Thực hiện giao diện `Runnable`

```Java
class MyRunnable implements Runnable {
    public void run() {
        System.out.println("Thread đang chạy...");
    }
    
    public static void main(String[] args) {
        Thread thread = new Thread(new MyRunnable());
        thread.start();  // Bắt đầu thread
    }
}

```


### 4. Synchronization

Synchronization được sử dụng để đảm bảo rằng chỉ một thread có thể truy cập vào một phương thức hoặc một khối mã tại một thời điểm.

#### Synchronized Method

```Java
class Counter {
    private int count = 0;
    
    public synchronized void increment() {
        count++;
    }
    
    public int getCount() {
        return count;
    }
}

class MyRunnable implements Runnable {
    private Counter counter;
    
    public MyRunnable(Counter counter) {
        this.counter = counter;
    }
    
    public void run() {
        for(int i = 0; i < 1000; i++) {
            counter.increment();
        }
    }
    
    public static void main(String[] args) throws InterruptedException {
        Counter counter = new Counter();
        Thread t1 = new Thread(new MyRunnable(counter));
        Thread t2 = new Thread(new MyRunnable(counter));
        
        t1.start();
        t2.start();
        
        t1.join();
        t2.join();
        
        System.out.println("Count: " + counter.getCount());  // Nên in ra 2000
    }
}

```

### 5. Advanced Concurrency

Java cung cấp một gói thư viện mạnh mẽ là `java.util.concurrent` để hỗ trợ lập trình concurrency cao cấp hơn. Một số lớp hữu ích bao gồm:

- **ExecutorService**: Quản lý một nhóm các thread làm việc với nhau.
- **ConcurrentHashMap**: Một phiên bản thread-safe của `HashMap`.
- **CountDownLatch**: Cho phép một hoặc nhiều thread chờ cho đến khi một số hoạt động hoàn thành.

```Java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

class MyRunnable implements Runnable {
    private String message;
    
    public MyRunnable(String message) {
        this.message = message;
    }
    
    public void run() {
        System.out.println(Thread.currentThread().getName() + " (Start) message = " + message);
        processMessage();
        System.out.println(Thread.currentThread().getName() + " (End)");
    }
    
    private void processMessage() {
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
    
    public static void main(String[] args) {
        ExecutorService executor = Executors.newFixedThreadPool(5);
        for (int i = 0; i < 10; i++) {
            Runnable worker = new MyRunnable("" + i);
            executor.execute(worker);
        }
        executor.shutdown();
        while (!executor.isTerminated()) {
        }
        System.out.println("Finished all threads");
    }
}

```

### Kết luận

Concurrency là một kỹ năng quan trọng trong lập trình Java, đặc biệt là trong các ứng dụng yêu cầu hiệu suất cao và khả năng phản hồi tốt. Bằng cách hiểu và sử dụng các công cụ và kỹ thuật mà Java cung cấp, bạn có thể tạo ra các ứng dụng mạnh mẽ và hiệu quả.