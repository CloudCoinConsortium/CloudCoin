<?php
/*Made by Sean worthington 7/9/2019
  This script shows you how to implement the "Payment Verifier" executable so that you can accept payments automatically.

  This is about a simple example as possible. It is very simulare to other languages. 

  To test this, you will need to send 1 CloudCoin to your Skywallet. The Payment Verifier will check for 1 CloudCoin with the memo of 
  "test".

  This assumes that you will put the logs into the root directory of your web server. You also need to put the paymentverifier.exe in the 
  same directory. 
*/

$timeout = 5; //Amount of times in seconds that the payment verifier will wait for response from the RAIDA.
$invoice = "test"; //The CloudCoin Wallet allows the user to specify a memo. This is important and used to identify the sender. 
$from = "change.skywallet.cc";//If your payments fail, the coins will be refunded here. So if you want your money back, replace this with your own address. 

//Change the from address into a serial number
$ip = gethostbyname($from.".");//add root to the senders account.
$to_sn = ip2long($ip) - 16777216; //Convert the senders account name to a serial number. 
$total_due = 1;//How many coins the customer should have sent. 

//Declare the location of your log files. The Paymentverifier.exe program will create a folder called "Log" there and track customer purhase attempts. 
  $Log_path = "C:\xampp\htdocs\";

//Declare location of your CloudCoin ID. This CLoudCoin stack file will be used to retrive or transfer your CloudCoin. Keep it safe. It is like a private key for a crypto currency. You don't want to lose it or have it stolen. 
  $Path_to_ID_coin = "C:\xampp\htdocs\1.CloudCoin.1.2..stack";

//Declare the name of the sub account within your skywallet where you want to move the coins (can be called anything)
  $move_to = "Received_from_".$from; 

//Put the "paymentverifier.exe" program in a place and start it. Note that we use the $total_due that we calculated
  $command = "C:\xampp\htdocs\paymentverifier.exe -getfrom=\"$invoice\" -payment=$total_due -refundto=$to_sn -logpath=\"$Log_path\" -idpath=\"$Path_to_ID_coin\"  -oldtag=\"$invoice\" -timeout=\"$timeout\" -newtag=\"$move_to\"";

  $results = exec($command); //This executes the command and puts the results in the results variable. 

 //Show the results
  echo "The command was: "+ $command + "<br>";
  echo "The results where: " + $results;


?>
