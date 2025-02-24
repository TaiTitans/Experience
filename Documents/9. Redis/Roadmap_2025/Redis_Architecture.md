
---
Redis có một kiến trúc đơn giản nhưng mạnh mẽ, giúp nó hoạt động cực kỳ nhanh và có thể mở rộng linh hoạt. Kiến trúc Redis bao gồm các thành phần sau:

## **1️⃣ Kiến trúc tổng quan**

Redis hoạt động theo mô hình **Client-Server**, nơi client gửi yêu cầu đến Redis Server và nhận phản hồi. Nó chủ yếu hoạt động trên **bộ nhớ RAM**, giúp truy vấn cực nhanh.

📌 **Các thành phần chính:**

1. **Client** – Ứng dụng hoặc người dùng gửi request đến Redis Server.
2. **Redis Server** – Máy chủ chính xử lý yêu cầu từ client.
3. **Storage Engine** – Bộ máy lưu trữ dữ liệu chủ yếu trên RAM.
4. **Persistence Mechanism** – Hệ thống lưu trữ dữ liệu ra ổ đĩa để đảm bảo không mất dữ liệu khi Redis restart.
5. **Replication** – Hỗ trợ kiến trúc **Master-Slave** để tăng độ tin cậy và hiệu suất.
6. **Cluster** – Cho phép Redis phân tán dữ liệu và xử lý nhiều node trong môi trường phân tán.
7. **Sentinel** – Cung cấp high availability bằng cách theo dõi các node và tự động failover.

## **2️⃣ Cấu trúc bộ nhớ**

Redis lưu dữ liệu chủ yếu trong RAM để tăng tốc truy vấn. Khi cần, nó có thể ghi dữ liệu xuống ổ cứng để tránh mất dữ liệu.

📌 **Các thành phần bộ nhớ chính:**

- **Main Memory (RAM):** Nơi Redis lưu trữ key-value data.
- **Disk (RDB & AOF):** Dùng để lưu trữ dữ liệu vĩnh viễn.
- **Virtual Memory (VM):** Khi Redis bị thiếu RAM, nó có thể swap dữ liệu ra đĩa.

## **3️⃣ Cơ chế xử lý request**

📌 **Redis sử dụng mô hình Single-threaded Event Loop**, giúp nó xử lý nhanh mà không cần lock.

🔥 **Các bước xử lý request trong Redis:**

1. **Client gửi request** đến Redis Server qua TCP (hoặc Unix Socket).
2. **Redis nhận request** và đưa vào hàng đợi event loop.
3. **Redis xử lý request**, thực thi lệnh trên bộ nhớ.
4. **Redis gửi phản hồi** về client.

🔹 Vì Redis chạy đơn luồng (single-threaded), nó không cần sử dụng locks hay context switching, giúp tốc độ nhanh hơn so với database truyền thống.

## **4️⃣ Cơ chế lưu trữ dữ liệu – Persistence**

Redis có hai cơ chế lưu trữ dữ liệu chính để tránh mất dữ liệu khi restart:

### ✅ **1. RDB (Redis Database Snapshot)**

- Chụp nhanh dữ liệu theo từng khoảng thời gian.
- Lưu dữ liệu dưới dạng file `.rdb`.
- Ít ảnh hưởng đến hiệu suất vì ghi dữ liệu định kỳ.
- Có thể mất dữ liệu nếu Redis crash giữa hai lần snapshot.

**Lệnh liên quan:**

- `SAVE`: Lưu dữ liệu ngay lập tức.
- `BGSAVE`: Lưu dữ liệu ở background.

### ✅ **2. AOF (Append-Only File)**

- Lưu tất cả lệnh ghi (write) vào file `.aof`.
- Khôi phục dữ liệu bằng cách chạy lại các lệnh từ file này.
- Chậm hơn RDB nhưng đảm bảo không mất dữ liệu.

**Lệnh liên quan:**

- `CONFIG SET appendonly yes`: Bật AOF.
- `BGREWRITEAOF`: Dọn dẹp file AOF để giảm kích thước.

**Nên dùng cái nào?**

- **Chỉ RDB** nếu hiệu suất quan trọng hơn an toàn dữ liệu.
- **Chỉ AOF** nếu cần lưu toàn bộ dữ liệu.
- **Cả hai** nếu cần tối ưu cả hiệu suất và độ an toàn.

## **5️⃣ Replication – Nhân bản dữ liệu**

Redis hỗ trợ **Master-Slave Replication**, giúp tăng hiệu suất và độ tin cậy.

📌 **Mô hình Master-Slave:**

- **Master** xử lý ghi (write).
- **Slave** chỉ xử lý đọc (read).
- Nếu Master chết, Slave có thể trở thành Master mới.

**Lệnh liên quan:**

- `SLAVEOF <host> <port>`: Cấu hình Slave.
- `INFO REPLICATION`: Kiểm tra trạng thái replication.

## **6️⃣ Redis Sentinel – High Availability**

Redis Sentinel là một hệ thống giúp giám sát, phát hiện lỗi và tự động failover cho Redis.

📌 **Chức năng chính:**

- Giám sát Redis Master-Slave.
- Tự động chuyển Slave thành Master nếu Master chết.
- Gửi thông báo khi có sự cố.

🔹 **Triển khai Redis Sentinel:**
```
redis-sentinel /etc/redis/sentinel.conf
```

## **7️⃣ Redis Cluster – Distributed Redis**

Redis Cluster giúp phân tán dữ liệu trên nhiều node để tăng hiệu suất và khả năng chịu lỗi.

📌 **Nguyên tắc hoạt động:**

- Dữ liệu được chia thành **16384 slots** và phân bổ cho các node.
- Hỗ trợ **Sharding** để chia tải.
- Có ít nhất **3 Master nodes** để đảm bảo không mất dữ liệu.

```bash
redis-cli --cluster create 192.168.1.1:7000 192.168.1.2:7001 192.168.1.3:7002 --cluster-replicas 1
```

## **8️⃣ Cơ chế giao tiếp trong Redis**

Redis hỗ trợ nhiều cơ chế giao tiếp khác nhau:  
✅ **Redis Protocol (RESP)** – Protocol riêng của Redis, tối ưu tốc độ.  
✅ **Pub/Sub** – Hệ thống message queue, cho phép các client subscribe nhận sự kiện real-time.  
✅ **Streams** – Hàng đợi tin nhắn mạnh hơn Pub/Sub, hỗ trợ lưu trữ.

## 🚀 **Tóm tắt kiến trúc Redis**

| Thành phần            | Chức năng                                                |
| --------------------- | -------------------------------------------------------- |
| **Client-Server**     | Kiến trúc TCP đơn giản, xử lý request cực nhanh.         |
| **In-Memory Storage** | Lưu dữ liệu trên RAM, truy xuất tốc độ cao.              |
| **Persistence**       | Lưu trữ bằng RDB và AOF để tránh mất dữ liệu.            |
| **Replication**       | Master-Slave để nhân bản dữ liệu, tăng tốc độ đọc.       |
| **Sentinel**          | Hỗ trợ High Availability, tự động failover.              |
| **Cluster**           | Phân tán dữ liệu trên nhiều node, tăng khả năng mở rộng. |
| **Pub/Sub & Streams** | Hỗ trợ message queue, event-driven.                      |