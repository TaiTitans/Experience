
---
**JUnit** là một framework kiểm thử phổ biến cho Java, hỗ trợ kiểm thử đơn vị (**Unit Testing**), kiểm thử tích hợp (**Integration Testing**) và có thể kết hợp với **Mockito** để mock dependencies.

- **JUnit 4**: Phiên bản cũ, sử dụng annotation như `@Test`, `@Before`, `@After`.
- **JUnit 5** (Jupiter): Cải tiến hơn với nhiều tính năng linh hoạt, tách thành 3 phần chính:
    - **JUnit Platform**: Cung cấp môi trường chạy test.
    - **JUnit Jupiter**: Framework chính của JUnit 5.
    - **JUnit Vintage**: Hỗ trợ chạy test của JUnit 4.

## 🔥 **2. Cấu trúc cơ bản của JUnit 5**

Cài đặt JUnit 5 (Spring Boot có sẵn trong `spring-boot-starter-test`):
```java
<dependency>
    <groupId>org.junit.jupiter</groupId>
    <artifactId>junit-jupiter-api</artifactId>
    <scope>test</scope>
</dependency>
```
Ví dụ test cơ bản:
```java
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class MathTest {
    @Test
    void testAddition() {
        int sum = 2 + 3;
        assertEquals(5, sum);
    }
}
```
## ✅ **3. Các Annotation Quan Trọng**

| Annotation           | Ý nghĩa                                                     |
| -------------------- | ----------------------------------------------------------- |
| `@Test`              | Đánh dấu một method là test case                            |
| `@BeforeEach`        | Chạy trước mỗi test (tương đương `@Before` của JUnit 4)     |
| `@AfterEach`         | Chạy sau mỗi test (tương đương `@After` của JUnit 4)        |
| `@BeforeAll`         | Chạy một lần trước tất cả test (tương đương `@BeforeClass`) |
| `@AfterAll`          | Chạy một lần sau tất cả test (tương đương `@AfterClass`)    |
| `@Disabled`          | Bỏ qua test                                                 |
| `@ParameterizedTest` | Dùng để chạy test với nhiều giá trị đầu vào                 |
Ví dụ:
```java
import org.junit.jupiter.api.*;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class LifecycleTest {

    @BeforeAll
    void setupAll() {
        System.out.println("Chạy trước tất cả test");
    }

    @BeforeEach
    void setup() {
        System.out.println("Chạy trước mỗi test");
    }

    @Test
    void test1() {
        System.out.println("Test 1");
    }

    @Test
    void test2() {
        System.out.println("Test 2");
    }

    @AfterEach
    void tearDown() {
        System.out.println("Chạy sau mỗi test");
    }

    @AfterAll
    void tearDownAll() {
        System.out.println("Chạy sau tất cả test");
    }
}
```
🔹 **Kết quả chạy test**:
```bash
Chạy trước tất cả test
Chạy trước mỗi test
Test 1
Chạy sau mỗi test
Chạy trước mỗi test
Test 2
Chạy sau mỗi test
Chạy sau tất cả test
```
## 🔄 **4. Test Driven Development (TDD)**

### **Nguyên tắc TDD:**

1. **Viết test trước** (test case fail vì chưa có code).
2. **Viết code đủ để test pass**.
3. **Refactor** để tối ưu code và test.

Ví dụ áp dụng TDD cho `Calculator`: 1️⃣ **Viết test trước**:
```java
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class CalculatorTest {

    @Test
    void testAddition() {
        Calculator calc = new Calculator();
        assertEquals(5, calc.add(2, 3));
    }
}
```
2️⃣ **Viết code cho test pass**:
```java
class Calculator {
    int add(int a, int b) {
        return a + b;
    }
}
```
3️⃣ **Refactor nếu cần**.
## 🔄 **5. Mocking với Mockito**

Trong thực tế, khi test service có database hoặc API call, ta cần **mock** dependencies.

Cài đặt Mockito:
```java
<dependency>
    <groupId>org.mockito</groupId>
    <artifactId>mockito-core</artifactId>
    <scope>test</scope>
</dependency>
```
Ví dụ:
```java
import static org.mockito.Mockito.*;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

class UserServiceTest {
    @Test
    void testGetUserById() {
        UserRepository mockRepo = mock(UserRepository.class);
        when(mockRepo.findById(1L)).thenReturn(new User(1L, "John Doe"));

        UserService userService = new UserService(mockRepo);
        User user = userService.getUserById(1L);

        assertEquals("John Doe", user.getName());
    }
}
```
## 🚀 **6. Test Controller trong Spring Boot**

Khi test REST API, ta dùng `@WebMvcTest` + `MockMvc` để mock request:
```java
@WebMvcTest(UserController.class)
class UserControllerTest {
    @Autowired private MockMvc mockMvc;
    @MockBean private UserService userService;

    @Test
    void testGetUser() throws Exception {
        when(userService.getUserById(1L)).thenReturn(new User(1L, "John Doe"));

        mockMvc.perform(get("/users/1"))
               .andExpect(status().isOk())
               .andExpect(jsonPath("$.name").value("John Doe"));
    }
}
```
## 🏆 **7. Test Database với Spring Boot & Testcontainers**

**Testcontainers** giúp kiểm thử với DB thật (PostgreSQL, MySQL, Redis, etc.). Cài đặt:
```java
<dependency>
    <groupId>org.testcontainers</groupId>
    <artifactId>postgresql</artifactId>
    <scope>test</scope>
</dependency>
```
Ví dụ:
```java
@Testcontainers
@SpringBootTest
class UserRepositoryTest {
    @Container
    static PostgreSQLContainer<?> postgres = new PostgreSQLContainer<>("postgres:14");

    @Autowired UserRepository userRepository;

    @Test
    void testSaveUser() {
        User user = new User(null, "John Doe");
        User savedUser = userRepository.save(user);

        assertNotNull(savedUser.getId());
        assertEquals("John Doe", savedUser.getName());
    }
}
```
## 🔥 **8. Coverage Report với JaCoCo**

JaCoCo giúp đo **coverage (mức độ bao phủ của test)**.

Cài đặt:
```java
<plugin>
    <groupId>org.jacoco</groupId>
    <artifactId>jacoco-maven-plugin</artifactId>
    <version>0.8.8</version>
    <executions>
        <execution>
            <goals><goal>prepare-agent</goal></goals>
        </execution>
    </executions>
</plugin>
```
Chạy:
```shell
mvn test
mvn jacoco:report
```
Mở file `target/site/jacoco/index.html` để xem báo cáo coverage.

## 📌 **Kết luận**

- **Viết unit test tốt** giúp giảm bug, cải thiện chất lượng code.
- **Mockito giúp mock dependencies** khi test service.
- **MockMvc giúp test API** dễ dàng.
- **Testcontainers giúp test với DB thật**.
- **JaCoCo đo coverage** để cải thiện test.