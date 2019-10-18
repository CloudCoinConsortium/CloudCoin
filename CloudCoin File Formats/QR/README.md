# QR Codes

The QR codes are designed to hold a bunch of data in a small size so that it is easy to read and easy to place on paper or packaging. 

The QR code uses the method of generating PANs based on a seed in the QR code. To turn the coin into a CloudCoin you must be able to run an algorithm against it. 

## New Compressed CloudCoin Format

The compressed CloudCoin Format has three parts:
1. Serial Number
2. Random Letters
3. Network Number

########$$$$$$$$$$##
```
Sample:16777216FOILJNPVUE2
```
## Serial Number
The serial number is a number that can be as large as needed. 

## Random Letters
The Random Letters are capital letters between A and Z. Usually their are 10 but there could be more or less. 

## Network Number
If there is no network number, the network number is assumed to be 1. Otherwise the network number just needs to be a number and can be as large as you can imagine. 

## CREATING A CODE FOR A COIN:
The Serial Number and the Network Numbers are taken right off the coin and are not changed. 

## Generating Random Number:
For unbeatable security, Generate a string of random capitalized letters between A and Z. 

Example
```
FZLKBNGMNW
16777215az!Q2w#E

RAIDA Random     AN
R0 FZLKBNGMNW  0FZLKBNGMNW  44E16EEFDDFBB4C02E28F61D888D66F2
R1 FZLKBNGMNW  1FZLKBNGMNW  D5D3540E35ADC72F28AC3FD927C86E94
R2 FZLKBNGMNW  2FZLKBNGMNW  20E9BCFD395B659C5B0CE1E5C9FFDE25
... R3 - R9 left out for brevity
R23 FZLKBNGMNW 23FZLKBNGMNW  7B3D77AB3214D247C9D46D4BCECDC9A7
R24 FZLKBNGMNW 24FZLKBNGMNW  143D6B3585DFCDACF8C721FDA684CD8C
``
GENERATING ANS BASED on 


5 16777216 ABCDE 6 9.7 million
6 16777216 ABCDEF 7 244 Million
7 16777216 ABCDEFG 7 6 Billion
8 16777216 ABCDEFGH 7 152 Billion
9 16777216 ABCDEFGHI 7
10 16777216 ABCDEFGHIK 7 95 Trillion
11 16777216 ABCDEFGHIKL 7
12 16777216 ABCDEFGHIKLM 7
13 16777216ABCDEFGHIKLMN 11
14 16777216FOILJNPVUEDA2

16777216ABCDEFGHIK2



![Without Logo](qr-code.png)

![QR code with a log0](qr-code-logo.png)
