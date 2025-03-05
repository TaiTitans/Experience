
---
## **1. Kiến thức cơ bản về PostgreSQL**

### 📌 **Cài đặt và cấu hình**

- Cài đặt PostgreSQL trên Windows, Linux, hoặc Docker.
- Hiểu về `postgresql.conf`, `pg_hba.conf`, và cách cấu hình cơ bản.
- Quản lý người dùng và quyền (`CREATE USER`, `GRANT`, `REVOKE`).

### 📌 **SQL cơ bản với PostgreSQL**

- `CREATE TABLE`, `INSERT`, `UPDATE`, `DELETE`, `SELECT`.
- Ràng buộc dữ liệu (`PRIMARY KEY`, `FOREIGN KEY`, `CHECK`, `UNIQUE`).
- Chỉ mục (`INDEX`), sử dụng chỉ mục để tối ưu hiệu suất.

---

## **2. Kiến thức nâng cao**

### 📌 **Cấu trúc dữ liệu nâng cao**

- **Partitioning Table**: `PARTITION BY RANGE/LIST/HASH`
- **Inheritance Tables**: Kế thừa bảng trong PostgreSQL.
- **JSON/JSONB**: Lưu trữ và truy vấn dữ liệu dạng JSON.
- **Array & Hstore**: Làm việc với kiểu dữ liệu mảng và key-value store.

### 📌 **Tối ưu hóa truy vấn (Query Optimization)**

- **EXPLAIN & EXPLAIN ANALYZE**: Phân tích kế hoạch thực thi.
- **Vacuum & Analyze**: Cách PostgreSQL dọn dẹp và tối ưu hóa.
- **Indexing nâng cao**:
    - B-tree, Hash, GIN, GiST, BRIN.
    - Covering Index, Partial Index, Expression Index.
- **CTE (WITH Queries) và Recursive Query**.
- **Parallel Query Execution**.

---

## **3. Quản lý giao dịch và đồng thời (Transaction & Concurrency)**

- **ACID và MVCC (Multi-Version Concurrency Control)**.
- **Isolation Levels (Read Committed, Repeatable Read, Serializable)**.
- **Deadlock detection & prevention**.
- **Locking Strategies**:
    - Row-level locks (`FOR UPDATE`, `FOR SHARE`).
    - Table-level locks (`LOCK TABLE`).
    - Advisory Locks (`pg_advisory_lock`).

---

## **4. Cấu hình, quản trị và High Availability**

### 📌 **Quản lý & Giám sát**

- **Logging & Monitoring**:
    
    - Cấu hình log (`log_statement`, `log_min_duration_statement`).
    - Sử dụng `pg_stat_statements`, `pg_stat_activity` để theo dõi truy vấn.
    - Kết hợp với Prometheus + Grafana để giám sát hiệu suất.
- **Backup & Restore**:
    
    - Logical Backup: `pg_dump`, `pg_restore`.
    - Physical Backup: `pg_basebackup`, WAL Archiving.

### 📌 **Replication & High Availability**

- **Streaming Replication** (Primary - Standby).
- **Logical Replication** (`pglogical`, `CREATE PUBLICATION/SUBSCRIPTION`).
- **Failover & Load Balancing** với Patroni, PgBouncer, HAProxy.

---

## **5. PostgreSQL trong hệ thống Microservices**

- **Kết nối từ Spring Boot, Node.js, Golang...**.
- **Connection Pooling với HikariCP, PgBouncer**.
- **Saga Pattern & Outbox Pattern trong giao dịch phân tán**.
- **Event-driven với PostgreSQL LISTEN/NOTIFY**.

---

## **6. Kiến thức chuyên sâu**

### 📌 **PostgreSQL Extensions**

- `PostGIS` (Xử lý dữ liệu địa lý).
- `TimescaleDB` (Làm việc với Time-Series Data).
- `Citus` (Phân tán dữ liệu trên nhiều node).

### 📌 **Tối ưu hiệu suất nâng cao**

- **Sharding** với `Citus` hoặc `Foreign Data Wrapper (FDW)`.
- **Partitioning chiến lược cho Big Data**.
- **Columnar Storage với `pg_columnar`**.

---

## **7. Thực hành & Dự án**

- **Tối ưu database trong hệ thống đang làm** (chạy `EXPLAIN ANALYZE`, tạo index hợp lý).
- **Xây dựng mô hình High Availability với Streaming Replication**.
- **Triển khai PostgreSQL trên Kubernetes với Patroni**.
- **Xây dựng Event Sourcing với LISTEN/NOTIFY**.

---

## **Tài nguyên học tập**

- 📚 **Sách**:
    
    - _PostgreSQL: Up & Running_ - Regina O. Obe, Leo S. Hsu.
    - _Mastering PostgreSQL in Application Development_ - Dimitri Fontaine.
    - _High Performance PostgreSQL_ - Gregory Smith.
- 🎓 **Khóa học**:
    
    - Udemy: _The Complete PostgreSQL Bootcamp_.
    - LinkedIn Learning: _PostgreSQL Essential Training_.
- 📝 **Docs chính thức**:
    
    - [https://www.postgresql.org/docs/](https://www.postgresql.org/docs/)