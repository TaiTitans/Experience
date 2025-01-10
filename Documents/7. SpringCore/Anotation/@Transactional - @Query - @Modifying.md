
---
Dưới đây là giải thích chi tiết về các annotation `@Transactional`, `@Query`, và `@Modifying` trong Spring Framework:

### 1. **`@Transactional`**

`@Transactional` là một annotation trong Spring được sử dụng để quản lý giao dịch (transaction) trong các phương thức của bean.

- **Mục đích**:
    - Quản lý giao dịch tự động, đảm bảo rằng các thao tác trên cơ sở dữ liệu được thực hiện như một đơn vị công việc duy nhất.
    - Hỗ trợ rollback khi có lỗi xảy ra trong quá trình thực hiện giao dịch.

**Cách sử dụng**:
```java
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class MyService {

    @Transactional
    public void performTransactionalOperation() {
        // Thao tác với cơ sở dữ liệu
        // Nếu có exception xảy ra, transaction sẽ bị rollback
    }
}
```
- **Giải thích**:
    
    - `@Transactional`: Spring sẽ bắt đầu một giao dịch trước khi thực hiện phương thức và kết thúc giao dịch sau khi phương thức hoàn tất.
    - Nếu có exception được ném ra trong quá trình thực hiện, giao dịch sẽ được rollback.
- **Lợi ích**:
    
    - Đảm bảo tính toàn vẹn của dữ liệu.
    - Giảm thiểu lỗi xảy ra khi các thao tác cơ sở dữ liệu không hoàn tất đầy đủ.
### 2. **`@Query`**

`@Query` là một annotation trong Spring Data JPA được sử dụng để định nghĩa các câu truy vấn SQL hoặc JPQL (Java Persistence Query Language) trực tiếp trong repository.

- **Mục đích**:
    - Tạo các truy vấn tùy chỉnh thay vì sử dụng các phương thức truy vấn được tạo tự động.
    - Hỗ trợ cả JPQL và Native SQL.

**Cách sử dụng**:
```java
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;

public interface MyRepository extends CrudRepository<MyEntity, Long> {

    @Query("SELECT m FROM MyEntity m WHERE m.name = ?1")
    List<MyEntity> findByName(String name);

    @Query(value = "SELECT * FROM my_entity WHERE name = ?1", nativeQuery = true)
    List<MyEntity> findByNameNative(String name);
}
```
- **Giải thích**:
    
    - `@Query("SELECT m FROM MyEntity m WHERE m.name = ?1")`: Sử dụng JPQL để tìm kiếm các thực thể theo tên.
    - `@Query(value = "SELECT * FROM my_entity WHERE name = ?1", nativeQuery = true)`: Sử dụng native SQL để thực hiện truy vấn trực tiếp trên bảng cơ sở dữ liệu.
- **Lợi ích**:
    
    - Cung cấp linh hoạt trong việc tạo các truy vấn phức tạp.
    - Hỗ trợ cả JPQL và Native SQL.
### 3. **`@Modifying`**

`@Modifying` là một annotation được sử dụng cùng với `@Query` để thực hiện các thao tác cập nhật hoặc xóa dữ liệu.

- **Mục đích**:
    - Đánh dấu các phương thức trong repository sử dụng `@Query` để thực hiện các thao tác thay đổi dữ liệu như `UPDATE`, `DELETE`.

**Cách sử dụng**:
```java
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.transaction.annotation.Transactional;

public interface MyRepository extends CrudRepository<MyEntity, Long> {

    @Modifying
    @Transactional
    @Query("UPDATE MyEntity m SET m.name = ?1 WHERE m.id = ?2")
    int updateNameById(String name, Long id);

    @Modifying
    @Transactional
    @Query("DELETE FROM MyEntity m WHERE m.id = ?1")
    void deleteById(Long id);
}
```
- **Giải thích**:
    
    - `@Modifying`: Đánh dấu phương thức thực hiện các thao tác thay đổi dữ liệu (không phải truy vấn).
    - `@Transactional`: Cần thiết để quản lý giao dịch cho các thao tác cập nhật hoặc xóa.
    - `int updateNameById(String name, Long id)`: Trả về số dòng bị ảnh hưởng bởi câu lệnh `UPDATE`.
- **Lợi ích**:
    
    - Hỗ trợ thực hiện các thao tác cập nhật, xóa dữ liệu trực tiếp từ repository.
    - Cung cấp cách làm việc linh hoạt với các câu lệnh SQL tùy chỉnh.