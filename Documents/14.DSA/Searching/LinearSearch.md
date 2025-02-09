
---
**Tìm kiếm tuyến tính (Linear Search)**: Tìm phần tử bằng cách duyệt qua từng phần tử trong mảng.
```java
public static int linearSearch(int[] arr, int target) {
    for (int i = 0; i < arr.length; i++) {
        if (arr[i] == target) {
            return i; // Trả về chỉ số nếu tìm thấy
        }
    }
    return -1; // Trả về -1 nếu không tìm thấy
}
```

