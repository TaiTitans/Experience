
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