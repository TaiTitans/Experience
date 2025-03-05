

---

Load balancing là một thành phần quan trọng của cơ sở hạ tầng thường được sử dụng để cải thiện hiệu suất và độ tin cậy của các trang web, các ứng dụng, cơ sở dữ liệu và các dịch vụ khác bằng cách phân phối khối lượng công việc trên nhiều máy chủ.

Một cơ sở hạ tầng web không có Load balancing có thể trông giống như sau:

![](https://images.viblo.asia/9ef3bb43-3d54-49fa-90be-d464751e5b47.png)

Trong ví dụ này, người sử dụng kết nối trực tiếp đến máy chủ web, tại [yourdomain.com](http://yourdomain.com/). Nếu máy chủ web duy nhất này down, người sử dụng sẽ không thể truy cập vào trang web. Ngoài ra, nếu nhiều người dùng cố gắng truy cập vào máy chủ cùng một lúc thì nó không thể xử lý tải, họ có thể gặp thời gian tải chậm hoặc không thể kết nối.

Đây là điểm có thể khắc phục bằng một load balancer và ít nhất một máy chủ web bổ sung trên backend. Thông thường, tất cả các máy chủ phụ trợ sẽ cung cấp nội dung giống hệt nhau để người dùng nhận được nội dung phù hợp bất kể là máy chủ nào đáp ứng.
![](https://images.viblo.asia/65fad7fe-c1b6-4798-8bdd-94232331f8f8.png)

**Những loại giao thức load balancers có thể xử lý:**

Quản trị Load balancer tạo quy định chuyển tiếp đối với bốn loại giao thức chính:

- HTTP - Chuẩn HTTP balancing chỉ đạo yêu cầu dựa trên cơ chế HTTP chuẩn. Load Balancer đặt X-Forwarded-For, X-Forwarded-Proto, và tiêu đề X-Forwarded-Port để cung cấp cho các thông tin backends về các yêu cầu ban đầu.
    
- HTTPS - HTTPS balancing với các chức năng tương tự như HTTP balancing, với sự bổ sung của mã hóa. Mã hóa được xử lý theo một trong hai cách: hoặc là với passthrough SSL duy trì mã hóa tất cả con đường đến backend hoặc chấm dứt SSL mà đặt gánh nặng giải mã vào load balancer nhưng gửi lưu lượng được mã hóa đến back end.
    
- TCP - Đối với các ứng dụng không sử dụng HTTP hoặc HTTPS, lưu lượng TCP cũng có thể được cân bằng. Ví dụ, lượng truy cập vào một cụm cơ sở dữ liệu có thể được lan truyền trên tất cả các máy chủ.
    
- UDP - Gần đây, một số load balancer đã thêm hỗ trợ cho cân bằng tải giao thức internet lõi như DNS và syslogd sử dụng UDP.
    

Những quy tắc chuyển tiếp sẽ xác định các giao thức và cổng vào load balancer và bản đồ chúng đến các giao thức và cổng load balancer sẽ sử dụng để định tuyến lưu lượng trên backend.

---
**Làm thế nào để load balancer chọn máy chủ backend?**

Load balancers chọn máy chủ để chuyển tiếp yêu cầu dựa trên sự kết hợp của hai yếu tố. Lần đầu tiên sẽ đảm bảo rằng bất kỳ máy chủ được lựa chọn có thể thực sự đáp ứng yêu cầu và sau đó sử dụng một quy tắc được cấu hình sẵn để lựa chọn trong số đó.

**Health Checks**

Load balancer chỉ chuyển tiếp lưu lượng đến "healthy" backend server. Để theo dõi sức khỏe của một backend server, kiểm tra sức khỏe thường xuyên bằng cách cố gắng kết nối với backend server sử dụng giao thức và cổng được định nghĩa bởi các quy tắc chuyển tiếp để đảm bảo rằng các máy chủ đang lắng nghe. Nếu một máy chủ không kiểm tra sức khỏe, và do đó không thể phục vụ yêu cầu, nó sẽ tự động loại bỏ khỏi vùng chứa, và request sẽ không được chuyển tiếp đến nó cho đến khi nó đáp ứng việc kiểm tra sức khỏe một lần nữa.

**Các thuật toán load balancer**

Các thuật toán load balancer được sử dụng xác định của máy chủ lành mạnh trên backend sẽ được lựa chọn. Một số các thuật toán thường được sử dụng là:

- Round Robin - Round Robin có nghĩa là các máy chủ sẽ được lựa chọn theo tuần tự. Bộ load balancer sẽ chọn máy chủ đầu tiên trong danh sách của mình đối với yêu cầu đầu tiên, sau đó di chuyển xuống trong danh sách theo thứ tự, bắt đầu lại ở đầu trang khi đi đến cuối cùng.
    
- Least Connections - load balancer sẽ chọn máy chủ với các kết nối ít nhất.
    
- Source - Với các thuật toán mã nguồn, load balancer sẽ chọn máy chủ để sử dụng dựa trên một hash của IP nguồn của yêu cầu, chẳng hạn như địa chỉ IP của người truy cập. Phương pháp này đảm bảo rằng một người dùng cụ thể sẽ luôn kết nối với cùng một máy chủ.
    

Các thuật toán có người quản lý khác nhau tùy thuộc vào công nghệ load balancer sử dụng.


**Làm thế nào để load balancer xử lý trạng thái?**

Một số ứng dụng yêu cầu người dùng tiếp tục kết nối đến cùng một backend server. Một thuật toán mã nguồn tạo ra một mối quan hệ dựa trên thông tin IP khách hàng. Một cách khác để đạt được điều này ở mức ứng dụng web là thông qua sticky sessions, nơi load balancer đặt một cookie và tất cả các requests từ sessions hướng đến một máy chủ vật lý.

**Load balancer dự phòng**

Để loại bỏ việc load balancer như một điểm truy cập duy nhất, một load balancer thứ hai có thể được kết nối với cái đầu tiên để tạo thành một cụm. Mỗi load balancer là đều có khả năng phát hiện lỗi và phục hồi.

![](https://images.viblo.asia/5f72fc77-5ba3-4dc5-9d73-8ff37f33fff9.png)

Trong trường hợp load balancer chính bị lỗi, DNS phải đưa người dùng đến các bộ load balancer thứ hai. Bởi vì thay đổi DNS có thể mất một lượng thời gian đáng kể để được tải lên Internet và để làm cho chuyển đổi dự phòng này tự động, nhiều quản trị viên sẽ sử dụng hệ thống, cho phép linh hoạt địa chỉ IP Remapping, chẳng hạn như các floating IPs. Theo yêu cầu địa chỉ IP Remapping giúp loại bỏ các vấn đề tuyên truyền, bộ nhớ đệm vốn có trong những thay đổi DNS bằng cách cung cấp một địa chỉ IP tĩnh có thể được dễ dàng ánh xạ lại khi cần thiết. Tên miền có thể duy trì liên kết với các địa chỉ IP, trong khi các địa chỉ IP của chính nó được di chuyển giữa các máy chủ.

Đây là cách một cơ sở hạ tầng sử dụng Floating IPs có thể xem xét:

![](https://images.viblo.asia/7f9b9acf-6ba0-4f4d-96c1-272cc7beb95c.gif)

---
Để load balancer xử lý trạng thái, có một vài cách như sau:

1. Session Persistence (Sticky Sessions):
    
    - Load balancer lưu trữ thông tin về phiên làm việc của người dùng và chuyển yêu cầu của người dùng đó đến cùng một máy chủ back-end.
    - Giúp duy trì tính toàn vẹn của dữ liệu phiên làm việc.
2. Session Replication:
    
    - Load balancer sao chép dữ liệu phiên làm việc của người dùng sang các máy chủ back-end.
    - Khi một máy chủ bị lỗi, người dùng vẫn có thể tiếp tục sử dụng dữ liệu phiên làm việc từ các máy chủ khác.
3. Database/Distributed Cache:
    
    - Lưu trữ dữ liệu phiên làm việc trong một hệ thống lưu trữ trung tâm như cơ sở dữ liệu hoặc bộ nhớ cache phân tán.
    - Load balancer có thể truy cập vào dữ liệu phiên làm việc từ bất kỳ máy chủ back-end nào.

Load balancer dự phòng là một cơ chế giúp duy trì hoạt động của hệ thống khi một load balancer chính bị lỗi hoặc không khả dụng. Các cách triển khai load balancer dự phòng gồm:

1. Failover Load Balancer:
    
    - Có một load balancer chính và một load balancer phụ.
    - Khi load balancer chính bị lỗi, load balancer phụ sẽ tự động tiếp quản và xử lý lưu lượng truy cập.
2. Active-Passive Load Balancing:
    
    - Có một load balancer active và một load balancer passive.
    - Load balancer active xử lý toàn bộ lưu lượng truy cập.
    - Khi load balancer active gặp sự cố, load balancer passive sẽ được kích hoạt để tiếp quản.
3. Active-Active Load Balancing:
    
    - Có hai hoặc nhiều load balancer cùng hoạt động song song.
    - Tất cả load balancer đều xử lý một phần lưu lượng truy cập.
    - Khi một load balancer gặp sự cố, các load balancer còn lại sẽ chia sẻ thêm lưu lượng.