## String
1 số kiểu xử lí String:
![](../../Assets/Images/JavaCore/String.png)
## Mảng 1 chiều

Mảng 1 chiều (1D array) là một cấu trúc dữ liệu trong lập trình, được sử dụng để lưu trữ một tập hợp các giá trị cùng kiểu dữ liệu. Mảng 1 chiều được đánh chỉ số bắt đầu từ 0 và kích thước của nó được xác định tại thời điểm khai báo.

Cú pháp khai báo:

```
<kiểu dữ liệu>[] <tên mảng> = new <kiểu dữ liệu>[<kích thước>];

int[] numbers = new int[5];

```

Mảng 1 chiều cung cấp phương thức:

- **length**: trả về số lượng phần tử trong mảng.
- **sort**: sắp xếp các phần tử trong mảng theo thứ tự tăng dần.
- **toString**: trả về một chuỗi đại diện cho mảng.

Ví dụ:
```
int[] numbers = {5, 2, 8, 1, 9};
int length = numbers.length; // trả về 5
Arrays.sort(numbers); // sắp xếp mảng theo thứ tự tăng dần
String str = Arrays.toString(numbers); // trả về chuỗi "[1, 2, 5, 8, 9]"
```

## Mảng 2 chiều

Mảng 2 chiều trong Java là một mảng có kích thước nhiều hơn so với mảng 1 chiều. Mảng 2 chiều có thể được coi như một bảng (hay một ma trận) với các phần tử được sắp xếp trong các hàng và cột. Mỗi phần tử trong mảng 2 chiều được truy cập thông qua chỉ số của hàng và cột.

Cú pháp khai báo:
```
dataType[][] arrayName = new dataType[rowSize][colSize];
```

Trong đó:

- **dataType** là kiểu dữ liệu của các phần tử trong mảng.
- **arrayName** là tên của mảng.
- **rowSize** là kích thước của hàng (số lượng phần tử trong hàng) trong mảng 2 chiều.
- **colSize** là kích thước của cột (số lượng phần tử trong cột) trong mảng 2 chiều

Ví dụ:
```
int[][] arr = new int[3][4];
```

Cách truy cập:
Ví dụ
```
arr[0][0] = 1;  // gán giá trị 1 cho phần tử ở hàng 0, cột 0
int x = arr[1][2];  // truy cập giá trị của phần tử ở hàng 1, cột 2 và lưu vào biến x
```

Một số phương thức của mảng 2 chiều:

- **length:** trả về kích thước của mảng.
- **clone():** tạo một bản sao của mảng.
- **equals():** so sánh hai mảng với nhau.

```
int[][] arr1 = {{1, 2}, {3, 4}};
int[][] arr2 = {{1, 2}, {3, 4}};
if (Arrays.deepEquals(arr1, arr2)) {
    System.out.println("Hai mảng bằng nhau.");
}
```