
---

**Tuần tự hóa (Serialization):**

- Quá trình chuyển đổi một đối tượng thành một chuỗi byte được gọi là tuần tự hóa đối tượng.

**Giải tuần tự hóa (Deserialization):**

- Quá trình khôi phục một chuỗi byte trở lại thành một đối tượng được gọi là giải tuần tự hóa đối tượng.
Bổ sung:
- **Tuần tự hóa**:
    - Mục đích: Lưu trữ trạng thái của đối tượng (ví dụ: vào file) hoặc truyền qua mạng.
    - Trong Java, một lớp cần triển khai giao diện Serializable để hỗ trợ tuần tự hóa. 
        
- **Giải tuần tự hóa**:
    - Mục đích: Tái tạo đối tượng từ dữ liệu byte đã lưu hoặc nhận được.