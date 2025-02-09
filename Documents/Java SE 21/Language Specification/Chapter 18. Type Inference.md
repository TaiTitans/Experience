
---
Chương này trong Java Language Specification (JLS) đề cập đến **suy luận kiểu (type inference)**—một cơ chế giúp Java **xác định kiểu dữ liệu** mà không cần lập trình viên phải chỉ rõ.

## **1️⃣ Type Inference là gì?**

Type inference giúp trình biên dịch Java suy luận kiểu mà không cần lập trình viên khai báo tường minh. Điều này giúp **giảm lặp lại code** và **cải thiện tính linh hoạt**.

Ví dụ với `var` (Java 10+):
```java
var name = "John";  // Java tự suy luận kiểu là String
var age = 25;       // Kiểu suy luận là int
```
## **2️⃣ Type Inference trong Generics**

Java hỗ trợ suy luận kiểu trong **generic methods và constructors** để tránh chỉ rõ kiểu dữ liệu khi gọi phương thức.

📌 **Ví dụ không dùng type inference:**
```java
Map<String, List<Integer>> map = new HashMap<String, List<Integer>>();
```
📌 **Dùng type inference với diamond operator (`<>`):**
```java
Map<String, List<Integer>> map = new HashMap<>();
```
✅ **Lợi ích**: Giảm lặp lại code, dễ đọc hơn.

## **3️⃣ Type Inference trong Method Calls**

Suy luận kiểu giúp gọi phương thức generic mà không cần chỉ rõ kiểu.

📌 **Ví dụ không dùng inference:**
```java
List<String> list = Collections.<String>emptyList();
```
📌 **Dùng inference (Java 7+):**
```java
List<String> list = Collections.emptyList(); // Kiểu String được suy luận
```
## **4️⃣ Type Inference với `var` (Java 10+)**

Từ Java 10, `var` giúp khai báo biến mà không cần chỉ rõ kiểu, miễn là trình biên dịch có thể suy luận.
```java
var list = new ArrayList<String>(); // list có kiểu ArrayList<String>
var count = 10; // Kiểu suy luận là int
```
✅ **Lưu ý:** `var` **chỉ dùng được với biến cục bộ**, không áp dụng cho tham số phương thức hoặc thuộc tính lớp.

## **5️⃣ Type Inference trong Lambda Expressions**

Java 8+ có thể suy luận kiểu của tham số trong lambda expressions.

📌 **Ví dụ:**
```java
// Kiểu của (a, b) được suy luận là Integer
BiFunction<Integer, Integer, Integer> add = (a, b) -> a + b;
```
## **6️⃣ Khi nào Type Inference không hoạt động?**

🚫 **Không thể suy luận kiểu nếu có nhiều cách hiểu khác nhau.**
```java
var obj; // ❌ Lỗi: Không thể suy luận kiểu nếu không có giá trị khởi tạo
```
🚫 **Không thể suy luận kiểu nếu kiểu không rõ ràng.**
```java
List<?> list = new ArrayList<>(); // ❌ Lỗi: Không thể suy luận `<?>`
```
## **7️⃣ Tổng kết**

✅ **Type Inference giúp code ngắn gọn hơn, ít trùng lặp.**  
✅ **Hỗ trợ tốt trong generics, lambda, `var`, method calls.**  
✅ **Không thể dùng nếu trình biên dịch không thể suy luận kiểu chính xác.**
