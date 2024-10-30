

---

Nuxt.JS có rất nhiều tính năng hữu ích, giúp bạn nhanh chóng xây dựng các ứng dụng web, có thể kể tới như:

- Automatic Code Splitting
- Hỗ trợ Vue hoàn hảo
- Static File Rendering
- Hỗ trợ phiên bản HTTP/2
- Hệ thống router và dữ liệu bất đồng bộ dễ sử dụng
- Hỗ trợ tính năng Hot reloading (rất hữu ích cho các bạn developer)
- .v.v…

---
## Khởi tạo một dự án với Nuxt.JS

`npx create-nuxt-app <project-name>`

**Lưu ý**: Bạn nhớ cài đặt công cụ `create-nuxt-app` trong máy tính đã nhé. Câu lệnh cài đặt: `npm install -g create-nuxt-app`. Và cả npx nữa nhé: `npm i npx`


Run app:
`npm run dev`

---

## Cấu trúc thư mục dự án Nuxt.JS

![](https://topdev.vn/blog/wp-content/uploads/2023/07/cau-truc-du-an-nuxt.jpg)

**Assets:** Chứa những tài nguyên phục vụ hiển thị trang web như ảnh, fonts chữ, hay CSS…

**Components:** Cũng giống như ứng dụng Vue thông thường, các component là các thành phần được tạo ra để bạn tái sử dụng trong ứng dụng như: Button, Input, Card, Dialog…

**Layouts:** thư mục layouts là nơi chứa các thành phần tạo nên layout của ứng dụng như layout dọc, layout ngang. Đây là nơi hợp lý nhất mà bạn có thể để các thành phần như Header, Footer, Theme…

**Middleware:** Middleware là nới bạn dùng để tạo các function mà chạy trước khi render trang.

**Pages:** Thư mục này chứa các view của ứng dụng cũng như định nghĩa routes cho ứng dụng luôn.

**Plugins:** Chứa các javascript plugin  mà bạn muốn chạy trước khi khởi tạo root vue.js application.

**Static:** Tương tự như thư mục assets nhưng mà nó cho phép truy cập trực tiếp, được map tự động với domain từ client mà không cần phải qua router hay biến môi trường. Ví dụ: /static/robots.txt sẽ được truy cập trực tiếp: http://localhost:3000/robots.txt

**Store:** Chứa các tệp của vuex, dùng quản lý state của ứng dụng. Vuex Store được cài đặt kèm với Nuxt nhưng mặc định thì lại bị disable. Muốn enable chúng, bạn chỉ cần tạo một tệp index.js trong thư mục store là được.

**nuxt.config.js:** Cấu hình ứng dụng Nuxt

**package.json:** Tương tự như dự án  react, vue, nodejs… là nơi cấu hình build, và chứa các dependencies và scripts

---

## Routing

Cách thức hoạt động của router trong Nuxt là nó tự động tạo cấu hình vue-router dựa trên cây các tệp .vue trong thư mục **pages**.

```
pages/
--|index.vue
--|product.vue
 --|index.vue
 --|one.vue
```

Với cấu trúc các tệp .vue như trên, khi generate ra route, bạn sẽ thu được kết quả như sau:

```javascript
router: {
  routes: [
    {
      name: 'index',
      path: '/',
      component: 'pages/index.vue'
    },
    {
      name: 'product',
      path: '/product',
      component: 'pages/product/index.vue'
    },
    {
      name: 'product-one',
      path: '/product/one',
      component: 'pages/product/one.vue'
    }
  ]
}
```

Bạn nên nhớ là việc tạo route là hoàn toàn tự động.


### Route lồng nhau


NuxtJS cho phép bạn tạo các route lồng nhau bằng cách sử dụng các route con của vue-router.

Để định nghĩa component cha của một route lồng nhau, bạn cần tạo một file vue với tên trùng với thư mục chứa các tệp vue con.

```
pages/
--| products/
-----| _id.vue
-----| index.vue
--| products.vue
```

Khi navigate giữa các page, Nuxt khuyến khích chúng ta sử dụng `nuxt-link` component thay vì dùng `router-link`.


---
## Triển khai ứng dụng Nuxt.JS

Nhìn vào file `package.json`, bạn sẽ thấy 4 câu lệnh sau:

![](https://vntalking.cdn.vccloud.vn/wp-content/uploads/2021/03/nuxt-deploy.png)


|   |   |
|---|---|
|Câu lệnh|Công dụng|
|dev|Chạy một development server trên localhost:8080, hỗ trợ hot-reloading|
|build|Build ứng dụng với Webpack và minify các tệp CSS, JS|
|start|Chạy server ở chế độ production (sau khi chạy nuxt build)|
|generate|Build ứng dụng và generate tất cả cá route tương ứng với các HTML files (sử dụng cho các static hosting)|
Như vậy, khi bạn muốn deploy ứng dụng Nuxt, bạn có thể chọn một trong 3 chế độ deploy:

Server-Side Rendering (Câu lệnh: npm run build)
Static Generated (Câu lệnh: npm run generate)
Single Page Applications


Riêng với chế độ Single Page Application, bạn cần thêm tham số này vào trong `nuxt.config.js`

```
export default {
  mode: 'spa'
}
```

Sau đó, quay trở lại package.json để thêm vào trong scripts.

```
"scripts": {
     "dev": "nuxt --spa",
     "build": "nuxt build --spa",
     "start": "nuxt start --spa",
     "generate": "nuxt generate --spa",
},
```

