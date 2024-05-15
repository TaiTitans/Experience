
---

### Tính chất của interface

- Interface không cung cấp việc kế thừa như Class hay Abstract Class mà nó chỉ khai báo phương thức/sự kiện để các lớp khác thực hiện nó.

- Nó không được khởi tạo nhưng nó có thể được tham chiếu bởi đối tượng của lớp thực hiện nó. ví dụ: IUserRepository user = new UserRepository();

- Không có interface lồng nhau (nested interface)

- Không có constructor, destructor, constants, static và variable.

- Một interface có thể kế thừa từ một hoặc nhiều interface khác. public interface IList : ICollection, IEnumerable, IEnumerable

- Một interface có thể mở rộng một interface khác.

- Một class có thể implement một hoặc nhiều interfaces.

- Mọi phương thức, property đều mặc định là public.


---

