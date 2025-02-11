
---
Java SE 21 cung cấp một bộ API mạnh mẽ và linh hoạt cho lập trình đồng thời, giúp các nhà phát triển xây dựng các ứng dụng hiệu suất cao và an toàn trong môi trường đa luồng. Các API này được thiết kế để đơn giản hóa việc quản lý luồng, đồng bộ hóa và thực thi tác vụ, đồng thời cung cấp các công cụ cần thiết cho lập trình đồng thời tiên tiến.

**Các thành phần chính của API đồng thời trong Java SE 21:**

1. **Virtual Threads (Luồng ảo):** Luồng ảo là các luồng nhẹ, giúp giảm thiểu công sức khi viết, duy trì và gỡ lỗi các ứng dụng đồng thời có thông lượng cao. Chúng cho phép tạo ra hàng triệu luồng với chi phí tài nguyên thấp, cải thiện hiệu suất và khả năng mở rộng của ứng dụng.
    
2. **Structured Concurrency (Đồng thời có cấu trúc):** Đồng thời có cấu trúc coi các nhóm tác vụ liên quan chạy trên các luồng khác nhau như một đơn vị công việc duy nhất, giúp đơn giản hóa việc xử lý lỗi và hủy bỏ, cải thiện độ tin cậy và tăng cường khả năng quan sát. Lớp `StructuredTaskScope` hỗ trợ các trường hợp trong đó một tác vụ được chia thành nhiều tác vụ con đồng thời và các tác vụ con phải hoàn thành trước khi tác vụ chính tiếp tục. Các chính sách tắt ngắn mạch như `ShutdownOnFailure` và `ShutdownOnSuccess` giúp quản lý vòng đời của các tác vụ con một cách hiệu quả.
    
3. **Task Scheduling Framework (Khung lập lịch tác vụ):** Giao diện `Executor` chuẩn hóa việc gọi, lập lịch, thực thi và kiểm soát các tác vụ bất đồng bộ theo một tập hợp các chính sách thực thi. Các triển khai có sẵn cho phép thực thi các tác vụ trong luồng gửi, trong một luồng nền duy nhất, trong một luồng mới được tạo, hoặc trong một nhóm luồng (thread pool). Các chính sách như giới hạn độ dài hàng đợi và chính sách bão hòa có thể cải thiện tính ổn định của ứng dụng bằng cách ngăn chặn việc sử dụng tài nguyên quá mức.
    
4. **Fork/Join Framework (Khung Fork/Join):** Dựa trên lớp `ForkJoinPool`, khung này là một triển khai của `Executor`, được thiết kế để chạy hiệu quả một số lượng lớn các tác vụ bằng cách sử dụng một nhóm các luồng công nhân. Kỹ thuật "work-stealing" được sử dụng để giữ cho tất cả các luồng công nhân bận rộn, tận dụng tối đa các bộ xử lý đa lõi.
    
5. **Concurrent Collections (Bộ sưu tập đồng thời):** Bao gồm các cấu trúc dữ liệu như `Queue`, `BlockingQueue` và `BlockingDeque`, được tối ưu hóa cho truy cập đồng thời. Ví dụ, `ConcurrentLinkedDeque` là một deque không giới hạn dựa trên các nút liên kết, cho phép các thao tác chèn, loại bỏ và truy cập đồng thời an toàn giữa các luồng.
    
6. **Atomic Variables (Biến nguyên tử):** Các lớp tiện ích cung cấp các hoạt động nguyên tử trên các biến, giúp tránh các vấn đề đồng bộ hóa phức tạp. Ví dụ, lớp `AtomicInteger` cho phép thực hiện các phép toán tăng, giảm và so sánh-thiết lập một cách nguyên tử.
    
7. **Synchronization Utilities (Tiện ích đồng bộ hóa):** Bao gồm các công cụ như `Semaphore`, `CountDownLatch`, và `CyclicBarrier`, hỗ trợ quản lý đồng bộ hóa giữa các luồng. Ví dụ, `Semaphore` duy trì một tập hợp các giấy phép, mỗi `acquire()` sẽ chặn nếu cần thiết cho đến khi một giấy phép có sẵn, và sau đó lấy nó; mỗi `release()` sẽ thêm một giấy phép, có thể giải phóng một luồng đang chờ.
    

Việc sử dụng các API đồng thời này giúp giảm thiểu công sức lập trình, tăng hiệu suất, độ tin cậy và khả năng bảo trì của ứng dụng. Thay vì phải tự phát triển các thành phần đồng thời phức tạp, các nhà phát triển có thể tận dụng các triển khai tiêu chuẩn, đã được kiểm tra kỹ lưỡng và tối ưu hóa cho hiệu suất cao.

Để biết thêm chi tiết, bạn có thể tham khảo tài liệu chính thức về [Concurrency](https://docs.oracle.com/en/java/javase/21/core/concurrency.html).