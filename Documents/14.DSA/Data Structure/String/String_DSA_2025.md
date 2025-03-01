
---
Học String trong DSA (Data Structures & Algorithms) bằng Java là một chủ đề quan trọng vì chuỗi (String) được sử dụng rộng rãi trong lập trình và thường xuất hiện trong các bài toán thuật toán. Tôi sẽ hướng dẫn bạn từ cơ bản đến nâng cao, bao gồm:

1. **Cơ bản về String trong Java**
2. **Các thuật toán và kỹ thuật xử lý String**
3. **Các cấu trúc dữ liệu quan trọng để xử lý String**
4. **Bài toán và cách tối ưu hóa xử lý String**

## 🔹 1. Cơ bản về String trong Java

### 🔹 1.1 Khai báo và tạo chuỗi

Trong Java, `String` là **immutable** (bất biến), tức là một khi tạo ra thì không thể thay đổi giá trị.
```java
String str1 = "Hello";  // String Literal
String str2 = new String("Hello"); // String Object
```
📌 **Lưu ý**:

- `"Hello"` được lưu trong **String Pool** (bộ nhớ đặc biệt trong Heap), nếu một chuỗi giống hệt đã tồn tại, nó sẽ tái sử dụng.
- `new String("Hello")` luôn tạo ra một đối tượng mới trong Heap, không tái sử dụng.

🔹 1.2 Các phương thức quan trọng
```java
String s = "Hello, World!";
System.out.println(s.length());        // Độ dài chuỗi
System.out.println(s.charAt(1));       // Truy xuất ký tự tại vị trí
System.out.println(s.substring(0, 5)); // Cắt chuỗi (Hello)
System.out.println(s.indexOf("o"));    // Vị trí của ký tự 'o'
System.out.println(s.contains("World")); // Kiểm tra chuỗi con
System.out.println(s.replace("World", "Java")); // Thay thế chuỗi con
System.out.println(s.toUpperCase());   // Viết hoa
System.out.println(s.toLowerCase());   // Viết thường
System.out.println(s.trim());          // Xóa khoảng trắng thừa đầu/cuối
```
### 🔹 1.3 StringBuilder & StringBuffer (Mutable Strings)

Nếu bạn cần thao tác nhiều trên chuỗi, hãy sử dụng `StringBuilder` hoặc `StringBuffer` để tối ưu hiệu suất.
```java
StringBuilder sb = new StringBuilder("Hello");
sb.append(" World");  // Thêm vào cuối
sb.insert(5, ",");    // Chèn tại vị trí 5
sb.delete(5, 6);      // Xóa ký tự
System.out.println(sb.toString());
```
📌 **Khác biệt giữa String, StringBuilder và StringBuffer**:

| Loại            | Tốc độ     | Thread-Safe | Mutable |
| --------------- | ---------- | ----------- | ------- |
| `String`        | Chậm       | Có          | ❌ Không |
| `StringBuilder` | Nhanh      | ❌ Không     | ✅ Có    |
| `StringBuffer`  | Trung bình | ✅ Có        | ✅ Có    |
## 🔹 2. Các thuật toán và kỹ thuật xử lý String

### 🔹 2.1 Đảo ngược chuỗi
```java
public static String reverseString(String s) {
    return new StringBuilder(s).reverse().toString();
}
```
📌 **Không dùng StringBuilder**:
```java
public static String reverseString(String s) {
    char[] arr = s.toCharArray();
    int left = 0, right = arr.length - 1;
    while (left < right) {
        char temp = arr[left];
        arr[left] = arr[right];
        arr[right] = temp;
        left++;
        right--;
    }
    return new String(arr);
}
```
### 🔹 2.2 Kiểm tra chuỗi đối xứng (Palindrome)
```java
public static boolean isPalindrome(String s) {
    int left = 0, right = s.length() - 1;
    while (left < right) {
        if (s.charAt(left) != s.charAt(right)) return false;
        left++;
        right--;
    }
    return true;
}
```
🔹 2.3 Đếm số lần xuất hiện của ký tự
```java
public static int countOccurrences(String s, char c) {
    int count = 0;
    for (char ch : s.toCharArray()) {
        if (ch == c) count++;
    }
    return count;
}
```
### 🔹 2.4 Kiểm tra chuỗi con

**KMP Algorithm** (O(n + m)):
```java
public static boolean containsSubstring(String text, String pattern) {
    return text.contains(pattern);
}
```
📌 Để tối ưu, ta có thể dùng thuật toán **KMP (Knuth-Morris-Pratt)**.

## 🔹 3. Các cấu trúc dữ liệu quan trọng để xử lý String

### 🔹 3.1 HashMap - Đếm số lần xuất hiện của ký tự
```java
public static Map<Character, Integer> countCharFrequency(String s) {
    Map<Character, Integer> map = new HashMap<>();
    for (char c : s.toCharArray()) {
        map.put(c, map.getOrDefault(c, 0) + 1);
    }
    return map;
}
```
### 🔹 3.2 Trie - Tìm kiếm chuỗi con nhanh

Trie giúp tìm kiếm chuỗi nhanh hơn so với HashMap.
```java
class TrieNode {
    Map<Character, TrieNode> children = new HashMap<>();
    boolean isEndOfWord = false;
}

class Trie {
    private final TrieNode root = new TrieNode();

    public void insert(String word) {
        TrieNode node = root;
        for (char c : word.toCharArray()) {
            node.children.putIfAbsent(c, new TrieNode());
            node = node.children.get(c);
        }
        node.isEndOfWord = true;
    }

    public boolean search(String word) {
        TrieNode node = root;
        for (char c : word.toCharArray()) {
            if (!node.children.containsKey(c)) return false;
            node = node.children.get(c);
        }
        return node.isEndOfWord;
    }
}
```
## 🔹 4. Bài toán và cách tối ưu hóa xử lý String

### 🔹 4.1 Tìm chuỗi con dài nhất không lặp ký tự (**Sliding Window**)

```java
public static int longestUniqueSubstring(String s) {
    Map<Character, Integer> map = new HashMap<>();
    int left = 0, maxLen = 0;
    for (int right = 0; right < s.length(); right++) {
        char c = s.charAt(right);
        if (map.containsKey(c)) left = Math.max(left, map.get(c) + 1);
        map.put(c, right);
        maxLen = Math.max(maxLen, right - left + 1);
    }
    return maxLen;
}
```
### 🔹 4.2 Tìm chuỗi con có tần suất ký tự lớn nhất (**Sliding Window**)

```java
public static int characterReplacement(String s, int k) {
    int[] count = new int[26];
    int left = 0, maxCount = 0, maxLength = 0;
    for (int right = 0; right < s.length(); right++) {
        maxCount = Math.max(maxCount, ++count[s.charAt(right) - 'A']);
        if (right - left + 1 - maxCount > k) count[s.charAt(left++) - 'A']--;
        maxLength = Math.max(maxLength, right - left + 1);
    }
    return maxLength;
}
```
## 🔹 Tổng kết

✅ **Cách xử lý String trong Java hiệu quả**  
✅ **Các thuật toán và kỹ thuật quan trọng**  
✅ **Ứng dụng cấu trúc dữ liệu như HashMap, Trie**  
✅ **Tối ưu thuật toán bằng Sliding Window, KMP**

---
## 🔹 **String được lưu trong bộ nhớ như thế nào trong Java?**

Trong Java, `String` là **immutable (bất biến)**, có nghĩa là khi một `String` được tạo, nó **không thể thay đổi nội dung**. Điều này ảnh hưởng đến cách `String` được lưu trong bộ nhớ.

## **1. Bộ nhớ Heap và String Pool**

Java lưu trữ `String` trong **Heap Memory**, cụ thể hơn là **String Pool (Bể nhớ chuỗi)**.

### 🏗 **Cách hoạt động của String Pool**

- Khi bạn tạo một `String` bằng cách **gán trực tiếp** (`String s = "Hello";`), Java sẽ **kiểm tra trong String Pool** xem chuỗi đó đã tồn tại chưa:
    - Nếu có, nó sẽ **tái sử dụng** địa chỉ bộ nhớ cũ.
    - Nếu chưa có, nó sẽ **tạo mới và lưu trong String Pool**.
- Khi bạn tạo `String` bằng `new`, chuỗi **luôn được tạo mới trong Heap Memory**, không vào Pool.

