
---
Mặc dù Golang không hỗ trợ lập trình hướng đối tượng (OOP) theo cách truyền thống như Java hay C++, nhưng nó vẫn cung cấp các công cụ mạnh mẽ để thực hiện các nguyên tắc OOP thông qua `struct`, `interface`, `method receiver`, và **interface embedding**.


### 1. **Hiểu về `struct` và `interface`**

#### 📌 `struct`

Golang không có `class`, nhưng `struct` đóng vai trò tương đương. Nó cho phép bạn tạo ra các đối tượng có thuộc tính (fields).

Ví dụ:
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name, p.Age) // Alice 30
}
```

#### 📌 `interface`

`interface` trong Go giúp định nghĩa hành vi mà một kiểu dữ liệu phải tuân theo mà không cần kế thừa.

```go
package main

import "fmt"

type Speaker interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

func main() {
    var s Speaker = Dog{Name: "Buddy"}
    fmt.Println(s.Speak()) // Woof!
}
```

### 2. **Method Receiver (`value receiver` vs `pointer receiver`)**

Trong Golang, bạn có thể định nghĩa method cho `struct` với hai loại receiver:

#### 📌 `Value receiver` (bản sao của giá trị)

Method sử dụng **value receiver** sẽ nhận một bản sao của struct, tức là thay đổi trên bản sao này không ảnh hưởng đến struct gốc.
```go
package main

import "fmt"

type Rectangle struct {
    Width, Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

func main() {
    r := Rectangle{Width: 10, Height: 5}
    fmt.Println(r.Area()) // 50
}
```
#### 📌 `Pointer receiver` (truy cập trực tiếp vào đối tượng gốc)

Dùng khi bạn muốn thay đổi giá trị của struct hoặc tránh copy struct lớn.

```go
package main

import "fmt"

type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}

func main() {
    c := Counter{Value: 0}
    c.Increment()
    fmt.Println(c.Value) // 1
}
```

**Khi nào dùng `pointer receiver`?**

- Khi method cần **thay đổi giá trị của struct**.
- Khi struct quá lớn, tránh copy tốn bộ nhớ.

### 3. **Kế thừa bằng Interface Embedding**

Go không hỗ trợ kế thừa như OOP truyền thống nhưng có **interface embedding**, cho phép một interface có thể chứa interface khác.

Ví dụ:
```go
package main

import "fmt"

type Animal interface {
    Speak() string
}

type Walker interface {
    Walk()
}

// Kết hợp nhiều interface vào một interface
type Dog interface {
    Animal
    Walker
}

type Bulldog struct {
    Name string
}

func (b Bulldog) Speak() string {
    return "Woof!"
}

func (b Bulldog) Walk() {
    fmt.Println(b.Name, "is walking")
}

func main() {
    var d Dog = Bulldog{Name: "Rocky"}
    fmt.Println(d.Speak()) // Woof!
    d.Walk() // Rocky is walking
}
```
🔹 **Lợi ích của Interface Embedding**:

- Dễ dàng mở rộng hành vi mà không cần kế thừa.
- Linh hoạt hơn trong thiết kế hệ thống.

### 4. **Polymorphism và Dependency Injection trong Go**

#### 📌 **Polymorphism** (Tính đa hình)

Polymorphism trong Go được thực hiện qua interface, cho phép một hàm xử lý nhiều kiểu khác nhau.

Ví dụ:
```go
package main

import "fmt"

type Notifier interface {
    Notify() string
}

type Email struct {
    Address string
}

func (e Email) Notify() string {
    return "Sending email to " + e.Address
}

type SMS struct {
    Number string
}

func (s SMS) Notify() string {
    return "Sending SMS to " + s.Number
}

// Hàm nhận một interface Notifier
func SendNotification(n Notifier) {
    fmt.Println(n.Notify())
}

func main() {
    email := Email{Address: "alice@example.com"}
    sms := SMS{Number: "+123456789"}

    SendNotification(email) // Sending email to alice@example.com
    SendNotification(sms)   // Sending SMS to +123456789
}
```

#### 📌 **Dependency Injection (DI) trong Go**

Golang không có DI container như Spring (Java), nhưng có thể thực hiện DI bằng cách truyền dependencies qua constructor hoặc setter.

Ví dụ:
```go
package main

import "fmt"

type Repository interface {
    GetData() string
}

type Database struct{}

func (d Database) GetData() string {
    return "Fetching data from database"
}

type Service struct {
    repo Repository
}

func NewService(r Repository) *Service {
    return &Service{repo: r}
}

func (s *Service) ProcessData() {
    fmt.Println(s.repo.GetData())
}

func main() {
    db := Database{}
    service := NewService(db) // Inject dependency
    service.ProcessData()     // Fetching data from database
}
```

🔹 **Lợi ích của DI trong Go**:

- Dễ dàng thay thế dependencies (ví dụ: thay đổi từ Database sang API).
- Code dễ kiểm thử hơn (mock dependencies trong unit test).

### 🔥 **Tóm Tắt**

✅ **Struct & Interface** giúp tổ chức dữ liệu và hành vi mà không cần kế thừa.  
✅ **Method Receiver** (`value receiver` vs `pointer receiver`) giúp kiểm soát cách struct tương tác với method.  
✅ **Interface Embedding** thay thế kế thừa trong Go, giúp mở rộng hành vi linh hoạt hơn.  
✅ **Polymorphism** được thực hiện qua interface, giúp code dễ mở rộng.  
✅ **Dependency Injection** giúp code dễ bảo trì và kiểm thử hơn.