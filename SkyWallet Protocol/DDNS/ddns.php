<?php

/*
This program allows someone to create a DNS record in Skywallet.cc. It requires that the user prove that they own the number
that they want to associate with the A record. 
Written May 13th 2019
By Sean Worthington
*/

$thisServerName = "Skywallet";
$trustedDomainName = "cloudcoin.global";
$date = date("Y-m-d H:i:s");

//SEE IF GET VARIABLES ARE NOT THERE
if (!isset($_GET['ticket']) || !isset($_GET['sn']) || !isset($_GET['raidanumber']) || !isset($_GET['username'])) 
{ 
        $date = date("Y-m-d H:i:s");
        die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide a ticket, sn, username and raida_numer.","time":"'.$date.'"}');
}


//Get the Username
$username = $_GET['username'];
//Get the SN
$sn = intval($_GET['sn']);

//Get the RAIDA number
$raidaNumber = intval($_GET['raidanumber']);
//Get the Ticket Number
$ticket = $_GET['ticket'];
//Validate
if(!is_valid_domain_name($username) || strpos($username, ".") !== false){//Name is bad
	$date = date("Y-m-d H:i:s");
        die('{"server":"' . $thisServerName . '","status":"fail","message":"Username is not valid","time":"'.$date.'"}');
	
}//if username fucked
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
 
        
        die('{"server":"' . $thisServerName . '","status":"error","message":"Raida Number: The unit\'s serial number was out of range.","time":"' . $date . '"}');  
}

if( strlen($ticket) != 44 ){  
die('{"server":"'.$thisServerName.'","status":"error","message":"Ticket: The ticket was not 44 characters long.","time":"'.$date.'"}');
}//end if 

//Check Ticket is valid
//GET DATA FROM WEB SERVERS USING MESSAGES
$raw_from_server_1 = curl_get_contents("https://raida$raidaNumber.$trustedDomainName/service/hints?rn=$ticket");
//echo $raw_from_server_1;
//die("https://raida$raidaNumber.$trustedDomainName/service/hints?rn=$ticket");

if( !$raw_from_server_1 ){//Server 1 did not return
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
$sn = $sn + 16777216;//Makes the network be 1
$dnsipgeldi = long2ip( $sn );//String will now start with 

//If good talk to the CloudFlare API
 /* Cloudflare.com | APİv4 | Api Ayarları */
    $apikey = '2a2efe54f GET FROM SEAN 0d6b263266a'; // Cloudflare Global API
    $email = 'cloudcoin@protonmail.com'; // Cloudflare Email Adress
    $domain = 'skywallet.cc';  // zone_name // Cloudflare Domain Name
    $zoneid = '561ed GET FROM SEAN 16fd2cdc234'; // zone_id // Cloudflare Domain Zone ID


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
    	    // 'Content-Type: multipart/form-data; charset=utf-8',
    	    'Content-Type:application/json',
    		'purge_everything: true'
    		
    		));
    	
    		// -d curl parametresi.
    		$data = array(
    		
    			'type' => 'A',
    			'name' => ''.$username.'',
    			'content' => ''.$dnsipgeldi.'',
    			'zone_name' => ''.$domain.'',
    			'zone_id' => ''.$zoneid.'',
    			'proxiable' => 'true',
    			'proxied' => false,
    			'ttl' => 1
    		);
    		
    		$data_string = json_encode($data);

    		curl_setopt($ch, CURLOPT_POST, true);
    		curl_setopt($ch, CURLOPT_POSTFIELDS, $data_string);	
    		//curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($data_string));

    		$sonuc = curl_exec($ch);

                 // If you want show output remove code slash.
        if( strpos($sonuc, 'success":true') !== false ){
            echo '{"status":"success","errors":[],"messages":[],"result":"' .$username.' created"}';
        }else{       
            
           // $sonuc = str_replace( '"success":true', '"status":"fail"', $sounce);
           $sonuc = str_replace( "success" , "status", $sonuc);
		  $sonuc = str_replace( "false" , "\"fail\"", $sonuc);
		  
		  print_r($sonuc);
}
    		curl_close($ch);

function curl_get_contents($url)
{
    $ch = curl_init();

    curl_setopt($ch, CURLOPT_HEADER, 0);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_URL, $url);

    $data = curl_exec($ch);
    curl_close($ch);

    return $data;
}
	 
	 
function is_valid_domain_name($domain_name)
{
    return (preg_match("/^([a-z\d](-*[a-z\d])*)(\.([a-z\d](-*[a-z\d])*))*$/i", $domain_name) //valid chars check
            && preg_match("/^.{1,253}$/", $domain_name) //overall length check
            && preg_match("/^[^\.]{1,63}(\.[^\.]{1,63})*$/", $domain_name)   ); //length of each label
}
?>
