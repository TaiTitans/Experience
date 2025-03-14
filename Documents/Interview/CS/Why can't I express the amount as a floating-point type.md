
---
**Vì số thập phân được lưu trữ trong máy tính thực chất chỉ là một giá trị xấp xỉ của số thập phân và không phải là giá trị chính xác, bạn không nên sử dụng số dấu phẩy động (floating-point numbers) trong mã của mình để biểu diễn các chỉ số quan trọng như số tiền.**  
Thay vào đó, nên sử dụng BigDecimal hoặc Long để biểu diễn số tiền.