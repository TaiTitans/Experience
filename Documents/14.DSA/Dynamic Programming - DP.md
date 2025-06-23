
---
**Quy hoạch động (Dynamic Programming - DP)** là một kỹ thuật giải thuật mạnh mẽ trong cấu trúc dữ liệu và giải thuật (DSA). Nó được sử dụng để giải quyết các bài toán tối ưu hóa hoặc bài toán có tính chất "lặp lại con" (overlapping subproblems) và "cấu trúc con tối ưu" (optimal substructure).

Dưới đây là hướng dẫn chi tiết để hiểu và sử dụng quy hoạch động:

## **1. Ý tưởng cơ bản**

- **Overlapping Subproblems (Lặp lại con):** Bài toán lớn có thể được chia nhỏ thành các bài toán con, và các bài toán con này được lặp lại nhiều lần.
    
- **Optimal Substructure (Cấu trúc con tối ưu):** Lời giải của bài toán lớn có thể được xây dựng từ lời giải của các bài toán con.
    

Thay vì giải đi giải lại các bài toán con giống nhau, quy hoạch động lưu trữ kết quả của chúng (memoization hoặc bảng) để tái sử dụng.

---

## **2. Cách tiếp cận**

Có hai cách tiếp cận chính trong quy hoạch động:

### **a. Top-Down (Memoization):**

- Dùng đệ quy để giải bài toán từ lớn đến nhỏ.
    
- Lưu trữ kết quả các bài toán con vào bảng (thường là mảng hoặc hashmap) để tránh tính toán lặp lại.

Ví dụ:
```java
// Fibonacci với Memoization
int[] dp = new int[n + 1];
Arrays.fill(dp, -1);

int fib(int n) {
    if (n <= 1) return n;
    if (dp[n] != -1) return dp[n]; // Nếu đã tính trước, trả kết quả từ bảng.
    dp[n] = fib(n - 1) + fib(n - 2); // Lưu kết quả vào bảng.
    return dp[n];
}
```

### **b. Bottom-Up (Tabulation):**

- Xây dựng lời giải từ bài toán nhỏ nhất lên bài toán lớn nhất.
    
- Sử dụng một bảng để lưu trữ kết quả các bài toán con.
    

Ví dụ:
```java
// Fibonacci với Tabulation
int fib(int n) {
    if (n <= 1) return n;
    int[] dp = new int[n + 1];
    dp[0] = 0;
    dp[1] = 1;
    for (int i = 2; i <= n; i++) {
        dp[i] = dp[i - 1] + dp[i - 2];
    }
    return dp[n];
}
```

## **3. Các bước giải bài toán quy hoạch động**

1. **Xác định bài toán con lặp lại (Overlapping Subproblems):**
    
    - Chia bài toán lớn thành các bài toán con nhỏ hơn.
        
    - Tìm cách tái sử dụng các kết quả của bài toán con.
        
2. **Xác định trạng thái (State):**
    
    - Trạng thái là một biểu diễn của bài toán con.
        
    - Ví dụ: Với bài toán đường đi ngắn nhất, trạng thái có thể là `(i, j)` biểu diễn chi phí đường đi từ điểm (i, j).
        
3. **Xác định công thức truy hồi (State Transition):**
    
    - Công thức truy hồi là cách tính toán lời giải từ các trạng thái con.
        
    - Ví dụ: Với bài toán Fibonacci, `fib(n) = fib(n-1) + fib(n-2)`.
        
4. **Xác định điều kiện cơ sở (Base Case):**
    
    - Đây là trường hợp đặc biệt hoặc nhỏ nhất mà lời giải đã biết sẵn.
        
    - Ví dụ: Fibonacci có `fib(0) = 0` và `fib(1) = 1`.
        
5. **Cài đặt bằng Memoization hoặc Tabulation.**


## ** Mẹo tối ưu**

- **Dùng không gian:** Nếu chỉ cần kết quả cuối cùng, thay vì lưu cả bảng, bạn có thể lưu các giá trị cần thiết (thường là hàng hoặc cột trước đó).
    
- **Nhận biết bài toán quy hoạch động:** Nếu bài toán có tính chất "lặp lại con" và "cấu trúc con tối ưu", thì đó có thể là bài toán dùng quy hoạch động.


