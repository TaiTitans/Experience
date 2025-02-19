
---
Moore Voting Algorithm là một thuật toán hiệu quả để tìm **majority element** (phần tử xuất hiện nhiều hơn một nửa số lần) trong một mảng có độ phức tạp **O(n)** và sử dụng **O(1) bộ nhớ**.

### 🟢 **Ý tưởng chính của thuật toán**

Thuật toán hoạt động theo hai bước:

1. **Bước 1: Tìm ứng viên tiềm năng**
    
    - Duyệt qua mảng và sử dụng một biến đếm (`count`) để theo dõi số lần xuất hiện của một phần tử "ứng viên".
    - Nếu `count == 0`, gán phần tử hiện tại làm ứng viên mới.
    - Nếu phần tử hiện tại giống ứng viên, tăng `count`, ngược lại giảm `count`.
2. **Bước 2: Kiểm tra ứng viên**
    
    - Sau khi tìm được ứng viên, duyệt lại mảng để kiểm tra xem nó có thực sự xuất hiện hơn `n/2` lần hay không.
### 📌 **Phân tích thuật toán**

- **Độ phức tạp thời gian**: `O(n)` (do chỉ duyệt qua mảng 2 lần)
- **Độ phức tạp không gian**: `O(1)` (chỉ sử dụng biến `count` và `candidate`)

### 📌 **Tại sao thuật toán hoạt động?**

- Khi `count` về `0`, phần tử hiện tại trở thành ứng viên mới.
- Nếu có một phần tử xuất hiện nhiều hơn `n/2` lần, nó sẽ **luôn tồn tại** đến cuối cùng sau khi giảm dần `count`.
- Kiểm tra lại để đảm bảo ứng viên thực sự là majority element.

### 🎯 **Ứng dụng thực tế**

- Tìm phần tử phổ biến nhất trong tập dữ liệu lớn.
- Phân tích dữ liệu streaming với bộ nhớ hạn chế.
- Hệ thống bỏ phiếu và đánh giá sản phẩm.

### VÍ DỤ 
[Majority Element - LeetCode](https://leetcode.com/problems/majority-element/description/?envType=problem-list-v2&envId=array)
```java
class Solution {

    public int majorityElement(int[] nums) {

        int candidate = nums[0];

        int count = 1;

        for(int num :  nums){

            if(num == candidate){

                count++;

            }else{

                count--;

                if(count == 0){

                    candidate = num;

                    count = 1;

                }

            }

        }

         return candidate;

    }

}
```

