
---
### **Divide and Conquer** (Chia để trị)

**Divide and Conquer** là một kỹ thuật giải quyết vấn đề trong khoa học máy tính và toán học. Ý tưởng chính là chia nhỏ bài toán lớn thành các bài toán con nhỏ hơn, giải từng bài toán con, và sau đó kết hợp kết quả để giải quyết bài toán ban đầu.

---

### **Các bước chính của Divide and Conquer**

1. **Divide** (Chia nhỏ):
    
    - Chia bài toán lớn thành nhiều bài toán con nhỏ hơn. Các bài toán con này thường là các phần độc lập hoặc tương tự với bài toán gốc nhưng có kích thước nhỏ hơn.
        
2. **Conquer** (Giải quyết):
    
    - Giải từng bài toán con nhỏ. Nếu bài toán con đã đủ nhỏ, giải trực tiếp (trường hợp cơ sở). Nếu không, tiếp tục chia nhỏ cho đến khi có thể giải.
        
3. **Combine** (Kết hợp):
    
    - Gộp kết quả của các bài toán con để giải bài toán lớn ban đầu.

### **Ví dụ minh họa**

#### 1. **Tìm phần tử lớn nhất trong mảng**

- **Divide**: Chia mảng thành hai nửa.
    
- **Conquer**: Tìm phần tử lớn nhất trong từng nửa (gọi đệ quy).
    
- **Combine**: So sánh hai phần tử lớn nhất từ hai nửa và trả về phần tử lớn nhất.
    

**Pseudo Code**:
```java
int findMax(int[] arr, int left, int right) {
    if (left == right) {
        return arr[left]; // Trường hợp cơ sở
    }
    int mid = (left + right) / 2;
    int leftMax = findMax(arr, left, mid); // Tìm max bên trái
    int rightMax = findMax(arr, mid + 1, right); // Tìm max bên phải
    return Math.max(leftMax, rightMax); // Kết hợp kết quả
}
```

#### 2. **Merge Sort** (Sắp xếp trộn)

- **Divide**: Chia mảng thành hai nửa.
    
- **Conquer**: Sắp xếp từng nửa (gọi đệ quy).
    
- **Combine**: Trộn hai nửa đã được sắp xếp lại với nhau.
    

**Pseudo Code**:
```java
void mergeSort(int[] arr, int left, int right) {
    if (left < right) {
        int mid = (left + right) / 2;
        mergeSort(arr, left, mid); // Sắp xếp nửa trái
        mergeSort(arr, mid + 1, right); // Sắp xếp nửa phải
        merge(arr, left, mid, right); // Trộn hai nửa
    }
}

void merge(int[] arr, int left, int mid, int right) {
    // Code trộn mảng con
}
```

#### 3. **Binary Search**

- **Divide**: Chia khoảng tìm kiếm thành hai nửa.
    
- **Conquer**: Tìm kiếm ở nửa phù hợp (dựa vào giá trị cần tìm so với phần tử giữa).
    
- **Combine**: Không cần (bài toán này không cần kết hợp).
    

**Pseudo Code**:
```java
int binarySearch(int[] arr, int left, int right, int target) {
    if (left > right) {
        return -1; // Không tìm thấy
    }
    int mid = (left + right) / 2;
    if (arr[mid] == target) {
        return mid; // Tìm thấy
    }
    if (arr[mid] > target) {
        return binarySearch(arr, left, mid - 1, target); // Tìm ở nửa trái
    } else {
        return binarySearch(arr, mid + 1, right, target); // Tìm ở nửa phải
    }
}
```
### **Ứng dụng thực tế của Divide and Conquer**

1. **Thuật toán sắp xếp**:
    
    - Merge Sort
        
    - Quick Sort (tuy là Divide and Conquer nhưng cách kết hợp hơi khác).
        
2. **Tìm kiếm**:
    
    - Binary Search.
        
3. **Xử lý dữ liệu lớn**:
    
    - Dùng để giải các bài toán như nhân ma trận lớn, tính toán FFT (Fast Fourier Transform).
        
4. **Tối ưu hóa các bài toán đệ quy**:
    
    - Tìm lũy thừa (`a^b`) bằng phương pháp chia đôi.
        
5. **Đồ họa máy tính**:
    
    - Tìm điểm gần nhau nhất trong một tập hợp điểm.
        

---

### **Lợi ích**

- Giúp giảm độ phức tạp của bài toán, thường đạt độ phức tạp thấp hơn so với cách tiếp cận tuyến tính.
    
- Dễ triển khai với đệ quy.
    

---

### **Nhược điểm**

- Sử dụng đệ quy có thể gây tốn bộ nhớ nếu không tối ưu.
    
- Việc gộp kết quả có thể phức tạp với một số bài toán.
