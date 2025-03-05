
---
Dưới đây là lộ trình học MongoDB theo hướng **Senior Developer**, tập trung vào ứng dụng thực tế, tối ưu hóa hiệu suất, và sử dụng trong hệ thống phân tán.

---

## **1. Kiến thức Cơ Bản về MongoDB**

> ⏳ **Thời gian: 1-2 tuần**  
> 🎯 **Mục tiêu**: Hiểu mô hình NoSQL của MongoDB, các thao tác CRUD, và cách thiết lập một cơ sở dữ liệu đơn giản.

✅ **Làm quen với MongoDB**

- Cơ bản về NoSQL và sự khác biệt với RDBMS.
- Cấu trúc dữ liệu trong MongoDB (Document, Collection).
- Cài đặt MongoDB (Docker, Local, Atlas).

✅ **CRUD Operations**

- `insertOne()`, `insertMany()`, `find()`, `findOne()`, `updateOne()`, `updateMany()`, `deleteOne()`, `deleteMany()`.
- Cách sử dụng các bộ lọc (`$eq`, `$gt`, `$lt`, `$in`, `$regex`, `$exists`).
- Indexing cơ bản (`createIndex()`, `explain()`, `hint()`).

✅ **Aggregation Framework (Cơ bản)**

- `match`, `group`, `project`, `sort`, `limit`, `unwind`.

✅ **Thực hành**

- Xây dựng một API CRUD đơn giản bằng **Spring Boot + MongoDB**.

---

## **2. Thiết Kế Dữ Liệu & Indexing**

> ⏳ **Thời gian: 2-3 tuần**  
> 🎯 **Mục tiêu**: Hiểu cách thiết kế dữ liệu hiệu quả và tối ưu truy vấn.

✅ **Schema Design Patterns**

- Embedded vs. Referenced Documents: Khi nào nên sử dụng?
- One-to-Many, Many-to-Many trong MongoDB.
- **Data Modeling Patterns**: Bucket Pattern, Outlier Pattern, Extended Reference Pattern...

✅ **Tối ưu hóa Indexing**

- Các loại index: **Single Field**, **Compound Index**, **TTL Index**, **Text Index**, **Hashed Index**.
- **Explain Plan** & `db.collection.explain("executionStats").find(query)`.
- Index ảnh hưởng đến hiệu suất đọc/ghi như thế nào?

✅ **Thực hành**

- Thiết kế database cho một hệ thống thương mại điện tử (**giống Shopee, Lazada**).

---

## **3. Aggregation Framework (Nâng cao)**

> ⏳ **Thời gian: 2-3 tuần**  
> 🎯 **Mục tiêu**: Thành thạo `Aggregation Framework` để xử lý dữ liệu phức tạp.

✅ **Các stage nâng cao**

- `$lookup` (JOIN dữ liệu giữa collections).
- `$facet` (Xử lý nhiều pipeline song song).
- `$bucket`, `$bucketAuto` (Nhóm dữ liệu tự động).
- `$graphLookup` (Tạo truy vấn đệ quy, hữu ích khi làm hệ thống phân cấp).
- `$merge` (Lưu kết quả vào collection khác).

✅ **Thực hành**

- Xây dựng API thống kê doanh thu theo tháng/quý/năm từ dữ liệu đơn hàng.

---

## **4. Replica Set, Sharding & High Availability**

> ⏳ **Thời gian: 3-4 tuần**  
> 🎯 **Mục tiêu**: Hiểu cách triển khai MongoDB trong hệ thống phân tán.

✅ **Replica Set (Đảm bảo HA - High Availability)**

- **Cách hoạt động**: Primary, Secondary, Arbiter.
- Cách cấu hình Replica Set (`rs.initiate()`, `rs.add()`, `rs.status()`).
- Failover & Read Preference (`nearest`, `secondaryPreferred`).

✅ **Sharding (Chia nhỏ dữ liệu để scale)**

- **Cơ chế hoạt động**: Config Server, Shard Server, Mongos.
- Cách chọn **Sharding Key** phù hợp (`Hashed`, `Range`).
- Cấu hình Sharding trong MongoDB.

✅ **Thực hành**

- Cài đặt MongoDB Cluster với Replica Set + Sharding trên **Docker Compose**.

---

## **5. Caching & Performance Optimization**

> ⏳ **Thời gian: 2-3 tuần**  
> 🎯 **Mục tiêu**: Tối ưu hóa hiệu suất khi làm việc với MongoDB.

✅ **Tối ưu hóa đọc/ghi**

- Tuning `writeConcern`, `readConcern`, `journal`.
- **Bulk Operations** (`bulkWrite()` để insert/update nhanh hơn).
- **Schema Optimization** (Giảm duplication, tránh deep nesting).

✅ **Kết hợp với Redis để cache dữ liệu nóng**

- Khi nào nên cache với Redis?
- Kết hợp **MongoDB + Redis** để tối ưu performance.

✅ **Thực hành**

- **Xây dựng hệ thống cache sản phẩm trên Redis**, chỉ truy vấn MongoDB khi cần.

---

## **6. Bảo Mật & Best Practices**

> ⏳ **Thời gian: 2 tuần**  
> 🎯 **Mục tiêu**: Hiểu cách bảo mật dữ liệu và triển khai MongoDB an toàn.

✅ **Authentication & Authorization**

- Xác thực với **SCRAM, X.509, LDAP, Kerberos**.
- Phân quyền với **Role-Based Access Control (RBAC)**.

✅ **Mã hóa dữ liệu & Audit Log**

- **Field-Level Encryption** (Encrypt dữ liệu quan trọng).
- Bật **Audit Logging** để theo dõi các thao tác quan trọng.

✅ **Backup & Restore**

- `mongodump`, `mongorestore`, `oplog`.
- Replica Set Backup Strategy.

✅ **Thực hành**

- Cấu hình **MongoDB với TLS/SSL + RBAC** trên Docker.

---

## **7. Kết Hợp MongoDB trong Microservices**

> ⏳ **Thời gian: 2-3 tuần**  
> 🎯 **Mục tiêu**: Tích hợp MongoDB vào hệ thống Microservices.

✅ **Event-Driven với MongoDB Change Streams**

- **Stream dữ liệu real-time** khi có thay đổi (`watch()`).
- **Kết hợp với Kafka, RabbitMQ** để đẩy sự kiện.

✅ **Sử dụng MongoDB với CQRS & Event Sourcing**

- Khi nào sử dụng MongoDB làm Read Model?
- Kết hợp **MongoDB + PostgreSQL** trong kiến trúc CQRS.

✅ **Thực hành**

- **Tích hợp MongoDB vào một Microservice** với **Spring Boot**.

---

## **Tổng Kết & Hướng Đi Tiếp**

Sau khi hoàn thành lộ trình này, bạn sẽ có thể:  
✅ Thiết kế hệ thống MongoDB hiệu quả cho dữ liệu lớn.  
✅ Xây dựng API tối ưu với Aggregation Framework.  
✅ Deploy MongoDB với Replica Set, Sharding.  
✅ Tích hợp MongoDB trong hệ thống Microservices.

⏭ **Hướng đi tiếp**:

- Nghiên cứu **MongoDB Atlas** (Managed Service).
- Tìm hiểu **GraphQL + MongoDB**.
- Học **Time Series Data** với MongoDB (IoT, Log Analytics).