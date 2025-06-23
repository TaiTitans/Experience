
---
**-> “What happens when my server gets turned off?”**
> Redis has two different forms of per sistence available for writing in-memory data to disk in a compact format. The first method is a point-in-time dump either when certain conditions are met (a number of writes in a given period) or when one of the two dump-to-disk commands is called. The other method uses an append-only file that writes every command that alters data in Redis to disk as it happens. Depending on how careful you want to be with your data, append-only writing can be configured to never sync, sync once per second, or sync at the completion of every operation.

**Nói tóm lại:**

- Khi máy chủ Redis bị tắt, dữ liệu trong bộ nhớ sẽ mất.
- Để tránh mất dữ liệu, Redis cung cấp hai cách lưu dữ liệu vào ổ cứng: sao lưu theo thời điểm (snapshot) và ghi nhật ký lệnh (AOF).
- Bạn có thể chọn cách phù hợp với nhu cầu về tốc độ và an toàn dữ liệu của mình.
---
>To support higher rates of read performance (along with handling failover if the server that Redis is running on crashes), Redis supports master/slave replication where slaves connect to the master and receive an initial copy of the full database. As writes are performed on the master, they’re sent to all connected slaves for updating the slave datasets in real time. With continuously updated data on the slaves, clients can then connect to any slave for reads instead of making requests to the master.

**Giải thích đơn giản:**

- **Master/Slave:**
    - Redis có thể hoạt động theo mô hình "chủ" (master) và "tớ" (slave).
    - Master là máy chủ chính, nơi tất cả các thao tác ghi (thay đổi dữ liệu) diễn ra.
    - Slave là các máy chủ phụ, chúng nhận bản sao dữ liệu từ master.
- **Sao chép dữ liệu:**
    - Khi một slave kết nối với master, nó sẽ nhận một bản sao đầy đủ của dữ liệu.
    - Khi dữ liệu trên master thay đổi (do các thao tác ghi), những thay đổi này sẽ được gửi ngay lập tức đến tất cả các slave.
    - Nhờ đó, các slave luôn có dữ liệu giống hệt master (hoặc gần giống, tùy độ trễ mạng).
- **Tăng hiệu suất đọc:**
    - Thay vì tất cả các client (ứng dụng) phải đọc dữ liệu từ master, chúng có thể đọc từ các slave.
    - Điều này giúp phân tải, làm giảm gánh nặng cho master và tăng tốc độ đọc dữ liệu.
- **Khả năng chịu lỗi (Failover):**
    - Nếu master bị lỗi (máy chủ bị tắt, gặp sự cố), các slave vẫn giữ bản sao dữ liệu.
    - Một trong các slave có thể được "thăng cấp" lên làm master mới, giúp hệ thống tiếp tục hoạt động.

---
