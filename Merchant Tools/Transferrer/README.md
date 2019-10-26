# Transferrer

Uses the ID key to tell the Trusted Transfer to move coins from its account to the account of another. 

It assumes that the RAIDA is ready and echoer was checked by the caller. 

It must know what coins it is to transfere. This requires Show.

## Flags

* fromTag: A string that specifies the sender's envelope/sub account/tag that the coins will be transfered out of. 

* amount: An int that tells amount of CloudCoins to be transfered to another person. 

* receiverID: A string that is aSerial number, IP or Account Name of the receiver. So 16777216 or 1.255.255.255 or Sean.CloudCoin.global

* tag: A string memo that will be included with the coins transfered so the receiver know where they came from. 

* logpath: A string that shows the path to the log folder

* transactionlogpath: A string that shows the path to the transaction log folder

* idpath: A string that shows the path to the id Coin

* timeout: A integer that says how many seconds the Transferer will wait for RAIDA to respond. 

## Sample usage:
```
transfer.exe -amount=236 -logto="C:\c c\Logs\Transfer" -id=C:\cc\ID\idcoin.stack -receiverID=16777216 -fromTag="change" -tag="ExportedCoins" -transactionlogpath=C:\user\bill\CloudCoin\Skywallet\Billy.Skywallet.cc\
```

## What the software does
Before the Transferrer is called, the caller should check that the RAIDA Echo works and that enough RAIDA are available.  

1. See if the receiverID is a decimal, IP Adress or account name such as billy.skywallet.cc. 
  case decimal: Validate it.
  case IP: Convert it to decimal
  case Account Name: Do a DNS lookup and and change the IP to the decimal. 

2. Validate other flags. 

3. Load the ID coin.

4. Show Coins in the Skywallet.
   The program should Pick two RAIDA by number. Then call the show service on each of the. These coins shoud be compared and only coins that exist in both arrays of coins should be used. 
   
5. Figure out if there are enough coins to send. If not, stop and send an error. 

6. See if there is exact change or if changes needs to be made. If chnage is needed, use the Transfer With Change Service. 
Otherwise you the Transfer service. NOTE: THE TRANSFER WITH CHANGE SERVICE CAN BE ADDED LATER. TO BEGIN WITH, IF THERE IS NO CHANGE, RESPOND WITH AN ERROR That it did not have the exact change. 

7. Call the Transfer Service. Check to see that at least 20 responded success. If less than 20, report money not fully send. 

8. Make sure all is logged. Who was sent money, who was suppose to be sent money but failed. 

## Example Call to the Transfer Service:
```
/**
 * https://raida0.cloudcoin.global/service/transfer?
 * sn=16777216&
 * nn=1&
 * an=8ae06de0f9ce4917b3309df71570f92c&
 * pan=8ae06de0f9ce4917b3309df71570f92c&
 * denomination=1&
 * to_sn=16225354&
 * nns[]=1&
 * nns[]=1&
 * nns[]=1&
 * sns[]16777214&
 * sns[]16777215&
 * sns[]16777216&
 * tag=8d0caf3
 */

```
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

There must be a differnt log file for every day. The names of the log files will be formatted like: 
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
