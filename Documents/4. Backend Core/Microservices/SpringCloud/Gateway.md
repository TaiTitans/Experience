
---
### Spring Cloud Gateway là gì?

Spring Cloud Gateway là một thư viện được sử dụng để xây dựng cổng API trên nền tảng Spring WebFlux. Dự án này nhằm mục đích cung cấp một cách đơn giản và hiệu quả để định tuyến đến các API và cung cấp các vấn đề liên quan đến bảo mật, giám sát/thống kê và tính chịu lỗi.

Nó là một proxy, cổng thông tin, một lớp trung gian giữa người dùng và dịch vụ của bạn.

Eureka server giải quyết vấn đề đặt tên cho các dịch vụ thay vì hardcode địa chỉ IP của chúng.

Tuy nhiên, chúng ta vẫn có thể có nhiều hơn một dịch vụ (instances) đang chạy trên các cổng khác nhau

1. **Spring Cloud Gateway** sẽ ánh xạ giữa một tiền tố đường dẫn, ví dụ như **/service02/** và một dịch vụ **Service02**. Nó sử dụng **Eureka server** để định tuyến dịch vụ được yêu cầu.
2. Nó cân bằng tải (sử dụng Ribbon) giữa các phiên bản của dịch vụ đang chạy trên các cổng khác nhau.
3. Chúng ta có thể **lọc các yêu cầu( filter requests)**, thêm **xác thực(authentication)** và nhiều hơn nữa.

**Tính năng của Spring Cloud Gateway:**

1. Xây dựng trên Spring Framework 5, Project Reactor và Spring Boot 2.0
2. Có thể khớp các tuyến đường với bất kỳ thuộc tính yêu cầu nào.
3. Các Predicates và Filters được chỉ định cho từng tuyến đường cụ thể.
4. Tích hợp Circuit Breaker.
5. Tích hợp DiscoveryClient của Spring Cloud.
6. Dễ dàng viết Predicates và Filters.
7. Giới hạn tốc độ yêu cầu (Request Rate Limiting).
8. Chuyển đổi lại đường dẫn (Path Rewriting).