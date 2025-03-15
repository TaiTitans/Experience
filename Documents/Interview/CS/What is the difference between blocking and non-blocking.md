
---
**Chặn (Blocking) và không chặn (Non-blocking) tập trung vào trạng thái của luồng (thread):**

- **Chặn (Blocking):** Khi thực hiện một lời gọi chặn, luồng hiện tại sẽ bị tạm dừng cho đến khi kết quả của lời gọi được trả về. Luồng gọi sẽ không tiếp tục chạy cho đến khi nhận được kết quả.
- **Không chặn (Non-blocking):** Lời gọi không chặn không làm tạm dừng luồng hiện tại khi kết quả chưa sẵn sàng ngay lập tức.

**Ví dụ để hiểu sự khác biệt giữa đồng bộ (synchronous), chặn (blocking), bất đồng bộ (asynchronous) và không chặn (non-blocking):**

- **Đồng bộ (Synchronous):** Giống như việc đun nước, bạn phải đứng đó và xem liệu nước đã sôi hay chưa.
- **Bất đồng bộ (Asynchronous):** Nước đang đun, và khi nước sôi, ấm nước sẽ kêu để thông báo cho bạn (thông báo qua callback).
- **Chặn (Blocking):** Trong quá trình đun nước, bạn không thể làm gì khác, phải đứng chờ bên cạnh.
- **Không chặn (Non-blocking):** Trong khi nước đang đun, bạn có thể làm những việc khác.
### Giải thích bổ sung:

1. **Blocking vs Non-blocking**:
    - **Blocking**: Luồng bị "chặn" (suspended), không thể thực hiện tác vụ khác. Ví dụ: Một thao tác I/O đồng bộ như InputStream.read() sẽ chặn luồng cho đến khi dữ liệu sẵn sàng.
    - **Non-blocking**: Luồng tiếp tục chạy mà không cần chờ. Ví dụ: Trong NIO, Selector cho phép kiểm tra dữ liệu mà không chặn luồng.
2. **Synchronous vs Asynchronous**:
    - **Synchronous**: Quy trình diễn ra tuần tự, phải chờ kết quả trước khi tiếp tục.
    - **Asynchronous**: Quy trình không cần chờ, kết quả được xử lý sau qua callback hoặc thông báo (như CompletableFuture trong Java).
3. **Kết hợp**:
    - Synchronous + Blocking: Đun nước và đứng chờ.
    - Asynchronous + Non-blocking: Đun nước, làm việc khác, ấm kêu thì quay lại. 