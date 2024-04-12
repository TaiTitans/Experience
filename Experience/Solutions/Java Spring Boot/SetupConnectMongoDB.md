# Connect to MongoDB

1. Adding dependency to build.gradle:

   1. ```java
      dependencies {
          implementation 'org.springframework.boot:spring-boot-starter-data-mongodb'
      }
      ```

2. setting uri in application.properties

   1. ```java
      spring.data.mongodb.uri=mongodb://username:password@host:port/database
      spring.data.mongodb.database=your-database-name
      ```

3. make class in config:

   1. ```
      import org.springframework.context.annotation.Configuration;
      import org.springframework.data.mongodb.config.AbstractMongoClientConfiguration;
      import org.springframework.data.mongodb.repository.config.EnableMongoRepositories;
      
      @Configuration
      @EnableMongoRepositories(basePackages = "your.package.name.repository")
      public class MongoConfig extends AbstractMongoClientConfiguration {
      
          @Override
          protected String getDatabaseName() {
              return "your-database-name";
          }
      
          // Thêm các cấu hình khác tùy theo nhu cầu
      }
      ```

4. make other respository interface 

   1. ```java
      import org.springframework.data.mongodb.repository.MongoRepository;
      import org.springframework.stereotype.Repository;
      import your.package.name.model.YourEntity;
      
      @Repository
      public interface YourEntityRepository extends MongoRepository<YourEntity, String> {
          // Thêm các phương thức tùy theo nhu cầu
      }
      ```

5. use respository in service layer

   1. ```
      import org.springframework.stereotype.Service;
      import your.package.name.model.YourEntity;
      import your.package.name.repository.YourEntityRepository;
      
      @Service
      public class YourEntityService {
          private final YourEntityRepository repository;
      
          public YourEntityService(YourEntityRepository repository) {
              this.repository = repository;
          }
      
          public void createEntity(YourEntity entity) {
              repository.save(entity);
          }
      
          // Thêm các phương thức tùy theo nhu cầu
      }
      ```