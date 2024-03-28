# Install flexible automatic support in retrieving data from the backend

## Install from NPM
```
npm install http-proxy-middleware
```

## Create proxy.js in src
```
const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use('/api', createProxyMiddleware({ target: 'http://localhost:3000/', changeOrigin: true }));
};
```
## Active in vue.config.js
```
  devServer: {
    before: require('./src/proxy.js')
  }
```
