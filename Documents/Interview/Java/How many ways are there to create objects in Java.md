
---
**Có nhiều cách để tạo đối tượng trong Java:**

- **Sử dụng câu lệnh new**: Đây là cách phổ biến nhất để tạo một đối tượng.
- **Sử dụng phản xạ (reflection)**: Dùng phương thức Class.newInstance() để tạo đối tượng.
- **Gọi phương thức clone() của đối tượng**: Tạo một bản sao của đối tượng hiện có.
- **Sử dụng giải tuần tự hóa (deserialization)**: Gọi phương thức readObject() của đối tượng java.io.ObjectInputStream để khôi phục đối tượng từ dữ liệu tuần tự hóa.