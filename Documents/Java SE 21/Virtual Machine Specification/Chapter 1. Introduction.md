
---
Chương này thuộc **Java Virtual Machine Specification (JVMS)**, tức là đặc tả về **Máy ảo Java (JVM)**. Đây là tài liệu mô tả cách JVM hoạt động, bao gồm cách nó thực thi mã Java, quản lý bộ nhớ, xử lý luồng, và tương tác với hệ điều hành.

## **1️⃣ Tổng quan về JVM**

**Java Virtual Machine (JVM)** là môi trường chạy mã Java, giúp chương trình Java có thể thực thi trên nhiều hệ điều hành khác nhau mà không cần biên dịch lại. JVM chịu trách nhiệm:

- **Biên dịch bytecode** thành mã máy (_machine code_) thông qua JIT (_Just-In-Time Compiler_).
- **Quản lý bộ nhớ** (heap, stack, method area, v.v.).
- **Xử lý luồng (threads)**.
- **Thực thi lệnh trong bộ mã lệnh JVM**.

## **2️⃣ Cấu trúc của tài liệu JVMS**

Tài liệu này bao gồm các chương chính:

1. **Introduction (Giới thiệu về JVM)**
2. **The Structure of the Java Virtual Machine (Cấu trúc JVM)**
3. **Compiling for the Java Virtual Machine (Quá trình biên dịch bytecode)**
4. **The class File Format (Định dạng tệp `.class`)**
5. **Loading, Linking, and Initialization (Quá trình tải và khởi tạo lớp)**
6. **Instruction Set (Tập lệnh JVM)**
7. **Stacks, Frames, and Heap (Quản lý bộ nhớ)**
8. **Threads and Locks (Luồng và khóa đồng bộ)**

