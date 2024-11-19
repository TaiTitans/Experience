
---

```
JavaMailSender
```

EmailSenderUtil

```Java
public class EmailSenderUtil{
@Autowire
private JavaMailSender javaMailSender;
public void sendTextEmail(String to, String subject, String content){
SimpleMailMessage message = new SimpleMailMessage();

message.setTo(to);
message.setSubject(subject);
message.setText(content);
message.setFrom(EMAIL_HOST);
try{
javaMailSender.send(message);
System.out.println("Email sent Successfully");
}catch(Exception e){
throw new RuntimeException(e);
}
}
public void sendHtmlEmail(String to, String subject, String content){
try{
MineMessage message = javaMailSender.createMimeMessage();

MimeMessageHelper helper = new MimMessageHelper(message, true);
helper.setFrom(EMAIL_HOST);
helper.setTo(to);
helper.setSubject(subject);
helper.setText(content, true);
javaMailSender.send(message);

System.out.println("Email sent Successfully");

}catch(Exception e){
throw new RuntimeException(e);
}

}

}
```

```Java
@Test
void sendHTMLEmail(){
...
Resource resource = new ClassPathResource("/templates/email/otp-auth.html");
String htmlContent = new String(resource.getInputStream().readAllBytes());
EmailSenderUtil.sendHtmlEmail(to, subject, htmlContent);
}

```
Folder: resource/template/email/htmlfile


