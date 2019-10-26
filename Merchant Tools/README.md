# Accepting Payments

You can quickly and reliably receive CloudCoins and automatically fill orders. 

Your customers will send CloudCoins to your Skywallet using their CloudCoin Wallet software. 

The customers will then send your servers a message specifying what they bought, how many CloudCoins they sent you, and what the "memo" was. 

Your servers will then check your Skywallet account and see if you received a transfer of that amount with the same memo.

Your server can call on our Payment Verifier exe to do all the work. Just tell the Payment Verifier, the amount, memo and a few other details. The Payment Verifier will return a true or false. 

The Payment Verifier is a executable program that runs on either Windows, Linux or Mac. It can be called from any langague that is able to call executables such as PHP, C#, JAVA, C++, Ruby, almost anything! 

# Payment Verifier

The Payment Verifier sees if a payment has been received based on the memo (aka "tag" or "envelope"). Then, if the payment has been receieved, it will rename the memo so you will not be confused with other orders. If the customer sent too many or too few CloudCoins, the Payment Verifier will return the coins to the sender. The Payment Verifier expects the customer either send the exact amount or resend. 

You can download the compiled Payment Verifier at: https://cloudcoinconsortium.com/exe/paymentverifier.exe

When you call the Payment Verifier, you will need to include some command line arguments also called "flags". 

## Sample Call to the Payment Verifier Executable:
```
paymentVerifier.exe -timeout=5 -payment_memo=sean4 -payment=100 -refundto=1371486 -logpath="C:\Logs\PaymentVerifier" -idpath="C:\cc\Accounts\Change\ID\1.CloudCoin.1.2..stack" -transactionlogpath="C:\user\bill\CloudCoin\Skywallet\bill.skywallet.cc"

```
## Flags Required by the Payment Verifier Executable. 
* payment_memo: A string that specifies the tag/envelope/memo that the sender specified.
* payment:  An integer showing the amount of CloudCoins the user claimed to have sent. 
* newtag: A string tag that will be given to the coins that have been receieved. Here you can tag the coins with a name. Like "Received". If you use the world "change" the coins will be used to create change for your customers.  
* logpath: A string specifiying the path to the Log files.
* transactionlogpath: A string specifiying the path to the tranaction log. These log files can be read by the CloudCoin Wallet. 
* idpath: A string specifying the path to the coin that will be used to access your skywallet account. 
* timeout A string that describes the number of seconds that the program will wait for RAIDA to respond. 5 seconds is good but if you have slow internet, you may need more time. 
* refundto: A string that is a Serial number, IP or Account Name of the receiver. So 16777216 or 1.255.255.255 or Sean.CloudCoin.global are all good.


## Logging
There are two types of logs. The event log and the transaction log. They each have their own log locations. 

There must be a differnt log file for every day. The names of the log files will be formatted like:

```
"Jan.26.2019.PaymentVerifier.txt"
```

The records inside the log file will include the hour(0-23) and minute (0-59), the SenderID, amount, receiverID, fromTag, new tag.
Here we see the merchant with Skywallet 2345 check for payments. 
```
11:45 2345 verified that billy.skywallet.cc sent 235 coins to tag "From Billy". New tag: "change"
12:32 2345 DID NOT verify that billy.skywallet.cc sent 235 coins to tag "From Billy". Too many coins sent.
15:27 2345 DID NOT verify that sean.cloudcoin.global sent 10000 coins from tag "change". No coins sent.
17:45 2345 verified that Jen.skywallet.cc sent 885 coins to tag "8522211225". New tag: "May.14.2019.Receipts"
```
Transaction Log name: "transactions.txt"

The format is: 
Tag, date, Total_expected, Withdraw, Total, new_tag

* The "Withdraw" is always empty for Payment Verifier. 

* The "Total" is figured out by doing a show call to the RAIDA to find the total amount in the account. 

Contents of Transaction.txt file:
```
For New Sneakers, 2019-Sep-13 8-15PM, 250, , 250,change
I am a customer, 2019-Sep-16 9-20PM, 6000, , 6250, May.14.2019.Receipts
order 29938, 2019-Sep-16 9-23PM, 500,,6750, May.14.2019.Receipts
```
