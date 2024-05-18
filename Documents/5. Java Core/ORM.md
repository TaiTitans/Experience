
---

**ORM** (Object-Relational Mapping) là một kỹ thuật lập trình giúp lập trình viên tương tác với cơ sở dữ liệu bằng cách sử dụng các đối tượng trong ngôn ngữ lập trình thay vì phải viết các câu lệnh SQL trực tiếp. ORM ánh xạ các lớp và đối tượng trong mã nguồn của ứng dụng với các bảng và các bản ghi trong cơ sở dữ liệu.


## Lợi ích của ORM

- **Trừu tượng hóa cơ sở dữ liệu**: Giảm bớt sự phụ thuộc vào loại cơ sở dữ liệu cụ thể.
- **Tăng năng suất**: Giúp lập trình viên làm việc hiệu quả hơn bằng cách giảm thiểu việc phải viết và bảo trì mã SQL.
- **An toàn hơn**: Giảm thiểu nguy cơ các lỗi bảo mật như SQL Injection.
- **Bảo trì dễ dàng hơn**: Mã nguồn dễ đọc và bảo trì hơn so với việc sử dụng các câu lệnh SQL trực tiếp.

---

### So sánh ORM và SQL Thông thường

### So sánh chi tiết:

1. **Khai báo và tạo bảng:**
    
    - **SQL Thông Thường**: Bạn phải viết các câu lệnh SQL để tạo bảng và quản lý schema.
    - **ORM**: Schema được tạo tự động dựa trên các lớp ánh xạ (Entity Class). Các thuộc tính của lớp sẽ được ánh xạ tới các cột trong bảng.
2. **Thao tác với dữ liệu:**
    
    - **SQL Thông Thường**: Bạn phải viết các câu lệnh SQL như SELECT, INSERT, UPDATE, DELETE.
    - **ORM**: Bạn thao tác với dữ liệu thông qua các đối tượng Java. ORM tự động chuyển đổi các thao tác này thành các câu lệnh SQL.
3. **Truy vấn dữ liệu:**
    
    - **SQL Thông Thường**: Bạn viết các câu lệnh SQL trực tiếp và xử lý kết quả trả về từ ResultSet.
    - **ORM**: Bạn sử dụng các phương thức của ORM như `session.createQuery` hoặc `session.find`.
4. **Độ phức tạp và bảo trì:**
    
    - **SQL Thông Thường**: Mã nguồn có thể phức tạp và khó bảo trì khi quy mô dự án lớn.
    - **ORM**: Mã nguồn dễ đọc, dễ bảo trì và dễ dàng mở rộng.