# Using Cookie

> Step by Step

```javascript
npm install cookie-parser
```

```javascript
const express = require('express');
const cookieParser = require('cookie-parser');
const app = express();

app.use(cookieParser());

```

In Frontend:

```
import axios from 'axios';

axios.post('/api/login', { username: 'example', password: '123' }, { withCredentials: true })
  .then(response => {
    console.log(response);
  })
  .catch(error => {
    console.error(error);
  });

```

In Controller Backend:

```javascript
   // Generate token here
         const token = generateAccessToken(customer);
        res.cookie('accessToken', token, { httpOnly: true, maxAge: 3600000 }); 
```

