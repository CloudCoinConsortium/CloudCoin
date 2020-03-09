# Bill Pay
## Distributing CloudCoins in Mass
CloudCoin maybe the only currency in the world that can be distributed without fees. The cost of sending Bitcoin and other crypto currencies, to massive amounts of users can be extreemly expensive. For most payment systems, it is less expensive to send cash via the plain old mail than electonically. The only advantage of payment systems over regular mail is the speed, accountability and dependability. With CloudCoin, you can pay thousands of people within minutes and pay no fees. You also have the addition of privacy which is not possible with most payment systems such as PayPal, Western Union, Bank Transfer and others.

If you want to be able to get CloudCoins out to a lot fo people quickly, use the CloudCoin Wallet's Bill Pay.
To get the CloudCoin Wallet's Bill Pay to work, you will open your CloudCoin Wallet and then go to Tools>BillPay The progrm will then ask you to choose what wallet you want to send money from. Then you will be asked to load a CSV (Comma Serperated Value) file. This "CSV" file will specify who you want to send money too, how much to send, and which transport system to use. This CSV file can be created and edited in a spreadsheet program such as Microsoft Excel, Google Sheets, Open Office Sheets or any other spreadsheet software). You can also just use a text editor.

To send you CloudCoin via encrypted email, you will need to have a Protonmail account with the Gold plan so that you can download the Proton Bridge software. The Proton Bridge software will place an SMTP mail server right on your computer.

Requirements: 
1. A professional subscrtiption to Protonmail
2. Download the ProtonBridge
3. You can only send 150 email each day. 
4. Create the email template file
5. Create the email configuration file
6. Make the eamil list csv file.
7. Start and monitor the emaling. 
8. Sky Wallet account if sending via Sky Wallet

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

Bill pay will read this file, make some notes in it and then move the file to the Sent folder once it has processed the file.

CSV means comma seperated values. It can be opened in any spreadsheet program such as google sheets and Micorosoft Excel. 
The mailing list can be located any where on your computer, clouddirves

| Send Method | Format | Amount | 1 | 5 | 25 | 100 | 250 | To Address | Subject, Memo or FileName | Special Instructions | status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
|Protonmail	| stack | 238000 | 0 | 0 | 0 | 0 | 0 | Billy@Protonmail.com	| January Pay. Thank you! |	body=c:\templates\monthlypay.html | Ready|
| Gmail	| stack	| 100 |	0 |0|0	|0	|0	|Billy@gmail.com	|Thanks for signing up!	|body=c:\templates\monthlypay.html	|Ready|

|Protonmail	| stack |	238000|	0	|0	|0	|0	|0	|Billy@Protonmail.com	|January Pay. Thank you!|	body=c:\templates\monthlypay.html	|Ready|
|Gmail	|stack	|100|	0	|0	|0	|0	|0	|Billy@gmail.com	|Thanks for signing up!	|body=c:\templates\monthlypay.html	|Ready|
|Protonmail	|debitcard|	0|	1|	0|	0|	0|	0|	Jenny@protonmail.com	|Your Skywallet ID card.	|Cardholder=jen.skywallet.cc& |body=c:\letter.html|	ready|
|SendToSkywallet	|(ignored)	|238000|	0|	0	|0	|0	|0	|larry.skywallet.cc|	Thanks for the work you did.|	cloudbank=Skywallet.cc|	sent|
|TransferToSkywallet	|(ignored)|	238000	|0	|0	|0	|0	|0	|larry.skywallet.cc|	Thanks for the work you did.|	cloudbank=Skywallet.cc|	ready|
|ExportFolder|	single	|card-data	|1	|0	|0	|0	|0	|c:\users\production\	|Welcome to our club!|	domain=https://my.com/claim.html?&body=c:\template.html|	skip|
|ExportFolder|	single|	multiurl | 35	|0|	0	|5	|4	|2	|C:\list\	Invoice 34565|	https://domain.com/mypage.html?|	ready|

Sample of tile contents (Scheduled to change in fuuture editions)
```
send,20000,0,0,0,0,0,Pfinal@protonmail.com,feb.txt
protonmail,12500,0,0,0,0,0,tyfern2000@protonmail.com,feb.txt
gmail,12500,0,0,0,0,0,Lstang38@protonmail.com,feb.txt
send,12500,0,0,0,0,0,ldfever@protonmail.com,feb.txt
transfer,12500,0,0,0,0,0,zulennon@protonmail.com,feb.txt
protonmail,12500,0,0,0,0,0,difidfchrod@protonmail.com,feb.txt
exportFolder,12500,0,0,0,0,0,Michadlkeinnis@protonmail.com,feb.txt

```



### Column Details

#### Send Method

Possible methods include:
* Protonmail (Sends encrypted mail using Protonmail Bridge)
* SendToSkywallet (From local bank folder to Skywallet)
Phase 2:
* Gmail (Sends plain text email)
* TransferToSwyallet (From Skywallet to Skywallet)
* ExportFolder (From Bank to local file system)

#### Format

Formates used in Bill Pay include:
* Raw URL ( To Embed One Coin in an Email )
* Compressed URL ( Uses the Compression Standard to Embed One Coin in an Email )
* Multi Compressed URL ( To Embed Many Coins in an Email )
* Compressed CSV ( To write many compressed coins to a csv file so it can be accessed later. )
Future Use:
* ID Coin PNG ( One PNG file attached to an Email )
* stack ( One '.stack' file attached to Email )
* QR Code File ( Saves codes to the local drive. These codes can generate graphical QR codes. )
* jpeg ( A few jpeg files attached to email )
* The Send and Transfer send methods do not need a format as they are sent directly to a CloudBank. See details about file formats.

#### Amount
The amount of CloudCoins to be sent. If this column is used the following five columns will be ignored.


#### 1s	
The number of 1 note CloudCoins to be sent.

#### 5s	
The number of 5 note CloudCoins to be sent.

#### 25s	
The number of 25 note CloudCoins to be sent.

#### 100s	
The number of 100 note CloudCoins to be sent.

#### 250s	
The number of 250 note CloudCoins to be sent.

#### To Address

1. For Protonmail or Gmail, use an email address
2. For Send and Transfer, use a skywallet account ID
3. For Export Folder, the path to the folder. Leave blank for the Default Folder.


#### Subject, Memo or Filename	
1. Subject: For Protonmail or Gmail
2. Memo: For Send and Transfer
3. Tag: For jpegs or stack files that are put on file system
4. File Name: For other files placed on the file system



#### Special Instructions
These are key value pairs seperated by "&" ampersams
1. Body: For Protonmail or Gmail, this is location of the template for the body of the email like: "Body=c:\body.txt&priority=high"
2. AttachmentUnique: For Protonmail or Gmail, this is location of a folder full of coins. The program will remove one coin from this folder and attach it to the email. One: 3. "AttachmentUnique=C:\Celebrium\Gallery"
4. For Send and Transfer do not need special instructions and can be left blank.
URL: For URLs this is the URL of the webpage that the process the coin like "url=https://cloudcoin.global/public/claimcoin.html



#### Status
The statuses are:
1. ready: Ready to be sent
2. skip: Bill Pay shoud skip this payment and not pay it.
3. sent: Payment was sent
4. error: Error means something whent wrong and the BillPay log files should be examined.



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

