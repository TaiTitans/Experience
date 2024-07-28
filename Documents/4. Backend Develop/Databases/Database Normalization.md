
---
Database Normalization là một quá trình tổ chức dữ liệu trong một cơ sở dữ liệu quan hệ nhằm giảm thiểu sự lặp và tăng tính nhất quán của dữ liệu. Quá trình này bao gồm chia dữ liệu thành các bảng riêng biệt và xác định mối quan hệ giữa các bảng này.

Có 5 mức chuẩn hóa cơ bản:

1. **First Normal Form (1NF)**: Yêu cầu rằng tất cả các giá trị trong bảng phải là nguyên tử (atomic), không được chứa tập hợp các giá trị. Mỗi cột phải chứa một giá trị duy nhất.
    
2. **Second Normal Form (2NF)**: Yêu cầu rằng bảng phải ở 1NF và mọi cột không phải khóa phải phụ thuộc hoàn toàn vào khóa chính.
    
3. **Third Normal Form (3NF)**: Yêu cầu rằng bảng phải ở 2NF và mọi cột không phải khóa phải phụ thuộc trực tiếp vào khóa chính, không được phụ thuộc vào cột không phải khóa khác.
    
4. **Boyce-Codd Normal Form (BCNF)**: Yêu cầu rằng bảng phải ở 3NF và mọi determinant (cột hoặc tập hợp cột xác định duy nhất một bộ) phải là sluperkey.
    
5. **Fourth Normal Form (4NF)**: Yêu cầu rằng bảng phải ở BCNF và không có sự phụ thuộc đa trị (multivalued dependency).
    

Lợi ích của chuẩn hóa cơ sở dữ liệu:

- Giảm sự dư thừa dữ liệu.
- Tăng tính nhất quán của dữ liệu.
- Dễ dàng thực hiện các thao tác CRUD.
- Giảm kích thước cơ sở dữ liệu.
- Dễ dàng bảo trì và nâng cấp cơ sở dữ liệu.

Tuy nhiên, quá trình chuẩn hóa cũng có một số hạn chế:

- Tăng số lượng bảng, dẫn đến nhiều Join.
- Có thể mất một số thông tin ngữ nghĩa.
- Không phù hợp với một số mô hình dữ liệu như NoSQL.

Trong thực tế, các nhà thiết kế cơ sở dữ liệu thường cân bằng giữa chuẩn hóa và denormalization (tối ưu hóa) để đáp ứng tốt nhất các yêu cầu của ứng dụng.