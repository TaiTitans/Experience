
---
Giả sử Shopee có 10000 đơn hàng 1 ngày
Vậy 1 tháng đã tốn 300,000 đơn hàng (record trong db)

Vậy nếu sử dụng 1 table sẽ bị phình to và truy vấn sẽ rất lâu

Vậy solution ở đây là gì ?

-> Cronjob tự động tạo thêm table mỗi tháng. Ví dụ vào 0h00 ngày 1 tháng 2 sẽ tự tạo ra bảng order_detail_202502.
Cứ thế mà kéo dài mỗi tháng.

Nhưng làm sao query theo tháng được ?

Trong mỗi mã đơn hàng sẽ có dạng: 202502..x.xx.x.

Dựa vào đó mà chọn table theo tháng chính xác và truy vấn.
