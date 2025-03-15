
---
**Câu trả lời không phải lúc nào cũng chắc chắn.**  
Có hai trường hợp mà khối finally sẽ không được thực thi:

1. **Chương trình không thực thi đến khối try:** Nếu mã không bao giờ đến được khối try (ví dụ: do thoát chương trình sớm), thì khối finally tương ứng sẽ không chạy.
2. **Nếu một luồng bị gián đoạn hoặc bị dừng (killed) trong khi thực thi khối try hoặc catch:** Trong trường hợp này, khối finally tương ứng có thể không được thực thi.

Ngoài ra, còn có những trường hợp cực đoan hơn:

- Khi một luồng đang chạy khối try hoặc catch mà đột nhiên bị crash (sập chương trình) hoặc mất điện, khối finally chắc chắn sẽ không được thực thi.