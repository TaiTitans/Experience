
---
**final được sử dụng để sửa đổi thuộc tính, phương thức và lớp:**

- Khi sửa đổi **thuộc tính**: Biểu thị rằng thuộc tính không thể được gán lại giá trị.
- Khi sửa đổi **phương thức**: Biểu thị rằng phương thức không thể bị ghi đè (override).
- Khi sửa đổi **lớp**: Biểu thị rằng lớp không thể được kế thừa.

**finally là một phần của cấu trúc câu lệnh xử lý ngoại lệ:**

- Thường xuất hiện dưới dạng một khối mã trong cấu trúc try-catch-finally, và khối mã này dường như luôn được thực thi bất kể ngoại lệ có xảy ra hay không.

**finalize là một phương thức của lớp Object:**

- Phương thức này thường được gọi bởi bộ thu gom rác (garbage collector). Khi chúng ta ghi đè phương thức này, bộ thu gom rác sẽ gọi nó để dọn dẹp tài nguyên trước khi đối tượng bị thu hồi.
- Tuy nhiên, JVM không đảm bảo rằng phương thức này sẽ luôn được gọi. Ví dụ, khi sử dụng System.gc() để gợi ý thu gom rác, finalize() có thể được gọi, nhưng không chắc chắn.