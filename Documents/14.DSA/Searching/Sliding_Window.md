
---
Sliding Window là một kỹ thuật tối ưu hóa quan trọng trong lập trình, đặc biệt hữu ích khi xử lý các bài toán về mảng, chuỗi hoặc dãy số có tính liên tục. Nó giúp giảm độ phức tạp thời gian từ O(N^2)xuống O(N) trong nhiều trường hợp bằng cách tránh lặp lại tính toán không cần thiết.

## 🔥 **1. Khái Niệm Sliding Window**

Sliding Window là một phương pháp trong đó chúng ta duy trì một cửa sổ con (subarray hoặc substring) trên một phần của cấu trúc dữ liệu (thường là mảng hoặc chuỗi) và di chuyển cửa sổ này dọc theo tập dữ liệu bằng cách thêm phần tử mới và loại bỏ phần tử cũ.

## 🎯 **2. Các Loại Sliding Window**

Có hai loại Sliding Window chính:

### **2.1. Fixed Size Sliding Window (Cửa sổ có kích thước cố định)**

- Cửa sổ có kích thước cố định kkk.
- Khi di chuyển, ta bỏ phần tử đầu và thêm phần tử mới vào cuối.
- Thường áp dụng cho các bài toán tìm giá trị lớn nhất, nhỏ nhất, tổng lớn nhất,... trong một cửa sổ có kích thước xác định.

**Ví dụ:** Tìm tổng lớn nhất của một dãy con có kích thước k.

🔹 **Bài toán:**  
Cho một mảng arr=[2,1,5,1,3,2]và một số nguyên k=3, hãy tìm tổng lớn nhất của một dãy con có kích thước k.

🔹 **Giải pháp Brute Force (O(N^2)):**  
Duyệt qua tất cả các cửa sổ con có kích thước kkk, tính tổng và chọn tổng lớn nhất.

🔹 **Giải pháp Sliding Window (O(N)):**

- Bước 1: Tính tổng của cửa sổ đầu tiên.
- Bước 2: Dịch cửa sổ bằng cách loại bỏ phần tử đầu và thêm phần tử mới.
- Bước 3: Cập nhật tổng lớn nhất nếu cần.
```java
public int maxSumSubarray(int[] arr, int k) {
    int maxSum = 0, windowSum = 0;
    
    // Tính tổng của cửa sổ đầu tiên
    for (int i = 0; i < k; i++) {
        windowSum += arr[i];
    }
    maxSum = windowSum;

    // Trượt cửa sổ
    for (int i = k; i < arr.length; i++) {
        windowSum += arr[i] - arr[i - k]; // Thêm phần tử mới, loại bỏ phần tử cũ
        maxSum = Math.max(maxSum, windowSum);
    }
    
    return maxSum;
}
```
⏳ **Độ phức tạp:** O(N) do mỗi phần tử được thêm vào và loại bỏ đúng 1 lần.

### **2.2. Dynamic Size Sliding Window (Cửa sổ có kích thước thay đổi)**

- Cửa sổ có thể mở rộng hoặc thu nhỏ dựa trên điều kiện bài toán.
- Dùng phổ biến khi làm việc với chuỗi hoặc dãy số có điều kiện.

**Ví dụ:** Tìm độ dài nhỏ nhất của dãy con có tổng >= SSS.

🔹 **Bài toán:**  
Cho mảng arr=[2,1,5,2,3,2] và S=7, tìm độ dài nhỏ nhất của dãy con có tổng lớn hơn hoặc bằng SSS.

🔹 **Giải pháp Sliding Window:**

- Bắt đầu với cửa sổ rỗng.
- Mở rộng cửa sổ bằng cách thêm phần tử vào.
- Khi tổng >= S, thử thu nhỏ cửa sổ để tìm dãy ngắn nhất.
```java
public int minSubarrayLen(int S, int[] arr) {
    int minLength = Integer.MAX_VALUE;
    int left = 0, windowSum = 0;

    for (int right = 0; right < arr.length; right++) {
        windowSum += arr[right];

        // Thu nhỏ cửa sổ khi tổng >= S
        while (windowSum >= S) {
            minLength = Math.min(minLength, right - left + 1);
            windowSum -= arr[left];
            left++; // Dịch cửa sổ sang phải
        }
    }

    return minLength == Integer.MAX_VALUE ? 0 : minLength;
}
```
⏳ **Độ phức tạp:** O(N) do mỗi phần tử chỉ được duyệt qua tối đa 2 lần (một lần khi mở rộng cửa sổ, một lần khi thu hẹp).


## 🔥 **3. Ứng Dụng Thực Tế**

Sliding Window cực kỳ hữu ích trong nhiều bài toán thực tế như:

1. **Tìm tổng lớn nhất của dãy con có độ dài kkk.**
2. **Tìm số lượng ký tự duy nhất lớn nhất trong một chuỗi con.**
3. **Tìm chuỗi con có tổng nhỏ nhất thỏa mãn điều kiện.**
4. **Kiểm tra một chuỗi có chứa tất cả ký tự của một chuỗi khác không.**
---
### Khi nào dùng Sliding Window?

- Bài toán yêu cầu xử lý trên **đoạn con liên tục** (contiguous subarray/substring).
- Có thể tính toán kết quả của cửa sổ mới dựa trên cửa sổ cũ mà không cần duyệt lại toàn bộ.
- Ví dụ:
    - Tìm tổng lớn nhất của đoạn con độ dài k.
    - Tìm số lần xuất hiện tối thiểu của một ký tự trong đoạn con.
    - Tìm đoạn con ngắn nhất chứa tất cả các ký tự cần thiết.

### Cách hoạt động của Sliding Window

#### Với Fixed Size Sliding Window:

1. **Khởi tạo**: Tính kết quả cho cửa sổ đầu tiên (k phần tử đầu tiên).
2. **Trượt**:
    - Loại bỏ phần tử đầu cửa sổ cũ.
    - Thêm phần tử mới vào cuối cửa sổ.
    - Cập nhật kết quả (tổng, số lượng, min/max, v.v.).
3. **Lặp lại**: Tiếp tục trượt cho đến khi duyệt hết mảng/chuỗi.

#### Với Variable Size Sliding Window:

1. **Khởi tạo**: Bắt đầu với cửa sổ rỗng hoặc nhỏ nhất.
2. **Mở rộng**: Thêm phần tử mới vào cửa sổ cho đến khi thỏa mãn điều kiện.
3. **Thu hẹp**: Nếu có thể, loại bỏ phần tử từ đầu để tối ưu hóa (vẫn giữ điều kiện).
4. **Lặp lại**: Tiếp tục mở rộng và thu hẹp cho đến hết.