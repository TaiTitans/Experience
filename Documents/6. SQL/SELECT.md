
---

# MySQL Thực Thi Lệnh SELECT Như Thế Nào?

## 1. Kiến Trúc Tổng Quan của MySQL

Hãy bắt đầu với sơ đồ tổng quan quá trình thực thi câu lệnh của MySQL trước, để chúng ta có cái nhìn toàn cảnh.

![](https://images.viblo.asia/93ec8ba7-a6f2-48ea-8ecb-14a263bf4289.png)

Kiến trúc của MySQL được chia thành hai layer chính là **Server Layer** và **Storage Engine Layer**.

**Server Layer: Đảm nhiệm vai trò tạo kết nối, phân tích và thực thi câu lệnh SQL.** Hầu hết các module chức năng cốt lõi của MySQL được triển khai ở tầng này. Bao gồm các module như connection manager, query cache, parser, preprocessor, optimizer, executor, ... Ngoài ra, tất cả các built-in functions (date, time, math,… ) và các chức năng khác chẳng hạn như stored procedures, triggers, views, … cũng được triển khai ở Server Layer.

**Storage Engine Layer: Đảm nhiệm vai trò lưu trữ và trích xuất dữ liệu**. MySQL hỗ trợ nhiều storage engine khác nhau như InnoDB, MyISAM, Memory, … Mỗi loại storage engine có đặc tính, tính năng khác nhau và performance khác nhau, phù hợp với các trường hợp khác nhau. Nhưng chúng sử dụng chung Server Layer. Storage engine được sử dụng phổ biến nhất hiện nay là InnoDB, nó cũng là storage engine mặc định kể từ MySQL 5.5.


Trước khi câu lệnh truy vấn SQL được thực thi, MySQL sẽ phân tích cú pháp câu lệnh SQL và công việc này được thực hiện bởi các module **Parser**.

Phần Parser sẽ làm hai việc.

Thứ nhất, **Lexical Scanner** sẽ xác định từ khóa dựa trên chuỗi bạn nhập vào. Ví dụ câu lệnh SQL `select Name from Member` sẽ nhận được 4 token sau khi phân tích. Trong đó có 2 keyword là `select` và `from` và 2 non-keyword là `Name` và `Member`.

Từ kết quả của Lexical Scanner, sau đó **Grammatical Checker** sẽ đánh giá xem câu lệnh SQL bạn nhập có thõa mãn cú pháp MySQL hay không. Nếu không có vấn đề gì, Grammatical Checker sẽ xây dựng SQL syntax tree. Parse Tree chứa các thông tin như loại SQL, tên bảng, ... và các module sau sẽ sử dụng những thông tin này.


Nếu chúng ta nhập câu lệnh SQL sai cú pháp, Parser sẽ báo lỗi.

```SQL
mysql> select * form Member;
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'form Member at line 1

```

Nhưng lưu ý rằng Parser chỉ chịu trách nhiệm kiểm tra cú pháp và xây dựng syntax tree, chứ không tra cứu bảng hoặc sự tồn tại của các trường.

Sau khi đi qua Parser, Optimizer sẽ tiếp nhận Parse Tree (SQL Query) và quá trình được chia thành 2 giai đoạn sau:

- Giai đoạn tiền xử lý (Preprocessor).
- Giai đoạn tối ưu hóa (Optimizer).


### Giai đoạn tiền xử lý (Preprocessor)

Trong giai đoạn tiền xử lý, Preprocessor sẽ thực hiện các công việc sau:

- **Kiểm tra xem bảng hoặc trường trong câu lệnh truy vấn SQL có tồn tại hay không**.
- **Convert từ dấu `*` thành tất cả cột của bảng tương ứng**.
- **Nếu một bảng không tồn tại, Preprocessor sẽ báo lỗi**

### Giai đoạn tối ưu hóa (Optimizer)

Sau giai đoạn tiền xử lý, **Optimizer sẽ xác định chiến lược thực thi (execution plan).**

Khi có nhiều index trong một bảng, Optimizer sẽ tính toán ra các chiến lược thực thi tương ứng với mỗi index. Ngoài thông tin về index, Optimizer phải dựa vào nhiều thông tin khác như các thông số dữ liệu (statistics) của bảng, cấu hình hệ thống, … để tính ra một chiến lược thực thi. Cuối cùng, **Optimizer sẽ chọn ra một execution plan có chi phi thấp nhất**.

Để biết Optimizer chọn index nào để truy vấn, ta có thể sử dụng từ khóa `explain` và xem giá trị của cột key để biết Optimizer chọn index nào.

*Nếu giá trị cột key là null, nghĩa là đang không có index nào được sử dụng. Toàn bộ bảng sẽ được scan (type = ALL), truy vấn kiểu này kém hiệu quả nhất.*

## Query Executor

Sau giai đoạn tối ưu hóa và xác định được chiến lược thực thi, đã đến lúc Executor thực hiện công việc của mình. Trong quá trình thực thi, Executor tương tác với storage engine.  
Query sẽ được thực thi theo 1 trong 3 cách sau:

- Truy vấn duyệt toàn bộ bảng
- Truy vấn sử dụng primary index
- Truy vấn sử dụng composite index

