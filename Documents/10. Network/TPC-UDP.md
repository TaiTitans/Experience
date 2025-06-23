
----
### 1. TCP (Transmission Control Protocol) - Giao thức điều khiển truyền vận
#### Khái niệm:
TCP là một giao thức tầng giao vận (Transport Layer) trong mô hình OSI/TCP-IP. Nó đảm bảo việc truyền dữ liệu giữa hai thiết bị (client và server) diễn ra **đáng tin cậy**, **theo thứ tự**, và **không bị mất mát**.

#### Đặc điểm chính:
- **Kết nối hướng (Connection-Oriented):** Trước khi truyền dữ liệu, TCP thiết lập một kết nối giữa hai thiết bị thông qua cơ chế "bắt tay 3 bước" (Three-Way Handshake).
- **Đáng tin cậy (Reliable):** TCP sử dụng cơ chế kiểm tra lỗi, gửi lại gói tin bị mất (retransmission), và đảm bảo dữ liệu đến đúng thứ tự.
- **Kiểm soát luồng (Flow Control):** Sử dụng cơ chế "sliding window" để tránh tình trạng gửi dữ liệu quá nhanh khiến bên nhận không xử lý kịp.
- **Kiểm soát tắc nghẽn (Congestion Control):** Điều chỉnh tốc độ gửi dữ liệu để tránh làm nghẽn mạng.
- **Truyền dữ liệu dạng luồng (Stream):** Dữ liệu được gửi dưới dạng một luồng byte liên tục, không có ranh giới rõ ràng giữa các gói tin.

#### Cơ chế hoạt động:
1. **Three-Way Handshake:**
   - Client gửi gói SYN (synchronize) đến server.
   - Server trả lời bằng gói SYN-ACK (synchronize-acknowledge).
   - Client gửi gói ACK (acknowledge) để xác nhận. Kết nối được thiết lập.
2. **Truyền dữ liệu:** Dữ liệu được chia thành các đoạn (segments), mỗi đoạn có số thứ tự (sequence number) để đảm bảo thứ tự và kiểm tra mất mát.
3. **Đóng kết nối:** Sử dụng "Four-Way Handshake" (FIN, ACK) để kết thúc.

#### Ưu điểm:
- Đảm bảo dữ liệu được gửi đến chính xác và đầy đủ.
- Phù hợp với các ứng dụng yêu cầu độ tin cậy cao như HTTP/HTTPS, FTP, SMTP (email).

#### Nhược điểm:
- Tốn tài nguyên hơn do phải thiết lập và duy trì kết nối.
- Độ trễ cao hơn vì phải xử lý kiểm tra lỗi và gửi lại.

#### Ứng dụng thực tế:
- Web (HTTP/HTTPS).
- Truyền file (FTP).
- Email (SMTP, IMAP).

---

### 2. UDP (User Datagram Protocol) - Giao thức dữ liệu người dùng
#### Khái niệm:
UDP cũng là một giao thức tầng giao vận, nhưng khác với TCP, nó **không thiết lập kết nối** và **không đảm bảo độ tin cậy**. Nó đơn giản hơn, nhanh hơn, nhưng dữ liệu có thể bị mất hoặc đến không đúng thứ tự.

#### Đặc điểm chính:
- **Không kết nối (Connectionless):** Không cần thiết lập kết nối trước khi gửi dữ liệu, chỉ cần gửi gói tin (datagram) đi.
- **Không đáng tin cậy (Unreliable):** Không có cơ chế kiểm tra lỗi hay gửi lại gói tin bị mất.
- **Nhanh và nhẹ:** Do không có overhead từ việc quản lý kết nối hay kiểm soát luồng.
- **Truyền dữ liệu dạng gói (Datagram):** Mỗi gói tin là một đơn vị độc lập, có ranh giới rõ ràng.
- **Không kiểm soát tắc nghẽn:** Gửi dữ liệu với tốc độ tối đa, có thể gây nghẽn mạng nếu không được quản lý.

#### Cơ chế hoạt động:
- Gửi dữ liệu trực tiếp dưới dạng các datagram mà không cần "bắt tay".
- Bên nhận tự xử lý nếu có lỗi hoặc mất gói tin (nếu cần).

#### Ưu điểm:
- Tốc độ nhanh, độ trễ thấp.
- Phù hợp với các ứng dụng không cần độ tin cậy tuyệt đối hoặc có thể tự xử lý lỗi ở tầng ứng dụng.

#### Nhược điểm:
- Không đảm bảo dữ liệu đến đúng thứ tự hoặc đầy đủ.
- Dễ bị mất gói tin trong mạng không ổn định.

#### Ứng dụng thực tế:
- Streaming video/audio (YouTube, Zoom) - mất vài gói tin không ảnh hưởng lớn.
- Game online - ưu tiên tốc độ hơn độ tin cậy.
- DNS (Domain Name System).

---

### 3. So sánh TCP và UDP
| Tiêu chí              | TCP                          | UDP                          |
|-----------------------|------------------------------|------------------------------|
| **Kết nối**           | Có (Connection-Oriented)     | Không (Connectionless)       |
| **Độ tin cậy**        | Cao (Reliable)               | Thấp (Unreliable)            |
| **Tốc độ**            | Chậm hơn                    | Nhanh hơn                    |
| **Kiểm soát lỗi**     | Có (Retransmission)          | Không                        |
| **Kiểm soát luồng**   | Có (Flow Control)            | Không                        |
| **Ứng dụng**          | Web, Email, FTP              | Streaming, Game, DNS         |

---

### 4. Khi nào dùng TCP, khi nào dùng UDP?
- **Dùng TCP** khi bạn cần:
  - Độ tin cậy cao (ví dụ: gửi file, giao dịch ngân hàng).
  - Dữ liệu phải đến đúng thứ tự và không bị mất.
- **Dùng UDP** khi bạn cần:
  - Tốc độ cao, chấp nhận mất mát nhỏ (ví dụ: video call, game).
  - Ứng dụng tự xử lý lỗi ở tầng cao hơn.

---

### 5. Một số lưu ý cho Backend Developer
- **Socket Programming:** Khi lập trình mạng (dùng Python, Java, Node.js, v.v.), bạn sẽ làm việc trực tiếp với TCP (qua `SOCK_STREAM`) hoặc UDP (qua `SOCK_DGRAM`).
- **Hiệu năng:** Với hệ thống lớn, cân nhắc dùng UDP nếu cần tối ưu tốc độ, nhưng phải tự xây dựng cơ chế kiểm tra lỗi nếu cần.
- **Debugging:** Khi gặp lỗi mạng, kiểm tra xem vấn đề nằm ở TCP (kết nối bị ngắt) hay UDP (mất gói tin).
