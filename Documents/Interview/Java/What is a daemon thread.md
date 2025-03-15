
---
**Daemon (luồng nền) là một loại tiến trình đặc biệt chạy ở chế độ nền.**

- Nó độc lập với thiết bị điều khiển (terminal) và định kỳ thực hiện một nhiệm vụ nhất định hoặc chờ xử lý các sự kiện cụ thể xảy ra.
- Trong Java, các luồng thu gom rác (garbage collection threads) là một loại luồng daemon đặc biệt.

### Giải thích bổ sung:

- **Daemon Thread trong Java**: Là các luồng chạy ở chế độ nền để hỗ trợ các luồng chính (user threads). Chúng có thể được thiết lập bằng phương thức setDaemon(true) trước khi khởi chạy luồng (ví dụ: Thread t = new Thread(); t.setDaemon(true); t.start();).
- **Ví dụ**: Luồng thu gom rác trong JVM là daemon thread, hoạt động ngầm để dọn dẹp bộ nhớ mà không cần sự can thiệp trực tiếp của chương trình.
- **Đặc điểm**: JVM sẽ thoát khi không còn user thread nào hoạt động, bất kể daemon thread có đang chạy hay không.
