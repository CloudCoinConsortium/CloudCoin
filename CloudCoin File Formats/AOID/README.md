# AOID ( All Other Information Desired )

The AOID is part of the CloudCoin that allows the coin to be extensible. The AOID is a an array of strings. 
As of yet, the AOID has not been used. However, we now need to use it to let people know about ID coins.

The information in the AOID is for the current owner and who they are trading with. That is why the AOID is ofter stripped out at every exchange. 


## One Standard single key value pair. 
```
"aoid":["id=sean.cloudoin.global"]
```
## Two Standard single key value pair. 
```
"aoid":["email=CloudCoin@protonamil.com","phone=(523)955-9854"]
```
## One multi key value pair. 
Do not do this: 
```
"aoid":["owner=Bill&owner=Betty"]
```
Here we hae two people who own the coin (Bill and Betty) and they have the same key but different values. 
Do this:
```
"aoid":["owner=Bill,Betty"]
```

## Serial Single key. 
```
"aoid":["shuffle=3,20,22,15,24,10,1,11,12,23,17,14,4,13,5,8,2,19,0,16,6,9,7,18,21"]
```
Here we see that the coin shuffles the ANs so they go on different RAIDA. This is encryption. 

# Stardard Key pairs:

### ID

Tells the reader the DNS name of the account that is associated with the coin 
```
"aoid":["id=bill.skywallet.cc"]
```

### Encrypted
Tells the reader if the coin is encrypted
```
"aoid":["encrypted=false"]
```

### Password Hash
Tells the reader what is the hash of the coin's password. This allows a program to check and see if the password is correct before it uses the key.  
```
"aoid":["passwordhash=bf5791868fd5433caaa0042831d2d9e1"]
```

### PAWG
Tells the reader what PANG (Proposed Authenticity Number Generator) is. 
```
"aoid":["pang=55114424555662221155226"]
```


