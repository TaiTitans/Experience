
---
1. Đảo ngược mảng (Reversing an Array)
```java
public class ArrayManipulation {
    public static void reverseArray(int[] array) {
        int left = 0, right = array.length - 1;
        while (left < right) {
            int temp = array[left];
            array[left] = array[right];
            array[right] = temp;
            left++;
            right--;
        }
    }
}
```
2. Xoay mảng (Rotating an Array)
**Xoay trái k vị trí:**
```java
public class ArrayManipulation {
    public static void rotateLeft(int[] array, int k) {
        int n = array.length;
        k = k % n;
        reverse(array, 0, k - 1);
        reverse(array, k, n - 1);
        reverse(array, 0, n - 1);
    }

    private static void reverse(int[] array, int start, int end) {
        while (start < end) {
            int temp = array[start];
            array[start] = array[end];
            array[end] = temp;
            start++;
            end--;
        }
    }
}
```
Xoay phải k vị trí:
```java
public class ArrayManipulation {
    public static void rotateRight(int[] array, int k) {
        int n = array.length;
        k = k % n;
        reverse(array, 0, n - 1);
        reverse(array, 0, k - 1);
        reverse(array, k, n - 1);
    }
}
```
### 3. Tách và hợp mảng (Splitting and Merging Arrays)

**Tách mảng:**
```java
public class ArrayManipulation {
    public static int[][] splitArray(int[] array, int index) {
        int[] firstPart = java.util.Arrays.copyOfRange(array, 0, index);
        int[] secondPart = java.util.Arrays.copyOfRange(array, index, array.length);
        return new int[][]{firstPart, secondPart};
    }
}
```
Hợp mảng:
```java
public class ArrayManipulation {
    public static int[] mergeArrays(int[] array1, int[] array2) {
        int[] mergedArray = new int[array1.length + array2.length];
        System.arraycopy(array1, 0, mergedArray, 0, array1.length);
        System.arraycopy(array2, 0, mergedArray, array1.length, array2.length);
        return mergedArray;
    }
}
```
4. Loại bỏ phần tử trùng lặp (Removing Duplicates)
```java
import java.util.HashSet;

public class ArrayManipulation {
    public static int[] removeDuplicates(int[] array) {
        HashSet<Integer> set = new HashSet<>();
        for (int num : array) {
            set.add(num);
        }
        int[] result = new int[set.size()];
        int index = 0;
        for (int num : set) {
            result[index++] = num;
        }
        return result;
    }
}
```
### 5. Chèn và xóa phần tử (Insertion and Deletion)

**Chèn phần tử:**
```java
public class ArrayManipulation {
    public static int[] insertElement(int[] array, int element, int index) {
        int[] newArray = new int[array.length + 1];
        System.arraycopy(array, 0, newArray, 0, index);
        newArray[index] = element;
        System.arraycopy(array, index, newArray, index + 1, array.length - index);
        return newArray;
    }
}
```
Xóa phần tử:
```java
public class ArrayManipulation {
    public static int[] deleteElement(int[] array, int index) {
        if (index < 0 || index >= array.length) {
            throw new IndexOutOfBoundsException("Index out of bounds");
        }
        int[] newArray = new int[array.length - 1];
        System.arraycopy(array, 0, newArray, 0, index);
        System.arraycopy(array, index + 1, newArray, index, array.length - index - 1);
        return newArray;
    }
}
```
