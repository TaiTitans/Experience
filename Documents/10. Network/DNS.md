# DNS 

> -> DNS là viết tắt của cụm từ Domain Name System, mang ý nghĩa đầy đủ là **hệ thống phân giải tên miền.** Hệ thống cho phép thiết lập tương ứng giữa địa chỉ IP và **tên miền**.



Cách hoạt động:

> **DNS** là viết tắt của cụm từ Domain Name System, mang ý nghĩa đầy đủ là **hệ thống phân giải tên miền**. Hiểu một cách ngắn gọn nhất, DNS cơ bản là một hệ thống chuyển đổi các [**tên miền website**](https://www.matbao.net/ten-mien/dang-ky-ten-mien.html#kiem-tra-ten-mien) mà chúng ta đang sử dụng, ở dạng *www.tenmien.com* sang một địa chỉ IP dạng số tương ứng với **tên miền** đó và ngược lại.

Các bản ghi DNS:

- CNAME Record: Cho phép bạn tạo một tên mới, điều chỉnh trỏ tới tên gốc và đặt TTL. Tóm lại, tên miền chính muốn đặt một hoặc nhiều tên khác thì cần có bản ghi này. 
- A Record: Bản ghi này được sử dụng phổ biến để trỏ tên Website tới một địa chỉ IP cụ thể. Đây là bản ghi DNS đơn giản nhất, cho phép bạn thêm Time to Live (thời gian tự động tái lại bản ghi), một tên mới và Points To ( Trỏ tới IP nào).
- MX Record:  Với bản ghi này, bạn có thể trỏ Domain đến Mail Server, đặt TTL, mức độ ưu tiên (Priority). MX Record chỉ định Server nào quản lý các dịch vụ Email của tên miền đó.
- AAAA Record: Để trỏ tên miền đến một địa chỉ IPV6 Address, bạn sẽ cần sử dụng AAA Record. Nó cho phép bạn thêm Host mới, TTL,IPv6.
- **TXT Record**: Bạn cũng có thể thêm giá trị TXT, Host mới, Points To, TTL. Để chứa các thông tin định dạng văn bản của Domain, bạn sẽ cần đến bản ghi này.
- **SRV Record**: Là bản ghi dùng để xác định chính xác dịch vụ nào chạy Port nào. Đay là Record đặc biệt trong DNS. Thông qua nó, bạn có thể thêm Name, Priority, Port, Weight, Points to, TTL.
- **NS Record**: Với bản ghi này, bạn có thể chỉ định Name Server cho từng Domain phụ. Bạn có thể tạo tên Name Server, Host mới, TTL.

Các DNS Server bao gồm:

- Root Name Server
- Local Name Server

##### **Root Name Servers là gì?**

> Máy chủ ROOT có thể đưa ra các truy vấn (query) để tìm kiếm tối thiểu các thông tin về địa chỉ của các máy chủ tên miền authority thuộc lớp top-level-domain chứa tên miền muốn tìm.

##### **Local Name Servers là gì?**

> Server này chứa thông tin, để tìm kiếm **máy chủ tên miền** lưu trữ cho các tên miền thấp hơn. Nó thường được duy trì bởi các doanh nghiệp, các nhà cung cấp dịch vụ Internet (ISPs).

More : [DNS là gì? Tầm quan trọng của DNS trong thế giới mạng (matbao.net)](https://wiki.matbao.net/dns-la-gi-tam-quan-trong-cua-dns-trong-the-gioi-mang/)