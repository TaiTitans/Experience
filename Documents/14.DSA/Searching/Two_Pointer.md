
---
### **1. Two Pointers là gì?**

Two Pointers (Hai con trỏ) là một kỹ thuật tối ưu hóa thuật toán, trong đó ta sử dụng hai biến con trỏ để duyệt qua dữ liệu một cách hiệu quả hơn so với các phương pháp thông thường như brute force (duyệt toàn bộ cặp phần tử).

Hai con trỏ có thể di chuyển theo nhiều cách khác nhau:

- **Two Pointers đối đầu (Opposite Direction Pointers):** Một con trỏ bắt đầu từ đầu, một con trỏ bắt đầu từ cuối và chúng tiến gần nhau.
- **Two Pointers cùng chiều (Same Direction Pointers):** Cả hai con trỏ cùng di chuyển về một hướng, thường là duyệt trên danh sách hoặc chuỗi.
### **2. Khi nào nên dùng Two Pointers?**

Two Pointers được sử dụng hiệu quả khi:

- Cần tìm cặp hoặc tập hợp phần tử thỏa mãn điều kiện nào đó (ví dụ: tổng bằng một giá trị, loại bỏ phần tử trùng lặp, tìm khoảng cách nhỏ nhất, ...).
- Có thể tối ưu hóa bằng cách tránh lặp lại phần tử hoặc tránh duyệt qua một vùng dữ liệu nhiều lần.
- Dữ liệu đã được sắp xếp hoặc có thể sắp xếp dễ dàng.

### **3. Các loại bài toán sử dụng Two Pointers**

Dưới đây là một số loại bài toán phổ biến có thể áp dụng kỹ thuật này:

Tìm hai số có tổng bằng một giá trị (`Two Sum`)

Xóa phần tử trùng lặp trong mảng (`Remove Duplicates`)

Tìm dãy con có tổng nhỏ nhất (`Smallest Subarray with Sum ≥ S`)

Sắp xếp màu (Dutch National Flag Problem)
### **Tổng kết**

|**Loại bài toán**|**Chiến lược**|**Độ phức tạp**|
|---|---|---|
|Tìm hai số có tổng bằng `target`|Hai con trỏ đối đầu|O(n)O(n)O(n)|
|Xóa phần tử trùng lặp|Hai con trỏ cùng chiều|O(n)O(n)O(n)|
|Dãy con có tổng ≥ `S`|Cửa sổ trượt|O(n)O(n)O(n)|
|Sắp xếp màu|Ba con trỏ|O(n)O(n)O(n)|

Two Pointers là một kỹ thuật mạnh mẽ giúp tối ưu hóa bài toán có liên quan đến tìm kiếm, sắp xếp và xử lý dãy số. Nó giảm đáng kể độ phức tạp của thuật toán từ O(n2)O(n^2)O(n2) xuống O(n)O(n)O(n) trong nhiều trường hợp.
