
---
a. Khi user lấy dữ liệu lần đầu
    
Không tìm thấy trong local cache.
Không tìm thấy trong Redis cache.
Truy xuất từ cơ sở dữ liệu, sau đó lưu vào Redis và local cache.

b. Khi user lấy dữ liệu với phiên bản không khớp
    
Dữ liệu từ local cache bị lỗi thời (phiên bản thấp hơn).
Truy xuất từ Redis cache.
Cập nhật lại dữ liệu vào local cache.

c. Khi user đặt vé (ghi dữ liệu)
    
Xóa dữ liệu khỏi local cache và Redis cache.
Dữ liệu mới sẽ được cập nhật khi user lấy lại thông tin.

Setup Circuit Breaker để phòng khi redis hoặc db gặp sự cố
Khi lấy dữ liệu từ db buộc sử dụng Lock để đảm bảo chỉ 1 instance query db.