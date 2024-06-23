	https://github.com/mapstruct/mapstruct
---

1. **Add MapStruct Dependency**: First, you need to add the MapStruct dependency to your project. If you're using Maven, you can add the following dependency to your `pom.xml` file:
```xml
<dependency>
    <groupId>org.mapstruct</groupId>
    <artifactId>mapstruct</artifactId>
    <version>1.5.3.Final</version>
</dependency>

```

2. **Create Source and Target Classes**: Suppose you have a `User` class and a `UserDTO` class that represent the same data but with different field names or structures. You want to map between these two classes:
```Java
// User class
public class User {
    private long id;
    private String firstName;
    private String lastName;
    private String email;
    // getters and setters
}

// UserDTO class
public class UserDTO {
    private long userId;
    private String name;
    private String userEmail;
    // getters and setters
}
```

3. **Create a Mapper Interface**: Create an interface that will be used by MapStruct to generate the mapping code. This interface should extend the `Mapper` interface provided by MapStruct.
```Java
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Mappings;
import org.mapstruct.factory.Mappers;

@Mapper
public interface UserMapper {
    UserMapper INSTANCE = Mappers.getMapper(UserMapper.class);

    @Mappings({
        @Mapping(source = "id", target = "userId"),
        @Mapping(source = "firstName", target = "name"),
        @Mapping(source = "email", target = "userEmail")
    })
    UserDTO userToUserDTO(User user);

    @Mappings({
        @Mapping(source = "userId", target = "id"),
        @Mapping(source = "name", target = "firstName"),
        @Mapping(source = "userEmail", target = "email")
    })
    User userDTOToUser(UserDTO userDTO);
}
```

4. **Use the Generated Mapper**: Now, you can use the generated `UserMapper` to perform the mapping between `User` and `UserDTO` objects.
```Java
User user = new User(1L, "John", "Doe", "john.doe@example.com");
UserDTO userDTO = UserMapper.INSTANCE.userToUserDTO(user);

// userDTO now contains the mapped values
System.out.println(userDTO.getUserId()); // 1
System.out.println(userDTO.getName()); // "John Doe"
System.out.println(userDTO.getUserEmail()); // "john.doe@example.com"
```

---
## Advance

OneToMany and ManyToOne

Example:

```Java
// Department class
public class Department {
    private Long id;
    private String name;
    private List<Employee> employees;
    // getters, setters, and constructors
}

// Employee class
public class Employee {
    private Long id;
    private String name;
    private Department department;
    // getters, setters, and constructors
}
```

And:

```Java
// DepartmentDTO class
public class DepartmentDTO {
    private Long id;
    private String name;
    private List<EmployeeDTO> employees;
    // getters, setters, and constructors
}

// EmployeeDTO class
public class EmployeeDTO {
    private Long id;
    private String name;
    private Long departmentId;
    // getters, setters, and constructors
}
```


Result:

```Java
@Mapper
public interface DepartmentMapper {
    DepartmentMapper INSTANCE = Mappers.getMapper(DepartmentMapper.class);

    @Mapping(source = "employees", target = "employees")
    DepartmentDTO departmentToDepartmentDTO(Department department);

    @InverseMapping
    Department departmentDTOToDepartment(DepartmentDTO departmentDTO);

    @Mapping(source = "department.id", target = "departmentId")
    EmployeeDTO employeeToEmployeeDTO(Employee employee);

    @InverseMapping
    Employee employeeDTOToEmployee(EmployeeDTO employeeDTO);
}
```

