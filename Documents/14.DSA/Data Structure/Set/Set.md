
---
## Set là gì?

Set là một giao diện (interface) trong Java, nằm trong gói java.util, đại diện cho một tập hợp các phần tử **không trùng lặp**. Điều này có nghĩa là mỗi phần tử trong Set là duy nhất, và nếu bạn cố thêm một phần tử đã tồn tại, nó sẽ bị bỏ qua.

### Đặc điểm chính của Set

- **Không trùng lặp**: Chỉ chứa các phần tử duy nhất. Ví dụ: thêm "Apple" hai lần vào Set, chỉ một "Apple" được giữ lại.
- **Không có thứ tự cụ thể**: Thông thường, Set không đảm bảo thứ tự của các phần tử, trừ khi bạn dùng một triển khai đặc biệt như TreeSet (sắp xếp) hoặc LinkedHashSet (giữ thứ tự thêm vào).
- **Cho phép null (tùy triển khai)**: Một số Set cho phép chứa null, nhưng không phải tất cả.

---

## Các triển khai phổ biến của Set

Trong Java, Set có ba triển khai chính mà bạn cần nắm:

### 1. HashSet

- **Cách hoạt động**: Dựa trên **bảng băm (hash table)**. Mỗi phần tử được ánh xạ tới một vị trí trong bảng băm dựa trên giá trị hashCode của nó.
- **Hiệu suất**: Trung bình O(1) cho các thao tác thêm (add), xóa (remove), và kiểm tra (contains).
- **Thứ tự**: Không đảm bảo thứ tự của các phần tử.
- **Null**: Cho phép chứa một phần tử null.
- **Khi nào dùng**: Khi bạn cần một tập hợp không trùng lặp, không quan tâm đến thứ tự, và muốn hiệu suất cao.

**Ví dụ:**
```java
import java.util.HashSet;
import java.util.Set;

public class HashSetExample {
    public static void main(String[] args) {
        Set<String> set = new HashSet<>();
        set.add("Apple");
        set.add("Banana");
        set.add("Apple"); // Bị bỏ qua vì trùng lặp
        set.add(null);    // Cho phép null

        System.out.println(set); // Có thể in: [null, Apple, Banana] hoặc thứ tự khác
    }
}
```

### 2. TreeSet

- **Cách hoạt động**: Dựa trên **cây đỏ-đen (red-black tree)**, tự động sắp xếp các phần tử theo thứ tự tự nhiên (natural order) hoặc theo một Comparator bạn cung cấp.
- **Hiệu suất**: O(log n) cho các thao tác add, remove, và contains.
- **Thứ tự**: Luôn được sắp xếp.
- **Null**: Không cho phép null nếu không có Comparator rõ ràng (từ Java 7 trở đi).
- **Khi nào dùng**: Khi bạn cần một tập hợp không trùng lặp và muốn các phần tử được sắp xếp.

**Ví dụ:**
```java
import java.util.Set;
import java.util.TreeSet;

public class TreeSetExample {
    public static void main(String[] args) {
        Set<String> set = new TreeSet<>();
        set.add("Banana");
        set.add("Apple");
        set.add("Cherry");

        System.out.println(set); // In ra: [Apple, Banana, Cherry]
    }
}
```
### 3. LinkedHashSet

- **Cách hoạt động**: Kết hợp **bảng băm** và **danh sách liên kết**, giữ thứ tự thêm vào của các phần tử.
- **Hiệu suất**: Gần giống HashSet, trung bình O(1) cho các thao tác cơ bản.
- **Thứ tự**: Giữ nguyên thứ tự thêm vào.
- **Null**: Cho phép chứa một phần tử null.
- **Khi nào dùng**: Khi bạn cần một tập hợp không trùng lặp và muốn giữ thứ tự chèn.

**Ví dụ:**
```java
import java.util.LinkedHashSet;
import java.util.Set;

public class LinkedHashSetExample {
    public static void main(String[] args) {
        Set<String> set = new LinkedHashSet<>();
        set.add("Apple");
        set.add("Banana");
        set.add("Cherry");

        System.out.println(set); // In ra: [Apple, Banana, Cherry]
    }
}
```

## Các phương thức quan trọng của Set

Dưới đây là các phương thức phổ biến bạn sẽ dùng với Set:

- **add(E e)**: Thêm một phần tử. Trả về true nếu thêm thành công (không trùng lặp).
- **remove(Object o)**: Xóa một phần tử. Trả về true nếu xóa thành công.
- **contains(Object o)**: Kiểm tra xem Set có chứa phần tử không.
- **size()**: Trả về số lượng phần tử trong Set.
- **isEmpty()**: Kiểm tra Set có rỗng không.
- **clear()**: Xóa toàn bộ phần tử trong Set.

---

## So sánh các triển khai của Set

|Triển khai|Thứ tự|Hiệu suất|Cho phép null|
|---|---|---|---|
|HashSet|Không đảm bảo|O(1)|Có|
|TreeSet|Sắp xếp|O(log n)|Không (thường)|
|LinkedHashSet|Thứ tự thêm vào|O(1)|Có|

---

## Lưu ý quan trọng khi sử dụng Set

### 1. HashSet và equals/hashCode

Khi dùng HashSet hoặc LinkedHashSet với các đối tượng tùy chỉnh (custom objects), bạn **phải ghi đè** hai phương thức:

- **equals(Object o)**: Xác định hai đối tượng có "bằng nhau" không.
- **hashCode()**: Trả về giá trị hash để HashSet định vị phần tử.

Nếu không làm điều này, HashSet sẽ không nhận diện được các phần tử trùng lặp đúng cách.

### 2. TreeSet và Comparable/Comparator

Với TreeSet, các phần tử phải:

- Triển khai giao diện Comparable (để xác định thứ tự tự nhiên), hoặc
- Cung cấp một Comparator khi tạo TreeSet.

Ví dụ: Nếu bạn thêm một lớp không triển khai Comparable vào TreeSet mà không có Comparator, chương trình sẽ抛出 (ném ra) lỗi ClassCastException.

### 3. Không truy cập theo chỉ số

Set không hỗ trợ truy cập phần tử theo chỉ số như List. Nếu cần, bạn có thể chuyển Set thành List bằng cách:
```java
List<String> list = new ArrayList<>(set);
```


---
## 1. HashSet: Cách lưu trữ và tối ưu

### Cách lưu trữ trong bộ nhớ

- **HashSet dựa trên bảng băm (hash table)**:
    - Một mảng (array) bên trong được sử dụng để lưu trữ các phần tử.
    - Mỗi phần tử được ánh xạ đến một ô (bucket) trong mảng dựa trên giá trị hashCode của nó.
    - Nếu có **va chạm (collision)** (tức là nhiều phần tử có cùng hashCode), các phần tử được lưu trong một **danh sách liên kết (linked list)** hoặc **cây (tree)** (từ Java 8, nếu số lượng va chạm lớn, danh sách liên kết sẽ chuyển thành cây đỏ-đen để cải thiện hiệu suất).
- **Cấu trúc bên trong**:
    - Một mảng các bucket (ví dụ: Object[] table).
    - Mỗi bucket có thể chứa một nút đơn (nếu không va chạm) hoặc một cấu trúc dữ liệu phụ (linked list hoặc tree).
    - Kích thước mảng được điều chỉnh động (resize) khi số phần tử vượt quá ngưỡng (load factor).
- **Ví dụ minh họa**:
    - Thêm "Apple" với hashCode = 123 → ánh xạ vào bucket 123 % mảng.length.
    - Thêm "Banana" với hashCode = 124 → ánh xạ vào bucket khác.
    - Nếu "Apple" và "Cat" cùng ánh xạ vào bucket 123, chúng được lưu trong danh sách liên kết hoặc cây.

### Độ phức tạp

- Thêm (add), xóa (remove), kiểm tra (contains): Trung bình **O(1)**, nhưng có thể tệ hơn nếu có nhiều va chạm (O(n) trong trường hợp xấu nhất).

### Cách tối ưu hiệu suất

1. **Tùy chỉnh kích thước ban đầu (initial capacity)**:
    - Khi tạo HashSet, bạn có thể chỉ định kích thước ban đầu để tránh việc mở rộng mảng nhiều lần (resize rất tốn kém).
    - Ví dụ: Nếu biết trước sẽ lưu 1000 phần tử, dùng new HashSet<>(1000) thay vì mặc định (16).
    - Công thức: initialCapacity = (số phần tử dự kiến / load factor) + 1. Load factor mặc định là 0.75.
2. **Điều chỉnh load factor**:
    - Load factor (tỷ lệ lấp đầy) quyết định khi nào mảng sẽ resize. Mặc định là 0.75 (75%).
    - Giảm load factor (ví dụ: 0.5) → ít va chạm hơn, nhưng tốn bộ nhớ hơn.
    - Tăng load factor (ví dụ: 0.9) → tiết kiệm bộ nhớ, nhưng tăng khả năng va chạm.
    - Ví dụ: new HashSet<>(1000, 0.5f).
3. **Tối ưu hashCode**:
    - Đảm bảo phương thức hashCode() của đối tượng phân bố đều giá trị hash để giảm va chạm.
    - Ví dụ: Nếu hashCode() trả về giá trị cố định (như luôn là 1), tất cả phần tử sẽ rơi vào một bucket, làm hiệu suất giảm xuống O(n).
4. **Sử dụng đối tượng immutable**:
    - Nếu đối tượng thay đổi trạng thái sau khi thêm vào HashSet (ví dụ: thay đổi thuộc tính ảnh hưởng đến hashCode), nó có thể không được tìm thấy nữa. Hãy dùng các lớp immutable (như String).

---

## 2. TreeSet: Cách lưu trữ và tối ưu

### Cách lưu trữ trong bộ nhớ

- **TreeSet dựa trên cây đỏ-đen (red-black tree)**:
    - Đây là một cây nhị phân tự cân bằng, đảm bảo chiều cao cây luôn là O(log n).
    - Mỗi nút trong cây chứa một phần tử, và các nút được sắp xếp theo thứ tự (dựa trên Comparable hoặc Comparator).
    - Mỗi nút có màu (đỏ hoặc đen) để duy trì tính cân bằng.
- **Cấu trúc bên trong**:
    - Một tập hợp các nút (node) liên kết với nhau.
    - Mỗi nút có: giá trị, con trái, con phải, nút cha, và màu.
- **Ví dụ minh họa**:
    - Thêm 5 → gốc cây.
    - Thêm 2 → con trái của 5.
    - Thêm 8 → con phải của 5.
    - Cây tự cân bằng để đảm bảo không có nhánh nào quá dài.

### Độ phức tạp

- Thêm, xóa, kiểm tra: **O(log n)** vì chiều cao cây luôn được giữ cân bằng.

### Cách tối ưu hiệu suất

1. **Tối ưu Comparator/Comparable**:
    - Nếu dùng Comparator hoặc Comparable, hãy đảm bảo phương thức so sánh (compare hoặc compareTo) chạy nhanh và chính xác.
    - Tránh các phép tính phức tạp trong so sánh, vì nó được gọi nhiều lần khi thêm/xóa phần tử.
2. **Giảm số lần cân bằng**:
    - Thêm phần tử theo thứ tự ngẫu nhiên thay vì thứ tự tăng dần/giảm dần liên tục. Thêm liên tục theo thứ tự có thể làm cây tạm thời mất cân bằng nhiều hơn, dù cuối cùng nó vẫn tự điều chỉnh.
3. **Sử dụng TreeSet khi thực sự cần sắp xếp**:
    - Nếu không cần sắp xếp, hãy dùng HashSet để đạt hiệu suất tốt hơn (O(1) thay vì O(log n)).

---

## 3. LinkedHashSet: Cách lưu trữ và tối ưu

### Cách lưu trữ trong bộ nhớ

- **LinkedHashSet là sự kết hợp của bảng băm và danh sách liên kết đôi (doubly-linked list)**:
    - Giống HashSet, nó dùng bảng băm để lưu trữ và kiểm tra trùng lặp.
    - Thêm vào đó, nó duy trì một danh sách liên kết đôi để ghi nhớ thứ tự thêm vào của các phần tử.
- **Cấu trúc bên trong**:
    - Một mảng bảng băm (giống HashSet).
    - Mỗi phần tử còn được liên kết với phần tử trước và sau trong danh sách liên kết.
- **Ví dụ minh họa**:
    - Thêm "Apple" → vào bucket X, thêm vào danh sách liên kết.
    - Thêm "Banana" → vào bucket Y, liên kết sau "Apple".
    - Duy trì cả thứ tự chèn và tính không trùng lặp.

### Độ phức tạp

- Thêm, xóa, kiểm tra: **O(1)** (giống HashSet), nhưng tốn thêm chút bộ nhớ và thời gian để duy trì danh sách liên kết.

### Cách tối ưu hiệu suất

1. **Tối ưu giống HashSet**:
    - Chỉ định kích thước ban đầu và load factor phù hợp.
    - Tối ưu hashCode để giảm va chạm.
2. **Tránh lặp không cần thiết**:
    - Vì LinkedHashSet giữ thứ tự, việc lặp qua các phần tử (iteration) nhanh hơn HashSet (O(n) nhưng không bị xáo trộn). Hãy tận dụng điều này nếu cần duyệt tuần tự.
3. **Chỉ dùng khi cần thứ tự**:
    - Nếu không cần thứ tự chèn, HashSet sẽ nhẹ hơn vì không phải duy trì danh sách liên kết.

---

## So sánh bộ nhớ và hiệu suất

|Triển khai|Cách lưu trữ|Bộ nhớ sử dụng|Hiệu suất trung bình|
|---|---|---|---|
|HashSet|Bảng băm|Nhẹ nhất|O(1)|
|TreeSet|Cây đỏ-đen|Nhiều hơn (con trỏ)|O(log n)|
|LinkedHashSet|Bảng băm + linked list|Nhiều hơn HashSet|O(1)|

---

## Chiến lược tối ưu chung cho Set

1. **Chọn đúng triển khai**:
    - Không cần thứ tự → HashSet.
    - Cần sắp xếp → TreeSet.
    - Cần giữ thứ tự chèn → LinkedHashSet.
2. **Dự đoán kích thước**:
    - Với HashSet và LinkedHashSet, cung cấp kích thước ban đầu để tránh resize.
3. **Tối ưu đối tượng**:
    - Với HashSet/LinkedHashSet: Tối ưu hashCode và equals.
    - Với TreeSet: Tối ưu compareTo hoặc Comparator.
4. **Tránh thay đổi đối tượng sau khi thêm**:
    - Thay đổi trạng thái ảnh hưởng đến hashCode hoặc thứ tự có thể làm Set hoạt động sai.
5. **Sử dụng bộ nhớ cache nếu cần**:
    - Nếu truy vấn nhiều lần (như contains), hãy cân nhắc giữ Set trong bộ nhớ đệm để tránh tạo lại.