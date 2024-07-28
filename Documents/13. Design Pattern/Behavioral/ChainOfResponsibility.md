
---

## Khái Niệm
---

Trong phát triển phần mềm, việc quản lý không hiệu quả các yêu cầu đa dạng và phức tạp có thể dẫn đến mã lập trình rối rắm, khó bảo trì. Nếu một đối tượng hoặc lớp đảm nhiệm quá nhiều nhiệm vụ, điều này tạo ra sự phụ thuộc lẫn nhau cao và làm giảm khả năng mở rộng của hệ thống, cũng như làm tăng độ phức tạp trong quản lý. Chain of Responsibility giải quyết vấn đề này bằng cách phân tán trách nhiệm xử lý yêu cầu qua một chuỗi các đối tượng, giúp giảm sự phụ thuộc và tăng tính linh hoạt.


### Tổng quan


- **Định Nghĩa của Pattern:** Chain of Responsibility Pattern cho phép một yêu cầu được chuyển qua một chuỗi các bộ xử lý. Mỗi bộ xử lý quyết định xử lý yêu cầu hoặc chuyển nó đến bộ xử lý tiếp theo trong chuỗi.
    
- **Mục Đích:** Mẫu thiết kế này giúp loại bỏ sự cứng nhắc trong việc chỉ định chính xác đối tượng xử lý một yêu cầu cụ thể. Nó giúp phân tán trách nhiệm xử lý và giảm sự phụ thuộc lẫn nhau giữa các đối tượng.
    
- **Ý Tưởng Cốt Lõi:** Trong Chain of Responsibility, không có đối tượng cụ thể nào được chỉ định trước để xử lý một yêu cầu. Thay vào đó, mỗi đối tượng trong chuỗi có thể xử lý yêu cầu hoặc chuyển nó đến đối tượng tiếp theo. Điều này tạo ra một hệ thống linh hoạt, nơi xử lý yêu cầu không phụ thuộc vào một đối tượng cố định.

## Đặt vấn đề

---
Bây giờ, hãy tưởng tượng bạn là một lập trình viên đang phát triển một hệ thống đặt hàng trực tuyến. Mục tiêu của bạn là hạn chế quyền truy cập vào hệ thống, chỉ cho phép những người đã xác thực mới có thể tạo đơn hàng. Đối với admin, họ có quyền truy cập toàn diện đến mọi đơn hàng.

Sau một thời gian, bạn nhận ra rằng các thao tác xác thực cần được thực hiện theo một trình tự nhất định. Hệ thống sẽ xác thực thông tin người dùng khi họ đăng nhập, nhưng nếu quá trình này thất bại, không cần thiết phải tiến hành các bước tiếp theo.

Một vài tháng sau, bạn cần phải thêm vào một số bước kiểm tra xác thực mới:

- Một đồng nghiệp gợi ý: "Em ơi, việc truyền dữ liệu trực tiếp vào cơ sở dữ liệu có thể rất nguy hiểm." Dựa trên lời khuyên này, bạn thêm một bước kiểm tra và lọc dữ liệu.
- Sau đó, một hacker mũ trắng chỉ ra rằng hệ thống của bạn dễ dàng bị tấn công bằng brute force. Nhận ra điều này, bạn nhanh chóng thêm một lớp kiểm tra để chặn các yêu cầu đăng nhập liên tiếp không thành công từ cùng một IP.

## Giải pháp

Chain of Responsibility dựa vào việc chuyển đổi các hành vi cụ thể thành các đối tượng hoạt động lập gọi là handlers. Trong vấn đề trên, với hoạt động kiểm thử bạn nên đổi chúng thành một lớp đối tượng cụ thể với một phương thức duy nhất là kiểm tra.


## Cách triển khai


### Bước 1: Tạo Interface Handler

```Java
interface Handler {
    void handleRequest(String request);
    void setNext(Handler nextHandler);
}
```

### Bước 2: Tạo Concrete Handlers


Mỗi `ConcreteHandler` sẽ triển khai `Handler` và quyết định liệu nó có thể xử lý yêu cầu hay chuyển nó đến handler tiếp theo.


```Java
class ConcreteHandler1 implements Handler {
    private Handler next;

    @Override
    public void handleRequest(String request) {
        if (request.equals("Handler1")) {
            System.out.println("ConcreteHandler1 đã xử lý yêu cầu.");
        } else if (next != null) {
            next.handleRequest(request);
        }
    }

    @Override
    public void setNext(Handler nextHandler) {
        this.next = nextHandler;
    }
}

class ConcreteHandler2 implements Handler {
    private Handler next;

    @Override
    public void handleRequest(String request) {
        if (request.equals("Handler2")) {
            System.out.println("ConcreteHandler2 đã xử lý yêu cầu.");
        } else if (next != null) {
            next.handleRequest(request);
        }
    }

    @Override
    public void setNext(Handler nextHandler) {
        this.next = nextHandler;
    }
}
```

### Bước 3: Tạo Client Class

Client sẽ tạo yêu cầu và gửi chúng qua chuỗi các handler.

```Java
class Client {
    private Handler handler;

    public Client(Handler handler) {
        this.handler = handler;
    }

    public void makeRequest(String request) {
        handler.handleRequest(request);
    }
}
```

### Bước 4: Thiết lập và Sử dụng Chain of Responsibility

Ở đây, chúng ta tạo các handler, thiết lập chuỗi trách nhiệm và sau đó là thực thi yêu cầu thông qua client.

```Java
public class Main {
    public static void main(String[] args) {
        Handler handler1 = new ConcreteHandler1();
        Handler handler2 = new ConcreteHandler2();

        handler1.setNext(handler2);

        Client client = new Client(handler1);
        client.makeRequest("Handler1");
        client.makeRequest("Handler2");
        client.makeRequest("Unknown"); // Sẽ không được xử lý bởi bất kỳ handler nào
    }
}
```

## Khi nào áp dụng Chain-Of-Responsibility Pattern


Mẫu thiết kế Chain-Of-Responsibility nên được áp dụng trong các tình huống sau:

1. **Xử lý Nhiều Loại Yêu Cầu Khác Nhau**: Khi chương trình của bạn cần xử lý nhiều loại yêu cầu khác nhau và bạn không thể hoặc không muốn xác định trước loại yêu cầu cụ thể cũng như thứ tự xử lý của chúng, mẫu này sẽ rất hữu ích. Nó cho phép bạn tổ chức một chuỗi các đối tượng xử lý, mỗi đối tượng sẽ xử lý một loại yêu cầu cụ thể hoặc chuyển nó đến đối tượng tiếp theo trong chuỗi.
    
2. **Xử lý Tuần Tự**: Mẫu này cũng thích hợp khi một tác vụ cần được xử lý một cách tuần tự. Trong trường hợp này, mỗi đối tượng trong chuỗi sẽ thực hiện một phần của tác vụ hoặc quyết định xem liệu có nên chuyển tác vụ đó đến đối tượng tiếp theo trong chuỗi hay không.
    
3. **Phân Cấp Trách Nhiệm**: Khi bạn muốn phân cấp trách nhiệm xử lý yêu cầu, mẫu này cũng rất hữu ích. Nó cho phép từng đối tượng trong chuỗi tập trung vào một phần nhỏ của tác vụ, làm cho việc xử lý trở nên quản lý và bảo trì dễ dàng hơn.
    
4. **Linh Hoạt trong Xử lý Yêu Cầu**: Cuối cùng, mẫu Chain-Of-Responsibility tạo điều kiện cho sự linh hoạt trong việc xử lý yêu cầu. Bạn có thể dễ dàng thay đổi hoặc mở rộng chuỗi xử lý mà không cần thay đổi mã nguồn của các đối tượng xử lý hiện có.
    

Kết hợp những điểm trên, Chain-Of-Responsibility là một lựa chọn tuyệt vời cho các ứng dụng cần một cách tiếp cận linh hoạt và mở rộng trong việc xử lý một loạt yêu cầu khác nhau, đồng thời giữ cho mã nguồn trở nên gọn gàng và dễ quản lý.