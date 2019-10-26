# QR Codes

The QR codes are designed to hold a bunch of data in a small size so that it is easy to read and easy to place on paper or packaging. 

The QR code uses the method of generating PANs based on a seed in the QR code. To turn the coin into a CloudCoin you must be able to run an algorithm against it. 

There are three types of QR Codes:

1. Entire Coin

2. Compressed Coin

3. Subdomain Tokken

# Format 1. Stack file as text on the QR.  

This is very simple, you just take a stack file and insert it into a QR code like it was text. However, it creates a large QR that may be difficult for some phones to read and will be difficult to put on packaging. 

# Format 2. Compressed CloudCoin Format

The compressed CloudCoin Format has three parts:
1. Serial Number
2. Random Letters
3. Network Number

########$$$$$$$$$$##
```
Sample:16777216FOILJNPVUE2
```
### Serial Number
The serial number is a number will usually be between 1 and 16777216 and is from the coin. 

### Random Letters
The Random Letters are capital letters between A and Z. Usually their are 10 but there could be more or less. 10 Provides 95 Trillion combinations. 

### Network Number
If there is no network number, the network number is assumed to be 1. Otherwise the network number just needs to be a number and can be as large as you can imagine. 

### CREATING A CODE FOR A COIN:
The Serial Number and the Network Numbers are taken right off the coin and are not changed. 

### Generating Random String:
For unbeatable security, Generate a string of random capitalized letters between A and Z. 

### Generating the Authenticity Numbers
Attache the Raida ID (from 0 to 24) in front of the Random String. Then run an MD5 has on that number. Use the results as the AN.

Example
```
Randome string: FZLKBNGMNW

RAIDA       Seed      Generatated AN From MD5 Hash of seed
R0   0FZLKBNGMNW  44E16EEFDDFBB4C02E28F61D888D66F2
R1   1FZLKBNGMNW  D5D3540E35ADC72F28AC3FD927C86E94
R2   2FZLKBNGMNW  20E9BCFD395B659C5B0CE1E5C9FFDE25
 R3 - R22 left out for brevity
R23  23FZLKBNGMNW  7B3D77AB3214D247C9D46D4BCECDC9A7
R24  24FZLKBNGMNW  143D6B3585DFCDACF8C721FDA684CD8C
```

Code used to make the QR Code: 
```
16777216ABCDEFGHIK2
```

## Storing the codes before turning them into PNGs. 

The codes can be stored in one file called CompressedQR.txt that is created in a folder called "CompressedQR" that is a sub folder of the "Export" folder. One code per line. Like:
```
compressedQR.txt

16777216ABCDEFGHIK1
9988433EFGHIKABCD1
1300923GHIKABCDEF1


```



# Format 3. Tokens that are not CloudCoins. (NOT FULLY DEVELOPED)

Tokens based on CloudCoin may used GUIDs instead of numbers for their serial number. 

Suppose a token's serial number is: ddf88ef84b244922

Now we can use this to generage ANs


### Generating the Authenticity Numbers
Attache the Raida ID (from 0 to 24) in front of the serial number. Then run an MD5 has on that number. Use the results as the AN.

Example
```
Randome string: FZLKBNGMNW

RAIDA       Seed      Generatated AN From MD5 Hash of seed
R0   0ddf88ef84b244922  44E16EEFDDFBB4C02E28F61D888D66F2
R1   1ddf88ef84b244922  D5D3540E35ADC72F28AC3FD927C86E94
R2   2ddf88ef84b244922  20E9BCFD395B659C5B0CE1E5C9FFDE25
 R3 - R22 left out for brevity
R23  23ddf88ef84b244922  7B3D77AB3214D247C9D46D4BCECDC9A7
R24  24ddf88ef84b244922  143D6B3585DFCDACF8C721FDA684CD8C
```

Code used to make the QR Code: 
```
16777216ABCDEFGHIK2
```


# Sample QR Code

![Without Logo](qr-code.png)

# Sample QR Code with Logo

![QR code with a log0](qr-code-logo.png)
