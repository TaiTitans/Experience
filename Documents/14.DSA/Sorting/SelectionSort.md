
---
**Selection Sort** là một thuật toán sắp xếp đơn giản nhưng không hiệu quả cho các mảng lớn. Nó hoạt động bằng cách chia mảng thành hai phần: phần đã sắp xếp và phần chưa sắp xếp. Mỗi lần lặp, thuật toán tìm phần tử nhỏ nhất trong phần chưa sắp xếp và hoán đổi nó với phần tử đầu tiên của phần chưa sắp xếp, rồi mở rộng phần đã sắp xếp.
### **Cách hoạt động của Selection Sort**

1. Bắt đầu với phần đầu tiên của mảng.
2. Tìm phần tử nhỏ nhất trong phần chưa sắp xếp của mảng.
3. Hoán đổi phần tử nhỏ nhất với phần tử đầu tiên của phần chưa sắp xếp.
4. Lặp lại quá trình này cho phần còn lại của mảng.

### **Độ phức tạp**

- **Thời gian tốt nhất**: O(n²).
- **Thời gian trung bình**: O(n²).
- **Thời gian tồi tệ nhất**: O(n²).
- **Không gian**: O(1) (không cần bộ nhớ phụ ngoài mảng đầu vào).

### **Ưu điểm**

- Dễ cài đặt.
- Không phụ thuộc vào trạng thái ban đầu của mảng.

### **Nhược điểm**

- Hiệu suất kém với các mảng lớn.
- Tốn nhiều thời gian hơn so với các thuật toán như Merge Sort hay Quick Sort.

```java
public class SelectionSort {

    // Hàm thực hiện thuật toán Selection Sort
    public static void selectionSort(int[] arr) {
        int n = arr.length;

        // Di chuyển ranh giới của mảng chưa sắp xếp
        for (int i = 0; i < n - 1; i++) {
            // Tìm phần tử nhỏ nhất trong mảng chưa sắp xếp
            int minIdx = i;
            for (int j = i + 1; j < n; j++) {
                if (arr[j] < arr[minIdx]) {
                    minIdx = j;
                }
            }

            // Hoán đổi phần tử nhỏ nhất với phần tử đầu tiên của mảng chưa sắp xếp
            int temp = arr[minIdx];
            arr[minIdx] = arr[i];
            arr[i] = temp;
        }
    }

    // Hàm in mảng
    public static void printArray(int[] arr) {
        for (int i : arr) {
            System.out.print(i + " ");
        }
        System.out.println();
    }

    // Hàm chính để chạy chương trình
    public static void main(String[] args) {
        int[] arr = {64, 25, 12, 22, 11};
        System.out.println("Mảng ban đầu:");
        printArray(arr);

        selectionSort(arr);

        System.out.println("Mảng sau khi sắp xếp:");
        printArray(arr);
    }
}
```
### **Giải thích**

1. **Tìm phần tử nhỏ nhất**: Vòng lặp `for` đầu tiên chọn từng phần tử một làm phần tử hiện tại. Vòng lặp `for` lồng bên trong tìm phần tử nhỏ nhất trong phần còn lại của mảng.
2. **Hoán đổi phần tử**: Sau khi tìm thấy phần tử nhỏ nhất, thuật toán hoán đổi phần tử này với phần tử đầu tiên của phần chưa sắp xếp.
3. **Lặp lại**: Quá trình này tiếp tục cho đến khi toàn bộ mảng được sắp xếp.

### **Ví dụ hoạt động**

Với mảng đầu vào `[64, 25, 12, 22, 11]`:

1. Ở bước đầu tiên, phần tử nhỏ nhất `11` được tìm thấy và hoán đổi với `64`.
2. Tiếp theo, phần tử nhỏ nhất trong phần còn lại là `12` và được hoán đổi với `25`.
3. Quá trình tiếp tục cho đến khi mảng được sắp xếp thành `[11, 12, 22, 25, 64]`.