
---
Heap là một dạng **cấu trúc dữ liệu cây** được sử dụng chủ yếu trong **hàng đợi ưu tiên** (Priority Queue). Có hai loại heap chính:

- **Max Heap**: Nút cha luôn lớn hơn hoặc bằng các nút con.
- **Min Heap**: Nút cha luôn nhỏ hơn hoặc bằng các nút con.

## **1. Đặc điểm chung của Heap**

- Heap thường được **cài đặt bằng một mảng** thay vì một cấu trúc cây liên kết.
- Nó là **cây nhị phân hoàn chỉnh** (Complete Binary Tree), nghĩa là:
    - Tất cả các mức (trừ mức cuối cùng) đều đầy đủ.
    - Các nút ở mức cuối cùng nằm **từ trái sang phải**, không được có khoảng trống.
- Heap thường được sử dụng trong:
    - **Hàng đợi ưu tiên (Priority Queue)**
    - **Thuật toán sắp xếp Heap Sort**
    - **Tìm phần tử lớn nhất hoặc nhỏ nhất nhanh chóng**

## **2. Max Heap**

- Ở **Max Heap**, phần tử lớn nhất nằm ở **gốc (root)**.
- Mỗi nút cha luôn có giá trị **lớn hơn hoặc bằng** các nút con.

**Ví dụ Max Heap:**
```
       50
      /  \
    30    40
   /  \   /  \
 10   20 15  35

```

- **50** là phần tử lớn nhất ở gốc.
- Mỗi nút cha lớn hơn hoặc bằng các nút con của nó.

### **Cài đặt Max Heap bằng Mảng**

Nếu ta lưu Max Heap trên mảng:
```
Index:  0   1   2   3   4   5   6
Array: [50, 30, 40, 10, 20, 15, 35]
```

Công thức tính vị trí:

- **Cha của node `i`**: `(i - 1) / 2`
- **Con trái của node `i`**: `2 * i + 1`
- **Con phải của node `i`**: `2 * i + 2`

## **3. Min Heap**

- Ở **Min Heap**, phần tử nhỏ nhất nằm ở **gốc (root)**.
- Mỗi nút cha luôn có giá trị **nhỏ hơn hoặc bằng** các nút con.

**Ví dụ Min Heap:**

```
       10
      /   \
    20     15
   /  \   /   \
 30   50 40   35

```

- **10** là phần tử nhỏ nhất ở gốc.
- Mỗi nút cha nhỏ hơn hoặc bằng các nút con của nó.

Cài đặt Min Heap bằng Mảng
```
Index:  0   1   2   3   4   5   6
Array: [10, 20, 15, 30, 50, 40, 35]
```

Sử dụng cùng công thức vị trí như Max Heap.

## **4. Các thao tác chính trong Heap**

### **a) Thêm phần tử (Insertion)**

- Thêm vào vị trí cuối cùng của mảng (duy trì cấu trúc cây hoàn chỉnh).
- **"Heapify Up"**: So sánh với cha, nếu vi phạm điều kiện Heap, hoán đổi.
- Độ phức tạp: **O(log N)**.

### **b) Xóa phần tử lớn nhất (Max Heap) hoặc nhỏ nhất (Min Heap)**

- Thay thế gốc bằng phần tử cuối cùng.
- **"Heapify Down"**: So sánh với con lớn hơn (Max Heap) hoặc nhỏ hơn (Min Heap), nếu vi phạm điều kiện Heap, hoán đổi.
- Độ phức tạp: **O(log N)**.

5. Cài đặt Max Heap trong Java
```java
import java.util.PriorityQueue;
import java.util.Collections;

public class MaxHeapExample {
    public static void main(String[] args) {
        // Khởi tạo Max Heap (hàng đợi ưu tiên với thứ tự đảo ngược)
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Collections.reverseOrder());

        // Thêm phần tử vào Max Heap
        maxHeap.add(50);
        maxHeap.add(30);
        maxHeap.add(40);
        maxHeap.add(10);
        maxHeap.add(20);
        maxHeap.add(15);
        maxHeap.add(35);

        // In ra phần tử lớn nhất (root)
        System.out.println("Max element: " + maxHeap.peek()); // 50

        // Xóa phần tử lớn nhất
        maxHeap.poll();
        System.out.println("After removing max, new max: " + maxHeap.peek()); // 40
    }
}
```

6. Cài đặt Min Heap trong Java
```java
import java.util.PriorityQueue;

public class MinHeapExample {
    public static void main(String[] args) {
        // Khởi tạo Min Heap (mặc định)
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();

        // Thêm phần tử vào Min Heap
        minHeap.add(50);
        minHeap.add(30);
        minHeap.add(40);
        minHeap.add(10);
        minHeap.add(20);
        minHeap.add(15);
        minHeap.add(35);

        // In ra phần tử nhỏ nhất (root)
        System.out.println("Min element: " + minHeap.peek()); // 10

        // Xóa phần tử nhỏ nhất
        minHeap.poll();
        System.out.println("After removing min, new min: " + minHeap.peek()); // 15
    }
}
```

## **7. Ứng dụng của Heap**

- **Dijkstra’s Algorithm**: Tìm đường đi ngắn nhất trong đồ thị.
- **Heap Sort**: Thuật toán sắp xếp hiệu quả với độ phức tạp O(N log N).
- **Tìm K phần tử lớn nhất hoặc nhỏ nhất** trong danh sách.
- **Hệ thống ưu tiên xử lý** (ví dụ: hệ thống đặt lịch CPU, hệ thống đường dây nóng).

---

## **8. So sánh Min Heap & Max Heap**

|Đặc điểm|Max Heap|Min Heap|
|---|---|---|
|Root|Lớn nhất|Nhỏ nhất|
|Được dùng khi|Cần lấy giá trị lớn nhất|Cần lấy giá trị nhỏ nhất|
|Ví dụ|Sắp xếp giảm dần, ưu tiên công việc quan trọng trước|Tìm đường đi ngắn nhất, xử lý công việc ít ưu tiên trước|

---

## **Kết luận**

- **Heap** là một **cây nhị phân hoàn chỉnh**, giúp truy xuất phần tử lớn nhất hoặc nhỏ nhất nhanh chóng (**O(1)**).
- **Max Heap** ưu tiên giá trị lớn nhất, **Min Heap** ưu tiên giá trị nhỏ nhất.
- **Java có `PriorityQueue` để cài đặt Min Heap, và `Collections.reverseOrder()` để biến nó thành Max Heap**.