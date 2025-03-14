
---
## **Hiểu sâu về Index trong SQL**

### **1. Index trong SQL là gì?**

Index (chỉ mục) trong SQL là một cấu trúc dữ liệu giúp tăng tốc độ truy vấn trong cơ sở dữ liệu bằng cách cho phép truy xuất dữ liệu nhanh hơn thay vì quét toàn bộ bảng (Full Table Scan). Index hoạt động giống như một mục lục trong sách, giúp tìm kiếm dữ liệu nhanh hơn.

### **2. Cách Index hoạt động**

Index thường được triển khai dưới dạng **B-Tree** hoặc **Hash Table**, trong đó:

- **B-Tree Index**: Cấu trúc dạng cây giúp tìm kiếm dữ liệu nhanh hơn với độ phức tạp O(log n).
- **Hash Index**: Sử dụng bảng băm để tìm kiếm khóa chính xác (Hash lookup), phù hợp với truy vấn so sánh trực tiếp (`=`) nhưng không tối ưu cho truy vấn phạm vi (`BETWEEN`, `LIKE`).

Khi có index, SQL sẽ sử dụng index để tìm vị trí của hàng thay vì phải quét toàn bộ bảng.

### **3. Các loại Index trong SQL**

#### **3.1. Primary Index (Clustered Index)**

- Mỗi bảng chỉ có một **Clustered Index**.
- Sắp xếp vật lý dữ liệu theo thứ tự của index.
- Tìm kiếm nhanh nhưng khi có nhiều cập nhật hoặc chèn dữ liệu thì có thể gây phân mảnh.
- Mặc định, khóa chính (`PRIMARY KEY`) sẽ tạo một **Clustered Index**.
```sql
CREATE TABLE users (
    id INT PRIMARY KEY,  -- Mặc định tạo Clustered Index
    name VARCHAR(255),
    email VARCHAR(255)
);
```
#### **3.2. Secondary Index (Non-Clustered Index)**

- Không thay đổi cách lưu trữ vật lý của bảng.
- Có thể tạo nhiều **Non-Clustered Index** trên một bảng.
- Thích hợp khi cần tăng tốc độ tìm kiếm trên các cột không phải khóa chính.
```sql
CREATE INDEX idx_email ON users(email);
```
#### **3.3. Unique Index**

- Đảm bảo rằng các giá trị trong một cột hoặc tập hợp cột là duy nhất.
- Thường được sử dụng với `UNIQUE CONSTRAINT`.
```sql
CREATE UNIQUE INDEX idx_unique_email ON users(email);
```
#### **3.4. Composite Index**

- Là index trên nhiều cột (`(col1, col2, col3)`).
- Hữu ích khi truy vấn thường xuyên sử dụng nhiều cột trong điều kiện tìm kiếm.
```sql
CREATE INDEX idx_name_email ON users(name, email);
```

📌 **Lưu ý**: Index này sẽ hỗ trợ tốt cho các truy vấn:

```sql
SELECT * FROM users WHERE name = 'John';  -- OK
SELECT * FROM users WHERE name = 'John' AND email = 'john@example.com'; -- OK
SELECT * FROM users WHERE email = 'john@example.com'; -- Không tối ưu (vì email đứng sau name)
```
#### **3.5. Full-Text Index**

- Được sử dụng cho tìm kiếm văn bản trong các cột kiểu `TEXT`, `VARCHAR`, `CHAR`.
- Tăng tốc độ tìm kiếm với các từ khóa trong dữ liệu dạng text.
```sql
CREATE FULLTEXT INDEX idx_fulltext_name ON users(name);
```
#### **3.6. Partial Index**

- Chỉ lập index trên một phần của bảng (ví dụ: chỉ index những dòng có trạng thái `active`).
#### **3.7. Covering Index**

- Index chứa tất cả các cột cần thiết cho một truy vấn cụ thể để tránh truy cập lại bảng chính.

### **4. Khi nào nên đánh Index?**

Bạn nên tạo index khi:

✅ **Truy vấn SELECT có điều kiện lọc (`WHERE`), `JOIN`, `ORDER BY`, `GROUP BY` trên một hoặc nhiều cột.**  
✅ **Các cột thường xuyên được tìm kiếm.**  
✅ **Cột được dùng làm khóa chính (`PRIMARY KEY`) hoặc ràng buộc duy nhất (`UNIQUE`).**  
✅ **Các truy vấn thường xuyên thực hiện tìm kiếm theo phạm vi (`BETWEEN`, `LIKE 'abc%'`).**  
✅ **Cột có nhiều giá trị khác nhau (độ phân tán dữ liệu cao).**

### **Không nên tạo Index khi:**

❌ Bảng có **ít dữ liệu** (vì lợi ích từ index không đáng kể).  
❌ Cột có **giá trị trùng lặp nhiều** (như giới tính, trạng thái đơn hàng).  
❌ Index quá nhiều có thể làm chậm **INSERT, UPDATE, DELETE** vì cần cập nhật index mỗi lần thay đổi dữ liệu.


## **5. Các câu hỏi phỏng vấn về Index trong SQL**

Dưới đây là một số câu hỏi phỏng vấn phổ biến về Index:

### **5.1. Kiến thức cơ bản về Index**

- Index trong SQL là gì?
- Có bao nhiêu loại index?
- Phân biệt **Clustered Index** và **Non-Clustered Index**?
- Một bảng có thể có bao nhiêu Clustered Index? (Chỉ 1)
- Khi nào nên sử dụng Unique Index?
- Khi nào nên dùng Composite Index?
- Index ảnh hưởng thế nào đến hiệu suất **INSERT, UPDATE, DELETE**?

### **5.2. Câu hỏi chuyên sâu**

- **Làm sao để kiểm tra một bảng đã có index nào?**  
    📌 **MySQL:**
```sql
SHOW INDEX FROM users;
```
📌 **PostgreSQL:**
```sql
SELECT * FROM pg_indexes WHERE tablename = 'users';
```

- **Làm sao để xóa index?**
```sql
DROP INDEX idx_email;
```

- **Index có hỗ trợ tìm kiếm giá trị NULL không?** (Có, nhưng hiệu suất thấp hơn so với giá trị không NULL)
- **Tại sao đôi khi index không được sử dụng?**
    - Query không phù hợp với index.
    - Optimizer quyết định Full Table Scan nhanh hơn.
    - Index bị hỏng (cần **ANALYZE** hoặc **REBUILD** index).

