
---
### **Tối ưu hiệu suất truy vấn và xử lý dữ liệu trong hệ thống**

1. **Tối ưu chỉ mục (Indexing)**
    
    - Thêm **chỉ mục** vào các cột được sử dụng trong điều kiện `WHERE` hoặc các cột sắp xếp theo `ORDER BY`.
    - Tránh lạm dụng chỉ mục, vì chỉ mục quá nhiều có thể làm giảm hiệu suất ghi dữ liệu.
2. **Tối ưu câu lệnh SQL**
    
    - **Tránh `SELECT *`**, chỉ lấy những cột cần thiết để giảm băng thông và tài nguyên xử lý.
    - **Hạn chế phân trang sâu (Deep Pagination)** bằng cách sử dụng `OFFSET` hợp lý hoặc lưu trữ dữ liệu phân trang trong bộ nhớ tạm.
    - **Tối ưu `GROUP BY`**, có thể sử dụng **index covering** để tăng tốc độ truy vấn nhóm.
3. **Tránh giao dịch lớn (Big Transactions)**
    
    - **Giao dịch lớn** có thể khóa bảng trong thời gian dài, gây tắc nghẽn hệ thống.
    - **Chia nhỏ giao dịch**, chỉ đặt các thao tác cần thiết trong một transaction.
    - **Xử lý logic ngoài giao dịch**, tránh các thao tác không cần thiết bên trong transaction.
4. **Xử lý bất đồng bộ (Asynchronous Processing)**
    
    - **Tách logic chính và phụ** để tránh làm chậm luồng xử lý chính.
    - Ví dụ: Khi đơn hàng được giao, **gửi SMS thông báo** có thể được thực hiện **bất đồng bộ** để không làm chậm tiến trình giao hàng.
5. **Giảm độ chi tiết của khóa (Lock Granularity)**
    
    - Trong hệ thống có **nhiều luồng (concurrent threads)** truy cập cùng lúc, **lock không hợp lý** có thể ảnh hưởng đến hiệu suất.
    - Sử dụng **optimistic locking** thay vì **pessimistic locking** nếu có thể.
    - Cân nhắc **khoá cấp cột (row-level locking)** thay vì **khoá cấp bảng (table-level locking)** để tối ưu hiệu suất.
6. **Sử dụng bộ nhớ đệm (Caching)**
    
    - **Redis / Memcached** giúp tăng tốc truy vấn dữ liệu **thường xuyên sử dụng** thay vì truy vấn trực tiếp từ database.
    - Cần xác định dữ liệu nào nên cache, thời gian sống của cache (`TTL`), và cơ chế **cache eviction** để đảm bảo dữ liệu luôn nhất quán.
7. **Chia nhỏ cơ sở dữ liệu và bảng (Database & Table Sharding)**
    
    - Khi lượng truy cập cao, chia nhỏ **cơ sở dữ liệu (database sharding)** giúp **giảm tải kết nối** và **cải thiện hiệu suất I/O**.
    - Khi dữ liệu bảng quá lớn, chia nhỏ **bảng (table partitioning)** giúp **giảm thời gian truy vấn**.
    - Áp dụng các chiến lược **horizontal partitioning, vertical partitioning, hoặc consistent hashing** tùy theo yêu cầu hệ thống.
8. **Tránh truy vấn trong vòng lặp (Avoid Loop Queries)**
    
    - Truy vấn **trong vòng lặp** làm tăng số lần gọi database, gây quá tải và giảm hiệu suất.
    - **Tối ưu bằng cách sử dụng truy vấn `IN (...)` hoặc `JOIN`**, thay vì lặp từng phần tử để truy vấn riêng lẻ.

➡ **Kết luận:** Khi xây dựng hệ thống hiệu suất cao, cần kết hợp nhiều kỹ thuật tối ưu: **cấu trúc dữ liệu, SQL tuning, caching, sharding, và xử lý bất đồng bộ** để đảm bảo hệ thống vận hành ổn định và nhanh chóng. 🚀