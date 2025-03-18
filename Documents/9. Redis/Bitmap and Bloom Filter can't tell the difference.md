
---

### Bitmap

Bitmap trong Redis là một loại dữ liệu đặc biệt lưu trữ dữ liệu ở đơn vị nhỏ nhất. Chúng ta biết rằng một byte bao gồm 8 bit. So với các cấu trúc dữ liệu truyền thống sử dụng byte, bitmap cực kỳ tiết kiệm không gian khi xử lý một lượng lớn trạng thái nhị phân (đúng, sai, hoặc 0, 1, v.v.). Tuy nhiên, đây không hoàn toàn là một loại dữ liệu mới, và cách triển khai bên dưới vẫn dựa trên kiểu bitString.

Để dễ hiểu hơn, bạn có thể hình dung cấu trúc cơ bản của bitmap như một mảng các bit, trong đó mỗi bit tương ứng với một offset (giống như chỉ số của mảng). Các trạng thái khác nhau được biểu thị bằng cách đặt giá trị bit tại một offset cụ thể thành 0 hoặc 1.

#### Ví dụ

Giả sử chúng ta muốn thiết kế một hệ thống trả lời câu hỏi với quy tắc như sau: Nếu người dùng trả lời đúng tất cả 7 câu hỏi, họ sẽ giành được giải thưởng lớn.

Mỗi người dùng có một khóa riêng, ví dụ: answer:user:X. Chúng ta sử dụng bitmap để ghi lại câu trả lời của người dùng, đặt số thứ tự câu hỏi làm offset tương ứng. Khi người dùng trả lời đúng một câu hỏi ✅, offset tương ứng được đặt thành 1; khi trả lời sai ❌, offset được đặt thành 0.

Nếu người dùng trả lời đúng các câu hỏi 2, 5 và 7, các offset tương ứng là 2, 5 và 7 sẽ được đặt thành 1, còn các vị trí khác mặc định là 0. Để xem người dùng đã trả lời câu hỏi như thế nào, chỉ cần duyệt qua cấu trúc dữ liệu theo offset. Nếu gặp giá trị 1 tại một offset, điều đó có nghĩa là câu hỏi tương ứng đã được trả lời đúng. Ví dụ với khóa answer:user:1.

#### Đếm người chiến thắng

Sau khi cuộc thi kết thúc, bước tiếp theo là đếm số người chiến thắng, tức là những người trả lời đúng tất cả 7 câu hỏi.

Để nhanh chóng kiểm tra xem một người dùng có trả lời đúng tất cả các câu hỏi hay không, bạn có thể sử dụng lệnh BITCOUNT để đếm số vị trí được đặt thành 1. Nếu kết quả bằng 7, điều đó có nghĩa là người dùng đã trả lời đúng toàn bộ câu hỏi. Cụ thể:  
BITCOUNT answer:user:X == 7

#### Ứng dụng với danh sách đen email

Nếu bạn muốn sử dụng bitmap để xác định liệu một địa chỉ email có nằm trong danh sách đen hay không, làm thế nào để đặt offset? Thật không may, bitmap không hỗ trợ trực tiếp việc sử dụng chuỗi ký tự làm offset. Tuy nhiên, chúng ta có thể băm (hash) địa chỉ email để lấy giá trị băm, từ đó tính toán offset.

Khi sử dụng phương pháp băm, va chạm dữ liệu (data collision) là điều không thể tránh khỏi, nghĩa là các chuỗi khác nhau có thể cho ra cùng một giá trị băm. Kết quả là việc đánh giá trạng thái có thể không chính xác. Đừng lo lắng, chúng ta sẽ thảo luận về Bộ lọc Bloom (Bloom Filter) sau đây để xem cách nó có thể tối ưu hóa vấn đề này.

### Các lệnh thao tác

Bitmap có ít lệnh và dễ sử dụng, chủ yếu bao gồm các lệnh sau: `SETBIT`, `GETBIT`, `BITCOUNT`, và `BITOP`.

#### `SETBIT`
Lệnh này cho phép bạn đặt giá trị bit tại một offset được chỉ định, với độ phức tạp thời gian là O(1). Ví dụ, khi một người dùng trả lời đúng câu hỏi số 7, bạn có thể đặt giá trị bit tại offset 7 thành 1 để biểu thị rằng câu hỏi đó đã được trả lời đúng.

```bash
# Khóa của người dùng: answer:user:1
# Offset: số câu hỏi 7
# Câu hỏi được trả lời đúng, đặt thành 1
SETBIT answer:user:1 7 1
```

#### `GETBIT`
Lệnh này dùng để lấy giá trị bit tại offset được chỉ định, cũng có độ phức tạp thời gian hiệu quả. Bạn có thể nhanh chóng truy vấn trạng thái trả lời của người dùng cho một câu hỏi cụ thể.

```bash
# Truy vấn trạng thái trả lời của người dùng cho câu hỏi số 7: 1 - đúng, 0 - sai
GETBIT answer:user:1 7
```

#### `BITCOUNT`
Lệnh này được sử dụng để đếm số lượng giá trị bit được đặt thành 1. Ví dụ, bạn có thể dễ dàng đếm số người dùng trả lời đúng tất cả các câu hỏi. Tuy nhiên, lệnh này chỉ cho biết tổng số câu hỏi được trả lời đúng, không thể xác định cụ thể đó là câu hỏi nào.

```bash
# Thống kê số lượng câu hỏi mà người dùng trả lời đúng (giá trị bit = 1)
BITCOUNT answer:user:1
```

#### `BITOP`
Lệnh này thực hiện các phép toán bit trên một hoặc nhiều bitmap và lưu kết quả vào một khóa mới. Nó hỗ trợ bốn phép toán: `AND` (và), `OR` (hoặc), `NOT` (phủ định), và `XOR` (hoặc độc quyền). Lệnh này được dùng để tính toán giá trị bit tại cùng offset trên nhiều bitmap. Chẳng hạn, nếu bạn muốn biết câu hỏi mà cả người dùng 1 và người dùng 2 đều trả lời đúng, sau khi thực hiện phép toán `AND`, nếu chỉ có câu hỏi số 1 là câu mà cả hai đều trả lời đúng, thì trong tập kết quả mới, chỉ offset tương ứng với câu hỏi số 1 sẽ có giá trị 1.

```bash
# Các câu hỏi mà cả người dùng 1 và người dùng 2 đều trả lời đúng, chỉ có câu hỏi số 1 được cả hai trả lời đúng
SETBIT answer:user:1 1 1
SETBIT answer:user:1 2 0
SETBIT answer:user:1 3 1

SETBIT answer:user:2 1 1
SETBIT answer:user:2 2 1
SETBIT answer:user:2 3 0

BITOP AND resultbitmap answer:user:1 answer:user:2
```

### Phát huy điểm mạnh và tránh điểm yếu

#### **Ưu điểm**
1. **Tiết kiệm không gian cực kỳ hiệu quả**: Bitmap thực sự giúp tiết kiệm không gian lưu trữ dữ liệu. Theo tính toán sơ bộ, một bitmap với 100 triệu bit chỉ chiếm khoảng 12MB bộ nhớ, giúp tiết kiệm đáng kể không gian lưu trữ so với các cấu trúc dữ liệu khác.
2. **Truy vấn nhanh**: Các thao tác trên bit trong bitmap thường nhanh hơn so với truy vấn trên các cấu trúc dữ liệu khác. Dù là đặt giá trị bit hay lấy giá trị bit, độ phức tạp thời gian đều là O(1), cho phép phản hồi nhanh chóng các yêu cầu truy vấn.
3. **Dễ vận hành**: Hỗ trợ thao tác trên từng bit, thống kê bit, và các phép toán logic trên bit, với hiệu suất cao, không cần thực hiện so sánh hay dịch chuyển phức tạp.

#### **Nhược điểm**
1. Do đặc tính của cấu trúc dữ liệu, bitmap chỉ phù hợp để biểu thị hai trạng thái: 0 và 1. Trong các trường hợp cần biểu thị nhiều trạng thái hơn, Bitmap không phải là lựa chọn thích hợp.
2. Nếu chỉ đặt giá trị bit tại ba offset như (20, 30, 888888888), chúng ta cần tạo một Bitmap có độ dài 99999999, nhưng thực tế chỉ lưu trữ 3 dữ liệu. Điều này dẫn đến lãng phí không gian lớn. Trong trường hợp gặp vấn đề này, bạn có thể sử dụng Roaring Bitmap để giải quyết.

#### **Các kịch bản ứng dụng**
Bitmap là một cấu trúc dữ liệu tương đối đơn giản, chiếm ít không gian và có hiệu suất truy vấn cao, rất phù hợp cho các kịch bản ghi lại trạng thái. Các ứng dụng của nó khá phổ biến, chẳng hạn như:
- **Trạng thái điểm danh của người dùng**: Theo dõi số ngày điểm danh liên tục.
- **Trạng thái trực tuyến của người dùng**: Thống kê số lượng người dùng hoạt động.
- **Trả lời bảng câu hỏi**: Ghi lại và đánh giá kết quả trả lời.

---
### Bộ lọc Bloom (Bloom Filter)

Như đã đề cập ở trên, khi Bitmap ghi lại trạng thái của một phần tử dạng ký tự (ví dụ: địa chỉ email), bạn cần sử dụng hàm băm (hashing) để lấy offset trước. Tuy nhiên, sau khi áp dụng thao tác băm, có thể xảy ra va chạm băm (hash collision), dẫn đến việc đánh giá trạng thái sai lệch.

Bộ lọc Bloom nâng vấn đề này lên một tầm cao mới và cho phép kiểm soát tỷ lệ dương tính giả (false positive rate). Khi chúng ta thêm một địa chỉ email vào tập hợp, một số hàm băm khác nhau sẽ ánh xạ địa chỉ email đó đến nhiều vị trí offset khác nhau trong bitmap, và đặt giá trị của các bit tại những vị trí này thành 1.

Để xác định xem một địa chỉ email có nằm trong tập hợp hay không, ta sử dụng cùng các hàm băm để ánh xạ nó đến nhiều vị trí trên bitmap. Nếu tất cả giá trị bit tại các vị trí này đều là 1, thì địa chỉ email *có thể* nằm trong tập hợp; nếu bất kỳ vị trí nào có giá trị 0, thì chắc chắn phần tử đó *không* nằm trong tập hợp. Đây là một đặc trưng nổi bật của Bộ lọc Bloom.

Mặc dù Bộ lọc Bloom vẫn tồn tại khả năng xảy ra dương tính giả, nhưng may mắn thay, chúng ta có thể kiểm soát tỷ lệ dương tính giả này bằng cách điều chỉnh kích thước của Bộ lọc Bloom và số lượng hàm băm được sử dụng.


### Các lệnh thao tác

Bộ lọc Bloom (Bloom Filter) không có quá nhiều lệnh, và các lệnh chính bao gồm như sau:

#### `BF.RESERVE`
Lệnh này dùng để tạo một bộ lọc Bloom mới và chỉ định dung lượng (`capacity`) cùng tỷ lệ dương tính giả (`error_rate`).

Cú pháp:  
`BF.RESERVE <key> <error_rate> <capacity>`  

Ví dụ:  
```bash
BF.RESERVE myfilter 0.000001 999999
```
Lệnh trên tạo một bộ lọc Bloom có tên `myfilter` với tỷ lệ dương tính giả là 0.000001 và dung lượng tối đa là 999999 phần tử.

#### `BF.INFO`
Lệnh này cung cấp thông tin về bộ lọc Bloom, bao gồm dung lượng, tỷ lệ dương tính giả, v.v.

Cú pháp:  
`BF.INFO <key>`  

Ví dụ:  
```bash
BF.INFO myfilter
```

#### `BF.ADD` và `BF.MADD`
- `BF.ADD`: Thêm một phần tử vào bộ lọc Bloom.  
- `BF.MADD`: Thêm nhiều phần tử vào bộ lọc Bloom cùng lúc (thêm hàng loạt).

Ví dụ:  
```bash
# Thêm một phần tử vào bộ lọc Bloom
BF.ADD myfilter hello
```

Cú pháp `BF.MADD`:  
`BF.MADD <key> <item> [item ...]`  

Ví dụ:  
```bash
BF.MADD myfilter item1 item2 item3
```

#### `BF.EXISTS` và `BF.MEXISTS`
- `BF.EXISTS`: Kiểm tra xem một phần tử có tồn tại trong bộ lọc Bloom hay không.  
- `BF.MEXISTS`: Kiểm tra sự tồn tại của nhiều phần tử trong bộ lọc Bloom cùng lúc.

Ví dụ:  
```bash
# Kiểm tra xem phần tử "hello" có trong bộ lọc Bloom hay không
BF.EXISTS myfilter hello
```

Cú pháp `BF.MEXISTS`:  
`BF.MEXISTS <key> <item> [item ...]`  

Ví dụ:  
```bash
BF.MEXISTS myfilter item1 item2 item3
```


### Phát huy điểm mạnh và tránh điểm yếu

#### **Ưu điểm**
1. **Chiếm ít không gian**: Bộ lọc Bloom sử dụng rất ít không gian lưu trữ. Nó không lưu trữ toàn bộ dữ liệu mà chỉ dùng các bit ở tầng cơ sở (tương tự như Bitmap) để biểu thị sự tồn tại của dữ liệu.
2. **Hiệu suất ổn định**: Hiệu suất của Bộ lọc Bloom khá ổn định bất kể số lượng phần tử trong tập hợp là bao nhiêu. Độ phức tạp thời gian của các thao tác chèn và truy vấn rất thấp, thường là O(k), trong đó k là số lượng hàm băm. Nói cách khác, khi xử lý dữ liệu quy mô lớn, hiệu suất của Bộ lọc Bloom không giảm mạnh khi lượng dữ liệu tăng lên.

#### **Nhược điểm**
1. **Tồn tại tỷ lệ nhận diện sai**: Bộ lọc Bloom có khả năng xảy ra dương tính giả (false positives), tức là khi một phần tử thực tế không nằm trong tập hợp, nó vẫn có thể bị đánh giá là có mặt. Điều này xảy ra do nhiều phần tử có thể được ánh xạ đến cùng một vị trí thông qua hàm băm, dẫn đến kết quả sai lệch. Tuy nhiên, khi Bộ lọc Bloom xác định một phần tử không nằm trong tập hợp, kết quả này chính xác 100%.
2. **Khó xóa phần tử**: Thông thường, bạn không thể xóa trực tiếp các phần tử khỏi Bộ lọc Bloom. Lý do là một vị trí bit có thể được ánh xạ bởi nhiều phần tử, và nếu đặt giá trị bit đó thành 0, điều này có thể ảnh hưởng đến việc đánh giá các phần tử khác.

#### **Các kịch bản ứng dụng**
Do Bộ lọc Bloom có một số dương tính giả, các kịch bản sử dụng nó phải chấp nhận được mức độ không chính xác nhất định:

1. **Khắc phục vấn đề xuyên thấu bộ đệm Redis**: Trong các chương trình giảm giá chớp nhoáng, danh sách sản phẩm thường được lưu trong bộ đệm Redis. Nếu có một lượng lớn yêu cầu độc hại truy vấn các sản phẩm không tồn tại, Bộ lọc Bloom có thể nhanh chóng xác định những sản phẩm này không tồn tại, từ đó tránh truy vấn đến cơ sở dữ liệu và giảm áp lực lên hệ thống.
2. **Lọc danh sách đen email**: Trong hệ thống thư điện tử, Bộ lọc Bloom có thể được dùng để lọc thư rác và thư độc hại. Địa chỉ hoặc đặc điểm của những kẻ gửi thư rác đã biết được lưu trong Bộ lọc Bloom, giúp xác định xem người gửi thư mới có nằm trong danh sách đen hay không.
3. **Lọc URL của trình thu thập dữ liệu**: Trong các bot thu thập dữ liệu, Bộ lọc Bloom có thể ghi lại các URL đã được thu thập để tránh lặp lại việc thu thập cùng một URL. Khi gặp URL mới, hệ thống sẽ kiểm tra xem URL đó đã được thu thập trước đó hay chưa.
4. **Và còn nhiều ứng dụng khác nữa...**


---
### Tóm tắt

Bài viết đã sắp xếp và trình bày các nguyên lý của **Bitmap** và **Bộ lọc Bloom**, cách sử dụng, cùng với ưu điểm, nhược điểm và các kịch bản ứng dụng của chúng. Trong bối cảnh môi trường chung không thuận lợi, việc nâng cao kỹ năng kỹ thuật là cần thiết. Hiện nay, các cuộc phỏng vấn thường không thể tránh khỏi các câu hỏi liên quan đến dữ liệu lớn và xử lý đồng thời cao (high concurrency). Để trả lời tự tin những câu hỏi này, bạn cần không chỉ chiều sâu mà còn cả chiều rộng kiến thức. Việc nắm vững hai kiến thức này (Bitmap và Bloom Filter) sẽ giúp bạn có thêm một lựa chọn trả lời hiệu quả. Dù bài viết có thể chưa hoàn hảo, nhưng hãy tận dụng và xử lý nó một cách linh hoạt nhé!