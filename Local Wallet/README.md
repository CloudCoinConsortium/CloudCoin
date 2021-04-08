# Local Wallet Standards
A Local Wallet stores CloudCoins on a local device. This is the opposite of a Skywallet that stores CloudCoins on a RAIDA. The main features of a Local Wallet are to check for authenticity, 
change ownership, place good coins in a protected folder and fix coins that are unsyncronized (fracked).

[Folder Standards](README.md#folder-standards)

[Functionality Standards](README.md#functionality-standards)

[Testing Standards](README.md#testing-standards)

## Folder Standards
The folder structure is the backbone of the Local Wallet and we would like to keep it the same forever. 
The reason we have the folders is to ensure the coins are not in RAM if and when the program crashes. 
This could happen if the electricity goes out. Writing to files is slower but safer. Coins travel through
the program by starting in the "Import" folder and then being unpacked and moved to the "Suspect" folder. The 
original file is moved to the "Imported" folder. Then the coins are checked against the raida and moved to the
"Detected" folder where they are graded. After grading, the coins will be moved into either the "Bank" if authentic
"Fracked" if authentic but unsynctonized, "Counterfeit" if bad or "Limbo" if there is no response from the RAIDA and
it is unknown if the coins are good or bad. 

Here is the folder structure:
```javascript
Program Name //sample name: "CloudCoin Wallet"
│   global.config
│
├───Accounts 
│   └───Account Name //sample name "Bill's Wallet"
│       │   transactions.txt
│       │
│       ├───Bank
│       │       1.CloudCoin.1.1287829.stack //sample file
│       │
│       ├───Config
│       │       config.txt
│       │
│       ├───Counterfeit
│       ├───Deposit
│       ├───Detected
│       ├───EmailOut
│       ├───EmailSent
│       ├───Export
│       ├───Fracked
│       ├───Gallery
│       ├───Import
│       ├───Imported
│       │       5.CloudCoin.1.4181672.stack //Sample file
│       ├───Logs
│       ├───Limbo
│       ├───Mind
│       ├───OtherCoins
│       ├───Partial
│       ├───PayForward
│       ├───Predetect
│       ├───Receipts
│       │       32DF1105AD003420F5A3C984CC50FBA8.txt //sample file
│       │       D2AFE032D2D4C54E69AC11DE2CB3EC54.txt
│       │
│       ├───Requests
│       ├───RequestsResponse
│       ├───Sent
│       │       5.CloudCoin.1.4181672.stack //sample file
│       │
│       ├───Suspect
│       ├───Trash
│       │       sean5.skywallet.cc.png //sample file
│       │
│       ├───TrustedTransfer
│       └───Vault
├───Backups
├───Brands
│   └───CloudCoin
│           arrow.png  //Sample files
│           arrowleft.png
│           arrowright.png
│           bglogo.png
│           card.png
│           gears.png
│           jpeg1.jpg
│           Montserrat-SemiBold.otf
│           OpenSans-Regular.ttf
│           OpenSans-Semibold.ttf
│           OverpassMono-Regular.ttf
│
├───CloudBank
│   └───Keys
├───Downloads
├───EmailTemplates
├───ID
│       sean4.skywallet.cc.png //Sample ID Coin for a Skywallet account
│
├───Logs
│   │   main.log
│   │   sentcoins.csv
│   │   sentwallets.txt
│   │
│   ├───Authenticator
│   ├───Backupper
│   ├───ChangeMaker
│   ├───Echoer
│   │       0_ready_808_0.129.txt
│   │       10_ready_995_0.083.txt
│   │       11_ready_552_0.138.txt
│   │       12_ready_1318_0.068.txt
│   │       13_ready_582_0.077.txt
│   │       14_ready_703_0.199.txt
│   │       15_ready_1863_0.152.txt
│   │       16_ready_1345_0.077.txt
│   │       17_ready_611_0.04.txt
│   │       18_ready_577_0.134.txt
│   │       19_ready_1193_0.307.txt
│   │       1_ready_728_0.112.txt
│   │       20_ready_673_0.107.txt
│   │       21_ready_1504_0.083.txt
│   │       22_ready_246_0.295.txt
│   │       23_ready_634_0.223.txt
│   │       24_ready_721_0.203.txt
│   │       2_ready_1986_0.151.txt
│   │       3_ready_636_0.503.txt
│   │       4_ready_654_0.299.txt
│   │       5_ready_797_0.175.txt
│   │       6_ready_850_0.104.txt
│   │       7_ready_871_0.257.txt
│   │       8_ready_1492_0.063.txt
│   │       9_ready_659_0.112.txt
│   │
│   ├───Emailer
│   ├───Eraser
│   ├───Exporter
│   ├───FrackFixer
│   ├───Grader
│   ├───LossFixer
│   ├───Receiver
│   ├───Recoverer
│   ├───Sender
│   ├───ShowCoins
│   │       ShowCoins.log
│   │
│   ├───ShowEnvelopeCoins
│   │       ShowEnvelopeCoins.log
│   │
│   ├───Transfer
│   ├───Unpacker
│   └───Vaulter
├───PaidForRecovered
├───Recovered
├───Recovery
├───Templates
│       card.png
│       jpeg1.jpg
│       jpeg100.jpg
│       jpeg25.jpg
│       jpeg250.jpg
│       jpeg5.jpg
│       template.png
│
└───Trash
```


## Functionality Standards


## Testing Standards
Testing is very difficult with CloudCoin Applications because of the astonomical combinations of things that can and will go wrong. 
The user may have no, slow or intermitten network connectivity. They may make a request and lose Internet service 
before the RAIDA can respond. Coins and become "Fracked" (out of sync) when some RAIDA think the coin is authentic and other RAIDA
thing the coin is counterfeit. Here is how we have figured out how to test. The program could crash or lose electricity before the tasks
are completed. 

CONFIG FILE
1. How many RAIDA need to be good to authenticate
2. How many seconds to wait to authenticate

### GENERAL
1. The program should have the excact file structure as specified above.
2. The program should have a disclaimer when starting: 
```
CloudCoin Founders Edition 
Version: April 9 2020 
Used to Authenticate, Store and Payout CloudCoins  
This software is provided as is with all faults, defects, errors, and without warranty of any kind. Free from the
CloudCoin Consortium. 			

```

3. All files in the bank folder and the fracked folder should use the nameing standard of a single coin.
4. The files in the folders should be formatted. http://cloudcoinconsortium.org/standards.html#JSON_File
5. The files in the "Lost" folder may cointian the "PAN" array. 

Echo RAIDA
1. The program should show the status of the RAIDA and let the person know if there are not enough RAIDA to operate. The colors are: Grey: Unchecked or in the process of checking connectivity. Green: Connected to the RAIDA. Red: Unable to connect to the RAIDA. You can test this by going to your host file and adding a record. The location in Windows;
```File
C:\Windows\System32\Drivers\etc\hosts
```
Record to disable RAIDA 12:
```
127.0.0.1 raida12.cloudcoin.global
```
Remember to disable this by putting a '#' infront of the line when you are finished testing. 
3. If there are more than four RAIDA offline, the program should tell the user that not-enough RAIDA can be contacted to authenticate.
You can test this by adding the following records to you host file:
```
127.0.0.1 raida12.cloudcoin.global
127.0.0.1 raida13.cloudcoin.global
127.0.0.1 raida14.cloudcoin.global
127.0.0.1 raida15.cloudcoin.global

```

SHOW BALANCE
1. The balance is the coins in the bank added to the coins in the fracked folder. 
   2. 

DEPOSIT / IMPORT
1. Import the same coin 
2. Import coin that is missing pown, aoid, and ed (Should still work)
3. Import coin that has quotes around the "1" network number or one that has no quotes. 
4. Import coin with bad formatting in stack file. 
5. Import coin that is in a jpeg
6. Import coin that is a PNG.
7. Import coin that is an ID coin if Skywallat.
8. Files are in a zip file and must be unzipped to work.
9. After the coin is unpacked, the files should be moved from where they are and put in the Imported folder. 
10. What if they try to import them two times? File should overwrite the last one that was there. 
11. Deposit 2000 CloudCoin Notes. (Large amounts)

FIX

WITHDRAW
1. The coins in the bank folder should be withdrawn before the coins in the fracked folder.
2. The withdraw options should be PNG, Stack and Single Stacks put in a zip file. 
3. If a person enters the world "Random" into the tag, the program should generate a random number for the tags.
4. Coins that are exported should be able to import again.
5. Import the coins into the old desktop (CloudCoin Wallet)
6. Import the coins 




