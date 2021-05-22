# NFTs
Each CloudCoin can hold one NFT. Data can only be uploaded once. Data can be deleted. Once deleted, it cannot be written to again because deletion deletes the data in a coin but not the coin.  

WARNING: If an NTFs is not authenticated at least once per five years, it will be considered lost and will be recovered by the RAIDA and sold if it has any value. 


[SQL](README.md#sql)

[Storage Limits](README.md#storage-limits)

[Storage Protocols](README.md#storage-protocols)

Services
[has_nft](README.md#has_nft)

[Read](README.md#read)

[Insert](README.md#insert)

[Delete](README.md#delete)

Mass NFT Upload 

[Mass NFT Upload](README.md#mass-nft-upload)

## NFTs SQL
```

mysql> describe nft;
+--------------+---------------+------+-----+-------------------+-----------------------------+
| Field        | Type          | Null | Key | Default           | Extra                       |
+--------------+---------------+------+-----+-------------------+-----------------------------+
| sn           | int(11)       | NO   | PRI | NULL              |                             |
| protocol     | int(11)       | YES  |     | NULL              |                             |
| stripe       | text          | YES  |     | NULL              |                             |
| mirror       | text          | YES  |     | NULL              |                             |
| mirror2      | text          | YES  |     | NULL              |                             |
| created      | timestamp     | NO   |     | CURRENT_TIMESTAMP |                             |
+--------------+---------------+------+-----+-------------------+-----------------------------+
6 rows in set (0.00 sec)

```

## Storage Rules
The larger the CloudCoin, the more data can store. 

Note Size | Max Amount of Stripe Data that can be stored | Stripe Data Per RAIDA
------|-----------------|----
250 | 5 MB | 200 KB
100 | 2 MB| 80 KB
 25 | 500 KB| 20 KB
 5 | 250 KB| 10 KB
1 | 50 KB| 2 KB

## Storage Protocols
The user may store their data any way that they like and there may be standards that develop over time. The user can specify the standard with a number called the
Storage Protocol.
The standard handles compression, RAID, encryption, file formatting and how to extract meta data. Now there is only one standard. 
Any protocol that is under 1000, will have its mirrors and stripes sycned. If you don't want this use a higher number than 1000. 

Protocol Number | Description
---|---
-1| Deleted. This means that the data that was in the NFT has been deleted. Now it is impossible to put new data there. 
0 | There is no compression or encryption. The data is striped, mirrored and then mirrored again. The stripe will be on RAIDA n, the mirror on RAIDA n-3 and the second mirror on RAIDA n-6. The data that is stored is broken into two parts that are seperated by an asterik * The first part is the meta data formatted in TOML and encoded in base64. The second part is the file data encoded using base64.  Sample ZmlsZV9uYW1lPUdvbmUgV0l0aCB0aGUgV2luZA==*SXQncyBhIGdyZWF0IG1vdmVp
1| Includes Proof.Same as protocol 0 but has a base64 image that is used as a proof. Sample: ZmlsZV9uYW1lPUdvbmUgV0l0aCB0aGUgV2luZA==*SXQncyBhIGdyZWF0IG1vdmVp*SW1hZ2Ugb2YgbWU= where the third and last part is the proof image. 

### Protocol 0 Meta Data
The Protocol 0 only requires one field: the file name.
```
file_name = "Fun in the Sun.png"
```
### Protocol 1 Meta Data
The Protocol 1 requires two fields: the file name and the proof name.
```
file_name = "Fun in the Sun.png"
proof_name = "SeanProffFirstNFT.jpg"
```

![](SeanProofFirstNFT.jpg)

## Services
There are four services
0. has_nft
1. Read
2. Insert
3. Delete


## has_nft
has_nft takes an array of Serial numbers and an array of authentication numbers. It returns a comma seperated value file. 
Sample GET Call:
```
https://raida11.cloudcoin.global/service/nft/has_nft?sn[]=14729186&an[]=C8E01BCA9A22FA0250DCC65C3DBC76E1&sn[]=1&an[]=50378FA6A36D9498D90C09D2608DA63E

```
Sample Reponse: 
```
true,false

```

## Read


Sample GET Call:
```
https://raida0.cloudcoin.global/service/nft/read?sn=8867&an=9dfa64eb6c774635b5ac3e643e8100f1

```
Note that nn, denomination are not needed and are being phased out. PAN is not needed because this service does not change the AN. This service only allows for one coin 
to be detected at once. 

Sample Reqsponse:
```json
{
	"server":"raida9",
	"status":"pass",
	"message": "The attached data belongs to the token", 
	"storage_protocol": 0,
	"stripe":"b3ZpcG9pd2VyO2xtZ",
	"mirror":"Um5wro7urS7LbunvC",
	"mirror2":"mpsO2VqcmxrZWpyIH",
	"time":"2019-11-08 05:56:32",
   	"ex_time":"8.577"
}

```


## Insert
Each RAIDA will write 

Sample POST Request: 
```
https://raida0.cloudcoin.global/service/ntf/insert?
sn=8867&
an=9dfa64eb6c774635b5ac3e643e8100f1
stripe=b3ZpcG9pd2VyO2xtZ
mirror=Um5wro7urS7LbunvC
mirror2=mpsO2VqcmxrZWpyIH
```

Sample Response
```
Success
```


## Delete
Deletes all the data in an NFT but not the record 
So the coin can never be used again for an NFT but does allow the owner to get rid of data they do not like. 

Sample GET Request 
```
https://Raida0.Cloudcoin.global/service/nft/delete?sn=677&an=afbb46743568964cf
```
Sample Response 
```
Success
```

# Mass NFT Upload

This is a protocol for uploading many NFTs at once. All of the files that are needed are first placed in a default folder. 
Then all the names of the files are placed in a CSV file. 
The Program will be given the csv and will look in the default location for the files. Then, one by one the files will be uploaded. 
A report is returned.

Example Table:

CloudCoin File |NFT File | Proof File (Optional
----------|---------|------------
250.CloudCoin.1.1478526.stack | Art1.png | Art_proof.jpg
250.CloudCoin.1.1478526.stack | Art2.png | Art_proof.jpg
250.CloudCoin.1.1478526.stack | Art3.png | Art_proof.jpg
250.CloudCoin.1.1478526.stack | Contract.pdf | Contract_proof.jpg
250.CloudCoin.1.1478526.stack | game.exe | game_proof.jpg
 
Example csv file:
```csv
250.CloudCoin.1.1478526.stack, Art1.png, Art_proof.jpg
250.CloudCoin.1.1478526.stack, Art2.png, Art_proof.jpg
250.CloudCoin.1.1478526.stack, Art3.png, Art_proof.jpg
250.CloudCoin.1.1478526.stack, Contract.pdf, Contract_proof.jpg
250.CloudCoin.1.1478526.stack, game.exe, game_proof.jpg

```
Example Report Returned to the user:
```
Attempting 250.CloudCoin.1.1478526.stack, Art1.png, Art_proof.jpg. SUCCESS
Attempting 250.CloudCoin.1.1478526.stack, Art2.png, Art_proof.jpg. ALL PASS EXCEPT RAIDA 5
Attempting 250.CloudCoin.1.1478526.stack | Art3.png | Art_proof.jpg ALL PASS
Attempting 250.CloudCoin.1.1478526.stack | Contract.pdf | Contract_proof.jpg ALL PASS EXCEPT RAIDA 5, 13
Attempting 250.CloudCoin.1.1478526.stack | game.exe | game_proof.jpg ERROR
```
