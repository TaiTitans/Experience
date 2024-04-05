# Verify Password

```javascript
const bcrypt = require('bcrypt');
```

=> In Schema

```javascript
CustomerSchema.methods.comparePassword = async function(candidatePassword) {
    try {
        // Use bcrypt to compare the candidate password with the actual password stored in the database
        return await bcrypt.compare(candidatePassword, this.MatKhau);
    } catch (error) {
        throw new Error(error);
    }
};

```

