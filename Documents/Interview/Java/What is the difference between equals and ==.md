
---
**Đối với các kiểu dữ liệu cơ bản, == so sánh giá trị của chúng.**  
Các kiểu dữ liệu cơ bản không có phương thức equals().

**Đối với các kiểu dữ liệu tổng hợp (composite data types), == so sánh địa chỉ lưu trữ của chúng (tức là kiểm tra xem chúng có phải là cùng một đối tượng hay không).**

- Theo mặc định, == so sánh giá trị địa chỉ.
- Nếu phương thức equals() được ghi đè, việc so sánh sẽ tuân theo logic của phần ghi đè đó.

### Giải thích bổ sung:

- **Kiểu dữ liệu cơ bản (primitive types)**: Như int, double, char, v.v., chỉ lưu giá trị trực tiếp, không phải đối tượng, nên == so sánh giá trị (ví dụ: 5 == 5 trả về true). Chúng không có phương thức như equals() vì không phải là đối tượng.
- **Kiểu dữ liệu tổng hợp (reference types)**: Như các lớp (String, Object, v.v.), == so sánh tham chiếu (địa chỉ bộ nhớ). Ví dụ: new String("a") == new String("a") trả về false vì chúng là hai đối tượng khác nhau trong heap.
- **equals()**: Mặc định (trong Object), so sánh địa chỉ giống ==, nhưng khi được ghi đè (như trong String), so sánh nội dung (ví dụ: "a".equals("a") trả về true).
