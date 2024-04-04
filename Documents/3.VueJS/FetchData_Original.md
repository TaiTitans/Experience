# Fetch Data Original

Để fetch data từ một API trên frontend website mà không sử dụng thư viện, bạn có thể sử dụng công cụ JavaScript như XMLHttpRequest hoặc fetch API có sẵn trong trình duyệt.

```javascript
function fetchDataFromAPI(url, callback) {
  var xhr = new XMLHttpRequest();

  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4 && xhr.status === 200) {
      var response = JSON.parse(xhr.responseText);
      callback(response);
    }
  };

  xhr.open("GET", url, true);
  xhr.send();
}

// Sử dụng fetchDataFromAPI để lấy dữ liệu từ API
fetchDataFromAPI("https://example.com/api/data", function(response) {
  // Xử lý dữ liệu nhận được từ API ở đây
  console.log(response);
});
```

Và đây là một ví dụ sử dụng `fetch` API:

```javascript
fetch("https://example.com/api/data")
  .then(function(response) {
    return response.json();
  })
  .then(function(data) {
    // Xử lý dữ liệu nhận được từ API ở đây
    console.log(data);
  })
  .catch(function(error) {
    // Xử lý lỗi nếu có
    console.log(error);
  });
```

--------

[JavaScript Fetch API: Everything You Need to Know | by Chris Ebube Roland | Medium](https://medium.com/@chrisebuberoland/javascript-fetch-api-everything-you-need-to-know-189c6ae49710)