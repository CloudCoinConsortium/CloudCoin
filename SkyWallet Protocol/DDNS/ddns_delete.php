<?php

/*
ddns_delete.php
This file deletes an A record associated with a CloudCoin. The user must show tha they are the owner of the cloudcoin by
presenting a ticket. 
Written by Sean Worthington
Jan 6, 2020

Sample Use:
https://mydomain.com/ddns_delete.php?nn=2&sn=1358923&raidanumber=2&username=Billy1234&ticket=12345678901234567890123456789012345678901234
*/

ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);
error_reporting(E_ALL);

$dns_zone          = "Skywallet.cc";
$thisServerName    = "Skywallet";
$trustedDomainName = "cloudcoin.global";
$date              = date("Y-m-d H:i:s");

// Cloudflare.com | APİv4 | Api Variables 
$apikey = 'ce5ec1e GET THIS FROM SEAN d107f6'; // Cloudflare Global API
$email  = 'cloudcoin@protonmail.com'; // Cloudflare Email Adress
$domain = 'skywallet.cc'; // zone_name // Cloudflare Domain Name
$zoneid = '561ed8e89cf6c8a10aa8416fd2cdc234'; // zone_id // Cloudflare Domain Zone ID

//SEE IF GET VARIABLES ARE NOT THERE
if (!isset($_GET['ticket'])) {
    die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide a ticket.","time":"' . $date . '"}');
}

if (!isset($_GET['sn'])) {
    die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide an sn.","time":"' . $date . '"}');
}

if (!isset($_GET['raidanumber'])) {
    die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide a raidanumber.","time":"' . $date . '"}');
}

if (!isset($_GET['username'])) {
    die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide a username.","time":"' . $date . '"}');
}



//Get the Username
$username = $_GET['username']; //username is the dns name\
$username = $username . "." . $dns_zone;
$nn       = intval($_GET['nn']);
$sn       = intval($_GET['sn']);
$raidaNumber = intval($_GET['raidanumber']);
$ticket      = $_GET['ticket'];

//Validate
if (!is_valid_domain_name($username)) { //Name is bad
    die('{"server":"' . $thisServerName . '","status":"fail","message":"Username is not a valid domain name. It should have three parts seperated by periods. Like john.cloudcoin.com","time":"' . $date . '"}');
} //if username bad

if ($sn > 16777216 || $sn < 1) {
    
        die('{"server":"' . $thisServerName . '","status":"error","sn":"' . $sn . '","message":"SN: The unit\'s serial number was out of range.","time":"' . $date . '"}');
}

if ($raidaNumber > 24 || $raidaNumber < 0) {
    die('{"server":"' . $thisServerName . '","status":"error","message":"Raida Number: The unit\'s serial number was out of range.","time":"' . $date . '"}');
}

if (strlen($ticket) != 44) {
    die('{"server":"' . $thisServerName . '","status":"error","message":"Ticket: The ticket was not 44 characters long.","time":"' . $date . '"}');
} //end if 

//Check Ticket is valid
//GET DATA FROM WEB SERVERS USING MESSAGES

$raw_from_server_1 = curl_get_contents("https://raida$raidaNumber.$trustedDomainName/service/hints?rn=$ticket");
//echo $raw_from_server_1;
//die("https://raida$raidaNumber.$trustedDomainName/service/hints?rn=$ticket");

if (!$raw_from_server_1) { //Server 1 did not return
    die('{"server":"' . $thisServerName . '","status":"error","message":"Connection: Could not connect to RAIDA ' . $raidaNumber . '.","time":"' . $date . '"}');
} //end if fail


$parts_1 = explode(":", $raw_from_server_1);

$sn_from_server_1 = $parts_1[0];

$seconds_from_server_1 = intval($parts_1[1]);

switch ($sn_from_server_1) {
    case -1:
        die('{"server":"' . $thisServerName . '","status":"error","message":"Remote Ticket: RAIDA ' . $raidaNumber . ' database said invalid ticket.","time":"' . $date . '"}');
        break;
    case -2:
        die('{"server":"' . $thisServerName . '","status":"error","message":"Remote Ticket: No Ticket ( ' . $ticket . ') found on Server ' . $raidaNumber . '.","time":"' . $date . '"}');
        break;
    case -3:
        die('{"server":"' . $thisServerName . '","status":"error","message":"Remote Ticket: RAIDA ' . $raidaNumber . ' said invalid ticket.","time":"' . $date . '"}');
        break;
} //end switch




//Turn SN into IP address. 
$sn         = $sn + ($nn * 16777216); //Makes the network be 1
$dnsipgeldi = long2ip($sn); //String will now start with 

//If good talk to the CloudFlare API


//Find the id of the coin so the delete command can be issued. 
//die("https://api.cloudflare.com/client/v4/zones/".$zoneid."/dns_records?type=A&name=".$username."&content=".$dnsipgeldi);

// A-record oluşturur DNS sistemi için.
$ch = curl_init("https://api.cloudflare.com/client/v4/zones/$zoneid/dns_records?type=A&name=$username&content=$dnsipgeldi&page=1&per_page=20&order=type&direction=desc&match=all");

//	$ch = curl_init("https://api.cloudflare.com/client/v4/zones/$zoneid/dns_records?type=A&page=1&per_page=20&order=type&direction=desc&match=all");   


//DELETE zones/:zone_identifier/dns_records/:identifier

curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "GET");
curl_setopt($ch, CURLOPT_HTTPHEADER, array(
    'X-Auth-Email: ' . $email . '',
    'X-Auth-Key: ' . $apikey . '',
    'Cache-Control: no-cache',
    // 'Content-Type: multipart/form-data; charset=utf-8',
    'Content-Type:application/json',
    'purge_everything: true'
    
));

// -d curl parametresi.
$data = array(
    
    'type' => 'A',
    'name' => '' . $username . '',
    'content' => '' . $dnsipgeldi . '',
    'zone_name' => '' . $domain . '',
    'zone_id' => '' . $zoneid . '',
    'proxiable' => 'true',
    'proxied' => false,
    'ttl' => 1
);

$data_string = json_encode($data);

//die($data_string);
//echo ("https://api.cloudflare.com/client/v4/zones/$zoneid/dns_records?type=A&name=$username&content=$dnsipgeldi&page=1&per_page=20&order=type&direction=desc&match=all");curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);

$sonuc = curl_exec($ch);
// If you want show output remove code slash.

if (strpos($sonuc, 'id') == false) {
    die('{"server":"' . $thisServerName . '","status":"fail","message":"A recored was not found for ' . $username . '.","time":"' . $date . '"}');
} else {
    
    //Get ID and send DELETE
    $record_id = substr($sonuc, 18, 32);
    //die("https://api.cloudflare.com/client/v4/zones/$zoneid/dns_records/$record_id");
    $ch        = curl_init("https://api.cloudflare.com/client/v4/zones/$zoneid/dns_records/$record_id");
    
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
    curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "DELETE");
    curl_setopt($ch, CURLOPT_HTTPHEADER, array(
        'X-Auth-Email: ' . $email . '',
        'X-Auth-Key: ' . $apikey . '',
        'Cache-Control: no-cache',
        // 'Content-Type: multipart/form-data; charset=utf-8',
        'Content-Type:application/json',
        'purge_everything: true'
        
    ));
    
    $result = curl_exec($ch);
    //echo $result;
    
    if (strpos($result, 'success":true') !== false) {
        die('{"status":"success","errors":[],"messages":[],"result":"' . $username . ' deleted"}');
        die('{"server":"' . $thisServerName . '","status":"success","message":"' . $username . ' was deleted.","time":"' . $date . '"}');
    } else {
        die('{"server":"' . $thisServerName . '","status":"fail","message":"' . $username . ' was not removed. CloudFlare returned a fail.","time":"' . $date . '"}');
        
    }
    
    
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
        && preg_match("/^[^\.]{1,63}(\.[^\.]{1,63})*$/", $domain_name)); //length of each label
}
?>
