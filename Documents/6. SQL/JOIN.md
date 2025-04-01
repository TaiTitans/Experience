
---
## JOIN trong SQL

**JOIN** trong SQL được sử dụng để kết hợp dữ liệu từ hai hoặc nhiều bảng dựa trên một điều kiện chung. Dưới đây là các loại **JOIN** phổ biến:
### 1. **INNER JOIN** (Kết hợp khớp dữ liệu giữa các bảng)

**INNER JOIN** chỉ trả về các bản ghi có dữ liệu khớp nhau giữa các bảng.

#### 📌 **Ví dụ**

Giả sử có hai bảng:
```sql
CREATE TABLE customers (
    customer_id INT PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE orders (
    order_id INT PRIMARY KEY,
    customer_id INT,
    order_date DATE,
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
```
🔹 **Truy vấn INNER JOIN** để lấy danh sách khách hàng có đơn hàng:
```sql
SELECT customers.customer_id, customers.name, orders.order_id, orders.order_date
FROM customers
INNER JOIN orders ON customers.customer_id = orders.customer_id;
```
✅ **Kết quả**: Chỉ hiển thị khách hàng có đơn hàng.

---
### 2. **LEFT JOIN** (Giữ dữ liệu bảng bên trái, nếu không có khớp thì trả về NULL)

**LEFT JOIN** trả về tất cả các bản ghi từ bảng bên trái và các bản ghi khớp từ bảng bên phải. Nếu không có dữ liệu khớp, nó sẽ trả về `NULL`.

#### 📌 **Ví dụ**
```sql
SELECT customers.customer_id, customers.name, orders.order_id, orders.order_date
FROM customers
LEFT JOIN orders ON customers.customer_id = orders.customer_id;
```
✅ **Kết quả**: Hiển thị tất cả khách hàng, kể cả những người chưa có đơn hàng (`order_id` sẽ là `NULL` nếu không có đơn hàng).

---
### 3. **RIGHT JOIN** (Giữ dữ liệu bảng bên phải, nếu không có khớp thì trả về NULL)

**RIGHT JOIN** hoạt động ngược lại với **LEFT JOIN**, giữ lại tất cả bản ghi của bảng bên phải.

#### 📌 **Ví dụ**
```sql
SELECT customers.customer_id, customers.name, orders.order_id, orders.order_date
FROM customers
RIGHT JOIN orders ON customers.customer_id = orders.customer_id;
```
✅ **Kết quả**: Hiển thị tất cả đơn hàng, kể cả khi không có khách hàng khớp (`name` sẽ là `NULL` nếu không có khách hàng).

---
### 4. **FULL OUTER JOIN** (Kết hợp tất cả, nếu không có khớp thì trả về NULL)

**FULL OUTER JOIN** trả về tất cả bản ghi từ cả hai bảng. Nếu không có khớp, giá trị sẽ là `NULL`.

#### 📌 **Ví dụ**
```sql
SELECT customers.customer_id, customers.name, orders.order_id, orders.order_date
FROM customers
FULL OUTER JOIN orders ON customers.customer_id = orders.customer_id;
```
✅ **Kết quả**: Hiển thị tất cả khách hàng và tất cả đơn hàng, ngay cả khi không có dữ liệu khớp nhau.

---
### 5. **CROSS JOIN** (Kết hợp tất cả bản ghi từ hai bảng)

**CROSS JOIN** tạo ra một tập hợp dữ liệu bằng cách kết hợp từng dòng của bảng A với từng dòng của bảng B.

#### 📌 **Ví dụ**
```sql
SELECT customers.name, orders.order_id
FROM customers
CROSS JOIN orders;
```
✅ **Kết quả**: Nếu có 5 khách hàng và 3 đơn hàng, sẽ tạo ra 5 × 3 = 15 bản ghi.

---
### 6. **SELF JOIN** (JOIN chính nó)

**SELF JOIN** là khi một bảng JOIN với chính nó.

#### 📌 **Ví dụ**

Tìm nhân viên có quản lý trong cùng bảng:
```sql
SELECT e1.employee_id, e1.name AS employee_name, e2.name AS manager_name
FROM employees e1
JOIN employees e2 ON e1.manager_id = e2.employee_id;
```
✅ **Kết quả**: Hiển thị danh sách nhân viên và tên của người quản lý.

---
### **Cách MySQL thực hiện lệnh JOIN bên trong**

Khi bạn chạy một lệnh `JOIN` trong MySQL, hệ thống sẽ tối ưu hóa truy vấn và chọn chiến lược tốt nhất để thực hiện. Dưới đây là cách MySQL xử lý JOIN bên trong, bao gồm các bước quan trọng và thuật toán được sử dụng.

## **1. Các bước MySQL xử lý một lệnh JOIN**

### **1️⃣ Phân tích cú pháp (Parsing)**

- MySQL kiểm tra cú pháp SQL để đảm bảo truy vấn hợp lệ.
    
- Nếu có lỗi (ví dụ: bảng không tồn tại, lỗi cú pháp), nó sẽ dừng ngay lập tức.
    

### **2️⃣ Lập kế hoạch thực thi (Query Optimization)**

- MySQL xây dựng **query execution plan** (kế hoạch thực thi truy vấn) bằng cách:
    
    - Kiểm tra có bao nhiêu bảng cần JOIN.
        
    - Xác định kiểu JOIN (`INNER JOIN`, `LEFT JOIN`, `RIGHT JOIN`, v.v.).
        
    - Chọn **thứ tự JOIN tối ưu** dựa trên số lượng bản ghi và index.
        

### **3️⃣ Chọn chiến lược JOIN (Join Execution Strategy)**

- Dựa vào cấu trúc bảng, index và kích thước dữ liệu, MySQL sẽ quyết định dùng thuật toán JOIN nào (Nested Loop, Hash Join, Sort-Merge Join).
    

### **4️⃣ Thực thi JOIN**

- Sau khi có kế hoạch, MySQL thực hiện JOIN bằng cách lặp qua các bảng và kết hợp dữ liệu.

---
## **2. Các thuật toán JOIN trong MySQL**

MySQL sử dụng ba chiến lược chính để thực hiện JOIN:

### **🔹 1. Nested Loop Join (Lặp lồng nhau)**

- Đây là phương pháp mặc định MySQL sử dụng.
    
- Nếu không có index, MySQL sẽ thực hiện một vòng lặp trên từng dòng của bảng A và tìm dòng khớp trong bảng B.
    
- Nếu có index, MySQL có thể tận dụng để tìm dòng khớp nhanh hơn.
🔹 **Ví dụ về Nested Loop Join (Không có index)**
```sql
SELECT * 
FROM customers c 
JOIN orders o ON c.customer_id = o.customer_id;
```
Cách hoạt động bên trong (giả lập):
```
FOR mỗi dòng trong bảng `customers`:
    FOR mỗi dòng trong bảng `orders`:
        IF c.customer_id = o.customer_id THEN
            TRẢ VỀ bản ghi khớp
```
⚡ **Nhược điểm**: Rất chậm nếu không có index (O(n²) trong trường hợp xấu nhất).

**✅ Cải thiện:** Đánh index vào `orders.customer_id` giúp MySQL có thể tìm nhanh hơn.

### **🔹 2. Index Nested Loop Join**

- Nếu cột trong điều kiện JOIN có **index**, MySQL sẽ sử dụng index để tìm dòng khớp nhanh hơn thay vì duyệt toàn bộ bảng.
    

🔹 **Ví dụ**
```sql
SELECT * 
FROM customers c 
JOIN orders o ON c.customer_id = o.customer_id;
```

Nếu `orders.customer_id` có index:
```
FOR mỗi dòng trong bảng `customers`:
    Dùng index trên `orders.customer_id` để tìm dòng khớp nhanh hơn (O(log n))
```

✅ **Tăng tốc độ đáng kể nếu bảng lớn**.

### **🔹 3. Block Nested Loop Join (BNL Join)**

- Khi không có index, MySQL có thể tải một nhóm dữ liệu nhỏ vào bộ nhớ và lặp qua dữ liệu còn lại để giảm số lần truy cập ổ cứng.
    

🔹 **Ví dụ**
```sql
SELECT * 
FROM big_table t1
JOIN big_table t2 ON t1.id = t2.id;
```

🔥 **Tối ưu hơn so với Nested Loop thông thường nếu không có index**.
### **🔹 4. Hash Join (Không có trong MySQL)**

- Một số hệ quản trị CSDL như PostgreSQL hoặc SQL Server có **Hash Join**, nhưng MySQL **không có** Hash Join.
    
- Trong Hash Join, hệ thống sẽ tạo một bảng băm (hash table) để tìm kiếm nhanh hơn.
    
- MySQL không dùng Hash Join vì nó tập trung vào **index-based execution**.

### **🔹 5. Sort-Merge Join**

- Sử dụng khi MySQL có **ORDER BY** trên cột JOIN.
    
- Nếu bảng đã được sắp xếp theo cột JOIN, MySQL có thể **duyệt song song** hai bảng thay vì phải duyệt toàn bộ.
    

🔹 **Ví dụ**
```sql
SELECT * 
FROM customers c
JOIN orders o ON c.customer_id = o.customer_id
ORDER BY c.customer_id;
```

✅ **Tốt khi hai bảng đã có index sắp xếp theo khóa JOIN**.

## **3. Cách kiểm tra MySQL thực hiện JOIN như thế nào?**

Bạn có thể dùng lệnh `EXPLAIN` để xem MySQL thực hiện JOIN như thế nào.
```sql
EXPLAIN SELECT * 
FROM customers c 
JOIN orders o ON c.customer_id = o.customer_id;
```

📌 **Ý nghĩa một số cột trong kết quả `EXPLAIN`**:

- `type`: Loại JOIN (ALL, index, range, ref, eq_ref).
    
- `possible_keys`: Các index có thể sử dụng.
    
- `key`: Index thực tế được sử dụng.
    
- `rows`: Số dòng MySQL dự kiến duyệt qua.
    
- `Extra`: Thông tin bổ sung (ví dụ: "Using index", "Using where", v.v.).

## **4. Khi nào nên dùng Index để tối ưu JOIN?**

- Khi bảng lớn (> 100.000 dòng), nên **đánh index vào cột JOIN**.
    
- Luôn kiểm tra với `EXPLAIN` để xem MySQL có sử dụng index không.
    
- Nếu JOIN nhiều bảng, hãy sắp xếp **thứ tự bảng tối ưu** để MySQL JOIN hiệu quả nhất.



