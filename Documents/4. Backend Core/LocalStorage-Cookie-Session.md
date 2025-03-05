# Local Storage & Cookie & Session

1. Local Storage
   - Khả năng lưu trữ vô thời hạn - chỉ bị xoá bằng JS, hoặc xoá bộ nhớ trình duyệt, hoặc xoá bằng localStorage API.
   - Lưu trữ được 5MB
   - Không gửi thông tin lên server như Cookie nên bảo mật tốt hơn.
   - Data type: string
   
2. Session 

   - Lưu trên Client: Cũng giống LS thì SS cũng dùng để lưu trữ dữ liệu trên trình duyệt khách (client)

   - Mất dữ liệu khi đóng tab.
   - Dữ liệu ko được gửi lên Server
   - Lữu trữ nhiều hơn cookie(>5MB)

3. Cookie
   - Thông tin được gửi lên server -> truyền từ server tới browser và được lưu trữ trên client -> mỗi khi người dùng tải ứng dụng -> trình duyệt sẽ gửi cookie để thông báo cho ứng dụng về hoạt động trước đó của user. -> bảo mật không cao
   - Cookie chủ yếu là để đọc phía máy chủ cũng có thể được đọc ở client.
   - Có thời gian sống: mỗi cookie sẽ được coder set timeout
   - Lưu trữ tối đa 4KB và vài chục cookie cho 1 domain.
   - Có thể lưu nhiều datatype

PS:

 - Chú ý đến vấn đề bảo mật nên sẽ sử dụng 1 cách hiệu quả.
 - LS và SS lữu trữ trên client nên sẽ không ảnh hưởng đến hiệu xuất của client nhưng sẽ làm nặng trình duyệt của client (ko đáng kể).
 - Về phạm vi: SS giới hạn trong cửa sổ của browser. Nếu 1 browser mở 2 tab cùng 1 trang web sẽ không thể truy xuất dữ liệu lẫn nhau. Còn LS thì có thể truy xuất lẫn nhau.



Cách sử dụng: [Local Storage, Session Storage và Cookie - Viblo](https://viblo.asia/p/local-storage-session-storage-va-cookie-ORNZqN3bl0n)