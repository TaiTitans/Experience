

---
**Thuật toán phát hiện chu trình Floyd**, hay còn gọi là **Thuật toán Tortoise and Hare (Rùa và Thỏ)**, là một phương pháp phát hiện chu trình trong một chuỗi hoặc danh sách liên kết. Thuật toán này thường được sử dụng để phát hiện chu trình trong danh sách liên kết hoặc xác định vòng lặp trong các quá trình lặp lại. Thuật toán này có độ phức tạp thời gian O(n)O(n)O(n) và độ phức tạp không gian O(1)O(1)O(1).


### **Khái niệm chính**

1. **Hai con trỏ:**
    
    - Con trỏ chậm (Rùa): Di chuyển từng bước một.
    - Con trỏ nhanh (Thỏ): Di chuyển hai bước một lúc.
2. **Phát hiện chu trình:**
    
    - Nếu có chu trình, hai con trỏ sẽ gặp nhau.
    - Nếu không có chu trình, con trỏ nhanh sẽ tới cuối danh sách.

### **Các bước của thuật toán**

1. **Khởi tạo:**
    
    - Bắt đầu cả hai con trỏ tại đầu danh sách liên kết.
2. **Duyệt danh sách:**
    
    - Di chuyển con trỏ chậm một bước.
    - Di chuyển con trỏ nhanh hai bước.
    - Nếu hai con trỏ gặp nhau, có chu trình.
3. **Xác định chu trình:**
    
    - Để tìm điểm bắt đầu của chu trình:
        - Di chuyển một con trỏ về đầu danh sách và giữ con trỏ kia tại điểm gặp.
        - Di chuyển cả hai con trỏ từng bước một cho đến khi chúng gặp nhau. Điểm gặp này là điểm bắt đầu của chu trình.
4. **Không có chu trình:**
    
    - Nếu con trỏ nhanh đến cuối danh sách (null), không có chu trình.


### **Tại sao nó hoạt động**

1. **Điểm gặp trong chu trình:**
    
    - Con trỏ nhanh di chuyển nhanh gấp đôi con trỏ chậm. Nếu có chu trình, con trỏ nhanh sẽ "bắt kịp" con trỏ chậm trong chu trình.
2. **Tìm điểm bắt đầu của chu trình:**
    
    - Khi con trỏ chậm được đưa về đầu danh sách và di chuyển từng bước, nó sẽ gặp lại con trỏ nhanh tại điểm bắt đầu của chu trình.

### **Ứng dụng**

- Phát hiện chu trình trong danh sách liên kết.
- Phát hiện trạng thái lặp lại trong các quá trình lặp.
- Giải quyết các bài toán như phát hiện số trùng lặp trong mảng.