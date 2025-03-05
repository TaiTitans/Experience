

---

SOAP (Simple Object Access Protocol) là một giao thức được sử dụng để trao đổi thông tin trong các dịch vụ web. Nó dựa trên XML và thường được sử dụng trong các hệ thống yêu cầu giao tiếp đồng nhất và bảo mật cao. Dưới đây là tài liệu hướng dẫn cơ bản về SOAP API từ góc nhìn của một Senior Java Backend Developer.


## **1. Tổng quan về SOAP API**

### **Đặc điểm chính của SOAP API:**

- **Protocol-based**: SOAP là giao thức, khác với REST là phong cách kiến trúc.
- **Dựa trên XML**: Mọi thông điệp SOAP được mã hóa dưới dạng XML.
- **Ràng buộc chặt chẽ**: Yêu cầu định nghĩa rõ ràng qua WSDL (Web Services Description Language).
- **Hỗ trợ bảo mật**: WS-Security cung cấp các cơ chế mã hóa, xác thực và chữ ký số.
- **Hỗ trợ trạng thái**: Dễ dàng duy trì trạng thái thông qua các header SOAP.

## **2. Cấu trúc thông điệp SOAP**

### **Cấu trúc chính:**

```xml
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <Header>
        <!-- Thông tin bảo mật hoặc metadata -->
    </Header>
    <Body>
        <!-- Nội dung chính của yêu cầu hoặc phản hồi -->
    </Body>
</Envelope>
```

Ví dụ yêu cầu SOAP:
```xml
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <Body>
        <GetProductDetails xmlns="http://example.com/product">
            <ProductId>123</ProductId>
        </GetProductDetails>
    </Body>
</Envelope>
```


Ví dụ phản hồi SOAP:

```xml
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <Body>
        <GetProductDetailsResponse xmlns="http://example.com/product">
            <ProductName>Soap</ProductName>
            <ProductPrice>10.99</ProductPrice>
        </GetProductDetailsResponse>
    </Body>
</Envelope>
```

## **3. Tạo SOAP API với Java (Spring-WS)**

### **3.1. Cài đặt Maven**

Thêm dependency vào `pom.xml`:

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web-services</artifactId>
</dependency>
```

### **3.2. Định nghĩa XSD**

Tạo file `product-details.xsd` để định nghĩa cấu trúc dữ liệu:

```xml
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           targetNamespace="http://example.com/product"
           xmlns="http://example.com/product"
           elementFormDefault="qualified">

    <xs:element name="GetProductDetailsRequest">
        <xs:complexType>
            <xs:sequence>
                <xs:element name="ProductId" type="xs:string"/>
            </xs:sequence>
        </xs:complexType>
    </xs:element>

    <xs:element name="GetProductDetailsResponse">
        <xs:complexType>
            <xs:sequence>
                <xs:element name="ProductName" type="xs:string"/>
                <xs:element name="ProductPrice" type="xs:decimal"/>
            </xs:sequence>
        </xs:complexType>
    </xs:element>
</xs:schema>
```

3.3. Tạo file JAXB Object

Sử dụng plugin để sinh các class Java từ XSD:

```xml
<plugin>
    <groupId>org.codehaus.mojo</groupId>
    <artifactId>jaxb2-maven-plugin</artifactId>
    <version>2.5.0</version>
    <executions>
        <execution>
            <goals>
                <goal>xjc</goal>
            </goals>
        </execution>
    </executions>
    <configuration>
        <schemaDirectory>${project.basedir}/src/main/resources</schemaDirectory>
        <outputDirectory>${project.basedir}/target/generated-sources/xjc</outputDirectory>
    </configuration>
</plugin>

```

### **3.4. Tạo Endpoint**

Tạo endpoint xử lý yêu cầu SOAP:
```java
@Endpoint
public class ProductEndpoint {

    private static final String NAMESPACE_URI = "http://example.com/product";

    @PayloadRoot(namespace = NAMESPACE_URI, localPart = "GetProductDetailsRequest")
    @ResponsePayload
    public GetProductDetailsResponse getProductDetails(@RequestPayload GetProductDetailsRequest request) {
        GetProductDetailsResponse response = new GetProductDetailsResponse();
        response.setProductName("Sample Product");
        response.setProductPrice(BigDecimal.valueOf(12.99));
        return response;
    }
}
```

### **3.5. Cấu hình SOAP Web Service**

Tạo lớp cấu hình:
```java
@EnableWs
@Configuration
public class WebServiceConfig extends WsConfigurerAdapter {

    @Bean
    public ServletRegistrationBean<MessageDispatcherServlet> messageDispatcherServlet(ApplicationContext context) {
        MessageDispatcherServlet servlet = new MessageDispatcherServlet();
        servlet.setApplicationContext(context);
        servlet.setTransformWsdlLocations(true);
        return new ServletRegistrationBean<>(servlet, "/ws/*");
    }

    @Bean(name = "products")
    public DefaultWsdl11Definition defaultWsdl11Definition(XsdSchema productsSchema) {
        DefaultWsdl11Definition definition = new DefaultWsdl11Definition();
        definition.setPortTypeName("ProductsPort");
        definition.setLocationUri("/ws");
        definition.setTargetNamespace("http://example.com/product");
        definition.setSchema(productsSchema);
        return definition;
    }

    @Bean
    public XsdSchema productsSchema() {
        return new SimpleXsdSchema(new ClassPathResource("product-details.xsd"));
    }
}
```

## **4. Test SOAP API**

### **4.1. Tool Test**

- Sử dụng **SoapUI** hoặc Postman (nếu hỗ trợ SOAP).

### **4.2. URL**

- Truy cập `http://localhost:8080/ws/products.wsdl` để xem file WSDL.
- Gửi request SOAP đến `http://localhost:8080/ws`.



## **5. Ưu và nhược điểm của SOAP**

### **Ưu điểm:**

- Tích hợp bảo mật tốt.
- Hỗ trợ giao thức chuẩn như HTTP, SMTP.
- Định nghĩa rõ ràng qua WSDL.

### **Nhược điểm:**

- Nặng hơn REST do sử dụng XML.
- Khó học hơn so với REST.
- Ít phổ biến trong các ứng dụng hiện đại.