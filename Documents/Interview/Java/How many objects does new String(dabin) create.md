
---
**Việc sử dụng phương thức này tạo ra hai đối tượng chuỗi (với điều kiện không có đối tượng chuỗi "dabin" trong vùng nhớ hằng chuỗi - string constant pool).**

- "dabin" là một chuỗi ký tự nguyên bản (string literal), vì vậy tại thời điểm biên dịch, một đối tượng chuỗi được tạo ra trong vùng nhớ hằng chuỗi để trỏ đến chuỗi ký tự "dabin" này.
- Việc sử dụng từ khóa new tạo ra một đối tượng chuỗi khác trong bộ nhớ heap.
