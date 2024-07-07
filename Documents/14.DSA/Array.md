
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

