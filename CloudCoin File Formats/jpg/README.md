# JPEG or JPG File Format
The JPG file format allows one CloudCoin to be stored into one Jpeg image. 

## Naming Convention

Denomination "CloudCoin", NN, SN, a user tag and ".jpg" separated by dots.

Example Filename:
```
250.cloudcoin.1.14589548.mintTagHere.jpg
```
Denomination "CloudCoin", NN, SN, a user tag and ".jpg" separated by dots.


![JPEG Format](jpegformat.jpg

NOTE: The first byte of the AOID tells if the coin is encrypted. EE means it is encrypted. 
If it is encrypted, then the next 25 hex numbers will represent if the parts of the AN are positive or negative. 
0000b or 0x0 means they are all positive. 1111b or 0xf means they are all negative. There are 16 possible combinations. 
for example 1010b or 0x9 means that the first and third AN parts are negative. See the vaulter for more information about this. 

## Explanation of Bytes

| Byte Numbers Starting with Zero | Name  | Bytes | File Bytes  | Mandatory Hex Valuer | Values vary? | 
| ------------- | ------------- | ------------- | ------------- | ------------- | ------------- |
| 0 & 1   |2 | SOI (Start of Image) | Always the first bytes  | FF D8   | No |
| 4 & 5  | Content Cell  | Content Cell  | Content Cell  | Content Cell  | Content Cell  |


## Offical CloudCoin Templates

![1 Template](jpeg1.jpg)
![5 Template](jpeg5.jpg)
![25 Template](jpeg25.jpg)
![100 Template](jpeg100.jpg)
![250 Template](jpeg250.jpg)
