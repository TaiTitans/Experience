
---
### **B-Tree**: Khái niệm và ứng dụng

**B-Tree** là một cấu trúc dữ liệu cây tìm kiếm cân bằng (balanced search tree), được thiết kế để hoạt động hiệu quả với hệ thống lưu trữ thứ cấp (như đĩa cứng) và tối ưu hóa cho các thao tác truy xuất và cập nhật dữ liệu.

---

### **Đặc điểm của B-Tree**

1. **Cân bằng chiều cao**:
    
    - B-Tree luôn duy trì cân bằng, nghĩa là tất cả các lá (leaf nodes) đều ở cùng một mức.
        
    - Chiều cao của cây được giữ ở mức tối thiểu, giúp tăng tốc các thao tác tìm kiếm, chèn, và xóa.
        
2. **Nút có nhiều con**:
    
    - Một nút trong B-Tree có thể có nhiều khóa (keys) và nhiều con (child nodes) hơn cây nhị phân.
        
3. **Tính chất của các nút**:
    
    - Một nút có thể chứa tối đa `m-1` khóa (với `m` là bậc của cây).
        
    - Một nút có thể có từ `⌈m/2⌉` đến `m` con (trừ nút gốc).
        
4. **Các khóa trong nút**:
    
    - Các khóa trong mỗi nút được sắp xếp theo thứ tự tăng dần.
        
    - Các khóa này chia không gian thành các khoảng, giúp định hướng tìm kiếm dữ liệu.
        

---

### **Tính chất của B-Tree**

Giả sử cây có bậc `m`:

1. Một nút có thể chứa tối đa `m-1` khóa và tối thiểu `⌈m/2⌉ - 1` khóa (trừ nút gốc, có thể chứa ít hơn).
    
2. Tất cả các nút lá đều nằm trên cùng một mức.
    
3. Độ cao của cây tăng chậm ngay cả khi số lượng phần tử lớn, giúp tối ưu hóa truy cập.
    

---

### **Hoạt động trên B-Tree**

#### 1. **Tìm kiếm (Search)**

- Duyệt tuần tự qua các khóa trong một nút để tìm khóa mong muốn hoặc xác định nhánh (child node) để tiếp tục tìm kiếm.
    
- Độ phức tạp: **O(log n)** (nhờ chiều cao nhỏ của cây).
    

#### 2. **Chèn (Insert)**

- Chèn khóa mới vào nút phù hợp:
    
    1. Tìm nút lá để chèn khóa.
        
    2. Nếu nút chứa đủ chỗ, thêm khóa vào.
        
    3. Nếu nút đầy (đạt `m-1` khóa), tiến hành **split** (chia nút):
        
        - Chia nút thành hai phần và đẩy khóa giữa lên nút cha.
            
        - Tiếp tục split nếu cần.
            

#### 3. **Xóa (Delete)**

- Xóa khóa khỏi cây:
    
    1. Nếu khóa ở nút lá, xóa trực tiếp.
        
    2. Nếu khóa ở nút không phải lá, thay thế bằng khóa gần nhất (tiền nhiệm hoặc hậu nhiệm).
        
    3. Nếu một nút không đủ khóa sau khi xóa, thực hiện **merge** (gộp nút) hoặc **borrow** (mượn khóa từ nút lân cận).
        

---

### **Ưu điểm của B-Tree**

1. **Hiệu quả với bộ nhớ ngoài**:
    
    - Tối ưu hóa cho các hệ thống lưu trữ thứ cấp (như ổ đĩa) nhờ giảm số lần truy cập bộ nhớ.
        
2. **Thích hợp cho dữ liệu lớn**:
    
    - B-Tree có thể lưu trữ và quản lý lượng dữ liệu lớn với chiều cao cây thấp.
        
3. **Hỗ trợ cập nhật hiệu quả**:
    
    - Các thao tác tìm kiếm, chèn và xóa đều có độ phức tạp **O(log n)**.
        

---

### **Nhược điểm của B-Tree**

1. **Phức tạp trong cài đặt**:
    
    - Chèn, xóa và chia/gộp nút đòi hỏi nhiều bước phức tạp.
        
2. **Không tối ưu với dữ liệu nhỏ**:
    
    - Đối với tập dữ liệu nhỏ, các cấu trúc như BST hoặc AVL Tree có thể phù hợp hơn.
        

---

### **Ứng dụng của B-Tree**

1. **Cơ sở dữ liệu (Databases)**:
    
    - B-Tree được sử dụng để tổ chức các chỉ mục (indexes) trong cơ sở dữ liệu.
        
    - Biến thể **B+ Tree** (dạng cải tiến của B-Tree) thường được sử dụng nhiều hơn trong cơ sở dữ liệu hiện đại.
        
2. **Hệ thống tệp (File Systems)**:
    
    - Một số hệ thống tệp như NTFS, HFS+ sử dụng B-Tree để tổ chức tệp và thư mục.
        
3. **Hệ thống lưu trữ (Storage Systems)**:
    
    - Được dùng để quản lý các cấu trúc dữ liệu trên đĩa cứng và SSD.

### **So sánh B-Tree và B+ Tree**

|Đặc điểm|B-Tree|B+ Tree|
|---|---|---|
|**Khóa trong nút lá**|Không bắt buộc lưu toàn bộ khóa|Toàn bộ khóa được lưu trong nút lá|
|**Tìm kiếm**|Tìm kiếm dừng ở bất kỳ nút nào|Tìm kiếm chỉ dừng ở nút lá|
|**Lá liên kết**|Không liên kết|Liên kết tuần tự các nút lá|

---
