
---
Thuật toán Knuth-Morris-Pratt (KMP) là một thuật toán tìm kiếm chuỗi (string matching) hiệu quả, giúp tìm một chuỗi con (pattern) trong một chuỗi lớn (text) với độ phức tạp thời gian là **O(n + m)**, trong đó:

- `n` là độ dài của chuỗi văn bản (`text`).
- `m` là độ dài của chuỗi mẫu (`pattern`).

KMP hiệu quả hơn thuật toán tìm kiếm chuỗi cơ bản (Brute Force), vốn có độ phức tạp **O(n * m)**, bằng cách sử dụng một bảng **LPS (Longest Prefix Suffix)** để tránh kiểm tra lại những ký tự đã so khớp.

## 1. Ý tưởng chính của thuật toán KMP

### 🌟 Vấn đề của phương pháp Brute Force

Giả sử chúng ta tìm `"ABCDABD"` trong `"ABC ABCDAB ABCDABCDABDE"`, với phương pháp brute force, khi gặp ký tự không khớp, chúng ta phải quay lại và kiểm tra lại các ký tự trước đó.

### ✅ KMP cải thiện bằng cách sử dụng bảng **LPS**

- **LPS (Longest Prefix Suffix)** cho biết tại vị trí nào trong `pattern`, chúng ta có thể tiếp tục kiểm tra mà không cần quay lại đầu chuỗi.
- Khi phát hiện ký tự không khớp, thay vì quay lại đầu pattern, ta có thể sử dụng giá trị trong bảng LPS để biết nên quay lại đâu.
## 2. Xây dựng bảng LPS

LPS giúp tránh kiểm tra lại những phần đã biết là khớp.  
Bảng LPS chứa giá trị của "độ dài lớn nhất của tiền tố cũng là hậu tố" cho mỗi vị trí trong pattern.

**Ví dụ:**  
Pattern: `"ABCDABD"`

|Index|0|1|2|3|4|5|6|
|---|---|---|---|---|---|---|---|
|Pattern|A|B|C|D|A|B|D|
|LPS|0|0|0|0|1|2|0|

### Cách tính:

1. Bắt đầu từ index `1` (vì `LPS[0] = 0`).
2. Nếu `pattern[i] == pattern[length]`, tăng `length`, gán `LPS[i] = length`.
3. Nếu không khớp, kiểm tra giá trị LPS trước đó (`LPS[length - 1]`) để thử lại, nếu `length == 0` thì `LPS[i] = 0`.

## 3. Thuật toán KMP

**Bước 1:** Tạo bảng LPS.  
**Bước 2:** Dùng LPS để duyệt qua chuỗi văn bản.

### **Cài đặt thuật toán KMP**
```java
public class KMP {
    // Hàm tính bảng LPS
    public static int[] computeLPS(String pattern) {
        int m = pattern.length();
        int[] lps = new int[m];
        int len = 0;  // Độ dài của tiền tố cũng là hậu tố
        int i = 1;

        while (i < m) {
            if (pattern.charAt(i) == pattern.charAt(len)) {
                len++;
                lps[i] = len;
                i++;
            } else {
                if (len != 0) {
                    len = lps[len - 1]; // Quay lại vị trí trước đó của LPS
                } else {
                    lps[i] = 0;
                    i++;
                }
            }
        }
        return lps;
    }

    // Thuật toán KMP tìm pattern trong text
    public static void KMPsearch(String text, String pattern) {
        int n = text.length();
        int m = pattern.length();
        int[] lps = computeLPS(pattern);

        int i = 0, j = 0; // i duyệt text, j duyệt pattern
        while (i < n) {
            if (pattern.charAt(j) == text.charAt(i)) {
                i++;
                j++;
            }
            if (j == m) { // Khi j = m, tìm thấy pattern trong text
                System.out.println("Pattern found at index " + (i - j));
                j = lps[j - 1]; // Tiếp tục tìm pattern khác
            } else if (i < n && pattern.charAt(j) != text.charAt(i)) {
                if (j != 0) {
                    j = lps[j - 1]; // Nhảy theo LPS
                } else {
                    i++;
                }
            }
        }
    }

    public static void main(String[] args) {
        String text = "ABC ABCDAB ABCDABCDABDE";
        String pattern = "ABCDABD";
        KMPsearch(text, pattern);
    }
}
```
## 4. Độ phức tạp

- **Tạo bảng LPS:** `O(m)`
- **Duyệt chuỗi văn bản:** `O(n)`

Tổng thể, thuật toán chạy trong **O(n + m)** thời gian, nhanh hơn so với brute force.

---

## 5. Ứng dụng của KMP

- **Tìm kiếm chuỗi trong văn bản lớn** (ví dụ: kiểm tra lỗi chính tả, kiểm tra sự xuất hiện của từ khóa).
- **Tìm kiếm DNA trong phân tích sinh học**.
- **Kiểm tra dữ liệu đầu vào trong xử lý văn bản**.

---

💡 **Tóm tắt lại KMP**

- Xây dựng **bảng LPS** để tối ưu các lần duyệt lại.
- Khi phát hiện sự không khớp, dùng **bảng LPS** để quyết định bước nhảy mà không cần quay lại.
- Độ phức tạp **O(n + m)** giúp xử lý hiệu quả trên văn bản lớn.