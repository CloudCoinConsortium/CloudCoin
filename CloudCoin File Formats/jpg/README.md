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

NOTE: The first byte of the AOID tells if the coin is encrypted. EE means it is encrypted.
If it is encrypted, then the next 25 hex numbers will represent if the parts of the AN are positive or negative.
0000b or 0x0 means they are all positive. 1111b or 0xf means they are all negative. There are 16 possible combinations.
for example 1010b or 0x9 means that the first and third AN parts are negative. See the vaulter for more information about this.

## Explanation of Bytes

| Byte Numbers Starting with Zero | Bytes  | Name | File Bytes  | Mandatory Hex Value | Values vary? |
| ------------- | ------------- | ------------- | ------------- | ------------- | ------------- |
| 0 & 1   |2 | SOI (Start of Image) | Always the first bytes  | FF D8   | No |
| 4 & 5  | 2  | APP0 Length  | Shows the Length of the APP0 Marker Excluding the first 2 bytes. In CloudCoins, this Equals 16+ 435 16 is fixed APPO stuff and 435 is the CloudCoin.  | 01 c3  | No |
| 6 => 10 | 5 | File identity | Writes "JFIF" in ASCII, terminated by a null byte | 4A 46 49 46 00 | No |
| 11 & 12 |	2 |	Version |	First byte for major version, second byte for minor version (01 02 for 1.02) | 01 01 |	No |
| 13 | 1 | Density units |	00 or 01 = pixels per inch. 02 = pixels per centimeter. |	01 | No |
| 14 & 15 |	2 |	Xdensity | ? | 00 60 | No |
| 16 & 17 |	2 |	Ydensity | ? | 00 60 | No |
| 18 | 1 | Thumbnail Horizontal Pixel Count |	CloudCoin data is in the thumbnail and is 29 pixels wide | 1D |	No |
| 19 | 1 | Thumbnail Vertical Pixel Count |	CloudCoin data is in the thumbnail and is 5 pixels high |	05 | No |
| 20 => 419 |	400 |	ANs: 25 Authenticity Numbers | 25 GUIDs without hyphens. The Jpeg thumbnail dimensions are 29 pixels * 5 pixels * 3 bytes each = 435 bytes. Each AN is 16 bytes and 25 x 16 = 400 bytes so there are 35 bytes left for other info (AOID, Has Comment, Health Status, Expiration Date, Network Number and Serial Number.) | Too Large to show | Yes |
| 420 => 435 | 16 | AOID: Account / Owner ID | Bytes can be used by the owner for their uses. | Too Large to show |	Yes |
| 436 => 448 | 13 |	POWN: Results of last pown attempt | 0 (unknown), 1 (pass),2 (no Response), E (error) or F (fail) '. | Hex 11111011001f11111111211e199 The last nibble can be ignorred | Yes |
| 449 |	1 |	HC: Has comments? |	Indicates if there are any other information in the comments section of the Jpeg that is used by applications. | 00 = No, More than zero = Yes (Program should look in the comments section of JPEG for more owner information). There is room for 255 codes so if you need to put some data it could fit here. | Yes |
| 450 |	1 |	ED: Expiration Date |	Months from zero month (August 2016). |	Hex 18 = 24 months since zero or August 2018 | Yes |
| 451 |	1 |	NN: Network Number | 1 Through 256 | 01 |	Yes |
| 452, 453 & 454 | 3 |	SN: Serial Number |	sn in Hex 6 bytes FF FFFF = 16,777,215 | FF FFFF |	Yes |


## Offical CloudCoin Templates

![1 Template](jpeg1.jpg)
![5 Template](jpeg5.jpg)
![25 Template](jpeg25.jpg)
![100 Template](jpeg100.jpg)
![250 Template](jpeg250.jpg)
