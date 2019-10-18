# SKYWALLET PROTOCOL

The Skywallet protocols can be devided into four different catagories. 
1. Trusted Transfer
2. Merchante Services
3. Donation Services
4. Change Maker Services

# The Trusted Transfer Services 

These services allow one person to pay another without the other person needed to "pown" the coins. This service is the third attempt to achieve this. 
The first was called "Trusted Transfer." The 2nd called "PayForward". This latest attempt does away with external servers completely and attempts to achieve the 
same internally. The solution does create new problems: Keys can be lost and stolen. This means the keys become a sentral way of failure just like the 
private keys of the blockchain. 

[Send](README.md#send)

[ShowEnvelopes](README.md#showenvelopes)

[ShowCoinsInEnvelope](README.md#showcoinsinenvelope)

[Rename Tag](README.md#renametag)

[Receive](README.md#receive)

[Transfer](README.md#transfer)

# Merchant Services

[VerifyPayment](README.md#verifypayment)

[SendWithChange](README.md#sendwithchange)

[TransferWithChange](README.md#transferwithchange)

# Donation Services
Allows the user to donate money to the change system or to the RAIDA itself (to help the RAIDA fund itself).

[DonateToRaida](README.md#donatetoraida)

[DonateChange](README.md#donatechange)

# The Change Maker Services

The change services allows one to break notes into smaller ones and consolidate small notes into larger ones. 

[ShowChange](README.md#showchange)

[Change](README.md#change)

[Exchange](README.md#exchange)


## Send

The Send allows many coins to be authenticated at the same time but their PANs are generated by the RAIDA instead of the owner. 
This allows CloudCoins to be transfered from person to person with no powning neccessary.

rules: 
"to_sn" must be a number between 1 and 16,777,215 because it is the serial number of a CloudCoin that a user must own in order to retrieve the coins. 

"envelope_name" must not start or end with any white space character and must use only UTF-8 characters. All the Characters must be English Alphanumeric. 


EXAMPLE POST SENDING THREE COINS TO SN 16225354
```
    https://s0.teleportnow.cc/service/send
    nn=1&
    sns[]=145895&sns[]=66585&sns[]=16589554&
    ans[]=8ae06de0f9ce4917b3309df71570f92c&ans[]=b25fc7a548c341c98cefbac35689aff1&ans[]=f193f1304ffc4344822c10be9309a4c3&
    denomination[]=250&denomination[]=1&denomination[]=250&
    to_sn=16225354&
    envelope_name=1Z 999 AA1 01 2345 6784

```

RESPONSE IF SUCCESS:
```
[{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"66585",
  "nn":"1",
  "message":"Authentic: 1-unit. Your coins have been sent to the address you specified.",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"66586",
  "nn":"1",
  "message":"Authentic: 1-unit. Your coins have been sent to the address you specified.",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"16589554",
  "nn":"1",
  "message":"Authentic: 250-unit. Your coins have been sent to the address you specified.",
  "time":"2016-44-19 7:44:PM"
}]
```

RESPONSE IF TOO MANY COINS SENT (OVER 400)
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Too many coins attached.",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ARRAY LENGTHS NOT THE SAME
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Arrays not all the same length (nn,sn,an,denominations).",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires you to supply the com_broker,
to_account_name_or_number, total_to_send and change_to_account_name_or_number."
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ANY OF THE sns ARE THE SAME AS THE to_sn
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"sns: no sns can be the same as to_sn"
  "time":"2016-44-19 7:44:PM"
}
```

## ShowEnvelopes
An Envelope is a grouping of coins usally grouped by a transaction number, purchase order number, memo, account number. 
For transactions, the seller gives the buyer their account number and an envelope name (which could be a receipt number) so the receiver can 
determine where the money came from. 
This service allows the user to see the names of the envelopes that belong to their account id (The serial number of their ID coin).

Envelop names could be like:

"From Billy Williams"

"1Z 999 AA1 01 2345 6784"

"1836843d928347fb22c2142b49d772b5"

"Use this money to obey god & destroy tyrants" // up to 44 characters



EXAMPLE GET REQUESTING TO SEE ALL THE ENVELOPES THAT BELONG TO THE USER. 
```
https://s0.teleportnow.cc/service/show_envelopes?nn=1&sn=1&an=1836843d928347fb22c2142b49d772b5&pan=1836843d928347fb22c2142b49d772b5&denomination=1

```

SUCCESSFUL RESPONSE SHOWING THAT THERE ARE THREE ENVELOPES IN THE TRANSFER TABLE:
```
{
	"server": "RAIDA1",
	"status": "success",
	"envelopes": [{
		"tag": "From Billy",
		"1s": 20,
		"5s": 15,
		"25s": 0,
		"100s": 18,
		"250s": 23,
		"total": 1874,
		"date_received": "2016-44-19 7:44:PM"
	}, {
		"tag": "From Jill",
		"1s": 20,
		"5s": 15,
		"25s": 0,
		"100s": 18,
		"250s": 23,
		"total": 1874,
		"date_received": "2016-44-19 7:44:PM"
	}, {
		"tag": "Payment for Car",
		"1s": 20,
		"5s": 15,
		"25s": 40,
		"100s": 18,
		"250s": 6,
		"total": 804,
		"date_received": "2016-44-19 7:44:PM"
	}],
	"message": "Showing: Envelope Names",
	"time": "2016-44-19 7:44:PM"
}
```

RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires you to supply the id_sn, id_nn and id_ans and id_pan parameters."
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ACCOUNT DID NOT LOG IN CORRECTLY
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Login: Authenticity Number was incorrect for the Serial Number .",
  "time":"2016-44-19 7:44:PM"
}
```


## ShowCoinsInEnvelope

This service shows how many coins you have that are in an envelope that you specify. 
Each coin will be listed so that the person will know what they have. 

"envelope_name" must not start or end with any white space character and must use only UTF-8 characters. 

Example GET asking to see all the coins they own in the transfer file. 
```
https://s0.teleportnow.cc/service/show_coins_in_envelope?nn=1&sn=1&an=1836843d928347fb22c2142b49d772b5&pan=1836843d928347fb22c2142b49d772b5&denomination=1&envelope_name=8rie

```

RESPONSE SHOWING THAT THERE ARE THREE COINS OWNED:
```
{
	"server": "RAIDA1",
	"status": "success",
	"tag":"From Billy",
	"coins": [{
			"nn": "1",
			"sn": "9955856",
			"ed": "2019-10-1",
			"denomination": "1"
		},
		{
			"nn": "1",
			"sn": "9955857",
			"ed": "2019-10-1",
			"denomination": "1"
		},
		{
			"nn": "1",
			"sn": "9955858",
			"ed": "2019-10-1",
			"denomination": "1"
		}

	],
	"message": "Coins: List of Coins has been returned.",
	"time": "2016-44-19 7:44:PM"
}
```


RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires you to supply the id_sn, id_nn and id_ans parameters."
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ACCOUNT DID NOT LOG IN CORRECTLY
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Login: Authenticity Number was incorrect for the Serial Number .",
  "time":"2016-44-19 7:44:PM"
}
```


## Show

This service shows how many coins you have that belong to you that you have not taken control of yet. 
Each coin will be listed so that the person will know what they have. 

"tag" must not start or end with any white space character and must use only UTF-8 characters. 

Example GET asking to see all the coins they own in the transfer file. 
```
https://s0.teleportnow.cc/service/show?nn=1&sn=1&an=1836843d928347fb22c2142b49d772b5&pan=1836843d928347fb22c2142b49d772b5&denomination=1
```

RESPONSE SHOWING THAT THERE ARE TWO COINS OWNED IN THE TRANSFER TABLE:
```
{
    "server": "RAIDA0",
    "status": "pass",
    "message": [
        {
            "sn": "11",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559780031"
        },
        {
            "sn": "12",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559780326"
        },
        {
            "sn": "13",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559780351"
        },
        {
            "sn": "14",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559780351"
        },
        {
            "sn": "15",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559780351"
        },
        {
            "sn": "16",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559780351"
        },
        {
            "sn": "17",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559784415"
        },
        {
            "sn": "18",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559785124"
        },
        {
            "sn": "19",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559785224"
        },
        {
            "sn": "20",
            "tag": "q1w2e3r4 asdasda asdasdf",
            "created": "1559785553"
        }
    ],
    "time": "2019-06-13 05:27:35",
    "nn": "1",
    "sn": "21"
}

```


RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires you to supply the id_sn, id_nn and id_ans parameters."
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ACCOUNT DID NOT LOG IN CORRECTLY
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Login: Authenticity Number was incorrect for the Serial Number .",
  "time":"2016-44-19 7:44:PM"
}
```



## RenameTag

Rename Tag allows you to take an entire envelope and either change its name or move the entire contents into an existing envelope. This way, CloudCoins that have been
sent to you may be moved to a different envelope. A person may want to do this to protect themselves from a 
person trying to use the same envelope name over and over again to prove that they paid for something. This is because the envelope name can be used by the sender as a proof
of payment. The Rename Tag service also tells how much was in the envelope so the receiver can tell if they have been paid enough. 


Example GET asking to change the name of an tag or envelope. 
```
https://s0.teleportnow.cc/service/rename_tag?
nn=1&
sn=1&
an=1836843d928347fb22c2142b49d772b5&
pan=1836843d928347fb22c2142b49d772b5&
denomination=1&
tag=1Z 999 AA1 01 2345 6784&
new_tag=6892ed132f8741a6ab1332eb5a4543ec

```



RESPONSE IF GOOD:
```
{
  "server":"RAIDA1",
  "status":"imported",
  "total_coins":"1562",
  "message":"Moved: The coins in the envelope specified are now in the new envelope and the old envelope has been removed",
  "time":"2016-44-19 7:44:PM"
}
```


RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires you to supply the sn, nn, an, pan, deonomination, envelope_name and new_envelope_name parameters."
  "version":"some version number here",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ACCOUNT DID NOT LOG IN CORRECTLY
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Login: Authenticity Number was incorrect for the Serial Number .",
  "version":"some version number here",
  "time":"2016-44-19 7:44:PM"
}
```


## Receive


The Receive service allows many coins to be downloaded from the RAIDA at the same time. The Receiver can download all of their
coins, or just some of there coins. But it must specify the coins it wants to download by serial number. This means 
the client will have to do a "ShowCoins" so they can get a list of the serial number that they have.
This allows the RAIDA to actaully store coins for people. The Receiver must supply proof that they are the owner.
They prove that they are the owner by supplying a serial number of a CloudCoin that they own (their account number) and the AN for that serial number. 
The RAIDA checks its tranfer table and starts giving coin until the amount of coins asked to be downloaded is reached. 


This allows CloudCoins to be transfered from person to bank with no powning necessary.

rules: 
"to_number" must be a number between 1 and 16,777,215 because it is the serial number of a CloudCoin that a user must own in order to retrieve the coins. 

"envelope_name" must not start or end with any white space character and must use only UTF-8 alpha numeric characters. 

Example POST asking for specific CloudCoins
```
https://s0.teleportnow.cc/service/recieve?
nn=1&
sn=1&
an=1836843d928347fb22c2142b49d772b5&
pan=1836843d928347fb22c2142b49d772b5&
denomination=1&
nns[]=1&
nns[]=1&
nns[]=1&
sns[]=152658&
sns[]=9955856&
sns[]=6652154&


```



RESPONSE IF GOOD:
```
[{
  "server":"RAIDA1",
  "status":"received",
  "sn":"152658",
  "nn":"1",
  "an":"c18889d9028240b796ea2389d0e36219",
  "message":"Please record the an provided within a CloudCoin file",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"received",
  "sn":"9955856",
  "nn":"1",
  "an":"e56507ecd05945d990bf6655546bdeff",
  "message":"Please record the an provided within a CloudCoin file",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"missing",
  "sn":"6652154",
  "nn":"1",
  "an":"",
  "message":"Nonexistant: That serial number has not been transfered to the id provided.",
  "time":"2016-44-19 7:44:PM"
}]
```

RESPONSE IF TOO MANY COINS SENT (OVER 400)
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Too many coins recieved.",
  "version":"some version number here",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ARRAY LENGTHS NOT THE SAME
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Arrays not all the same length (nns,sns,pans,).",
  "version":"some version number here",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires you to supply the id_sn, id_nn, id_ans, nns[]=1, sns[] and pans[] parameters."
  "version":"some version number here",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ACCOUNT DID NOT LOG IN CORRECTLY
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Login: Authenticity Number was incorrect for the Serial Number .",
  "version":"some version number here",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF THERE WAS A DATABASE ERROR
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Database: Operation failed. Your Authenticity Number has been left unchanged.",
  "version":"some version number here",
  "time":"2016-44-19 7:44:PM"
}
```


## Transfer

Example POST asking to transfer 3 CloudCoin notes to user 16225354
```
https://s0.teleportnow.cc/service/receive?
nn=1&
sn=16777216&
an=8ae06de0f9ce4917b3309df71570f92c&
pan=8ae06de0f9ce4917b3309df71570f92c&
denomination=1&
to_sn=16225354&
nns[]=1&
nns[]=1&
nns[]=1&
sns[]16777214&
sns[]16777215&
sns[]16777216&
tag=8d0caf3
```


RESPONSE IF GOOD:
```
[{
  "server":"RAIDA1",
  "status":"transfered",
  "sn":"16777214",
  "nn":"1",
  "message":"transfered: This coin has been sent to the address you specified.",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"transfered",
  "sn":"16777215",
  "nn":"1",
  "message":"Authentic: 1-unit. Your coins have been sent to the address you specified.",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"fail",
  "sn":"16777216",
  "nn":"1",
  "message":"Nonexistant: The serial Number does not belong to id given.",
  "time":"2016-44-19 7:44:PM"
}]
```

RESPONSE IF TOO MANY COINS SENT (OVER 400)
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Too many coins attached.",
  "time":"2016-44-19 7:44:PM"
}
```




RESPONSE IF ARRAY LENGTHS NOT THE SAME
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Arrays not all the same length (nn,sn,an,denominations).",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires you to supply the com_broker,
to_account_name_or_number, total_to_send and change_to_account_name_or_number."
  "time":"2016-44-19 7:44:PM"
}
```
RESPONSE IF ANY OF THE sns ARE THE SAME AS THE to_sn
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"sns: no sns can be the same as to_sn"
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ACCOUNT DID NOT LOG IN CORRECTLY
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Login: Authenticity Number was incorrect for the Serial Number .",
  "time":"2016-44-19 7:44:PM"
}
```

## VerifyPayment

Verify Payment allows a merchant to see if a payment was recieved, make sure the correct amount was recieved, move that payment to another envelope and provide change if needed. The merchant must have 
a special "change" envelope to issue change. 

payment_envelope is the tag that the sender has used when he sent money to the merchant. 

payment_expected is the number of CloudCoins that the sender was suppose to send. The merchant will check to see if the actual payment from the sender was correct. 

return_to_sn is the address that any change or refundes should be sent to. If the sender does not send enough money, they will get their money back at this account. If the
sender sends too much money then they will get the change back at this account. 

storage_envelope is the place that the senders payment will be moved to. This is done by renaming the envelope. 

Example GET asking to change the name of an tag or envelope. 
```
https://s0.teleportnow.cc/service/rename_tag?
---ID COIN---
nn=1&
sn=1&
an=1836843d928347fb22c2142b49d772b5&
pan=1836843d928347fb22c2142b49d772b5&
denomination=1&
---Payment Information---
tag=1Z 999 AA1 01 2345 6784&
new_tag=6892ed132f8741a6ab1332eb5a4543ec
total_expected=33
refund_sn=16555897


```


RESPONSE IF GOOD WITH NO CHANGE RETURNED:
```
{
  "server":"RAIDA1",
  "status":"verified",
  "message":"The coins in the payment envelope are now in the storage envelope. The payment envelope has been removed.",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF GOOD WITH CHANGE RETURNED:
```
{
  "server":"RAIDA1",
  "status":"verified",
  "message":"The coins in the payment envelope are now in the storage envelope. The payment envelope has been removed. Change was returned.",
  "time":"2016-44-19 7:44:PM"
}
```


RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"The request requires you to supply the sn, nn, an, pan, deonomination, payment_envelope, payment_expected, return_to_sn and storage_envelope parameters."
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ACCOUNT DID NOT LOG IN CORRECTLY
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Login: Authenticity Number was incorrect for the Serial Number .",
  "time":"2016-44-19 7:44:PM"
}
```
RESPONSE IF PAYMENT ENVELOPE HAD LESS THAN THE EXPECTED AMOUNT
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"$payment_expected Coins were expected but only $coin_count were supplied",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF PAYMENT ENVELOPE HAD MORE THAN THE EXPECTED AMOUNT AND COULD NOT MAKE CHANGE
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"$payment_expected Coins were expected but $coin_count were supplied. Cannot make change",
  "time":"2016-44-19 7:44:PM"
}
```


## SendWithChange
This service allows you to send money from your account to another and make change if needed. 

This service requires an account that has a "public_change" envelope. Account SN 2 has a public_change envelope but any account could allow for one. 

The sender sends a note such as 250 to the public change agent. 

The public change agent breaks the 250 as specifeid and sends some to the payee and the rest back to the sender's account as change. 


The Send allows many coins to be authenticated at the same time but their PANs are generated by the RAIDA instead of the owner. 
This allows CloudCoins to be transfered from person to person with no powning necessary.

rules: 
"to_sn" must be a number between 1 and 16,777,215 because it is the serial number of a CloudCoin that a user must own in order to retrieve the coins. 

"envelope_name" must not start or end with any white space character and must use only UTF-8 characters. All the Characters must be English Alphanumeric. 


EXAMPLE POST SENDING THREE COINS TO SN 16225354
```
    https://s0.teleportnow.cc/service/send_with_change
    nn=1&
    sns[]=145895&sns[]=66585&sns[]=16589554&
    ans[]=8ae06de0f9ce4917b3309df71570f92c&ans[]=b25fc7a548c341c98cefbac35689aff1&ans[]=f193f1304ffc4344822c10be9309a4c3&
    denomination[]=250&denomination[]=1&denomination[]=250&
    public_change_maker=16225354&
    send_to_sn=9225362&
    amount_sent=250
    payment_required=33
    payment_envelope=7URE
    return_to_sn=16555897

```

RESPONSE IF SUCCESS:
```
[{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"66585",
  "nn":"1",
  "message":"Authentic: 1-unit. Your coins have been sent to the address you specified.",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"66586",
  "nn":"1",
  "message":"Authentic: 1-unit. Your coins have been sent to the address you specified.",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"16589554",
  "nn":"1",
  "message":"Authentic: 250-unit. Your coins have been sent to the address you specified.",
  "time":"2016-44-19 7:44:PM"
}]
```

RESPONSE IF TOO MANY COINS SENT (OVER 400)
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Too many coins attached.",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ARRAY LENGTHS NOT THE SAME
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Arrays not all the same length (nn,sn,an,denominations).",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires you to supply the com_broker,
to_account_name_or_number, total_to_send and change_to_account_name_or_number."
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ANY OF THE sns ARE THE SAME AS THE to_sn
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"sns: no sns can be the same as to_sn"
  "time":"2016-44-19 7:44:PM"
}
```



## TransferWithChange

Example POST asking to transfer 3 CloudCoin notes to user 9225362


```
https://s0.teleportnow.cc/service/receive?
---ID---
nn=1&
sn=16777216&
an=8ae06de0f9ce4917b3309df71570f92c&
pan=8ae06de0f9ce4917b3309df71570f92c&
denomination=1&
---Coins to Send---
nn=1&
sns[]16777214&
sns[]8245591&
sns[]1366648&
---Instructions---
public_change_maker=16225354&
send_to_sn=9225362&
payment_required=251
payment_envelope=A8UP
```

RESPONSE IF SUCCESS:
```
{
  "server":"RAIDA1",
  "status":"transfered",
  "message":"Your coins have been transfered to the address you specified. Change was returned.",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF VALUE OF CLOUDCOINS TRANSFERED IS LESS THAN PAYMENT REQUIRED
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Transferd coins were less than the coins required.",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF CHANGE MAKER IS NOT A CHANGE MAKER
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"change_maker is not a Change Maker.",
  "time":"2016-44-19 7:44:PM"
}
```


RESPONSE IF CHANGE SERVER WILL CAN NOT MAKE CHANGE
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Change maker does not have enough change to make change.",
  "time":"2016-44-19 7:44:PM"
}
```


RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED (One for each parameter)
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"The request requires you to supply the payment_envelope parameter."
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ANY OF THE sns ARE THE SAME AS THE to_sn
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Attempt to send sns[] cannot be done because it is the user's ID sn"
  "time":"2016-44-19 7:44:PM"
}
```




## DonateToRAIDA (We can implement this later)
This allows people to donate CloudCoins to pay the costs of running the servers. 

People can donate as much as 400 coins. 

There is a folder on the RAIDA called the "Donations" folder. 

Example POST donating three coins to the RAIDA  
```
https://s0.teleportnow.cc/service/donate_to_raida
nns[]=1&nns[]=1&nns[]=1&
sns[]=145895&sns[]=66585&sns[]=16589554&
ans[]=8ae06de0f9ce4917b3309df71570f92c&ans[]=b25fc7a548c341c98cefbac35689aff1&ans[]=f193f1304ffc4344822c10be9309a4c3&
denomination[]=250&denomination[]=1&denomination[]=250&


```

RESPONSE:
```
[{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"145895",
  "nn":"1",
  "message":"Authentic: 1-unit. Your coins have been donated.",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"66585",
  "nn":"1",
  "message":"Authentic: 1-unit. Your coins have been donated.",
  "time":"2016-44-19 7:44:PM"
},
{
  "server":"RAIDA1",
  "status":"pass",
  "sn":"16589554",
  "nn":"1",
  "message":"Authentic: 250-unit. Your coins have been donated.",
  "time":"2016-44-19 7:44:PM"
}]
```

RESPONSE IF TOO MANY COINS SENT (OVER 400)
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Too many coins attached.",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF ARRAY LENGTHS NOT THE SAME
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Length: Arrays not all the same length (nn,sn,an,denominations).",
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters: The request requires nns[],sns[], ans[] and denomination[]."
  "time":"2016-44-19 7:44:PM"
}
```



## ShowChange

ShowChange service returns the serial numbers of CloudCoins that are the specified account's "public_change" folder.
This allows the client to request to make change from 25 different RAIDA but use the same 
Serial Number. Otherwise, every RAIDA could return a different SN and all the coins would be counterfeit.

Anyone can created an envelope named "public_change." It may make it easier on the server side to have five envelopes named "public_change_1s",
"public_change_5s", "public_change_25s", "public_change_100s" and "public_change_250s". 

Each of these envelopes will contain the denominations for that type of CloudCoin.

Example GET asking for specific CloudCoins
```
https://RAIDA1.CloudCoin.Global/service/show_change?sn=778832

```

RESPONSE IF THE RESPONSE WAS GOOD:
```
{
  "server":"RAIDA1",
  "status":"shown",
  "250s":[16230602,16675880,16192311,...show as many as ten ... 15169770],
  "100s":[13230602,13675880,16192311,...show as many as ten ... 15169770],
  "25s":[10230602,10675880,10192311,...show as many as ten ... 15169770],
  "5s":[8230602,8675880,6192311,...show as many as ten ... 15169770],
  "1s":[230602,675880,192311,...show as many as ten ... 15169770],
  "message":"Change:This report shows the serial numbers that are available to make change now.",
  "version":"some version number here",//Optional 
  "time":"2016-44-19 7:44:PM"
}
```

RESPONSE IF CHANGE NOT OFFERED OR AVAILABLE
```
{
  "server":"RAIDA1",
  "status":"nochange",
  "message":"Message. Sorry the account specified does not provide chane to the public."
  "version":"some version number here",//Optional
  "time":"2016-44-19 7:44:PM"
}
```



## Change
Take one CloudCoin and turns it into many smaller ones.  This is almost exactly like Receive except that instead of providing authentication, coins are provided instead. 
The caller must suggest which SNs it would like to recieve as change. This means that in order to get change, 
the caller must use the "ShowChange" service to supply a list of SNs that it expects to get in return.

As a rule, it can be required that the client always ask for change as folows:
```

250:
    100x2
    25x1
    5x4
    1x5
    
100:
    25x3
    5x4
    1x5

25:
    5x4
    1x5

5:
    1x5
    
```
NOTE: The above combination ensure that the client will always have the correct change to make any transaction. 

The service will go to the "public_change" envelope for the account spedified and take the requested coins out of the transfer table and return the ANs requested. 


EXAMPLE GET SENDS ONE 250 NOTE AND REQUESTS THAT IT IS BROKEN INTO 2 100s and 2 25s. 
THE COINS TO BE RETURNED ARE SPECIFIED IN THE SNS[] PARAMETERS
```
https://s0.teleportnow.cc/service/change?nn=1&sn=16777202&an=8ae06de0f9ce4917b3309df71570f92c&denomination=250sns[]=13555555&sns[]=5966558&sns[]=5556665&sns[]=8887372
```


RESPONSE IF SUCCESS BREAKING COIN INTO SEVERAL SMALLER COINS:
```
{
  "server":"RAIDA1",
  "status":"change",
  "nns":["1","1","1","1"],
  "sns":["13555555","5966558","5556665","8887372"],
  "ans":["a91c5b6456b74217a27e5e3d518ab49b","e58e8620310c4ea3bd4bb6c3999deb64","0f915257f5cf40bca2e4f94e054bdd32","6f24bd539ce941feaf4a01b10ac59559"],
  "message":"Change: 1-unit. Save these ANs to you CloudCoins.",
  "time":"2016-44-19 7:44:PM"
}

```


RESPONSE IF THE COIN SENT WAS COUNTERFEIT:
```
{
  "server":"RAIDA1",
  "status":"fail",
  "message":"Counterfeit: The Unit you sent sent was counterfeit.",
  "time":"2016-44-19 7:44:PM"
}

```


RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"error",
  "message":"Parameters: The request requires parameters such as nn,sn, an, denomination and sns[] but thery were not supplied,"
  "time":"2016-44-19 7:44:PM"
}
```




## Exchange
This takes many smaller notes and exchanges them for one big note. 
The number of notes provided must add up to either 250, 100, 25 or 5.
This is almost exactly like Receive except that instead of providing authentication, coins are provided instead. 

The user must specify the big coin that they hope to get in return. This helps
the server check to see if it has that change. There are five return: 250,100,25 and 5 ( 1 is not a type).

EXAMPLE POST EXCHANGING FOUR COINS TO MAKE A BIGGER 250 NOTE  
```
https://s0.teleportnow.cc/service/exchange?
return=250&
nns[]=1&nns[]=1&nns[]=1&nns[]=1&
sns[]=145895&sns[]=66585&sns[]=16589554&sns[]=16589554&
ans[]=8ae06de0f9ce4917b3309df71570f92c&ans[]=b25fc7a548c341c98cefbac35689aff1&ans[]=f193f1304ffc4344822c10be9309a4c3&ans[]=f193f1304ffc4344822c10be9309a4c3&
denomination[]=100&denomination[]=100&denomination[]=25&&denomination[]=25
```

RESPONSE IF GOOD:
```
{
  "server":"RAIDA1",
  "status":"exchange",
  "nns":"1",
  "sns":"14589005",
  "ans":"6f24bd539ce941feaf4a01b10ac59559",
  "message":"Exchanged. Record as a CloudCoin.",
  "time":"2016-44-19 7:44:PM"
}

```


RESPONSE IF JUST ONE OF THE COINS WAS COUNTERFEIT:
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Counterfeit. At least one of the units sent was counterfeit. No ANs were changed.",
  "time":"2016-44-19 7:44:PM"
}

```


RESPONSE IF SOME PARAMETERS WERE NOT SUPPLIED
```
{
  "server":"RAIDA1",
  "status":"dud",
  "message":"Parameters. The request requires type, sns[], ans[], nn[], denomination[]"
  "time":"2016-44-19 7:44:PM"
}
```



RESPONSE IF THE TYPE WAS INVALID
```
{
  "server":"RAIDA1",
  "status":"dud",The request requires return of either 250, 100, 25 or 5."
  "time":"2016-44-19 7:44:PM"
}
```