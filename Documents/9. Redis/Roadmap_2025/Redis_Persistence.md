
---
Redis Persistence là một chủ đề quan trọng khi bạn sử dụng Redis trong môi trường production. Nó giúp đảm bảo rằng dữ liệu không bị mất khi Redis bị tắt hoặc gặp sự cố. Có hai cơ chế chính trong Redis Persistence:

- **RDB (Redis Database File - Snapshotting)**
- **AOF (Append-Only File - Logging từng thao tác ghi)**

Mỗi phương pháp có ưu và nhược điểm riêng, và nhiều hệ thống kết hợp cả hai để đạt hiệu suất tối ưu. 

## 1. **RDB (Redis Database File - Snapshotting)**

### 🛠 **Cách hoạt động**

- RDB lưu trạng thái toàn bộ dataset vào file `.rdb` tại một thời điểm nhất định.
- Cách thức hoạt động:
    1. Redis **forks** (tạo một tiến trình con).
    2. Tiến trình con sẽ sao chép toàn bộ dữ liệu hiện tại và ghi nó vào file RDB.
    3. Tiến trình chính vẫn tiếp tục phục vụ các request mà không bị block đáng kể.
    4. Sau khi ghi xong, tiến trình con kết thúc, file RDB mới sẽ thay thế file cũ.

### 🔧 **Cấu hình RDB trong Redis**

File cấu hình `redis.conf` có thể được chỉnh sửa để điều chỉnh tần suất snapshot:
```
save 900 1    # Snapshot mỗi 900 giây nếu có ít nhất 1 thay đổi
save 300 10   # Snapshot mỗi 300 giây nếu có ít nhất 10 thay đổi
save 60 10000 # Snapshot mỗi 60 giây nếu có ít nhất 10,000 thay đổi
```

### ⚡ **Ưu điểm của RDB**

✔ **Hiệu suất cao**: Do tiến trình fork chạy riêng biệt, không ảnh hưởng quá nhiều đến hiệu suất Redis chính.  
✔ **Khởi động nhanh**: Vì chỉ lưu trữ snapshot nên quá trình load lại dữ liệu nhanh hơn so với AOF.  
✔ **Chiếm ít dung lượng hơn**: Vì chỉ lưu trạng thái tại các thời điểm cụ thể, không ghi lại từng thao tác nhỏ lẻ như AOF.

### ❌ **Nhược điểm của RDB**

✖ **Có thể mất dữ liệu**: Vì Redis chỉ snapshot sau một khoảng thời gian, dữ liệu thay đổi giữa các lần snapshot có thể bị mất nếu hệ thống gặp sự cố.  
✖ **Tốn tài nguyên khi snapshot**: Tiến trình fork tiêu tốn CPU và RAM do phải sao chép dữ liệu, có thể gây vấn đề trong hệ thống có bộ nhớ lớn.


## 2. **AOF (Append-Only File - Logging toàn bộ thao tác ghi)**

### 🛠 **Cách hoạt động**

- AOF ghi lại từng lệnh thay đổi dữ liệu (write operation) vào file `.aof`.
- Khi Redis khởi động lại, nó đọc file này và thực hiện lại từng thao tác để khôi phục dữ liệu.
- Có ba chế độ ghi file AOF:
    1. **always**: Ghi dữ liệu ngay lập tức sau mỗi thao tác (an toàn nhất nhưng chậm nhất).
    2. **everysec**: Ghi dữ liệu mỗi giây một lần (cân bằng giữa hiệu suất và an toàn).
    3. **no**: Để hệ điều hành tự quyết định khi nào ghi (có thể mất dữ liệu nếu Redis crash).

### 🔧 **Cấu hình AOF trong Redis**
```
appendonly yes
appendfsync everysec  # Cấu hình an toàn nhất mà vẫn giữ hiệu suất tốt
```
### ⚡ **Ưu điểm của AOF**

✔ **An toàn hơn**: Mất ít dữ liệu hơn vì nó ghi lại từng thao tác.  
✔ **Dễ đọc và chỉnh sửa**: File AOF lưu trữ ở dạng text với các lệnh Redis, có thể chỉnh sửa thủ công.  
✔ **Tương thích với các phiên bản Redis khác nhau**: Dễ dàng đồng bộ dữ liệu giữa các hệ thống Redis.

### ❌ **Nhược điểm của AOF**

✖ **Tốn dung lượng hơn**: File AOF thường lớn hơn file RDB vì nó lưu mọi thao tác thay đổi dữ liệu.  
✖ **Khởi động chậm hơn**: Khi Redis restart, nó phải đọc file AOF và thực hiện lại tất cả lệnh trước đó.  
✖ **Cần compact dữ liệu thường xuyên**: File AOF có thể phình to theo thời gian, cần chạy `BGREWRITEAOF` để giảm kích thước file.


## Kết hợp cả RDB và AOF
Trong production, ta có thể bật cả hai để tận dụng ưu điểm của mỗi loại:
```
save 900 1
save 300 10
save 60 10000
appendonly yes
appendfsync everysec
```
- Khi Redis khởi động, nó sẽ ưu tiên sử dụng AOF nếu có cả hai file AOF và RDB.
- Trong trường hợp cần **khởi động nhanh**, ta có thể vô hiệu hóa AOF bằng cách chạy:
```
redis-cli config set appendonly no
```

## **Redis Persistence và High Availability**

- **Replication** (Master-Slave) kết hợp với AOF có thể tăng tính an toàn cho dữ liệu.
- **Sentinel** giúp giám sát Redis Master-Slave và tự động failover khi Master bị down.
- **Cluster Mode** giúp phân mảnh dữ liệu trên nhiều node để mở rộng quy mô.

## **Khi nào nên chọn RDB hoặc AOF?**

- Nếu bạn **ưu tiên hiệu suất** và **chấp nhận mất một ít dữ liệu** → Chọn RDB.
- Nếu bạn **muốn đảm bảo không mất dữ liệu quan trọng** → Chọn AOF.
- Nếu bạn **muốn cân bằng giữa hiệu suất và an toàn** → Kết hợp cả RDB + AOF (`appendfsync everysec`).

