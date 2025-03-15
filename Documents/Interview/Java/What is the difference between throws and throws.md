
---
- **throw:** Được sử dụng để ném một đối tượng ngoại lệ cụ thể.
- **throws:** Được sử dụng trong chữ ký phương thức để khai báo các ngoại lệ mà phương thức có thể ném ra.
    - Các phương thức của lớp con chỉ có thể ném một phạm vi ngoại lệ nhỏ hơn (hoặc không ném ngoại lệ nào) so với phương thức của lớp cha.
### Giải thích bổ sung:

1. **throw:**
    - Là câu lệnh thực thi trong mã để chủ động ném ngoại lệ.
    - Ví dụ: throw new IOException("Lỗi đọc tệp"); – ném một ngoại lệ cụ thể khi có lỗi xảy ra.
2. **throws:**
    - Là khai báo trong định nghĩa phương thức, cho biết phương thức có thể ném những ngoại lệ nào để người gọi phương thức xử lý.
```java
public void myMethod() throws IOException {
    throw new IOException("Lỗi");
}
```

**Quy tắc ghi đè (override)**: Khi lớp con ghi đè phương thức của lớp cha, phạm vi ngoại lệ trong throws của lớp con phải hẹp hơn hoặc bằng lớp cha (hoặc không khai báo throws). Ví dụ:

- Lớp cha: void method() throws Exception
- Lớp con: void method() throws IOException (được phép vì IOException là con của Exception), hoặc void method() (không ném ngoại lệ nào).

