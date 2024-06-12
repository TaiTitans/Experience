
---
Khi xây dựng một chương trình với **Spring Boot** đôi lúc chúng ta một **Bean** chỉ được load lên hoặc khởi tạo theo một điều kiện nào đó. Ví dụ như tạo một **Bean** trong môi trường Test, còn môi trường thật sẽ không cần nữa.

**Spring Boot** hỗ trợ chúng ta làm điều này với Annotation `@Conditional`.


### Cách tạo bean có điều kiện

Trong **Spring Boot**, có rất nhiều cách để tạo ra `Bean`.

`@Component`, `@Configuration`, `@Bean`, `@Service`, v.v...

Với tất cả các cách tạo ra `Bean`, bạn đều có thể thêm **một hoặc nhiều điều kiện** để cho việc tạo ra `Bean` đó chỉ xảy ra nếu nó thỏa mãn một điều kiện cho trước.

**SpringBoot** hỗ trợ rất nhiều loại điều kiện khác nhau, tất cả đều sẽ bắt đầu bằng tiền tố `@Conditional...`


Cách thức hoạt động của tất cả `@Condition` như sau:

```Java
@Configuration
public class ConditionalOnBeanExample {
    /*
    ABeanWithCondition chỉ được tạo ra khi @Condition thỏa mãn
    */
    @Bean
    @Conditional...
    ABeanWithCondition aBeanWithACondition() {
        return new ABeanWithCondition();
    }
}

```

Nếu bạn gắn nó trên một `@Configuration` thì toàn bộ các `@Bean` bên trong sẽ bị chịu tác động.

---

### @ConditionalOnBean

`@ConditionalOnBean` sử dụng khi chúng ta muốn tạo ra một Bean, chỉ khi có một Bean khác đang tồn tại

### @ConditionalOnProperty


Dùng `@ConditionalOnProperty` khi bạn muốn quyết định sự tồn tại Bean thông qua cấu hình property.

### @ConditionalOnExpression

Trong trường hợp bạn muốn thỏa mãn nhiều điều kiện trong property, hãy sử dụng `@ConditionalOnExpression`


### @ConditionalOnMissingBean

Nếu trong `Context` chưa tồn tại bất kỳ Bean nào tương tự, thì `@ConditionalOnMissingBean` sẽ thỏa mãn điều kiện và tạo ra một Bean như thế.

### @ConditionalOnResource

`@ConditionalOnResource` thỏa mãn khi có một resources nào đó do bạn chỉ định tồn tại.

### @ConditionalOnClass

`@ConditionalOnClass` thỏa mãn khi trong classpath có tồn tại class mà bạn yêu cầu.


### @ConditionalOnMissingClass

`@ConditionalOnMissingClass` ngược lại với `@ConditionalOnClass`. thỏa mãn khi trong classpath **không** tồn tại class mà bạn yêu cầu.


### @ConditionalOnJava

[](https://github.com/loda-kun/spring-boot-learning/tree/master/spring-boot-%40Conditional#conditionalonjava)

`@ConditionalOnJava` thỏa mãn nếu môi trường chạy version Java đúng với điều kiện.


