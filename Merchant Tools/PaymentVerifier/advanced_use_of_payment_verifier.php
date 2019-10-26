<?php
/*Made by Sean worthington 7/9/2019
This script shows you how to implement the "Payment Verifier" executable so that you can accept payments automatically.

This page works with the simply_payment_sample.html. The html page has a form and posts its data to this page. 

Note: Before the user submits the form on the html page, they should send one CloudCoin to your skywallet account 
with a invoice number that you specify. The exact amount is required. The payment verifies send the customer a refund immediatly
if the amount is inccorrect.

In this example, we pretend that the merchant 

Sample GET Request 
 https://raida18.cloudcoin.global/service/verify_payment?from=billy@Skywallet.cc&invoice=b26b&total_coins_sent=250
  *The "from" is the account that the requester wants to receive their refund if one is due (like the sender sends the wrong amount).
  *The "Invoice" is a random number that the sender made up so that you could recognize their payment.
  *The "total_coins_sent" is the number of coins the customer claims to have sent you. This may or may not be the correct amount. 

Extended GET Request that includes your merchant fields
 https://raida18.cloudcoin.global/service/verify_payment?from=billy@Skywallet.cc&invoice=b26b&total_coins_sent=250&book_88729=1&art_99882=2&product_998823=4
	*The book_88719 would be one of your books with its id (like if you were selling books). The =1 means the customer is buying one of these. 
	*The art_99882 would be if you were selling art product number 99882. The =2 means the customer is buying two of these. 
	*The product_998823 would be one of your products with its product ID. The =4 means the customer is buying four of these. 
 
*/

//Enable errors for development
ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);
error_reporting(E_ALL);

//Set timezone
date_default_timezone_set('Etc/UTC');
$date = date("Y-m-d H:i:s");

//Set timeout for request to run payment verifier
$timeout = 5; //5 seconds

//Get the manditory GET Variables
$invoice          = $_GET['invoice'];
$from             = $_GET['from'];
$ip = gethostbyname($from.".");//add root 
$to_sn = ip2long($ip) - 16777216;
$total_coins_sent = $_GET['total_coins_sent'];
$move_to = "from buy.html";//tag to give incomming income

//Redirect them back to the shopping page if the total coins sent is zero
if($total_coins_sent == "0"){ header('Location:buy.html');}

//Get your custom GET Variables
$book_88719       = $_GET['book_88719'];
$art_99882   	  = $_GET['art_99882'];
$product_998823	  = $_GET['product_998823'];

//Test to see if the prices of the items ordered match the CloudCoins sent. 

//pretend this is a database of products and prices:
//------Item ----+---Price----//
$price_book_88719 = 24;//in CloudCoins
$price_art_99882 = 240;
$price_product_998823 = 1;

//Calculate Total due:
$total_due = ($book_88719 * $price_book_88719) + ($price_art_99882 * $art_99882) + ($product_998823 * $price_product_998823);


//Declare the location of your log files. The Paymentverifier.exe program will create a folder called "Log" there and track customer purhase attempts. 
$Log_path = "C:\\Invoices\\OnlineSales";

//Declare location of your CloudCoin ID. This CLoudCoin stack file will be used to retrive or transfer your CloudCoin. Keep it safe. It is like a private key for a crypto currency. You don't want to lose it or have it stolen. 
$Path_to_ID_coin = "C:\\CloudCoin\\Accounts\\Change\ID\\1.CloudCoin.1.2..stack";

//Declare the name of the sub account within your skywallet where you want to move the coins (can be called anything)
$move_to = "Received_from_".$from; 

//Put the "paymentverifier.exe" program in a place and start it. Note that we use the $total_due that we calculated
$command = "C:\\xampp\htdocs\production_go\PaymentVerifier\paymentverifier.exe -getfrom=\"$invoice\" -payment=$total_due -refundto=$to_sn -rootpath=\"$Log_path\" -idpath=\"$Path_to_ID_coin\" -timeout=\"$timeout\" -newtag=\"$move_to\"";

//die( $command );
$Results = exec($command);
//$Results = "fail";
//The PaymentVerifier will now check to see if you have recieved payment. If the amount is wrong, it will refund the money.
//If the amount is true it will keep the money and create a sub account with the Move To name you specified. 

//Check results to see if the payment was verified.
if (strpos($Results, 'success') !== false) 
{
	//The payment send to your invoice matched the total due. 
	//You may now send the customer their product
	//Tell the customer that the payment has been verified:
	
	

	//Redirect to Success Page or Send JSON
	
	
	//Success Page
		header('Location: success.html?book=' . $book_88719 . '&art=' . $art_99882 . '&product=' . $product_998823 );
		
	
	//JSON: use if you are just goig to reply the outcome. 
	//$formattedRename= str_replace("[","",$Results);//Format the JSON without the square brakets. 
    //echo str_replace("]","",$formattedRename);

}
else
{
	//Sorry, the payment was not in your skywallet. Redirect to a fail page or return ajax
	     header('Location: fail.html');
	//Return JSON
		// $formattedRename= str_replace("[","",$Results);
		// echo str_replace("]","",$formattedRename);
		// echo $command;
	
}//End if else success


?>
