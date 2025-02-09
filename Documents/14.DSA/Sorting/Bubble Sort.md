
---
### 1. Hiểu Cách Hoạt Động Của Bubble Sort

**Bubble Sort** là một thuật toán sắp xếp đơn giản, hoạt động bằng cách lặp qua mảng và hoán đổi các phần tử kề nhau nếu chúng không theo thứ tự mong muốn.

- **Ý tưởng chính:** So sánh từng cặp phần tử liền kề và hoán đổi chúng nếu chúng ở sai thứ tự.
- **Quá trình này tiếp tục:** Cho đến khi không còn cặp phần tử nào cần hoán đổi.

### 2. Cách Hoạt Động

- **Pass đầu tiên:** So sánh từng cặp phần tử liền kề, hoán đổi nếu cần. Sau pass này, phần tử lớn nhất "nổi" lên cuối mảng.
- **Pass tiếp theo:** Lặp lại quá trình nhưng không cần kiểm tra phần tử cuối cùng (vì nó đã ở đúng vị trí).
- **Lặp lại:** Quá trình cho đến khi mảng được sắp xếp hoàn toàn.

```java
public class BubbleSort {

    public static void bubbleSort(int[] arr) {
        int n = arr.length;
        boolean swapped;
        for (int i = 0; i < n - 1; i++) {
            swapped = false;
            for (int j = 0; j < n - 1 - i; j++) {
                if (arr[j] > arr[j + 1]) {
                    // Hoán đổi arr[j] và arr[j + 1]
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                    swapped = true;
                }
            }
            // Nếu không có phần tử nào được hoán đổi, mảng đã được sắp xếp
            if (!swapped) {
                break;
            }
        }
    }

    public static void main(String[] args) {
        int[] arr = {64, 34, 25, 12, 22, 11, 90};
        System.out.println("Original array:");
        printArray(arr);

        bubbleSort(arr);

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

1. **bubbleSort()**:
    
    - Sử dụng vòng lặp lồng nhau, vòng lặp bên ngoài kiểm soát số lần pass, vòng lặp bên trong thực hiện so sánh và hoán đổi các phần tử liền kề.
    - Biến `swapped` được sử dụng để theo dõi liệu có sự hoán đổi nào xảy ra hay không, giúp tối ưu hóa bằng cách thoát khỏi vòng lặp sớm nếu mảng đã được sắp xếp.
2. **main()**:
    
    - Tạo mảng và gọi hàm `bubbleSort` để sắp xếp.
    - Sử dụng `printArray()` để hiển thị mảng trước và sau khi sắp xếp.
3. **printArray()**:
    
    - In các phần tử của mảng ra màn hình.

### Phân Tích Độ Phức Tạp

- **Thời gian trung bình và trường hợp xấu nhất:** O(n²).
- **Thời gian trường hợp tốt nhất:** O(n), khi mảng đã được sắp xếp.
- **Không gian:** O(1), vì Bubble Sort là thuật toán sắp xếp tại chỗ.

---
### Khi Nên Sử Dụng Bubble Sort:

1. **Dữ liệu rất nhỏ hoặc đã gần sắp xếp:**
    
    - Bubble Sort hoạt động tốt nhất khi làm việc với các bộ dữ liệu nhỏ hoặc khi mảng đã gần sắp xếp.
    - Trường hợp tốt nhất của Bubble Sort là O(n), nếu mảng đã được sắp xếp hoặc chỉ cần một vài lần hoán đổi.
2. **Yêu cầu về thuật toán đơn giản và dễ hiểu:**
    
    - Bubble Sort dễ cài đặt và hiểu, rất hữu ích cho mục đích giáo dục, đặc biệt là khi dạy hoặc học về các thuật toán cơ bản.
    - Khi bạn chỉ cần một thuật toán đơn giản để minh họa cách hoạt động của sắp xếp, Bubble Sort là một lựa chọn tốt.
3. **Cần sắp xếp tại chỗ (In-Place Sorting):**
    
    - Bubble Sort không cần không gian phụ lớn (O(1)), vì nó chỉ hoán đổi các phần tử trong mảng, phù hợp khi bộ nhớ hạn chế.
4. **Khi ổn định là quan trọng:**
    
    - Bubble Sort là một thuật toán sắp xếp ổn định, nghĩa là các phần tử bằng nhau sẽ giữ nguyên thứ tự ban đầu của chúng trong mảng.

### Khi Không Nên Sử Dụng Bubble Sort:

1. **Khi làm việc với dữ liệu lớn:**
    
    - Bubble Sort có độ phức tạp thời gian là O(n²), không hiệu quả cho dữ liệu lớn.
    - Các thuật toán khác như Quick Sort hoặc Merge Sort có độ phức tạp trung bình là O(n log n) sẽ tốt hơn nhiều.
2. **Khi hiệu suất quan trọng:**
    
    - Nếu hiệu suất là một yếu tố quan trọng, bạn nên chọn các thuật toán tối ưu hơn như Quick Sort, Merge Sort hoặc Heap Sort.
3. **Khi cần tối ưu hóa thời gian thực thi:**
    
    - Các thuật toán như Insertion Sort hoặc Selection Sort thường nhanh hơn Bubble Sort, đặc biệt là đối với dữ liệu không lớn.