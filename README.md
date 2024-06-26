<h1 align="center">SSL Notifier</h1>

<h4 align="center">Looking for a reliable SSL notification system? Look no further!!</h4>

<p align="center">Its simple, pull the image with the appropriate tag of "aws" or "azure" based on your cloud provider and initiate the docker container with the much required access and notification information.</p>

<br>
<h3 align="center">Example:</h3>

<p><b>AWS:</b></p>

    docker run --rm -e ACCESS_KEY_ID=" " -e ACCESS_KEY_SECRET=" " -e AWS_REGION=" " -e EMAIL_SENDER_ID=" " -e EMAIL_SENDER_PASSWORD=" " -e RECEIPIENTS="ABC@example.com,123@example.com,XYZ@example.com" -e DAYS_TO_NOTIFY_BEFORE_EXPIRY="" -e IS_CRON="true | false" -e CRON="* * * * *" devopsguyxd/sslnotifier:aws
    
![2023-03-11 21_30_42-Window](https://user-images.githubusercontent.com/77780574/224494726-48f26a04-1905-4f66-8cdc-061d3bd247ab.png)

<p><b>AZURE:</b></p>

    docker run --rm -e CLIENT_ID=" " -e TENANT_ID=" " -e SECRET_VALUE=" " -e EMAIL_SENDER_ID=" " -e EMAIL_SENDER_PASSWORD=" " -e RECEIPIENTS="ABC@example.com,123@example.com,XYZ@example.com" -e DAYS_TO_NOTIFY_BEFORE_EXPIRY=" " -e IS_CRON="true | false" -e CRON="* * * * *" devopsguyxd/sslnotifier:azure
    
![2023-03-02 23_24_21-NowFloats-API-monitoring - Chat](https://user-images.githubusercontent.com/77780574/222514102-3aaa8fd8-e09c-428d-a0fd-a71ca851b543.png)

<br>
<p align="center"><b>Note:</b><p>

1. Make sure the user has at the least "Read" access to "AWS Certificate Manager" in AWS.
2. For Azure the service principal being used needs to have at the least "Read" access to "App Service Certificates" and "Key Vault"
3. Make sure your email provider trusts this application in order for it to send notifications successfully to the specified email addresses.
4. In this current version, the notification can only be sent to a main email address along with two more addresses in CC at the max.
