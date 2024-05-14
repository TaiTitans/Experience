
---
	Spring Boot Annotations is a form of metadata that provides data about a program. In other words, annotations are used to provide **supplemental** information about a program. It is not a part of the application that we develop. It does not have a direct effect on the operation of the code they annotate. It does not change the action of the compiled program.

**@Autowired**: Spring provides annotation-based auto-wiring by providing @Autowired annotation. It is used to autowire spring bean on setter methods, instance variable, and constructor. When we use @Autowired annotation, the spring container auto-wires the bean by mathcing data-type.

<details> <summary>Example</summary> 

```
1. @Component  
2. public class Customer  
3. {  
4. private Person person;  
5. @Autowired  
6. public Customer(Person person)   
7. {   
8. this.person=person;  
9. }  
10. }
```

</details>

**@Configuration**: It is a class-level annotation. The class annotated with @Configuration used by Spring Containers as a source of bean definitions.

**@Component:** It is a class-level annotation. It is used to mark a Java class as a bean. A Java class annotated with **@Component** is found during the classpath.


**@Controller:** The @Controller is a class-level annotation. It is a specialization of **@Component**. It marks a class as a web request handler. It is often used to serve web pages. By default, it returns a string that indicates which route to redirect. It is mostly used with **@RequestMapping** annotation.

<details>
<summary>Example</summary>

```
1. @Controller  
2. @RequestMapping("books")  
3. public class BooksController   
4. {  
5. @RequestMapping(value = "/{name}", method = RequestMethod.GET)  
6. public Employee getBooksByName()   
7. {  
8. return booksTemplate;  
9. }  
10. }
```

</details>

**@Service:** It is also used at class level. It tells the Spring that class contains the business logic.

<details>
<summary>Example</summary>
```
1. package com.javatpoint;  
2. @Service  
3. public class TestService  
4. {  
5. public void service1()  
6. {  
7. //business code  
8. }  
9. }
```
</details>

**@Repository:** It is a class-level annotation. The repository is a **DAOs** (Data Access Object) that access the database directly. The repository does all the operations related to the database.

<details>
<summary>Example</summary>
```
1. package com.javatpoint;  
2. @Repository   
3. public class TestRepository  
4. {  
5. public void delete()  
6. {     
7. //persistence code  
8. }  
9. }
```
</details>



And more Annotations : [Spring Boot Annotations - javatpoint](https://www.javatpoint.com/spring-boot-annotations)