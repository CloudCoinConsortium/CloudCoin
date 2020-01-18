# QR Codes

There are different ways to create QR Codes including just turning a URL into a QR code. Those QR codes can be very large. 

To create a small QR code that is easy to read, we use the following standard. 

## Application QR

The QR codes are designed to hold a bunch of data in a small size so that it is easy to read and easy to place on paper or packaging. 

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
We do not need the first three because the app will know that it must be "901". We can get rid of the Lhun number because we do not have to check for correction. So the "4" right before the comman can be disguarded.  

The line above becomes:
```
4 5245 9836 515, 8925
```
Then we remove the formatting characters and get:
```
4524598365158925
```
Now we have 16 characters. You will need to validate this because there should always be 16 characters, no more and no less. The we will add the coin's serial number to the end: (Here the SN is 16777216).
```
452459836515892516777216

```
Now, to generate a QR code,  you can go to a site like https://www.qrcode-monkey.com/#text and change into text mode if needed. Than paste the code above and create the QR code. You will get a small code that can be placed on any item. 

[Creating a CSV File](https://github.com/CloudCoinConsortium/CloudCoin/blob/master/CloudCoin%20File%20Formats/card/README.md#saving-the-coins-to-file)
# Sample QR Code

![Without Logo](qr-code.png)

# Sample QR Code with Logo

![QR code with a log0](qr-code-logo.png)
