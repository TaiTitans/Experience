
---

# Time-To-Live (TTL) trong Redis hoạt động như nào?

## Định nghĩa Time-To-Live (TTL)

Time-To-Live là khái niệm phổ biến được sử dụng trong 2 lĩnh vực:

- **Mạng máy tính (Computer Networking)**: Mỗi packet (network layer) chứa 1 trường TTL dùng để đếm số router (hop) mà packet có thể đi qua được. Khi đi qua mỗi router (hop), TTL sẽ bị trừ đi 1 đơn vị. Nếu TTL = 0 thì packet sẽ bị drop.
- **Bộ nhớ đệm (Caching)**: TTL ám chỉ khoảng thời gian mà dữ liệu được lưu trước khi dữ liệu đó bị xoá đi. Việc xoá dữ liệu này phản ánh **dữ liệu có thể còn không hợp lệ** và giúp **giải phóng bộ nhớ**.


## Time-To-Live trong Redis hoạt động như nào?

*Khi một key được cài đặt Time-To-Live, sau khoảng thời gian TTL đó, key đó sẽ tự động được xoá đi.*

**Cách tiếp cận ban đầu**
Redis sử dụng và kết hợp đồng thời 2 cách để xoá key khi hết hạn:

**Passive Way:**

Redis không xoá key ngay khi hết hạn. Cho tới khi ta lấy (GET) key đó, Redis mới check key đã hết hạn chưa. Nếu đã hết hạn thì sẽ trả về null trước rồi xoá key đó đi một cách bất đồng bộ.

Cách này khá đơn giản nhưng gặp 1 vấn đề. Đó là có những key mà sẽ bao giờ được động tới, như vậy nó sẽ chiếm không gian nhớ.
Để khắc phục tình trạng này, Redis kết hợp với Active Way - xoá chủ động.

**Active Way:**

Redis thực hiện job xoá key chủ động 10 lần / giây (trong code gọi là activeExpireCycle)

Bước 1: Lấy random 20 keys có gắn TTL. Trong có thể có keys hết hạn và key chưa hết hạn.
Bước 2: Thử xoá 20 keys đó
Bước 3: Nếu có nhiều hơn 25% của 20 keys bị hết hạn, lặp lại bước 1.
Do Redis hoạt động dựa trên single-thread, khi job xoá key chủ động chạy thì sẽ block các request khác. Do đó, Redis có cấu hình timeout cho job này tránh tình trạng bị treo do quá nhiều key đồng thời hết hạn.

Việc kết hợp 2 cách này hoạt động rất tốt trong hầu hết các trường hợp thực tế. Tuy nhiên vẫn có trường số lượng key expire lớn và việc lấy sampling random khiến job activeExpireCycle hoạt động không hiệu quả.

**Cách tiếp cận mới**
Redis khắc phục nhược điểm của cách tiếp cận trên bằng cách: **các key sẽ được sắp xếp bởi thời gian hết hạn (expiration time) và lưu trong Sorted Set (ZSET)** giúp cho việc tìm kiếm key hết hạn nhanh chóng hơn.

***Best Practices khi đặt TTL***

**Dựa trên yêu cầu nghiệp vụ**: dựa vào ý nghĩa của dữ liệu mà ta đặt TTL cho phù hợp. Ví dụ dữ liệu về session thì có TTL ngắn hơn so với static content như ảnh, đường dẫn, …
**Đối với những TTL dài:** thì ta cần random các giá trị này để chúng được dải đều ra. Tránh tình trạng “thundering herd" xảy ra khi tất cả key đồng thời hết hạn, gây ra load cao ở các tầng app và DB
**Sử dụng đúng kiểu dữ liệu:** Không phải tất cả kiểu dữ liệu điều có cơ chế TTL giống nhau. Ví dụ, List và Set hỗ trợ TTL theo từng phần tử, còn Sorted Set và Hash thì không.
**Chú ý cấu hình Eviction Policy dựa trên yêu cầu của hệ thống:**. Khi Redis ăn hết mem, nó sẽ evict key dựa trên cấu hình (LRU, LFU, …)
***Nếu không có yêu cầu đặc biệt từ nghiệp vụ thì thông thường nên đặt TTL ngắn để tránh chiếm mem và dữ liệu được update.***