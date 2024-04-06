# Generate Access Token

```javascript
const jwt = require('jsonwebtoken');

// generate access token

function generateAccessToken(username){
    const payload = {
        _id: Customer._id,
        MaDocGia: Customer.MaDocGia
    }
    const accessToken = jwt.sign(payload, process.env.ACCESS_TOKEN_SECRET, {expiresIn:'1d'})
    return accessToken
}

module.exports = generateAccessToken
```

=> Sử dụng ở Controller Đăng Nhập