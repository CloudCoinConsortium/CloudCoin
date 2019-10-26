# Accepting Payments

You can quickly and reliably receive CloudCoins so you can automatically fill your orders. 

Your customers will send CloudCoins to your Skywallet address using the CloudCoin Wallet software or one of our web payment systems. 

The customer will send your servers a message specifying what they bought, how many CloudCoins they sent you, and what "memo" was included. 

Your servers with then go to your Skywallet account and see if a transfer has been received with the memo and amount specified by the customer. 

For this, your server can call on the Payment Verifier to do that work. The Payment Verifier will return a true or false. 

Payment Verifier is a executable program that runs on either Windows of Linux. It can be called from any langague that is able to call executable. So, you can call the Payment Verifier from PHP, C#, JAVA, C++, Ruby, almost anything! 

# Payment Verifier

The Payment Verifier sees if a payment has been received based on the memo (aka "tag" or "envelope"). Then, if the payment has been receieved, it will rename the memo so you will not that it has been handled. If the customer sent too many or too few CloudCoins, the Payment Verifier will return the coins to the sender. 

You can download the compiled Payment Verifier at: https://cloudcoinconsortium.com/exe/paymentverifier.exe

When you call the Payment Verifier, you will need to include some command line arguments also called "flags". 

## Sample Call to the Payment Verifier Executable:
```
paymentVerifier.exe -timeout=5 -payment_memo=sean4 -payment=100 -refundto=1371486 -logpath="C:\Logs\PaymentVerifier" -idpath="C:\cc\Accounts\Change\ID\1.CloudCoin.1.2..stack" -transactionlogpath="C:\user\bill\CloudCoin\Skywallet\bill.skywallet.cc"

```
## Flags Required by the Payment Verifier Executable. 
* payment_memo: A string that specifies the tag/envelope/memo that the sender provided
* payment:  A string that is an integer showing how many CloudCoins the user claimed to have sent. 
* newtag: A string tag that will be given to the coins that have been receieved.  
* logpath: A string specifiying the path to the Log files. 
* transactionlogpath: A string specifiying the path to the Log files. 
* idpath: A string specifying the path to the coin that will be used to access the receiver's account. 
* timeout A string that describes the number of seconds that the program will wait for RAIDA to respond. 
* refundto: A string that is a Serial number, IP or Account Name of the receiver. So 16777216 or 1.255.255.255 or Sean.CloudCoin.global

## How it works
Before the PaymentVerifier is called, the caller should check that the RAIDA Echo works and that enough RAIDA are available.

1. See if the refundto is a decimal, IP Adress or account name such as billy.skywallet.cc. case decimal: Validate it. case IP: Convert it to decimal case Account Name: Do a DNS lookup and and change the IP to the decimal.

2. Validate other flags.

3. Load the ID coin.

4. Show Coins in the Skywallet. THIS IS THE HARD PART: Then program must call the verify_payment service on all the RAIDA. 

## Example GET asking to verity the amount and envelope name of the claimed payment.
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

## Logging
There are two types of logs. The event log and the transaction log. They each have their own log locations. 

There must be a differnt log file for every day. The names of the log files will be formatted like:

```
"Jan.26.2019.PaymentVerifier.txt"
```

The records inside the log file will include the hour and minute, the SenderID, amount, receiverID, fromTag, tag. Like this: NOTE the use of millitary time.
```
11:45 2 verified that billy.skywallet.cc sent 235 coins to tag "From Billy". New tag: "change"
12:32 2 DID NOT verify that billy.skywallet.cc sent 235 coins to tag "From Billy". Too many coins sent.
15:27 2 DID NOT verify that sean.cloudcoin.global sent 10000 coins from tag "change". No coins sent.
```
Transaction Log name: "transactions.txt"

The format is: 
Tag, date, Total_expected, Withdraw, Total, new_tag

* The "Withdraw" is always empty for Payment Verifier. 

* The "Total" is figured out by doing a show call to the RAIDA to find the total amount in the account. 

Contents of Transaction.txt file:
```

For New Sneakers,2019-Sep-13 8-15PM,250,,250,change
I am a customer,2019-Sep-16 9-20PM,6000,,6250,change
order 29938,2019-Sep-16 9-23PM,500,,6750,change



```
