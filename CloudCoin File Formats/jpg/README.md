# JPEG or JPG File Format
The JPG file format allows one CloudCoin to be stored into one Jpeg image. 

## Naming Convention

Denomination "CloudCoin", NN, SN, a user tag and ".jpg" separated by dots.

Example Filename:
```
250.cloudcoin.1.14589548.mintTagHere.jpg
```
Denomination "CloudCoin", NN, SN, a user tag and ".jpg" separated by dots.


![JPEG Format](jpegformat.jpg)

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
