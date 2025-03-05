
---
# Rate Limiter là gì?

Một cách hiểu khái quát nhất, Rate Limiter giới hạn **số lượng** request truy cập một tài nguyên trên hệ thống từ một **tác nhân** (người dùng, trình duyệt, máy chủ khác...), trong một khoảng **thời gian** xác định.

Dựa trên một vài kỹ thuật rate limiter, khi vi phạm hoặc đạt một ngưỡng xác định, các requests đó sẽ bị chặn truy cập vào hệ thống.

Ba tác nhân mà mình nhấn mạnh ở trên: **số lượng, tác nhân, thời gian** cũng chính là `Core Concepts`của Rate Limiter.

# Rate Limiter với những ví dụ thực tiễn

Chắc hẳn các bạn cũng từng đối mặt hoặc chứng kiến một trong các trường hợp sau:

- Mã PIN của thẻ đã nhập sai quá 5 lần, vui lòng liên hệ với ngân hàng để mở khoá thẻ
    
- Ở một số ứng dụng, khi nhập sai mật khẩu quá nhiều lần bạn sẽ nhận thông báo, đại loại như: `Số lượng yêu cầu vượt quá mức cho phép` (như lỗi 2027 của một ứng dụng khá nổi tiếng ở VN)
    
- Thông thường IP Việt Nam dễ bị block bởi các website thương mại của nước ngoài do các vụ tấn công trước đây
    
- Hay khi yêu cầu nhận OTP quá nhiều trong 1 phút, bạn phải đợi hết thời gian quy định mới có thể yêu cầu một tin nhắn OTP mới

# Tại sao lại cần Rate Limiter

Nói là cần nhưng thực chất tuỳ thuộc vào hệ thống của bạn, nếu ở quy mô vừa và nhỏ, nghiễm nhiên có hay không một hệ thống phục vụ cho các tác vụ trên không quá quan trọng.

Nhưng khi hệ thống đạt scale lớn hơn, chắc chắn các bạn cần đặt vấn đề và quan tâm hơn dành cho Rate Limiter, như với các vấn đề sau:

- Bảo mật: kể đến đầu tiên phải là DOS và DDOS, còn đó brute force, credential stuffing attacks hay web scraping... chẳng hạn
- Cân bằng tài nguyên: Đảm bảo server không bị quá tải, phân chia tài nguyên hợp lý / công bằng cho từng user trên hệ thống
- Tiết kiệm chi phí: việc truy cập tài nguyên "có kiểm soát", góp phần giảm thiểu sự gia tăng về mặt chi phí hệ thống.

# Một số thuật toán được sử dụng trong Rate Limiter

Có thể kể đến một số thuật toán thường dùng nhất như sau:

- Leaky Bucket
- Fixed Window Counter
- Sliding Window Log
- Sliding Window Counter

---
# Leaky Bucket

Về mặt khái niệm, Leaky Bucket hoạt động như sau. Server sử dụng 1 cái xô (bucket) có kích thước cố định và đáy chứa 1 cái lỗ (output) để đựng một số lượng tokens nhất định. Token ở đây tượng trưng cho những request được gửi đến server.

Token được gởi tới server bởi user (client) và server là tác nhân cho phép token đi vào/đi ra khỏi bucket.

Để dễ hình dung hơn có thể xem hình minh hoạ bên dưới:

![](https://images.viblo.asia/b7874761-2d7a-407f-affe-6118a788b721.png)

Thuật toán Leaky Bucket không quá phức tạp để triển khai cũng như bảo trì trên một server hay load balancer, việc quản lý bộ nhớ cũng đơn giản vì mọi thứ được cấu hình từ ban đầu. Vì độ chính xác cao trong khi sử dụng tài nguyên hiệu quả, rất nhiều hệ thống lớn sử dụng Leaky Bucket để làm một Rate Limiter, tiêu biểu như [NGINX](https://www.nginx.com/blog/rate-limiting-nginx/)

Mặc dù có nhiều ưu điểm nhưng thuật toán này cũng mang một số nhược điểm sau:

1. Burstiness: Một nhược điểm của thuật toán Leaky Bucket là nó không xử lý tốt các tải trọng requests đột ngột hoặc bursty. Khi một số lượng lớn yêu cầu đến cùng một thời điểm, thuật toán sẽ phải xử lý tất cả chúng trong cùng một khoảng thời gian. Điều này có thể dẫn đến quá tải và giảm hiệu suất của hệ thống.
2. Chậm trong phản hồi: Thuật toán Leaky Bucket có thể dẫn đến việc chậm trong phản hồi đối với yêu cầu truy cập được thực hiện sau một khoảng thời gian đầy "đứt gãy" (interval), trong đó không có yêu cầu nào được thực hiện. Trong trường hợp này, việc "chảy" (leak) của các yêu cầu trong Bucket chỉ diễn ra khi interval kế tiếp bắt đầu, dẫn đến thời gian chờ lâu hơn để yêu cầu được xử lý.
3. Thiếu linh hoạt trong việc xử lý các yêu cầu cần ưu tiên hoặc khẩn cấp.
4. Lỗi triển khai thuật toán: nếu lỗi xảy ra và Bucket không được xử lý đúng cách, một số requests có thể được chấp nhận và xử lý mặc dù tỷ lệ requests đã vượt quá giới hạn được thiết lập. Điều này có thể cho phép một số lượng lớn requests tràn vào hệ thống, gây ra quá tải và có thể làm cho server trở nên không ổn định hoặc bị chậm.

---
