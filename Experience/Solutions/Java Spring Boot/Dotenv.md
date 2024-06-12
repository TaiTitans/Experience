
---
`build.gradle`

```Java
implementation 'io.github.cdimascio:dotenv-java:3.1.0'
```

Make .env 

```
DB_USERNAME=myusername
DB_PASSWORD=mypassword

```

Update `application.properties`
```java
spring.datasource.url=jdbc:mysql://localhost:3306/
spring.datasource.username=${DB_USERNAME}
spring.datasource.password=${DB_PASSWORD}

```

Load environment:
`main`

```Java
package com.jewelrymanagement;

import io.github.cdimascio.dotenv.Dotenv;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class JewelryManagementApplication {

    public static void main(String[] args) {
        // Load .env file
        Dotenv dotenv = Dotenv.load();

        // Set environment variables
        System.setProperty("DB_USERNAME", dotenv.get("DB_USERNAME"));
        System.setProperty("DB_PASSWORD", dotenv.get("DB_PASSWORD"));

        SpringApplication.run(JewelryManagementApplication.class, args);
    }
}

```