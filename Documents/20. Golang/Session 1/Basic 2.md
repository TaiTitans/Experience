

---
### 📌 **2. Cấu trúc ngôn ngữ và cú pháp cơ bản**

#### 2.1. **Biến & Hằng số**: var, const

Trong Go, biến và hằng số là cách bạn lưu trữ dữ liệu.

- **Biến (var)**:
    - Khai báo biến bằng cú pháp: var tên_biến kiểu_dữ_liệu.
    - Go hỗ trợ khai báo ngắn gọn bằng := (chỉ dùng trong hàm).
    - Ví dụ:
```go
package main

import "fmt"

func main(){
var age int  = 25
name := "Nguyen"
var height float64

fmt.Println("Age:", age)
fmt.Println("Name:", name)
fmt.Println("Height:", height)
}
```

**Hằng số (const)**:

- Dùng để khai báo giá trị không thay đổi.
- Không thể dùng := với const.

```go
const PI = 3.14
const Greeting = "Hello"

func main(){
fmt.Println(PI, Greeting)
}
```

**Tips**: Go rất nghiêm ngặt, nếu khai báo biến mà không dùng, compiler sẽ báo lỗi. Điều này giúp code sạch hơn.

#### 2.2. **Kiểu dữ liệu**

Go có các kiểu dữ liệu cơ bản và nâng cao:

- **Cơ bản**:
    - int, int32, int64: Số nguyên (kích thước phụ thuộc kiến trúc máy hoặc chỉ định rõ).
    - float32, float64: Số thực.
    - string: Chuỗi ký tự.
    - bool: True/False.

- **Tập hợp**:

- array: Mảng cố định kích thước.
```go
var numbers [3]int = [3]int{1,2,3}
```

- slice: Mảng động (linh hoạt hơn array).
```go
slice := []int{1,2,3}
slice = append(slice, 4)
```

**Dữ liệu phức tạp**:

- map: Từ điển (key-value pair).
```go
scores := map[string]int{"Nam": 90, "Lan": 85}
fmt.Println(scores["Nam"])
```

- struct: Kiểu dữ liệu tự định nghĩa.
```go
type Person struct{
	Name string
	Age int
}
p := Person{Name: "An", Age: 30}
fmt.Println(p.name) // "An
```
- interface: Kiểu trừu tượng, dùng để định nghĩa hành vi.
```go
type Speaker interface{
	Speak() string
}
```

**Ví dụ tổng hợp**:
```go
func main() {
    slice := []string{"Go", "is", "fun"}
    person := struct{ Name string }{Name: "Grok"}
    fmt.Println(slice, person.Name)
}
```

#### 2.3. **Cấu trúc điều kiện & vòng lặp**

Go đơn giản hóa các cấu trúc điều khiển luồng.

- **if-else**:
    - Không cần dấu ngoặc cho điều kiện, nhưng bắt buộc có {}.
    - Hỗ trợ khai báo biến trong if.
    - Ví dụ:
```go
age := 20

if age >= 18 {
	fmt.Println("OK")
} else {
 fmt.Println("COOK")
}
//Khai bao bien trong if
if num := 10; num > 0{
	fmt.Println("Positive")
}
```

**switch-case**:

- Không cần break (Go tự động thoát sau mỗi case).
- Ví dụ:

```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Ngày đầu tuần")
case "Sunday":
    fmt.Println("Ngày cuối tuần")
default:
    fmt.Println("Ngày thường")
}
```


**for**:

- Go chỉ có for (không có while hay do-while).
- Dạng cơ bản:
```go
for i := 0;i < 5; i++{
fmt.Println(i)
}
```

Dạng vòng lặp vô hạn:
```go
for {
    fmt.Println("Vòng lặp mãi mãi")
    break // Thoát vòng lặp
}
```

Duyệt slice/map:
```go
name := []string{"A", "B", "C"}
for index, value := range names{
fmt.Println(index, value)
}
```


#### 2.4. **Hàm & Defer**: func, defer, panic, recover

Hàm là cốt lõi của Go, rất mạnh mẽ và linh hoạt.

- **Hàm cơ bản**:
    - Có thể trả về nhiều giá trị.
```go
func add(a ,b int) (int, string) {
	 return a + b, "Ket Qua"
}

func main() {
 sum, msg := add (3, 5)
 fmt.Println(sum, msg)
}
```
**defer**:

- Hoãn thực thi một lệnh đến khi hàm kết thúc.
```go
func main() {
    defer fmt.Println("Kết thúc")
    fmt.Println("Bắt đầu")
}
// Output: Bắt đầu -> Kết thúc
```


**panic & recover**:

- panic: Dừng chương trình khi gặp lỗi nghiêm trọng.
- recover: Bắt lỗi từ **panic** để chương trình không crash.
```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Đã phục hồi từ:", r)
        }
    }()
    panic("Lỗi xảy ra!")
}
```

#### 2.5. **Package & Import**

Go tổ chức code theo package để tái sử dụng và bảo trì tốt hơn.

- **Tạo package**:
    - Tạo thư mục mathutils và file math.go:
```go
package mathutils

func Add(a, b int) int {
    return a + b
}
```

Dùng trong main.go:
```go
package main

import (
    "fmt"
    "./mathutils" // Import package local
)

func main() {
    fmt.Println(mathutils.Add(2, 3)) // 5
}
```

**Với module**:

- Nếu dùng module, thêm module vào go.mod và import bằng tên module.


### 📌 **Bài tập thực hành**

1. Viết chương trình khai báo một struct Student với các trường Name (string) và Score (int), sau đó in thông tin của 3 học sinh.
2. Viết hàm max nhận 2 số int và trả về số lớn hơn, dùng if-else.
3. Dùng for để in bảng cửu chương của 5.
4. Tạo một map lưu tên và tuổi của 3 người, sau đó duyệt và in ra bằng range