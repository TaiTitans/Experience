
---
Ví dụ 1 table có 10tr record.
Rules: Table (quá 5tr record hoặc >= 2GB data) 

Cách chia ngang: 
-> Tách thành 2 table (1 bảng userId chẵn, 1 bảng userId lẻ) -> Mỗi bảng 5tr record
-> Sử dụng UNION để truy vấn 2 bảng


Cách chia dọc:
- Bảng 1 lấy các trường thường query
- Bảng 2 các trường phụ
- Dùng Join + ON để lấy full 2 bảng

Cách chia 2 table khác database:


---
=> Nếu muốn sửa dữ liệu phải sửa 2 bảng.
=> Tính tổng phải query nhiều bảng.

