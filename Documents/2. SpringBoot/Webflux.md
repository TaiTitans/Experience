
---

Spring Webflux Framework là một phần của Spring 5 và cung cấp [**Reactive Programming**][reactive-programming] nhằm hỗ trợ cho việc xây dựng ứng dụng web.

### Reactive Streams API

**Reactive Stream API** được tạo bởi các kỹ sư từ Netflix, Pivotal, Lightbend, RedHat, Twitter, and Oracle và bây giờ là một phần của Java 9. Nó định nghĩa 4 interface:

`Publisher:` Phát ra một chuỗi các sự kiện đến `subscriber` theo yêu cầu của người mà `subscriber` đến nó. Một `Publisher` có thể phục vụ nhiều `subscriber`. Interface này chỉ có một phương thức:

``` Java
public interface Publisher<T>
{
    public void subscribe(Subscriber<? super T> s);
}
```

`Subscriber:` Nhận và xử lý sự kiện được phát ra bởi `Publisher`. Chú ý rằng không có gì xảy ra cho tới khi Subscription – nó được gọi là báo hiệu yêu cầu cho `Publisher`.

```Java
public interface Subscriber<T>
{
    public void onSubscribe(Subscription s);
    public void onNext(T t);
    public void onError(Throwable t);
    public void onComplete();
}
```

`Subscription:` Định nghĩa mỗi quan hệ 1-1 giữa `Publisher` và `Subscriber`. Nó chỉ có thể được sử dụng bởi một `Subsriber` duy nhất và được sử dụng để báo hiệu yêu cầu (request) hoặc hủy (cancel) data.

```Java
public interface Subscription<T>
{
    public void request(long n);
    public void cancel();
}
```

`Processor:` Đại diện cho giai đoạn xử lý gồm cả `Publisher` và `Subscriber` đồng thời tuân thủ nguyên tắc của cả 2.


```Java
public interface Processor<T, R> extends Subscriber<T>, Publisher<R>
{
}
```

Bản chất, một `Subscriber` tạo một  `Subscription` tới `Publisher`, sau đó `Publisher` gửi một sự kiện cho `Subsriber` với một luồng các phần tử.


----
### Spring WebFlux

**Spring Webflux** là một phiên bản song song với **Spring MVC** và hỗ trợ non-blocking reactive streams. Nó hỗ trợ khái niệm **back pressure** và sử dụng Server Netty để run hệ thống reactive.

Spring webflux sử dụng project reactor ( Thư viện implements phỗ biến nhất) như một thư viện reactive, vì thế nó cung cấp 2 kiểu `Publisher`:

`Mono`: Phát ra 0 hoặc 1 phần tử.

`Flux` : Phát ra 0..N phần tử.

### Kết luận

Cả `Spring MVC` và Spring `Webflux` để hỗ trợ kiến trúc Client-Server nhưng điểm khác nhau chính là mô hình `concurrency` và hành động mặc định trong tính chất non-blocking và threads.

Trong Spring MVC, nó mặc định rằng ứng dụng có thể bị block tại thread hiện tại, trong khi webflux thì mặc định threads là non-blocking .

Reactive và non-blocking nhìn chung thì không làm cho ứng dụng chạy nhanh hơn. Lợi ích mà nó được kỳ vọng là mở rộng ứng dụng với số luồng nhỏ và yêu cầu ít bộ nhớ hơn. Nó làm cho các ứng dụng trở nên linh hoạt hơn khi tải.