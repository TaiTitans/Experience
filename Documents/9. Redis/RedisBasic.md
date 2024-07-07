
---
Redis (Remote Dictionary Server) là một cơ sở dữ liệu NoSQL lưu trữ dữ liệu dưới dạng key-value, thường được sử dụng làm bộ nhớ đệm (cache), hệ thống quản lý phiên (session), hàng đợi tin nhắn (message broker), và lưu trữ dữ liệu tạm thời. Redis nổi bật với hiệu suất cao và tính năng phong phú, bao gồm hỗ trợ các cấu trúc dữ liệu như strings, hashes, lists, sets, sorted sets, bitmaps, hyperloglogs và các geo-spatial indexes.

**Redis** là 1 hệ thống lưu trữ key-value in-memory rất mạnh mẽ và phổ biến hiện nay.

**Redis** nổi bật bởi việc hỗ trợ nhiều cấu trúc dữ liệu khác nhau (hash, list, set, sorted set, string), giúp việc thao tác với dữ liệu cực kì nhanh và thuận tiện.

Các hệ thống ngày nay luôn tìm cách tối ưu performance và **Redis** gần như là một mảnh ghép không thể thiếu trong đó.

---
`lettuce-core`
`spring-data-redis`

Trong đó, `spring-data-redis` là thư viện của Spring giúp chúng ta thao tác với **Redis** dễ dàng hơn.

Còn `lettuce-core` là một thư viện mã nguồn mở, giúp kết nối tới Redis một cách thread-safe bằng nhiều hình thức như synchronous, asynchronous and reactive usage.


---
**Spring Data** hỗ trợ chúng ta thao tác với Redis thông qua các Operations như sau:

1. `opsForValue()`: Kiểu Key-Value thông thường. Với Value là 1 giá trị String tùy ý.
2. `opsForHash()`: Tương ứng với cấu trúc Hash trong Redis. Value là một Object có cấu trúc
3. `opsForList()`: Tương ứng với cấu trúc List trong Redis. Value là một list.
4. `opsForSet()`: Tương ứng với cấu trúc Set trong Redis.
5. `opsForZSet()`: Tương ứng với cấu trúc ZSet trong Redis.


----
**Áp dụng Redis trong một dự án Spring Boot:**

Redis có thể được sử dụng trong nhiều trường hợp khác nhau trong một dự án Spring Boot để cải thiện hiệu suất và khả năng mở rộng của ứng dụng. Dưới đây là một số trường hợp sử dụng phổ biến:


1. **Bộ nhớ đệm (Caching):** Redis thường được sử dụng để lưu trữ dữ liệu truy xuất thường xuyên nhằm giảm tải cho cơ sở dữ liệu chính và cải thiện tốc độ phản hồi của ứng dụng. Ví dụ, bạn có thể cache kết quả của các truy vấn phức tạp hoặc các đối tượng thường xuyên được yêu cầu.
    
2. **Quản lý phiên (Session Management):** Redis có thể được sử dụng để lưu trữ thông tin phiên (session) của người dùng trong các ứng dụng web. Điều này giúp duy trì trạng thái người dùng một cách hiệu quả và đảm bảo rằng thông tin phiên có thể được truy cập nhanh chóng.
    
3. **Hàng đợi công việc (Task Queue):** Redis hỗ trợ các cấu trúc dữ liệu như lists và sorted sets, giúp dễ dàng triển khai các hàng đợi công việc, xử lý các tác vụ không đồng bộ, và điều phối các tiến trình background.
    
4. **Đếm và thống kê:** Redis có thể được sử dụng để theo dõi số lượt truy cập, đếm số lượng sự kiện, và thực hiện các thao tác thống kê thời gian thực một cách nhanh chóng và hiệu quả.
    
5. **Hệ thống pub/sub:** Redis hỗ trợ mô hình publish/subscribe, cho phép các thành phần của ứng dụng giao tiếp với nhau thông qua các kênh tin nhắn, rất hữu ích trong các hệ thống phân tán và microservices.



---

#### Các Lệnh Cơ Bản

- **SET và GET:**
    
    `SET key value 
    `GET key`
    
- **Xóa Dữ Liệu:**
    
    `DEL key`
    
- **Kiểm Tra Sự Tồn Tại của Key:**
    
    `EXISTS key`
    
- **Tăng Giá Trị:**
    
    `INCR key`
    
- **Danh Sách (List):**
    
    `LPUSH list_name value LRANGE list_name 0 -1`

---

### 2. Expiration Policies

Redis cho phép thiết lập thời gian sống (TTL) cho các key, giúp tự động xóa dữ liệu sau một khoảng thời gian nhất định.

- **Thiết Lập TTL:**
    
    `SET key value EX 60  # key sẽ hết hạn sau 60 giây`
    
- **Kiểm Tra TTL:**
    
    `TTL key`
    
- **Xóa TTL:**
    
    `PERSIST key`


---
### 3. Persistence

Redis hỗ trợ hai phương pháp lưu trữ dữ liệu để đảm bảo tính bền vững (durability):

- **Snapshotting (RDB):** Lưu trữ toàn bộ dữ liệu theo khoảng thời gian định kỳ.
    
    - Cấu hình trong `redis.conf`:
        
        
        `save 900 1  # Mỗi 900 giây nếu có ít nhất 1 thay đổi`
        
- **Append-Only File (AOF):** Lưu trữ mỗi lệnh ghi vào một file log.
    
    - Bật AOF trong `redis.conf`:
        
        
        `appendonly yes`


---
### 4. Cache Penetration

Cache Penetration xảy ra khi có nhiều yêu cầu tìm kiếm các key không tồn tại trong cache và database, dẫn đến nhiều truy vấn đến database.

- **Giải Pháp:**
    - Cache các giá trị null.
    - Sử dụng Bloom Filter để kiểm tra sự tồn tại của key trước khi truy vấn cache hoặc database.

### 5. Cache Avalanche

Cache Avalanche xảy ra khi nhiều key hết hạn cùng một lúc, dẫn đến một lượng lớn truy vấn đổ dồn về database.

- **Giải Pháp:**
    - Thiết lập thời gian hết hạn khác nhau cho các key.
    - Sử dụng cơ chế làm mới cache trước khi nó hết hạn.

### 6. Cache Warm-Up

Cache Warm-Up là quá trình tải dữ liệu vào cache trước khi hệ thống bắt đầu xử lý yêu cầu, giúp giảm thiểu độ trễ khi truy cập dữ liệu lần đầu.

- **Cách Thực Hiện:**
    - Tải trước các dữ liệu thường xuyên truy cập khi hệ thống khởi động.
    - Sử dụng các công cụ tự động làm ấm cache.

### 7. Distributed Lock

Redis có thể được sử dụng để triển khai cơ chế khóa phân tán nhằm ngăn chặn các race condition trong môi trường phân tán.

- **Cách Triển Khai:**
    - Sử dụng lệnh `SETNX` (set if not exists) để thiết lập khóa.
    - Thiết lập thời gian sống cho khóa để tránh deadlock.
    - Sử dụng các thư viện như Redlock để triển khai khóa phân tán an toàn.

### 8. Data Sharding

Data Sharding là quá trình phân chia dữ liệu thành các phần nhỏ hơn và lưu trữ chúng trên nhiều máy chủ Redis khác nhau.

- **Cách Thực Hiện:**
    - Sử dụng các công cụ như Redis Cluster hoặc các giải pháp sharding của bên thứ ba.
    - Cấu hình Redis Cluster:
```SHELL
redis-server --cluster-enabled yes
redis-server --cluster-config-file nodes.conf
```

### 9. High Availability

Redis hỗ trợ nhiều phương pháp để đảm bảo tính sẵn sàng cao của hệ thống.

- **Replication:** Redis hỗ trợ cấu hình master-slave để sao chép dữ liệu từ máy chủ chính (master) sang máy chủ phụ (slave).
- **Sentinel:** Redis Sentinel cung cấp cơ chế giám sát, thông báo và failover tự động cho các cụm Redis.