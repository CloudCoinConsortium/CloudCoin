#  Root RAIDA and the Celebrium

The current RAIDA shall now be called the Root RAIDA. and have the name ".raida"and network number 1.
The Celebrium RAIDA is network 2.

## Root RAIDA Services

[Add Coins To Domain](README.md#add-coins-to-domain)

[Audit Domain](README.md#audit-domain) To see if the coins in Celebrium are what they should be.

[Show Domain](README.md#show-domain)  Shows Details about the Celebrium Domain



## New RAIDA Top Level Domain Services

[Count](README.md#count) Counts how many coins are in the Celebrium Domain

[Top](README.md#top) Returns the top IDs for the coins and for the Denominations

[Mint](README.md#mint) Make new Celebrium

[Detect Denom](README.md#detect-denom) Pown Celebrium and see the meta and image data too. 

[Show Denominations](README.md#show-denominations) Show all the denomination of Celebrium

[Show Denomination](README.md#show-denomination) Show details about one denomination

[Make Change](README.md#make_change) (For Future)

[Exchange](README.md#exchange) (For Future)



The difference between a Root RAIDA coin and a RAIDA Top Level Domain coins can be substantial.  RAIDA Domains are not 
required to follow any standard but there own. 
However, we have a standard that should be useful. The only difference is that the RAIDA Sub Domain coin must 
include a "denomination".

SAMPLES OF POSSIBLE RAIDA Domain Coin formats
NOTE: It is assumed that there is a .raida after the token name. So "RAIDATech" would be raidatech.raida
```{
	"Celebrium": [
	{
		"nn": "2",
		"sn": "1340",
		"an": ["6A415F1CC541353C9AA5CE6FEAFC15B3", "1F07C89ACEF44EFC936C3685EF013D32", "6A415F1CC541353C9AA5CE6FEAFC15B3"."6A415F1CC541353C9AA5CE6FEAFC15B3", "1F07C89ACEF44EFC936C3685EF013D32", ],
    "pown":"ppepp",
    "ed": "0-0000",
		"aoid": []
    }
                ]
}
```


# Root RAIDA Services

## Add Coins To Domain

NOTE: The user can only refer an even number of 250s that are on network 1. Therefor nn and denomination parameters are not required.

The refer service allows a Root RAIDA to know the CloudCoins that are being used as referals. The user can refer all the coins they want without paying. 

The caller must have a domain network number first. 


Example POST requesting three referals
```
Usage: https://RAIDA13.cloudcoin.global/service/refer
domain_name=RAIDATechStock
domain_nn=3
sns[]=16777001&sns[]=16777002&sns[]=16777003&
ans[]=8ae06de0f9ce4917b3309df71570f92c&ans[]=b25fc7a548c341c98cefbac35689aff1&ans[]=f193f1304ffc4344822c10be9309a4c3

```

The RAIDA will check that the coin is authentic, and then check to see if that coins owns the domain specified. 

Each coin needs to be authenticated. If the coin is authentic, the Root RAIDA will update the nn colum in the ANs talble so that it matches the new nn. The AN of the coin will be replaced by the RAIDA's PAN that will be unknown to the owner of the Sub Domain.

RESPONSE IF GOOD
```
[{
  "server":"RAIDA1",
  "status":"added",
  "sn":"16777001",
  "nn":"1",
  "message":"Refrenced:The coin was added as part of your domain. Please mint ($exchange_rate * number of cc sent) tokens",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"added",
  "sn":"16777002",
  "nn":"1",
  "message":"Refrenced:The coin was added as part of your domain. Please mint ($exchange_rate * number of cc sent) tokens",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"fail",
  "sn":"16777003",
  "nn":"1",
  "message":"Conterfiet:The coin was not authentic",
  "time":"2016-44-19 7:44:PM"
}]
```

EXAMPLE OF FAIL BECAUSE THERE IS A FIXED AMOUNT AND THE LIMIT HAS BEEN REACHED

```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Fixed Amount:You have added all the coins that you are allowed to add. Your domain has a fixed supply.",
  "time":"2016-44-19 7:44:PM"
}

```

## Audit Domain

Give it a domain name. It will return the amount of coins it has and the amount of coins it is suppose to have. 

Once every day, a Chron job will call the Count service on corrasponding RAIDA in sub domains. The Chron job will then count
the number of CloudCoins dedicated to that nn in the Ans talbe. It will need to use the domains table to know the conversion rate. 

This information will be stored in the title of a text file. "Celebrium had 1000 coins and is authorized to have 1000 coins.txt"

It will look at the last update of the file and send that too. 


EXAMPLE GET:
```
https://raida9.cloudcoin.global/service/audit_domain?domain=celebrium

```
EXAMPLE RESPONSE:
```

{
  "server":"RAIDA1",
  "status":"audit",
  "message":"Celebrium had 1000 coins and is authorized to have 1000 coins. Checked 2019-10-29",
  "time":"2016-44-19 7:44:PM"
}

```

 
//This is a chron job that runs on each root RAIDA. It goes to the domain's RAIDA servers and calls their "Audit" service. 
//The Audit Domain chron job creates a web page that can be looked at by the public. 
//We need to design this webpage. 

## Show Domain

 FUTURE USE

Shows all the information about a domain including its last Audit

EXAMPLE GET REQUEST
```
https://raida9.cloucoin.global/service/show_domain?domain=RAIDAtech
```
Note: it does not show 

EXAMPLE SUCCESS RESPONSE

```

{
  "server":"RAIDA1",
  "status":"domain",
  "message":"The information is provided",
  ""
  "time":"2016-44-19 7:44:PM"
}
```
EXAMPLE FAIL RESPONSE
```

{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Conterfiet:The coin was not authentic",
  "time":"2016-44-19 7:44:PM"
}
```

# RAIDA Top Level Domain Services

## Top

The top service returns the highest SN in the ANs table. It also returns the highest id in the Denominations table.  
This allows the client to coordinat the RAIDA so that new coins created will be of the same serial numbers.

EXAMPLE GET TOP:
```
https://raida9.raida.tech/service/top?password=j3jkod8ikokfs&
```
EXAMPLE JSON RESPONSE IF SUCCESS:
```
{"server":"raida18","status":"topped","message":"*16777216*1* highest sn in ans and highest id in denominations", "time":"2019-11-09 05:10:45" }

```
EXAMPLE JSON RESPONSE IF FAIL:

```
{"server":"","status":"fail","message":"DB Error", "time":"2019-11-08 05:56:32" }

```


## Count

The count service counts all the rows in the ANs table.It will return the total number of coins in that RAIDA. 
This allows the auditing of coins to make sure the sub raida has no more than allowed by the root RAIDA. 

EXAMPLE GET COUNT:
```
https://raida9.raida.tech/service/count?password=j3jkod8ikokfs&
```
EXAMPLE JSON RESPONSE IF SUCCESS:
```
{"server":"raida0","status":"counted","message":"*1* rows in ans, "time":"2019-12-11 23:01:37" }

```
EXAMPLE JSON RESPONSE IF FAIL:

```
{"server":"","status":"fail","message":"DB Error", "time":"2019-11-08 05:56:32" }

```

## Mint

The Mint allows the Admin of the operator of the RAIDA to create new tokens of the same denomination. 

The ANs for these new coins will be generated based on the seed that is provided. 

The seed will be mixed with the SN of the coin to be created and an MD5 hash of the SN and the Seed will be used as the AN. 

This way, the Client and the RAIDA will know the AN without them needing to pass data back and forth. 

The base 64 will hold the data (such as an image) that is associated with the denomination. 

### Denomination RAIDA:
THe purpose of the RAIDA RAID method is to:
1. Make it fast to upload data. 
2. Make it possible for the user to choose their RAIDA protection
3. Reduce the size of data stored on the mail servers. 
4. Conceal the meaning of the data to the RAIDA Admins. 
5. Make it fast to download the data. 
6. Make it possible for the RAIDA to synchtonize data without the users help. 

In order to achieve this. the user may store their data in three classesd: 
Class 0 is striped. 
Class 01 is stripped and then a mirror of that stripe is stored offsent by one RAIDA
Class 011 (or just lcass 2_) The data is stored with a third mirror offsent by two RAIDA. 


WHERE DATA WILL BE STORED


|CLASS|R0|R1|R2|R3|R4|R5|R6|R7|R8|R9|R10|R11|R12|R13|R14|R15|R16|R17|R18|R19|R20|R21|R22|R23|R24|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
|RAID 0|S0|S1|S2|S3|S4|S5|S6|S7|S8|S9|S10|S11|S12|S13|S14|S15|S16|S17|S18|S19|S20|S21|S22|S23|S24|
|RAID 10|S1|S2|S3|S4|S5|S6|S7|S8|S9|S10|S11|S12|S13|S14|S15|S16|S17|S18|S19|S20|S21|S22|S23|S24|S0|
|RAID 110|S2|S3|S4|S5|S6|S7|S8|S9|S10|S11|S12|S13|S14|S15|S16|S17|S18|S19|S20|S21|S22|S23|S24|S0|S1|

EXAMPLE OF HOW THE STRIPES ARE PUT TOGETHER AFTER THE CLIENT DOWNLOADS THEM
```
//Put all the stripes together. mparts is an array of stripes
function unstripeData(mparts) {

   var unstriped ="";
   
  for(var j = 0; j< mparts[1].length; j++){ 
       for(var i=0; i<25; i++){
           unstriped +=mparts[i].charAt(j);
       }//for all 25 servers
  }//end for every character
   
    // Trim padding at the end
    unstriped = unstriped.replace(/-/g, "");//remove all - padding. 
    return unstriped
}

```

EXAMPLE POST REQUEST:
```
https://raida9.celebrium.com/service/mint_new?
test=true&
denomination_id=4&
title=Kelly Goes To Hollywood&
first_sn=44558&
an_seed=e07550ccfff748d091388c00f44ea5fd&
coins_to_mint=10000&
mint_password=j3jkod8ikokfs&
meta_stripe=kfdkflkd&
meta_mirror=89834&
meta_mirror2=9384983
data_stripe=kfdkflkd&
data_mirror=89834&
data_mirror2=9384983
```
NOTE: If test=true, the service will test to see if all parameters are correct on the RAIDA. You should run it in test mode before minting for real. 

EXAMPLE OF SUCCESS
```
{
    "server":"raida9",
	"status":"success",
	"message":"The denomination titled Kelly Goes To Hollywood has been uploaded and 10000 coins have been created with the first SN being 44558",
	"time":"2019-11-08 05:56:32" 
}
				
```
EXAMPLE OF ERRORS
```
		{
			"server":"raida9",
			"status":"fail",
			"message": "Sorry, your meta data is too large.",
			"time":"2019-11-08 05:56:32" 
		}

		{
			"server":"raida9",
			"status":"fail",
			"message": "Sorry, the denomination already exists.",
			"time":"2019-11-08 05:56:32" 
		}

		{
			"server":"raida9",
			"status":"failure",
			"message": "Sorry, your base64data was not base64",
			"time":"2019-11-08 05:56:32" 
		}
		
```

## Detect Denom

The Detect Denomination service is like detect except it returns the data associated with the denomination. 

You can request three types of data:

1. Stripe:   data=stripe
2. Mirror: data=mirror
3. mirror2: data=mirror2



EXAMPLE GET REQUEST:
```
https://RAIDA18.raida.tech/service/detect_denom?nn=2&sn=1&an=1836843d928347fb22c2142b49d772b5&pan=1836843d928347fb22c2142b49d772b5&data=stripe

```
EXAMPLE RESPONSE IF GOOD:
```
		{
			"server":"raida9",
			"status":"pass",
			"message": "The attached data belongs to the item",
			"title","Kelly Drinks Milk"
			"meta":"Hkjro7urS7LbunvCG",
			"base64data":"Um5wro7urS7LbunvCG8mr0UvY",
			"time":"2019-11-08 05:56:32" 
		}
```

## Show Denominations 
Shows a list of all the names of the denominations on a domain.

You can request three types of data:

1. Stripe:   data=stripe
2. Mirror: data=mirror
3. mirror2: data=mirror2


EXAMPLE GET REQUEST
```
http://raida5.raida.tech/service/show_denominations?read_password=0e9iios8ose&data=stripe

```
Sample Response if successful:
```
{
    "server":"raida9",
    "status":"success",
    "message": "Inventory returned",
    "denominations": [1,2, 3],
    "titles": ["Kelly Goes To Hollywood","Brian Making Bread", "Bill with His Dog"],
    "time":"2019-11-08 05:56:32" 
}
```

## Show Denomination
Shows the details about the specified domain

EXAMPLE GET REQUEST
```
http://raida0.raida.tech/service/show_denomination.php?id=5&password=746193b6dc5249c0bbaa282414aedff0&data=stripe
```
EXAMPLE RESPONSE IF GOOD
```
{
    "server":"raida9",
    "status":"success",
    "message": "Inventory returned",
    "id":1, 
    "title":"Kelly Goes To Hollywood", 
    "metadata":"raid 7, comprsson5 we likethsi ",
    "base64data":"d3920cca24fa47bbb2b784ab7ea42df1",
    "time":"2019-11-08 05:56:32" 
    
}
``` 

## Make change
Make change allows a person to give a CloudCoin note and get back smaller notes from that domain.
TO BE DEVELOPED LATER

## Exchange
Exchanges allows a user to send a bunch of smaller notes and either get a CloudCoin or a bigger denomination
To BE Developed

## Client Images

![Screenshot 1](https://github.com/CloudCoinConsortium/CloudCoin/blob/master/Celebrium/c5.png?raw=false)

![Screenshot 2](https://github.com/CloudCoinConsortium/CloudCoin/blob/master/Celebrium/c4.png?raw=true)

![Screenshot 3](https://github.com/CloudCoinConsortium/CloudCoin/blob/master/Celebrium/c3.png?raw=true)

![Screenshot 4](https://github.com/CloudCoinConsortium/CloudCoin/blob/master/Celebrium/c2.png?raw=true)

![Screenshot 5](https://github.com/CloudCoinConsortium/CloudCoin/blob/master/Celebrium/c1.png?raw=true)
