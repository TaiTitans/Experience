
---
**Virtual Thread** (VT) là một khái niệm mới được giới thiệu trong Java 19 và được hỗ trợ bởi Spring Framework. Nó cung cấp một cách thức mới để thực thi các tác vụ song song trong Java, hứa hẹn mang lại nhiều lợi ích so với các phương pháp truyền thống như Thread và ForkJoinPool.

**1. Virtual Thread trong Java:**

- **Khái niệm:** VT là một khái niệm trừu tượng hóa việc thực thi song song, giúp đơn giản hóa việc viết mã đồng thời và giảm thiểu chi phí quản lý luồng.
- **Hoạt động:** VT được triển khai dựa trên nền tảng co-routine, cho phép nhiều tác vụ chia sẻ cùng một luồng JVM. Nhờ vậy, nó có thể giảm thiểu overhead so với việc sử dụng nhiều luồng riêng biệt.
- **Lợi ích:**
    - **Giảm thiểu overhead:** VT sử dụng ít tài nguyên hơn so với luồng truyền thống, giúp cải thiện hiệu suất cho các ứng dụng có nhiều tác vụ I/O chờ đợi.
    - **Đơn giản hóa mã:** VT cung cấp API đơn giản hơn Thread, giúp dễ dàng viết mã đồng thời và giảm nguy cơ xảy ra lỗi.
    - **Tăng khả năng bảo trì:** VT giúp quản lý tài nguyên bộ nhớ hiệu quả hơn, dẫn đến mã dễ bảo trì hơn.

**2. Virtual Thread trong Spring Boot:**

- **Spring Framework cung cấp hỗ trợ cho VT thông qua các thư viện sau:**
    - `spring-async`: Cho phép sử dụng VT để thực thi các tác vụ bất đồng bộ trong các ứng dụng Spring.
    - `webflux`: Hỗ trợ VT trong các ứng dụng web phản hồi.
- **Lợi ích:**
    - **Tự động hóa việc quản lý VT:** Spring Boot tự động hóa việc tạo và quản lý VT, giúp giảm thiểu gánh nặng cho lập trình viên.
    - **Tích hợp với các thành phần Spring khác:** VT được tích hợp tốt với các thành phần Spring khác như Spring MVC và Spring Data, giúp dễ dàng áp dụng vào các ứng dụng Spring hiện có.
- **Ví dụ:** Sử dụng `@Async` annotation để đánh dấu một phương thức là có thể thực thi bởi VT.

## So sánh Virtual Thread với Thread thông thường

**Virtual Thread (VT)** và **Thread** là hai phương pháp chính để thực thi các tác vụ song song trong Java. Mỗi phương pháp đều có những ưu và nhược điểm riêng, phù hợp với những trường hợp sử dụng khác nhau.

**Dưới đây là bảng so sánh chi tiết giữa VT và Thread:**


| Đặc đểm                   | Virtual Thread                                                                                                                 | Thread                                                                                                                         |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| **Hoạt động**             | VT được triển khai dựa trên nền tảng co-routine, cho phép nhiều tác vụ chia sẻ cùng một luồng JVM.                             | Thread được quản lý bởi hệ điều hành và có trạng thái riêng biệt.                                                              |
|  **Overhead**             | VT có overhead thấp hơn do giảm thiểu chi phí chuyển đổi ngữ cảnh và quản lý luồng.                                            | Thread có overhead cao hơn do chi phí quản lý luồng bởi hệ điều hành.                                                          |
|  **Khả năng quản lý**     | VT dễ quản lý hơn do Spring Boot tự động hóa việc tạo và quản lý VT.                                                           | Thread khó quản lý hơn do lập trình viên phải tự tạo và quản lý luồng.                                                         |
|  **Khả năng bảo trì**     | VT giúp quản lý tài nguyên bộ nhớ hiệu quả hơn, dẫn đến mã dễ bảo trì hơn.                                                     | Thread có thể dẫn đến rò rỉ bộ nhớ và các vấn đề đồng bộ hóa nếu không được quản lý cẩn thận.                                  |
|  **Hiệu suất I/O**        | VT phù hợp cho các ứng dụng có nhiều tác vụ I/O chờ đợi do nó có thể xử lý nhiều tác vụ I/O song song mà không cần chặn luồng. | Thread cũng có thể phù hợp cho các ứng dụng có nhiều tác vụ I/O chờ đợi, nhưng overhead của nó có thể ảnh hưởng đến hiệu suất. |
|  **Hiệu suất tính toán**  | VT có thể không phù hợp cho các ứng dụng có nhiều tác vụ tính toán nặng do nó chia sẻ tài nguyên CPU với các tác vụ khác.      | Thread phù hợp cho các ứng dụng có nhiều tác vụ tính toán nặng do mỗi thread có thể sử dụng một lõi CPU riêng biệt.            |
|  **Khả năng mở rộng**     | VT có thể mở rộng tốt hơn do chi phí tạo và quản lý VT thấp hơn.                                                               | Thread có thể khó mở rộng do chi phí quản lý luồng bởi hệ điều hành tăng lên khi số lượng luồng tăng.                          |
## So sánh Virtual Thread với Thread thông thường

**Virtual Thread (VT)** và **Thread** là hai phương pháp chính để thực thi các tác vụ song song trong Java. Mỗi phương pháp đều có những ưu và nhược điểm riêng, phù hợp với những trường hợp sử dụng khác nhau.

**Dưới đây là bảng so sánh chi tiết giữa VT và Thread:**

|Đặc điểm|Virtual Thread|Thread thông thường|
|---|---|---|
|**Hoạt động**|VT được triển khai dựa trên nền tảng co-routine, cho phép nhiều tác vụ chia sẻ cùng một luồng JVM.|Thread được quản lý bởi hệ điều hành và có trạng thái riêng biệt.|
|**Overhead**|VT có overhead thấp hơn do giảm thiểu chi phí chuyển đổi ngữ cảnh và quản lý luồng.|Thread có overhead cao hơn do chi phí quản lý luồng bởi hệ điều hành.|
|**Khả năng quản lý**|VT dễ quản lý hơn do Spring Boot tự động hóa việc tạo và quản lý VT.|Thread khó quản lý hơn do lập trình viên phải tự tạo và quản lý luồng.|
|**Khả năng bảo trì**|VT giúp quản lý tài nguyên bộ nhớ hiệu quả hơn, dẫn đến mã dễ bảo trì hơn.|Thread có thể dẫn đến rò rỉ bộ nhớ và các vấn đề đồng bộ hóa nếu không được quản lý cẩn thận.|
|**Hiệu suất I/O**|VT phù hợp cho các ứng dụng có nhiều tác vụ I/O chờ đợi do nó có thể xử lý nhiều tác vụ I/O song song mà không cần chặn luồng.|Thread cũng có thể phù hợp cho các ứng dụng có nhiều tác vụ I/O chờ đợi, nhưng overhead của nó có thể ảnh hưởng đến hiệu suất.|
|**Hiệu suất tính toán**|VT có thể không phù hợp cho các ứng dụng có nhiều tác vụ tính toán nặng do nó chia sẻ tài nguyên CPU với các tác vụ khác.|Thread phù hợp cho các ứng dụng có nhiều tác vụ tính toán nặng do mỗi thread có thể sử dụng một lõi CPU riêng biệt.|
|**Khả năng mở rộng**|VT có thể mở rộng tốt hơn do chi phí tạo và quản lý VT thấp hơn.|Thread có thể khó mở rộng do chi phí quản lý luồng bởi hệ điều hành tăng lên khi số lượng luồng tăng.|


**Kết luận:**

- Nên sử dụng VT cho các ứng dụng có nhiều tác vụ I/O chờ đợi, yêu cầu hiệu suất cao và dễ dàng quản lý.
- Nên sử dụng Thread cho các ứng dụng có nhiều tác vụ tính toán nặng, yêu cầu hiệu suất cao và cần kiểm soát chính xác việc sử dụng CPU.

---

## Overhead trong Virtual Thread Java

**Overhead** trong Virtual Thread (VT) Java là **chi phí chung** liên quan đến việc sử dụng VT để thực thi các tác vụ song song.

[Xem thêm](https://viblo.asia/p/virtual-threads-nen-tang-moi-cho-ung-dung-java-quy-mo-lon-5pPLkxA8VRZ)