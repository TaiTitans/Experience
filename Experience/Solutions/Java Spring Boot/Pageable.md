
---
Ví dụ triển khai:
Tạo ra class `User` và insert sẵn 100 `User` vào Database.

_User.java_

```Java
@Entity
@Data
@Builder
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String name;
}
```

_UserRepository.java_

```Java
public interface UserRepository extends JpaRepository<User, Long> {

}
```

_DatasourceConfig.java_

```Java
@Configuration
@RequiredArgsConstructor
public class DatasourceConfig {
    // inject bởi RequiredArgsConstructor
    private final UserRepository userRepository;

    // Chỉ áp dụng trong demo :D 
    @PostConstruct
    public void initData() {
        // Insert 100 User vào H2 Database sau khi
        // DatasourceConfig được khởi tạo
        userRepository.saveAll(IntStream.range(0, 100)
                                        .mapToObj(i -> User.builder()
                                                           .name("name-" + i)
                                                           .build())
                                        .collect(Collectors.toList())
        );
    }
}
```

Để có thể query lấy dữ liệu theo dạng Page, **Spring Data JPA** hỗ trợ chúng ta bằng đối tượng `Pageable`.

Hàm `findAll(Pageable pageable)` là có sẵn và trả về đối tượng `Page<T>`

```Java
// Lấy ra 5 user đầu tiên
// PageRequest.of(0,5) tương đương với lấy ra page đầu tiên, và mỗi page sẽ có 5 phần tử
// PageRequest là một đối tượng kế thừa Pageable
Page<User> page = userRepository.findAll(PageRequest.of(0, 5));
```

Để lấy ra 5 `User` tiếp theo, chúng ta có thể làm theo 2 cách:

```Java
// tận dụng đối tượng Page trước đó
page.nextPageable()

// Sử dụng PageRequest mới
PageRequest.of(1, 5)
```

### Sorting

Bạn có thể query dạng `Page` kèm theo yêu cầu sorting theo một trường nào đó.

```Java
Page<User> page = userRepository.findAll(PageRequest.of(1, 5, Sort.by("name").descending()));
```

### Note

Ngoài ra, bạn hoàn toàn có thể tự custom để hàm trả về `Page<T>`, `Slice<T>`, `List<T>`.

### DEMO:

```Java
@SpringBootApplication
@RequiredArgsConstructor
public class App {
    public static void main(String[] args) {
        SpringApplication.run(App.class, args);
    }

    private final UserRepository userRepository;

    @Bean
    CommandLineRunner run() {
        return args -> {
            // Lấy ra 5 user đầu tiên
            // PageRequest.of(0,5) tương đương với lấy ra page đầu tiên, và mỗi page sẽ có 5 phần tử
            Page<User> page = userRepository.findAll(PageRequest.of(0, 5));
            // In ra 5 user đầu tiên
            System.out.println("In ra 5 user đầu tiên: ");
            page.forEach(System.out::println);
            // Lấy ra 5 user tiếp theo
            page = userRepository.findAll(page.nextPageable());

            System.out.println("In ra 5 user tiếp theo: ");
            page.forEach(System.out::println);

            System.out.println("In ra số lượng user ở page hiện tại: " + page.getSize());
            System.out.println("In ra tổng số lượng user: " + page.getTotalElements());
            System.out.println("In ra tổng số page: " + page.getTotalPages());

            // Lấy ra 5 user ở page 1, sort theo tên
            page = userRepository.findAll(PageRequest.of(1, 5, Sort.by("name").descending()));
            System.out.println("In ra 5 user page 1, sắp xếp theo name descending:");
            page.forEach(System.out::println);

            // Custom method
            List<User> list = userRepository.findAllByNameLike("name-%", PageRequest.of(0, 5));
            System.out.println(list);
        };
    }

}

```

### Kết

`Pageable` là một cách tiếp cận rất tiện lợi, tuy nhiên bạn cần biết rằng các hàm sử dụng `Pageable` sẽ thực hiện query 2 lần xuống DB. Một lần là để lấy ra tổng số lượng bản ghi, một lần là query lấy ra page bạn cần