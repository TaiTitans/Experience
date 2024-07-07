
---

```Java
@Service  
public class MailjetEmailService implements EmailSender {  
    private static final Logger logger = LoggerFactory.getLogger(MailjetEmailService.class);  
  
    private final MailjetClient mailjetClient;  
  
    public MailjetEmailService(@Value("${mailjet.api-key}") String apiKey,  
                               @Value("${mailjet.secret-key}") String secretKey) {  
        ClientOptions options = ClientOptions.builder()  
                .apiKey(apiKey)  
                .apiSecretKey(secretKey)  
                .build();  
  
        this.mailjetClient = new MailjetClient(options);  
  
    }  
    @Override  
    public void sendOTPEmail(String to, String otp) {  
        try {  
            TransactionalEmail message = TransactionalEmail  
                    .builder()  
                    .to(new SendContact(to))  
                    .from(new SendContact("talktowntechnology@gmail.com", "Talk Town Technology"))  
                    .subject("Your OTP Code")  
                    .htmlPart("<html><body>" +  
                            "<table width='100%' cellpadding='0' cellspacing='0' style='background-color: #f5f5f5;'>" +  
                            "  <tr>" +  
                            "    <td align='center' style='padding: 20px 0;'>" +  
                            "      <img src='https://res.cloudinary.com/duwzchnsp/image/upload/v1720011446/ea0nbypctt5ure1tgvxr.png' alt='Talk Town Technology' style='max-width: 200px;'>" +  
                            "    </td>" +  
                            "  </tr>" +  
                            "  <tr>" +  
                            "    <td align='center' style='padding: 20px;'>" +  
                            "      <h1 style='color: #333;'>Hello, customer!</h1>" +  
                            "      <p style='font-size: 16px;'>We've received a request to verify your identity. Your one-time password (OTP) is:</p>" +  
                            "      <p style='font-size: 24px; font-weight: bold; color: #007bff;'>" + otp + "</p>" +  
                            "      <p>Please enter this code on the verification page to complete the process.</p>" +  
                            "    </td>" +  
                            "  </tr>" +  
                            "</table>" +  
                            "</body></html>")  
                    .build();  
  
            SendEmailsRequest request = SendEmailsRequest  
                    .builder()  
                    .message(message)  
                    .build();  
  
            SendEmailsResponse response = request.sendWith(this.mailjetClient);  
  
        } catch (MailjetException e) {  
            logger.error("Error sending email: {}", e.getMessage());  
            throw new MailParseException("Could not parse mail", e);  
        }    }}
```


