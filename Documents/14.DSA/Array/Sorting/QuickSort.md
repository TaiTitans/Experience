
---

So với thuật toán sắp xếp nổi bọt (bubble sort) thì thuật toán sắp xếp nhanh có tốc độ nhanh hơn. Thay vì đi theo sắp xếp từng cặp như bubble sort, chúng ta có thể chia dữ liệu ra thành 22 danh sách, rồi so sánh từng phần tử của danh sách với một phần tử được chọn (gọi là phần tử chốt) và mục đích của chúng ta là đưa phần tử chốt về đúng vị trí của nó.
## II. Miêu tả thuật toán
Chắc hẳn bạn vẫn còn khá mông lung với thuật toán, để giúp bạn hiểu rõ hơn, chúng ta hãy cùng đến với một trò chơi "hành quân" sau:

Xét một dãy số như sau:

6,1,2,7,9,3,4,5,10,86,1,2,7,9,3,4,5,10,8

Yêu cầu là sắp xếp dãy trên theo thứ tự không giảm từ trái qua phải.
![](https://images.viblo.asia/9737ec24-ae95-43c0-905b-84520e17b7ac.png)


https://www.geeksforgeeks.org/quick-sort-algorithm/

---
QuickSort là một thuật toán sắp xếp theo nguyên lý chia để trị (Divide and Conquer). Đây là một trong những thuật toán sắp xếp nhanh và hiệu quả. Quá trình hoạt động của QuickSort có thể tóm tắt như sau:
1. **Chọn một phần tử làm Pivot**: Đầu tiên, chọn một phần tử từ danh sách (mảng). Phần tử này gọi là "pivot".
2. **Chia mảng**: Phân chia mảng thành 2 phần:
    - Mảng con bên trái của pivot chứa các phần tử nhỏ hơn hoặc bằng pivot.
    - Mảng con bên phải của pivot chứa các phần tử lớn hơn pivot.
3. **Đệ quy**: Sau khi phân chia, áp dụng đệ quy để sắp xếp các phần tử ở hai mảng con.
4. **Dừng khi mảng có một phần tử hoặc không có phần tử nào**.

### **Ưu và nhược điểm của QuickSort**

- **Ưu điểm**:
    
    - Thời gian trung bình của QuickSort là **O(n log n)**, nhanh hơn nhiều so với các thuật toán sắp xếp khác như Bubble Sort hoặc Selection Sort.
    - Là thuật toán sắp xếp tại chỗ (in-place), không cần bộ nhớ phụ trợ lớn.
- **Nhược điểm**:
    
    - Trường hợp tồi tệ nhất (khi mảng đã được sắp xếp hoặc gần sắp xếp) có thể lên tới **O(n²)**.
    - Tuy nhiên, nếu chọn pivot một cách ngẫu nhiên hoặc tối ưu, ta có thể tránh được trường hợp này.
### **Cách hoạt động của QuickSort**

1. Chọn một phần tử làm pivot.
2. Di chuyển các phần tử nhỏ hơn pivot sang bên trái, và các phần tử lớn hơn pivot sang bên phải.
3. Đệ quy áp dụng cho phần tử bên trái và bên phải của pivot.
4. Cuối cùng, các phần tử được sắp xếp.

```java
import java.util.Random;

public class QuickSort {
    
    public static int partition(int[] arr, int low, int high) {
        Random rand = new Random();
        // Chọn một chỉ số ngẫu nhiên giữa low và high
        int pivotIndex = rand.nextInt(high - low + 1) + low;
        // Hoán đổi phần tử pivot với phần tử cuối cùng
        int temp = arr[pivotIndex];
        arr[pivotIndex] = arr[high];
        arr[high] = temp;

        int pivot = arr[high];
        int i = low - 1;

        for (int j = low; j < high; j++) {
            if (arr[j] <= pivot) {
                i++;
                temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }
        temp = arr[i + 1];
        arr[i + 1] = arr[high];
        arr[high] = temp;
        return i + 1;
    }

    public static void quickSort(int[] arr, int low, int high) {
        if (low < high) {
            int pi = partition(arr, low, high);
            quickSort(arr, low, pi - 1);
            quickSort(arr, pi + 1, high);
        }
    }

    public static void printArray(int[] arr) {
        for (int i = 0; i < arr.length; i++) {
            System.out.print(arr[i] + " ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        int[] arr = {10, 7, 8, 9, 1, 5};
        System.out.println("Mảng ban đầu:");
        printArray(arr);

        quickSort(arr, 0, arr.length - 1);

        System.out.println("Mảng sau khi sắp xếp:");
        printArray(arr);
    }
}


```

```
Mảng ban đầu:
10 7 8 9 1 5 
Mảng sau khi sắp xếp:
1 5 7 8 9 10 
```
### **Giải thích đoạn mã**

1. **partition()**: Phân chia mảng thành hai phần. Nó chọn phần tử cuối cùng làm pivot và đảm bảo rằng các phần tử nhỏ hơn pivot nằm ở bên trái, còn các phần tử lớn hơn pivot nằm ở bên phải.
2. **quickSort()**: Là hàm đệ quy sắp xếp mảng. Nó gọi hàm `partition()` để phân chia mảng, sau đó đệ quy sắp xếp các phần tử bên trái và bên phải của pivot.
3. **printArray()**: Hàm này chỉ để in mảng ra màn hình.

### **So sánh các phương pháp chọn Pivot**

| Phương pháp chọn pivot | Ưu điểm                                      | Nhược điểm                                    |
| ---------------------- | -------------------------------------------- | --------------------------------------------- |
| Chọn phần tử đầu tiên  | Đơn giản, dễ thực hiện                       | Có thể dẫn đến trường hợp tồi tệ **O(n²)**    |
| Chọn phần tử cuối cùng | Dễ thực hiện                                 | Có thể gặp phải trường hợp tồi tệ **O(n²)**   |
| Chọn phần tử giữa      | Giảm khả năng gặp phải trường hợp tồi tệ     | Không phải lúc nào cũng tối ưu                |
| Chọn ngẫu nhiên        | Tránh được các trường hợp tồi tệ             | Phức tạp hơn và cần một bộ sinh số ngẫu nhiên |
| Chọn median của ba     | Thường phân chia mảng đều, tránh được tồi tệ | Phức tạp hơn và cần một số phép toán          |