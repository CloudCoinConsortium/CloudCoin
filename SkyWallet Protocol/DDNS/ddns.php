<?php

/*
ddns.php

This program allows someone to create a DNS record in Skywallet.cc. It requires that the user prove that they own the number
that they want to associate with the A record. 
Written May 13th 2019
By Sean Worthington

Sample Use:
https://yourhost.com/ddns.php?nn=1&sn=1358923&raidanumber=2&ticket=12345678901234567890123456789012345678901234username=Billy1234&


*/

$thisServerName    = "yourhost";// This is the name of the server that will be return in all JSON responses. 
$trustedDomainName = "cloudcoin.global";//This is the domain name that will be used to contact the raida to check tickets.  
$date              = date("Y-m-d H:i:s");

/* Cloudflare.com | APİv4 | Api variables */
$apikey = '2a2efe54f GET FROM SEAN 0d6b263266a'; // Cloudflare Global API
$email  = 'cloudcoin@protonmail.com'; // Cloudflare Email Adress
$domain = 'skywallet.cc'; // zone_name // Cloudflare Domain Name
$zoneid = '561ed GET FROM SEAN 16fd2cdc234'; // zone_id // Cloudflare Domain Zone ID


//SEE IF GET VARIABLES ARE NOT THERE
if (!isset($_GET['nn']) || !isset($_GET['ticket']) || !isset($_GET['sn']) || !isset($_GET['raidanumber']) || !isset($_GET['username'])) {
    die('{"server":"' . $thisServerName . '","status":"fail","message":"GET Parameters: You must provide a ticket, nn, sn, username and raida_numer.","time":"' . $date . '"}');
}


$username = $_GET['username'];
$sn       = intval($_GET['sn']);
$nn = intval($_GET['nn']);//Get network number. 1 = root RAIDA. 2 =  Celebrium RAIDA. 
$raidaNumber = intval($_GET['raidanumber']);//Get the RAIDA number that the ticket came from
$ticket      = $_GET['ticket'];//This came from a get_ticket call to the raidaNumber above. 


//Validate
if (!is_valid_domain_name($username) || strpos($username, ".") !== false) { //Name is bad
    die('{"server":"' . $thisServerName . '","status":"fail","message":"Username is not valid","time":"' . $date . '"}');   
} //if username bad

//Make sure sn is within range.
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
$raw_from_server_1 = curl_get_contents("https://raida$raidaNumber.$trustedDomainName/service/hints?rn=$ticket");
//echo $raw_from_server_1;//debug
//die("https://raida$raidaNumber.$trustedDomainName/service/hints?rn=$ticket");//debug

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
$sn         = $sn + (16777216 * $nn); //Makes the first octet of the dotted decimal IP address be 1 or 2. 
$dnsipgeldi = long2ip($sn); //String will now start with 

//If good talk to the CloudFlare API
// A-record oluşturur DNS sistemi için.
$ch = curl_init("https://api.cloudflare.com/client/v4/zones/" . $zoneid . "/dns_records");
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "POST");
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

curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $data_string);
//curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($data_string));

$sonuc = curl_exec($ch);

// If you want show output remove code slash.
if (strpos($sonuc, 'success":true') !== false) {
    echo '{"status":"success","errors":[],"messages":[],"result":"' . $username . ' created"}';
} else {
    
    $sonuc = str_replace("success", "status", $sonuc);
    $sonuc = str_replace("false", "\"fail\"", $sonuc);
    
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
        && preg_match("/^[^\.]{1,63}(\.[^\.]{1,63})*$/", $domain_name)); //length of each label
}
?>
