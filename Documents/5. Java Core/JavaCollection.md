1. List

`List trong Java là một interface của Java Collection Framework, cung cấp các phương thức để thao tác với danh sách các phần tử. List cho phép lưu trữ các phần tử trùng nhau và theo thứ tự chèn vào`

Một số phương thức cơ bản của List:

- add(E element): thêm phần tử vào cuối danh sách.
- add(int index, E element): thêm phần tử vào vị trí được chỉ định trong danh sách.
- remove(Object obj): xoá phần tử được chỉ định khỏi danh sách.
- get(int index): trả về phần tử ở vị trí chỉ định trong danh sách.
- size(): trả về số lượng phần tử trong danh sách.
- clear(): xoá tất cả các phần tử trong danh sách.
-> ArrayList, LinkedList, Vector...
Ví dụ:
```
import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        // Khởi tạo một List kiểu String
        List<String> myList = new ArrayList<>();

        // Thêm phần tử vào List
        myList.add("Java");
        myList.add("Python");
        myList.add("C++");

        // Truy cập các phần tử trong List
        System.out.println("Các phần tử trong List là:");
        for (String element : myList) {
            System.out.println(element);
        }

        // Xóa phần tử khỏi List
        myList.remove("Python");
        System.out.println("Sau khi xóa phần tử Python, các phần tử còn lại trong List là:");
        for (String element : myList) {
            System.out.println(element);
        }

        // Kiểm tra sự tồn tại của phần tử trong List
        boolean containsJava = myList.contains("Java");
        System.out.println("List có chứa phần tử Java không? " + containsJava);

        // Lấy kích thước của List
        int size = myList.size();
        System.out.println("Kích thước của List là: " + size);
    }
}


```

2. Set

`Trong Java, Set là một interface trong Collection framework được sử dụng để lưu trữ tập hợp các phần tử không trùng lặp. Nó kế thừa interface Collection và bổ sung các tính năng bổ sung để quản lý các tập hợp các phần tử không trùng lặp.`

Các đặc tính của Set:

- Không chứa các phần tử trùng lặp.
- Có thể chứa một số phần tử null (nếu được phép).
- Không đảm bảo thứ tự của các phần tử.

Một số phương thức phổ biến của Set:

- add(E e): Thêm một phần tử vào Set.
- remove(Object o): Xóa một phần tử khỏi Set.
- contains(Object o): Kiểm tra xem Set có chứa phần tử được chỉ định không.
- size(): Trả về số lượng phần tử trong Set.
- iterator(): Trả về một Iterator để lặp lại các phần tử trong Set.


```
Set<String> set = new HashSet<>();

set.add("apple");
set.add("banana");
set.add("orange");

System.out.println(set); // Output: [orange, banana, apple]

set.remove("banana");

System.out.println(set); // Output: [orange, apple]

System.out.println(set.contains("apple")); // Output: true

System.out.println(set.size()); // Output: 2

Iterator<String> iterator = set.iterator();

while (iterator.hasNext()) {
   System.out.println(iterator.next());
}

```

3. Map

`Trong Java Collections Framework, Map là một interface được sử dụng để lưu trữ một tập hợp các cặp key-value, trong đó mỗi key được liên kết với một giá trị value. Map là một interface con của Collection và được sử dụng để tạo ra các bản đồ dữ liệu trong Java.`

Một số đặc điểm của Map:

- Key trong Map là duy nhất và không thể trùng nhau.
- Giá trị value có thể bị trùng nhau.
- Map không bảo đảm thứ tự của các cặp key-value.
-> HashMap, TreeMap, LinkedHash

```
import java.util.*;

public class MapExample {
    public static void main(String[] args) {
        // Khởi tạo một đối tượng Map sử dụng lớp HashMap
        Map<String, Integer> myMap = new HashMap<>();
        
        // Thêm các cặp key-value vào Map
        myMap.put("apple", 1);
        myMap.put("banana", 2);
        myMap.put("orange", 3);
        
        // Truy xuất giá trị value dựa trên key
        int value = myMap.get("apple");
        System.out.println("Value of key 'apple' is: " + value);
        
        // Lặp qua tất cả các cặp key-value trong Map
        for (Map.Entry<String, Integer> entry : myMap.entrySet()) {
            String key = entry.getKey();
            int val = entry.getValue();
            System.out.println("Key: " + key + ", Value: " + val);
        }
    }
}

```

4. Queue
`Queue là một cấu trúc dữ liệu trong Java Collections Framework, giúp lưu trữ các phần tử theo cơ chế FIFO (First In First Out), tức là phần tử đầu tiên được thêm vào sẽ là phần tử đầu tiên được lấy ra khỏi hàng đợi.`

Một số phương thức quan trọng của Queue:

- add(element): thêm một phần tử vào hàng đợi, nếu hàng đợi không đủ chỗ thì sẽ ném ra ngoại lệ.
- offer(element): thêm một phần tử vào hàng đợi, trả về true nếu thêm thành công, false nếu hàng đợi không đủ chỗ.
- remove(): lấy ra và xóa phần tử đầu tiên của hàng đợi, nếu hàng đợi rỗng sẽ ném ra ngoại lệ.
- poll(): lấy ra và xóa phần tử đầu tiên của hàng đợi, trả về null nếu hàng đợi rỗng.
- peek(): trả về phần tử đầu tiên của hàng đợi mà không xóa nó, trả về null nếu hàng đợi rỗng.
Ví dụ:
```

import java.util.LinkedList;
import java.util.Queue;

public class QueueExample {
    public static void main(String[] args) {
        Queue<String> queue = new LinkedList<>();
        queue.offer("apple");
        queue.offer("banana");
        queue.offer("orange");
        System.out.println("Queue: " + queue); // Queue: [apple, banana, orange]
        System.out.println("Size of queue: " + queue.size()); // Size of queue: 3
        System.out.println("First element of queue: " + queue.peek()); // First element of queue: apple
        System.out.println("Remove first element of queue: " + queue.poll()); // Remove first element of queue: apple
        System.out.println("Queue after removing first element: " + queue); // Queue after removing first element: [banana, orange]
    }
}

```

5. Stack
`Stack (Ngăn xếp) là một cấu trúc dữ liệu dạng danh sách (List) đặc biệt trong đó thao tác chèn và xóa đều chỉ diễn ra ở một đầu của danh sách. Điểm đặc biệt của Stack là người ta gọi nó là cấu trúc dữ liệu LIFO (Last-In-First-Out), tức là phần tử được chèn vào cuối cùng sẽ được lấy ra đầu tiên. Ngăn xếp thường được sử dụng trong các thuật toán liên quan đến việc phân tích cú pháp, lưu trữ trạng thái tạm thời, tìm đường đi (DFS) và các thuật toán sắp xếp, tìm kiếm, đảo ngược chuỗi, thực hiện undo-redo trong các ứng dụng văn bản, đồ họa.`


Stack có 2 phương thức cơ bản là push() để thêm một phần tử vào đỉnh ngăn xếp, và pop() để lấy phần tử ở đỉnh ngăn xếp ra. Ngoài ra, Stack cũng cung cấp một số phương thức khác như empty() để kiểm tra ngăn xếp có rỗng hay không, peek() để lấy giá trị của phần tử đầu tiên mà không xóa nó khỏi ngăn xếp, search() để tìm kiếm vị trí của một phần tử trong ngăn xếp.

Ví dụ:
```
import java.util.*;

public class ReverseString {
    public static void main(String[] args) {
        String str = "Hello world!";
        Stack<Character> stack = new Stack<>();

        // Push characters of string to stack
        for (int i = 0; i < str.length(); i++) {
            stack.push(str.charAt(i));
        }

        // Pop characters from stack to reverse the string
        StringBuilder reversedStr = new StringBuilder();
        while (!stack.empty()) {
            reversedStr.append(stack.pop());
        }

        System.out.println("Original string: " + str);
        System.out.println("Reversed string: " + reversedStr.toString());
    }
}

```

6. Tree

`Tree trong Java Collections Framework thường được sử dụng để tham chiếu đến các cấu trúc dữ liệu dạng cây như Binary Search Tree (BST), AVL Tree, Red-Black Tree, và nhiều loại cây khác.`

Một Tree trong Java là một cấu trúc dữ liệu được tổ chức thành các nút (nodes) và liên kết với nhau bằng các cạnh (edges) để tạo thành một cây. Nút gốc (root node) là nút trên cùng của cây, trong khi nút lá (leaf node) là các nút không có nút con. Các nút khác là các nút nội (internal nodes).

Trong Java, chúng ta có thể sử dụng interface "Tree" để thao tác với các cây như BST hoặc TreeMap, và các phương thức chính bao gồm: add, remove, contains, size, và clear.

```
import java.util.TreeSet;

public class TreeExample {
    public static void main(String[] args) {
        // Tạo một TreeSet để lưu trữ các số nguyên
        TreeSet<Integer> treeSet = new TreeSet<>();

        // Thêm các số vào TreeSet
        treeSet.add(5);
        treeSet.add(2);
        treeSet.add(9);
        treeSet.add(1);
        treeSet.add(3);

        // In ra các phần tử trong TreeSet theo thứ tự tăng dần
        System.out.println("TreeSet: " + treeSet);
    }
}

```