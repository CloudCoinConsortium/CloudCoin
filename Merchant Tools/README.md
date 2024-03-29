# RECEIVING AND SENDING CLOUDCOINS PROGRAMMATICALLY

# NOTE: THIS IS FOR THE ORIGINAL CLOUDCOIN AND NOT FOR THE NEWEST VERSION. YOU SHOULD NOT USE THIS 

There are many ways to send and receive CloudCoins including by email, downloading from web sites and Skywallet. We suggest using Skywallet because of its ability to show the sender and the receiver that money was sent and received. It is also simple to implement. Skywallet has one drawback: it is pseudo anonymous in that RAIDA Administrators can see the account numbers and transactions. This is similar to the Blockchain.  Unlike crypto however, Skywallet transactions are not public.

# Receiving CloudCoins via Skywallet
You and your users will need to have a Skywallet account with debit card. You can get an account at: https://www.skywallet.cc/debit_card.html Skywallet debit cards are scheduled to be included in the CloudCoin Wallet 4.0 due to be released August 24, 2020. 

The major components of receiving CloudCoins via Skywallet are: 

1. Skywallet HTML widget
2. The backend Action Page. 
3. The raida_go command line program that your action page can call. 

## 1. Skywallet HTML widget
To learn about the HTML widget see: https://github.com/CloudCoinConsortium/POSJS/blob/master/sample/index.html

Here is the minimum HTML needed for the widget: 
```html
<!DOCTYPE html>
<html lang="en">
        <head>
                <script src="https://cloudcoin.global/assets/posjs.min.v004.js" type="text/javascript"></script>
                <script type="text/javascript">
                        var pos = new POSJS({
				'timeout':'5000',
				'action': 'https://yourdomain.com/your_action_page.php', 
				'merchant_skywallet' : 'Your.skywallet.cc'
			})
                         var get_parameters = {}; // Create a get_parameters object to hold GET parameters for your action page.
                         get_parameters.amount = 100;  // Mandatory GET variable specifies the amount of CloudCoins to be paid.              
                         get_parameters.customerID = '1554887'; // Optional Merchant Variable customized by you but must be a String. 
                </script>
        </head>
        <body>
	     <img onclick='pos.show(get_parameters)' src='https://cloudcoin.global/assets/paywithcc.001.png' width='100' alt='Pay with Cloud Coin'>
		
        </body>
</html>

```
## 2. Backend Action Page
You can use any language such as PHP, Python, Ruby, Java, C# etc. Your language will call the raida_go command line program and this executable will tell you how many CloudCoins the customer sent you. For more information read https://github.com/CloudCoinConsortium/POSJS/blob/master/sample/action.php

Here is a very simple example of an action page calling the raida_go executable:
```php
<?php
	
	$my_wallet = $_GET['merchant_skywallet'];
	$received_from = $_GET['sender_skywallet'];
	$amount_due = $_GET['amount'];
	$receipt_guid = $_GET['guid'];
	$merchantData = $_GET['merchantData'];//this can be customized. 

	$command = "E:/Documents/pos/raida_go.exe receive $receipt_guid $my_wallet"; //This is for Windows 
	//$command = "./raida_go receive $receipt_guid"; //This is for Linux. 
	$json_obj = exec($command); //Returns something like: {"amount_verified":100}
	$arr = json_decode($json_obj, true);
	$amount_verified = $arr["amount_verified"];
	echo "The amount received in $my_wallet is $amount_verified";

?>
```
## 3. The raida_go Executable
 The raida_go executable is a command line program that your action page can call in order to see if your customer has sent you CloudCoins. 
 This program comes compiled for Linux and Windows in a zip here https://cloudcoin.global/asseets/raida_go.zip. To learn more about raida_go click https://github.com/CloudCoinConsortium/raidaGo. 
 
 
Example of calling this CLI for Windows to check how many coins were sent to Skywallet 'Sean.CloudCoin.Global':
```dos
E:\CloudCoin\raida_go.exe receive 2cb825ee32a847d68650617cc6a3862a sean.cloudcoin.global
```
Example of the same call for Linux
```bash
./CloudCoin/RAIDA_GO.elf receive 2cb825ee32a847d68650617cc6a3862a sean.cloudcoin.global
```
Reponse from the raida_go if 100 coins were received: 
```json
{
  "amount_verified": 100
}
```

Reponse from the raida_go if no coins were received: 
```json
{
   "amount_verified": 0
}
```

# Sending CloudCoins From Your Skywallet to a Customer's Skywallet
You can quickly send CloudCoins from your wallet to the Customer's wallet by having your program call the 
raida_go executable' transfer command. You will need to include a memo. These memos can use our memo standard that allows them to show up in your transaction log as well as your customer's transaction log. Note that this technology is being implemented now and is not ready. 

```bash
./CloudCoin/raida_go send "From sean.cloudcoin.global VGhpcyBpcyBhIGZpbmUgd2F5ICB0byBkbyB0aGluZ3M=" billy.skywallet.cc
```
NOTE: This standard is in flux. This is an ini file that will be converted into base64 and concatenated after the memo. 
```ini
from=sean.cloudcoin.global
to=billy.skywallet.cc
description=Four pairs of shoes, Socks, 
transaction_url=https://er.miroch.ru/c/6200c62cc6a94aa39f98894ad0347f35.html
transaction_img=https://er.miroch.ru/img/6200c62cc6a94aa39f98894ad0347f35.im
```



# Skipping raida_go and Talking Directly to the RAIDA
This is not recommended because there are 25 of them that should be contacted in parallel and this can be tricky. Also the raida_go program will syncronize your Skywallet account and talking directly to the RAIDA will not. This could cause your Skywallet to become unsyncronized amount all the RAIDA. 

SAMPLE CALL TO VIEW RECEIPT
```
https://raida7.cloudcoin.global/service/view_receipt?account=554452&tag=623e88186f3c4a4694c02230abe72666
```

SAMPLE RESPONSE FROM SKYWALLET
```
{
	"server": "raida18",
	"total_received": 454,
	"serial_numbers": "1308799, 1308800, 1308868, 1308983, 14339990, 14340140, 16458252",
	"version": "2020-02-13",
	"time": "2020-04-20 23:31:40",
	"execution_time": 0.0022211554454
}
```

PARAMETERS
```
tag: guid used in the memo of the transaction (MUST BE A GUID). 
account: Serial Number of your SkyWallet ID coin.
```
NOTES ON PARAMETERS

Tag: The GUID can be generated by your customer or you can tell the customer what the GUID will be.  

Account: You can get this by looking in the ID folder on your CloudCoin Wallet and then opening your ID coin in a text editor. You must first have a SkyWallet account which can be created using the CloudCoin Wallet. To make things easier, we convert these numbers into IP addresses and then store them on DNS servers so that we can use friendly names that are easy to use and provide confidence in transactions. For example, if you owned Walmart, you would want to want to convert your skyWallet account number to an IP and put it on Walmart DNS servers. So then your user SkyWallet account may be "payments.Walmart.com" and this would translate to IP address 1.232.70.172. If you convert that IP address to a decimal, it becomes 15222444. Because only you have access to the Walmart DNS servers, it proves that Walmart is the owner of account 15222444 and that it is safe to send money there.   

EVALUATING RAIDA RESPONSES
You must call the "view_receipt" service on 25 RAIDA in parallel. As a simple rule, you need 20 responses confirming the total to be certain that you received payment. However, if you use "Advanced Grading" you can be safe with 16 positive responses or less.  

REFERENCE

###### Stadard for ViewReceipt Service
###### https://github.com/CloudCoinConsortium/CloudCoin/blob/master/SkyWallet%20Protocol/README.md#viewreceipt

###### Javajs.js API for creating a "RAIDA" class that can access the RAIDA and SkyWallet.
###### https://github.com/CloudCoinConsortium/raidajslibrary

###### Download the CloudCoin Wallet with RPC (Requires version 3.0.3 or beyond)
###### http://CloudCoinConsortium.com/use.html

###### Sample Skywallet Store Accepting Coins:
###### https://pownesium.com/coke_sample_store.php

###### Backend of the Store:
###### https://github.com/CloudCoinConsortium/CloudCoin/blob/master/Merchant%20Tools/PaymentVerifier/sample_store/payment_verifier.php

###### Skywallet ATM (All static javascript, html, css using the RAIDA class from RaidaJs.js
###### http://pownesium.com/atm/

###### Tool to create Skywallet ATM Card
###### http://pownesium.com/atm/debitcard.html

###### PHP Store that accepts CloudCoins as payment
###### https://github.com/CloudCoinConsortium/PHPDemoOnlineStore

###### Coke Sample Store
###### https://github.com/CloudCoinConsortium/CokeSampleStore
# Sending Coins

Your customers may want to download coins from your website. Or they might want to receive coin in their email. If they have a Skywallet Account, they may want to receive coins there. You may also want to give your customers a link that allows them to download their coins. 

The easyist way to send Coins to your customers is to use CloudCoin Wallet with RPC. Download the CloudCoin Wallet and run it on your desktop computer or server. The CloudCoin wallet will act as your depository and hold your coins. The CloudCoin Wallet with RPC is written in java and will run on any platform with a GUI.  You will need to enable CloudCoin Wallet RPC (Remote Procedure Calls). This allows your web servers to command your CloudCoin Wallet to give it a stack of CloudCoins. You can then put these coins on your website and email the coins to your customers. 

SAMPLE REQUEST FOR A STACK OF COINS
```
https://195.45.84.78:33/service/withdraw_one_stack?amount=254&pk=ef50088c8218afe53ce2ec&account=532fe9e5dc3932650cfa
```
SAMPLE REQUEST FOR A STACK OF COINS FOR WALLET THAT IS ENCRYPTED IN THE WALLET
```
https://195.45.84.78:33/service/withdraw_one_stack?amount=254&pk=ef50088c8218afe53ce2ec&account=532fe9e5dc3932650cfa&ek=1170b354e097f2d
```
SAMPLE REQUEST WITH OPTIONAL MEMO
```
https://195.45.84.78:33/service/withdraw_one_stack?amount=254&pk=ef50088c8218afe53ce2ec&account=532fe9e5dc3932650cfa&base64=TransactionNo887734
```
SAMPLE RESPONSE IF GOOD
```
{
	"cloudcoin": [
		{ 
		"nn":"1", 
		"sn":"1112240", 
		"an": ["f5a52ee881daaae548c24a8eaff7176c", "415c2375a6fa48c4661f5af8d7c95541", "73e067b7b47c1556deebdca33f9a09fb", "9b90d265d102a565a702813fa2211f54", "e3e191ca987c8010a3adc49c6fc18417",
			"baa7578e207b7cfaa0b8336d7ed4a4f8", "6d8a5c66a589532fe9e5dc3932650cfa", "1170b354e097f2d90132869631409bd3", "b7bc83e8ee7529ff9f874866b901cf15", "a37f6c4af8fbcfbc4d77880fc29ddfbc",
			"277668208e9bafd9393aebd36945a2c3", "ef50088c8218afe53ce2ecd655c2c786", "b7bbb01fbe6c3a830a17bd9a842b46c0", "737360e18596d74d784f563ca729aaea", "e054a34f2790fd3353ea26e5d92d9d2f",
			"7051afef36dc388e65e982bc853be417", "ea22cbae0394f6c6918691f2e2f2e267", "95d1278f54b5daca5898c62f267b6364", "b98560e11b7142d1addf5b9cf32898da", "e325f615f93ed682c7aadf6b2d77c17a",
			"3e8f9d74290fe31d416b90db3a0d2ab1", "c92d1656ded0a4f68e5171c8331e0aea", "7a9cee66544934965bca0c0cb582ba73", "7a55437fa98c1b10d7f47d84f9accdf0", "c3577cced2d428f205355522bc1119b6"],
		"ed":"7-2019",
		"pown":"ppppppppppppppppppppppppp",
		"aoid": []
		},
		{ 
		"nn":"1", 
		"sn":"1112241", 
		"an": ["f5a52ee881daaae548c24a8eaff7176c", "415c2375a6fa48c4661f5af8d7c95541", "73e067b7b47c1556deebdca33f9a09fb", "9b90d265d102a565a702813fa2211f54", "e3e191ca987c8010a3adc49c6fc18417",
			"baa7578e207b7cfaa0b8336d7ed4a4f8", "6d8a5c66a589532fe9e5dc3932650cfa", "1170b354e097f2d90132869631409bd3", "b7bc83e8ee7529ff9f874866b901cf15", "a37f6c4af8fbcfbc4d77880fc29ddfbc",
			"277668208e9bafd9393aebd36945a2c3", "ef50088c8218afe53ce2ecd655c2c786", "b7bbb01fbe6c3a830a17bd9a842b46c0", "737360e18596d74d784f563ca729aaea", "e054a34f2790fd3353ea26e5d92d9d2f",
			"7051afef36dc388e65e982bc853be417", "ea22cbae0394f6c6918691f2e2f2e267", "95d1278f54b5daca5898c62f267b6364", "b98560e11b7142d1addf5b9cf32898da", "e325f615f93ed682c7aadf6b2d77c17a",
			"3e8f9d74290fe31d416b90db3a0d2ab1", "c92d1656ded0a4f68e5171c8331e0aea", "7a9cee66544934965bca0c0cb582ba73", "7a55437fa98c1b10d7f47d84f9accdf0", "c3577cced2d428f205355522bc1119b6"],
		"ed":"7-2019",
		"pown":"ppppppppppppppppppppppppp",
		"aoid": []
		}

	]
}
```
REQUIRED PARAMETERS
```
amount: The amount of CloudCoins to be returned in the stack.
pk: The password that your CloudBank has generated for you. 
account: The account ID that your CloudBank has generated for you.
```
OPTIONAL   PARAMETERS
```
ek: The password used to decrypt your CloudCoin Wallet. This is needed if you have choosen to password protect your wallet. 
base64: A memo or note that will be placed in the transaction file on the CloudCoin Wallet. 
```

Your Wallet will automatically make change if needed. You can put your CloudCoin Wallet on a server located in a physically secure place (like your home). You can even have lots of CloudCoin Wallets on different machines around the world to mitigate the risk of theft and loss. You can encrypt your Wallet so that the person who is the custodian/administrator of the CloudCoin Wallet does not have the ability to rob them. In the future, the CloudCoin Wallet with CloudBank will allow you to configure limits on sending money, restrictions on who can request money, times of day that money can be sent and even if you want manual confirmation.

Once your server gets your stack file, it can put it on a web page, stored with a link for the customer to download, emailed to the user or even turned into a QR code, png or jpeg file (using the RAIDA class functions in the RaidaJs.js).  

REFERENCE:
###### Withdraw one Stack standard
###### https://github.com/CloudCoinConsortium/CloudBank-V2#withdraw-one-stack-service

###### Sample Withdraw One Stack on a web server using Javascript
###### https://bank.cloudcoin.global/banking/test_withdraw_one_stack.html

###### Sample webpage that creates a jpe from a stack file (Note that the CloudCoin is in the GET variables. 
###### https://pownesium.com/coke_claims.php?nn=1&sn=1304167&pown=pupppppppuppppppppppppppp&ed=0-2022&an0=2acd1e906917a1a088a25cb35453635a&an1=086f74522bc612dc2b119c6643094da5&an2=f0a48b139d4983c9128b7f064c1f5176&an3=58bac3f0bfc558dfb9ba7ed43f3e5084&an4=90591c58319e8a34d2166de3585ebb82&an5=a99dff2b61b69090d84131da012e5e6d&an6=92218b885b2fa625516112a1a4a3276c&an7=d2994a2d4aa659e7a998a95e0e014d62&an8=2ba775297a55283ba46085877dbd46e1&an9=0c357f93c12e0b88a32c3667beffde63&an10=52fa09a2bd86431e876b0f6088c3f5fd&an11=58f743a11db34f8e2cab88fd47b588fc&an12=6c298b9b0e1e64154bd7b4c96c55bed4&an13=3c788d749abb2ae92748e23191cb4f08&an14=14f9d272385f559eb1dcbb125852156b&an15=ca576b3ab19168c4ac8dd87e52de274b&an16=f054af9c350e12cca1be14328d1e6aec&an17=a23961be33d977433795240a6945a839&an18=9e5e9ee3a4b8438faffb3e95e3a83e5c&an19=d78e171850513fd197d8142cf4f80b18&an20=103a6de638333eefb3b4cd47a313ca45&an21=efabf0011f1901604fe36d468b121958&an22=3a907d76d6b4fdc1250a73f53e5cab0b&an23=49df92630c84c7f4620299376f1931a7&an24=25c75855568f1ba12e7703c041bfe03b&

###### Javajs.js API for creating a "RAIDA" class that can access the RAIDA and SkyWallet.
###### https://github.com/CloudCoinConsortium/raidajslibrary

###### Download the CloudCoin Wallet with RPC (Requires version 3.0.3 or beyond)
###### http://CloudCoinConsortium.com/use.html

# Other Ways to Send and Receive

### Skywallet to Skywallet using "Transfer.go"
If you want to sent CloudCoins directly from your SkyWallet to a customer's SkyWallet you can use our "Transfer.go" program that can be called from any language. However, this "Transfer.go" program is still in Alpha but it may be done by the time you read this.

### Skywallet to Skywallet using CloudCoin Wallet
If you want to command you CloudWallet CloudBank to transfer funds from your SkyWallet you can use our "Transfer" RPC. This is under deveoplpment. 

### CloudCoin Wallet to Skywallet using RPC
If you want to command your CloudWallet CloudBank to send CloudCoins from your local wallet to a SkyWallet, you can use our "SendToSkyWallet" CloudBank service. This is underdevelopment. 

### Accepting Skywallet Debit Card for Skywallet to Skywallet
You can accept a Skywallet Debit Card in the near future. It will allow a web page to transfer coins from their Skywallet to yours. 

## Other API
The CloudBank is designed to allow you to do all the things that the CloudCoin Wallet can do through remote procedure calls. We have APIs written in Java, JavaScript, Python, C# and Unity, C++, Unreal Engine, GoDot, Ruby and others. These APIs allow you to control your CloudBank remotly. We also have the RaidaJS API that allows you to access Skywallet via Javascript.  Most of these APIs need some work. If you have need for APIs in these languages, let us know so we can However, they are not well developed. We can schedule them to be improved and ready for use. 

Javascript SDK
https://github.com/CloudCoinConsortium/raidajslibrary

PHP SDK
https://github.com/CloudCoinConsortium/CloudBank_PHP_SDK

Python SDK
https://github.com/CloudCoinConsortium/SDK-Python

C# SDK
https://github.com/CloudCoinConsortium/CloudCoin-Csharp-SDK

Godot SDK
https://github.com/CloudCoinConsortium/CloudCoin-GodotEngine-SDK

Go Lang
https://github.com/CloudCoinConsortium/CloudCoin

Java SDK
https://github.com/CloudCoinConsortium/CloudCoin-Java-SDK

C++ SDK (Created by Volunteers, untested, disorganized)
https://github.com/CloudCoinConsortium/CloudCoin-CPP-SDK/blob/master/old/DemoProgram/DemoProgram.cpp

UnrealEngine C++ SDK
https://github.com/CloudCoinConsortium/CloudCoin-UnrealEngine-SDK



# Contact us If you Need Help
Please email Support@CloudCoin.global
or 
CloudCoin@Protonmail.com 


