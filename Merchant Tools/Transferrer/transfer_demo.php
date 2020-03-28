
<?php

/*

WARNING: THIS CODE COULD BE USED TO TRANSFER ALL FUNDS FROM YOUR SKYWALLET TO ANOTHER PERSON SO KEEP IT SECURE!
This code shows how you can call the "Transferrer" Servant from a php page. 
The transferrer servant moves coins from one Skywallet to another skywallet.

@Author Sam Leary
@Version 10/2/2019


Sample Use:
https://myserver.com/transfer_demo.php?invoice=88322&from=sean.Skywallet.cc&amount=22&to=Billy.Skywallet.cc&password=9812481556564824824856

*/
//Show errors
ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);
error_reporting(E_ALL);

//the tag that the sent coins will have
$invoice          = $_GET['invoice'];
//which tag the service will send from 
$from             = $_GET['from'];
// how much the service will send
$amount           = $_GET['amount'];

//The Receiver ID
$receiverID       = $_GET['to'];

$password = $_GET['password'];

//We password protect this demo just to make sure this is not called from an outside person. 
//You will need to protect this php code from being called
if ($password == '9812481556564824824856'){

  
  //Now you can call the transferrer.exe and get its response. 
  //Download transferrer at: http://raida18.cloudcoin.global/exchange/teller/transferrer.zip
  //For possible responses, see https://github.com/worthingtonse/servants/tree/master/Teller/Transfer
  //To see the source code for Transferrer see: https://github.com/worthingtonse/servants/blob/master/Teller/Transfer/transferrer.go
  /*
  Flags
    fromTag: Every account on the Skywallet has envelopes. These are sub accounts and each envelope has is labeled with a tags. A tag is a string that identifies an envelope. This flag specifies which of the sender's envelope/sub account/tag that the coins will be transfered out of. Coins are almost alwasy transfered from the "change" envelope which is used as a general fund. 

    amount: An int that tells amount of CloudCoins to be transfered from the sender's account to the receivers account.

    receiverID: A string that is aSerial number, IP or Account Name of the receiver. So IDs like 16777216 or 1.255.255.255 or Sean.CloudCoin.global are possible. The transferrer will always convert and ID into a decimal. 

    tag: The Sender is allowed to create an envelope in the Receiver's account. The tag is what that new envelope will be named. This allows the Sender to tell the recever where to check to find their payment. Tags are often random strings that uniquely identify the senders payment.    

    logpath: A string that shows the path to the log folder where the Transferrer will make log files that can be used in troubleshooting problems. 

    transactionPath: A string that shows the path to the transaction log folder. This is where client software can look to see transactions that the Transferrer has completed to show the user.  

    idpath: To transfer CloudCoin, the Sender must have a ID coin. This is the path to that coin. You will need to mark a coin as an ID coin. Use the CloudCoin Wallet to create sending and receiving Skywallet accounts to make testing easier. 

    timeout: A integer that says how many seconds the Transferer will wait for RAIDA to respond. I recomend about 10 seconds. 

  */
  
  $command = "C:\\xampp\\htdocs\\exchange\\teller\\transferrer.exe -timeout=5 -transactionPath=TransactionLog\\ -logpath=Logs\\transferrer -idpath=C:\\CloudCoin\\Accounts\\Change\\ID\\1.CloudCoin.1.11..stack -receiverID=$receiverID -fromtag=$from -tag=$invoice -amount=$amount";
  $results = exec($command);  
  echo $results;
  
}else{
  echo "incorrect password";
}




?>
