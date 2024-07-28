Factory - Builder - Singleton - Observer - Iterator  - Strategy - Adapter - Facade


---


	Design Pattern là một giải pháp chung để giải quyết các vấn đề phổ biến khi thiết kế phần mềm trong lập trình hướng đối tượng OOP.


## Design Patterns hỗ trợ developers như thế nào

### 1. Tăng tốc độ phát triển phần mềm
### 2. Code tường minh, dễ dàng team work
### 3. Tái sử dụng code
### 4. Hạn chế lỗi tiềm ẩn, dễ dàng nâng cấp

## Phân loại Design Patterns

Hệ thống các mẫu design pattern hiện nay rất nhiều, nhưng thường tóm gọn bằng 23 mẫu được định nghĩa trong cuốn “Design patterns Elements of Reusable Object Oriented Software”. Hệ thống các mẫu design pattern được chia thành 3 nhóm, được phân loại theo mục đích sử dụng:

- Nhóm Creational
- Nhóm Structural
- Nhóm Behavioral
### 1. Creational Pattern (Nhóm khởi tạo)

Nhóm Creational Pattern gồm 5 mẫu:

- [Singleton](https://viblo.asia/p/signleton-desgin-pattern-tro-thu-dac-luc-cua-developers-Qbq5QBkJKD8)
- [Factory Method](https://viblo.asia/p/factory-method-design-pattern-tro-thu-dac-luc-cua-developers-924lJBLYlPM)
- [Abstract Factory](https://viblo.asia/p/abstract-factory-design-pattern-tro-thu-dac-luc-cua-developers-maGK7B4M5j2)
- [Builder](https://viblo.asia/p/builder-design-pattern-tro-thu-dac-luc-cua-developers-bWrZnowwlxw)
- [Prototype](https://viblo.asia/p/prototype-design-pattern-tro-thu-dac-luc-cua-developers-GrLZDBQO5k0)

Các patterns loại này cung cấp giải pháp để tạo ra các đối tượng và che giấu được logic của việc tạo ra nó thay vì tạo ra đối tượng theo cách trực tiếp (sử dụng từ khoá new). Điều này giúp chương trình trở nên mềm dẻo hơn trong việc quyết định đối tượng nào cần được tạo ra trong những tình huống khác nhau.

### 2. Structural Pattern (Nhóm cấu trúc)

Nhóm Structural Pattern gồm 7 mẫu:

- [Adapter](https://viblo.asia/p/adapter-design-pattern-tro-thu-dac-luc-cua-developers-Az45bqYQlxY)
- [Bridge](https://viblo.asia/p/bridge-design-pattern-tro-thu-dac-luc-cua-developers-gDVK2oG2ZLj)
- [Composite](https://viblo.asia/p/composite-design-pattern-tro-thu-dac-luc-cua-developers-Qbq5QBk3KD8)
- [Decorator](https://viblo.asia/p/decorator-design-pattern-tro-thu-dac-luc-cua-developers-1VgZvQ1OKAw)
- [Facade](https://viblo.asia/p/facade-design-pattern-tro-thu-dac-luc-cua-developers-924lJBLNlPM)
- [Flyweight](https://viblo.asia/p/flyweight-design-pattern-tro-thu-dac-luc-cua-developers-maGK7B4b5j2)
- [Proxy](https://viblo.asia/p/proxy-design-pattern-tro-thu-dac-luc-cua-developers-RQqKLB2bl7z)

Những patterns loại này liên quan tới class và các thành phần của đối tượng. Nó dùng để thiết lập, định nghĩa quan hệ giữa các đối tượng. Hệ thống càng lớn thì mẫu này càng đóng vai trò quan trọng. Ta có thể dựa vào class diagram để theo dõi mẫu này.


###  3. Behavioral Pattern (Nhóm hành vi)

Nhóm Behavioral Pattern gồm 11 mẫu:

- [Iterpreter](https://viblo.asia/p/interpreter-design-pattern-tro-thu-dac-luc-cua-developers-djeZ1d43KWz)
- [Template Method](https://viblo.asia/p/template-method-design-pattern-tro-thu-dac-luc-cua-developers-Az45bqYLlxY)
- [Chain of Responsibility](https://viblo.asia/p/chain-of-responsibility-design-pattern-tro-thu-dac-luc-cua-developers-yMnKMBNDZ7P)
- [Command](https://viblo.asia/p/command-design-pattern-tro-thu-dac-luc-cua-developers-4dbZNBqkZYM)
- [Iterator](https://viblo.asia/p/iterator-design-pattern-tro-thu-dac-luc-cua-developers-jvElaNwY5kw)
- [Mediator](https://viblo.asia/p/mediator-design-pattern-tro-thu-dac-luc-cua-developers-m68Z0jVj5kG)
- [Memento](https://viblo.asia/p/memento-design-pattern-tro-thu-dac-luc-cua-developers-gGJ59BzrKX2)
- [Observer](https://viblo.asia/p/observer-design-pattern-tro-thu-dac-luc-cua-developers-gAm5y7WAZdb)
- [State](https://viblo.asia/p/state-design-pattern-tro-thu-dac-luc-cua-developers-3P0lPB9PKox)
- [Strategy](https://viblo.asia/p/strategy-design-pattern-tro-thu-dac-luc-cua-developers-bJzKmdwP59N)
- [Visitor](https://viblo.asia/p/visitor-design-pattern-tro-thu-dac-luc-cua-developers-gDVK2oGeZLj)

Nhóm này liên quan đến các quan hệ hành vi để xử lí các chức năng giữa các đối tượng trong hệ thống. Đối với các mẫu thuộc nhóm này ta có thể dựa vào collaboration và sequence diagram để theo dõi.