
---
**Quotes**:
*Chạy đua với cuộc sống rất mệt tôi biết. Nhưng trở thành kẻ thất bại thì sẽ khổ.*

## **🔥 1. Giới thiệu về Mảng trong Java**

- **Mảng (Array)** là một **tập hợp các phần tử** có **cùng kiểu dữ liệu** và được lưu trữ liên tiếp trong bộ nhớ.
- Mảng có **kích thước cố định** khi khởi tạo và không thể thay đổi sau đó.
- Mảng trong Java là **đối tượng** và được quản lý bởi **Garbage Collector (GC)**.

📌 **Ví dụ về khai báo mảng:**
```java
int[] numbers;  // Khai báo mảng số nguyên
String[] names; // Khai báo mảng chuỗi
```
## **🔥 2. Khai báo, Khởi tạo và Truy xuất Mảng**

### **a) Khai báo và khởi tạo mảng**

Có 3 cách để khởi tạo mảng trong Java:

1️⃣ **Khai báo trước, khởi tạo sau:**
```java
int[] arr;       // Khai báo mảng
arr = new int[5]; // Khởi tạo mảng có 5 phần tử
```
2️⃣ **Khai báo và khởi tạo cùng lúc:**
```java
int[] arr = new int[5]; // Mảng gồm 5 phần tử mặc định là 0
```
3️⃣ **Khởi tạo bằng danh sách phần tử:**
```java
int[] arr = {1, 2, 3, 4, 5}; // Mảng có giá trị sẵn
```
### **b) Truy xuất phần tử trong mảng**

- Chỉ số (index) của mảng bắt đầu từ `0`.
- Dùng **toán tử `[]`** để truy xuất phần tử.

📌 **Ví dụ:**
```java
int[] numbers = {10, 20, 30, 40, 50};
System.out.println(numbers[0]); // 10
System.out.println(numbers[2]); // 30
```
⚠ **Lỗi phổ biến:** Nếu truy xuất ngoài phạm vi (`index < 0` hoặc `index >= length`), Java sẽ ném lỗi `ArrayIndexOutOfBoundsException`.
```java
System.out.println(numbers[5]); // Lỗi: Index 5 out of bounds
```
## **🔥 3. Thuộc tính `length` của mảng**

Dùng **`array.length`** để lấy độ dài của mảng.

📌 **Ví dụ:**
```java
int[] arr = {1, 2, 3, 4, 5};
System.out.println(arr.length); // Output: 5
```
## **🔥 4. Duyệt mảng (Loop through Arrays)**

Có 3 cách để duyệt mảng:

### **a) Dùng vòng lặp `for` (Cách truyền thống)**
```java
int[] arr = {1, 2, 3, 4, 5};

for (int i = 0; i < arr.length; i++) {
    System.out.println(arr[i]);
}
```
### **b) Dùng vòng lặp `foreach` (Enhanced for-loop)
```java
for (int num : arr) {
    System.out.println(num);
}
```
### c) Dùng `Arrays.stream()` (Java 8+)
```java
import java.util.Arrays;

Arrays.stream(arr).forEach(System.out::println);
```
## **🔥 5. Mảng nhiều chiều (Multidimensional Arrays)**

Java hỗ trợ **mảng 2D, 3D, … (nhiều chiều)**.

📌 **Khai báo mảng 2 chiều:**
```java
int[][] matrix = new int[3][3]; // Mảng 3x3 (9 phần tử)
```
📌 **Gán giá trị ban đầu:**
```java
int[][] matrix = {
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9}
};
```
📌 **Truy xuất phần tử trong mảng 2D:**
```java
System.out.println(matrix[1][2]); // Output: 6
```
📌 **Duyệt mảng 2D bằng vòng lặp lồng nhau:**
```java
for (int i = 0; i < matrix.length; i++) {
    for (int j = 0; j < matrix[i].length; j++) {
        System.out.print(matrix[i][j] + " ");
    }
    System.out.println();
}
```
## **🔥 6. Mảng không đối xứng (Jagged Array)**

- **Mảng không đối xứng** là mảng có số phần tử mỗi hàng khác nhau.

📌 **Ví dụ:**
```java
int[][] jaggedArray = new int[3][]; 
jaggedArray[0] = new int[]{1, 2}; 
jaggedArray[1] = new int[]{3, 4, 5}; 
jaggedArray[2] = new int[]{6};

System.out.println(jaggedArray[1][2]); // Output: 5
```
## **🔥 7. Ép kiểu với Mảng**

### **a) Ép kiểu ngầm định**

Java tự động chuyển đổi kiểu dữ liệu khi mảng có thể chứa kiểu con của lớp cha.

📌 **Ví dụ:**
```java
class Animal {}
class Dog extends Animal {}
class Cat extends Animal {}

Animal[] animals = new Dog[5]; // OK vì Dog là Animal
```
### **b) Ép kiểu tường minh**

Dùng **ép kiểu tường minh** khi chuyển từ `Object[]` về kiểu cụ thể.

📌 **Ví dụ:**
```java
Object[] objArray = new String[]{"Hello", "World"};
String[] strArray = (String[]) objArray; // OK
```
⚠ Nếu ép kiểu sai, sẽ có lỗi `ClassCastException`:
```java
Object[] objArray = new Integer[]{1, 2, 3};
String[] strArray = (String[]) objArray; // Lỗi ClassCastException
```
## **🔥 8. Mảng và Generics**

Java **không thể tạo mảng kiểu generic trực tiếp** vì không an toàn kiểu (`Type Safety`).

📌 **Lỗi phổ biến khi dùng Generics với mảng:**
```java
List<String>[] array = new ArrayList<String>[10]; // Lỗi biên dịch
```
📌 **Cách giải quyết bằng `@SuppressWarnings("unchecked")`:**
```java
@SuppressWarnings("unchecked")
List<String>[] array = (List<String>[]) new ArrayList[10]; // Hợp lệ
```
## **🔥 9. Mảng trong `java.util.Arrays`**

Lớp `java.util.Arrays` hỗ trợ nhiều phương thức tiện ích cho mảng.

📌 **Ví dụ về các phương thức phổ biến:**
```java
import java.util.Arrays;

int[] numbers = {5, 2, 8, 1, 3};

// Sắp xếp mảng
Arrays.sort(numbers); 
System.out.println(Arrays.toString(numbers)); // [1, 2, 3, 5, 8]

// Tìm kiếm phần tử (mảng phải được sắp xếp trước)
int index = Arrays.binarySearch(numbers, 3);
System.out.println("Found at index: " + index); // Output: 2

// Sao chép mảng
int[] copy = Arrays.copyOf(numbers, numbers.length);
System.out.println(Arrays.toString(copy)); // [1, 2, 3, 5, 8]
```

## **🔥 10. Tóm tắt**

✅ Mảng là tập hợp các phần tử **cùng kiểu dữ liệu**.  
✅ **Dùng `array.length`** để lấy kích thước mảng.  
✅ **Mảng nhiều chiều** được khai báo bằng `[][]`.  
✅ **Mảng không đối xứng** có số phần tử mỗi hàng khác nhau.  
✅ **Không thể tạo mảng Generic trực tiếp**.  
✅ **Dùng `Arrays` để xử lý mảng dễ dàng hơn**.

---
# BONUS

### **Mảng trong Java được lưu ở đâu trong bộ nhớ máy tính?**

Trong Java, mảng là một đối tượng (`Object`), vì vậy nó **được lưu trữ trên Heap Memory (Bộ nhớ Heap)**.

> 💡 **Heap Memory** là vùng nhớ dành cho các đối tượng trong Java và được quản lý bởi **Garbage Collector (GC)**.

## **📌 1. Bộ nhớ trong Java**

|**Vùng nhớ**|**Mô tả**|
|---|---|
|**Stack Memory**|Lưu trữ biến cục bộ (local variables), con trỏ tham chiếu, dữ liệu của từng thread.|
|**Heap Memory**|Chứa đối tượng (Object), bao gồm cả các mảng, được quản lý bởi **Garbage Collector**.|
|**Method Area**|Chứa metadata của class, phương thức tĩnh, biến tĩnh.|
|**PC Register**|Lưu trữ địa chỉ lệnh hiện tại của thread.|
|**Native Stack**|Dùng cho các lệnh native, thường liên quan đến JNI (Java Native Interface).|
## **📌 2. Cách mảng được lưu trong bộ nhớ**

### **🔹 a) Khi khai báo và khởi tạo mảng**

Khi bạn tạo một mảng trong Java, trình biên dịch sẽ cấp phát bộ nhớ trên Heap và lưu một tham chiếu đến nó trong Stack.
📌 **Ví dụ 1:**
```java
int[] arr = new int[5]; // Khởi tạo mảng có 5 phần tử
```
📌 **Cách lưu trữ trong bộ nhớ:**

- **`arr` (tham chiếu) lưu trong Stack.**
- **Mảng thực tế (`{0, 0, 0, 0, 0}`) được lưu trên Heap.**

📌 **Hình minh họa:**
```lua
Stack                      Heap
--------------------       ------------------
arr (tham chiếu)  --->    [0, 0, 0, 0, 0]
```

### **🔹 b) Khi mảng chứa kiểu dữ liệu nguyên thủy (Primitive)**

Các kiểu dữ liệu nguyên thủy (`int`, `double`, `boolean`, v.v.) được lưu trực tiếp trên **Heap**, vì mảng là một đối tượng.

📌 **Ví dụ 2:**
```java
int[] numbers = {1, 2, 3, 4, 5};
```
📌 **Cách lưu trữ:**

- **Tham chiếu `numbers` lưu trên Stack.**
- **Mảng `{1, 2, 3, 4, 5}` lưu trên Heap.**
```lua
Stack                      Heap
--------------------       ------------------
numbers (tham chiếu)  ---> [1, 2, 3, 4, 5]
```
### **🔹 c) Khi mảng chứa kiểu dữ liệu tham chiếu (Object)**

Nếu mảng chứa các đối tượng (ví dụ `String`), thì **các phần tử trong mảng chỉ chứa tham chiếu đến đối tượng trên Heap**, chứ không chứa giá trị thực tế.

📌 **Ví dụ 3:**
```java
String[] names = {"Alice", "Bob", "Charlie"};
```
📌 **Cách lưu trữ:**

- **`names` (tham chiếu) lưu trên Stack.**
- **Mảng `String[]` (chứa địa chỉ của từng `String`) nằm trên Heap.**
- **Mỗi `String` cũng là một đối tượng nằm riêng trên Heap.**

📌 **Hình minh họa:**
```java
Stack                      Heap
--------------------       -----------------------------
names (tham chiếu)  --->   [ref1, ref2, ref3] (Mảng)
                          /      |      \
                         v       v       v
                      "Alice"  "Bob"  "Charlie"
```
### **🔹 d) Mảng nhiều chiều (Multidimensional Arrays)**

Mảng nhiều chiều thực chất là **mảng của mảng**. Mỗi phần tử trong mảng **trỏ đến một mảng khác** trên Heap.

📌 **Ví dụ 4:**
```java
int[][] matrix = {
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9}
};
```
📌 **Cách lưu trữ:**

- **`matrix` (tham chiếu) lưu trên Stack.**
- **Mảng `matrix[0]`, `matrix[1]`, `matrix[2]` là các mảng con trên Heap.**
- **Các giá trị số nguyên được lưu trực tiếp trong các mảng con trên Heap.**
```lua
Stack                          Heap
--------------------       -----------------------------
matrix (tham chiếu)  --->   [ref1, ref2, ref3] (Mảng chính)
                          /      |      \
                         v       v       v
                     [1, 2, 3] [4, 5, 6] [7, 8, 9] (Mảng con)
```
## **📌 3. Garbage Collection và Mảng**

Do mảng là **đối tượng trên Heap**, nó sẽ được **Garbage Collector (GC) thu hồi** nếu **không còn tham chiếu nào trỏ đến nó**.

📌 **Ví dụ 5:**
```java
int[] arr = new int[]{1, 2, 3};
arr = null; // Mảng bị GC thu hồi
```
Mảng `{1, 2, 3}` không còn tham chiếu nào → GC sẽ giải phóng bộ nhớ.
## **📌 4. Tổng kết**

✅ **Mảng trong Java là đối tượng, được lưu trên Heap.**  
✅ **Biến tham chiếu của mảng nằm trên Stack.**  
✅ **Nếu mảng chứa kiểu nguyên thủy, giá trị lưu trực tiếp trên Heap.**  
✅ **Nếu mảng chứa đối tượng, chỉ lưu tham chiếu đến đối tượng trên Heap.**  
✅ **Mảng nhiều chiều là mảng của mảng, mỗi hàng là một mảng riêng trên Heap.**  
✅ **Mảng sẽ bị thu hồi nếu không còn tham chiếu nào trỏ đến nó (Garbage Collection).**