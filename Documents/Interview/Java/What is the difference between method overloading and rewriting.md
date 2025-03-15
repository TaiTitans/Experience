
---
- **Nạp chồng (Overloading)**:
    - Xảy ra trong cùng một lớp, các phương thức có cùng tên nhưng khác tham số (số lượng, kiểu, hoặc thứ tự).
    - Ví dụ: setPerson(), setPerson(String), và setPerson(String, int) là các phương thức nạp chồng.
    - Quyết định phương thức nào được gọi dựa trên tham số truyền vào lúc biên dịch (compile-time).
- **Ghi đè (Overriding)**:
    - Xảy ra giữa lớp cha và lớp con, khi lớp con muốn thay đổi hành vi của phương thức kế thừa từ lớp cha.
    - Tên phương thức, danh sách tham số và kiểu trả về phải giống nhau (trừ trường hợp covariant return type).
    - Quyết định phương thức nào được gọi dựa trên đối tượng thực tế lúc chạy (runtime polymorphism).