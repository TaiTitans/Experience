
---
**Mục đích:** Chỉ mục (index) giúp tăng tốc độ truy vấn dữ liệu bằng cách tạo ra các cấu trúc dữ liệu (như cây B+, hash) để dễ dàng tìm kiếm và truy xuất dữ liệu trong bảng.

**Các loại chỉ mục:**

- **Chỉ mục chính (Primary Index):** Tự động được tạo khi bạn định nghĩa một cột là PRIMARY KEY. Đây là chỉ mục duy nhất và không thể có giá trị NULL.
```SQL
CREATE TABLE sinh_vien (
    ma_sinh_vien INT PRIMARY KEY,
    ten VARCHAR(100)
);
```

**Chỉ mục duy nhất (Unique Index):** Đảm bảo rằng tất cả các giá trị trong cột chỉ mục là duy nhất.

```SQL
CREATE UNIQUE INDEX idx_ma_sv ON sinh_vien(ma_sinh_vien);
```
**Chỉ mục toàn văn (Full-text Index):** Dùng để tìm kiếm toàn văn bản, chủ yếu sử dụng cho các cột kiểu TEXT hoặc VARCHAR lớn.
```SQL
CREATE FULLTEXT INDEX idx_ten ON sinh_vien(ten);
```
**Chỉ mục không duy nhất (Non-unique Index):** Được sử dụng để tăng tốc độ truy vấn cho các cột không phải là khóa chính hay duy nhất.
```SQL
CREATE INDEX idx_ten ON sinh_vien(ten);
```

**Lợi ích của indexing:**

- Tăng tốc độ truy vấn SELECT.
- Giảm chi phí tính toán khi JOIN các bảng.

**Hạn chế của indexing:**

- Tốn thêm dung lượng lưu trữ.
- Cập nhật, chèn và xóa dữ liệu có thể chậm hơn vì cần cập nhật chỉ mục.