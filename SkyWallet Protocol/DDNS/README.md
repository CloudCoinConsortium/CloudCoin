# DDNS
The DDNS services allow you to associate an coin serial number with an DNS name. You can also delete this association. 

This php code works with the API of CloudFlare. In order for this code to work, you must have CloudFlare control your DNS system and have a CloudFlare
API Key to create and delete your DNS Records. 

## DDNS Service
Creates a new record in your DNS server. 

This code must be customized for your domain. 

You will need to get a ticket from one of the RAIDA to prove that you own the coin using the get_ticket service on your RAIDA. 

Note: The nn (network number) parameter is 1 for the root RAIDA and 2 for Celebrium. 
EXAMPLE GET REQUEST: 

```
https://yourhost.com/ddns.php?nn=1&sn=1358923&raidanumber=2&ticket=12345678901234567890123456789012345678901234username=Billy1234&

```






## Delete DDNS Service
Deletes a DNS Recored from a CloudFlare controlled DNS Server. 

You will need to get a ticket from one of the RAIDA to prove that you own the coin. 

EXAMPLE GET REQUEST: 
```
https://yourhost.com/ddns_delete.php?nn=1&sn=1358923&raidanumber=2&ticket=12345678901234567890123456789012345678901234username=Billy1234&


```
