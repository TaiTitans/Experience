
---
### **Giải thích chi tiết về Prefix Sum**

**1. Định nghĩa**  
Prefix Sum (còn gọi là mảng tổng tiền tố) là một kỹ thuật giúp **tính tổng một đoạn con của mảng một cách nhanh chóng**. Thay vì tính tổng các phần tử trong một khoảng `(left, right)` theo cách thông thường (duyệt từng phần tử trong khoảng đó), chúng ta **tính trước tổng của các phần tử từ đầu mảng đến mỗi vị trí** và lưu vào một mảng riêng.

---

### **2. Cách xây dựng Prefix Sum**

Giả sử ta có mảng `nums = [-2, 0, 3, -5, 2, -1]`.

#### **Bước 1: Xây dựng mảng `prefixSum`**

Chúng ta tạo một mảng `prefixSum` sao cho:

- `prefixSum[i]` lưu tổng các phần tử từ `nums[0]` đến `nums[i]`.
    
- Công thức:
    
    prefixSum[i]=prefixSum[i−1]+nums[i]
    
    với điều kiện `prefixSum[-1] = 0`.
    

|`i`|`nums[i]`|`prefixSum[i]`|
|---|---|---|
|0|-2|-2|
|1|0|-2 + 0 = -2|
|2|3|-2 + 0 + 3 = 1|
|3|-5|-2 + 0 + 3 - 5 = -4|
|4|2|-2 + 0 + 3 - 5 + 2 = -2|
|5|-1|-2 + 0 + 3 - 5 + 2 - 1 = -3|

Vậy `prefixSum = [-2, -2, 1, -4, -2, -3]`.

---

### **3. Cách sử dụng Prefix Sum để tính tổng nhanh**

Sau khi đã có mảng `prefixSum`, để tính tổng của một đoạn `[left, right]`, ta có công thức:

sumRange(left,right)=prefixSum[right]−prefixSum[left−1]

- Nếu `left = 0`, ta chỉ cần lấy `prefixSum[right]`.
    

Ví dụ:

- `sumRange(0, 2) = prefixSum[2] = 1`
    
- `sumRange(2, 5) = prefixSum[5] - prefixSum[1] = (-3) - (-2) = -1`
    
- `sumRange(0, 5) = prefixSum[5] = -3`
    

**Lợi ích**:

- Sau khi xây dựng mảng `prefixSum` trong **O(N)**, ta có thể trả lời mỗi truy vấn `sumRange(left, right)` trong **O(1)**.

```java
class NumArray {
    private int[] prefixSum;

    public NumArray(int[] nums) {
        int n = nums.length;
        prefixSum = new int[n];
        prefixSum[0] = nums[0];
        for (int i = 1; i < n; i++) {
            prefixSum[i] = prefixSum[i - 1] + nums[i];
        }
    }

    public int sumRange(int left, int right) {
        if (left == 0) return prefixSum[right];
        return prefixSum[right] - prefixSum[left - 1];
    }
}
```
### **5. Độ phức tạp**

|**Giai đoạn**|**Độ phức tạp**|
|---|---|
|Xây dựng `prefixSum`|O(N)|
|Truy vấn `sumRange()`|O(1)|

---

### **6. Khi nào nên dùng Prefix Sum?**

✅ **Nên dùng khi**:

- Có nhiều truy vấn tính tổng (`sumRange`).
    
- Không có hoặc có rất ít phép cập nhật giá trị của `nums[i]`.
    

❌ **Không nên dùng nếu**:

- Cần thay đổi giá trị phần tử trong `nums` nhiều lần (vì `prefixSum` phải cập nhật lại toàn bộ, mất O(N)) → Dùng **Segment Tree** sẽ tốt hơn (`O(logN)` cho cập nhật).
