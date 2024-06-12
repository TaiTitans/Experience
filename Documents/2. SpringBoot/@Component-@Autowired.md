
----

Một trong những nhiệm vụ chính của **Spring** là tạo ra một cái **_Container_** chứa các **_Dependency_** cho chúng ta.

`SpringApplication.run(App.class, args)` chính là câu lệnh để tạo ra **_container_**. Sau đó nó tìm toàn bộ các **_dependency_** trong project của bạn và đưa vào đó.

Spring đặt tên cho **_container_** là `ApplicationContext`

và đặt tên cho các **_dependency_** là `Bean`.


@Component

@Component là Annotation (chú thích) đánh dấu trên các Class để Spring biết đó là 1 `Bean`
**Spring Boot** khi chạy sẽ dò tìm toàn bộ các _Class cùng cấp_ hoặc ở trong các _package thấp hơn_ (Chúng ta có thể cấu hình việc tìm kiếm này, sẽ đề cập sau). Trong quá trình dò tìm này, khi gặp một class được đánh dấu `@Component` thì nó sẽ tạo ra một instance và đưa vào `ApplicationContext` để quản lý.


**@Autowired** là một annotation trong Spring Boot được sử dụng để tự động tiêm phụ thuộc vào các thành phần của ứng dụng. Nó giúp đơn giản hóa việc quản lý phụ thuộc và làm cho mã dễ đọc và bảo trì hơn.

**Cách thức hoạt động**

Khi bạn sử dụng @Autowired trên một trường hoặc phương thức, Spring Boot sẽ tự động tìm kiếm một bean phù hợp để tiêm vào. Bean này có thể được xác định theo tên bean, kiểu bean hoặc cả hai. Nếu có nhiều bean phù hợp, Spring Boot sẽ sử dụng các quy tắc ưu tiên để chọn bean phù hợp nhất.

**Ưu điểm**

- Giúp đơn giản hóa việc quản lý phụ thuộc
- Làm cho mã dễ đọc và bảo trì hơn
- Giảm thiểu lỗi liên quan đến việc tiêm phụ thuộc thủ công
### Singleton

Điều đặc biệt là các `Bean` được quản lý bên trong `ApplicationContext` đều là singleton.

Tất cả những `Bean` được quản lý trong `ApplicationContext` đều chỉ được tạo ra **một lần duy nhất** và khi có `Class` yêu cầu `@Autowired` thì nó sẽ lấy đối tượng có sẵn trong `ApplicationContext` để _inject_ vào.

Trong trường hợp bạn muốn mỗi lần sử dụng là một instance hoàn toàn mới. Thì hãy đánh dấu `@Component` đó bằng `@Scope("prototype")`


More:
Trong trường hợp Autowired không biết phải chọn đối tượng nào khi 1 class được implement thì có 2 cách giải quyết:
1. @Primary nó dánh đấu luôn được ưu tiên lựa chọn trong trường hợp nhiều Bean cùng loại trong Context.
2. @Qualifier xác định tên của 1 `Bean` mà bạn muốn chỉ định inject.
Ví dụ:
```Java
@Component("bikini")
public class Bikini implements Outfit {
    @Override
    public void wear() {
        System.out.println("Mặc bikini");
    }
}

@Component("naked")
public class Naked implements Outfit {
    @Override
    public void wear() {
        System.out.println("Đang không mặc gì");
    }
}

@Component
public class Girl {

    Outfit outfit;

    // Đánh dấu để Spring inject một đối tượng Outfit vào đây
    @Autowired
    public Girl(@Qualifier("naked") Outfit outfit) {
        // Spring sẽ inject outfit thông qua Constructor đầu tiên
        // Ngoài ra, nó sẽ tìm Bean có @Qualifier("naked") trong context để ịnject
        this.outfit = outfit;
    }

    // GET
    // SET
}
```


