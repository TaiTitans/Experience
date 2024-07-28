
---

Networking models cung cấp các cấu trúc và tiêu chuẩn giúp tổ chức và quản lý truyền thông dữ liệu giữa các thiết bị trong mạng. Hai mô hình mạng phổ biến nhất là OSI (Open Systems Interconnection) và TCP/IP (Transmission Control Protocol/Internet Protocol). Dưới đây là chi tiết về cả hai mô hình:

### OSI Model

Mô hình OSI được phát triển bởi Tổ chức Tiêu chuẩn hóa Quốc tế (ISO) và có 7 lớp, mỗi lớp phục vụ một chức năng cụ thể trong việc truyền tải dữ liệu qua mạng.

1. **Lớp 1: Physical (Vật lý)**
    
    - Chịu trách nhiệm về việc truyền tải bit qua phương tiện truyền dẫn vật lý (dây cáp, sóng vô tuyến).
    - Các thiết bị như hub, repeater thuộc lớp này.
2. **Lớp 2: Data Link (Liên kết dữ liệu)**
    
    - Đảm bảo truyền tải dữ liệu không lỗi giữa hai nút lân cận trong cùng một mạng.
    - Chia thành hai lớp con: LLC (Logical Link Control) và MAC (Media Access Control).
    - Các thiết bị như switch, bridge thuộc lớp này.
3. **Lớp 3: Network (Mạng)**
    
    - Chịu trách nhiệm về định tuyến và chuyển tiếp gói tin từ nguồn đến đích qua nhiều mạng khác nhau.
    - Sử dụng địa chỉ IP.
    - Các thiết bị như router thuộc lớp này.
4. **Lớp 4: Transport (Vận chuyển)**
    
    - Cung cấp truyền tải dữ liệu đáng tin cậy giữa các thiết bị.
    - Các giao thức như TCP (đảm bảo) và UDP (không đảm bảo).
5. **Lớp 5: Session (Phiên)**
    
    - Quản lý và kiểm soát các phiên kết nối giữa các ứng dụng.
    - Đảm bảo các phiên kết nối được thiết lập, duy trì và kết thúc đúng cách.
6. **Lớp 6: Presentation (Trình bày)**
    
    - Chịu trách nhiệm về mã hóa, giải mã, và chuyển đổi định dạng dữ liệu để đảm bảo tính tương thích giữa các hệ thống khác nhau.
7. **Lớp 7: Application (Ứng dụng)**
    
    - Cung cấp các dịch vụ mạng trực tiếp cho ứng dụng của người dùng cuối.
    - Các giao thức như HTTP, FTP, SMTP thuộc lớp này.

### TCP/IP Model

Mô hình TCP/IP được phát triển bởi Bộ Quốc phòng Hoa Kỳ và có 4 lớp, mỗi lớp tương ứng với một hoặc nhiều lớp trong mô hình OSI.

1. **Lớp 1: Link (Liên kết)**
    
    - Tương ứng với lớp Physical và Data Link của mô hình OSI.
    - Quản lý truyền tải dữ liệu giữa các nút trong cùng một mạng cục bộ.
2. **Lớp 2: Internet (Mạng)**
    
    - Tương ứng với lớp Network của mô hình OSI.
    - Quản lý định tuyến gói tin và sử dụng giao thức IP để định địa chỉ và chuyển tiếp gói tin.
3. **Lớp 3: Transport (Vận chuyển)**
    
    - Tương ứng với lớp Transport của mô hình OSI.
    - Cung cấp truyền tải dữ liệu đáng tin cậy (TCP) hoặc không đảm bảo (UDP) giữa các thiết bị.
4. **Lớp 4: Application (Ứng dụng)**
    
    - Tương ứng với lớp Session, Presentation, và Application của mô hình OSI.
    - Cung cấp các dịch vụ mạng trực tiếp cho ứng dụng của người dùng cuối, bao gồm HTTP, FTP, SMTP, và nhiều giao thức khác.

### So sánh OSI và TCP/IP

- **Số lượng lớp**: OSI có 7 lớp, trong khi TCP/IP có 4 lớp.
- **Thiết kế và Phát triển**: OSI là một mô hình lý thuyết chuẩn hóa, trong khi TCP/IP được phát triển dựa trên thực tiễn và là bộ giao thức được sử dụng rộng rãi trên Internet.
- **Chức năng**: Cả hai mô hình đều cung cấp các chức năng tương tự nhau, nhưng cách chúng tổ chức và đặt tên các lớp có sự khác biệt.
- **Khả năng mở rộng**: TCP/IP được xem là linh hoạt và dễ triển khai hơn so với mô hình OSI.

Cả hai mô hình này đều đóng vai trò quan trọng trong việc thiết kế, triển khai và quản lý mạng máy tính hiện đại.