
---
**MapStruct** là một công cụ mapping (ánh xạ) Java mạnh mẽ được sử dụng để chuyển đổi dữ liệu giữa các loại đối tượng khác nhau, chẳng hạn từ **DTO (Data Transfer Object)** sang **Entity** và ngược lại. Đặc điểm nổi bật của MapStruct là nó tạo mã mapping tự động trong quá trình biên dịch, mang lại hiệu suất cao và không cần phụ thuộc vào thời gian chạy.

### 1. **Lợi Ích Của MapStruct**

- **Hiệu Suất Cao**: Mã mapping được tạo tại thời điểm biên dịch, vì vậy không có chi phí hiệu năng tại runtime.
- **Ít Lỗi Hơn**: Các mapping được xác định rõ ràng và kiểm tra tại compile-time, giúp phát hiện lỗi sớm.
- **Dễ Bảo Trì**: Giảm thiểu mã boilerplate bằng cách sử dụng các annotation, giúp code dễ bảo trì hơn.
- **Tích Hợp Tốt**: Tích hợp dễ dàng với Spring Framework và các framework khác.
### C**ách Sử Dụng MapStruct**

####  **Định Nghĩa DTO và Entity**

Giả sử bạn có hai lớp: `UserDTO` và `UserEntity`.
```java
public class UserDTO {
    private Long id;
    private String name;
    private String email;
    // getters and setters
}

public class UserEntity {
    private Long userId;
    private String fullName;
    private String emailAddress;
    // getters and setters
}
```
#### **Tạo Mapper Interface**

MapStruct sử dụng các interface để định nghĩa mapping giữa các đối tượng.
```java
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.factory.Mappers;

@Mapper
public interface UserMapper {
    UserMapper INSTANCE = Mappers.getMapper(UserMapper.class);

    @Mapping(source = "userId", target = "id")
    @Mapping(source = "fullName", target = "name")
    @Mapping(source = "emailAddress", target = "email")
    UserDTO toUserDTO(UserEntity userEntity);

    @Mapping(source = "id", target = "userId")
    @Mapping(source = "name", target = "fullName")
    @Mapping(source = "email", target = "emailAddress")
    UserEntity toUserEntity(UserDTO userDTO);
}
```
### **Chạy MapStruct**

Khi bạn biên dịch dự án, MapStruct sẽ tự động tạo ra một lớp triển khai của `UserMapper` với tên `UserMapperImpl`. Bạn có thể sử dụng mapper này như sau:
```java
public class Main {
    public static void main(String[] args) {
        UserEntity userEntity = new UserEntity();
        userEntity.setUserId(1L);
        userEntity.setFullName("John Doe");
        userEntity.setEmailAddress("john.doe@example.com");

        UserDTO userDTO = UserMapper.INSTANCE.toUserDTO(userEntity);
        System.out.println(userDTO.getName());  // Output: John Doe
    }
}
```
### **Các Tính Năng Nâng Cao**

####  **Mapping Tự Động**

Nếu các thuộc tính của DTO và Entity có cùng tên, MapStruct sẽ tự động ánh xạ chúng mà không cần chỉ định `@Mapping`.
```java
@Mapper
public interface UserMapper {
    UserDTO toUserDTO(UserEntity userEntity);
    UserEntity toUserEntity(UserDTO userDTO);
}
```
#### **Mapping Phức Tạp**

- **Nested Objects**: Ánh xạ các thuộc tính trong các đối tượng lồng nhau.
```java
public class AddressDTO {
    private String city;
    private String country;
    // getters and setters
}

public class UserDTO {
    private Long id;
    private String name;
    private AddressDTO address;
    // getters and setters
}
```

```java
@Mapper
public interface UserMapper {
    UserDTO toUserDTO(UserEntity userEntity);
}
```
**Ánh xạ danh sách**:
```java
List<UserDTO> toUserDTOList(List<UserEntity> userEntities);
```
#### **Custom Mapping Logic**

Bạn có thể cung cấp logic ánh xạ tùy chỉnh bằng cách thêm phương thức `default` hoặc `expression`:
```java
@Mapping(target = "fullName", expression = "java(user.getFirstName() + \" \" + user.getLastName())")
UserDTO toUserDTO(UserEntity userEntity);
```
#### **Reverse Mapping và Update Mapping**

- **Reverse Mapping**: Dễ dàng tạo các ánh xạ ngược bằng cách sử dụng `@InheritInverseConfiguration`.
```java
@Mapper
public interface UserMapper {
    @Mapping(source = "userId", target = "id")
    @Mapping(source = "fullName", target = "name")
    @Mapping(source = "emailAddress", target = "email")
    UserDTO toUserDTO(UserEntity userEntity);

    @InheritInverseConfiguration
    UserEntity toUserEntity(UserDTO userDTO);
}
```
**Update Mapping**: Cập nhật một đối tượng hiện có.
```java
@Mapping(target = "id", ignore = true)
void updateUserEntityFromDTO(UserDTO userDTO, @MappingTarget UserEntity userEntity);
```
### **Tích Hợp với Spring**

MapStruct có thể được tích hợp với Spring thông qua annotation `@Mapper` với `componentModel` là `"spring"`:
```java
@Mapper(componentModel = "spring")
public interface UserMapper {
    UserDTO toUserDTO(UserEntity userEntity);
}
```
Sau đó, bạn có thể `@Autowired` mapper trong các lớp Spring:
```java
@Service
public class UserService {

    @Autowired
    private UserMapper userMapper;

    public UserDTO getUserDTO(UserEntity userEntity) {
        return userMapper.toUserDTO(userEntity);
    }
}
```
### **Hiệu Năng và Ưu Nhược Điểm**

#### **Ưu Điểm**:

- **Hiệu Suất Cao**: Không có chi phí runtime vì mã được tạo tại compile-time.
- **Dễ Bảo Trì**: Mã rõ ràng và dễ dàng bảo trì.
- **Phát Hiện Lỗi Sớm**: Lỗi được phát hiện sớm trong quá trình biên dịch.

#### **Nhược Điểm**:

- **Hạn Chế về Cấu Trúc Phức Tạp**: Khi ánh xạ các cấu trúc rất phức tạp, có thể cần viết nhiều mã mapping tùy chỉnh.
- **Cần Biên Dịch Lại**: Nếu cấu trúc DTO hoặc Entity thay đổi, cần phải biên dịch lại dự án để cập nhật mapping.