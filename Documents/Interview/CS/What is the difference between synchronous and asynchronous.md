
---
**Đồng bộ (Synchronization):**

- Khi một lời gọi được thực hiện, lời gọi đó không trả về cho đến khi kết quả được nhận.

**Bất đồng bộ (Asynchronous):**

- Sau khi lời gọi được thực hiện, bên được gọi sẽ thông báo cho bên gọi khi kết quả được trả về, hoặc lời gọi sẽ được xử lý thông qua một hàm gọi lại (callback function).


### Giải thích bổ sung:

1. **Đồng bộ (Synchronous):**
    - Trong mô hình này, luồng gọi phải chờ đợi (blocking) cho đến khi thao tác hoàn tất.
    - Ví dụ: Một phương thức đọc tệp đồng bộ sẽ không trả về cho đến khi toàn bộ tệp được đọc xong.
    - Ưu điểm: Dễ hiểu và triển khai. Nhược điểm: Có thể lãng phí thời gian chờ.
2. **Bất đồng bộ (Asynchronous):**
    - Luồng gọi không cần chờ, có thể tiếp tục thực hiện các tác vụ khác trong khi chờ kết quả. Kết quả sẽ được xử lý sau qua thông báo hoặc callback.
    - Ưu điểm: Tăng hiệu suất, tận dụng tài nguyên tốt hơn. Nhược điểm: Mã phức tạp hơn.