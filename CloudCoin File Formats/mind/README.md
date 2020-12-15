
## STORE IN MIND

This standard allows us to store an ID coin or a whole bunch of coins in or mind based on two passwords and a serial number or group of serial numbers. 

We can also use this to create a CloudCoin "Credit Card" that contains a serial number and ANs of
a CloudCoin that looks like the familiar Credit Card and can be used to access a Skywallet account. 

# Generating a Random Seed

We need to generate a 16 digit random number using the user's two passwords as a seed. 

Random number: 452 459 836 515 0000

We will ask the user to give us two passwords that they can remember. We want the passwords to be long and complicated. Minumum of 8 charactes each is good. Let the user have lots of freedom but it is a good idea to see that the passwords are not in the list of 1,000 most used passwords. 

Now we mix the passwords together. The first password's first character is used first and then they are woven together like this: NOTE: text to decimal is done like: https://cryptii.com/pipes/text-decimal


```
Password 1: Million2
Password 2: BooBooYou
Weave: Million2BooBooYou
Decimal: 77 66 105 111 108 111 108 66 105 111 111 111 110 85 50 111 117
Random number: 776610511110 (First 12 of Decimal)
ccv: 8111 (13 to 16 of Decimal)

Password 1: OLardy1987
Password 2: Fink
Weave: OFLianrkdy1987
Decimal: 79 70 76 105 97 110 114 107 100 121 49 57 56 55
Random number: 797076105971 (First 12 of Decimal)
ccv: 1011 (13 to 16 of Decimal)

Password 1: Football
Password 2: ChuckBandit3398
Weave: FootballChuckBandit3398
Decimal: 70 67 111 104 111 117 116 99 98 107 97 66 108 97 108 110 100 105 116 51 51 57 56
Random number: 706711110411 (First 12 of Decimal)
ccv: 1117 (13 to 16 of Decimal)

Password 1: g^8KtDeN
Password 2: 0Ik$e^RR
Weave: g0^I8kK$teD^eRNR
Decimal: 103 48 94 73 56 107 75 36 116 101 68 94 101 82 78 82
Random number: 103489473561  (First 12 of Decimal)
Pin/ccv: 0775 (13 to 16 of Decimal)
```

Now we can use this random number to create a card. 

Example CARD:
```
Sean.CloudCoin.Global
9011 5985 4567 8522
EXP 09/22

Back of Card:
CVV 5485

```




### GENERATING THE CARD BASED ON A SKYWALLET ID

* The SN comes from a resolving the SkyWallet name to IP then to decimal. 

* First number must be '9'. Nine means CloudCoin (NOTE: To make it so these card numbers work with Google Card Managers, they can start with the number '4' instead of '9'). 

* The Next two digits are the NN (Network Number) and now it will be 01. 

* The next 12 digits are the first part of the PAN-Generator called 'R' for Random. 

* The last digit is the mod10 Check sum. Add up all the 15 numbers in the card number and the CVV. Then figure out what number is necessary for the number to be divisible by 10. Luhn Algorithm. Since we include the cvv, it is not a true Luhn Algorith:
Sample Card Number: 9010 0307 2080 301*  CVV 5131
Add the numbers up and it equals 44. Devide by 10 any you get 6 remainer. You must add 6 to make it divisable by 10. So you make the last number of the card =6. The card number becomes: 9010 0307 2080 3016 CVV 5101

* Exp Month 01-12 ( 2 digit month - one month before the real exp month).

* Exp Year ( 2 digit year of coin expiration).

* CVV (Must be at least 4 digit in length ). Part of the PAN-Generator coded as 'C'. 

### PAN-Generator 
The PAN-Generator parts are "9" for the CloudCoin sysem, N(two digit network number), 12 digit random number, Luhn number, and a the 4 digit CVV. 9, N(Nework Number, R(andom), L(Luhn Number) and C(VV). 
Here is an example with 452459836515 being random and 8925 being the CVV:
```
Pattern: 9NNR RRRR RRRR RRRL CCCC (Spaces have been added for easy reading).

Random number: 452459836515 

Pin/ccv: 8925

Luhn:

9+0+1+4+5+2+4+5+9+8+3+6+5+1+5+8+9+2+5=86 We must add 4 to make the card number and pin number divisable by 10. So the Luhn number is 
4. 
PAN-Generator: 9014 5245 9836 5154 8925
```
### Generating PANs from the PAN-Generator
Add the RAIDA number and the SN in front of the PAN-Generator and run an MD5 Hash:
NOTE: Here the SN is 16,777,216.
```
PAN for RAIDA 0: 01677721690145245983651548925 MD5 Hash c5eea8d38295aa8a60e766091e55b816
PAN for RAIDA 1: 11677721690145245983651548925 MD5 Hash e385cc046dd21eb1c05c00fea43ffea4
PAN for RAIDA2 through RAIDA22 left out for brevity..
PAN for RAIDA23 231677721690145245983651548925 MD5 Hash 424563309BF5D153C2209412BBC95DF8
PAN for RAIDA24 241677721690145245983651548925 MD5 Hash  8C55ABEAFEDA28A9F875BE816C7BE1EB
```
Now you have the PANs. These PANs must be put in the RAIDA by calling a detect service and using the original ANs. 

### Adding Recovery by Mail
Credit cards can be recovered by email if the an email hash is added to the end of PAN. 
First, calculate the hash of the email by concatinating the serial number of the coin, the RAIDA number 0 through 24 and the email. 
Here we have coin serial number 16777200 on RAIDA14 with the email Billy@Someplace.com. 
We will take the last 8 characters of the PAN and replace it with the first 8 characters of the hash. 
```php
    $sn = 16777200;
    $node_number = 14;
    $email = strtolower("Billy@SomePlace.com");
    $seed = $sn . $node_number . $email; // 1677720014billy@Someplace.com
		$hash = MD5( $seed );//21b289e5492baca352dcce9338fa491d
    //Get the first 8 characters from the Hash. 
    $hash_first_8 = substr($hash,0,8); // 21b289e5
    
    $pan = "8C55ABEAFEDA28A9F875BE816C7BE1EB";
    $pan = substr($string, 0, -8); //Removes last 8 = 8C55ABEAFEDA28A9F875BE81
    $pan = $pan .  $hash_first_8; //8C55ABEAFEDA28A9F875BE8121b289e5
```
You will need to ask the user for their recover email every time they want to do a transaction


## Saving the coins to file:
The Coins can then be saved in a CSV credit Card file. The fields are the PAN-G, Exp. Date, CVV, NN, SN and the QR. (all of them put together in one string that can be used to create a URL for a QR code).
"Credit Cards 12 11 2019 04 04 12.csv"


### This CSV File cointains one coin or card per line.
```
9014 5245 9836 5154, 02-25, 8925, 1, 16777200, n=1&s=16777200&r=9014524598365154&d=02-25&c=8925& 
9014 8563 1894 5726, 02-25, 0505, 1, 16777201, n=1&s=16777201&r=90148563189457266&d=02-25&c=0505& 
9016 5216 9524 9265, 02-25, 1289, 1, 16777202, n=1&s=16777202&r=9016521695249265&d=02-25&c=1289& 
```

### This CSV file is for multi-urls because all the coins have the same Card Number. This cannot be used for creating cards but it can be used to create multiple coins. Like when you want to send a lot of coins to one person. 

```
9014 5245 9836 5154, 02-25, 8925, 1, 16777200, n=1&r=9014524598365154&d=02-25&c=8925&s[]=16777200&s[]=485225&s[]=3895484&s[]=184554&s[]=96544

```


![Card Front](card-front.png)


![Card Front](card-back.png)

![Card Templet](card.png)


Sample of what it should look like when done: 

![Card Front](card.jpg)

https://github.com/CloudCoinConsortium/CloudCoin/tree/master/CloudCoin%20File%20Formats/jpg






