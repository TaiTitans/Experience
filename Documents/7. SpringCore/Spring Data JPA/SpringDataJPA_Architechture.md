
---
## **1. Kiến Trúc và Cơ Chế Hoạt Động của Spring Data JPA**

Spring Data JPA là một module của **Spring Data** giúp đơn giản hóa việc tương tác với cơ sở dữ liệu bằng cách sử dụng **JPA (Java Persistence API)**. Nó cung cấp một lớp trừu tượng để giảm bớt sự phức tạp khi sử dụng **EntityManager**, đồng thời hỗ trợ tạo các truy vấn động mà không cần viết nhiều code SQL hoặc JPQL.

Dưới đây là các phần quan trọng trong kiến trúc và cách Spring Data JPA hoạt động:

## **1.1. Spring Data JPA là gì?**

### **Mối quan hệ giữa Spring Data JPA, JPA, Hibernate**

- **JPA (Java Persistence API)**: Một **specification** (chuẩn) của Java EE để làm việc với ORM (Object-Relational Mapping). Nó chỉ định các quy tắc nhưng không cung cấp một implementation cụ thể.
- **Hibernate**: Một implementation phổ biến của JPA, cung cấp nhiều tính năng nâng cao so với chuẩn JPA.
- **Spring Data JPA**: Một abstraction trên JPA giúp đơn giản hóa việc tương tác với Hibernate (hoặc các provider khác như EclipseLink).

📌 **Tóm lại**:

- Spring Data JPA **không thay thế Hibernate** mà hoạt động **bên trên** JPA và Hibernate để giúp code dễ đọc và bảo trì hơn.
- Bạn vẫn có thể sử dụng **Native Query, JPQL, Criteria API** nếu cần truy vấn phức tạp.

---
## **1.2. Cấu Trúc Hoạt Động của Spring Data JPA**

Spring Data JPA hoạt động dựa trên các thành phần chính sau:

### **1.2.1. Entity (Model)**

- Là đại diện của bảng trong database, được ánh xạ bằng annotation **@Entity**.
- Có thể sử dụng **@Table, @Column, @Id, @GeneratedValue** để tùy chỉnh mapping.

Ví dụ:
```java
@Entity
@Table(name = "users")
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false, unique = true)
    private String username;

    @Column(nullable = false)
    private String password;

    // Getters and Setters
}
```

### **1.2.2. Repository Layer**

- Spring Data JPA cung cấp các interface như **JpaRepository, CrudRepository, PagingAndSortingRepository** để thao tác với database mà không cần viết SQL.
- Spring tự động **generate implementation** của repository dựa trên interface.

Ví dụ:
```java
public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByUsername(String username);
}
```
Phương thức `findByUsername(String username)` sẽ được Spring Data JPA **tự động sinh ra câu query** như sau:
`SELECT * FROM users WHERE username = ?;`

### **1.2.3. Service Layer**

- Là lớp trung gian giữa **Repository** và **Controller** để xử lý logic nghiệp vụ.
- Có thể sử dụng **@Transactional** để đảm bảo các thao tác database được thực hiện nguyên tử.

Ví dụ:
```java
@Service
public class UserService {
    private final UserRepository userRepository;

    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @Transactional
    public User createUser(User user) {
        return userRepository.save(user);
    }
}
```

### **1.2.4. Controller Layer**

- Là lớp tiếp nhận request từ client và gọi service để xử lý.

Ví dụ:
```java
@RestController
@RequestMapping("/users")
public class UserController {
    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @PostMapping
    public ResponseEntity<User> createUser(@RequestBody User user) {
        return ResponseEntity.ok(userService.createUser(user));
    }
}
```

## **1.3. EntityManager và Unit of Work**

### **1.3.1. EntityManager là gì?**

Trong JPA, `EntityManager` là một API chính để thao tác với các entity trong database. Nó chịu trách nhiệm cho:

- **Thêm, cập nhật, xóa entity**.
- **Quản lý transaction**.
- **Quản lý vòng đời của entity**.

📌 Khi sử dụng **Spring Data JPA**, bạn **không cần thao tác trực tiếp** với `EntityManager` vì Spring đã trừu tượng hóa nó thông qua `JpaRepository`.

### **1.3.2. Unit of Work Pattern**

- **Unit of Work** là một **design pattern** giúp đảm bảo các thay đổi được thực hiện trong **một transaction duy nhất**.
- Trong JPA, `EntityManager` quản lý **Persistence Context**, nghĩa là:
    - Nếu bạn lấy một entity từ database, sửa nó, và gọi `save()`, nó sẽ tự động **merge** entity đó mà không cần gọi `update`.

Ví dụ:
```java
@Transactional
public void updateUser(Long userId, String newUsername) {
    User user = userRepository.findById(userId).orElseThrow();
    user.setUsername(newUsername); 
    // Không cần gọi save() vì entity đã nằm trong Persistence Context
}
```

🔥 **Lưu ý:** Nếu `@Transactional` không có, `EntityManager` sẽ không quản lý entity, và bạn phải gọi `save()` để cập nhật dữ liệu.

## **1.4. Transactional Context trong Spring Data JPA**

### **1.4.1. @Transactional Hoạt Động Như Thế Nào?**

- `@Transactional` đảm bảo các thao tác trong method sẽ được thực hiện **nguyên tử** (atomic).
- Nếu có lỗi xảy ra, toàn bộ giao dịch sẽ **rollback** thay vì cập nhật một phần dữ liệu.

Ví dụ:
```java
@Transactional
public void transferMoney(Long fromAccountId, Long toAccountId, double amount) {
    Account fromAccount = accountRepository.findById(fromAccountId).orElseThrow();
    Account toAccount = accountRepository.findById(toAccountId).orElseThrow();
    
    fromAccount.setBalance(fromAccount.getBalance() - amount);
    toAccount.setBalance(toAccount.getBalance() + amount);
}
```
🔥 Nếu có lỗi xảy ra giữa chừng (ví dụ: database lỗi), cả hai tài khoản sẽ không bị thay đổi.

### **1.4.2. Propagation trong Transaction**

Spring hỗ trợ nhiều kiểu propagation cho `@Transactional`, phổ biến gồm:

- `REQUIRED` (Mặc định): Nếu đã có transaction, sẽ dùng transaction hiện tại. Nếu chưa có, Spring sẽ tạo mới.
- `REQUIRES_NEW`: Luôn tạo một transaction mới, bất kể có transaction hay không.
- `MANDATORY`: Bắt buộc phải có transaction, nếu không có sẽ throw Exception.
- `SUPPORTS`: Nếu có transaction thì dùng, nếu không có thì chạy bình thường.
- `NOT_SUPPORTED`: Luôn chạy ngoài transaction.
- `NEVER`: Nếu có transaction sẽ throw Exception.

Ví dụ:
```java
@Transactional(propagation = Propagation.REQUIRES_NEW)
public void createOrder(Order order) {
    orderRepository.save(order);
}
```

## **Tóm Tắt**

✅ **Spring Data JPA** giúp đơn giản hóa việc thao tác database bằng cách trừu tượng hóa **JPA & Hibernate**.  
✅ **EntityManager** giúp quản lý entity và thực hiện **Unit of Work Pattern**.  
✅ **@Transactional** giúp đảm bảo dữ liệu nhất quán bằng cách quản lý **transactions** tự động.  
✅ **Propagation** trong transaction cho phép kiểm soát cách một transaction được thực thi.


