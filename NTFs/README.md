# NFTs
Each CloudCoin can hold one NFT. Data can only be uploaded once. Data can be deleted. Once deleted, it cannot be written to again because deletion deletes the data in a coin but not the coin.  

WARNING: If an NTFs is not authenticated at least once per five years, it will be considered lost and will be recovered by the RAIDA and sold if it has any value. 


[SQL](README.md#sql)

[Storage Limits](README.md#storage-limits)

[Storage Protocols](README.md#storage-protocols)

Services

[Read](README.md#read)

[Insert](README.md#insert)

[Delete](README.md#delete)


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
100 | 2 MB| 200 KB
 25 | 500 KB| 100 KB
 5 | 250 KB| 25 KB
1 | 50 KB| 2 KB

## Storage Protocols
The user may store their data any way that they like and there may be standards that develop over time. The user can specify the standard with a number called the
Storage Protocol.
The standard handles compression, RAID, encryption, file formatting and how to extract meta data. Now there is only one standard. 
Any protocol that is under 1000, will have its mirrors and stripes sycned. If you don't want this use a higher number than 1000. 

Protocol Number | Description
---|---
0 | There is no compression or encryption. The data is stiped, mirrored and then mirrored again. The stripe will be on RAIDA n, the mirror on RAIDA n-3 and the second mirror on RAIDA n-6. 

### Protocol 0 
The Protocol 0 only requires two fields: the file name and the data. 
Code before it is turned into base64:
```
file_name = "Fun in the Sun.png"
data = mpsO2VqcmxrZWpyIH
```


## Services
There are three services

1. Read
2. Insert
3. Delete

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




