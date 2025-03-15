
---
**Mối quan hệ giữa equals và hashCode:**

- Nếu hai đối tượng gọi phương thức equals và trả về true, thì giá trị hashCode của chúng phải giống nhau.
- Nếu hai đối tượng có cùng giá trị hashCode, chúng không nhất thiết phải giống nhau (tức là equals không nhất thiết trả về true).

Phương thức hashCode chủ yếu được sử dụng để cải thiện hiệu suất so sánh đối tượng. Trước tiên, so sánh hashCode():

- Nếu giá trị hashCode khác nhau, không cần so sánh equals, điều này giảm đáng kể số lần so sánh bằng equals.
- Khi số lượng đối tượng cần so sánh lớn, việc này có thể cải thiện hiệu suất đáng kể.

### Giải thích bổ sung:

- **equals và hashCode trong Java**: Đây là hai phương thức quan trọng từ lớp Object. Chúng cần được ghi đè đồng bộ trong các lớp tự định nghĩa để đảm bảo tính nhất quán: nếu hai đối tượng được coi là bằng nhau (equals trả về true), thì hashCode của chúng phải trả về cùng một giá trị.
- **Ứng dụng thực tế**: Trong các cấu trúc dữ liệu như HashMap hoặc HashSet, hashCode được dùng để xác định vị trí lưu trữ, còn equals được dùng để kiểm tra chính xác sự trùng lặp khi có va chạm băm (hash collision).

---
### **hashCode là gì?**

- Phương thức hashCode() trong Java trả về một giá trị số nguyên (int) đại diện cho đối tượng. Nó được dùng để "băm" nội dung của đối tượng thành một giá trị ngắn gọn, thường để hỗ trợ các cấu trúc dữ liệu như HashMap hoặc HashSet.
- Giá trị hashCode không phải là duy nhất cho mỗi đối tượng. Vì kiểu int chỉ có 2³² giá trị (khoảng 4,3 tỷ), trong khi số lượng đối tượng có thể lớn hơn rất nhiều, nên không thể đảm bảo mọi đối tượng đều có hashCode khác nhau.