

-----
### Tổng quan


Như mặc định, Spring sẽ tạo tất cả các `singleton Bean` trong quá trình startup `Application Context`. Tuy nhiên, có những trường hợp chúng ta chưa dùng đến `Bean` khi mới startup `Application Context`, mà khi nào yêu cầu thì nó sẽ được tạo ra. Để làm được điều đó, Spring cho ra đời annotation `@Lazy` .

_ApplicationConfig.class_
```Java
@Lazy
@Configuration
public class ApplicationConfig {

    @Bean
    public FirstBean firstBean() {
        return new FirstBean();
    }

    @Bean
    public SecondBean secondBean() {
        return new SecondBean();
    }
}
```

_ExampleApplication.class_

```Java
@SpringBootApplication
public class ExampleApplication {
    public static void main(String[] args) {
        AnnotationConfigApplicationContext context = new AnnotationConfigApplicationContext(ExampleApplication.class);
        System.out.println("ApplicationContext đã khởi tạo!");
        FirstBean firstBean = context.getBean(FirstBean.class);
        SecondBean secondBean = context.getBean(SecondBean.class);
        context.close();
    }
}
```


Khi chúng ta đặt annotation `@Lazy` trong class `@Configuration` , nó chỉ ra rằng tất các phương thức với annotation `@Bean` sẽ tạm thời chưa được khởi tạo khi `ApplicationContext` startup. Sau khi start xong, thì chúng tôi yêu cầu getBean() thì chúng mới bắt đầu được khởi tạo, điều này được chứng minh ở output khi run chương trình:

```
/* OUPUT:
ApplicationContext đã khởi tạo!
Bean FirstBean đã được khởi tạo!
Bean SecondBean đã được khởi tạo!

*/
```


### @Lazy và @Bean

Để áp dụng `Bean` được chỉ định riêng, đơn giản chỉ cần đặt annotation `@Lazy` vào `Bean` được yêu cầu. Khi đó chúng tôi thay đổi code class `_ApplicationConfig` như sau:
__ApplicationConfig.class_

```Java
@Configuration
public class ApplicationConfig {

    @Lazy
    @Bean
    public FirstBean firstBean() {
        return new FirstBean();
    }

    @Bean
    public SecondBean secondBean() {
        return new SecondBean();
    }
}

```

Khi đó `Bean` SecondBean sẽ được khởi tạo đồng thời khi hệ thống running, còn `Bean` FirstBean sẽ được khởi tạo sau khi được yêu cầu.

### @Lazy và @Component

Thay vì tạo `Bean` trong class `Config`, chúng tôi cũng có thể khởi tạo Class sử dụng annotation `@Component`. Khi đó trường hợp này tương tự với trường hợp kết hợp `@Lazy` và `@Bean`. _FirstBean.class_

### @Lazy và @Autowired
Chú ý rằng trong trường hợp này `@Lazy` được đặt cả ở `@Component` và `@Autowired`.

```Java
public class ServiceBean {

    @Lazy
    @Autowired
    private FirstBean firstBean;

    public FirstBean getFirstBean() {
        return firstBean;
    }
}
```

Khi đó `Bean` FirstBean sẽ được khởi tạo khi chúng tôi gọi hàm `getFirstBean()`.