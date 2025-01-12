
---
**Aggregation Framework** là một công cụ mạnh mẽ trong MongoDB cho phép thực hiện các phép tính phức tạp và chuyển đổi dữ liệu. Nó được thiết kế để xử lý dữ liệu và trả về kết quả tổng hợp thông qua một hoặc nhiều giai đoạn (stages), mỗi giai đoạn áp dụng một thao tác nhất định trên dữ liệu.

### Hiểu Sâu Về Aggregation Framework Trong MongoDB

**Aggregation Framework** là một công cụ mạnh mẽ trong MongoDB cho phép thực hiện các phép tính phức tạp và chuyển đổi dữ liệu. Nó được thiết kế để xử lý dữ liệu và trả về kết quả tổng hợp thông qua một hoặc nhiều giai đoạn (stages), mỗi giai đoạn áp dụng một thao tác nhất định trên dữ liệu.

### 1. **Tổng Quan Về Aggregation Pipeline**

**Aggregation Pipeline** bao gồm một chuỗi các giai đoạn (stages) mà mỗi giai đoạn thực hiện một thao tác trên dữ liệu đầu vào và chuyển kết quả cho giai đoạn tiếp theo.

#### Các Stage Phổ Biến:

- **`$match`**: Lọc tài liệu (documents) để phù hợp với các điều kiện nhất định.
- **`$group`**: Nhóm tài liệu theo một hoặc nhiều trường và thực hiện các phép tổng hợp.
- **`$project`**: Thay đổi cấu trúc của tài liệu, có thể chọn hoặc tạo mới các trường.
- **`$sort`**: Sắp xếp tài liệu theo một hoặc nhiều trường.
- **`$limit`**: Giới hạn số lượng tài liệu trả về.
- **`$skip`**: Bỏ qua một số tài liệu.
- **`$unwind`**: Tách một mảng trong tài liệu thành nhiều tài liệu riêng biệt.
- **`$lookup`**: Thực hiện join với một collection khác.
- **`$addFields`**: Thêm các trường mới vào tài liệu.

### 2. **Ví Dụ Chi Tiết**

Giả sử bạn có collection `sales` với cấu trúc sau:
```
{
    "_id": 1,
    "item": "apple",
    "price": 10,
    "quantity": 5,
    "date": "2025-01-10"
}
```

2.1. **Tính Tổng Doanh Thu Theo Mỗi Ngày**

```
db.sales.aggregate([
    {
        $group: {
            _id: "$date",
            totalRevenue: { $sum: { $multiply: ["$price", "$quantity"] } },
            totalQuantity: { $sum: "$quantity" }
        }
    }
])
```
**Giải Thích:**

- `$group` tạo nhóm theo trường `date`.
- `totalRevenue` là tổng của phép nhân `price` và `quantity` cho mỗi tài liệu.
- `totalQuantity` là tổng của `quantity` cho mỗi ngày.
#### 2.2. **Lọc Và Tính Toán**

Tính tổng doanh thu chỉ cho các sản phẩm có giá trên 15.

```
db.sales.aggregate([
    { $match: { price: { $gt: 15 } } },
    {
        $group: {
            _id: "$item",
            totalRevenue: { $sum: { $multiply: ["$price", "$quantity"] } }
        }
    }
])
```

**Giải Thích:**

- `$match` lọc các tài liệu có `price` lớn hơn 15.
- `$group` tính tổng doanh thu cho từng `item`.

#### 2.3. **Thêm Trường Mới Sau Khi Tính Toán**

Thêm một trường mới để hiển thị doanh thu trung bình trên mỗi sản phẩm.
```
db.sales.aggregate([
    {
        $group: {
            _id: "$item",
            totalRevenue: { $sum: { $multiply: ["$price", "$quantity"] } },
            totalQuantity: { $sum: "$quantity" }
        }
    },
    {
        $addFields: {
            averageRevenuePerItem: { $divide: ["$totalRevenue", "$totalQuantity"] }
        }
    }
])

```

**Giải Thích:**

- `$addFields` thêm trường `averageRevenuePerItem`, tính bằng cách chia `totalRevenue` cho `totalQuantity`.

### 3. **Kết Hợp Nhiều Giai Đoạn**

Một pipeline có thể kết hợp nhiều giai đoạn để đạt được kết quả mong muốn.

#### Ví Dụ: Lọc, Nhóm, Và Sắp Xếp

```
db.sales.aggregate([
    { $match: { price: { $gt: 10 } } },
    {
        $group: {
            _id: "$item",
            totalRevenue: { $sum: { $multiply: ["$price", "$quantity"] } }
        }
    },
    { $sort: { totalRevenue: -1 } }
])

```

