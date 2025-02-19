
---
`Muốn thông thạo thứ gì đó phải bỏ sức ra mà cày.`

## 🔥 **I. Tổng quan về Array trong Java**

### 1️⃣ **Khái niệm**

- **Array (Mảng)** là một cấu trúc dữ liệu lưu trữ tập hợp các phần tử có cùng kiểu dữ liệu trong một vùng nhớ liên tiếp.
- Array có **kích thước cố định** sau khi khởi tạo.
2️⃣ **Cách khai báo & khởi tạo**
```java
int[] arr1 = new int[5]; // Mảng có kích thước 5, giá trị mặc định là 0
int[] arr2 = {1, 2, 3, 4, 5}; // Khởi tạo trực tiếp
```
3️⃣ **Truy cập phần tử**
```java
arr1[0] = 10; // Gán giá trị
System.out.println(arr2[2]); // In phần tử tại index 2 (3)
```
### 4️⃣ **Duyệt mảng**

- **Dùng vòng lặp for**
```java
for (int i = 0; i < arr2.length; i++) {
    System.out.print(arr2[i] + " ");
}
```
- Dùng vòng lặp for-each
```java
for (int num : arr2) {
    System.out.print(num + " ");
}
```
## 🚀 **II. Các kỹ thuật xử lý Array quan trọng**

### 1️⃣ **Tìm phần tử lớn nhất/nhỏ nhất**
```java
int max = arr2[0];
for (int i = 1; i < arr2.length; i++) {
    if (arr2[i] > max) {
        max = arr2[i];
    }
}
System.out.println("Max: " + max);
```
👉 **Độ phức tạp: O(n)**
2️⃣ **Đảo ngược mảng**
```java
int left = 0, right = arr2.length - 1;
while (left < right) {
    int temp = arr2[left];
    arr2[left] = arr2[right];
    arr2[right] = temp;
    left++;
    right--;
}
```
👉 **Độ phức tạp: O(n)**
3️⃣ **Xóa phần tử khỏi mảng (Dùng List)**
```java
List<Integer> list = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
list.remove(Integer.valueOf(3)); // Xóa giá trị 3
System.out.println(list);
```
👉 **Độ phức tạp: O(n)**
4️⃣ **Tìm phần tử xuất hiện nhiều nhất (HashMap)**
```java
int[] nums = {1, 3, 3, 2, 1, 3, 4, 2, 2};
Map<Integer, Integer> freq = new HashMap<>();
for (int num : nums) {
    freq.put(num, freq.getOrDefault(num, 0) + 1);
}

int maxFreq = 0, mostFrequent = -1;
for (var entry : freq.entrySet()) {
    if (entry.getValue() > maxFreq) {
        maxFreq = entry.getValue();
        mostFrequent = entry.getKey();
    }
}
System.out.println("Most Frequent: " + mostFrequent);
```
👉 **Độ phức tạp: O(n)**
## 🔥 **III. Các thuật toán quan trọng liên quan đến Array**

### 1️⃣ **Tìm cặp số có tổng bằng target (Two Sum)**
```java
public static int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    for (int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];
        if (map.containsKey(complement)) {
            return new int[]{map.get(complement), i};
        }
        map.put(nums[i], i);
    }
    return new int[]{-1, -1}; // Không tìm thấy
}
```
👉 **Độ phức tạp: O(n)**
2️⃣ **Dịch chuyển mảng (Rotate Array)**
```java
public static void rotate(int[] nums, int k) {
    k %= nums.length;
    reverse(nums, 0, nums.length - 1);
    reverse(nums, 0, k - 1);
    reverse(nums, k, nums.length - 1);
}

private static void reverse(int[] nums, int left, int right) {
    while (left < right) {
        int temp = nums[left];
        nums[left] = nums[right];
        nums[right] = temp;
        left++;
        right--;
    }
}
```
👉 **Độ phức tạp: O(n)**
3️⃣ **Tìm mảng con có tổng lớn nhất (Kadane’s Algorithm)**
```java
public static int maxSubArray(int[] nums) {
    int maxSum = nums[0], curSum = nums[0];
    for (int i = 1; i < nums.length; i++) {
        curSum = Math.max(nums[i], curSum + nums[i]);
        maxSum = Math.max(maxSum, curSum);
    }
    return maxSum;
}
```
👉 **Độ phức tạp: O(n)**
4️⃣ **Sắp xếp mảng (Quick Sort)**
```java
public static void quickSort(int[] arr, int left, int right) {
    if (left >= right) return;
    int pivot = partition(arr, left, right);
    quickSort(arr, left, pivot - 1);
    quickSort(arr, pivot + 1, right);
}

private static int partition(int[] arr, int left, int right) {
    int pivot = arr[right], i = left - 1;
    for (int j = left; j < right; j++) {
        if (arr[j] < pivot) {
            i++;
            swap(arr, i, j);
        }
    }
    swap(arr, i + 1, right);
    return i + 1;
}

private static void swap(int[] arr, int i, int j) {
    int temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}
```
👉 **Độ phức tạp: O(n log n)**

---
## 🔥 **Cấp độ 1: Hiểu bản chất Array từ góc độ bộ nhớ**

> Khi dùng `int[] arr = new int[5]`, Java thực sự lưu dữ liệu như thế nào?

### 📌 **1️⃣ Bộ nhớ & Cache Efficiency**

- **Array trong Java được lưu trên Heap**, nhưng tham chiếu đến nó lại nằm trên Stack.
- **Truy cập mảng rất nhanh** do các phần tử liên tiếp nhau trong bộ nhớ (tận dụng CPU Cache).
- **Tìm kiếm phần tử mất O(1)** (nếu có index), nhưng tìm kiếm giá trị thì mất **O(n)**.
- **So sánh với LinkedList**:
    - Array truy xuất nhanh hơn do **locality of reference** (liên tiếp trong RAM).
    - LinkedList tốn bộ nhớ do mỗi phần tử có thêm con trỏ (next, prev).

### 📌 **2️⃣ Array là Immutable hay Mutable?**

- **Array trong Java là Mutable** (có thể thay đổi giá trị).
- **Kích thước là Fixed** (không thể thay đổi sau khi tạo).
- **Dùng `ArrayList` thay vì Array nếu cần dynamic size**.
📌 **Ví dụ về hiệu suất bộ nhớ của Array**
```java
int[] arr = new int[1000000]; // Lưu trên Heap ~ 4MB (int chiếm 4 byte)
```
**1. `int[] arr = new int[1000000];`**

- **`int` là kiểu dữ liệu nguyên thủy (primitive type).** Trong Java, `int` là một trong tám kiểu dữ liệu nguyên thủy (byte, short, int, long, float, double, boolean, char). Kiểu dữ liệu nguyên thủy lưu trữ trực tiếp giá trị của nó trong bộ nhớ.
- **`int[]` là một mảng của các kiểu dữ liệu nguyên thủy `int`.** Khi bạn khai báo `int[] arr = new int[1000000];`, bạn đang tạo ra một mảng có 1 triệu phần tử, và mỗi phần tử trong mảng này sẽ lưu trữ một giá trị `int` trực tiếp.
- **Bộ nhớ:** Khi bạn tạo một mảng `int[]`, bộ nhớ được cấp phát **liên tục** để chứa 1 triệu giá trị `int`. Kích thước của mỗi `int` thường là 4 byte (tùy thuộc vào môi trường JVM), vì vậy mảng này sẽ chiếm khoảng 4 triệu byte (4MB) bộ nhớ.
- **Giá trị mặc định:** Khi bạn khởi tạo mảng `int[]`, các phần tử sẽ được gán giá trị mặc định là `0`.
- **Hiệu suất:** Mảng `int[]` thường nhanh hơn và hiệu quả hơn về mặt bộ nhớ khi làm việc với các số nguyên, đặc biệt trong các phép toán số học, vì nó truy cập trực tiếp vào các giá trị số nguyên trong bộ nhớ.
Nhưng với **Integer[]** (wrapper class) thì khác:
```java
Integer[] arr = new Integer[1000000]; // Tốn nhiều bộ nhớ hơn vì chứa Object reference
```
**2. `Integer[] arr = new Integer[1000000];`**

- **`Integer` là một lớp Wrapper (wrapper class).** Trong Java, `Integer` là một lớp bao bọc (wrapper class) cho kiểu dữ liệu nguyên thủy `int`. Nó là một **đối tượng** đại diện cho một giá trị số nguyên.

- **`Integer[]` là một mảng của các đối tượng `Integer`.** Khi bạn khai báo `Integer[] arr = new Integer[1000000];`, bạn đang tạo ra một mảng có 1 triệu phần tử, và mỗi phần tử trong mảng này sẽ là một **tham chiếu** (reference) đến một đối tượng `Integer`.

- **Bộ nhớ:** Khi bạn tạo mảng `Integer[]`, bộ nhớ được cấp phát **để lưu trữ các tham chiếu**. Ban đầu, mảng chỉ chứa các tham chiếu `null` (giá trị mặc định cho tham chiếu đối tượng). Để lưu trữ một giá trị số nguyên, bạn cần tạo một **đối tượng `Integer` riêng biệt** cho mỗi phần tử và gán tham chiếu của đối tượng đó vào mảng. Mỗi đối tượng `Integer` sẽ được cấp phát bộ nhớ trên heap (vùng nhớ động). Điều này dẫn đến việc sử dụng bộ nhớ **lớn hơn đáng kể** so với `int[]`. Ngoài bộ nhớ cho bản thân giá trị số nguyên, còn có chi phí bộ nhớ cho đối tượng `Integer` (ví dụ: overhead cho header của đối tượng, tham chiếu đến lớp, v.v.).

- **Giá trị mặc định:** Khi bạn khởi tạo mảng `Integer[]`, các phần tử sẽ được gán giá trị mặc định là `null`. Điều này có nghĩa là ban đầu, mảng chứa 1 triệu tham chiếu `null`, và bạn cần gán các đối tượng `Integer` thực tế vào từng vị trí nếu bạn muốn lưu trữ các giá trị số nguyên.

- **Hiệu suất:** Mảng `Integer[]` thường chậm hơn và kém hiệu quả hơn về mặt bộ nhớ so với `int[]`. Khi bạn truy cập một phần tử trong `Integer[]`, bạn thực sự đang làm việc với một **tham chiếu** đến một đối tượng trên heap, cần phải dereference tham chiếu này để lấy giá trị số nguyên. Ngoài ra, việc tạo và quản lý các đối tượng `Integer` có thêm chi phí hiệu suất.

👉 **Lời khuyên**: Nếu cần tối ưu bộ nhớ, **luôn ưu tiên sử dụng `int[]` thay vì `Integer[]`**.

**Khi nào nên sử dụng loại nào?**

- **`int[]`**:
    
    - Khi bạn cần hiệu suất cao và tối ưu hóa bộ nhớ, đặc biệt khi làm việc với số lượng lớn các số nguyên.
    - Trong các phép toán số học, xử lý dữ liệu số, thuật toán,...
    - Khi bạn không cần lưu trữ giá trị `null` trong mảng.
- **`Integer[]`**:
    
    - Khi bạn cần lưu trữ giá trị `null` trong mảng. Ví dụ: để biểu thị một giá trị "không có" hoặc "chưa xác định".
    - Khi bạn cần sử dụng các tính năng của đối tượng `Integer`, ví dụ như các phương thức của lớp `Integer` (ví dụ: `toString()`, `parseInt()`, ...), hoặc khi bạn cần sử dụng mảng trong các cấu trúc dữ liệu hoặc API yêu cầu mảng của đối tượng (ví dụ: Collections trong Java chỉ làm việc với đối tượng).
    - Khi bạn làm việc với Generic Collections (ví dụ: `List<Integer>`, `ArrayList<Integer>`), vì Generic Collections chỉ có thể chứa các đối tượng, không thể chứa các kiểu dữ liệu nguyên thủy trực tiếp.

**Ví dụ minh họa về bộ nhớ (ước tính):**

Giả sử một tham chiếu đối tượng (reference) chiếm 8 byte và overhead của một đối tượng `Integer` là khoảng 16 byte (ước tính, có thể thay đổi tùy thuộc JVM).

- **`int[] arr = new int[1000000];`**:
    
    - Bộ nhớ ≈ 1,000,000 * 4 byte (cho `int`) = 4,000,000 byte ≈ 4 MB
- **`Integer[] arr = new Integer[1000000];`**:
    
    - Bộ nhớ cho mảng tham chiếu ≈ 1,000,000 * 8 byte (cho tham chiếu) = 8,000,000 byte ≈ 8 MB
    - Nếu bạn gán giá trị cho từng phần tử, bạn sẽ tạo 1 triệu đối tượng `Integer`: 1,000,000 * (overhead đối tượng + kích thước `int`) ≈ 1,000,000 * (16 byte + 4 byte) = 20,000,000 byte ≈ 20 MB
    - **Tổng cộng ≈ 8 MB (tham chiếu mảng) + 20 MB (đối tượng Integer) = 28 MB (ước tính, có thể lớn hơn)**

Như bạn thấy, `Integer[]` có thể chiếm bộ nhớ lớn hơn nhiều so với `int[]` khi lưu trữ cùng số lượng phần tử.

## 🚀 **Cấp độ 2: Các thuật toán nâng cao với Array**

### 📌 **1️⃣ Sliding Window (O(n))**

> Tìm dãy con có tổng lớn nhất trong O(n).

```java
public static int maxSubArray(int[] nums) {
    int maxSum = nums[0], curSum = nums[0];
    for (int i = 1; i < nums.length; i++) {
        curSum = Math.max(nums[i], curSum + nums[i]);
        maxSum = Math.max(maxSum, curSum);
    }
    return maxSum;
}
```
⏳ **Ứng dụng**: Tìm dãy con có tổng bằng K, tìm dãy con thỏa mãn điều kiện.

### 📌 **2️⃣ Monotonic Stack (O(n))**

> Tìm phần tử lớn nhất/nhỏ nhất gần nhất về bên trái/phải.

```java
public static int[] nextGreaterElement(int[] nums) {
    Stack<Integer> stack = new Stack<>();
    int[] result = new int[nums.length];
    Arrays.fill(result, -1);

    for (int i = 0; i < nums.length; i++) {
        while (!stack.isEmpty() && nums[stack.peek()] < nums[i]) {
            result[stack.pop()] = nums[i];
        }
        stack.push(i);
    }
    return result;
}
```
⏳ **Ứng dụng**: Stock span, histogram, subarray problems.

### 📌 **3️⃣ Sparse Table (O(1) query, O(n log n) build)**

> Tìm **giá trị nhỏ nhất/lớn nhất trong một khoảng [L, R]** siêu nhanh.

- **Chuẩn bị bảng `st[i][j]`**, mỗi cell lưu giá trị nhỏ nhất của đoạn `2^j` tại vị trí `i`.

```java
int[][] st; 

void build(int[] arr) {
    int n = arr.length;
    int k = (int) (Math.log(n) / Math.log(2)) + 1;
    st = new int[n][k];

    for (int i = 0; i < n; i++) st[i][0] = arr[i];

    for (int j = 1; (1 << j) <= n; j++) {
        for (int i = 0; i + (1 << j) - 1 < n; i++) {
            st[i][j] = Math.min(st[i][j - 1], st[i + (1 << (j - 1))][j - 1]);
        }
    }
}
```
⏳ **Ứng dụng**: Tìm min/max trong một đoạn **siêu nhanh**, đặc biệt cho **hệ thống real-time**.