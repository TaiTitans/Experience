
---

![](https://statics.cdn.200lab.io/2024/01/xcaching.png.pagespeed.ic.wUSv2w_pmc.png)

Trong điện toán, **bộ nhớ đệm (cache)** là một phần cứng hoặc phần mềm lưu trữ dữ liệu tại các điểm truy cập nhanh hơn. Dữ liệu được lưu trữ trong bộ nhớ đệm có thể là kết quả của tính toán trước đó hoặc bản sao dữ liệu được lưu trữ ở nơi khác.

**Caching** là hành động thực hiện lưu trữ và quản lý dữ liệu trong cache để tối ưu hóa quá trình truy cập dữ liệu trong tương lai. Caching bao gồm việc xác định dữ liệu nào nên được lưu trữ, cũng như quản lý việc lưu trữ và cập nhật dữ liệu đó trong cache để đảm bảo dữ liệu luôn được cập nhật và phản ánh chính xác trạng thái hiện tại.

Caching có thể được áp dụng ở nhiều cấp độ khác nhau, bao gồm phần cứng (ví dụ, cache CPU), hệ điều hành, ứng dụng web và cơ sở dữ liệu. Ví dụ, trong phát triển web, caching có thể được sử dụng để lưu trữ các trang web tĩnh, kết quả truy vấn cơ sở dữ liệu, hoặc các đối tượng dữ liệu thường xuyên được truy cập để giảm thời gian tải trang và giảm tải cho máy chủ.

---
Các chiến lược caching được chia thành hai loại chính:

1. **Chiến lược đọc dữ liệu**: bao gồm **Cache aside** và **Read through**.
2. **Chiến lược ghi dữ liệu**: bao gồm **Write around**, **Write back** và **Write through**.

### Cache aside

![](https://statics.cdn.200lab.io/2024/02/Screenshot-2024-02-19-at-22.44.56.png)

**Chiến lược "Cache aside"**, còn được biết đến với tên gọi "Lazy loading" hay "Load-through cache". Mô hình này đặc biệt hữu ích trong các trường hợp mà việc cập nhật dữ liệu không quá thường xuyên, nhưng yêu cầu truy cập dữ liệu nhanh chóng.

**Cách hoạt động cơ bản của chiến lược Cache aside:**

- **Bước 1: Kiểm tra cache**

Nếu dữ liệu có trong cache, dữ liệu sẽ được trả về từ cache, giúp giảm thiểu độ trễ và tải trên hệ thống lưu trữ chính hoặc cơ sở dữ liệu. Ngược lại, nếu dữ liệu không có trong cache, hệ thống sẽ thông báo tới ứng dụng để tiếp tục với bước tiếp theo.

- **Bước 2: Truy cập dữ liệu từ nguồn chính**

Nếu dữ liệu yêu cầu không có trong cache, hệ thống sẽ truy cập dữ liệu từ nguồn dữ liệu chính, thường là một cơ sở dữ liệu hoặc một dịch vụ dữ liệu bên ngoài.

- **Bước 3: Cập nhật Cache**

Sau khi dữ liệu được truy cập từ nguồn chính, hệ thống sẽ lưu trữ (cache) dữ liệu đó cho các yêu cầu trong tương lai và trả về dữ liệu cho người dùng hoặc ứng dụng yêu cầu. Việc này đảm bảo rằng dữ liệu sẽ được nhanh chóng truy cập trong lần yêu cầu tiếp theo mà không cần phải truy cập lại nguồn dữ liệu chính.

- **Bước 4: Xử lý dữ liệu lỗi thời**

Dữ liệu trong cache cần được cập nhật hoặc loại bỏ khi nó không còn đồng bộ với dữ liệu trong nguồn chính (bị lỗi thời). Điều này thường được quản lý thông qua cơ chế hết hạn (TTL - Time to Live) hoặc bằng cách chủ động loại bỏ cache khi dữ liệu nguồn được cập nhật.

**Ưu điểm của chiến lược cache aside:**

1. **Tiết kiệm tài nguyên:** Cache-aside chỉ lưu trữ dữ liệu vào cache khi cần thiết, giảm thiểu việc sử dụng tài nguyên bộ nhớ cho dữ liệu không được truy cập thường xuyên.
2. **Tính linh hoạt:** Chiến lược này cho phép tự do quản lý việc đưa dữ liệu vào và lấy dữ liệu ra khỏi cache, giúp tối ưu hóa hiệu suất và sử dụng tài nguyên.
3. **Dễ triển khai:** Cache-aside là một phương pháp đơn giản và dễ triển khai, không đòi hỏi nhiều công sức để tích hợp vào các hệ thống phần mềm.

**Nhược điểm của chiến lược cache aside:**

1. **Hiện tượng đọc không đồng nhất:** Cần phải tự quản lý việc đồng bộ hóa dữ liệu giữa cache và nguồn dữ liệu chính. Có thể xảy ra khi dữ liệu trong cache bị cũ hoặc không được đồng bộ với nguồn gốc, dẫn đến sự không nhất quán giữa các phiên bản của dữ liệu.
2. **Rủi ro về độ trễ:** Khi dữ liệu không có sẵn trong cache, việc truy vấn từ nguồn gốc có thể tạo ra độ trễ đáng kể, ảnh hưởng đến thời gian phản hồi của hệ thống.

**Trường hợp sử dụng:** Web caching cho các trang web hoặc dịch vụ API có lượng truy cập cao nhưng dữ liệu không thay đổi thường xuyên.

**Ví dụ:** Một trang web tin tức có thể sử dụng cache aside để lưu trữ các bài báo được truy cập nhiều, giảm tải cho cơ sở dữ liệu bằng cách tránh truy vấn lặp lại cho cùng một nội dung.

---
### Read through


![](https://statics.cdn.200lab.io/2024/02/Screenshot-2024-02-19-at-22.51.05.png)

Chiến lược **Read through** là một mô hình caching được thiết kế để _tự động hóa quá trình tải và lưu trữ dữ liệu vào cache._ Khác với chiến lược "Cache Aside", nơi ứng dụng phải chủ động kiểm tra cache, sau đó tải dữ liệu vào cache nếu nó không tồn tại, chiến lược Read through sử dụng một lớp trung gian (thường là một _cache proxy_ hoặc _cache library_) để tự động xử lý việc tải dữ liệu vào cache.

**Các bước hoạt động của chiến lược read through bao gồm:**

- **Bước 1: Yêu cầu dữ liệu**

Khi một ứng dụng cần truy cập dữ liệu, nó sẽ yêu cầu dữ liệu thông qua lớp cache thay vì trực tiếp từ nguồn dữ liệu chính (ví dụ: cơ sở dữ liệu).

- **Bước 2: Kiểm tra cache**

Nếu dữ liệu có sẵn trong cache, cache sẽ trả về dữ liệu ngay lập tức cho ứng dụng, giảm thiểu độ trễ và tải trên nguồn dữ liệu chính. Ngược lại, nếu dữ liệu không có trong cache,  Quá trình tiếp theo sẽ được kích hoạt.

- **Bước 3: Tải dữ liệu từ nguồn chính**

Nếu dữ liệu yêu cầu không tồn tại trong cache, lớp trung gian sẽ tự động truy cập nguồn dữ liệu chính để lấy dữ liệu.

- **Bước 4: Cập nhật cache**

Sau khi dữ liệu được lấy từ nguồn chính, nó sẽ được tự động lưu vào cache. Điều này đảm bảo rằng trong lần truy cập tiếp theo, dữ liệu có thể được trả về ngay lập tức từ cache mà không cần truy cập lại nguồn chính.

- **Bước 5: Trả về dữ liệu**

Dữ liệu sau khi đã được lưu vào cache sẽ được trả về cho ứng dụng, giống như nó được trả về từ cache trong trường hợp dữ liệu đã sẵn có.


**Ưu điểm của chiến lược Read through:**

1. **Tự động cập nhật cache:** Hệ thống tự động đọc dữ liệu từ nguồn gốc và cập nhật vào cache mỗi khi cần thiết, giúp đảm bảo tính nhất quán của dữ liệu trong cache.
2. **Giảm thiểu lỗi thời của dữ liệu:** Bằng cách tự động cập nhật cache, chiến lược này giúp giảm thiểu tình trạng dữ liệu lỗi thời trong cache.
3. **Giảm công sức quản lý:** Không cần phải quản lý việc cập nhật cache một cách thủ công.

**Nhược điểm của chiến lược Read through:**

1. **Cần có lớp trung gian:** Cần có lớp trung gian hoặc thư viện hỗ trợ, có thể làm tăng độ phức tạp của hệ thống.

**Trường hợp sử dụng:** Ứng dụng cần tự động hóa việc tải dữ liệu vào cache, đảm bảo dữ liệu luôn sẵn có trong cache khi cần.

**Ví dụ:** Một hệ thống quản lý hàng tồn kho tự động cập nhật cache khi có yêu cầu dữ liệu về số lượng hàng, đảm bảo thông tin được truy cập nhanh chóng mà không cần truy vấn cơ sở dữ liệu mỗi lần.


---
### Write around

![](https://statics.cdn.200lab.io/2024/02/Screenshot-2024-02-19-at-22.44.45.png)
Trong chiến lược Write around, khi dữ liệu mới được viết, nó được ghi trực tiếp vào nguồn dữ liệu chính mà không cập nhật dữ liệu đó vào cache ngay lập tức.

- **Bước 1: Yêu cầu ghi dữ liệu**

Dữ liệu mới hoặc cập nhật được ghi trực tiếp vào nguồn dữ liệu chính, bỏ qua việc ghi vào cache.

- **Bước 2: Bỏ qua cache**

Không cập nhật cache ngay sau khi ghi, giúp giảm tải trên bộ nhớ cache và tăng hiệu suất ghi dữ liệu.

- **Bước 3: Yêu cầu đọc dữ liệu**

Khi có yêu cầu đọc dữ liệu, hệ thống kiểm tra trong cache. Nếu không có, dữ liệu mới từ nguồn chính sẽ được tải vào cache.

- **Bước 4: Cập nhật cache**

Dữ liệu được cập nhật vào cache từ nguồn chính, sẵn sàng cho các yêu cầu đọc tiếp theo.

**Ưu điểm của chiến lược Write around:**

- **Giảm độ trễ ghi:** Giảm bớt việc ghi dữ liệu không cần thiết vào cache bằng cách loại bỏ bước đồng bộ hóa dữ liệu giữa cache và nguồn dữ liệu chính, tiết kiệm tài nguyên cache cho dữ liệu thường xuyên được truy cập.

**Nhược điểm của chiến lược Write around:**

- **Rủi ro về độ trễ đọc:** Có thể gây ra độ trễ khi đọc dữ liệu lần đầu tiên sau khi nó được ghi, do dữ liệu phải được tải từ nguồn dữ liệu chính vào cache.

**Trường hợp sử dụng:** Hệ thống cần tối ưu hiệu suất ghi, giảm bớt việc sử dụng cache cho dữ liệu không thường xuyên được truy cập.

**Ví dụ:** Một hệ thống lưu trữ dữ liệu lớn, nơi dữ liệu mới được ghi vào cơ sở dữ liệu nhưng chỉ được cache khi có yêu cầu đọc, nhằm tối ưu hóa không gian cache và hiệu suất.

---

### Write back

![](https://statics.cdn.200lab.io/2024/02/Screenshot-2024-02-19-at-22.44.29.png)
Chiến lược **Write back** (còn được gọi là **Write behind**) là chiến lược caching mà dữ liệu được viết vào cache trước và sau đó được đồng bộ hóa với nguồn dữ liệu chính sau một khoảng thời gian nhất định hoặc dựa trên một sự kiện nhất định.

**Các bước hoạt động của chiến lược Write back bao gồm:**

- **Bước 1: Yêu cầu ghi dữ liệu**

Khi có yêu cầu ghi dữ liệu, dữ liệu mới hoặc được cập nhật được ghi vào cache trước.

- **Bước 2: Đánh dấu dữ liệu**

Dữ liệu trong cache được đánh dấu là "đã thay đổi" hoặc "dirty", để biểu thị rằng nó chưa được đồng bộ hóa với nguồn dữ liệu chính.

- **Bước 3: Đồng bộ hóa dữ liệu**

Dữ liệu được đồng bộ hóa với nguồn dữ liệu chính dựa trên một chính sách định kỳ hoặc sự kiện cụ thể (ví dụ, cache đầy hoặc sau một khoảng thời gian nhất định).

- **Bước 4: Xác nhận đồng bộ**

Sau khi dữ liệu được đồng bộ hóa thành công với nguồn dữ liệu chính, dấu "dirty" được gỡ bỏ.

**Ưu điểm của chiến lược Write back:**

- **Giảm độ trễ ghi:** Cải thiện đáng kể hiệu suất ghi bằng cách giảm độ trễ, vì việc ghi vào cache nhanh hơn nhiều so với ghi vào nguồn dữ liệu chính.
- **Tối ưu hóa tài nguyên:** Giảm tải trên nguồn dữ liệu chính, cho phép hệ thống xử lý hiệu quả hơn các yêu cầu khác.

**Nhược điểm của chiến lược Write back:**

- **Rủi ro mất dữ liệu:** Nếu hệ thống gặp sự cố trước khi dữ liệu được đồng bộ hóa, có nguy cơ mất dữ liệu chưa được lưu trữ vào nguồn dữ liệu chính.
- **Quản lý phức tạp:** Yêu cầu cơ chế quản lý cache và đồng bộ hóa dữ liệu hiệu quả để đảm bảo tính nhất quán và độ tin cậy của dữ liệu.

**Trường hợp sử dụng:** Ứng dụng đòi hỏi hiệu suất ghi cao và có thể chấp nhận được độ trễ trong việc đồng bộ hóa dữ liệu với nguồn dữ liệu chính.

**Ví dụ:** Hệ thống ghi nhật ký (logging) nơi dữ liệu được ghi nhanh chóng vào cache và sau đó đồng bộ hóa với cơ sở dữ liệu log một cách định kỳ, giảm độ trễ ghi và tối ưu hóa hiệu suất.


---
###  Write through


![](https://statics.cdn.200lab.io/2024/02/Screenshot-2024-02-19-at-22.50.31.png)
Chiến lược Write Through là một kỹ thuật quản lý cache trong đó dữ liệu được ghi vào cả cache và nguồn dữ liệu chính đồng thời.

**Các bước hoạt động của chiến lược Write through bao gồm:**

- **Bước 1: Yêu cầu ghi dữ liệu**

Khi có yêu cầu ghi dữ liệu, hệ thống sẽ thực hiện ghi dữ liệu vào cả cache và nguồn dữ liệu chính. Quá trình này đảm bảo rằng dữ liệu trong cache luôn đồng bộ với nguồn dữ liệu chính.

- **Bước 2: Ghi vào cache**

Dữ liệu mới hoặc được cập nhật ghi vào cache. Điều này giúp đảm bảo rằng các yêu cầu đọc tiếp theo cho dữ liệu đó sẽ truy cập nhanh chóng từ cache mà không cần phải truy vấn nguồn dữ liệu chính.

- **Bước 3: Ghi vào nguồn dữ liệu chính**

Đồng thời, dữ liệu cũng được ghi vào nguồn dữ liệu chính. Điều này đảm bảo tính nhất quán và độ tin cậy của dữ liệu, vì mọi thay đổi đều được phản ánh ngay lập tức trong cả hai vị trí.

**Trường hợp sử dụng:** Các ứng dụng cần đảm bảo dữ liệu luôn được cập nhật và đồng bộ giữa cache và nguồn dữ liệu chính một cách ngay lập tức.

**Ví dụ:** Hệ thống xử lý giao dịch tài chính, nơi cần đảm bảo dữ liệu giao dịch được cập nhật tức thì vào cả cache và cơ sở dữ liệu để đảm bảo tính nhất quán và độ tin cậy của dữ liệu.


---
##  Kết luận

- **Bộ nhớ đệm (cache)** là một phần cứng hoặc phần mềm lưu trữ dữ liệu tại các điểm truy cập nhanh hơn.
- Caching là quá trình lưu trữ tạm thời dữ liệu vào cache để tăng tốc độ truy cập dữ liệu sau này.
- Có 5 chiến lược caching phổ biến chia làm hai loại đọc và ghi, bao gồm: cache aside, read through, write around, write back và write through. Mỗi chiến lược có những ưu và nhược điểm riêng, việc lựa chọn chiến lược phù hợp phụ thuộc vào yêu cầu cụ thể của hệ thống và ứng dụng