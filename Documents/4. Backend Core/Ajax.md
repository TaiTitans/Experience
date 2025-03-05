# AJAX

> AJAX viết tắt từ Asynchronous JavaScript and XML, là bộ công nghệ giúp tạo ra các web động hay các ứng dụng giàu tính Internet, cho phép tăng tốc độ ứng dụng web bằng cách cắt nhỏ dữ liệu và chỉ hiển thị những gì cần thiết, thay vì tải đi tải lại toàn bộ trang web, làm như vậy trang của bạn sẽ muợt và đẹp hơn. AJAX không phải một công nghệ đơn lẻ mà là sự kết hợp một nhóm công nghệ với nhau

Trong đó:

- HTML (hoặc XHTML) và CSS
- mô hình DOM (Document Object Model) được thực hiện thông qua JavaScript, nhằm hiển thị thông tin động và tương tác với những thông tin được hiển thị.
- Đối tượng XMLHttpRequest để trao đổi dữ liệu một cách không đồng bộ với máy chủ web. 
- XML thường là định dạng cho dữ liệu truyền, mặc dầu bất cứ định dạng nào cũng có thể dùng, bao gồm HTML định dạng trước, văn bản thuần (plain text), JSON và ngay cả EBML.

##### Cách Ajax hoạt động:

Trong mô hình truyền thống, khi client gửi một HTTP Request lên server, server phải thực hiện các khâu xử lý và trả về một trang HTML hoàn chỉnh cho client. Tuy nhiên, phương pháp này có nhược điểm là bất tiện và tốn thời gian, vì khi server đang xử lý, người dùng phải chờ đợi.

Để khắc phục nhược điểm này, phương pháp AJAX được giới thiệu. AJAX cho phép tạo ra một trung gian (Ajax Engine) giữa client và server để giảm quá trình "đi lại" thông tin và tương tác. Thay vì tải lại toàn bộ trang, AJAX chỉ nạp những thông tin được thay đổi, giữ nguyên các phần khác. Khi sử dụng AJAX, người dùng không nhìn thấy cửa sổ trắng và biểu tượng đồng hồ cát, và thời gian chờ có thể được thay thế bằng thông báo "đang tải dữ liệu".

Với AJAX, khi có một thay đổi trên trang web, chỉ cần cập nhật phần thông tin thay đổi mà không cần tải lại toàn bộ trang. Điều này giúp tạo ra các giao dịch trơn tru và nhanh chóng.

Tóm lại, AJAX là một cơ chế trung gian giữa client và server, giúp giảm thời gian tương tác và tải lại trang web toàn bộ, bằng cách chỉ cập nhật những phần thay đổi.

**Ưu điểm của Ajax:**

Nhanh chóng và tương tác mượt mà: Ajax cho phép tải lên và tải xuống các dữ liệu và tài nguyên chỉ khi cần thiết, giúp tăng tốc độ tải trang và cải thiện trải nghiệm người dùng.

Tương tác không đồng bộ: Ajax cho phép tương tác không đồng bộ với server, tức là người dùng có thể tiếp tục tương tác với trang web trong khi các yêu cầu Ajax đang được gửi và nhận kết quả từ server.

Tiết kiệm băng thông: Chỉ cần cập nhật phần thông tin thay đổi, Ajax giúp giảm lượng dữ liệu phải truyền giữa client và server, giúp tiết kiệm băng thông mạng.

Tăng tính tương tác và đáp ứng: Ajax cho phép thay đổi nội dung trang web mà không cần tải lại trang hoàn toàn, giúp tạo ra các giao diện tương tác đáp ứng và linh hoạt.

Tích hợp dễ dàng: Ajax có thể dễ dàng tích hợp vào các ngôn ngữ lập trình web phổ biến như JavaScript, HTML và CSS.

**Nhược điểm của Ajax:**

Phụ thuộc vào JavaScript: Ajax cần sử dụng JavaScript để gửi và nhận dữ liệu thông qua các yêu cầu HTTP, điều này có nghĩa là nếu người dùng tắt JavaScript, chức năng Ajax sẽ không hoạt động.

Quản lý lỗi phức tạp: Ajax có thể gây ra những thách thức về quản lý lỗi và xử lý ngoại lệ. Vì các yêu cầu Ajax không gửi lại toàn bộ trang, việc xử lý lỗi và thông báo lỗi trở nên phức tạp hơn.

Vấn đề về bảo mật: Do Ajax cho phép tương tác không đồng bộ với server, việc kiểm soát và bảo mật dữ liệu trở nên quan trọng. Sự thiếu sót trong việc xử lý bảo mật có thể gây ra các vấn đề về an ninh.

Khả năng tương thích: Một số trình duyệt cũ không hỗ trợ đầy đủ các tính năng Ajax, điều này có thể gây khó khăn trong việc phát triển và duy trì tính tương thích trên các nền tảng khác nhau.

[Mastering AJAX in JavaScript: A Beginner’s Guide with Examples | by JavaScript-World | Medium](https://medium.com/@JavaScript-World/mastering-ajax-in-javascript-a-beginners-guide-with-examples-6111aa53e690)