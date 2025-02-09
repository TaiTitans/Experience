
---

`# Iterative and Recursive Implementation`


**Binary Search*** ***Algorithm*** is a [searching algorithm](https://www.geeksforgeeks.org/searching-algorithms/) used in a sorted array by ***repeatedly dividing the search interval in half***. The idea of binary search is to use the information that the array is sorted and reduce the time complexity to O(log N).

![](https://media.geeksforgeeks.org/wp-content/uploads/20240506155201/binnary-search-.webp)

---
**1. Khái niệm:**

- Binary Search là một thuật toán tìm kiếm nhanh, được sử dụng trên các mảng đã được sắp xếp (ascending hoặc descending).
- Thuật toán chia nhỏ mảng tìm kiếm thành hai phần và loại bỏ một nửa phần tử mỗi lần lặp, do đó có hiệu quả cao hơn so với Linear Search (tìm kiếm tuyến tính).

**2. Nguyên lý hoạt động:**

- Tìm phần tử giữa (mid) của mảng.
- So sánh giá trị của phần tử `mid` với phần tử cần tìm (key).
    - Nếu `key` bằng `mid`, trả về vị trí của `mid`.
    - Nếu `key` nhỏ hơn `mid`, tìm kiếm trong nửa bên trái của mảng.
    - Nếu `key` lớn hơn `mid`, tìm kiếm trong nửa bên phải của mảng.
- Lặp lại quá trình cho đến khi tìm thấy phần tử hoặc mảng rỗng (kết thúc tìm kiếm).

**3. Điều kiện tiên quyết:**

- Mảng phải được sắp xếp trước khi áp dụng Binary Search.

**4. Độ phức tạp:**

- **Thời gian (Time Complexity):**
    - **Best case:** O(1) (tìm thấy ngay tại phần tử giữa).
    - **Worst case:** O(log n) (chia đôi mảng mỗi lần).
- **Không gian (Space Complexity):**
    - O(1) cho phiên bản lặp (Iterative).
    - O(log n) cho phiên bản đệ quy (Recursive) do sử dụng ngăn xếp đệ quy.

**5. Các bước thực hiện:**

1. Khởi tạo hai biến chỉ mục `low` (đầu mảng) và `high` (cuối mảng).
2. Trong khi `low` nhỏ hơn hoặc bằng `high`:
    - Tính toán `mid` (giữa mảng): `mid = low + (high - low) / 2`.
    - So sánh phần tử ở `mid` với `key`.
    - Điều chỉnh `low` hoặc `high` dựa trên kết quả so sánh.
3. Nếu `key` không được tìm thấy sau khi kết thúc vòng lặp, trả về -1 hoặc thông báo "không tìm thấy".

```java
public static int binarySearch(int[] arr, int target) {
    int left = 0, right = arr.length - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (arr[mid] == target) {
            return mid;
        }
        if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return -1;
}
```

**6. Phiên bản sử dụng:**

- **Iterative**: Sử dụng vòng lặp `while`.
- **Recursive**: Gọi lại hàm với mảng con.

