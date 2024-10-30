
---
12 yếu tố này bao gồm:

1. **Codebase**: Một codebase nằm trong source control, deploy lên nhiều nơi
2. **Dependencies**: Depedencies như library/framework/extension phải được ghi rõ ràng, tách biệt từng app
3. **Config**: Lưu trữ thiết lập vào biến môi trường (environment variable)
4. **Backing services**: Xem các service đi kèm (database, API, …) như là resource của app
5. **Build, release, run**: Tách riêng quá trình release, build và run
6. **Processes**: App nên chạy dưới dạng stateless processes
7. **Port binding**: Một service có thể được access thông qua 1 port cố định
8. **Concurrency**: Một app nên được chia tách thành nhiều process nhỏ để tăng concurrency
9. **Disposability**: Process của web app nên sống nhanh, chết vội, để có thể dễ dàng chạy/kill process nhanh chóng
10. **Dev/prod parity**: Các [môi trường dev/staging/production](https://toidicodedao.com/2019/07/02/environment-trong-lap-trinh/) nên giống nhau hết sức có thể
11. **Logs**: Logs nên được viết ra dạng stream ở stdout
12. **Admin Processes**: Một số task dạng admin (tạo database, fix dữ liệu) nên được chạy trong cùng môi trường với app đang chạy


### **1. [Codebase](https://12factor.net/codebase): Một source code trong source control, có nhiều deploy**

Mỗi app chỉ nên có 1 source code (codebase) duy nhất, nằm trong source control như Git/SVN.

Mỗi khi code được build và [chạy trên một môi trường nào đó,](https://toidicodedao.com/2019/07/02/environment-trong-lap-trinh/) ta gọi nó là một bản deploy.

Code chạy trên môi trường Production gọi là Production deploy, chạy trên Staging gọi là Staging deploy. Deploy trên staging có thể có nhiều commit hơn, nằm ở branch khác Production, nhưng 2 code này đều nằm trong 1 source control.
Nhờ vậy, khi có developer mới gia nhập, ta có thể dễ dàng lấy và đọc source code từ source control.

Gần đây, khi [kiến trúc microservice](https://toidicodedao.com/2017/02/21/tong-quan-micro-service/) đang thịnh hành, ta có thể dùng [monorepo](https://medium.com/@mattklein123/monorepos-please-dont-e9a279be011b) – 1 repo chứa toàn bộ source code của các service. Hoặc có thể coi mỗi service là 1 app, lưu source code của service đó vào 1 repo riêng.

### **2. [Dependencies](https://12factor.net/dependencies): Dependencies rõ ràng, tách biệt**

Cho các bạn chưa biết, depedencies tức là các package/thư viện/tool của bên thứ 3, mà app của bạn cần có để chạy được.

Một app cần phải **thiết lập rõ ràng những dependency mà nó sử dụng, không lệ thuộc vào các dependency có sẵn của hệ thống.** Các dependencies này phải tách biệt trong từng app.

Giả sử thế này, bạn đang viết ứng dụng nhận dạng JAV Idol, sử dụng thư viện IdolRec ver 2.1. Thằng bạn của bạn cũng viết ứng dụng nhận diện EU Idol, sử dụng IdolRec ver 1.3:

- Cách thức đúng ở đây là: app sẽ ghi rõ mình dùng thư viện IdolRec + version. Thư viện này sẽ bỏ vào thư mục của app, chạy trong app.
- Nếu ứng dụng phụ thuộc **vào dependencies ngầm trong máy**, sẽ dễ dẫn đến tình trạng “It works on my machine” => Chạy được trên máy của dev, nhưng lên Production thì tèo…
- Nếu không ghi rõ, tách biệt dependeices, khi 2 ứng dụng chạy trên 1 máy sẽ **bị conflict** vì 2 thư viện khác nhau. Hoặc có thể … không chạy được vì máy kia chưa xài IdolRec

Hiện tại, đa phần các ngôn ngữ lập trình đều có cách [package manager](https://toidicodedao.com/2017/05/30/package-manager-la-gi/) hỗ trợ chuyện này (_NodeJS_ thì lưu dependencies vào file package.json, _Ruby_ thì có bundler lưu depedencies vào gemfile, C# thì lưu vào _Web.config hoặc App.config_ ….).

Nhờ vậy, khi ta deploy ứng dụng lên server, hoặc developer mới gia nhập pull code về, ta chỉ cần chạy _npm install_ hoặc _bundle install_ để tải thư viện về, là ứng dụng có thể chạy bình thường, không cần cài dependencies ngầm.

