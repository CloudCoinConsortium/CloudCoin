<?php


/*
ddns.php

This program allows someone to create a DNS record in Skywallet.cc. It requires that the user prove that they own the number
that they want to associate with the A record. 
Written May 22th 2020
By Sean Worthington

Sample Use:
https://yourhost.com/ddns.php?nn=1&sn=1358923&raidanumber=2&ticket=12345678901234567890123456789012345678901234username=Billy1234&

*/

header("Access-Control-Allow-Origin: *");

$thisServerName = "Skywallet";
$trustedDomainName = "cloudcoin.global";
$date = date("Y-m-d H:i:s");

 /* Cloudflare.com | APİv4 | Api Ayarları */
$apikey = 'ce5eYour CLoudFlare API here 6e9d107f6'; // Cloudflare Global API
$email = 'youremailnmail.com'; // Cloudflare Email Adress
$domain = 'yourdomain.com';  // zone_name // Cloudflare Domain Name
$zoneid = '561ed8e your zone id 2cdc234'; // zone_id // Cloudflare Domain Zone ID


//SEE IF GET VARIABLES ARE NOT THERE
    if ( !isset($_GET['sn']) ) 
    { 
            $date = date("Y-m-d H:i:s");
            die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide a sn","time":"'.$date.'"}');
    }
    
       if ( !isset($_GET['raidanumber']) ) 
    { 
            $date = date("Y-m-d H:i:s");
            die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide  raida_numer.","time":"'.$date.'"}');
    }
    
    
       if (  !isset($_GET['username'])) 
    { 
            $date = date("Y-m-d H:i:s");
            die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide username.","time":"'.$date.'"}');
    }
    
    



    $username = $_GET['username'];
    $sn = intval($_GET['sn']);
    $raidaNumber = intval($_GET['raidanumber']);
    $ticket = $_GET['ticket'];

//VALIDATE
    if(!is_valid_domain_name($username) || strpos($username, ".") !== false){//Name is bad
    	  $date = date("Y-m-d H:i:s");
        die('{"server":"' . $thisServerName . '","status":"fail","message":"Username is not valid","time":"'.$date.'"}');
    }//if username invalide
    
    if ($sn > 16777216 || $sn < 1) 
    { 
        if (isset($_GET['b'])) //is it a brief request?
        {
            die('fail');
        } 
    	else 
    	{   
            $date = date("Y-m-d H:i:s");
            die('{"server":"' . $thisServerName . '","status":"error","sn":"' . $sn . '","message":"SN: The unit\'s serial number was out of range.","time":"' . $date . '"}');  
        }
    }

    if ($raidaNumber > 24 || $raidaNumber < 0) 
    { 
        die('{"server":"' . $thisServerName . '","status":"error","message":"Raida Number was out of range.","time":"' . $date . '"}');  
    }

//WORK ON TICKET

       if ( !isset($_GET['ticket'])) 
    { 
            $date = date("Y-m-d H:i:s");
            die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide a ticket","time":"'.$date.'"}');
    }

    if( strlen($ticket) != 44 )
    {  
        die('{"server":"'.$thisServerName.'","status":"error","message":"Ticket: The ticket was not 44 characters long.","time":"'.$date.'"}');
    }//end if 

    $raw_from_server_1 = file_get_contents("https://raida$raidaNumber.$trustedDomainName/service/hints?rn=$ticket");
    //echo "Response from RAIDA$raidaNumber" .$raw_from_server_1;

    //die("https://raida$raidaNumber.$trustedDomainName/service/hints?rn=$ticket");

if( !$raw_from_server_1 ){//Server did not return anything
	die( '{"server":"'.$thisServerName.'","status":"error","message":"Connection: Could not connect to RAIDA '.$raidaNumber.'.","time":"'.$date.'"}');
}//end if fail


$parts_1 = explode(":", $raw_from_server_1);

$sn_from_server_1 = $parts_1[0];

$seconds_from_server_1 = intval( $parts_1[1] );

switch( $sn_from_server_1 ){
	case -1: 
	die( '{"server":"'.$thisServerName.'","status":"error","message":"Remote Ticket: RAIDA '.$raidaNumber.' database said invalid ticket.","time":"'.$date.'"}');
	break;
	case -2: 
	die('{"server":"'.$thisServerName.'","status":"error","message":"Remote Ticket: No Ticket ( '.$ticket.') found on Server '.$raidaNumber.'.","time":"'.$date.'"}');
	break;
	case -3: 
	die('{"server":"'.$thisServerName.'","status":"error","message":"Remote Ticket: RAIDA '.$raidaNumber.' said invalid ticket.","time":"'.$date.'"}');
	break;
}//end switch


//Turn SN into IP address. 
//$sn = $sn + 16777216;//Makes the network be 1
$sn = $sn + 16777216; //Makes the first octet of the dotted decimal IP address be 1 
$dnsipgeldi = long2ip( $sn );//String will now start with 


/* SEE IF IP ADDRESS IS ALREAYD IN THE DNS */
//Do not allow the same account to have two ips. 


 $get = curl_init("https://api.cloudflare.com/client/v4/zones/$zoneid/dns_records?type=A&content=$dnsipgeldi&match=all");
			curl_setopt($get, CURLOPT_RETURNTRANSFER, true);
    		curl_setopt($get, CURLOPT_FOLLOWLOCATION, true);
    		curl_setopt($get, CURLOPT_SSL_VERIFYPEER,false);
    		curl_setopt($get, CURLOPT_CUSTOMREQUEST, "GET");                                                                     
    		curl_setopt($get, CURLOPT_HTTPHEADER, array(
														'X-Auth-Email: '.$email.'',
														'X-Auth-Key: '.$apikey.'',
														'Cache-Control: no-cache',
														'Content-Type:application/json',
														'purge_everything: true'
														));

        $return = curl_exec($get);
	
        curl_close($get);
 
		// See if the IP address is already used. 
        if( contains($return, 'count": 0') == "false" || contains($return, 'count":0') == "false"  )
		{
		   //The number of DNS recoreds that match are zero
            
		}
		else
		{
		    die('{"server":"' . $thisServerName . '","status":"fail","message":"IP address already has a DNS name.","time":"'.$date.'"}');    
		}
		
   /*   */
    //echo "Your ip is not used";

    // A-record oluşturur DNS sistemi için.
    		$ch = curl_init("https://api.cloudflare.com/client/v4/zones/".$zoneid."/dns_records");
    		curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    		curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
    		curl_setopt($ch, CURLOPT_SSL_VERIFYPEER,false);
    		curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "POST");                                                                     
    		curl_setopt($ch, CURLOPT_HTTPHEADER, array(
														'X-Auth-Email: '.$email.'',
														'X-Auth-Key: '.$apikey.'',
														'Cache-Control: no-cache',
														'Content-Type:application/json',
														'purge_everything: true'
														));
        	
    // -d curl parametresi.
        $data = array(
            'type' => 'A',
            'name' => '' . $username . '',
            'content' => '' . $dnsipgeldi . '',
            'ttl' => 1,
            'priority' => 10,
            'proxied' => false
        );
    		
$data_string = json_encode($data);
//echo "$data_string is : ".$data_string;
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $data_string);
//curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($data_string));

$sonuc = curl_exec($ch);
 //echo "Sonuc is ".$sonuc;


//These CloudFlare fuckers keep changing the protcol by adding a space or changing the word. Check for all
if ( strpos($sonuc, 'status":true') !== false) {
    die( '{"status":"success","errors":[],"messages":[],"result":"' . $username . ' created"}');
}

if (strpos($sonuc, 'status": true') !== false) {
    die ('{"status":"success","errors":[],"messages":[],"result":"' . $username . ' created"}');
}

if (strpos($sonuc, 'success": true') !== false) {
    die ('{"status":"success","errors":[],"messages":[],"result":"' . $username . ' created"}');
}

if (strpos($sonuc, 'success":true') !== false) {
    die ('{"status":"success","errors":[],"messages":[],"result":"' . $username . ' created"}');
}


    $sonuc = str_replace("success", "status", $sonuc);
    $sonuc = str_replace("false", "\"fail\"", $sonuc);
    
    print_r($sonuc);

curl_close($ch);

function curl_get_contents($url)
{
    $ch = curl_init();

    curl_setopt($ch, CURLOPT_HEADER, 0);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_URL, $url);

    $data = curl_exec($ch);
    curl_close($ch);

    return  $data;
}
	 
	 
function is_valid_domain_name($domain_name)
{
    return (preg_match("/^([a-z\d](-*[a-z\d])*)(\.([a-z\d](-*[a-z\d])*))*$/i", $domain_name) //valid chars check
            && preg_match("/^.{1,253}$/", $domain_name) //overall length check
            && preg_match("/^[^\.]{1,63}(\.[^\.]{1,63})*$/", $domain_name)   ); //length of each label
}

function contains( $haystack, $needle )
{
     if ( strpos( $haystack, $needle ) !== false ) 
    {
        return 'true';
    }
        return 'false';
}//end contains


?>
