
---

CORS (Cross-Origin Resource Sharing) là một cơ chế bảo mật được triển khai trong trình duyệt web, nhằm kiểm soát cách các tài nguyên (như JavaScript, CSS, hình ảnh) trên một trang web được truy cập từ các nguồn khác (cross-origin).

Cơ chế hoạt động của CORS:

1. Khi một trang web (origin) cố gắng truy cập tài nguyên từ một nguồn khác (cross-origin), trình duyệt sẽ gửi một yêu cầu OPTIONS đến máy chủ chứa tài nguyên.
2. Máy chủ sẽ trả về các header cho phép hoặc từ chối yêu cầu truy cập, dựa trên cấu hình CORS.
3. Trình duyệt sẽ kiểm tra các header này và quyết định cho phép hoặc từ chối yêu cầu truy cập.

Các header CORS quan trọng:

- Access-Control-Allow-Origin: Chỉ định các nguồn được phép truy cập tài nguyên.
- Access-Control-Allow-Methods: Chỉ định các phương thức HTTP được phép sử dụng.
- Access-Control-Allow-Headers: Chỉ định các header được phép sử dụng.

Ưu điểm của CORS:

- Bảo vệ khỏi các cuộc tấn công Cross-Site Scripting (XSS) và Cross-Site Request Forgery (CSRF).
- Cho phép các ứng dụng web truy cập tài nguyên từ các nguồn khác một cách có kiểm soát.
- Giúp ngăn chặn việc truy cập trái phép vào các tài nguyên nhạy cảm.

Triển khai CORS:

- Trên phía máy chủ: Cấu hình các header CORS phù hợp với yêu cầu bảo mật.
- Trên phía client (trình duyệt): Sử dụng các API như Fetch hoặc XMLHttpRequest để tuân thủ các quy tắc CORS.

Lưu ý:

- CORS chỉ áp dụng cho các trình duyệt web, không áp dụng cho các API được gọi trực tiếp.
- Cấu hình CORS cần được thực hiện cẩn thận để tránh các vấn đề về bảo mật và khả năng truy cập.

Tóm lại, CORS là một cơ chế bảo mật quan trọng, giúp kiểm soát việc truy cập tài nguyên từ các nguồn khác, góp phần tăng cường bảo mật cho các ứng dụng web.