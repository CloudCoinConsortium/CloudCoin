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
│   └───sean4.skywallet.cc
│       sean4.skywallet.cc.png //Sample ID Coin for a Skywallet account
│       15587724.config //serial number of the ID coin contains hash of password
│       transaction.txt //Transaction log
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

CONFIG FILE (Ignore this for now)
1. How many RAIDA need to be good to authenticate
2. How many seconds to wait to authenticate

### GENERAL
1. The program should have the excact file structure as specified above.
2. The program should have a disclaimer when starting: 
### Disclaimer
```
CloudCoin Founders Edition 
Version: April 9 2020 
Used to Authenticate, Store and Payout CloudCoins  
This software is provided as is with all faults, defects, errors, and without warranty of any kind. Free from the
CloudCoin Consortium. 			

```
3. The program should never delete coin files, only move them to places like Imported and Trash. 
3. All files in the bank folder and the fracked folder should use the nameing standard of a single coin.
4. The files in the folders should be formatted. http://cloudcoinconsortium.org/standards.html#JSON_File
5. The files in the "Lost" folder may cointian the "PAN" array. 

### ECHO RAIDA
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

### SHOW COINS
Show coins is a command that returns how many coins are in the program. 
1. The balance is the sum of the coins in the "Bank", "Vault" and "Fracked" folders. To test this, put some properly named CloudCoins in to the Bank and Fracked folder and Vault folders see if their sum is displayed.  
2. Some programs may also show the number of coins in the Limbo folder but not as authentic coins that are owned. Some programs show Fracked coins but should make it clear that they are authentic. 
3. Show Coins should never try to track the coins in RAM. Instead, everytime the ShowCoins is called, the folders should be checked. To test this remove one of the stack files in the Bank folder and call the show coins again. See if the balance has changed accourdinly. 

### SHOW FOLDERS
If this feature is enabled, it should show the location of the following: 
1. The root program file.
2. The Logs folder (in the root) that tracks all activites.
3. The account folder (of the active account)
4. The Templates folder of the active account.
5. ID folder for Skywallet ID Coins


### DEPOSIT / IMPORT
This is the process of changing the coins ANs to new PANs thus authenticaing and powning (password owning) them in the process. 
1. Import the same coin. When you import the same coin, the program should check the Bank, Fracked, Limbo, Vault and ID folder. If they are in one of those folders, the user should be told that the coin was already in the "--" folder and that it was moved to Trash. You should find the coin in the Trash Folder. This process should be very fast because with the exception of the ID coins, only the file name has to be read.  
2. Only files with the following extensions should be allowed to be imported: .stack, .png, .jpeg, .zip. All other file formats should be ignored. Test this by brining in other file formats.
3. Zipped files should be unzipped. Then if they have files within them that are not.stack, .png, or .jpeg, those files should be ignored. 
5. .stack files should be able to be imported using four different formats, The old format, new format, compact format, jumbled format and the oops format (when developers do not follow standard). You can copy and past these into a file and import them to see if they work. You just need to give them a different serial number. They will all be counterfeit but that is not what we are testing. 

#### MAIN FORMAT
```javascript
{
	"cloudcoin": [{
		"nn": "1",
		"sn": "16776215",
		"an": [
			"be0f94e584584d85ab86c301b8f1906d", "1e539633bc974bf88ebaeb7fd23279d7", "d735e4d16ff54eda8019e1a59ad33032", "b394553ff4e94f78a5f4c04df4102a20", "85b8f6956deb4dc28bc40b192ef5b0bc",
			"6775856024a14d10af30181477f8957a", "31274a1103c7406fb0f60736604f23b6", "6eb2c9bfb21e484b89f35c8124a8f3c0", "407b9d6b0b034b23b1d160d7da96a8ae", "2dbf38c2b1834287aed0fdcf4ba32274",
			"4da74b2a8c784731826716dd36296271", "0fe3fb2ded9648eeae64bdcf29abdf15", "01bdfe3b43b74b618c5db44e6b2febb6", "e0bdf8a1b2eb4174bcca6fda0c3609a1", "c8636145d53d42008c84432872816d4c",
			"f5258cb2cd2f4afca333bc7cd64f01eb", "c89eb2da6897489db2581d0b17d4d959", "89d5dba265c84a488106840a79316940", "8dba35338678464a9ef9a1686cd500af", "0a39da32d43546009147127443793297",
			"109a745bb88946a3ac9720bc8f9fee66", "a56933d0ab01400d93782c4eba0c6849", "ef6025e09f88499fb0490edf7ae21eba", "6014a91a938f449cb8f72658e0b9d61e", "cb64b024fe8c47d19d3fcc5c797fa0b3"
		],
		"pown": "ppeppppppfppppppnpppupppp",
		"ed": "9-2021",
		"aoid": []
	}
}

```
#### Compact Standard
```javascript
{"cloudcoin":[
{"nn":"1",
"sn":"16777213","an":[
"be0f94e584584d85ab86c301b8f1906d","1e539633bc974bf88ebaeb7fd23279d7","d735e4d16ff54eda8019e1a59ad33032","b394553ff4e94f78a5f4c04df4102a20","85b8f6956deb4dc28bc40b192ef5b0bc","6775856024a14d10af30181477f8957a","31274a1103c7406fb0f60736604f23b6","6eb2c9bfb21e484b89f35c8124a8f3c0","407b9d6b0b034b23b1d160d7da96a8ae","2dbf38c2b1834287aed0fdcf4ba32274","4da74b2a8c784731826716dd36296271","0fe3fb2ded9648eeae64bdcf29abdf15","01bdfe3b43b74b618c5db44e6b2febb6","e0bdf8a1b2eb4174bcca6fda0c3609a1","c8636145d53d42008c84432872816d4c","f5258cb2cd2f4afca333bc7cd64f01eb","c89eb2da6897489db2581d0b17d4d959","89d5dba265c84a488106840a79316940","8dba35338678464a9ef9a1686cd500af","0a39da32d43546009147127443793297","109a745bb88946a3ac9720bc8f9fee66","a56933d0ab01400d93782c4eba0c6849","ef6025e09f88499fb0490edf7ae21eba","6014a91a938f449cb8f72658e0b9d61e","cb64b024fe8c47d19d3fcc5c797fa0b3"],
"pown":"ppeppppppfppppppnpppupppp",
"ed":"9-2021","aoid":[]}
]}
```
#### Old Standard
```javascript
{"cloudcoin":[
{"nn":"1",
"sn":"16777213","an":[
"be0f94e584584d85ab86c301b8f1906d","1e539633bc974bf88ebaeb7fd23279d7","d735e4d16ff54eda8019e1a59ad33032","b394553ff4e94f78a5f4c04df4102a20","85b8f6956deb4dc28bc40b192ef5b0bc","6775856024a14d10af30181477f8957a","31274a1103c7406fb0f60736604f23b6","6eb2c9bfb21e484b89f35c8124a8f3c0","407b9d6b0b034b23b1d160d7da96a8ae","2dbf38c2b1834287aed0fdcf4ba32274","4da74b2a8c784731826716dd36296271","0fe3fb2ded9648eeae64bdcf29abdf15","01bdfe3b43b74b618c5db44e6b2febb6","e0bdf8a1b2eb4174bcca6fda0c3609a1","c8636145d53d42008c84432872816d4c","f5258cb2cd2f4afca333bc7cd64f01eb","c89eb2da6897489db2581d0b17d4d959","89d5dba265c84a488106840a79316940","8dba35338678464a9ef9a1686cd500af","0a39da32d43546009147127443793297","109a745bb88946a3ac9720bc8f9fee66","a56933d0ab01400d93782c4eba0c6849","ef6025e09f88499fb0490edf7ae21eba","6014a91a938f449cb8f72658e0b9d61e","cb64b024fe8c47d19d3fcc5c797fa0b3"]}
]}
```
#### Jumbled Version (The elements are in a different order)
```javascript
{"cloudcoin":[
{"sn":"16777213","an":[
"be0f94e584584d85ab86c301b8f1906d","1e539633bc974bf88ebaeb7fd23279d7","d735e4d16ff54eda8019e1a59ad33032","b394553ff4e94f78a5f4c04df4102a20","85b8f6956deb4dc28bc40b192ef5b0bc","6775856024a14d10af30181477f8957a","31274a1103c7406fb0f60736604f23b6","6eb2c9bfb21e484b89f35c8124a8f3c0","407b9d6b0b034b23b1d160d7da96a8ae","2dbf38c2b1834287aed0fdcf4ba32274","4da74b2a8c784731826716dd36296271","0fe3fb2ded9648eeae64bdcf29abdf15","01bdfe3b43b74b618c5db44e6b2febb6","e0bdf8a1b2eb4174bcca6fda0c3609a1","c8636145d53d42008c84432872816d4c","f5258cb2cd2f4afca333bc7cd64f01eb","c89eb2da6897489db2581d0b17d4d959","89d5dba265c84a488106840a79316940","8dba35338678464a9ef9a1686cd500af","0a39da32d43546009147127443793297","109a745bb88946a3ac9720bc8f9fee66","a56933d0ab01400d93782c4eba0c6849","ef6025e09f88499fb0490edf7ae21eba","6014a91a938f449cb8f72658e0b9d61e","cb64b024fe8c47d19d3fcc5c797fa0b3"],
"nn":"1"}
]}
```
#### Oops Version (The programmer did not notice that the network number and serial number have quotation marks around them)
```javascript
{"cloudcoin":[
{"nn":"1",
"sn":16777213,an:[
"be0f94e584584d85ab86c301b8f1906d","1e539633bc974bf88ebaeb7fd23279d7","d735e4d16ff54eda8019e1a59ad33032","b394553ff4e94f78a5f4c04df4102a20","85b8f6956deb4dc28bc40b192ef5b0bc","6775856024a14d10af30181477f8957a","31274a1103c7406fb0f60736604f23b6","6eb2c9bfb21e484b89f35c8124a8f3c0","407b9d6b0b034b23b1d160d7da96a8ae","2dbf38c2b1834287aed0fdcf4ba32274","4da74b2a8c784731826716dd36296271","0fe3fb2ded9648eeae64bdcf29abdf15","01bdfe3b43b74b618c5db44e6b2febb6","e0bdf8a1b2eb4174bcca6fda0c3609a1","c8636145d53d42008c84432872816d4c","f5258cb2cd2f4afca333bc7cd64f01eb","c89eb2da6897489db2581d0b17d4d959","89d5dba265c84a488106840a79316940","8dba35338678464a9ef9a1686cd500af","0a39da32d43546009147127443793297","109a745bb88946a3ac9720bc8f9fee66","a56933d0ab01400d93782c4eba0c6849","ef6025e09f88499fb0490edf7ae21eba","6014a91a938f449cb8f72658e0b9d61e","cb64b024fe8c47d19d3fcc5c797fa0b3"]}
]}
```
#### LIMBO / LOST FORMAT
```javascript
{
	"cloudcoin": [{
		"nn": "1",
		"sn": "16776215",
		"an": [
			"be0f94e584584d85ab86c301b8f1906d", "1e539633bc974bf88ebaeb7fd23279d7", "d735e4d16ff54eda8019e1a59ad33032", "b394553ff4e94f78a5f4c04df4102a20", "85b8f6956deb4dc28bc40b192ef5b0bc",
			"6775856024a14d10af30181477f8957a", "31274a1103c7406fb0f60736604f23b6", "6eb2c9bfb21e484b89f35c8124a8f3c0", "407b9d6b0b034b23b1d160d7da96a8ae", "2dbf38c2b1834287aed0fdcf4ba32274",
			"4da74b2a8c784731826716dd36296271", "0fe3fb2ded9648eeae64bdcf29abdf15", "01bdfe3b43b74b618c5db44e6b2febb6", "e0bdf8a1b2eb4174bcca6fda0c3609a1", "c8636145d53d42008c84432872816d4c",
			"f5258cb2cd2f4afca333bc7cd64f01eb", "c89eb2da6897489db2581d0b17d4d959", "89d5dba265c84a488106840a79316940", "8dba35338678464a9ef9a1686cd500af", "0a39da32d43546009147127443793297",
			"109a745bb88946a3ac9720bc8f9fee66", "a56933d0ab01400d93782c4eba0c6849", "ef6025e09f88499fb0490edf7ae21eba", "6014a91a938f449cb8f72658e0b9d61e", "cb64b024fe8c47d19d3fcc5c797fa0b3"
		],
		"pan": [
			"be0f94e584584d85ab86c301b8f1906d", "1e539633bc974bf88ebaeb7fd23279d7", "d735e4d16ff54eda8019e1a59ad33032", "b394553ff4e94f78a5f4c04df4102a20", "85b8f6956deb4dc28bc40b192ef5b0bc",
			"6775856024a14d10af30181477f8957a", "31274a1103c7406fb0f60736604f23b6", "6eb2c9bfb21e484b89f35c8124a8f3c0", "407b9d6b0b034b23b1d160d7da96a8ae", "2dbf38c2b1834287aed0fdcf4ba32274",
			"4da74b2a8c784731826716dd36296271", "0fe3fb2ded9648eeae64bdcf29abdf15", "01bdfe3b43b74b618c5db44e6b2febb6", "e0bdf8a1b2eb4174bcca6fda0c3609a1", "c8636145d53d42008c84432872816d4c",
			"f5258cb2cd2f4afca333bc7cd64f01eb", "c89eb2da6897489db2581d0b17d4d959", "89d5dba265c84a488106840a79316940", "8dba35338678464a9ef9a1686cd500af", "0a39da32d43546009147127443793297",
			"109a745bb88946a3ac9720bc8f9fee66", "a56933d0ab01400d93782c4eba0c6849", "ef6025e09f88499fb0490edf7ae21eba", "6014a91a938f449cb8f72658e0b9d61e", "cb64b024fe8c47d19d3fcc5c797fa0b3"
		],
		"pown": "ppeppppppfppppppnpppupppp",
		"ed": "9-2021",
		"aoid": []
	}
}

```

8. Import coin that is in a jpeg
9. Import coin that is a PNG.
10. Import coin that is an ID coin if Skywallats are received. If the program can go and get CloudCoins from a Skywallet, put the Skywallet coin into the ID folder and then try to import a copy of this coin. The program should not allow it. 
13. After the coin is unpacked, the files should be moved from where they are and put in the Imported folder. 
14. What if they try to import them two times? File should overwrite the last one that was there in the Imported folder.  
15. Deposit 2000 CloudCoin Notes. (Large amounts). 
16. After the deposit, the user should be shown a result that show how many coins were authentic, counterfiet, Limbo and sometime fracked. Users should also be able to see details somewhere in something called "Receipts". The Receipts should show each coin and what the outcome was including the pown which will look something like "pppppppppfppppppppupppp".
17. If the coin has 13 or more passes, it is authentic. If a coin has 13 or more fails, it is counterfeit. Otherwise, the coin is in Limbo. You can go into a stack file and change the numbers of the ANs and that will cause the coin to become fracked on that RAIDA.  


### WITHDRAW

1. When withdrawing coins, the program should take the coins from the Bank or Vault folder first and only take from the Fracked folder if the Bank and Vault are empty. This by putting coins in the Bank and Fracked folder. withdraw one note and see if it comes from the Bank or Fracked.  

2. When withdrawing, copy of the coins removed should be put in the "Exported" folder. 
3. Sometime, coins are withdrawn by sending Email or by Uploading to a Skywallet (Send). Coins sent by email should be placed in the EMAIL folder. Coins Sent to a Skywallet should be placed in the "Sent" folder. 

Users should have three choices of withdraw:
1. Stack file
2. PNG
3. Single Stack Files zipped together

#### Stack Files
1. No more than 3000 notes per stack file. Try to withdraw more. 
2. Formatted using the Main Format above. Open the file using a text editor. Note that all the white spaces (tabs, spaces and line breaks) should be exactly the same. See file formatting: https://github.com/CloudCoinConsortium/CloudCoin/tree/master/CloudCoin%20File%20Formats/stack#stack-files 
3. Named using the naming convention for many coins if there is more than one note in the stack. If there is only one coin in the stack then it should use the single not naming convention. See file formating above. 
4. You should be able to import this coin back into the software and into other CloudCoin Wallets. 

#### PNG format
1. The standard PNG template should be used: https://github.com/CloudCoinConsortium/CloudCoin/tree/master/CloudCoin%20File%20Formats/png You can open the PNG file with a text editor to check this. (NOTE: The PNG template is under development at the time of this writing.)
Should have the writing on the front showing how many coin are in the PNG.

#### Single Stack Files zipped togehter
Instead of all the notes being in one stack file, there should be one stack file for every coin. The tag of these stack files should be a random guid so that their names cannot be guessed. The coins should be all zipped together in one file. 


3. If a person enters the world "Random" into the tag, the program should generate a random number for the tags.
4. Coins that are exported should be able to import again.


### LIST SERIALS
This lists all the serial numbers of coins in a specified folder. 
1. Serial Number file format is .csv with the serail number and the denomination like this:
```csv
16777214,250
16777215,250
16777216,250
```
2. The program should be able to read all the serial numbers in a stack file. 
3. The program should be able to unzip files and then read all the serial numbers in the files within those files. 
4. Only .PNG, .stack and .zip file extensions should be looked at. All other extensions should be ignored.  

### FIX

