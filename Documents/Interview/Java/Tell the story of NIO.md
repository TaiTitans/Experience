
---
### **Ví dụ minh họa**

Giả sử một ngân hàng chỉ có 10 nhân viên. Quy trình kinh doanh của ngân hàng được chia thành 4 bước sau:

1. Khách hàng điền vào mẫu đơn (5 phút).
2. Nhân viên xem xét (1 phút).
3. Nhân viên yêu cầu bảo vệ đến kho lấy tiền (3 phút).
4. Nhân viên in phiếu và trả tiền cùng phiếu cho khách hàng (1 phút).

Hãy cùng xem xét cách các phương thức làm việc khác nhau của ngân hàng ảnh hưởng đến năng suất của họ.

#### **Cách tiếp cận BIO (Blocking I/O)**

Mỗi khi một khách hàng đến, một nhân viên sẽ ngay lập tức tiếp nhận và xử lý, và nhân viên này sẽ chịu trách nhiệm thực hiện cả 4 bước trên. Khi có hơn 10 khách hàng, những khách hàng còn lại phải xếp hàng chờ đợi.

- Một nhân viên mất 10 phút (5 + 1 + 3 + 1) để xử lý một khách hàng.
- Trong một giờ (60 phút), một nhân viên có thể xử lý 6 khách hàng, và với 10 nhân viên, ngân hàng chỉ có thể phục vụ tối đa 60 khách hàng.

Có thể thấy, trạng thái làm việc của nhân viên ngân hàng không bão hòa. Ví dụ, ở bước đầu tiên, họ thực sự chỉ đang chờ khách hàng điền đơn.  
Cách làm việc này tương tự như mô hình **BIO**: mỗi khi có một yêu cầu (khách hàng) đến, nó được giao cho một luồng (thread) trong một nhóm luồng (thread pool) để xử lý. Nếu vượt quá giới hạn tối đa của nhóm luồng (10), yêu cầu sẽ được đưa vào hàng đợi và chờ.

#### **Làm thế nào để tăng thông lượng của ngân hàng?**

Ý tưởng là: **chia để trị**, phân chia nhiệm vụ và để những người chuyên biệt xử lý các nhiệm vụ cụ thể.  
Cụ thể:

- Ngân hàng phân công một nhân viên A chỉ làm nhiệm vụ phát mẫu đơn cho khách hàng điền mỗi khi có khách đến.
- Sau khi khách hàng điền xong, A giao ngẫu nhiên đơn đó cho một trong 9 nhân viên còn lại để hoàn thành các bước tiếp theo.

Với cách này, giả sử có rất nhiều khách hàng và công việc của nhân viên A đã bão hòa, A liên tục đưa các khách hàng đã điền đơn đến quầy để xử lý:

- Một nhân viên tại quầy mất 5 phút để xử lý một khách hàng (1 + 3 + 1, vì bước 1 đã được A làm).
- Trong một giờ, 9 nhân viên có thể xử lý: 9 * (60/5) = 108 khách hàng.

Có thể thấy, việc thay đổi cách làm việc mang lại sự gia tăng hiệu quả rất lớn.  
Cách làm việc này thực chất là tư duy của **NIO (Non-blocking I/O)**.

#### **Sơ đồ NIO cổ điển**

Trong mô hình NIO:

- Một luồng (**mainReactor**) chịu trách nhiệm lắng nghe trên socket máy chủ, nhận các kết nối mới và phân phối socket đã thiết lập đến **subReactor**.
- **subReactor** (có thể là một luồng hoặc một nhóm luồng) chịu trách nhiệm tách các socket đã kết nối và đọc/ghi dữ liệu mạng. Dữ liệu mạng đọc/ghi ở đây tương tự như hành động tốn thời gian khi khách hàng điền đơn.
- Các chức năng xử lý nghiệp vụ cụ thể được giao cho nhóm luồng công việc (**worker thread pool**) để hoàn thành.

Có thể thấy, trong mô hình NIO điển hình có ba loại luồng:

- **mainReactor**: Luồng chính lắng nghe và phân phối.
- **subReactor**: Luồng phụ xử lý kết nối và dữ liệu mạng.
- **worker**: Luồng công việc xử lý nghiệp vụ.

Các luồng khác nhau làm những việc chuyên biệt, cuối cùng mỗi luồng không bị rảnh rỗi, và thông lượng của hệ thống tự nhiên tăng lên.

#### **Còn điều gì có thể cải thiện trong quá trình này không?**

Như bạn thấy, ở bước thứ ba của quy trình, nhân viên quầy yêu cầu bảo vệ đi lấy tiền từ kho (3 phút). Trong 3 phút này, nhân viên quầy chỉ chờ đợi, và họ có thể tận dụng khoảng thời gian này.  
Một lần nữa, áp dụng ý tưởng **chia để trị**:

- Phân công một nhân viên B chuyên trách bước 3.
- Khi nhân viên quầy hoàn thành bước 2, họ thông báo cho nhân viên B liên lạc với bảo vệ để lấy tiền. Lúc này, nhân viên quầy có thể chuyển sang phục vụ khách hàng tiếp theo.
- Khi nhân viên quầy quay lại phục vụ khách hàng, họ thấy 3 bước đầu của khách hàng đã hoàn thành và có thể thực hiện ngay bước 4.

Trong các dịch vụ web ngày nay, việc gọi dịch vụ bên thứ ba qua RPC hoặc HTTP thường tương ứng với bước 3. Nếu bước này mất nhiều thời gian, việc sử dụng chế độ bất đồng bộ (asynchronous mode) sẽ giảm đáng kể lãng phí tài nguyên.

#### **NIO + Bất đồng bộ**

Cách tiếp cận NIO kết hợp bất đồng bộ cho phép một số lượng nhỏ luồng thực hiện rất nhiều việc. Điều này áp dụng cho nhiều kịch bản ứng dụng, như:

- Dịch vụ proxy.
- Dịch vụ API.
- Dịch vụ kết nối liên tục (persistent connection).  
    Nếu sử dụng chế độ đồng bộ (synchronous), các ứng dụng này sẽ tiêu tốn rất nhiều tài nguyên.

Tuy nhiên, mặc dù NIO + bất đồng bộ có thể tăng thông lượng hệ thống, nó không giảm thời gian chờ của một yêu cầu mà thậm chí có thể làm tăng thời gian chờ.

#### **Tóm lại**

Ý tưởng cơ bản của NIO có thể được tóm gọn như sau: **chia để trị**, phân chia nhiệm vụ và giao cho người chuyên trách xử lý nhiệm vụ cụ thể.

### Giải thích bổ sung:

- **BIO (Blocking I/O)**: Mỗi yêu cầu được xử lý tuần tự bởi một luồng, dẫn đến lãng phí tài nguyên khi luồng phải chờ (blocking).
- **NIO (Non-blocking I/O)**: Một luồng chính xử lý nhiều kết nối, giao nhiệm vụ cho các luồng phụ hoặc worker, tối ưu hóa hiệu suất.
- **Bất đồng bộ (Asynchronous)**: Không chờ đợi kết quả ngay lập tức, cho phép xử lý song song nhiều tác vụ, đặc biệt hữu ích khi có các bước tốn thời gian (như gọi dịch vụ bên ngoài).