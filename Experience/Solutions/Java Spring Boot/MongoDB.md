
---
```
<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-data-mongodb</artifactId>
    </dependency>
</dependencies>
```

```
spring:
  data:
    mongodb:
      uri: mongodb://localhost:27017/your_database_name

```

```java
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "users")
public class User {

    @Id
    private String id;
    private String name;
    private String email;

    // getters and setters
}

```

```java
import org.springframework.data.mongodb.repository.MongoRepository;

public interface UserRepository extends MongoRepository<User, String> {
    // Custom query methods can be defined here
    User findByEmail(String email);
}

```

```java
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserService {

    @Autowired
    private UserRepository userRepository;

    public List<User> getAllUsers() {
        return userRepository.findAll();
    }

    public User getUserByEmail(String email) {
        return userRepository.findByEmail(email);
    }

    public User saveUser(User user) {
        return userRepository.save(user);
    }

    public void deleteUser(String id) {
        userRepository.deleteById(id);
    }
}

```

Custom: 
```java
import org.springframework.data.mongodb.repository.Query;

public interface UserRepository extends MongoRepository<User, String> {

    @Query("{ 'name' : ?0 }")
    List<User> findUsersByName(String name);

    @Query(value = "{ 'email' : ?0 }", fields = "{ 'name' : 1, 'email' : 1 }")
    User findUserWithSelectiveFields(String email);
}

```

**Aggregation**

```java
import org.springframework.data.mongodb.core.aggregation.Aggregation;
import org.springframework.data.mongodb.core.aggregation.AggregationResults;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Service;

@Service
public class UserService {

    @Autowired
    private MongoTemplate mongoTemplate;

    public List<User> getUsersWithAggregation() {
        Aggregation aggregation = Aggregation.newAggregation(
            Aggregation.match(Criteria.where("name").is("John")),
            Aggregation.group("email").count().as("total")
        );

        AggregationResults<User> results = mongoTemplate.aggregate(aggregation, "users", User.class);
        return results.getMappedResults();
    }
}
```

