
---
### Key Points

- SSH, hoặc Secure Shell, là một giao thức an toàn để kết nối từ xa, được tạo ra vào năm 1995 bởi Tatu Ylonen, sử dụng cổng 22 để bảo vệ dữ liệu bằng mã hóa.
- Nghiên cứu cho thấy SSH thay thế Telnet không an toàn bằng cách sử dụng mã hóa bất đối xứng và đối xứng, đảm bảo dữ liệu được bảo vệ trong quá trình truyền tải.
- Có vẻ như SSH bắt đầu bằng một handshake TCP, sau đó trao đổi khóa và sử dụng kênh để quản lý các loại giao tiếp khác nhau.

#### Tổng quan về SSH

SSH cho phép bạn kết nối an toàn đến máy tính từ xa, giống như bạn đang ngồi trước nó. Nó được phát triển để khắc phục lỗ hổng bảo mật của Telnet, vốn gửi dữ liệu dưới dạng văn bản thuần không mã hóa. Với SSH, tất cả dữ liệu được mã hóa, bảo vệ khỏi việc bị nghe lén.

#### Quá trình Kết nối

Khi bạn kết nối qua SSH, quá trình bắt đầu bằng một handshake TCP (SYN, SYN/ACK, ACK) để thiết lập kết nối. Sau đó, máy khách và máy chủ trao đổi thông tin về phiên bản SSH và thuật toán mã hóa. Tiếp theo, họ thực hiện trao đổi khóa bằng mã hóa bất đối xứng để tạo ra một khóa bí mật chung, sau đó dùng mã hóa đối xứng để truyền dữ liệu thực tế. Dữ liệu được chia thành các gói, mỗi gói có phần đầu, kích thước đệm, nội dung chính và đệm, với tất cả trừ một số byte xác thực được mã hóa.

#### Mã hóa và Kênh

Mã hóa bất đối xứng được sử dụng trong trao đổi khóa để thiết lập khóa đối xứng, và cũng có thể dùng cho xác thực người dùng bằng khóa công khai. SSH sử dụng các kênh để tách biệt các loại giao tiếp, như một kênh cho phiên làm việc và một kênh cho xác thực.

#### Ý Nghĩa

SSH là tiêu chuẩn cho quản lý máy tính từ xa an toàn, thay thế Telnet và bảo vệ dữ liệu trong các mạng hiện đại. Một chi tiết bất ngờ là SSH không nhất thiết phải chạy trên TCP, mặc dù thường là như vậy.

---

### Báo cáo Chi tiết

SSH, hoặc Secure Shell, là một giao thức quan trọng trong quản lý máy tính từ xa, cung cấp một kênh an toàn để kết nối và truyền dữ liệu. Báo cáo này sẽ khám phá lịch sử, chức năng, và các cơ chế mã hóa của SSH, dựa trên một bài blog cụ thể và các nguồn bổ sung, nhằm cung cấp cái nhìn toàn diện cho người dùng.

#### Bối cảnh và Lịch sử

SSH được phát triển vào năm 1995 bởi Tatu Ylonen, một nhà nghiên cứu từ Phần Lan, như một giải pháp an toàn thay thế cho Telnet. Telnet, phổ biến từ những năm 1960 và vẫn được sử dụng cho các thiết bị mạng vào đầu những năm 2000, không có mã hóa, khiến dữ liệu dễ bị nghe lén qua kỹ thuật phân tích gói tin. SSH, sử dụng cổng 22 (giữa cổng 23 của Telnet và cổng 21 của FTP), đã cách mạng hóa quản lý từ xa bằng cách áp dụng mã hóa để bảo vệ dữ liệu. Bạn có thể tìm hiểu thêm về người tạo ra SSH tại [trang cá nhân của Tatu Ylonen](https://ylonen.org/index.html).

#### Chức năng và Quá trình Kết nối

SSH hoạt động trên mô hình khách-máy chủ, thường qua giao thức TCP, mặc dù không nhất thiết giới hạn ở TCP. Dưới đây là các bước chi tiết trong quá trình kết nối:

- **Handshake TCP**: Quá trình bắt đầu với một handshake ba bước tiêu chuẩn: máy khách gửi gói SYN, máy chủ trả lời bằng SYN/ACK, và máy khách xác nhận bằng ACK. Điều này thiết lập kết nối cơ bản.
- **Trao đổi Thông tin**: Sau handshake, cả hai bên trao đổi thông tin về phiên bản SSH và các thuật toán mã hóa được hỗ trợ, đảm bảo họ đồng ý về phương pháp mã hóa.
- **Trao đổi Khóa**: SSH sử dụng mã hóa bất đối xứng trong giai đoạn này. Cả máy khách và máy chủ tạo ra cặp khóa công khai và khóa riêng tư tạm thời (ephemeral), sau đó trao đổi khóa công khai. Quá trình này, thường dựa trên các thuật toán như Diffie-Hellman, cho phép họ tạo ra một khóa bí mật chung mà không cần truyền trực tiếp qua mạng. Khóa này sau đó được sử dụng cho mã hóa đối xứng trong suốt phiên làm việc.
- **Truyền Dữ liệu**: Dữ liệu được chia thành các gói, mỗi gói bao gồm:
    - **Phần đầu (Header)**: Chứa thông tin như loại gói hoặc số thứ tự.
    - **Kích thước đệm (Padding Size)**: Chỉ định lượng đệm được thêm vào.
    - **Nội dung chính (Payload)**: Dữ liệu thực tế cần truyền.
    - **Đệm (Padding)**: Dữ liệu bổ sung để đạt kích thước khối phù hợp cho mã hóa.
    - Tất cả các phần trừ một số byte xác thực đều được mã hóa bằng mã hóa đối xứng, đảm bảo tính bảo mật.

SSH cũng sử dụng các kênh để quản lý các loại giao tiếp khác nhau, chẳng hạn như một kênh cho phiên làm việc (interactive session) và một kênh cho xác thực (authentication).

#### Mã hóa Bất Đối Xứng và Ứng Dụng

Mã hóa bất đối xứng, dựa trên cặp khóa công khai và khóa riêng tư, đảm bảo rằng chỉ người nhận có khóa riêng tư mới có thể giải mã thông điệp được mã hóa bằng khóa công khai. Trong SSH, mã hóa bất đối xứng được sử dụng ở hai bối cảnh chính:

- **Trao đổi Khóa**: Để thiết lập khóa bí mật chung cho mã hóa đối xứng, sử dụng các khóa tạm thời (ephemeral) được tạo và hủy sau mỗi phiên.
- **Xác thực Người Dùng**: Nếu sử dụng xác thực bằng khóa công khai, khóa công khai của người dùng được lưu trên máy chủ, và người dùng chứng minh họ sở hữu khóa riêng tư tương ứng mà không cần tiết lộ nó. Điều này khác với các khóa ephemeral, vì khóa công khai cho xác thực thường là cố định và được sử dụng nhiều phiên.

#### Kết Luận và Tầm Quan Trọng

SSH đã thay thế Telnet không an toàn, trở thành tiêu chuẩn cho quản lý máy tính từ xa an toàn. Bằng cách sử dụng mã hóa và các cơ chế xác thực mạnh mẽ, SSH bảo vệ dữ liệu trong quá trình truyền tải, làm cho nó trở nên thiết yếu trong quản lý hệ thống và các dịch vụ mạng hiện đại. Một chi tiết thú vị là, mặc dù thường chạy trên TCP, SSH có thể được triển khai trên các giao thức khác, mở ra khả năng linh hoạt trong các kịch bản cụ thể.