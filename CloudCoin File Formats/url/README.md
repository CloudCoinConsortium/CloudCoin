# URLS: RawURL, CompressedURL, MultiCompressedURL

There are three standards of creating URLs:

[Raw URL](README.md#raw-url)

[Compressed URL](README.md#compressed-url)

[Mulit Compressed URL](README.md#multi-compressed-url)


CloudCoins can be stored as URLs in order for them to be sent to webpages that can change them into other files. 
The format is for GET Requests. The smaller urls can be encoded in QR codes for quick reading. 

## Raw URL

The Raw URL is simple to create. It just creates an HTTPS GET command with the parameters of a CloudCoin. 

Sample CloudCoin in a GET Request: 

```html
https://cloudcoin.global/cashout?nn=1&sn=16777216&an0=a795533814144f2d9f4e86bfc30f5760&an1=a795533814144f2d9f4e86bfc30f5760&an2=a795533814144f2d9f4e86bfc30f5760&an3=a795533814144f2d9f4e86bfc30f5760&an4=a795533814144f2d9f4e86bfc30f5760&an5=a795533814144f2d9f4e86bfc30f5760&an6=a795533814144f2d9f4e86bfc30f5760&an7=a795533814144f2d9f4e86bfc30f5760&an8=a795533814144f2d9f4e86bfc30f5760&an9=a795533814144f2d9f4e86bfc30f5760&an10=a795533814144f2d9f4e86bfc30f5760&an11=a795533814144f2d9f4e86bfc30f5760&an12=a795533814144f2d9f4e86bfc30f5760&an13=a795533814144f2d9f4e86bfc30f5760&an14=a795533814144f2d9f4e86bfc30f5760&an15=a795533814144f2d9f4e86bfc30f5760&an16=a795533814144f2d9f4e86bfc30f5760&an17=a795533814144f2d9f4e86bfc30f5760&an18=a795533814144f2d9f4e86bfc30f5760&an19=a795533814144f2d9f4e86bfc30f5760&an20=a795533814144f2d9f4e86bfc30f5760&an21=a795533814144f2d9f4e86bfc30f5760&an22=a795533814144f2d9f4e86bfc30f5760&an23=a795533814144f2d9f4e86bfc30f5760&an24=a795533814144f2d9f4e86bfc30f5760&ed=02-25&pown=uuuuuuuuuuuuuuuuuuuuuuuuu

```
The GET parameters are:
```html
nn,sn,an0,an1,an2,an3,an4,an5,an6,an7,an8,an9,an10,an11,an12,an13,an14,an15,an16,an17,an18,an19,an20,an21,an22,an23,an24,ed,pown
```
## Compressed URL

The compressed url format is designed to allow the user to either quickly create a debit card or use the seed included to create a coin that can be downloaded froma website or send to a SkyWallet account. The algorithm for creating these urls is in the Card format: 
https://github.com/CloudCoinConsortium/CloudCoin/tree/master/CloudCoin%20File%20Formats/card

EXAMPLE OF A COMPRESSED URL using GET:
```http
https://myDomain.com/coinclean?n=1&s=16777200&r=9014524598365154&d=02-25&c=8925
```
The GET parameters are:
```html
'n' is the network number, 's' is the serial number, 'r' is the card number, 'd' is the expiration date. 'c' is the CVV or PIN number. 
```

## Mulit Compressed URL

