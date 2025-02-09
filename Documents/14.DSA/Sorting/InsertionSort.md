

---

The Insertion Sort algorithm uses one part of the array to hold the sorted values, and the other part of the array to hold values that are not sorted yet.


The algorithm takes one value at a time from the unsorted part of the array and puts it into the right place in the sorted part of the array, until the array is sorted.



**How it works:**

1. Take the first value from the unsorted part of the array.
2. Move the value into the correct place in the sorted part of the array.
3. Go through the unsorted part of the array again as many times as there are values.

![](https://he-s3.s3.amazonaws.com/media/uploads/46bfac9.png)

---
**Insertion Sort** là một thuật toán sắp xếp đơn giản và hiệu quả cho các mảng nhỏ hoặc danh sách gần như đã được sắp xếp. Thuật toán hoạt động bằng cách chia mảng thành hai phần: phần đã sắp xếp và phần chưa sắp xếp. Các phần tử từ phần chưa sắp xếp được chèn vào vị trí chính xác trong phần đã sắp xếp.
### **Cách hoạt động của Insertion Sort**

1. Bắt đầu từ phần tử thứ hai của mảng.
2. So sánh phần tử hiện tại với các phần tử trong phần đã sắp xếp.
3. Dịch chuyển các phần tử trong phần đã sắp xếp sang bên phải để tạo chỗ trống cho phần tử hiện tại.
4. Chèn phần tử hiện tại vào đúng vị trí trong phần đã sắp xếp.
5. Lặp lại quá trình cho đến khi toàn bộ mảng được sắp xếp.

### **Độ phức tạp**

- **Thời gian tốt nhất**: O(n) (khi mảng đã được sắp xếp).
- **Thời gian trung bình**: O(n²).
- **Thời gian tồi tệ nhất**: O(n²) (khi mảng được sắp xếp ngược).
- **Không gian**: O(1) (không cần bộ nhớ phụ ngoài mảng đầu vào).

### **Ưu điểm**

- Dễ cài đặt.
- Hiệu quả với các mảng nhỏ hoặc mảng gần như đã được sắp xếp.

### **Nhược điểm**

- Hiệu suất kém khi xử lý các mảng lớn và không có sự sắp xếp ban đầu.
```java
public class InsertionSort {

    // Hàm thực hiện thuật toán Insertion Sort
    public static void insertionSort(int[] arr) {
        int n = arr.length;
        for (int i = 1; i < n; i++) {
            int key = arr[i];  // Lấy phần tử hiện tại
            int j = i - 1;

            // Dịch chuyển các phần tử lớn hơn key sang bên phải
            while (j >= 0 && arr[j] > key) {
                arr[j + 1] = arr[j];
                j = j - 1;
            }

            // Chèn key vào vị trí chính xác
            arr[j + 1] = key;
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
        int[] arr = {12, 11, 13, 5, 6};
        System.out.println("Mảng ban đầu:");
        printArray(arr);

        insertionSort(arr);

        System.out.println("Mảng sau khi sắp xếp:");
        printArray(arr);
    }
}
```
