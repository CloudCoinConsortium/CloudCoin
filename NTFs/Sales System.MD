# NFT Sales Site Phase 1.

The Sales site will consist of the following Components:

1. A modified "Buy Page" that will be put on "Celebrium.com"

2. A "Create NFT" page that will also located on Celebrium.com

3. A modified Skywallet ATM (Celebrium Wallet) that will allow people to see their NFTs. 



## Phase 1 Details

1. Modified buy page. 
Upon loading, the php page will check the NFT folder where all the NFTs will be place. NFTs are just stack files but with a different naming conention. The naming convention for the NFTs will be the name of the NFT, the price of the NFT and the .stack extention. So if I have a NFT jpeg called "Victoria Worthington" that is forsale for $1,000, then three files will be created called:
```
"Victoria Worthington.1000.stack"  
 "Victoria Worthington.description.txt"
 "Victoria Worthington.download.jpg"
 "Victoria Worthington.consignment.txt
 ```
 
 
It could be that people want to make many copies of their NFT like a limited series of 100. Then they would have to name them "Victoria Worthington 1.1000.stack"

Users would see the NFTs. If the NFTs were not images but media files such as MP3s or PDFs, people could download them and listen or look. 

There is a check box next to every NFT. People check it, the total is added below (no cart). Then they choose PayPal or Stripe. Click the Payment Button. Make their payments, then they get their stack files. 

2. "Create Your Own NFT" page. 
The user goes to the Create NFT page. They see the following except nicer and with instructions. https://www.cloudcoin.digital/nft Upon completion, they are shown their NFT and they can click on it and download it. The components of this page include the CloudCoin upload and MB indicator. The File Uploader, the Meta Data Advanced option and the instruction video. 

3. Skywallet Modification. 
For phase one we will use the existing Skywallet. 
We will add a button to the SKywallet that says "Safe Deposit Box". If the click on that, they will see a list of all the NFTs in there. If the click on an NFT, they can see or download the media of that NFT. 

# Phase 2 "Consignment" page.
NFT Sales Site Phase 2. To begin immeditaly after completion of phase 1. 

The addition of a "Consignment" page to the Celebrum site that allows people to upload their NFTs and name their price. The can use the "Create NFT" page to immeditaly put their NFTs on the market. 


People go to the Consignment page and can upload their NFTs with their price in CC or in $$ and description. 
These are placed in the NFT folder. 

Another file will go into the "Consignments" folder that shows who is the owner the NFT, their email and their address to send them a check. And or Their Skywalle address. This file will be in TOML format. When they sell, there will be a record in PayPal or Stripe. 

We will have a person who will go through this list every week and send payments to the seller. 

