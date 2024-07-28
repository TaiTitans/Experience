
----
# CompletableFuture là gì?

CompletableFuture được sử dụng cho lập trình bất đồng bộ trong Java. Lập trình bất đồng bộ là cách viết code không chặn bằng cách chạy một tác vụ trên một Thread riêng biệt khác với luồng chính và thông báo cho luồng chính về tiến trình, hoàn thành hoặc lỗi của nó.

Bằng cách này, chương trình chính không bị chặn / chờ đợi để hoàn thành nhiệm vụ và nó có thể thực thi các tác vụ khác song song. Các xử lý song song này cải thiện đáng kể hiệu suất của các chương trình của bạn.

---
# So sánh Future vs CompletableFuture

Một Future được sử dụng như là một tham chiếu đến kết quả của một tác vụ bất đồng bộ. Nó cung cấp một phương thức isDone() để kiểm tra xem việc tính toán có được thực hiện hay không, và phương thức get() để lấy kết quả của phép tính khi nó được thực hiện.

Như đã giới thiệu trong bài viết Lập trình đa luồng với Callable và Future trong Java , Future cung cấp các API rất tốt cho việc lập trình bất đồng bộ. Tuy nhiên nó có một vài hạn chế sau:

- Không thể quản lý kết quả trả về của Future. Ví dụ: khi main thread hoàn thành, chúng ta cũng mong muốn hoàn thành Future bằng gán cho nó một giá trị mặc định. Chúng ta không thể làm được điều này với Future.
    
- Không thể thực hiện thêm hành động nào đối với kết quả của Future mà không bị chặn:
    

- Future không thông báo cho về việc nó hoàn thành. Nó cung cấp phương thức get() để nhận kết quả. Tuy nhiên, phương thức này chặn main thread xử lý tác vụ khác cho đến khi kết quả có sẵn.
    
- Không thể thêm callback function khi Future hoàn thành.
    

- Nhiều Future không thể xử lý dây chuyền (chain). Ví dụ chương trình của chúng ta có nhiều tác vụ, khi tác vụ này hoành thành thì tác vụ khác xử lý. Với Future chúng ta không thể thực hiện: Future1 -> Future2 -> Future3 -> …
    
- Không thể kết hợp nhiều Future với nhau: Giả sử chúng ta có 10 Future khác nhau, chúng ta mong muốn khi tất cả chúng hoàn thành thì sẽ thực thi tiếp 1 Future khác. Chúng ta không thể làm điều này với Future.
    
- Không hỗ trợ xử lý ngoại lệ (Exception). Với Java 8, cung cấp một API mới giải quyết tất cả các giới hạn trên của Future đó chính là CompletableFuture.
    

CompletableFuture Implement các Interface Future và CompletionStage. Nó cung cấp nhiều các phương thức tiện lợi để: create, chain và combine nhiều Future, xử lý ngoại lệ, … Trong phần tiếp theo chúng ta sẽ lần lượt tìm hiểu các phương thức này qua các ví dụ.

`public class CompletableFuture<T> implements Future<T>, CompletionStage<T> {}`

---
# CompletionStage là gì?

Một CompletionStage là một lời hứa (Promise). Nó hứa hẹn rằng việc tính toán cuối cùng sẽ được thực hiện.

CompletionStage cung cấp nhiều phương thức khác nhau cho phép khai báo các callback sẽ được thực hiện khi hoàn thành. Bằng cách này, chúng ta có thể xây dựng hệ thống theo cách không bị chặn (block).
## Chạy bất đồng bộ với runAsync() và không cần kết quả trả về

Cú pháp:

```none
public static CompletableFuture<Void> runAsync(Runnable runnable);
```

## Chạy bất đồng bộ với supplyAsync() và cần nhận kết quả trả về

Cú pháp:

```none
public static <U> CompletableFuture<U> supplyAsync(Supplier<U> supplier);
```


# Chuyển đổi và thao tác trên một CompletableFuture

Phương thức CompletableFuture.get() là chặn (blocking), nó đợi cho đến khi Future được trả về kết quả sau khi hoàn thành.

CompletableFuture cung cấp phương thức cho phép đính kèm một phương thức khác, phương thức này được gọi lại (callback) khi Future hoàn thành. Bằng cách đó, chúng ta sẽ không cần phải đợi kết quả, và chúng ta có thể viết logic cần được thực thi sau khi hoàn thành tương lai bên trong hàm gọi lại của chúng ta.

Một số phương thức được sử dụng để callback là: thenApply(), thenAccept() và thenRun().

---
Một số tính năng chính của CompletableFuture:

1. **Tạo và hoàn thành CompletableFuture**:
    
    - Tạo một CompletableFuture bằng cách sử dụng các phương thức tĩnh như `CompletableFuture.supplyAsync()` hoặc `CompletableFuture.runAsync()`.
    - Hoàn thành CompletableFuture bằng cách sử dụng các phương thức như `complete()` hoặc `completeExceptionally()`.
2. **Xử lý kết quả bất đồng bộ**:
    
    - Sử dụng `thenApply()` để áp dụng một hàm xử lý kết quả của CompletableFuture.
    - Sử dụng `thenAccept()` để tiêu thụ kết quả mà không trả về kết quả mới.
    - Sử dụng `thenRun()` để thực hiện một tác vụ sau khi CompletableFuture hoàn thành.
3. **Xử lý nhiều CompletableFutures**:
    
    - Sử dụng `allOf()` để chờ đợi nhiều CompletableFutures hoàn thành.
    - Sử dụng `anyOf()` để chờ đợi một trong nhiều CompletableFutures hoàn thành.
4. **Xử lý lỗi**:
    
    - Sử dụng `exceptionally()` để xử lý ngoại lệ trong CompletableFuture.
    - Sử dụng `handle()` để xử lý cả kết quả và ngoại lệ.
5. **Kết hợp các CompletableFutures**:
    
    - Sử dụng `thenCompose()` để kết hợp các CompletableFutures theo thứ tự.
    - Sử dụng `thenCombine()` để kết hợp hai CompletableFutures song song.

