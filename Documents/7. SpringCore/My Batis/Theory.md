
---

## **1. Kiến trúc MyBatis**

### **1.1. Cấu trúc MyBatis**

MyBatis được chia thành 4 thành phần chính:

- **Core (Lõi):** Phần chính của MyBatis chịu trách nhiệm quản lý toàn bộ luồng dữ liệu.
    
- **SQL Mapping (Ánh xạ SQL):** Kết nối giữa câu lệnh SQL và đối tượng Java thông qua file XML hoặc annotation.
    
- **Session (SqlSession):** Là giao diện để tương tác với cơ sở dữ liệu.
    
- **Caching (Bộ nhớ đệm):** Hỗ trợ caching để tăng hiệu năng truy vấn dữ liệu.

---
### **1.2. Quy trình hoạt động**

Quy trình làm việc của MyBatis gồm các bước:

1. **Cấu hình:** MyBatis đọc file cấu hình (`mybatis-config.xml`) để thiết lập thông tin cơ sở dữ liệu, mapper, caching, và các cài đặt khác.
    
2. **SqlSessionFactory:** MyBatis sử dụng **SqlSessionFactoryBuilder** để tạo ra các phiên làm việc (**SqlSession**).
    
3. **SqlSession:** Sử dụng phiên này để gửi câu lệnh SQL và nhận kết quả từ cơ sở dữ liệu.
    
4. **Mapper:** Mapper ánh xạ các phương thức Java với các câu lệnh SQL.
    
5. **Trả về kết quả:** Kết quả từ cơ sở dữ liệu được ánh xạ (mapping) thành đối tượng Java hoặc danh sách đối tượng.

---
## **2. Cấu hình cơ bản trong MyBatis**

### **2.1. `mybatis-config.xml`**

Đây là file cài đặt chính, giúp cấu hình MyBatis. Một số cấu hình quan trọng:

- **Database Configuration:** Thông tin kết nối cơ sở dữ liệu.
    
- **Settings:** Tùy chỉnh các cài đặt chung.
    
- **Mappers:** Nơi khai báo mapper XML hoặc package.
    

Ví dụ:
```xml
<configuration>
    <environments default="development">
        <environment id="development">
            <transactionManager type="JDBC" />
            <dataSource type="POOLED">
                <property name="driver" value="com.mysql.cj.jdbc.Driver" />
                <property name="url" value="jdbc:mysql://localhost:3306/mydb" />
                <property name="username" value="root" />
                <property name="password" value="password" />
            </dataSource>
        </environment>
    </environments>
    <mappers>
        <mapper resource="mapper/UserMapper.xml" />
    </mappers>
</configuration>
```

---
## **3. Thành phần cốt lõi**

### **3.1. SqlSession**

`SqlSession` là giao diện trung tâm để tương tác với cơ sở dữ liệu. Nó hỗ trợ các thao tác sau:

- **CRUD:** Thực thi các câu lệnh `INSERT`, `SELECT`, `UPDATE`, `DELETE`.
    
- **Transaction Management:** Quản lý giao dịch (commit, rollback).
    
- **Caching:** Quản lý bộ nhớ đệm (level 1 và level 2).
    

Ví dụ:
```java
try (SqlSession session = sqlSessionFactory.openSession()) {
    UserMapper mapper = session.getMapper(UserMapper.class);
    User user = mapper.findById(1);
    System.out.println(user.getUsername());
}
```

---
### **3.2. Mapper**

Mapper ánh xạ giữa phương thức Java và câu lệnh SQL trong cơ sở dữ liệu. Có hai cách định nghĩa mapper:

1. **Annotation-based Mapper:** SQL được định nghĩa trực tiếp trong các annotation.
    
2. **XML-based Mapper:** SQL được định nghĩa trong file XML.
    

**So sánh:**

| Đặc điểm         | Annotation-based     | XML-based                   |
| ---------------- | -------------------- | --------------------------- |
| **Tính dễ đọc**  | Phức tạp nếu SQL dài | Dễ đọc hơn cho SQL phức tạp |
| **Tái sử dụng**  | Khó tái sử dụng      | Dễ tái sử dụng              |
| **Tính bảo trì** | Khó bảo trì          | Dễ bảo trì hơn              |

---
## ** Quản lý bộ nhớ đệm (Caching)**

### **. Caching cấp 1 (Level 1 Cache)**

- Là bộ nhớ đệm mặc định, lưu trữ trong phạm vi **SqlSession**.
    
- Tự động được bật.
    

### **. Caching cấp 2 (Level 2 Cache)**

- Bộ nhớ đệm ở phạm vi **SqlSessionFactory** (toàn bộ ứng dụng).
    
- Phải được bật thủ công

---
## ** Quản lý giao dịch**

MyBatis hỗ trợ hai kiểu quản lý giao dịch:

- **JDBC Transaction:** Giao dịch được quản lý trực tiếp bởi JDBC.
    
- **Managed Transaction:** Giao dịch được quản lý bởi container (như Spring).
    

Ví dụ với giao dịch:

```java
try (SqlSession session = sqlSessionFactory.openSession()) {
    UserMapper mapper = session.getMapper(UserMapper.class);
    mapper.insertUser(new User("newUser", "newEmail"));
    session.commit(); // Commit giao dịch
} catch (Exception e) {
    session.rollback(); // Rollback nếu lỗi
}
```
## ** Lợi ích khi sử dụng MyBatis**

1. **Tối ưu hóa hiệu năng:** Bạn có toàn quyền kiểm soát SQL.
    
2. **Dễ dàng bảo trì:** SQL phức tạp có thể được tổ chức tốt trong file XML.
    
3. **Tích hợp với Spring:** Hỗ trợ tuyệt vời khi làm việc trong Spring Framework.

