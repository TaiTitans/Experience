# Axios

Cài đặt Axios:

```javascript
npm install axios
```

##### GET:

```javascript
axios.get('https://api.example.com/data')
  .then(function (response) {
    // Xử lý dữ liệu nhận được từ API ở đây
    console.log(response.data);
  })
  .catch(function (error) {
    // Xử lý lỗi nếu có
    console.log(error);
  });
```

##### POST

```javascript
axios.post('https://api.example.com/data', {
    name: 'John Doe',
    email: 'johndoe@example.com'
  })
  .then(function (response) {
    // Xử lý phản hồi từ server ở đây
    console.log(response.data);
  })
  .catch(function (error) {
    // Xử lý lỗi nếu có
    console.log(error);
  });
```

##### Làm việc với form:

```javascript
const formData = new FormData();
formData.append('name', 'John Doe');
formData.append('email', 'johndoe@example.com');

axios.post('https://api.example.com/data', formData)
  .then(function (response) {
    // Xử lý phản hồi từ server ở đây
    console.log(response.data);
  })
  .catch(function (error) {
    // Xử lý lỗi nếu có
    console.log(error);
  });
```

##### Method + x-www-form-urlencoded

```javascript
const params = new URLSearchParams();
params.append('name', 'John Doe');
params.append('email', 'johndoe@example.com');

axios.post('https://api.example.com/data', params)
  .then(function (response) {
    // Xử lý phản hồi từ server ở đây
    console.log(response.data);
  })
  .catch(function (error) {
    // Xử lý lỗi nếu có
    console.log(error);
  });
```

[Bắt đầu | Tài liệu Axios (axios-http.com)](https://axios-http.com/vi/docs/intro)