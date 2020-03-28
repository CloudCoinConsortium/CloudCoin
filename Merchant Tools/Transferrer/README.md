# Transferer

The Transferer allows you to send CloudCoins from your your SkyWallet account to the SkyWallet account of your customers. Naturally, it requires that you and your customer have a SkyWallet. 

It assumes that the RAIDA is ready and echoer was checked to see that you are connected to the RAIDA. 

The Transferer needs access to the ID key so it can use the account. 

The Transferrer is an execuatble program that can run on Windows, Linux or Mac by any programming language that can call an executable. 

## Sample usage:
```
transfer.exe -amount=236 -logpath="C:\user\merchant\Logs\Transfer" -idpath=C:\user\merchant\ID\idcoin.stack -receiverID=16777216 -fromTag="generalfund" -tag="From Merchant for Purchase of Shoes" -transactionlogpath=C:\user\merchant\CloudCoin\Skywallet\merchant.Skywallet.cc\
```

## Flags

* fromTag: A string that specifies from what pool of coins labeled with a memo that you want the coins to come from. 

* amount: The amount of CloudCoins to be transfered to the other account. 

* receiverID: A string that is Serial number, IP or Account Name of the receiver. So 16777216 or 1.255.255.255 or Sean.CloudCoin.global will work. 

* tag: A string memo that will be included with the coins transfered so the receiver know where they came from. 

* logpath: The path to the folder that the log files will be created in.

* transactionlogpath: A string that shows the path to the transaction log folder

* idpath: A string that shows the path to the id Coin that is used to access your account.

* timeout: A integer that says how many seconds the Transferer will wait for RAIDA to respond. 5 seconds should be good. It takes more time for slower internet connection and larger transferrs. 

## Output to command line or calling process
A successful output will look like the following 
```
{"status":"success","message":"Successfully sent 5 coins to sean12345.skywallet.cc"," Execution Time = 2.2229534s","time":"2019-10-4 18:10:44"}
```
if the transfer fails the output will look like the following:
```
{"status":"failed","message":"could not fully send."," Execution Time = 2.2229534s","time":"2019-10-4 18:10:44"}
```


## Log Files

There will be a differnt log file for every day. The names of the log files will be formatted like: 
```
"Jan.26.2019.Transfers.txt"
```
The records inside the log file will include the hour and minute, the SenderID, amount, receiverID, fromTag, tag. Like this: NOTE the use of millitary time. 
```
11:45 2 sent billy.skywallet.cc 235 coins from tag "change" to tag "Your Order 88734"
12:32 2 sent jill.skywallet.cc 100 coins from tag "change" to tag "Your Order 88735"
15:27 2 sent sean.cloudcoin.global 10000 coins from tag "change" to tag "Your Order 88736"
```

Transaction Log name: "transactions.txt"

The format is: Tag, date, Deposited, Withdraw, Total, new_tag

The "Deposited" is always empty for Transfer unless the sn and to_sn are the same. Then the Deposited and Withdrawn will be the same.

The "Total" is figured out by doing a show call to the RAIDA to find the total amount in the account.

Contents of Transaction.txt file:
```
For New Sneakers,2019-Sep-13 8-15PM,,500,8000,change
I am a customer,2019-Sep-16 9-20PM,,2000,6000,change
order 29938,2019-Sep-16 9-23PM,,700,5300,change
```
