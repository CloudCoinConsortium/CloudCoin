# Stack files

Stack files have a ".stack" extention. They represent a stack of coins. They may hold many coins. 

## Naming Convention

Stack files names start with the amount of coins in the file. Then the name of the coins in the file. Then the network number. and then the tag. The tag may be an unguessable guid if the stack is to be placed on a web server. Or it may have some interesting information. 

Examples of stack file names. 
```
12750.cloudcoin.mytag.stack

800.cloudcoin.d735e4d16ff54eda8019e1a59ad33032.stack

5000.cloudcoin.ForBillyJenkins.stack

```
If the coin is used as an ID coin the name will be the DNS named followed by the .stack
```
Sean.CloudCoin.Global.stack

Billy.Skywallet.cc.stack


```
If there is only one coin in the stack file we have the denomination, CloudCoin, network number, serial number, tag and .stack
```
100.CloudCoin.1.7012421..stack

```


IMPORTANT: When creating an application that imports CloudCoin Stack files, you should ignor the parameters 'ed', 'aoid' and 'pown'. These parameters are for use inside the program but may not be there or maybe wrong. Theses parameters should not be imported. If you try to import them, you will fail with some coins. 

NOTE: It is ok to Capitalize in the stack file. 

Many CloudCoins can be stored in one text file as JSON (Java Script Object Notation).

## Standard Stack File Format
```
{
	"cloudcoin": [{
		"nn": "1",
		"sn": "16777215",
		"an": [
			"be0f94e584584d85ab86c301b8f1906d", "1e539633bc974bf88ebaeb7fd23279d7", "d735e4d16ff54eda8019e1a59ad33032", "b394553ff4e94f78a5f4c04df4102a20", "85b8f6956deb4dc28bc40b192ef5b0bc",
			"6775856024a14d10af30181477f8957a", "31274a1103c7406fb0f60736604f23b6", "6eb2c9bfb21e484b89f35c8124a8f3c0", "407b9d6b0b034b23b1d160d7da96a8ae", "2dbf38c2b1834287aed0fdcf4ba32274",
			"4da74b2a8c784731826716dd36296271", "0fe3fb2ded9648eeae64bdcf29abdf15", "01bdfe3b43b74b618c5db44e6b2febb6", "e0bdf8a1b2eb4174bcca6fda0c3609a1", "c8636145d53d42008c84432872816d4c",
			"f5258cb2cd2f4afca333bc7cd64f01eb", "c89eb2da6897489db2581d0b17d4d959", "89d5dba265c84a488106840a79316940", "8dba35338678464a9ef9a1686cd500af", "0a39da32d43546009147127443793297",
			"109a745bb88946a3ac9720bc8f9fee66", "a56933d0ab01400d93782c4eba0c6849", "ef6025e09f88499fb0490edf7ae21eba", "6014a91a938f449cb8f72658e0b9d61e", "cb64b024fe8c47d19d3fcc5c797fa0b3"
		],
		"pown": "ppeppppppfppppppnpppupppp",
		"ed": "9-2021",
		"aoid": []
	}, {
		"nn": "1",
		"sn": "24589",
		"an": [
			"aeab4ff7(113c0871)252a56d249e19bb4", "34847b78a16680f45b0188a1378359e1", "9bf4f0fd441c3fc826dd07f2907bbd7f", "9e4a1a3f371955db(0ecee5c4)54c6327d", "68d49a2e0ae642fd34aea4c169d1d20e",
			"ebb14c3c0c8e898262b11056686fce31", "fe948b05(1e58ab9f)98481e5cb27d7228", "9a9bda52a965ffc2(02c7406c)50e30f05", "ea9c3f42(1d203678)44ea11fee166da00", "9d263c13a3ce498f(3d9221ff)c05f7c4d",
			"da880eb9(07212c18)73592ebe7aa38907", "8c7164cea9456fcf(085f8100)e59089a2", "071dcc9c755e2acc3b5493ac8dc4a57a", "46163edf(0adeee1a)9990fd4fef4ae443", "1a4d1834(237ffe38)(484c881d)baac3ef0",
			"91760a5e0ce5fed94a23017ed8f77e40", "1ddf3387bd32f7ab(29e4fef7)89faeb5f", "91cb6bfda9b0ce2472d0e56a893b0ee2", "af4da2752ee288b5ad263200aa2509fa", "e94bbf83(36fbfd39)(3c014fb2)61c8ba5a",
			"47377bb013a3d87c(42ac04b2)ff41d623", "3da8ce859b5c6927(4bc34ae5)732a1a17", "a09ab60322d29420(040b8d5c)e26b3ee6", "1ef01cb084a8e500(08d61e24)4a991f7d", "d42f923269cf5d93(12711bae)41d1de64"
		],
		"pown": "uuuuuuuuuuuuuuuuuuuuuuuuu",
		"ed": "9-2021",
		"aoid": []
	}]
}
```
IMPORTANT NOTE: You must treat this format with precision. You must have the white space exactly the same. You must have everything in the same order. You must use the same quotation marks, even around things you may think are numbers. 

With white spaces:
NOTE: This format below allows the CloudCoin stack files to be 999 bytes with a single coin and 1000 bytes with the coma that seperates each note. This allows us to easily know how many notes are in a file by now many kilibytes it has. The pargraph symbos are CRLF and the arrors are tabs. 
```
{"cloudcoin":[¶
{"nn":"1",¶
"sn":"16777215","an":[¶
"be0f94e584584d85ab86c301b8f1906d","1e539633bc974bf88ebaeb7fd23279d7","d735e4d16ff54eda8019e1a59ad33032","b394553ff4e94f78a5f4c04df4102a20","85b8f6956deb4dc28bc40b192ef5b0bc",→"6775856024a14d10af30181477f8957a","31274a1103c7406fb0f60736604f23b6","6eb2c9bfb21e484b89f35c8124a8f3c0","407b9d6b0b034b23b1d160d7da96a8ae","2dbf38c2b1834287aed0fdcf4ba32274",→"4da74b2a8c784731826716dd36296271","0fe3fb2ded9648eeae64bdcf29abdf15","01bdfe3b43b74b618c5db44e6b2febb6","e0bdf8a1b2eb4174bcca6fda0c3609a1","c8636145d53d42008c84432872816d4c",→"f5258cb2cd2f4afca333bc7cd64f01eb","c89eb2da6897489db2581d0b17d4d959","89d5dba265c84a488106840a79316940","8dba35338678464a9ef9a1686cd500af","0a39da32d43546009147127443793297",→"109a745bb88946a3ac9720bc8f9fee66","a56933d0ab01400d93782c4eba0c6849","ef6025e09f88499fb0490edf7ae21eba","6014a91a938f449cb8f72658e0b9d61e","cb64b024fe8c47d19d3fcc5c797fa0b3"],¶
"pown":"ppeppppppfppppppnpppupppp",¶
"ed":"9-2021","aoid":[]}¶
]}


```
## Standard CloudCoin Stack File Format including White Space
```
{¶
→"cloudcoin":∙[{¶
→→"nn":∙"1",¶
→→"sn":∙"16777215",¶
→→"an":∙[¶
→→→"be0f94e584584d85ab86c301b8f1906d",∙"1e539633bc974bf88ebaeb7fd23279d7",∙"d735e4d16ff54eda8019e1a59ad33032",∙"b394553ff4e94f78a5f4c04df4102a20",∙"85b8f6956deb4dc28bc40b192ef5b0bc",¶
→→→"6775856024a14d10af30181477f8957a",∙"31274a1103c7406fb0f60736604f23b6",∙"6eb2c9bfb21e484b89f35c8124a8f3c0",∙"407b9d6b0b034b23b1d160d7da96a8ae",∙"2dbf38c2b1834287aed0fdcf4ba32274",¶
→→→"4da74b2a8c784731826716dd36296271",∙"0fe3fb2ded9648eeae64bdcf29abdf15",∙"01bdfe3b43b74b618c5db44e6b2febb6",∙"e0bdf8a1b2eb4174bcca6fda0c3609a1",∙"c8636145d53d42008c84432872816d4c",¶
→→→"f5258cb2cd2f4afca333bc7cd64f01eb",∙"c89eb2da6897489db2581d0b17d4d959",∙"89d5dba265c84a488106840a79316940",∙"8dba35338678464a9ef9a1686cd500af",∙"0a39da32d43546009147127443793297",¶
→→→"109a745bb88946a3ac9720bc8f9fee66",∙"a56933d0ab01400d93782c4eba0c6849",∙"ef6025e09f88499fb0490edf7ae21eba",∙"6014a91a938f449cb8f72658e0b9d61e",∙"cb64b024fe8c47d19d3fcc5c797fa0b3"¶
→→],¶
→→"pown":∙"ppeppppppfppppppnpppupppp",¶
→→"ed":∙"9-2021",¶
→→"aoid":∙[]¶
→},∙{¶
→→"nn":∙"1",¶
→→"sn":∙"24589",¶
→→"an":∙[¶
→→→"aeab4ff7(113c0871)252a56d249e19bb4",∙"34847b78a16680f45b0188a1378359e1",∙"9bf4f0fd441c3fc826dd07f2907bbd7f",∙"9e4a1a3f371955db(0ecee5c4)54c6327d",∙"68d49a2e0ae642fd34aea4c169d1d20e",¶
→→→"ebb14c3c0c8e898262b11056686fce31",∙"fe948b05(1e58ab9f)98481e5cb27d7228",∙"9a9bda52a965ffc2(02c7406c)50e30f05",∙"ea9c3f42(1d203678)44ea11fee166da00",∙"9d263c13a3ce498f(3d9221ff)c05f7c4d",¶
→→→"da880eb9(07212c18)73592ebe7aa38907",∙"8c7164cea9456fcf(085f8100)e59089a2",∙"071dcc9c755e2acc3b5493ac8dc4a57a",∙"46163edf(0adeee1a)9990fd4fef4ae443",∙"1a4d1834(237ffe38)(484c881d)baac3ef0",¶
→→→"91760a5e0ce5fed94a23017ed8f77e40",∙"1ddf3387bd32f7ab(29e4fef7)89faeb5f",∙"91cb6bfda9b0ce2472d0e56a893b0ee2",∙"af4da2752ee288b5ad263200aa2509fa",∙"e94bbf83(36fbfd39)(3c014fb2)61c8ba5a",¶
→→→"47377bb013a3d87c(42ac04b2)ff41d623",∙"3da8ce859b5c6927(4bc34ae5)732a1a17",∙"a09ab60322d29420(040b8d5c)e26b3ee6",∙"1ef01cb084a8e500(08d61e24)4a991f7d",∙"d42f923269cf5d93(12711bae)41d1de64"¶
→→],¶
→→"pown":∙"uuuuuuuuuuuuuuuuuuuuuuuuu",¶
→→"ed":∙"9-2021",¶
→→"aoid":∙[]¶
→}]¶
}¶
```

### Use of JSON. This coin is designed to be parsed without special json software. It must be exactly like this with the same order and whitespace. 
### Spaces: There are spaces, tabs and carriage returns. Keep them exactly the same as above. 
### Order of elements. The order of key value pairs must not be changed. 
### Capitalization. Everything is lowercase. 

### Components

1. cloudcoin is the name of the coin. CloudCoin can have sub coins like celebrium, eossteath.
2. nn is the Network Number or the Network Name. The root is "1". THIS MUST BE TREATED LIKE A STRING AND BE SURROUNDED BY QUOTES. 
3. sn is the Serial Number but it may not be a number. THIS MUST BE TREATED LIKE A STRING AND BE SURROUNDED BY QUOTES. 
4. an means authenticity numbers. These are 32 hexidecimal characters. We put them in rows of five with a carrage return at the end of each row. If the coin is encrypted, some of the parts of the ANs will be surrounded with parenthesis. You can tell if the coin is encrypted by looking for parenthesis in the ANs. 
5. pown shows the status of the last pown. There are 25 letters in a row. The only letters that are use are p, f, e, n and u. (pass, fail, error, no reply and untried or unknown). The default is all 'u's.
6. ed is the expiration date. Five years from the last month it was authenticated. month seperated by year with a hyphen. No zeros before the month. 
7. aoid is the All Other Info Dump. It can hold an array of strings. By default, it is empty. 

