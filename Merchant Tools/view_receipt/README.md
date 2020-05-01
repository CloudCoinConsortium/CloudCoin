# View Receipt

View Receipt allows you and your customer to see that you have received the payment. This is a simple version of Verify Payment. 
The differance is that View Receipt will not change the envelope that the received money is in. Also, the view_receipt requires that the 
sender use a GUID as the memo/envelope name. 

This software is written in Go and it will run on any plateform. Here we see it being called as an executable on Windows:

```
view_receipt.exe -expected_amount=236 -logpath="C:\merchant\receipts\" -guidtag="b5d8b40d02154fd69f7003978e790fa3" -timeout=5
```
## Flags
*
* expected_amount: The amount of CloudCoins to be transfered to the other account.
* logpath: The path to the folder that the log files will be created in.
* guidtag: The guid that the customer included in the tag or memo when they submitted payment. 
* timeout: A integer that says how many seconds the Transferer will wait for RAIDA to respond. 5 seconds should be good. It takes more time for slower internet connection and larger transferrs.

## Output to command line or calling process
```
{
	"server": "raida18",
	"total_received": 454,
	"serial_numbers": "1308799,1308800,1308868,1308983,14339990,14340140,16458252",
	"version": "2020-02-13",
	"time": "2020-04-20 23:31:40"
	"execution_time":.023
}
```
