# QR Codes

There are different ways to create QR Codes including just turning a URL into a QR code. URL QR codes can be very large because they contain the domain name and web pages and not just coin codes.

However, if we have a custom cell phone application that is reads the QR code, then we don't need a url for a browser. This gives us a much smaller QR code that is easier for cell phones to read. 

## Generated from the common CSV forat. 

The QR codes are generated from the CSV files created by the CARD Generator.

So we can take a line like this:
```
9014 5245 9836 5154, 02-25, 8925, 1, 16777200, n=1&s=16777200&r=9014524598365154&d=02-25&c=8925& 
```
And then strip away all the non-essential informtion. We only need the card number and the serial number. The application that we use to read the QR code will need to know the network number. The app will not need to have a url because the QR code is pure text (not a URL). 

So the line above becomes: 
```
9014 5245 9836 5154, 8925
```
Then we remove the formatting characters and get:
```
90145245983651548925
```
Now we have 20 characters. You will need to validate this because there should always be 20 characters, no more and no less. The we will add the coin's serial number to the end: (Here the SN is 16777216).Now the code will look like this:
```
9014524598365154892516777216

```
Now, to generate a QR code,  you can go to a site like https://www.qrcode-monkey.com/#text and change into text mode if needed. Then paste the code above and create the QR code. You will get a small code that can be placed on any item. 

[Creating a CSV File](https://github.com/CloudCoinConsortium/CloudCoin/blob/master/CloudCoin%20File%20Formats/card/README.md#saving-the-coins-to-file)
# Sample QR Code

![Without Logo](qr-code.png)

# Sample QR Code with Logo

![QR code with a log0](qr-code-logo.png)
