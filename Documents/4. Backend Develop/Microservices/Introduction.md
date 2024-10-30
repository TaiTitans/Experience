
---
_Một microservice là một phương pháp kỹ thuật tập trung vào việc **phân tách (decomposing)** ứng dụng thành các module **chức năng đơn (single-function )** với các **giao diện được xác định rõ ràng (well-defined interfaces)**, được triển khai và vận hành **độc lập (Independent)** bởi **các nhóm nhỏ (Small Teams)** có quản lý **toàn bộ vòng đời (Entire Lifecycle)** của dịch vụ._

_Microservices giúp gia tăng tốc độ phát triển bằng cách **giảm thiểu sự giao tiếp (Minimizing Communication)** và phối hợp giữa các thành viên trong nhóm, đồng thời **giảm phạm vi và rủi ro của các thay đổi (The scope and risk of change)** ._


---
### Phân tách (Decomposing) là gì?

**Decomposing** trong ngữ cảnh của microservices nghĩa là phân tách ứng dụng thành các module chức năng đơn. Tức là, thay vì có một ứng dụng lớn thì chúng ta sẽ tách thành nhiều ứng dụng nhỏ hơn, mỗi ứng dụng nhỏ này có chức năng rõ ràng và phục vụ cho một nhu cầu cụ thể của ứng dụng lớn. Việc phân tách này giúp cho việc phát triển và bảo trì ứng dụng trở nên dễ dàng hơn và tăng tính linh hoạt của hệ thống.


---
### Single-function (Chức năng đơn) là gì?

Single-function là thuật ngữ trong kiến trúc Microservices, chỉ định rằng mỗi microservice chỉ nên chứa một chức năng cố định, thực hiện một nhiệm vụ cụ thể của hệ thống. Điều này giúp tách biệt các chức năng khác nhau của ứng dụng và giúp cho việc phát triển và bảo trì các microservice trở nên đơn giản hơn.

---
### Giao diện được xác định rõ ràng (well-defined interfaces) là gì?

Well-defined interfaces là các giao diện được định nghĩa rõ ràng và cung cấp các quy tắc về cách các thành phần khác của hệ thống có thể tương tác với chúng. Các giao diện này giúp đảm bảo tính tương thích và khả năng mở rộng trong hệ thống phân tán, bởi vì các thành phần khác nhau có thể giao tiếp với nhau thông qua các giao diện được định nghĩa trước đó mà không cần biết chi tiết bên trong của nhau.

---
### Độc lập (Independent) là gì?

Trong ngữ cảnh của microservices, độc lập (independent) thường được hiểu là các service có thể hoạt động hoàn toàn độc lập với nhau, mà không bị phụ thuộc vào các service khác. Nói cách khác, các service có thể triển khai và vận hành độc lập, và không cần phải liên tục phối hợp với nhau để thực hiện một chức năng nào đó. Điều này giúp giảm thiểu sự phụ thuộc giữa các service và làm cho hệ thống dễ dàng mở rộng và bảo trì hơn.

---
### Các nhóm nhỏ (Small Teams) là gì?

- Chúng ta chia công việc và phân nhóm cho các dịch vụ. Mỗi nhóm tập trung vào một dịch vụ cụ thể, họ không cần biết về các hoạt động nội bộ của các nhóm khác.
- Những nhóm này có thể làm việc hiệu quả, giao tiếp dễ dàng và mỗi dịch vụ có thể triển khai nhanh chóng ngay khi sẵn sàng.

---
### Toàn bộ vòng đời (Entire Lifecycle) là gì?

- Đội ngũ chịu trách nhiệm cho toàn bộ vòng đời của dịch vụ; từ việc viết mã, kiểm thử, phát triển, triển khai, gỡ lỗi và bảo trì.
- Trong một ứng dụng truyền thống, chúng ta có thể có một nhóm chịu trách nhiệm cho việc viết mã, và một nhóm khác chịu trách nhiệm cho việc triển khai. Tuy nhiên, trong kiến trúc microservices, điều này không đúng.

---
### Giảm thiểu sự giao tiếp (Minimizing Communication) là gì?

- Giảm thiểu sự giao tiếp không có nghĩa là các thành viên trong nhóm không cần quan tâm đến nhau. Điều này có nghĩa là chỉ giao tiếp cần thiết giữa các nhóm nên thông qua giao diện mà mỗi dịch vụ cung cấp.
- Các nhóm cần đồng ý về giao diện bên ngoài để việc giao tiếp giữa các dịch vụ được xác định rõ ràng.

---
### Giảm phạm vi và rủi ro của các thay đổi (The scope and risk of change) là gì?

- Các dịch vụ nên được thay đổi mà không làm hỏng các dịch vụ khác. Và miễn là chúng ta không thay đổi giao diện bên ngoài thì sẽ không có vấn đề cho các dịch vụ khác.
- Kết quả của các thay đổi là phiên bản của các dịch vụ được cập nhật một cách độc lập và không có mối quan hệ giữa chúng.