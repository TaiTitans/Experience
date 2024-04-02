# Server Side Rendering - Client Side Rendering

Server- Side Rendering (SSR) → Kết xuất phía máy chủ:

→ Máy chủ sẽ gửi một trang HTML mới cho mỗi đường dẫn được truy xuất.

SPA → Máy chủ gửi duy nhất một trang HTML cho mọi đường dẫn được truy xuất.

![](img\image-20240402161047841.png)

**Server Side Rendering (SSR):**

- Trong SSR, các trang web được tạo ra hoàn toàn trên máy chủ trước khi được gửi đến trình duyệt của người dùng.
- Khi người dùng yêu cầu một trang, máy chủ tạo ra trang web hoàn chỉnh và gửi nó đến trình duyệt của người dùng.
- Thông thường, SSR được sử dụng cho các ứng dụng web yêu cầu SEO cao vì các trang được tạo ra với dữ liệu đầy đủ từ đầu, các công cụ tìm kiếm có thể dễ dàng tìm kiếm và lập chỉ mục chúng.

**Client Side Rendering (CSR):**

- Trong CSR, trình duyệt của người dùng tải một trang HTML trống ban đầu và sau đó sử dụng JavaScript để tải nội dung và hiển thị trang web.
- Thay vì gửi một trang web hoàn chỉnh từ máy chủ, máy chủ chỉ cung cấp dữ liệu cần thiết (thường là dưới dạng API) và mã JavaScript.
- Trình duyệt sau đó sử dụng mã JavaScript để tạo ra và hiển thị nội dung trang web dựa trên dữ liệu nhận được từ máy chủ.
- CSR thường được sử dụng trong các ứng dụng web động, nơi mà dữ liệu thay đổi thường xuyên và cần cập nhật mà không cần tải lại trang.

Ví dụ:

- SSR: Sử dụng templating engine như EJS để viết HTML trực tiếp trong ứng dụng Node.js.
- CSR: Tạo một project frontend với React hoặc Vue.js trong một thư mục riêng biệt và sử dụng API từ ứng dụng Node.js để lấy dữ liệu và hiển thị trên client.
