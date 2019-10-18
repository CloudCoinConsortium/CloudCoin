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

| yte Numbers Starting with Zero   | Second Header |First Header  | Second Header |First Header  | Second Header |
| ------------- | ------------- | ------------- | ------------- | ------------- | ------------- |
| Content Cell  | Content Cell  |Content Cell  | Content Cell  |Content Cell  | Content Cell  |
| Content Cell  | Content Cell  | Content Cell  | Content Cell  | Content Cell  | Content Cell  |


| Byte Numbers Starting with Zero        | Bytes    | File Bytes  | Mandatory Hex Value | Values vary? |
| ------------- |:-------------:| -----:|
| 0 & 1     | 2 | SOI (Start of Image)  | Always the first bytes | FF D8 | No |
| 4 & 5     | 2 |   APP0 Length | Shows the Length of the APP0 Marker Excluding the first 2 bytes. In CloudCoins, this Equals 16+ 435 16 is fixed APPO stuff and 435 is the CloudCoin. | 01 C3 | No |
|6 => 10 | 5    | 	File identity | Writes "JFIF" in ASCII, terminated by a null byte | 4A 46 49 46 00 | No |
