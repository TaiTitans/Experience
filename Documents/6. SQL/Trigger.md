
---
**Mục đích:** Trigger là các khối mã SQL tự động thực hiện trước hoặc sau một hành động (INSERT, UPDATE, DELETE) trên một bảng.

**Cú pháp tạo trigger:**

```SQL
CREATE TRIGGER ten_trigger
    { BEFORE | AFTER } { INSERT | UPDATE | DELETE }
    ON ten_bang
    FOR EACH ROW
BEGIN
    -- Nội dung trigger
END;
```

**Ví dụ:**

```SQL
CREATE TRIGGER trg_sau_insert
AFTER INSERT ON sinh_vien
FOR EACH ROW
BEGIN
    INSERT INTO log_sinh_vien (ma_sinh_vien, hanh_dong) VALUES (NEW.ma_sinh_vien, 'inserted');
END;
```

