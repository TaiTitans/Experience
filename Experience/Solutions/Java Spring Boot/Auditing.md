
----
Auditing đơn giản chỉ là việc cập nhật giá trị `created_at` và `updated_at` của đối tượng trong database.

### @EnableJpaAuditing

```Java

@SpringBootApplication
// Bạn phải kích hoạt chức năng Auditing bằng annotation @EnableJpaAuditing
@EnableJpaAuditing
public class App {
    public static void main(String[] args) {
        SpringApplication.run(App.class, args);
    }
}
```


### @EntityListeners

**Spring JPA** tự động lắng nghe các sự kiện này và xử lý giúp chúng ta bằng class `AuditingEntityListener`
```Java
@Data
@Entity
@Table(name = "app_params")
// Bạn phải đánh dấu class bởi @EntityListeners(AuditingEntityListener.class)
// Đây là Một đối tượng Listener, lắng nghe sự kiện insert hoặc update của đối tượng
// Để từ đó tự động cập nhật các trường @CreatedDate và @LastModifiedDate
@EntityListeners(AuditingEntityListener.class)
public class AppParams {
}
```

### @CreatedDate & @LastModifiedDate

`@CreatedDate` và `@LastModifiedDate` được chú thích trên trường `Date` của `Entity`

Khi `Entity` insert vào database, JPA sẽ cập nhật giá trị của trường có chú thích bởi `@CreatedDate`

Khi `Entity` insert hoặc update vào database, JPA sẽ cập nhật giá trị của trường có chú thích bởi `@LastModifiedDate`.

```Java
@EntityListeners(AuditingEntityListener.class)
public class AppParams {
    /*
    đánh dấu trường Date với @CreatedDate
    Nếu đối tượng được insert vào database lần đầu -> nó sẽ tự động lấy thời điểm đó đưa vào createdAt
     */
    @CreatedDate
    private Date createdAt;

    /*
    đánh dấu trường Date với @LastModifiedDate
    Nếu đối tượng được insert vào database lần đầu -> nó sẽ tự động lấy thời điểm đó đưa vào updatedAt
    Nếu đối tượng được update vào database lần tiếp theo -> nó sẽ tự động lấy thời điểm đó cập nhật vào updatedAt
     */
    @LastModifiedDate
    private Date updatedAt;
}

```