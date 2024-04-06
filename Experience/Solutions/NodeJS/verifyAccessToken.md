# Verify Access Token

```javascript
const jwt = require('jsonwebtoken')

function verifyAccessToken(req,res,next) {
    const accessToken = req.cookies.accessToken || req.headers['authorization']

    if(!accessToken){
        return res.status(401).json({error:'Access Token is missing'})
    }

    jwt.verify(accessToken,process.env.ACCESS_TOKEN_SECRET, (err, decoded)=>{
        if(err){
            return res.status(403).json({error:'Invalid access token'})
        }
        req.user = decoded
        next()
    })
}

module.exports = verifyAccessToken
```

=> Sử dụng ở router nhầm phân quyền cho user 