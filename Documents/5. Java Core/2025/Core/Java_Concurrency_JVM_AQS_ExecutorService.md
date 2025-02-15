
---
### **📌 Phần 1: Java Concurrency và JVM**

#### 1️⃣ **JVM và Concurrency**

Khi chạy các ứng dụng Java đa luồng, JVM quản lý các **Thread**, **Heap Memory**, và **Garbage Collection (GC)**. Những điểm quan trọng:

- **Thread States**: NEW, RUNNABLE, BLOCKED, WAITING, TIMED_WAITING, TERMINATED.
- **Heap & Stack Memory**: Stack dành cho từng Thread, Heap dùng chung giữa các Thread.
- **Synchronization**: Monitor Lock, Synchronized Blocks, Volatile, Atomic Variables.
- **Garbage Collection (GC)** và ảnh hưởng đến Performance.
### **📌 Phần 2: AbstractQueuedSynchronizer (AQS)**

#### 2️⃣ **AQS là gì?**

AQS là nền tảng của nhiều cơ chế đồng bộ trong Java như **ReentrantLock**, **Semaphore**, **CountDownLatch**, **ReadWriteLock**.

- **Cấu trúc Queue-based FIFO**: Sử dụng **CLH Queue (Craig, Landin, Hagersten)**.
- **State Management**: `acquire()`, `release()`, `tryAcquire()`, `tryRelease()`.
- **Exclusive vs Shared Mode**: Dùng cho các lock như **ReentrantLock** (exclusive) hay **Semaphore** (shared).

### **📌 Phần 3: ExecutorService**

#### 3️⃣ **Thread Pool và Executor Framework**

Java cung cấp `ExecutorService` để quản lý **Thread Pool**, giúp tối ưu hiệu suất:

- **FixedThreadPool**: Giới hạn số lượng thread chạy đồng thời.
- **CachedThreadPool**: Tạo thread mới khi cần, reuse thread cũ nếu có sẵn.
- **ScheduledThreadPool**: Lên lịch chạy task theo thời gian.
- **ForkJoinPool**: Chia nhỏ công việc theo mô hình **Work Stealing**.

---
