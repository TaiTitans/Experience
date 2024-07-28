
---
### Java Learning Path

#### Basic Knowledge

- **Basic Concepts**
    - Class and Object
    - Inheritance, Polymorphism, Interfaces
    - Encapsulation
    - Exception Handling
    - Collections
    - Generics
    - Multithreading
    - IO
    - Object-Oriented Programming (OOP)
- **Basic Syntax**
    - Variables
    - Data Types
    - Strings
    - Keywords
    - Loops
    - Arrays
    - Conditional Statements
    - Serialization
- **Tools**
    - Git
    - Maven/Gradle
    - IntelliJ IDEA

#### Advanced Knowledge

- **Design Patterns**
    - Factory Pattern
    - Singleton Pattern
    - Proxy Pattern
    - Decorator Pattern
    - Composite Pattern
    - Strategy Pattern
    - Observer Pattern
    - Template Pattern
    - Chain of Responsibility Pattern
    - State Pattern
- **Frameworks**
    - Spring
    - SpringBoot
    - MyBatis
    - SpringMVC
    - SpringCloud
    - Dubbo
    - Netty
- **Performance Optimization**
    - JVM
    - Memory Model
    - Garbage Collection Mechanisms
    - High-Concurrency and High-Performance

#### Database (MySQL)

- **Fundamentals**
    - Table Creation
    - Indexing
    - Stored Procedures and Triggers
    - Data Types
    - Transactions
    - Joins
    - Partitioning
- **Advanced**
    - Query Optimization
    - Index Optimization
    - Database Normalization
    - Backup and Recovery
    - Cluster and Replication
    - Performance Tuning

#### Middleware

- **Message Queue (MQ)**
    - Kafka
    - RabbitMQ
    - ActiveMQ
    - ZeroMQ
    - RocketMQ
- **Cache**
    - Redis
        - Basic Usage
        - Expiration Policies
        - Persistence
        - Cache Penetration
        - Cache Avalanche
        - Cache Warm-Up
        - Distributed Lock
        - Data Sharding
        - High Availability

#### Computing Knowledge

- **Networking**
    - TCP/IP
    - HTTP/HTTPS
    - WebSocket
    - RESTful APIs
- **Operating Systems**
    - Linux
    - Shell Scripting
- **Algorithms and Data Structures**
    - Sorting
    - Searching
    - Trees, Graphs, Heaps
    - Dynamic Programming
- **Data Communication**
    - Protocols
    - Error Detection and Correction
    - Networking Models

KeyCloak
KISS
Clean architecture
Học cách xử lý logging trong hệ thống
Naming convention
Cqrs
DDD


**Khi gặp await thì luồng xử lý sao** ?

> [!Result] Result
Khi gặp từ khóa `await` trong một hàm asynchronous (hàm có từ khóa `async`), luồng xử lý sẽ dừng lại tại vị trí đó cho đến khi promise mà `await` đang chờ được giải quyết (resolved hoặc rejected). Dưới đây là chi tiết về cách `await` hoạt động:
	1. `Dừng lại`: Khi gặp `await`, luồng xử lý sẽ dừng lại tại điểm này và đợi promise được giải quyết.
	2. **Tiếp tục khi promise hoàn thành**: Khi promise được giải quyết, luồng xử lý sẽ tiếp tục với giá trị được trả về từ promise nếu nó được giải quyết thành công, hoặc bắt lỗi nếu promise bị từ chối (rejected).
	3. **Không chặn luồng xử lý**: Mặc dù luồng xử lý trong hàm async dừng lại tại `await`, nó không chặn các luồng xử lý khác trong chương trình. Các tác vụ khác có thể tiếp tục chạy.



**Hàm async ko dùng await thì sao ?**


> [!NOTE] Result
> Nếu một hàm async không sử dụng `await`, nó vẫn là một hàm bất đồng bộ và sẽ trả về một `Promise`. Điều này có nghĩa là hàm đó sẽ được thực thi đồng bộ đến khi kết thúc và trả về một `Promise` mà không chờ đợi bất kỳ tác vụ bất đồng bộ nào.


**Thread ở đâu sinh ra để xử lý**


> [!NOTE] Result
> Trong lập trình, luồng (thread) là một đơn vị nhỏ nhất của quá trình xử lý mà hệ điều hành có thể quản lý một cách độc lập. Luồng được sinh ra để thực hiện các tác vụ đồng thời, giúp tăng hiệu suất và tận dụng tài nguyên hệ thống hiệu quả hơn. Dưới đây là một số ngữ cảnh cụ thể về nơi và cách thức luồng được sinh ra để xử lý:

### 1. **Trong Java:**

Java cung cấp hỗ trợ mạnh mẽ cho đa luồng thông qua lớp `Thread` và giao diện `Runnable`.

- **Tạo luồng bằng cách mở rộng lớp `Thread`:**
```Java
class MyThread extends Thread {
    public void run() {
        System.out.println("Thread is running.");
    }
}

public class Main {
    public static void main(String[] args) {
        MyThread t1 = new MyThread();
        t1.start();  // Bắt đầu luồng mới
    }
}

```
- **Tạo luồng bằng cách triển khai giao diện `Runnable`:**
```Java
class MyRunnable implements Runnable {
    public void run() {
        System.out.println("Thread is running.");
    }
}

public class Main {
    public static void main(String[] args) {
        Thread t1 = new Thread(new MyRunnable());
        t1.start();  // Bắt đầu luồng mới
    }
}

```
**Máy tính cấp gì cho vụ xử lý thread này**
Khi máy tính xử lý các luồng (threads), nó cần cấp phát một số tài nguyên và cơ chế quản lý để đảm bảo các luồng có thể hoạt động một cách hiệu quả và không xung đột với nhau. Dưới đây là các tài nguyên và cơ chế chính mà máy tính cấp phát và quản lý khi xử lý luồng:

### 1. **CPU Time (Thời gian CPU)**

- **Chia sẻ CPU**: Hệ điều hành sử dụng một bộ lập lịch (scheduler) để chia sẻ thời gian CPU giữa các luồng. Mỗi luồng sẽ nhận được một lượng thời gian CPU nhất định để thực thi các tác vụ của mình.
- **Chuyển ngữ cảnh (Context Switching)**: Khi hệ điều hành chuyển đổi từ luồng này sang luồng khác, nó phải lưu lại trạng thái hiện tại của luồng đang chạy (như các thanh ghi CPU, con trỏ ngăn xếp) và khôi phục trạng thái của luồng tiếp theo. Quá trình này gọi là chuyển ngữ cảnh và có thể tốn tài nguyên.

### 2. **Bộ nhớ (Memory)**

- **Ngăn xếp (Stack)**: Mỗi luồng có ngăn xếp riêng để lưu trữ các biến cục bộ và địa chỉ trả về khi thực hiện các lời gọi hàm. Ngăn xếp này được cấp phát khi luồng được tạo ra.
- **Không gian địa chỉ chung (Shared Address Space)**: Các luồng trong cùng một quá trình chia sẻ không gian địa chỉ của quá trình đó, bao gồm vùng nhớ dữ liệu, heap và các tài nguyên khác.

### 3. **Bộ nhớ Cache**

- **Cache Locality**: Khi nhiều luồng truy cập vào cùng một vùng nhớ, hệ điều hành phải quản lý các bản sao bộ nhớ cache để đảm bảo tính nhất quán và hiệu suất. Điều này có thể bao gồm việc sử dụng các kỹ thuật như MESI protocol để quản lý cache coherency.

### 4. **I/O Resources (Tài nguyên I/O)**

- **Quản lý I/O**: Hệ điều hành quản lý truy cập đến các thiết bị I/O (như đĩa cứng, mạng) để đảm bảo rằng các luồng có thể truy cập và sử dụng các tài nguyên này một cách hiệu quả và không xung đột.

### 5. **Synchronization Primitives (Các nguyên thủy đồng bộ)**

- **Mutexes và Semaphores**: Để tránh các vấn đề về đồng bộ hóa và tranh chấp tài nguyên, hệ điều hành cung cấp các cơ chế như mutexes và semaphores để các luồng có thể đồng bộ hóa truy cập đến các tài nguyên chung.
- **Condition Variables**: Được sử dụng để cho phép các luồng chờ đợi một điều kiện nhất định trước khi tiếp tục thực thi.

### 6. **Quản lý Luồng (Thread Management)**

- **Thread Control Block (TCB)**: Mỗi luồng có một khối điều khiển luồng (TCB) chứa thông tin quản lý về luồng đó, bao gồm trạng thái luồng, con trỏ ngăn xếp, và các thông tin khác cần thiết cho việc quản lý luồng.
- **Thread Creation and Destruction**: Hệ điều hành cung cấp các hàm và cơ chế để tạo và hủy các luồng. Quá trình này bao gồm việc cấp phát và giải phóng các tài nguyên liên quan đến luồng.

### 7. **Bảo mật và Quyền truy cập (Security and Access Rights)**

- **Access Control**: Hệ điều hành quản lý quyền truy cập của các luồng đến các tài nguyên hệ thống để đảm bảo rằng chỉ những luồng được cấp phép mới có thể truy cập và sử dụng các tài nguyên này.
- **Isolation**: Trong các hệ thống yêu cầu bảo mật cao, các luồng có thể được cô lập để đảm bảo rằng một luồng không thể ảnh hưởng đến hoặc truy cập dữ liệu của luồng khác một cách trái phép.

### Tóm lại

Hệ điều hành cung cấp và quản lý nhiều tài nguyên và cơ chế để xử lý luồng, bao gồm thời gian CPU, bộ nhớ, bộ nhớ cache, tài nguyên I/O, các nguyên thủy đồng bộ, quản lý luồng, và bảo mật. Điều này đảm bảo rằng các luồng có thể hoạt động một cách hiệu quả và an toàn, giúp tăng cường hiệu suất và khả năng phản hồi của các ứng dụng đa luồng.

**black list và white list**
**race condition**
- **Race Condition**: Xảy ra khi hai hoặc nhiều luồng truy cập và thay đổi dữ liệu dùng chung mà không có sự kiểm soát, dẫn đến kết quả không mong muốn.
- **Cách Giải Quyết**: Sử dụng các cơ chế đồng bộ hóa như `synchronized`, `Lock`, và các biến nguyên tử để đảm bảo tính toàn vẹn dữ liệu và tránh race condition.

----
**Idempotent** (tính chất bất biến) là một thuật ngữ trong khoa học máy tính và toán học, đặc biệt quan trọng trong thiết kế hệ thống và API. Một thao tác hoặc hàm được gọi là idempotent nếu thực hiện nó nhiều lần cũng cho kết quả giống như thực hiện nó một lần.

### Ví dụ về Idempotent

1. **Toán học**:
    
    - Phép toán `max(x, x)` luôn trả về `x`, bất kể bạn thực hiện bao nhiêu lần.
    - Hàm `abs(x)` (hàm giá trị tuyệt đối) cũng là idempotent vì `abs(abs(x))` luôn bằng `abs(x)`.
2. **HTTP Methods**:
    
    - **GET**: Lấy dữ liệu từ server. Gọi nhiều lần không thay đổi trạng thái server và luôn trả về cùng một dữ liệu.
    - **PUT**: Thay thế hoặc cập nhật tài nguyên trên server. Nếu bạn gửi cùng một yêu cầu PUT nhiều lần, trạng thái của tài nguyên sẽ không thay đổi sau lần thực hiện đầu tiên.
    - **DELETE**: Xóa tài nguyên trên server. Gọi nhiều lần sẽ không gây ra thay đổi sau lần xóa đầu tiên (tài nguyên sẽ bị xóa hoặc không tồn tại).

### Tại sao Idempotent quan trọng?

- **Đảm bảo tính ổn định và đáng tin cậy**: Trong các hệ thống phân tán, yêu cầu có thể bị gửi lại nhiều lần do lỗi mạng hoặc các vấn đề khác. Nếu các thao tác là idempotent, hệ thống sẽ không bị ảnh hưởng bởi các yêu cầu trùng lặp này.
- **Thiết kế API tốt hơn**: Các API idempotent dễ sử dụng và bảo trì hơn vì người dùng API không cần lo lắng về việc gửi lại yêu cầu có thể gây ra thay đổi không mong muốn.
