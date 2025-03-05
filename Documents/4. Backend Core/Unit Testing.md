
---
Unit Testing là một phương pháp kiểm thử phần mềm, trong đó các đơn vị/module nhỏ nhất của ứng dụng được kiểm tra riêng lẻ, độc lập với các thành phần khác.

Các đặc điểm chính của Unit Testing:

1. Phạm vi kiểm thử:
    
    - Kiểm tra từng class, method, function hoặc module nhỏ nhất trong ứng dụng.
    - Mỗi test case chỉ kiểm tra một chức năng cụ thể của đơn vị được kiểm tra.
2. Tính độc lập:
    
    - Test case không phụ thuộc vào các thành phần khác của ứng dụng.
    - Sử dụng các mock object để mô phỏng các phụ thuộc.
3. Tự động hóa:
    
    - Unit Test thường được viết và chạy tự động bằng các framework kiểm thử.
    - Giúp phát hiện và sửa chữa lỗi sớm trong quá trình phát triển.
4. Tái sử dụng:
    
    - Các test case có thể được tái sử dụng khi có thay đổi trong ứng dụng.
    - Đảm bảo rằng các chức năng cơ bản vẫn hoạt động đúng.
5. Bao phủ code:
    
    - Mục tiêu là đạt được mức độ bao phủ code (code coverage) càng cao càng tốt.
    - Giúp đảm bảo rằng hầu hết các đoạn code được kiểm tra.

Các kỹ thuật Unit Testing phổ biến:

1. Assertion:
    
    - Kiểm tra các giá trị đầu ra của đơn vị được kiểm tra với giá trị mong đợi.
2. Stub:
    
    - Thay thế các phụ thuộc bên ngoài bằng các stub (giả lập) để cách ly đơn vị.
3. Mock:
    
    - Tạo ra các mock object để mô phỏng hành vi của các phụ thuộc.
4. Spy:
    
    - Theo dõi và kiểm tra các cuộc gọi đến các phương thức của một lớp.
5. Dependency Injection:
    
    - Sử dụng Dependency Injection để tiêm các phụ thuộc, giúp dễ dàng mô phỏng chúng.

Các framework Unit Testing phổ biến trong Java:

- JUnit: Framework kiểm thử phổ biến nhất cho Java.
- Mockito: Framework mocking và stubbing.
- PowerMock: Mở rộng Mockito để mock các lớp/phương thức tĩnh, final, v.v.

Unit Testing là một phần quan trọng trong quá trình phát triển phần mềm, giúp đảm bảo chất lượng code và tăng tính bảo trì của ứng dụng.


---

