
---
`Socket.IO là một thư viện mạnh mẽ dùng để giao tiếp hai chiều (real-time bidirectional communication) giữa client và server thông qua giao thức WebSocket và các kỹ thuật dự phòng khác như long-polling (đối với những trình duyệt không hỗ trợ WebSocket). Socket.IO được xây dựng trên nền tảng Node.js và giúp phát triển các ứng dụng thời gian thực trở nên dễ dàng hơn.`

### Ưu điểm của Socket.IO:

- **Dễ sử dụng**: Cung cấp giao diện dễ hiểu để thực hiện các chức năng phức tạp.
- **Đa nền tảng**: Tự động hỗ trợ fallback sang các kỹ thuật dự phòng nếu WebSocket không được hỗ trợ.
- **Sự kiện tùy chỉnh**: Hỗ trợ việc tạo và lắng nghe các sự kiện tùy chỉnh, không chỉ đơn thuần là gửi/nhận tin nhắn.
- **Quản lý phòng (rooms) và kênh (namespaces)**: Hỗ trợ tính năng chia phòng (rooms) để quản lý nhiều kết nối dễ dàng hơn, rất hữu ích cho các ứng dụng chat hoặc game đa người chơi.

