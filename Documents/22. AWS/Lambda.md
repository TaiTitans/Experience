
---
1. AWS Lambda là một dịch vụ **serverless** cho phép chạy code mà không cần quản lý máy chủ. Lambda tự động mở rộng theo nhu cầu, chỉ chạy khi có sự kiện kích hoạt và tính phí dựa trên thời gian thực thi. Nó **hỗ trợ** nhiều ngôn ngữ như **Python, Node.js, Java, Go, v.v.** Lambda thường dùng để **xử lý sự kiện từ API Gateway, S3, DynamoDB, hoặc các dịch vụ AWS khác**
    
    - Tại sao lại là Lambda?
    
    Trước tiên ta sẽ đặt EC2 và Lambda lên bàn cân để phân tích EC2 + **Là một máy chủ ảo (Virtual Host)** trên điện toán đám mây + **Bị giới hạn bởi RAM và CPU**: Nếu tăng RAM hoặc CPU lên thì tiền sẽ tăng + **Chạy liên tục**: bạn phải **tắt bằng tay** nhưng vẫn **sẽ bị charge tiền EBS** attached với EC2 bạn vừa tạo + **Scale**: Khi Scale ta sẽ có Auto Scaling Group (ASG) nhưng điều đó có nghĩa là phải đặt thêm **rule** cho ASG để có thể thêm hoặc xóa EC2 đó trong ASG (rule thường sẽ là: scale lúc mấy giờ, scale dựa trên thông số các EC2 trong ASG đó như là RAM/CPU đang hoạt động hay scale vào 1 ngày mà ta tự đặt ra etc...)  
    Lambda + **Là một hàm ảo (Virtual Function)** trên điện toán đám mây: bạn chỉ cần viết code và còn lại AWS sẽ lo hết từ phần hạ tầng và môi trường chạy + **Bị giới hạn thời gian thực thi**: một lambda function chỉ có thể **thực thi tối đa 15 phút**, sau 15P nếu chưa thực thi xong **AWS sẽ tự hủy function đó** + **Chạy theo yêu cầu**: khi ta không sử dụng Lambda thì Lambda sẽ không chạy và bạn chỉ bị tính phí hàm của được chạy + **Tự động scale**: Nếu chúng ta cần xử lý một tác vụ đồng thời cao như request từ người dùng thì AWS sẽ tự động cung cấp thêm cho bạn Lambda function để xử lý tác vụ đó
    
    
    - Lợi ích khi dùng Lambda
    
    **+ **Giá cả siêu rẽ**: Bạn chỉ phải **trả mỗi khi request 1 lần và thời gian thực thi của hàm đó** (được tính theo đơn vị 1ms) - **$0.2/1 triệu request và $1.0 cho 600k GB-Seconds**, **Free tier** thì AWS cung cấp **1 TRIỆU Lambda request và 400k GB-Seconds thời gian thực thi** VD: Nếu một hàm **Lambda** sử dụng **1GB RAM**, bạn sẽ có **tổng cộng 400k giây** thực thi miễn phí VD: Nếu một hàm **Lambda** sử dụng **128MB RAM**, bạn sẽ có tổng cộng **3,2 triệu giây thực thi (400k Giây miễn phí còn lại thì tính tiền)** + **Giám sát**: Ta có thể **giám sát** thông số của **Lambda** function một cách **dễ dàng** thông qua **CloudWatch** + **Khả năng cung cấp**: Ta có thể **linh hoạt** trong việc **cung cấp tài nguyên** cho **Lambda** function lên tới **10GB RAM**, đồng thời việc cung cấp **RAM** càng nhiều thì **CPU và Network** cũng sẽ cải thiện chất lượng xử lý.**
    
    - Khả năng tương tích nhiều ngôn ngữ
    
    **+ **Lambda hỗ trợ cho nhiều ngôn ngữ như**: Java, Py thông, Gô lăng, NodeJS(JavaScript), .Net(C#), Rust + **Không có ngôn ngữ bạn tìm**: Nếu không thấy bạn có thể tìm trong Custom Runtime API (Do cộng đồng hỗ trợ)**
    
    - Sử dụng container image trên Lambda
    
    **(Advance) + Container image của bạn phải triển khai trên **Lambda Runtime API** để có thể hoạt động trên AWS Lambda. + AWS Lambda hỗ trợ container image có dung lượng tối đa **10 GB**. + Bạn có thể sử dụng các **runtime có sẵn của AWS** hoặc tự **xây dựng runtime tùy chỉnh** bằng cách sử dụng **AWS Lambda Runtime API**. + Container image có thể được lưu trữ trên **Amazon Elastic Container Registry (ECR)** hoặc các container registry khác được hỗ trợ. + AWS Lambda chỉ hỗ trợ các kiến trúc **x86_64 và ARM64**, nên bạn cần chọn đúng kiến trúc khi xây dựng image. + Khi sử dụng container trên Lambda, bạn vẫn có thể tận dụng các tính năng như **provisioned concurrency, IAM permissions, và VPC integration**.
    
    - Tích hợp với nhiều service
    
    **Lambda được tích hợp rất nhiều với các service của AWS, nhưng mình chỉ sẽ nói những service nổi tiếng (Khá chắc có trong Exam) + **API Gateway**: Là một dịch vụ tạo ra **REST API** và có thể tích hợp gọi tới Lambda + **Kinesis**: Là một dịch vụ **stream data thời gian thực** đến cho các service hoặc doanh nghiệp khác, **kinesis** có thể sử dụng **Lambda** để **chuyển đổi dữ liệu** (cắt xén hay thay đổi) **trong quá trình vận chuyển data** + **DyamoDB**: Là một **cơ sở dữ liệu NO SQL** có tích hợp **stream, Cache và TTL**, khi **một sự kiện** CRUD hay gì đó **tác động lên DyamoDB** thì DyamoDB sẽ **phát ra một event** và ta có thể **tận dụng event** đó để tích hợp **với Lambda** và gọi + **S3**: Là một dịch vụ lưu trữ thông tin dạng Object **Simple Storage Service** khi **S3 có sự kiện** CRUD thì ta có thể tích hợp với **Lambda** function để **thay đổi một số thứ trong file** + **CloudWatch Log**: Là một dịch vụ **lưu trữ và streaming log**, ta có thể **tích hợp Lambda để ghi log vào CloudWatch Log hoặc xử lý stream** + **SNS**: Là một dịch vụ chuyển phát thông báo đơn giản **(Simple Notification Service)** và có thể **gửi thông báo đến nhiều service** hoặc là **các thiết bị của người khác** và SNS có tích hợp với **Lambda để xử lý khi có notification** + **SQS**: Là một dịch vụ queue đơn giản **(Simple Queue Service)**, **Lambda** có thể được **kích hoạt bởi hàng đợi** SQS(event source) hoặc **gửi message đến SQS** + **Cognito**: Là một dịch vụ bảo mật cung cấp thông tin, Lambda có thể được dùng với Cognito để **xử lý authentication, authorization và user management** . VD: Xác thực thông tin trước khi tạo user (PreSignUp trigger).
    
    
    - Giới hạn AWS Lambda cần biết - theo khu vực (Có trong exam)
    
    ****Thực thi**: + **Phân bổ bộ nhớ**: 128 MB – 10GB (tăng theo mức 1 MB) + **Thời gian thực thi tối đa**: 900 giây (15 phút) + **Biến môi trường**: (4 KB) + **Dung lượng đĩa trong "container chức năng" (trong /tmp)** : 512 MB đến 10GB + **Số lượng thực thi đồng thời**: 1000 (có thể được tăng lên) **Triển khai**: + **Kích thước triển khai hàm Lambda (dạng nén .zip)**: 50 MB + **Kích thước triển khai giải nén (code + dependencies)**: 250 MB + **Kích thước triển khai giải nén (code + dependencies)**: 4 KB + Có thể sử dụng thư mục /tmp để tải các tệp khác khi khởi động**
    
    - Deploy Lambda
    
    ****Có 2 cách deploy Lambda function: Deploy ngoài VPC, Deploy trong VPC**
    
    - **Deploy ngoài VP**C: Ta chỉ có thể tương tác được với các Service cấp độ cao nhất **Region**(S3, SQS, DynamoDB, etc...) và không thể đi vào các VPC để tương tác (Sử dụng VPC ENDPOINT để giải quyết)
    - **Deploy trong VPC**: Deploy trong VPC thì ta có thể tương tác được với các services trong VPC nhưng ngoài VPC thì lại không được (VPC ENDPOINT), nhưng xử dụng Lambda trong VPC sẽ gặp một lỗi khá phổ biển mà nhiều doanh nghiệp đã dính phải (ENI Exhausion)
    
    + Để deploy Lambda ta sẽ vào bên trong AWS console và kiếm Lambda trên thanh tìm kiếm + **Trong bảng Lambda** thì bên trái là **thanh điều hướng** còn bên phải là **bảng thông số dữ liệu** + Bên phải sẽ có nút **create function** + Ta sẽ có **3 loại** Lambda function là Author from sratch(Tự mình viết), Use a blueprint(Sử dụng của người khác), Container Image(Sử dụng với container image) + Khi vào **Author from sratch** thì ta cần phải thên **tên function**, **runtime gì(Gia va, Gô lăng, etc...)**, **kiến trúc(x86_64, arm64)**, **Quyền thực thi(dùng để tương tác với các service khác)**, một vài cấu hình thêm như **Enable function URL(cho phép chạy lambda bằng URL trên browser hay POSTMAN)**, **Enable tags(Dùng để quản lý tài nguyên của Lambda)**, **Enable VPC**
    
    - Lambda Life cycle
    
    **Vòng đời của một Lambda sẽ bao gồm **AWS và Lambda của chúng ta**: + **AWS**: **Boot runtime -> Load Dependencies(Mất một thời gian, nhưng không tốn tiền)**
    
    - Bản chất Lambda không cung cấp cơ sở hạ tầng, nhưng đằng sau mà là các Container để xử lý code của bạn
    - Đây được gọi là giai đoạn **Cold Start của lambda** và giai đoạn này sẽ **tốn một khoảng thời gian để khởi động** và sau một thời gian (tùy vào ngôn ngữ runtime) nếu không có Lambda call nào nữa thì toàn bộ giai đoạn trên sẽ bị xóa và đây điểm trừ cho các phần core quan trọng yêu cầu tốc độ cao
    - Khi mà ta đã gọi Lambda và xong quá trình Cold start (Warm up) thì các request sau ta sẽ **không cần giai đoạn** này nữa, nhưng nếu **2 request cùng một lúc** thì lúc này ta cần **2 Lambda function** để xử lý và ta **phải tạo thêm** một Lambda function nữa và tính chất của Lambda là **isolate** nên lambda mới sẽ phải trải qua giai đoạn Cold start
    - Solution cho vấn đề này là dùng **CloudWatch ping** qua để container này luôn chạy hoặc **cung cấp cho lambda số lượng concurrent** hay dùng CloudWatch Events(cronjob) -> giải pháp này sẽ tốn thêm tiền
    
    + **Lambda của chúng ta**: **Invoke, xử lý -> Shutdown** + **Init + Invoke** là hai cái AWS sẽ căn cứ vào **để tính phí** (Init hay Invoke càng lâu phí càng nhiều) **-Lambda SnapStart** Là một trong những chức năng của Lambda nhằm **tăng tốc độ và giảm chi phí** của Lambda lên bằng cách AWS sẽ khởi tạo và gọi Lambda trước đó 1 lần, xong sau đó sẽ lưu **Cold start** vào **memory hoặc disk** để khi sau này khi có gọi lại thì ta sẽ **bỏ qua bước Cold start** và tới thẳng bước **Invoke** + Hiện tại Lambda SnapStart chỉ support: **Java, Pi Tông và .Net**, nhưng có **một solution khác** cũng sẽ đảm bảo tận dụng được cơ chế này
    
    - Layer
    
    **Được dùng để chứa các thư viện bên thứ ba, khi Lambda function nào cần dùng thì chỉ cần vào **Layer** để lấy ra sử dụng việc này giúp: + Khiến cho các mã nguồn Lambda đơn giản, tinh lọc hơn + Khi cần thay đổi ta chỉ cần vào Layer và sửa chứ không phải đi từng Lambda để sửa + Giảm kích thước deployment package + Tối ưu hóa thời gian cold start + Tận dung Cache Layer trên container runtime: Nếu nhiều Lambda function dùng cùng một Layer, chúng có thể tận dụng Layer đã cache ở giai đoạn Cold start thay vì tải lại, giúp giảm thời gian khởi chạy. Nếu biết **tận dụng Cache** thì đây cũng chính là **một giải pháp giảm thiểu chi phí** khi khởi tạo Lambda, nhưng đây là **một con dao hai lưỡi** khi: + Layer quá lớn **>250MB unzipped** thì bạn sẽ bị **tính thêm chi phí** + Làm tăng thời gian tải ban đầu: Nếu Layer quá lớn **>100MB** thì Lambda **cần thời gian** tải Layer từ S3 **khi cold start** → có thể tăng chi phí gián tiếp do thời gian chạy lâu hơn.