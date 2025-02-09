
---
### 1. Hiểu Cách Hoạt Động Của Merge Sort
**Merge Sort** là một thuật toán sắp xếp chia để trị (divide and conquer).
**Merge Sort** chia mảng thành hai nửa, sắp xếp từng nửa bằng cách gọi đệ quy và sau đó gộp hai nửa đã sắp xếp lại với nhau.

### 2. Cấu Trúc Merge Sort

- **Chia nhỏ (Divide):** Chia mảng thành hai nửa nhỏ hơn.
- **Sắp xếp (Conquer):** Sắp xếp từng nửa bằng đệ quy.
- **Gộp (Combine):** Gộp hai nửa đã sắp xếp lại với nhau.

```java
public class MergeSort {

    public static void mergeSort(int[] arr) {
        if (arr.length < 2) {
            return;
        }
        int mid = arr.length / 2;
        int[] left = new int[mid];
        int[] right = new int[arr.length - mid];

        // Copy data to temporary arrays
        for (int i = 0; i < mid; i++) {
            left[i] = arr[i];
        }
        for (int i = mid; i < arr.length; i++) {
            right[i - mid] = arr[i];
        }

        mergeSort(left);
        mergeSort(right);

        merge(arr, left, right);
    }

    private static void merge(int[] arr, int[] left, int[] right) {
        int i = 0, j = 0, k = 0;
        while (i < left.length && j < right.length) {
            if (left[i] <= right[j]) {
                arr[k++] = left[i++];
            } else {
                arr[k++] = right[j++];
            }
        }
        while (i < left.length) {
            arr[k++] = left[i++];
        }
        while (j < right.length) {
            arr[k++] = right[j++];
        }
    }

    public static void main(String[] args) {
        int[] arr = {38, 27, 43, 3, 9, 82, 10};
        System.out.println("Original array:");
        printArray(arr);

        mergeSort(arr);

        System.out.println("Sorted array:");
        printArray(arr);
    }

    private static void printArray(int[] arr) {
        for (int num : arr) {
            System.out.print(num + " ");
        }
        System.out.println();
    }
}
```

### Phân Tích Mã Java

1. **mergeSort()**:
    
    - Chia mảng thành hai phần: `left` và `right`.
    - Đệ quy gọi `mergeSort` trên cả hai phần.
    - Gọi `merge()` để gộp hai phần đã sắp xếp lại.
2. **merge()**:
    
    - Gộp hai mảng con `left` và `right` vào mảng `arr` chính.
    - Sử dụng hai con trỏ `i` và `j` để duyệt qua `left` và `right`, so sánh từng phần tử và chèn phần tử nhỏ hơn vào `arr`.
3. **printArray()**:
    
    - In mảng ra màn hình để kiểm tra kết quả.


### Phân Tích Độ Phức Tạp

- **Thời gian:** O(n log n).
- **Không gian:** O(n), do cần không gian phụ để lưu các mảng con.

---
Merge Sort nên được sử dụng trong các trường hợp sau:

### 1. **Khi cần độ ổn định (Stable Sorting)**

- Merge Sort là một thuật toán sắp xếp ổn định, nghĩa là nó giữ nguyên thứ tự của các phần tử có giá trị bằng nhau trong mảng.
- Nếu bạn cần sắp xếp dữ liệu mà việc bảo toàn thứ tự gốc của các phần tử bằng nhau là quan trọng (ví dụ: sắp xếp theo họ tên sau khi đã sắp xếp theo tuổi), Merge Sort là một lựa chọn tốt.

### 2. **Khi làm việc với dữ liệu lớn**

- Merge Sort có độ phức tạp thời gian là O(n log n), ổn định và hiệu quả cho việc sắp xếp dữ liệu lớn, đặc biệt là khi dữ liệu không thể được giữ trong bộ nhớ và phải xử lý trên ổ đĩa (External Sorting).
- Trong trường hợp này, Merge Sort có thể được sử dụng để sắp xếp từng phần nhỏ trong bộ nhớ trước khi gộp lại.

### 3. **Khi cần hiệu quả về thời gian hơn không gian**

- Merge Sort có thời gian chạy ổn định O(n log n) ngay cả trong trường hợp xấu nhất, trong khi các thuật toán khác như Quick Sort có thể rơi vào O(n²) nếu chọn pivot không tốt.
- Điều này làm cho Merge Sort phù hợp với các ứng dụng cần đảm bảo thời gian chạy ổn định.

### 4. **Khi cần sắp xếp dữ liệu được lưu trữ trên đĩa**

- Merge Sort rất hữu ích khi sắp xếp các tệp dữ liệu lớn hơn bộ nhớ (External Sorting), vì nó hoạt động bằng cách chia nhỏ dữ liệu, xử lý từng phần trong bộ nhớ và sau đó gộp lại.

### 5. **Khi cần một thuật toán đơn giản và dễ triển khai**

- Mặc dù không phải là thuật toán tối ưu về không gian, nhưng Merge Sort có một cấu trúc đơn giản, dễ hiểu và dễ triển khai, đặc biệt trong các ngôn ngữ hỗ trợ đệ quy tốt.

### 6. **Khi dữ liệu không thể sắp xếp tại chỗ (Not In-Place)**

- Nếu bạn có các hạn chế về việc sắp xếp tại chỗ và có đủ bộ nhớ để chứa các bản sao của mảng, Merge Sort là một lựa chọn phù hợp do nó không thực hiện sắp xếp tại chỗ.

###  **Khi không nên sử dụng Merge Sort**:

- Khi bạn cần một thuật toán sắp xếp tại chỗ để tiết kiệm bộ nhớ.
- Khi dữ liệu nhỏ và bạn cần một thuật toán sắp xếp nhanh hơn về mặt thực tế như Quick Sort hoặc Insertion Sort.
- Khi không muốn sử dụng thêm không gian phụ lớn (O(n)).