
---
	Mảng là một cấu trúc dữ liệu lưu trữ một tập hợp các phần tử có cùng kiểu dữ liệu. Các phần tử trong mảng được lưu trữ liên tiếp trong bộ nhớ và có thể truy cập thông qua chỉ số (index).

#### Khai Báo Mảng

Khai báo mảng trong Java có thể thực hiện bằng cách:

```Java
// Khai báo mảng một chiều
int[] arr;

// Khai báo mảng hai chiều
int[][] matrix;

```

#### Khởi Tạo Mảng

Khi khai báo mảng, bạn cần khởi tạo mảng để cấp phát bộ nhớ cho nó. Có thể khởi tạo mảng bằng cách:
```Java
// Khởi tạo mảng một chiều với kích thước 5
int[] arr = new int[5];

// Khởi tạo mảng một chiều với các giá trị cụ thể
int[] arr = {1, 2, 3, 4, 5};

// Khởi tạo mảng hai chiều 3x3
int[][] matrix = new int[3][3];

// Khởi tạo mảng hai chiều với các giá trị cụ thể
int[][] matrix = {
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9}
};

```
### Truy Cập Phần Tử Trong Mảng

Bạn có thể truy cập các phần tử trong mảng bằng chỉ số (bắt đầu từ 0):

```Java
int[] arr = {1, 2, 3, 4, 5};

// Truy cập phần tử đầu tiên
int firstElement = arr[0];

// Truy cập phần tử thứ ba
int thirdElement = arr[2];

// Cập nhật giá trị phần tử thứ hai
arr[1] = 10;

```

### Duyệt Qua Mảng

Có nhiều cách để duyệt qua các phần tử của mảng:

#### Sử Dụng Vòng Lặp For

```Java
int[] arr = {1, 2, 3, 4, 5};

for (int i = 0; i < arr.length; i++) {
    System.out.println(arr[i]);
}

```

Sử Dụng Vòng Lặp For-Each
```Java
int[] arr = {1, 2, 3, 4, 5};

for (int num : arr) {
    System.out.println(num);
}

```

---
Dưới đây là danh sách các kỹ thuật và thuật toán xử lý mảng phổ biến:

### Kỹ thuật tìm kiếm (Searching Techniques):

1. Tìm kiếm tuyến tính (Linear Search)
2. Tìm kiếm nhị phân (Binary Search)
3. Tìm kiếm nội suy (Interpolation Search)
4. Tìm kiếm Fibonacci (Fibonacci Search)

### Kỹ thuật sắp xếp (Sorting Techniques):

1. Bubble Sort <
2. Selection Sort <
3. Insertion Sort <
4. Merge Sort <
5. Quick Sort <
6. Heap Sort <
7. Radix Sort
8. Counting Sort
9. Bucket Sort
10. Shell Sort

### Kỹ thuật thao tác trên mảng (Array Manipulation):

1. Đảo ngược mảng (Reversing an Array)
2. Xoay mảng (Rotating an Array)
3. Tách và hợp mảng (Splitting and Merging Arrays)
4. Loại bỏ phần tử trùng lặp (Removing Duplicates)
5. Chèn và xóa phần tử (Insertion and Deletion)

### Kỹ thuật tìm kiếm nâng cao (Advanced Searching Techniques):

1. Tìm số xuất hiện lẻ (Find the Odd Occurring Element)
2. Tìm cặp số có tổng bằng giá trị cho trước (Pair with a Given Sum)
3. Tìm mảng con có tổng lớn nhất (Maximum Subarray Sum - Kadane’s Algorithm)
4. Tìm phần tử phổ biến nhất (Most Frequent Element)

### Kỹ thuật sắp xếp nâng cao (Advanced Sorting Techniques):

1. Sắp xếp bằng thư viện (Library Sort)
2. Sắp xếp kết hợp (Hybrid Sort - Tim Sort)

### Các thuật toán tối ưu khác (Other Optimization Algorithms):

1. Thuật toán Kadane’s (Kadane's Algorithm for Maximum Subarray Sum)
2. Thuật toán phân loại nhanh hơn (Optimized Sorting Algorithms)
3. Thuật toán tìm kiếm và sắp xếp đồng thời (Simultaneous Searching and Sorting)