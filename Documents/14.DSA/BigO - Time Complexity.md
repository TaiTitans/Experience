
---

![](https://media.geeksforgeeks.org/wp-content/cdn-uploads/mypic.png)

Đo lường lượng thời gian mà một thuật toán cần để hoàn thành dựa trên kích thước đầu vào. Một số ký hiệu Big-O phổ biến bao gồm:

- O(1) - Hằng số: Thời gian thực thi không thay đổi dù kích thước đầu vào thay đổi.
- O(log⁡n)- Logarit: Thời gian thực thi tỷ lệ với logarit của kích thước đầu vào.
- O(n)- Tuyến tính: Thời gian thực thi tỷ lệ trực tiếp với kích thước đầu vào.
- O(nlog⁡n) - Tuyến tính-logarit: Thường gặp trong các thuật toán sắp xếp hiệu quả như Merge Sort và Quick Sort.
- O(n^2) - Bình phương: Thời gian thực thi tỷ lệ với bình phương của kích thước đầu vào, thường thấy trong các thuật toán sắp xếp kém hiệu quả như Bubble Sort, Insertion Sort.
- O(2^n) - Exponential: Thời gian thực thi tăng theo hàm mũ, thường gặp trong các thuật toán liên quan đến tìm kiếm trong không gian lớn.

----
### 1. **Xác định thao tác cơ bản (Basic Operation)**

- Xác định thao tác nào là thao tác được thực hiện thường xuyên nhất hoặc tốn kém nhất về mặt thời gian. Thao tác này sẽ được sử dụng làm cơ sở để đo lường.

### 2. **Đếm số lần thao tác cơ bản được thực hiện**

- Đếm số lần thao tác cơ bản được thực hiện dựa trên kích thước đầu vào (thường ký hiệu là `n`). Điều này có thể bao gồm việc phân tích các vòng lặp, đệ quy, hoặc các điều kiện rẽ nhánh.

### 3. **Xác định dạng công thức (Expression)**

- Viết một công thức biểu diễn số lượng thao tác cơ bản theo `n`.

### 4. **Xác định hàm thời gian lớn nhất (Big-O Notation)**

- Dùng ký hiệu Big-O để biểu diễn biểu thức time complexity. Big-O thể hiện mức độ tăng trưởng của thuật toán khi `n` trở nên lớn.
- Bỏ qua các hằng số và các bậc thấp hơn vì chúng không ảnh hưởng nhiều khi `n` lớn.

---

