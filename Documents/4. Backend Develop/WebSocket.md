# Socket Networking

Socket là khái niệm trong lập trình mạng. Sử dụng để thiết lập và quản lý kết nối giữ các ứng dụng chạy trên các máy tính khác nhau qua internet. → Socket cho phép ứng dụng giao tiếp thông qua mạng, bằng cách truyền và nhận dữ liệu qua kết nối mạng.

→ Coi socket như một đầu nối ảo giữa 2 thiết bị.

Có 2 loại socket:

- Socket TCP: Kênh kết nối có thứ tự và đáng tin cậy. Dữ liệu được chia làm các gói và đảm bảo đến đích một cách có thứ tự và không bị mất. → Thích hợp ứng dụng cần truyền dữ liệu chính xác tin cậy như game, trình duyệt web….
- Socket UDP: giao thức ko kết nối, k thứ tự, ko đảm bảo → Thích hợp cho các ứng dụng truyền dữ liệu nhanh chóng như tin nhắn, streaming ….

---
`WebSocket là một giao thức truyền thông giúp thiết lập kết nối hai chiều (full-duplex) giữa trình duyệt và máy chủ trên một kết nối TCP duy nhất. Khác với giao thức HTTP chỉ cho phép kết nối đơn hướng (client gửi yêu cầu, server trả về phản hồi), WebSocket cho phép cả client và server có thể gửi dữ liệu lẫn nhau bất kỳ lúc nào sau khi kết nối đã được thiết lập.
`
### 1. Ưu điểm của WebSocket:

- **Real-time**: WebSocket được sử dụng chủ yếu cho các ứng dụng yêu cầu giao tiếp thời gian thực như chat, trò chơi trực tuyến, hệ thống thông báo.
- **Kết nối liên tục**: Kết nối chỉ thiết lập một lần và sau đó có thể truyền dữ liệu qua lại mà không cần phải thiết lập lại kết nối cho mỗi yêu cầu như HTTP.
- **Tiết kiệm băng thông**: Giao thức WebSocket không cần phải gửi đi các thông tin dư thừa như tiêu đề HTTP trong mỗi lần truyền dữ liệu.

### 2. Cách hoạt động của WebSocket:

- **Bắt đầu bằng HTTP**: Kết nối WebSocket ban đầu được khởi tạo bởi một yêu cầu HTTP tiêu chuẩn (HTTP Handshake) từ client. Sau đó, nếu server hỗ trợ WebSocket, nó sẽ trả về phản hồi cho phép nâng cấp kết nối từ HTTP sang WebSocket.
- **Giao tiếp hai chiều**: Sau khi kết nối được nâng cấp, cả server và client có thể gửi tin nhắn cho nhau mà không cần client phải gửi request trước.

![](https://images.viblo.asia/c013ee65-3408-4a36-a234-ca8b687189bc.png)

Nói một cách đơn giản với HTTP truyền thống khi ta truyền tin đến thì con đường truyền tin cũng sẽ biến mất dẫn tới việc tập tin trả lời rất khó tìm được lại con đường cũ đó để đi còn với webSocket nó sẽ tạo ra 1 con đường mòn giao thương nhờ vậy các tập tin có thể biết đường đi đâu về đâu.

### Cấu tạo của webSocket


Về bản chất, WebSocket khác với [HTTP](https://vietnix.vn/http-https-la-gi/), mặc dù cả giao thức đều ở trên layer 7 của [mô hình OSI](https://vietnix.vn/mo-hinh-osi-la-gi/), và cùng phụ thuộc vào TCP ở layer 4. Tuy nhiên, RFC 6455 cho biết WebSocket “được thiết kế để hoạt động trên các cổng HTTP 443 và 80, cũng như để hỗ trợ proxy HTTP và làm trung gian”. Do đó, WebSocket hoàn toàn có khả năng tương thích với giao thức HTTP. Để có được sự tương thích này, WebSocket handshake sẽ sử dụng một header HTTP Upgrade để thay đổi giao thức HTTP thành WebSocket.


Giao thức WebSocket hỗ trợ tương tác giữa một trình duyệt web (hay ứng dụng client) với một web server, nhưng chi phí sẽ thấp hơn so với các lựa chọn half-duplex khác như HTTP polling. Do đó, việc truyền dữ liệu theo thời gian thực với server sẽ trở nên hiệu quả hơn.


Vậy cách hỗ trợ tương tác trên của WebSocket là gì? WebSocket sẽ cung cấp một cách được tiêu chuẩn để server có thể gửi content đến client mà không cần được request bởi client. Đồng thời cho phép các thông báo được truyền qua lại trong khi vẫn giữ cho kết nối được mở. Từ đó tạo nên một giao tiếp hai chiều giữa client và server.


Các giao tiếp này thường được xử lý qua [port](https://vietnix.vn/port-la-gi/) TCP 443 (hoặc port 80 với những kết nối không được bảo mật). Do đó, các môi trường có thể dễ dàng chặn những kết nối non-web bằng firewall. Bên cạnh đó, các giao tiếp browser-server không tiêu chuẩn cũng có thể được thiết lập bằng các công nghệ stopgap như Comet.


## Thiết lập kết nối WebSocket


Vậy cách để thiết lập kết nối WebSocket là gì? Trước hết, WebSocket Open Handshake không sử dụng lược đồ `http://` hay `https://`, vì chúng không tuân theo giao thức HTTP. Thay vào đó, URI WebSocket sẽ sử dụng giao thức `ws:` (hay `wss:` với WebSocket bảo mật). Phần còn lại của URI sẽ tương tự với HTTP URI: gồm host, port, path và bất kỳ tham số truy vấn nào.

`"ws:" "//" host [ ":" port ] path [ "?" query ]`
`"wss:" "//" host [ ":" port ] path [ "?" query ]`

Các kết nối WebSocket chỉ có thể được thiết lập cho những URI tuân theo lược đồ này. Do đó, khi thấy một URI với lược đồ `ws://` (hoặc `wss://`), cả client lẫn server đều phải tuân theo giao thức WebSocket.


Kết nối WebSocket được thiết lập bằng cách upgrading một cặp HTTP request/repsonse. Nếu một client có hỗ trợ WebSocket muốn thiết lập một kết nối, client nãy sẽ gửi một HTTP request có bao gồm các header bắt buộc sau:

- `**Connection: Upgrade**`
    - Header `Connection` thường kiểm soát việc mở kết nối mạng sau khi transaction kết thúc. Giá trị phổ biến của header này là `keep-alive`, giúp đảm bảo kết nối được liên tục cho những request sau đến cùng server. Trong WebSocket opening handshake, ta đặt header thành `Upgrade`, cho biết rằng ta muốn giữ cho kết nối luôn mở, và dùng cho những request non-HTTP.


- `**Upgrade: websocket**`
    - Header `Upgrade` được client sử dụng để yêu cầu server chuyển sang các giao thức khác trong danh sách, sắp xếp theo thứ tự ưu tiên giảm dần. Trong ví dụ này là `websocket`, cho biết rằng client muốn thiết lập một kết nối WebSocket.
- `**Sec-WebSocket-Key: q4xkcO32u266gldTuKaSOw==**`
    - `Sec-WebSocket-Key` là một giá trị ngẫu nhiên, dùng một lần (nonce), được tạo ra bởi client. Giá trị của header là một giá trị 16 byte được mã hóa base64, được tạo ngẫu nhiên
- `**Sec-WebSocket-Version: 13**`
    - Cho biết phiên bản WebSocket duy nhất được chấp nhận là `13`. Bất kỳ phiên bản nào khác được liệt kê trong header này đều không hợp lệ.
Các header trên kết hợp với nhau sẽ tạo ra một request HTTP GET từ client đến một URI ws:// như sau:

```
GET ws://example.com:8181/ HTTP/1.1 
Host: localhost:8181 
Connection: Upgrade 
Pragma: no-cache 
Cache-Control: no-cache 
Upgrade: websocket 
Sec-WebSocket-Version: 13 
Sec-WebSocket-Key: q4xkcO32u266gldTuKaSOw==
```

Sau khi client gửi request ban đầu để mở một kết nối WebSocket, nó sẽ đợi server reply. Reply từ server phải chứa một response code `HTTP 101 Switching Protocols`. Response này cho biết server đang chuyển sang giao thức mà client đã yêu cầu ở trong request header `Upgrade`. Bên cạnh đó, server phải bao gồm cả các header HTTP xác thực các kết nối được upgrade thành công:

```
HTTP/1.1 101 Switching Protocols 
Upgrade: websocket 
Connection: Upgrade 
Sec-WebSocket-Accept: fA9dggdnMPU79lJgAE3W4TRnyDM=
```


Tiếp đến, sau khi client nhận được server response, kết nối WebSocket sẽ được mở để bắt đầu việc truyền dữ liệu.

## Giao thức Handshake

Để thiết lập một giao thức WebSocket, trước tiên client sẽ gửi một WebSocket handshake request. Sau đó server sẽ trả về một WebSocket handshake response như ví dụ ở dưới.

- Client request (giống với HTTP, mỗi dòng sẽ kết thúc bằng `\r\n` và thêm một khoảng trắng bắt buộc ở cuối:
```
GET /chat HTTP/1.1 
Host: server.example.com 
Upgrade: websocket 
Connection: Upgrade 
Sec-WebSocket-Key: x3JJHMbDL1EzLkh9GBhXDw== 
Sec-WebSocket-Protocol: chat, superchat 
Sec-WebSocket-Version: 13 
Origin: http://example.com
```

```
HTTP/1.1 101 Switching Protocols 
Upgrade: websocket 
Connection: Upgrade 
Sec-WebSocket-Accept: HSmrc0sMlYUkAGmm5OPpG2HaGWk= 
Sec-WebSocket-Protocol: chat
```

Giao thức handshake sẽ bắt đầu bằng một HTTP request/response, cho phép server xử lý các kết nối HTTP cũng như kết nối WebSocket ở trên cùng một cổng. Khi kết nối đã được thiết lập, giao tiếp sẽ chuyển sang giao thức nhị phân hai chiều, không phù hợp với HTTP.


Bên cạnh các header Upgrade, client cũng sẽ gửi một header Sec-WebSocket-Key chứa các byte ngẫu nhiên đã được mã hóa base64, và server sẽ phản hồi bằng một hash của key trong header Sec-WebSocket-Accept. Việc này nhằm ngăn chặn caching proxy gửi lại giao tiếp WebSocket trước đó. Đồng thời không cung cấp bất kỳ xác thực, quyền riêng tư hay tính toàn vẹn nào. Hàm hash sẽ append chuỗi cố định (một UUID) 258EAFA5-E914-47DA-95CA-C5AB0DC85B11 vào giá trị từ header Sec-WebSocket-Key (không được giải mã từ base64). Cùng với đó là áp dụng hàm băm (hashing fucntion) SHA-1 và giải mã kết quả bằng base64.

Sau khi kết nối được thiết lập thành công, client và server có thể gửi dữ liệu WebSocket hoặc text frame qua lại trong chế độ full-duplex. Dữ liệu được frame tối thiểu, với một header nhỏ ở trước payload. Việc truyền WebSocket được mô tả như các “messages”. Ở đó, mỗi message có thể được split trên nhiều data frame. Việc này sẽ cho phép gửi các message khi có sẵn dữ liệu ban đầu, nhưng độ dài của message chưa được xác định (nó sẽ gửi từng data frame cho đến khi kết thúc với bit FIN).

Ta cũng có thể sử dụng các phần mở rộng của giao thức này để ghép (multiplex) nhiều stream đồng thời. Chẳng hạn như để tránh việc sử dụng duy nhất một socket cho payload lớn.


## Cân nhắc về bảo mật

![](https://static-xf1.vietnix.vn/wp-content/uploads/2021/06/giao-thuc-hand-shake-websocket.webp)

Khác với các cross-domain HTTP request, WebSocket request không bị giới hạn bởi chính sách Same-Origin. Do đó, các server WebSocket phải xác thực header “Origin” so với origin dự kiến trong quá trình thiết lập kết nối, nhằm tránh tấn công Cross-Site WebSocket Hijacking (một loại tấn công tương tự như Cross-Site Request Forgery – CRF). Loại tấn công này có thể xảy ra khi kết nối được xác thực bằng cookies hay HTTP. Do đó, tốt nhất là hãy sử dụng token hay những cơ chế bảo mật tương tự để xác thực kết nối WebSocket khi chuyển các dữ liệu riêng tư qua WebSocket.


## Proxy traversal

Client triển khai giao thức WebSocket muốn kiểm tra xem user agent có được cấu hình để sử dụng proxy khi kết nối đến host đích và host hay không. Nếu có thì sẽ sử dụng phương thức HTTP CONNECT để thiết lập tunnel.

Mặc dù bản thân giao thức WebSocket không hề biết về các proxy server hay firewall, nó có giao thức handshake tương thích với HTTP. Do đó sẽ cho phép các HTTP server chia sẻ các port HTTP và HTTPS của họ (80 và 443) bằng một WebSocket gateway hay server. Giao thức WebSocket xác định một tiền tố ws:// và wss:// để xác định kết nối WebSocket và WebSocket Secure. Cả hai lược đồ đều sử dụng cơ chế HTTP uprade để upgrade lên giao thức WebSocket.

Một số proxy server có thể hoạt động tốt với WebSocket, tuy nhiên một số khác cũng có thể ngăn WebSocket hoạt động bình thường. Từ đó gây ảnh hưởng đến các kết nối. Đôi khi còn có thể yêu cầu thêm một số proxy server bổ sung, hoặc yêu cầu upgrade để có thể hỗ trợ WebSocket.

Nếu lưu lượng không được mã hóa đi qua một transparent proxy không hỗ trợ WebSocket, nhiều khả năng kết nối sẽ không thể thực hiện.

Mặt khác, với những lưu lượng WebSocket được mã hóa, thì việc sử dụng [TLS](https://vietnix.vn/tls-la-gi/) giúp đảm bảo lệnh HTTP CONNECT được gửi khi trình duyệt được cấu hình để sử dụng proxy server. Việc này sẽ thiết lập một tunnel, cung cấp giao tiếp TCP low-level, end-to-end thông qua HTTP proxy, ở giữa WebSocket Secure client và WebSocket server.

Đối với transparent proxy, trình duyệt sẽ không biết về proxy server, nên HTTP CONNECT sẽ không được gửi đi. Tuy nhiên, vì traffic được mã hóa, nên các transparent proxy trung gian có thể cho phép lưu lượng được mã hóa đi qua. Do đó, nếu sử dụng WebSocket Secure thì khả năng thiết lập kết nối WebSocket sẽ cao hơn. Việc sử dụng mã hóa dù tốn chi phí tài nguyên, nhưng sẽ mang lại tỷ lệ thành công cao hơn vì nó đi qua tunnel an toàn hơn.