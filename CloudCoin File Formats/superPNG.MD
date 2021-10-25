
# The PNG format used in the SuperRAIDA (Non-experimental version of CloudCoin)

You will work from the past game. 

You will open and read a PNG file that has secret codes embedded in it. 
```javascript
file = file_bin_open("myfile.png", 0);
data = file_bin_read_byte(file);
file_bin_close(file);
```
You will create an object called "Header"

The header starts at byte number 0x21(This is hex)

Name | Number of Bytes | Data type | Description
---|---|---|---
Chunk Length | 4 | Integer | Tells how many bytes are in the chunk of CloudCoins.
Chunk Type | 4| String | Should be cLDc (CloudCoin)
Format Type | 1 | int | Describes how the following bytes are formatted
CloudID | 1 | int | 0 for now
Coin ID | 2 | int | 0 for now
Password Hash | 16 | Bytes converted to String | This will be hexidecimal characters. Each byte has two hex characters. The length of the string will be 32. Used to see if the password is correct. 
Flags | 1 | Int | This is a bitfield 
Receipt Number | 16 | Bytes converted to String | This will be hexidecimal characters. Each byte has two hex characters. The length of the string will be 32. Used to identify the transaction.



You will create an object called "CloudCoin".
The CloudCoin will have variables that will be created by reading the bytes from the PNG file:

The coin starts at byte number 0x21(This is hex)
Name | Number of Bytes | Data type | Description
---|---|---|---
Serial Number | 3 | Integer | The ID for one coin. 
POWN | 13 | Iteger | Bytes converted to String || This will be hexidecimal characters. Each byte has two hex characters. The length of the string will be 26. Used to identify the transaction.
Authenticity Numbers | An array 25 16 byte numbers | bytes converted bo  String of hex characers | The secret passwords embedded in the coin. 

![png](https://github.com/CloudCoinConsortium/CloudCoin/blob/master/CloudCoin%20File%20Formats/PNG2.png)

