
---

## MVC - Model View Controller:

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#mvc---model-view-controller)

Model: chịu trách nhiệm về business logic của ứng dụng. Nó quản lý trạng thái của ứng dụng. Điều này cũng bao gồm đọc và ghi dữ liệu, duy trì trạng thái ứng dụng và thậm chí nó có thể bao gồm các tác vụ liên quan đến quản lý dữ liệu như xác thực dữ liệu và kết nối mạng.

View: thành phần này có hai nhiệm vụ quan trọng: trình bày dữ liệu cho người dùng và xử lý tương tác của người dùng.

Controller: View class và Model class giao tiếp với nhau thông qua một hoặc nhiều Controller (quan hệ một nhiều).

### **Ưu điểm MVC**

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#%C6%B0u-%C4%91i%E1%BB%83m-mvc)

- Nhẹ, tiết kiệm tài nguyên.
- Kết cấu đơn giản.
- Dễ dàng kiểm tra và phát hiện lỗi.
- Dễ dàng phân tách các phần Model và View.

### **Nhược điểm MVC**

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#nh%C6%B0%E1%BB%A3c-%C4%91i%E1%BB%83m-mvc)

- Chỉ phù hợp với các dự án lớn. Đối với các dự án nhỏ mô hình sẽ trở nên cồng kềnh và tốn nhiều thời gian để trung chuyển dữ liệu.
- Controller sẽ trở nên phức tạp theo thời gian.
- Controller liên quan với View, vì thế View thay đổi Controller phải thay đổi theo.

## MVP - Model View Presenter

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#mvp---model-view-presenter)

MVP architecture pattern provides an easy way to structure the project codes. The reason why MVP is widely accepted is that it provides modularity, testability, and a more clean and maintainable codebase. It is composed of the following three components:

- **Model:** Model đại diện cho một tập hợp các lớp mô tả dữ liệu và business logic. Nó cũng xác định các  business rules cho dữ liệu có nghĩa là cách dữ liệu có thể được thay đổi và thao tác..
    
- **View:** The View is used for making interactions with users like XML, Activity, fragments and keep a track of the user’s action in order to notify the Presenter.
    
- **Presenter:** Presenter đóng vai trò làm trung gian giữa View và Model. Presenter nhận dữ liệu từ Model và hiển thị nó trên View, cũng như xử lý các sự kiện từ người dùng trên View và truyền nó đến Model để thực hiện các thao tác xử lý. Presenter cũng có trách nhiệm thực hiện các logic xử lý phức tạp hoặc logic nghiệp vụ, tránh việc để View xử lý các công việc này.
    
    Trong mô hình MVP, Presenter không trực tiếp điều khiển view. Thay vào đó, Presenter sử dụng Interface để tương tác với View và chỉ định các hành động cần thực hiện. Điều này cho phép View độc lập với Presenter và giúp giảm thiểu sự phụ thuộc giữa các thành phần trong ứng dụng. Nó cũng giúp cho việc kiểm thử và bảo trì ứng dụng dễ dàng hơn.
    

### **Ưu điểm MVP**

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#%C6%B0u-%C4%91i%E1%BB%83m-mvp)

- Cấu trúc rõ ràng và trực quan hơn MVC.
- Dễ dàng để viết unit test cho Presenter vì MVP hoạt động độc lập với View và không gắn với bất cứ API nào của Android.

### **Nhược điểm MVP**

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#nh%C6%B0%E1%BB%A3c-%C4%91i%E1%BB%83m-mvp)

- Mô hình sẽ to dần theo thời gian.
- Presenter sẽ lớn thêm khi thêm các logic nghiệp vụ. Người dùng sẽ khó kiểm soát và chia nhỏ đoạn code khi Presenter đã quá lớn.

[![](https://github.com/Meonix/SummaryDocument/raw/master/resources/mvp.png)](https://github.com/Meonix/SummaryDocument/blob/master/resources/mvp.png)

Dưới đây là một số sự khác biệt giữa MVC và MVP:

- Liên kết – Layer View và Model trong MVC liên kết chặt chẽ hơn trong MVP
- Giao tiếp – Giao tiếp giữa View-Presenter và Presenter-Model diễn ra thông qua một Interface, Presenter chịu trách nhiệm xử lý logic hiển thị cho tầng View. Còn Layer controller trong MVC chỉ là cầu nối cho View và Model, nó không chịu trách nhiệm xử lý logic cho tầng View.
- Xử lý hành vi người dùng – Trong MVC hành vi người dùng được xử lý bởi Controller nó hướng dẫn Model thực hiện các hành động cần thiết để cập nhật dữ liệu. Nhưng trong MVP, hành vi người dùng được xử lý bởi View và nó hướng dẫn cho Presenter thực thi các hàm cần thiết.
- Mối quan hệ **many-to-one** là mối quan hệ giữa View và Controller trong MVC, nhưng trong MVP thì Presenter và View có mối quan hệ **one-to-one**.
- Thành phần chính – Trong MVC, layer Controller là thành phần chính, nó quản lý View và tương tác với Model để phục vụ các yêu cầu của người dùng. Còn trong MVP View là thành phần chính, nó trực tiếp nắm bắt các hành vi người dùng và gọi đến các method trong Presenter tương ứng để thực hiện yêu cầu của người dùng.


## MVVM:

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#mvvm)

        This pattern based on MVC and MVP which attempts to more clearly separate the development of UI from that of the business logic and behavior in an application. However, the drawbacks of the MVP pattern have been solved by MVVM. It suggests separating the data presentation logic(Views or UI) from the core business logic part of the application. The separate code layers of MVVM are:

- **Model:** This layer is responsible for the abstraction of the data sources. Model and ViewModel work together to get and save the data.
- **View:** The purpose of this layer is to inform the ViewModel about the user’s action. This layer observes the ViewModel and does not contain any kind of application business logic.
- **ViewModel:** It exposes those data streams which are relevant to the View. Moreover, it serves as a link between the Model and the View.

[![](https://github.com/Meonix/SummaryDocument/raw/master/resources/mvvm.png)](https://github.com/Meonix/SummaryDocument/blob/master/resources/mvvm.png)

### **Ưu điểm MVVM**

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#%C6%B0u-%C4%91i%E1%BB%83m-mvvm)

- Người dùng có thể thực hiện unit testing mà không phụ thuộc vào View.
- Khi test không cần phải tạo mockup như MVP chỉ cần xác nhận biến observable thích hợp.
- Sử dụng XAML cho View, vì vậy có thể chỉnh sửa giao diện, không gây ảnh hưởng đến code.
- Phân rất rõ ràng ba phần, vì vậy sẽ rất dễ dàng lập trình và kiểm tra sửa lỗi code.

### **Nhược điểm MVVM**

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#nh%C6%B0%E1%BB%A3c-%C4%91i%E1%BB%83m-mvvm)

- Khi gán nhiều biến vào View các logic sẽ rải rác tăng dần theo thời gian gây khó khăn cho việc kiểm soát code.
- Data binding hai chiều gây hao tốn tài nguyên bộ nhớ.

## Sự khác nhau giữa MVVM và MVP

[](https://github.com/Meonix/SummaryDocument/blob/master/General/Architecture.md#s%E1%BB%B1-kh%C3%A1c-nhau-gi%E1%BB%AFa-mvvm-v%C3%A0-mvp)

| MVP(Model View Presenter)                                                                                                                | MVVM(Model View ViewModel)                                                                                                                      |     |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | --- |
| It resolves the problem of having a dependent **View** by using **Presenter** as a communication channel between **Model** and **View.** | This architecture pattern is more event-driven as it uses data binding and thus makes easy separation of core business logic from the **View.** |     |
| The one-to-one relationship exists between the **Presenter** and the **View**.                                                           | Multiple **View** can be mapped with single **ViewModel**.                                                                                      |     |
| In this, observables are not needed as a **Presenter** layer.                                                                            | In this, observables are needed due to the absence of **Presenter** layer.                                                                      |     |
| The architecture is heavyweight.                                                                                                         | The architecture is leightweight as compared to MVP.                                                                                            |     |
| Testing is done easily as business logic and UI are separate.                                                                            | Testing is difficult as Business logic and UI are not separate.                                                                                 |     |
| In this case, debugging is easy.                                                                                                         | In this case, debugging is not easy.                                                                                                            |     |
| The **Presenter** has knowledge about the **View.**                                                                                      | **ViewModel** has no reference to the **View**.                                                                                                 |     |
| **Model** layer returns the response of the user’s input to the **Presenter** which forwards it to **View**.                             | After performing operations according to the user’s input, the **Model** layer returns the response to the **View**.                            |     |
| **Presenter** handles the application flow and the **View** is the actual application.                                                   | **ViewModel** is the actual application and **View** is the interface for the user in order to interact with the app.                           |     |
| The project file will contain more classes as well as code.                                                                              | The Project file will contain more classes but less code per class.                                                                             |     |
| Maintenance is not ensured in the architecture.                                                                                          | Maintenance of the layers is very good.                                                                                                         |     |
| It does not work well with android applications or software.                                                                             | It works extremely well with android applications or software.                                                                                  | ư   |
|                                                                                                                                          |                                                                                                                                                 |     |


## Clean Architecture

![[Pasted image 20240515225814.png]]

### The Dependency Rule

Quy tắc phụ thuộc là quy tắc giúp kiến trúc này hoạt động. Quy tắc này cho rằng sự phụ thuộc trong mã nguồn chỉ có thể đi vào hướng bên trong. Các thành phần trong vòng tròn bên trong không được biết bất cứ điều gì về những thứ thuộc về vòng tròn bên ngoài. Điều này bao gồm cả các hàm, lớp, biến hoặc bất kỳ thực thể phần mềm nào có tên.

- Mô hình sẽ chia thành nhiều vòng tròn, mỗi vòng tròn tương ứng với 1 layer.
- Vị trí của vòng tròn nói lên cấp (level) của chúng, vòng tròn càng xa tâm thì có cấp càng cao.
- Dependency rule là quy tắc quy định rằng:
    - Layer trong (level thấp hơn) không được chứa source code của layer ngoài (có level cao) hơn.
    - Bất cứ thay đổi ở layer ngoài không làm ảnh hưởng tới layer trong, có nghĩa là sự phụ thuộc hướng vào trong.
- Clean architecture thông thường chia thành 4 layer: entities, use cases, interface adapters và Frameworks and drivers. Nhưng tuỳ theo hệ thống, chúng ta có thể chia thành nhiều layer hơn nhưng phải tuyệt đối tuân thủ vào dependency rule.
- Khi chúng ta đi vào trong từ layer level cao xuống các layer thấp hơn, mức độ trừu tượng hoá sẽ được tăng lên (mức độ phụ thuộc giảm đi).

### Entities

---
- Entities chứa business logic của cả hệ thống (tập các app) chứ không bị bó buộc trong 1 app nào cả. Nói một cách dễ hiểu Entities sẽ đảm nhận nhiệm vụ định nghĩa các POJO chứa các thông tin liên quan đến business của ứng dụng.
- Do entities có phạm vị hệ thống nên sự thay đổi hoạt động (không phải business logic) của một ứng dụng không làm ảnh hưởng tới entities. Do đó, entities có thể chạy độc lập mà không sợ bị lỗi.

Ví dụ: Chúng ta xây dựng app tin tức. Ta tạo ra 1 class News, thì News chính là một entity của hệ thống.


---
### Use Cases
---
- Use cases chứa application specific business rules, thực hiện các use case của 1 ứng dụng dựa trên các entity object của entities.
- Use cases không chứa source code của interface adapters và frameworks and drivers.
- Use cases cũng không bị ảnh hưởng bởi sự thay đổi interface adapters và frameworks and drivers.
- Use cases thể hiện business logic của ứng dụng. Do đó khi requirement của ứng dụng thay đổi, code trong layer này cũng sẽ bị thay đổi.

Ví dụ: Use Cases thêm, sửa, xóa bài viết.


---
### Interface adapters
---

- Interface adapters là layer dùng để chuyển đổi dữ liệu của ở entites và use cases thành định dang thuận tiện nhất để làm việc với các tác vụ như database, web service… Do đó, interface adapters cũng có thể chứa các framework có sẵn để làm việc.
- Use cases, entities và các framework được sử dụng để chuyển đổi dữ liệu.

Ví dụ: Ứng dụng sử dụng Room Database

- Entities và use cases không thể chứa source code về Room Database.
- Interface adapter tạo ra các đối tượng như table, SQL query… để làm việc thuận tiện nhất với yêu cầu của ứng dụng.
- Ngoài interface adapters, không thành phần nào của clean architecture làm việc với Room Database.


---
### Framework and drivers
---

- Đây là layer ngoài cùng, nó là nơi chúng ta đi vào chi tiết hoá yêu cầu của ứng dụng như UI, application behavior, interaction… và cũng có thể bao gồm các Framework được xây dựng để phục vụ việc chi tiết hoá như UI framework…
- Chúng ta chỉ viết code liên quan đến việc chi tiết hoá như UI design…

Ví dụ 4: Với WEB, trong frameworks and drivers:

- Chúng ta sẽ import những UI dependency như TailwindCSS… và sử dụng code trong interface adapter để tạo ra UI.
- Ngoài ra, chúng ta cũng sẽ thực hiện các tác vụ như điều hướng giữa các router, thực hiện tương tác với user, chỉnh setting cho web…

---

## Ưu điểm của Clean Architecture
---
- Giúp business logic trở nên rõ ràng
- Việc tách biệt thành các layer và tuân theo dependecy rule, hệ thống của chúng ta sẽ dễ dàng test, maintain và thay đổi.
- Khi một thành phần giao diện, database bị lỗi thời, chúng ta có thể dễ dàng thay thế nó với effort bỏ ra là ít nhất.

---
## Nhược điểm
---
- Phân tách cấu trúc ra nhiều tầng nên dẫn đến việc số lượng code sinh ra là rất lớn, không phù hợp cho các dự án nhỏ.
- Khó hiểu, khó áp dụng với những người mới.

**!! Lưu ý :**

    Trong các mô hình MVC, MVVM, MVP, Repository không được liệt kê là một layer bởi vì nó không là một phần của kiến trúc mô hình, mà là một design pattern. Repository pattern được sử dụng để giải quyết vấn đề truy cập và quản lý dữ liệu trong ứng dụng, và thường được sử dụng cùng với các mô hình kiến trúc khác như MVC, MVVM, MVP. Repository pattern đóng vai trò quan trọng trong việc quản lý truy cập và xử lý dữ liệu, giúp cho việc tách biệt các thành phần trong ứng dụng trở nên dễ dàng hơn.