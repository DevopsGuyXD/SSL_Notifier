<h1 align="center">SSL Notifier</h1>

<h4 align="center">Looking for a reliable SSL notification system? Look no further!!</h4>

<p align="center">Its simple, pull the image with the appropriate tag of "aws" or "azure" based on your cloud provider and initiate the docker container with the much required access and notification information.</p>

<br>
<h3 align="center">Example:</h3>

<p><b>AWS:</b></p>

    docker run --rm -e ACCESS_KEY_ID="AKIAYUOXBHDKW5SDQELB" -e ACCESS_KEY_SECRET="OdreftDd+2h/9FLiqq+hnJxAmY+iuTUwJ52fieNs" -e AWS_REGION="ap-south-1" -e EMAIL_SENDER_ID="nfcloudsecurity@nowfloats.com" -e EMAIL_SENDER_PASSWORD="P67%4urG123" -e RECEIPIENT_MAIN="bharath.dundi@nowfloats.com" -e RECEIPIENT_CC_1="vishal.sharma@nowfloats.com" -e RECEIPIENT_CC_2="saurabh.verma@nowfloats.com" -e IS_CRON="false" -e CRON="* * * * *" devopsguyxd/sslnotifier:aws
    
![2023-03-11 21_30_42-Window](https://user-images.githubusercontent.com/77780574/224494726-48f26a04-1905-4f66-8cdc-061d3bd247ab.png)

<p><b>AZURE:</b></p>

<p>One time execution:</p>

    docker run --rm -e CLIENT_ID=" " -e TENANT_ID=" " -e SECRET_VALUE=" " -e EMAIL_SENDER_ID=" " -e EMAIL_SENDER_PASSWORD=" " -e RECEIPIENT_MAIN=" " -e RECEIPIENT_CC_1=" " -e RECEIPIENT_CC_2=" " sslnotifier:azure
    
<p>CRON:</p>

    docker run --rm -e CLIENT_ID=" " -e TENANT_ID=" " -e SECRET_VALUE=" " -e EMAIL_SENDER_ID=" " -e EMAIL_SENDER_PASSWORD=" " -e RECEIPIENT_MAIN=" " -e RECEIPIENT_CC_1=" " -e RECEIPIENT_CC_2=" " -e CRON="* * * * *" devopsguyxd/sslnotifier:azure-cron
    
![2023-03-02 23_24_21-NowFloats-API-monitoring - Chat](https://user-images.githubusercontent.com/77780574/222514102-3aaa8fd8-e09c-428d-a0fd-a71ca851b543.png)

<br>
<p align="center"><b>Note:</b><p>

1. The notification will be sent when the certificate has less than 15 days to expire.
2. Make sure the user has at the least "Read" access to "AWS Certificate Manager" in AWS.
3. For Azure the service principal being used needs to have at the least "Read" access to "App Service Certificates" and "Key Vault"
4. Make sure your email provider trusts this application in order for it to send notifications successfully to the specified email addresses.
5. In this current version, the notification can only be sent to a main email address along with two more addresses in CC at the max.
