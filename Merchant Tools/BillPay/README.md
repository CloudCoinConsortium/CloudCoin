#Bill Pay

Requirements: 
1. A professional subscrtiption to Protonmail
2. Download the ProtonBridge
3. You can only send 150 email each day. 
4. Create the email template file
5. Create the email configuration file
6. Make the eamil list csv file.
7. Start and monitor the emaling. 

## Create the Email Template file

You must have a file called etemp.txt located in your username CloudCoinWallet:
```
C:\Users\sean\CloudCoinWallet\EmailTemplates\etemp.txt
```
SAMPLE etemp.txt contents:
```txt
Hello,
Find your %amountToSend% coins attached below
Sean
```
## Make the email list .csv file. 

CSV means comma seperated values. It can be opened in any spreadsheet program such as google sheets and Micorosoft Excel. 
The mailing list can be located any where on your computer, clouddirves


Sample of tile contents (Scheduled to change in fuuture editions)
```
send,20000,0,0,0,0,0,Pfinal@protonmail.com,feb.txt
send,12500,0,0,0,0,0,tyfern2000@protonmail.com,feb.txt
send,12500,0,0,0,0,0,Lstang38@protonmail.com,feb.txt
send,12500,0,0,0,0,0,ldfever@protonmail.com,feb.txt
send,12500,0,0,0,0,0,zulennon@protonmail.com,feb.txt
send,12500,0,0,0,0,0,difidfchrod@protonmail.com,feb.txt
send,12500,0,0,0,0,0,Michadlkeinnis@protonmail.com,feb.txt
send,12500,0,0,0,0,0,iiuskidfdse@protonmail.com,feb.txt
send,12500,0,0,0,0,0,Sdfedfdnghorn@protonmail.com,feb.txt
send,12500,0,0,0,0,0,rrrepid43@protonmail.com,feb.txt
send,12500,0,0,0,0,0,bdboat@protonmail.com,feb.txt

```




or usb drives.
## Create a mailsettings.ini file

You must put a file here:
```
C:\Users\sean\CloudCoinWallet\mailsettings.ini
```
Contents of mailsettings.ini
```
[smtp]
from=CloudCoin Banker<CloudCoinBanker@protonmail.com>
mail_from=<CloudCoinBanker@protonmail.com>
username=CloudCoinBanker@protonmail.com
password=_wb5Uv7RKj3ou-DZ6GEELQ
smptServerAddress=localhost
port=1025
maxEmailsPerRun=10

```
Use protonbridge to see your username and password, port. Your smptServerAddress will be localhost. 
You can increaes the number of email per run to 150 if you plan on mailing one day. 

