

---

**Mục đích:** Giao dịch (transaction) là một nhóm các thao tác SQL được thực hiện như một đơn vị logic duy nhất. Tất cả các thao tác trong giao dịch phải được hoàn thành, nếu không sẽ không có thao tác nào được thực hiện.
**Cú pháp:**

```SQL
START TRANSACTION;
-- Các lệnh SQL
COMMIT; -- Hoặc ROLLBACK;
```

**Ví dụ:**
```SQL
START TRANSACTION;
    UPDATE sinh_vien SET lop = '12A1' WHERE ma_sinh_vien = 1;
    DELETE FROM sinh_vien WHERE ma_sinh_vien = 2;
COMMIT;
```

**Tính chất của giao dịch (ACID):**

- **Atomicity:** Tính nguyên tử - giao dịch phải được thực hiện toàn bộ hoặc không có phần nào được thực hiện.
- **Consistency:** Tính nhất quán - giao dịch phải chuyển hệ thống từ trạng thái nhất quán này sang trạng thái nhất quán khác.
- **Isolation:** Tính cô lập - giao dịch này không ảnh hưởng đến giao dịch khác.
- **Durability:** Tính bền vững - khi giao dịch đã được xác nhận, thay đổi sẽ được lưu trữ vĩnh viễn.

