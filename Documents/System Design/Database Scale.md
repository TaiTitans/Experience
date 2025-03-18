
---
**Mở rộng cơ sở dữ liệu**  
Có hai cách tiếp cận chính để mở rộng cơ sở dữ liệu: mở rộng theo chiều dọc và mở rộng theo chiều ngang.

**Mở rộng theo chiều dọc**  
Mở rộng theo chiều dọc, còn được gọi là "scaling up", là việc mở rộng bằng cách tăng thêm sức mạnh (CPU, RAM, ổ cứng, v.v.) cho một máy chủ hiện có. Hiện nay, có một số máy chủ cơ sở dữ liệu rất mạnh mẽ. Theo Dịch vụ Cơ sở dữ liệu Quan hệ Amazon (RDS) [12], bạn có thể sở hữu một máy chủ cơ sở dữ liệu với dung lượng RAM lên tới 24 TB. Loại máy chủ cơ sở dữ liệu mạnh mẽ này có thể lưu trữ và xử lý một lượng lớn dữ liệu. Ví dụ, vào năm 2013, stackoverflow.com có hơn 10 triệu lượt truy cập độc nhất mỗi tháng, nhưng chỉ sử dụng một cơ sở dữ liệu chính [13]. Tuy nhiên, mở rộng theo chiều dọc cũng đi kèm với một số nhược điểm nghiêm trọng:  
• Bạn có thể bổ sung thêm CPU, RAM, v.v. cho máy chủ cơ sở dữ liệu, nhưng sẽ có giới hạn phần cứng. Nếu bạn có một lượng người dùng lớn, một máy chủ duy nhất sẽ không đủ.  
• Nguy cơ xảy ra lỗi tại một điểm duy nhất (single point of failure) cao hơn.  
• Chi phí tổng thể của việc mở rộng theo chiều dọc rất cao. Các máy chủ mạnh mẽ có giá thành đắt đỏ hơn nhiều.

**Mở rộng theo chiều ngang**  
Mở rộng theo chiều ngang, còn được gọi là "sharding", là phương pháp mở rộng bằng cách bổ sung thêm nhiều máy chủ. Hình 1-20 so sánh giữa mở rộng theo chiều dọc và mở rộng theo chiều ngang.  
Sharding chia các cơ sở dữ liệu lớn thành những phần nhỏ hơn, dễ quản lý hơn, được gọi là các "shard". Mỗi shard có cùng cấu trúc (schema), nhưng dữ liệu thực tế trên mỗi shard là duy nhất và riêng biệt.  
Hình 1-21 đưa ra một ví dụ về cơ sở dữ liệu được phân mảnh (sharded). Dữ liệu người dùng được phân bổ vào một máy chủ cơ sở dữ liệu dựa trên ID người dùng. Mỗi khi bạn truy cập dữ liệu, một hàm băm (hash function) được sử dụng để tìm shard tương ứng. Trong ví dụ của chúng ta, hàm băm user_id % 4 được sử dụng. Nếu kết quả bằng 0, shard 0 sẽ được dùng để lưu trữ và truy xuất dữ liệu. Nếu kết quả bằng 1, shard 1 sẽ được sử dụng. Logic tương tự áp dụng cho các shard khác.

Yếu tố quan trọng nhất cần xem xét khi triển khai chiến lược phân mảnh (sharding) là việc lựa chọn **khóa phân mảnh (sharding key)**. Khóa phân mảnh (còn gọi là khóa phân vùng - partition key) bao gồm một hoặc nhiều cột quyết định cách dữ liệu được phân phối. Như được minh họa trong Hình 1-22, "user_id" là khóa phân mảnh. Khóa phân mảnh cho phép bạn truy xuất và chỉnh sửa dữ liệu một cách hiệu quả bằng cách định tuyến các truy vấn cơ sở dữ liệu đến đúng cơ sở dữ liệu. Khi chọn khóa phân mảnh, một trong những tiêu chí quan trọng nhất là chọn một khóa có khả năng phân phối dữ liệu đồng đều.

Phân mảnh là một kỹ thuật tuyệt vời để mở rộng cơ sở dữ liệu, nhưng nó không phải là một giải pháp hoàn hảo. Nó mang lại sự phức tạp và những thách thức mới cho hệ thống:

- **Tái phân mảnh dữ liệu**: Việc tái phân mảnh dữ liệu là cần thiết khi 1) một shard đơn lẻ không còn chứa được thêm dữ liệu do sự tăng trưởng nhanh chóng; 2) một số shard có thể bị cạn kiệt (shard exhaustion) nhanh hơn các shard khác do dữ liệu phân bố không đồng đều. Khi tình trạng cạn kiệt shard xảy ra, cần cập nhật hàm phân mảnh và di chuyển dữ liệu. **Hàm băm nhất quán (consistent hashing)**, sẽ được thảo luận trong Chương 5, là một kỹ thuật phổ biến để giải quyết vấn đề này.
- **Vấn đề người nổi tiếng (Celebrity problem)**: Còn được gọi là vấn đề khóa điểm nóng (hotspot key). Việc truy cập quá mức vào một shard cụ thể có thể gây quá tải máy chủ. Hãy tưởng tượng dữ liệu của Katy Perry, Justin Bieber và Lady Gaga đều nằm trên cùng một shard. Đối với các ứng dụng xã hội, shard đó sẽ bị quá tải bởi các thao tác đọc. Để giải quyết vấn đề này, chúng ta có thể cần phân bổ một shard riêng cho mỗi người nổi tiếng. Thậm chí, mỗi shard này có thể cần được phân vùng thêm.
- **Kết hợp (Join) và phi chuẩn hóa (De-normalization)**: Khi một cơ sở dữ liệu đã được phân mảnh trên nhiều máy chủ, việc thực hiện các thao tác kết hợp (join) giữa các shard trở nên rất khó khăn. Một cách giải quyết phổ biến là phi chuẩn hóa cơ sở dữ liệu để các truy vấn có thể được thực hiện trong một bảng duy nhất.

Trong Hình 1-23, chúng ta phân mảnh cơ sở dữ liệu để hỗ trợ lưu lượng dữ liệu tăng nhanh. Đồng thời, một số chức năng không liên quan đến cơ sở dữ liệu quan hệ được chuyển sang kho dữ liệu NoSQL để giảm tải cho cơ sở dữ liệu. Đây là một bài viết bao quát nhiều trường hợp sử dụng của NoSQL [14].

**Hàng triệu người dùng và hơn thế nữa**  
Việc mở rộng một hệ thống là một quá trình lặp đi lặp lại. Việc áp dụng những gì chúng ta đã học trong chương này có thể đưa chúng ta tiến xa. Tuy nhiên, để mở rộng vượt qua con số hàng triệu người dùng, cần có thêm sự tinh chỉnh và các chiến lược mới. Chẳng hạn, bạn có thể cần tối ưu hóa hệ thống và chia nhỏ hệ thống thành các dịch vụ nhỏ hơn nữa. Tất cả các kỹ thuật đã học trong chương này sẽ cung cấp một nền tảng vững chắc để đối mặt với những thách thức mới. Để kết thúc chương này, chúng tôi tóm tắt cách mở rộng hệ thống để hỗ trợ hàng triệu người dùng như sau:  
• Giữ tầng web không trạng thái (stateless).  
• Xây dựng dự phòng (redundancy) ở mọi tầng.  
• Lưu trữ dữ liệu vào bộ nhớ đệm (cache) nhiều nhất có thể.  
• Hỗ trợ nhiều trung tâm dữ liệu (data centers).  
• Lưu trữ các tài nguyên tĩnh trên CDN (Content Delivery Network).  
• Mở rộng tầng dữ liệu bằng cách phân mảnh (sharding).  
• Chia các tầng thành các dịch vụ riêng lẻ.  
• Theo dõi hệ thống của bạn và sử dụng các công cụ tự động hóa.

Chúc mừng bạn đã đi đến đây! Hãy tự vỗ vai khen ngợi bản thân nhé. Làm tốt lắm!