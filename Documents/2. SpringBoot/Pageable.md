
---

### Các khái niệm liên quan:

- **Pageable**: Giao diện đại diện cho thông tin phân trang (trang hiện tại, số lượng phần tử trên mỗi trang, sắp xếp, v.v.).
- **Page**: Đối tượng chứa dữ liệu của một trang cùng với các thông tin về tổng số trang, trang hiện tại, tổng số phần tử, v.v.
- **Sort**: Cung cấp khả năng sắp xếp dữ liệu dựa trên một hoặc nhiều thuộc tính theo thứ tự tăng hoặc giảm.

---
Repository:
```Java
public interface UserRepostiory extends JpaRepository<UserEntity, Long> JpaSpecificationExecutor<UserEntity>{

Page<UserEntity> findByUserNameContain(String name, Pageable pageable);


}
```

Service:

UserService:
```Java
@Service
public interface UserService{
Page<UserEntity> findAllUsers(Pageable pageable);

Page<UserEntity> findByUserNameContain(String name, Pageable pageable);
}
```
UserServiceImpl:
```Java
public class UserServiceImpl implements UserService{
@Override
public Page<UserEntity> findAllUsers(Pageable pageable){
return userRepository.findAllUsers(pageable);
}

@Override
public Page<UserEntity> findByUserName(String userName, Pageable pageable){
return userRepository.findByUserNameContain(userName, pageable);
}

}
```

Controller:
UserCRUDController
```Java
//Sort tang dan
@GetMapping("/getAll)
public Page<UserEntity> getAll(
@RequestParam int page,
@RequestParam int size,
@RequestParam(defaultValue = "id") String sort,
@RequestParam(defaultValue = "desc") String direction,
){
Sort.Direction sortDirection = direction.equalsIgnoreCase("asc") ? Sort.Direction.ASC : Sort.Direction.DESC;
Sort sortBy = Sort.by(sortDirection, sort);
Pageable pageable = PageRequest.of(page, size, sortBy);
return userService.findAllUser(pageable);
}

@GetMapping("/searchPage)
public Page<UserEntity> getAll(
@RequestParam String name,
@RequestParam int page,
@RequestParam int size,
@RequestParam(defaultValue = "id") String sort,
@RequestParam(defaultValue = "desc") String direction,
){
Sort.Direction sortDirection = direction.equalsIgnoreCase("asc") ? Sort.Direction.ASC : Sort.Direction.DESC;
Sort sortBy = Sort.by(sortDirection, sort);
Pageable pageable = PageRequest.of(page, size, sortBy);
return userService.findByUserNameContain(name, pageable);
}
```

#### 1. **API: `/getAll`**

API này trả về danh sách tất cả các user với phân trang và sắp xếp.

- **Tham số**:
    
    - `page`: Số trang bắt đầu từ 0.
    - `size`: Số lượng phần tử trên mỗi trang.
    - `sort`: Thuộc tính dùng để sắp xếp (mặc định là "id").
    - `direction`: Hướng sắp xếp, tăng dần ("asc") hoặc giảm dần ("desc").
- **Logic**:
    
    - Dựa vào `direction`, API kiểm tra xem sắp xếp tăng dần (ASC) hay giảm dần (DESC).
    - Tạo đối tượng `Pageable` bằng cách sử dụng `PageRequest.of(page, size, sortBy)` để phân trang và sắp xếp.
    - Gọi phương thức `userService.findAllUser(pageable)` để trả về kết quả được phân trang.

#### 2. **API: `/searchPage`**

API này tìm kiếm các user dựa trên tên chứa một chuỗi ký tự nhất định và cũng hỗ trợ phân trang và sắp xếp.

- **Tham số**:
    
    - `name`: Chuỗi ký tự mà tên user phải chứa.
    - `page`, `size`, `sort`, `direction`: Giống như trong API `/getAll`.
- **Logic**:
    
    - Tương tự như API đầu tiên, kiểm tra và xác định hướng sắp xếp.
    - Sử dụng `PageRequest.of(page, size, sortBy)` để tạo đối tượng `Pageable`.
    - Gọi phương thức `userService.findByUserNameContain(name, pageable)` để trả về kết quả được phân trang với tên chứa chuỗi `name`.

### CURL

1. **Test API `/getAll`**:

`curl -X GET "http://localhost:8080/getAll?page=0&size=10&sort=id&direction=desc"`

Lệnh này sẽ gọi API `/getAll`, yêu cầu trang đầu tiên (`page=0`), với 10 phần tử trên mỗi trang (`size=10`), sắp xếp theo `id` theo thứ tự giảm dần (`direction=desc`).

2. **Test API `/searchPage`**:

`curl -X GET "http://localhost:8080/searchPage?name=John&page=0&size=5&sort=name&direction=asc"`

Lệnh này sẽ gọi API `/searchPage`, tìm các user có tên chứa chuỗi "John", trang đầu tiên (`page=0`), với 5 phần tử trên mỗi trang (`size=5`), sắp xếp theo `name` theo thứ tự tăng dần (`direction=asc`).