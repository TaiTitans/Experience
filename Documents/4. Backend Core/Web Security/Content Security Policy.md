
---
Content Security Policy (CSP) là một tính năng bảo mật web, giúp giảm thiểu rủi ro từ các cuộc tấn công như Cross-Site Scripting (XSS) và các loại tấn công khác. CSP hoạt động bằng cách chỉ định các nguồn tin cậy mà trình duyệt được phép tải tài nguyên từ đó.

Cách thức hoạt động của CSP:

1. Trang web cấu hình CSP bằng cách thêm một HTTP header hoặc một `<meta>` tag vào trang web.
2. Header CSP chỉ định các nguồn tin cậy (như domains, CDN, etc.) mà trình duyệt được phép tải tài nguyên từ đó.
3. Khi trình duyệt tải tài nguyên (như JavaScript, CSS, hình ảnh, v.v.), nó sẽ kiểm tra xem nguồn tài nguyên đó có nằm trong danh sách tin cậy hay không.
4. Nếu tài nguyên không nằm trong danh sách tin cậy, trình duyệt sẽ chặn việc tải tài nguyên đó và ghi lại sự kiện vào bảng điều khiển của nhà phát triển.

Các chính sách CSP phổ biến:

- `default-src 'self'`: Chỉ cho phép tải tài nguyên từ nguồn gốc hiện tại.
- `script-src 'self' https://example.com`: Chỉ cho phép tải script từ nguồn gốc hiện tại và domain `example.com`.
- `img-src 'self' data:`: Chỉ cho phép tải hình ảnh từ nguồn gốc hiện tại và từ data URIs.
- `object-src 'none'`: Không cho phép tải bất kỳ tài nguyên dạng object/embed nào.

Ưu điểm của CSP:

- Giảm thiểu rủi ro từ các cuộc tấn công XSS.
- Cung cấp một lớp bảo vệ bổ sung cho các ứng dụng web.
- Giúp phát hiện và ngăn chặn các tấn công liên quan đến nội dung không đáng tin cậy.
- Cung cấp một cơ chế để kiểm soát và giám sát việc tải tài nguyên trên trang web.

Triển khai CSP:

- Trên phía máy chủ: Thêm header CSP hoặc `<meta>` tag vào trang web.
- Trên phía client: Trình duyệt sẽ tự động thực hiện kiểm tra và chặn các tài nguyên không đáng tin cậy.

Tóm lại, Content Security Policy (CSP) là một công cụ bảo mật web quan trọng, giúp giảm thiểu các rủi ro từ các cuộc tấn công như XSS bằng cách kiểm soát các nguồn tài nguyên được phép tải trên trang web.