<?php
/*Made by Samuel Leary 6/9/2020
  This script shows you how to implement the "view_receipt" executable so that you can accept payments automatically.

  To test this, we will pretent that a customer sent you some CloudCoins with a memo of "623e88186f3c4a4694c02230abe72666".  
  To test this, send one CloudCoin to your Skywallet. The Payment Verifier will check for 1 CloudCoin by looking for the GUID. 
 
  This assumes that you will put the logs into the root directory of your web server. You also need to put the view_receipta.exe in the 
  same directory. 
*/

$timeout = 5; //Amount of times in seconds that the payment verifier will wait for response from the RAIDA.
$guid = "623e88186f3c4a4694c02230abe72666"; //this is the GuidTag of the sent coins.
$account = billy.Skywallet.cc; //The senders 
$total_due = 1;


//TODO Change the $account into a integer between 1 and 16,777,216

//Declare the location of your log files. The Paymentverifier.exe program will create a folder called "Log" there and track customer purhase attempts. 
$Log_path = "C:\xampp\htdocs\";

//Declare location of your CloudCoin ID. This CLoudCoin stack file will be used to retrive or transfer your CloudCoin. Keep it safe. It is like a private key for a crypto currency. You don't want to lose it or have it stolen. 
  $Path_to_ID_coin = "C:\xampp\htdocs\1.CloudCoin.1.2..stack";

//Put the "paymentverifier.exe" program in a place and start it. Note that we use the $total_due that we calculated
  $command = "C:\xampp\htdocs\view_receipt.exe -guidtag=\"$guid\" -expected_amount=$total_due -account=$account -logpath=\"$Log_path\" -timeout=\"$timeout\"";

  $results = exec($command); //This executes the command and puts the results in the results variable. 

 //Show the results
  echo "The command was: "+ $command + "<br>";
  echo "The results where: " + $results;


?>
