
---
##  Xác định & phân tích vấn đề

- Đầu tiên là thu thập thông tin từ khách hàng để khoanh vùng trang nào, phần nào bị chậm, từ đó truy vết ra nguyên nhân gốc rễ.
- Setup hệ thống `log`, `monitor` để phát hiện những API nào chậm, từ đó sẽ optimize những cái có response time cao nhất.

## Các kỹ thuật tối ưu (Database)

### 1 .Không lấy dư cột
- Giúp giảm tải băng thông. (Tuy nhiên chỉ thực sự áp dụng với các câu cần optimize, đôi khi bạn phải đánh đổi 1 phần nhỏ hiệu suất để các function trong backend có thể được tái sử dụng nhiều. Nếu cứ tối ưu bằng việc loại bỏ các cột dư thừa thì code của bạn lại bị vi phạm vấn đề `DRY` (don't repeat yourself).
### 2. Đánh index

> [!NOTE]
> Nếu dữ liệu nhiều thì nên được đánh index.

- Tuy nhiên đánh index cũng phải sao cho đúng, `tránh đánh index mổ cỏ`, tức là đánh nhiều index trong table mà mỗi index chỉ có 1 column, cũng không nên đánh index trên quá nhiều cột (chọn tối đa từ 2-4 cột là đẹp).
- Thứ tự đánh index của các cột là rất quan trọng, nên chọn cột nào càng loại bỏ được nhiều data không khớp trước càng tốt. Index xong thì nhớ dùng WHERE với các column đã index để tận dụng tối đa hiệu quả.
- Không đánh index đơn lẻ cho những trường nào có giá trị trùng lặp nhiều (`low cardinality`), nếu vẫn muốn đánh thì nên kết hợp thêm với cột khác. Ví dụ cột giới tính chỉ có 3 dạng (male, female, other) thì dùng index cho riêng cột này sẽ không đạt hiệu quả cao, lúc này có thể thêm 1 column khác nữa như company_id chẳng hạn.
- - Các database sẽ sử dụng nhiều thuật toán index khác nhau, hãy lựa chọn đúng trong từng trường hợp. Ví dụ với postgres có các loại index như `b-tree`, `hash-index`, `GIN`... tuỳ mỗi hoàn cảnh mà mình sẽ lựa chọn 1 index phù hợp để tận dụng tối đa thuật toán index đó.
- Áp dụng `partial index` cho những table có dữ liệu lớn.

> [!NOTE] Chỉ mục một phần (Partial Index) là gì?
>**Chỉ mục một phần** hay **chỉ mục được lọc** là một loại chỉ mục đặc biệt trong cơ sở dữ liệu, chỉ bao gồm một phần dữ liệu của một bảng, chứ không phải toàn bộ. Phần dữ liệu này được xác định bởi một điều kiện cụ thể (gọi là **điều kiện lọc**). Nói cách khác, chỉ mục này chỉ xây dựng trên các hàng đáp ứng điều kiện lọc đó.
- Chú ý đến một số câu lệnh `order by`, mình từng gặp trường hợp lúc đánh index các kiểu cho một table nhưng lúc query ra vẫn chậm rì, đi lần mò `explain analyze` mãi mới biết vì mình dùng order by để sort lại dữ liệu làm nó không ăn được index. Để giải quyết thì có nhiều cách, 1 là đánh index luôn cho cột đang sort đó, hoặc bạn có thể thử biện pháp `defered join` chẳng hạn.

> [!NOTE] **Deferred Join** 
> **Deferred Join** (hay còn gọi là **Delayed Join**) là một kỹ thuật tối ưu hóa cơ sở dữ liệu, cho phép trì hoãn việc thực hiện một hoặc nhiều liên kết (join) trong một truy vấn SQL cho đến khi cần thiết. Điều này có thể giúp cải thiện hiệu suất truy vấn, đặc biệt là trong các truy vấn phức tạp với nhiều bảng liên kết.

### 3. Join

- Sử dụng `filtered join`: giống partial index, trong một vài trường hợp mình sẽ thêm vài điều kiện để loại bỏ bớt các record được join vào.
- Sử dụng `subquery` với `EXISTS hoặc IN`.
- `Lateral Joins` (trong PostgreSQL): Đây là một loại join cho phép bạn tham chiếu đến các cột từ các bảng được join trước đó trong mệnh đề FROM, cho phép các join phức tạp và có điều kiện.
- `Defered joins`: Nôm na là mình select kết quả ra trước để lấy id rồi mới đem join lại với chính bảng đó để lấy toàn bộ record. Cách này có thể dùng trong một số trường hợp như khi bạn sử dụng order by khiến câu query không ăn được index.

### 4. Xử lý các tác vụ nặng

- Phần lớn việc nghẽn cổ chai (bottle neck) đến từ database chứ không phải nơi nào khác. Đặc biệt với các database SQL khó nâng cấp theo chiều ngang được (khó chứ không phải không). Backend server thì có thể scale up tốt hơn. Nên việc gì khó cứ lấy về backend rồi để nó xử lý.
- Lưu trước kết quả: Có một số thống kê mình phải join vào nhiều bảng khác nhau, tính toán các thứ để lấy ra kết quả cuối cùng trả về cho người dùng, việc này thường mất nhiều thời gian. Lúc này mình sẽ lựa chọn tính toán trước dữ liệu rồi lưu vào 1 table nào đó, khi nào cần chỉ việc lấy ra cho nhanh. Cơ mà cái gì cũng phải đánh đổi, bạn nhanh hơn ở bước đọc thì phải tốn công viết code insert, update khi có gì đó thay đổi, tăng write vô database.
- Những dữ liệu không cần thiết real-time, nếu được nên sài `cronjob` để tính toán rồi ghi vào database, hãy chọn thời gian ít người dùng nhất, giảm tải cho database trong giờ cao điểm.

### 5.Connection

- Mỗi database sẽ có số lượng Input/Output Operations Per Second (`IOPS`), cần được phân chia và tận dụng tốt. Ngày đi học thì chúng ta hay được thầy cô dạy dùng `singleton pattern` để kết nối đến cơ sở dữ liệu. Singleton được cái tiết kiệm tài nguyên, chỉ tạo 1 kết nối duy nhất đến database, tránh phải tạo dư thừa connection. Nhưng nhược điểm là nó sẽ gây ra độ trễ vì chỉ có 1 connection duy nhất đến database để xử lý, các yêu cầu sẽ phải chờ nhau trong khi database thì còn quá trời `IOPS`.
- Bình thường mình hay sử dụng `connection pooling` để kết nối đến database. Mình set một số lượng `min connection` được mở để dùng, các tác vụ có thể tận dụng connection từ pool, dùng xong thì trả lại cho thằng khác sài, giúp tiết kiệm được tài nguyên, bên cạnh đó mình cũng set cả `max connection` có thể kết nối đồng thời để tránh làm quá tải database.

### 6. Scale up

- Tác vụ **read thường gấp 9 lần write**. Khi hệ thống đạt đến ngưỡng scale up dọc (tăng server vật lý), khi không dọc được nữa mình scale ngang. Mô hình replica gồm 1 `primary node` để write và nhiều replica node để read. Việc tách biệt read/write giúp tăng khả năng xử lý. Có nhiều RAM hơn để lưu cache, query sẽ nhanh hơn.
- Với các noSQL thì được cái khá mạnh về scale chiều ngang với replica set và sharding nên có thể tận dụng yếu tố này. Thường mình nói phần này ra nhà tuyển dụng hay vặn tiếp về distribute system như thế nào, nên ai phỏng vấn mà lỡ nói thì hãy chắc là bạn hiểu cách scale up ngang nhé.

### 7.Search & filter

- Thống nhất cách filter và search trong hệ thống, giúp tiết kiệm thời gian build source code.
- Nếu bạn cần 1 tính năng search mạnh mẽ mà database bạn sử dụng lại không mạnh về search thì nên dùng database search riêng (như Elasticsearch, Sorl). 

### 8.Partition & remove old data

- Dữ liệu sẽ ngày càng nhiều lên dẫn đến việc query sẽ bị chậm dần. Lúc này việc partition là 1 việc rất tốt, nó giúp chia nhỏ bảng (hoặc index) thành các bảng nhỏ hơn được lưu trong các phân vùng vật lý khác đễ dễ dàng truy xuất. Ví dụ bạn xây dựng trang facebook, bài đăng mới thì ngày nào cũng tăng lên, nhưng thực tế người dùng sẽ ít khi coi lại những bài đăng cách đây 5 năm, 10 năm của họ. Lúc này mình sẽ dùng partition theo range time.
- Note: `partition` thì không được tự động tạo, thường sẽ phải viết `trigger` hoặc `cronjob` để sau một khoảng thời gian (hoặc theo một điều kiện nào đó của bạn setup) thì nó sẽ tạo thêm partition mới.
- Với những data rất rất lâu rồi, khách hàng gần như không sài nhưng vẫn phải giữ lại nếu họ cần trong tương lai. Lúc này có thể di chuyển nó qua một vùng chưa dài hạn khác (ổ cứng, S3...) nhằm giảm tải cho database chính.
### 9.Cache ngu là cook

- Sử dụng cơ chế `materialize view` để cache trước những dữ liệu báo cáo (ít thay đổi), cần là có đem ra sài luôn.
- Thông thường thì database có sẵn `cache in-memory` nhưng thường nó sẽ tự chọn thứ mà nó muốn cache, nếu bạn muốn cache chủ động hơn, đồng thời giảm bớt việc cho database chính, thì có thể cân nhắc sử dụng các database chuyên biệt cho caching như `redis`.
### 11.Dữ liệu bị phân mảnh

- Dữ liệu trong SQL (ví dụ postgreSQL) sẽ được lưu trữ trong các page có dung lượng 8kb, đôi khi do việc đọc ghi, xoá quá nhiều dẫn đến dữ liệu bị lưu trữ trên nhiều page khác nhau, mặc dù dữ liệu lại chả có bao nhiêu. Giống như kiểu bạn mua 1 quyển sổ 200 trang nhưng mỗi trang bạn chỉ viết 1 2 từ, điều này gây lãng phí không gian và giảm hiệu suất.
- Để khắc phục thì trước tiên phải tìm hiểu xem có dữ liệu có bị phân mảnh không thông qua câu lệnh:
```
SELECT 
    schemaname || '.' || relname AS table_name,
    pg_size_pretty(pg_total_relation_size(relid)) AS total_size,
    pg_size_pretty(pg_table_size(relid)) AS table_size,
    pg_size_pretty(pg_indexes_size(relid)) AS index_size,
    pg_size_pretty(pg_total_relation_size(relid) - pg_table_size(relid) - pg_indexes_size(relid)) AS toast_size,
    round(100 * (pg_total_relation_size(relid) - pg_table_size(relid) - pg_indexes_size(relid)) / 
        NULLIF(pg_total_relation_size(relid), 0), 2) AS fragmentation_percent
FROM pg_stat_user_tables
WHERE relname = '&TABLE_NAME'
ORDER BY pg_total_relation_size(relid) DESC;

```

- Nếu kiểm tra thấy có table thì tức là dữ liệu đã bị phân mảnh rồi, lúc này có thể ngâm cứu tiến hành chữa bệnh cho nó.
- Giải quyết thì có thể tạo lại bảng mới, chạy lại index.
- Thiết kế schema hiệu quả hơn để giảm tải phân mảnh ngay từ đầu.