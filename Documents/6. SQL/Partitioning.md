

---

**Mục đích:** Phân vùng giúp chia một bảng lớn thành nhiều phần nhỏ hơn dựa trên một số tiêu chí (như giá trị của một cột) để quản lý và truy vấn dữ liệu hiệu quả hơn.

**Các loại phân vùng:**

- **RANGE Partitioning:** Chia dữ liệu thành các phần dựa trên một khoảng giá trị.
```SQL
PARTITION BY RANGE (cot_chia_phan)
(
    PARTITION p0 VALUES LESS THAN (value1),
    PARTITION p1 VALUES LESS THAN (value2),
    ...
);
```

**LIST Partitioning:** Chia dữ liệu dựa trên một danh sách các giá trị.

```SQL
PARTITION BY LIST (cot_chia_phan)
(
    PARTITION p0 VALUES IN (value1, value2),
    PARTITION p1 VALUES IN (value3, value4),
    ...
);
```

**HASH Partitioning:** Chia dữ liệu dựa trên giá trị băm của cột.
```SQL
PARTITION BY HASH (cot_chia_phan)
PARTITIONS so_phan_vung;
```

**KEY Partitioning:** Giống như HASH Partitioning nhưng sử dụng hàm băm mặc định của MySQL.

```SQL
PARTITION BY KEY (cot_chia_phan)
PARTITIONS so_phan_vung;
```

VD:
```SQL
CREATE TABLE doanh_thu (
    id INT,
    ngay DATE,
    so_tien FLOAT
)
PARTITION BY RANGE (YEAR(ngay)) (
    PARTITION p0 VALUES LESS THAN (2020),
    PARTITION p1 VALUES LESS THAN (2021),
    PARTITION p2 VALUES LESS THAN (2022)
);
```

