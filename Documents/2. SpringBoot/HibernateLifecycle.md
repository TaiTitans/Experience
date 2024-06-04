
---


![](https://static.javatpoint.com/hibernatepages/images/hibernate-lifecycle.png)

### Các trạng thái vòng đời của Hibernate

Một thực thể trong Hibernate có thể tồn tại ở một trong bốn trạng thái chính trong suốt vòng đời của nó:

1. **Transient (Tạm thời)**:
    
    - Một đối tượng ở trạng thái transient nếu nó được khởi tạo bằng toán tử `new` nhưng không được liên kết với một phiên (session) Hibernate.
    - Đối tượng chưa được lưu trữ trong cơ sở dữ liệu và bất kỳ thay đổi nào trên nó sẽ không được phản ánh trong cơ sở dữ liệu.
    - Ví dụ: `Person person = new Person();`
2. **Persistent (Tồn tại)**:
    
    - Một đối tượng ở trạng thái persistent khi nó được liên kết với một phiên Hibernate và được quản lý bởi Hibernate.
    - Bất kỳ thay đổi nào trên đối tượng sẽ được đồng bộ hóa với cơ sở dữ liệu một cách tự động trong quá trình flush hoặc commit phiên làm việc.
    - Ví dụ: `session.save(person);`
3. **Detached (Tách rời)**:
    
    - Một đối tượng trở thành detached khi phiên Hibernate mà nó được liên kết bị đóng.
    - Đối tượng vẫn đại diện cho một hàng trong cơ sở dữ liệu, nhưng không còn được quản lý bởi phiên Hibernate.
    - Các thay đổi trên đối tượng detached sẽ không được tự động lưu vào cơ sở dữ liệu.
    - Ví dụ: `session.close();`
4. **Removed (Xóa bỏ)**:
    
    - Một đối tượng ở trạng thái removed nếu nó được đánh dấu để xóa khỏi cơ sở dữ liệu.
    - Thực thể sẽ bị xóa khi phiên làm việc được flush hoặc commit.
    - Ví dụ: `session.delete(person);`

### Các phương thức vòng đời Hibernate

Hiểu các phương thức liên quan đến việc chuyển đổi giữa các trạng thái này là điều cần thiết:

1. **save()**:
    
    - Chuyển từ Transient sang Persistent
    - Lưu một thực thể transient, làm cho nó trở thành persistent.
    - Ví dụ: `session.save(person);`
2. **persist()**:
    
    - Chuyển từ Transient sang Persistent
    - Tương tự như `save()`, nhưng không đảm bảo việc trả về khóa chính (identifier) được tạo.
    - Ví dụ: `session.persist(person);`
3. **get() / load()**:
    
    - Lấy một đối tượng từ cơ sở dữ liệu và trả về nó ở trạng thái persistent.
    - Ví dụ: `Person person = session.get(Person.class, id);`
4. **update() / merge()**:
    
    - Chuyển từ Detached sang Persistent
    - `update()`: Kết nối lại một đối tượng detached với một phiên mới, làm cho nó trở thành persistent.
    - `merge()`: Hợp nhất trạng thái của một đối tượng detached với ngữ cảnh persistent hiện tại.
    - Ví dụ: `session.update(person);`
5. **saveOrUpdate()**:
    
    - Chuyển từ Transient hoặc Detached sang Persistent
    - Lưu một thực thể transient hoặc cập nhật một thực thể detached.
    - Ví dụ: `session.saveOrUpdate(person);`
6. **delete()**:
    
    - Chuyển từ Persistent sang Removed
    - Đánh dấu đối tượng để xóa khỏi cơ sở dữ liệu.
    - Ví dụ: `session.delete(person);`

### Các chuyển đổi giữa các trạng thái vòng đời

Dưới đây là tóm tắt về cách một thực thể có thể chuyển đổi giữa các trạng thái sử dụng các phương thức khác nhau:

- **Transient sang Persistent**: `save()`, `persist()`, `saveOrUpdate()`
- **Persistent sang Detached**: `session.close()`, `session.evict()`, `session.clear()`
- **Detached sang Persistent**: `update()`, `merge()`, `saveOrUpdate()`
- **Persistent sang Removed**: `delete()`
- **Removed sang Transient**: Đối tượng bị xóa khỏi cơ sở dữ liệu và có thể bị thu gom rác (garbage collected) nếu không có tham chiếu nào tồn tại.

```Java
Session session = sessionFactory.openSession();
Transaction transaction = session.beginTransaction();

// Trạng thái Transient
Person person = new Person("John", "Doe");

// Trạng thái Persistent
session.save(person); // Bây giờ person đang ở trạng thái persistent

// Thực hiện một số thay đổi
person.setLastName("Smith");

// Những thay đổi sẽ được tự động đồng bộ với cơ sở dữ liệu khi commit hoặc flush
transaction.commit(); // Đồng bộ hóa các thay đổi với cơ sở dữ liệu

session.close(); // Bây giờ person ở trạng thái detached

// Mở lại phiên và gắn lại person
session = sessionFactory.openSession();
transaction = session.beginTransaction();

session.update(person); // person trở lại trạng thái persistent
person.setFirstName("Johnny");

transaction.commit(); // Các thay đổi được đồng bộ với cơ sở dữ liệu

session.close(); // person lại ở trạng thái detached

// Xóa person khỏi cơ sở dữ liệu
session = sessionFactory.openSession();
transaction = session.beginTransaction();

session.delete(person); // person được đánh dấu để xóa
transaction.commit(); // person bị xóa khỏi cơ sở dữ liệu

session.close();

```