# Local Wallet Standards
A Local Wallet stores CloudCoins on a local device. This is the opposite of a Skywallet that stores CloudCoins on a RAIDA. The main features of a Local Wallet are to check for authenticity, 
change ownership, place good coins in a protected folder and fix coins that are unsyncronized (fracked).

[Folder Standards](README.md#folder-standards)

[Functionality Standards](README.md#functionality-standards)

[Testing Standards](README.md#testing-standards)

## Folder Standards
The folder structure is the backbone of the Local Wallet and we would like to keep it the same forever. 

Here is the folder structure:
```html
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
│       ├───Lost
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
