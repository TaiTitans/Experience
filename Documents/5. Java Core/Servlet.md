
---

	Servlet được sử dụng để tạo ra ứng dụng web ( nằm ở phía máy chủ và tạo ra trang web động )

Sử dụng servlet để thu nhập thông tin đầu vào người dùng thông qua trang web, hiển thị các bản ghi từ một cơ sở dữ liệu hoặc từ một nguồn khác.

Servlet mạnh mẽ và có khả năng mở rộng.

## Servlet là gì?

Servlet có thể được mô tả bằng nhiều cách, tùy thuộc vào ngữ cảnh:

- Servlet là một công nghệ được sử dụng để tạo ra ứng dụng web.
- Servlet là một API cung cấp các interface và lớp bao gồm các tài liệu.
- Servlet là một thành phần web được triển khai trên máy chủ để tạo ra trang web động.

Có nhiều interface và các lớp trong API servlet như Servlet, GenericServlet, HttpServlet, ServletRequest, ServletResponse, ...


Kiến trúc Servlet:

![](../../Assets/Images/kien-truc-servlet.png)

## Nhiệm vụ của Servlet:

- Đọc dữ liệu rõ ràng do khách hàng (trình duyệt) gửi. Điều này bao gồm một mẫu HTML trên một trang Web hoặc nó cũng có thể đến từ một applet hoặc một chương trình khách hàng HTTP tùy chỉnh.
- Đọc dữ liệu yêu cầu HTTP ẩn được gửi bởi khách hàng (trình duyệt). Điều này bao gồm các cookie, loại phương tiện truyền thông và các chương trình nén mà trình duyệt hiểu được, v.v.
- Xử lý dữ liệu và tạo ra các kết quả. Quá trình này có thể yêu cầu nói chuyện với một cơ sở dữ liệu, thực hiện một cuộc gọi RMI hoặc CORBA, gọi một dịch vụ Web, hoặc tính trực tiếp phản hồi.
- Gửi dữ liệu rõ ràng (tức là tài liệu) tới khách hàng (trình duyệt). Tài liệu này có thể được gửi bằng nhiều định dạng, bao gồm văn bản (HTML hoặc XML), nhị phân (hình ảnh GIF), Excel, v.v ...
- Gửi phản hồi HTTP ẩn cho khách hàng (trình duyệt). Điều này bao gồm nói với trình duyệt hoặc các trình khách khác loại tài liệu đang được trả về (ví dụ, HTML), thiết lập cookie và các tham số bộ nhớ đệm, và các tác vụ khác.
### Gói Servlet Java

Java Servlets là các lớp Java chạy bởi một máy chủ web có một trình thông dịch hỗ trợ đặc tả Java Servlet.

Servlets có thể được tạo ra bằng cách sử dụng các **gói javax.servlet** và **javax.servlet.http** , đây là một phần chuẩn của phiên bản Enterprise của Java, một phiên bản mở rộng của thư viện lớp Java hỗ trợ các dự án phát triển quy mô lớn.

Các lớp này thực hiện các đặc tả Java Servlet và JSP.

## Nhược điểm của Servlet

Phải viết code java + html trong cùng file.

## Tại sao nên biết Servlet

Nếu bạn đang đi theo hướng java web thì bạn nên biết về servlet vì nó là core của java web.
