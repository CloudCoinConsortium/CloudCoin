# Payment Verifier

This code allows client software to connect to the Skywallet, also known as Trusted Transfer or CloudBank. The Skywallet is the first distributed application built on RAIDA technology. The Skywallet can authenticate, transfer and store CloudCoins making it:
1. Easy to sent CloudCoins from one person to another or to make payments.
2. Store CloudCoins so that they can be available any where and on any platform. 

The drawbacks are:
1. Skywallet requires the use of an ID Coin. This means that unlike CloudCoin, the Skywallet is pseudo anonmymous like crypto currencies and not 100% anonmymous like CloudCoin. It also means that the Skywallet has system risk of loss and theft due to its need for a single private key (like Crypto currencies). However, we are implementing key storage and retreival systems that will reduce the risk of theft and loss.  

Everytime a CloudCoin is sent or transfered from one entity to another using the Skywallet, a "memo" is included. The memo allows the sender to include information that will allow the receiver to know what the payment if for. 

The Payment Verifier sees if a payment has been received based on the memo. Then, if the payment has been receieved, the Payment Verifier may rename the memo. If the sender sent too many or too few CloudCoins, the Payment Verifier will
return the coins to the sender. The Payment Verifier expects the exact amount. It will not make change. (Although that may be a features in the future). 

## Sample Usage:
```
paymentVerifier.exe -timeout=5 -oldtag=general_fund_e51c28025eb44caa82c77f85d69fe0f9 -payment=100 -refundto=1371486 -logpath="C:\user\bill\CloudCoin\Skywallet\bill.skywallet.cc" -idpath="C:\cc\Accounts\Change\ID\1.CloudCoin.1.2..stack" 
```

The Payment Verifier is a command line program written in GO that can be executed by other languages including php, Java, C# and almost all others. The comand line arguments include flags.  


## Flags
* oldtag: A string that specifies the memo that the sender provided. The Payment Verifier will search to see if a payment with this memo is in your SkyWallet. You can often give this memo to the customer ahead of time to make sure it is a random number and to automatically be on the look out for it. Or you can wait for the customer to tell you that they have sent a payment and what the memo is. 

* payment:  A string that is an integer showing how many CloudCoins the user was required to send. Once the Payment Verifier finds a payment that matches the memo, it will check to see if the amount of the payment matches the amount that was suppose to be received.

* newtag: A string that specifies the fund or subaccount the payment will go into. Once you recieve a payment, you will want to change the memo so that the person cannot claim they sent the payment twice. We suggest that you combine fixed information and add a random number at the end of your tag such as: "general_fund_c77f85d69fee51c28025eb44caa820f9. Your new tag can be up to 255 characters of Unicode. 

* logpath: A string specifiying the path to the folder where you want the event log to be stored. This shows all errors. See logging below. 

* idpath: A string specifying the path to the "ID coin" that will be used to access the receiver's account. You must have a CloudCoin that is used as your authentication with the SkyWallet. The Serial Number of this coin will become your account number. You must place this somewhere secure so that the Payment Verifier can use it to access your Skywallet. 

* timeout A string that describes the number of seconds that the program will wait for RAIDA to respond. The program will attempt to contact 25 RAIDA. In general, they will all respond within 2 seconds. However, there may be one or two slow ones. You do not need all 25 to respond. You only need 20 reponses. So you can cut off the slowest five. Generally, we recommendt 5 seconds. If your Internet connection is slow, then you may need to increase this. If your Internet connection is very fast, you may decrease this.

* refundto: A string that is a Serial number, IP or Account Name of the receiver. e.g. 16777216, 1.255.255.255 or Sean.CloudCoin.global. When your customer sends you a payment, they have an option of including their address. This allows you to reject payments that are too large or too small and refund them instantly. 

## How it works
You will place the paymentVerifier.exe on your computer/server and point your software to it. You will include a command in your software that calls the paymentVerifier.exe to be executed with flags. Before the PaymentVerifier is called, the caller should check to make sure they can connect to the RAIDA. The "Echoer" servant can do this. Echo ensures that enough RAIDA can be contacted. Sometimes, local routers will block connection to the RAIDA. Also, sometimes computers are too underpowered to create the 25 threads that the Payment Verifier uses to contact the SkyWallet. If you have an old computer over 7 years old or are connecting from work, you may be concerned about the availabilty. You can skip calling echoer if you have a reliable connection and good computer that has been proven to work. You can test if it works by downloading the CloudCoin Wallet and running the echo command from the tools menu. The CloudCoin Wallet can be found at https://CloudCoinConsortium.com/use.html  You will need the CloudCoin Wallet to create a SkyWallet account for yourself. 

Your Payment Verifier will return a JSON message if successful.     
```
{"status":"success","message":"25 good replies. Execution Time = 25ms","time":"2006-1-2 15:04:0"}
```

Otherwise it will return a fail message like: 
```
{"status":"fial","message":"error with command line argument.","time":"2006-1-2 15:04:0"}
```
Fail Errors may inclue:
1  error with command line argument
3  error cannot open Keys file
4  error Cannot read Keys File
5 error cannot unmarshal json
6  error cannot create log
7  error cannot create log request
12 error Not Enough Good Replies
15 error with command line flags: missing flags
20 error with command line flags: invalid Oldtag Supplied. Oldtag cannot contain the words "change" or "public_change"

## Logging
There are two types of logs. The event log and the transaction log. They each have their own log locations. 

There will be a different log file created for each day. The names of the log files will be formatted like:

```
Jan.25.2019.PaymentVerifier.txt
Jan.26.2019.PaymentVerifier.txt
Jan.27.2019.PaymentVerifier.txt
```

The records inside the log file will include the hour and minute, the SenderID, amount, receiverID, fromTag, tag. Like this: NOTE the use of military time.
```
11:45 16777216 verified that billy.skywallet.cc sent 235 coins to tag "From Billy". New tag: "general_fund_e51c28025eb44caa82c77f85d69fe0f9"
12:32 16777216 DID NOT verify that sara.skywallet.cc sent 235 coins to tag "From Sara". Too many coins sent.
15:27 16777216 DID NOT verify that sean.cloudcoin.global sent 10000 coins from tag "invoice 8878734". No coins sent.
```
Transaction Log name: "transactions.txt"

The format is: 
Tag, date, Total_expected, Withdraw, Total, new_tag

* The "Withdraw" is always empty for Payment Verifier. 

* The "Total" is figured out by doing a show call to the RAIDA to find the total amount in the account. 

Contents of Transaction.txt file:
```

For New Sneakers,2019-Sep-13 8-15PM,250,,250,general_fund_e51c28025eb44caa82c77f85d69fe0f9
I am a customer,2019-Sep-16 9-20PM,6000,,6250,general_fund_e51c28025eb44caa82c77f85d69fe0f9
order 29938,2019-Sep-16 9-23PM,500,,6750,Accounts Receivable

```

## Details about the GO code. 

Unless you are trying to modify the Go code, you probably don't need to consider the following information.

Here are the steps that go on in the Payment Verifier. 

1. See if the refundto is a decimal, IP Adress or account name such as billy.skywallet.cc. case decimal: Validate it. case IP: Convert it to decimal case Account Name: Do a DNS lookup and and change the IP to the decimal.

2. Validate other flags.

3. Load the ID coin.

4. Show Coins in the Skywallet. THIS IS THE HARD PART: Then program must call the verify_payment service on all the RAIDA. 

### Example GET asking to verity the amount and envelope name of the claimed payment.
```

https://s0.teleportnow.cc/service/verify_payment?
---ID COIN---
nn=1&
sn=1&
an=1836843d928347fb22c2142b49d772b5&
pan=1836843d928347fb22c2142b49d772b5&
denomination=1&
---Payment Information---
tag=1Z 999 AA1 01 2345 6784&
new_tag=change
total_expected=33
refund_sn=16555897

```

5. See if 20 or more payments were verified. 

6. Call the Show service and count the total number of coins in the wallet. 

7. Make sure all is logged. Who was sent money, who was suppose to be sent money but failed.

