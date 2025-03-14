
---
**Truyền giá trị (Value Transfer)** áp dụng cho các biến cơ bản, trong đó một bản sao của biến được truyền đi. Việc thay đổi bản sao này không ảnh hưởng đến biến gốc.

**Truyền tham chiếu (Reference Transfer)** thường áp dụng cho các biến kiểu đối tượng, trong đó một bản sao của địa chỉ của đối tượng được truyền đi, chứ không phải bản thân đối tượng gốc. Cả hai cùng trỏ đến một vùng bộ nhớ. Vì vậy, việc thao tác trên đối tượng được tham chiếu sẽ đồng thời thay đổi đối tượng gốc.

**Trong Java không có truyền tham chiếu, chỉ có truyền giá trị.** Điều này có nghĩa là không tồn tại trường hợp biến A trỏ đến biến B và biến B trỏ đến một đối tượng.