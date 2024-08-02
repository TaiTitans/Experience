
---
**Mục đích:** Thủ tục lưu trữ là các khối mã SQL được lưu trữ và thực thi trên máy chủ cơ sở dữ liệu. Chúng giúp tái sử dụng mã và quản lý các hoạt động phức tạp.

**Cú pháp tạo thủ tục lưu trữ:**

```SQL
CREATE PROCEDURE ten_thu_tuc ([IN|OUT|INOUT] tham_so KIEU_DU_LIEU)
BEGIN
    -- Các câu lệnh SQL
END;

```

**Ví dụ:**
```SQL
CREATE PROCEDURE them_sinh_vien(
    IN ma INT, 
    IN ten VARCHAR(100), 
    IN tuoi INT, 
    IN lop VARCHAR(10)
)
BEGIN
    INSERT INTO sinh_vien (ma_sinh_vien, ten, tuoi, lop) VALUES (ma, ten, tuoi, lop);
END;
```

**Gọi thủ tục:**

```SQL
CALL them_sinh_vien(1, 'Nguyen Van A', 20, '12A1');
```

