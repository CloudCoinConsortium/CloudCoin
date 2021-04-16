import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import RaidaJS from 'raidajs';
import * as CryptoJS from 'crypto-js';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RaidaServiceService {
  Options: any = {
    // The main domain for the RAIDA
    domain: 'cloudcoin.global',

    // Prefix used to construct the final url
    prefix: 'raida',

    // Protocol schema
    protocol: 'https',

    // Read Timeout for HTTP requests in milliseconds
    timeout: 10000,

    // Default Network Number
    defaultCoinNn: 1,

    // Debug. If set, an additional 'details' field will be set in the response data.
    // The field contains a raw response from the RAIDA server
    debug: true,

    // RAIDA to query when we create SkywWallets
    defaultRaidaForQuery: 7,

    // DDNS service for SkyWallets
    ddnsServer: 'ddns.cloudcoin.global'
  };
  public raidaJS = new RaidaJS(this.Options);

  // public cryptoJS = CryptoJS;
  constructor(private http: HttpClient) {
    this.raidaJS.setTimeout(15000);
    this.raidaJS.setProtocol('https');
    this.raidaJS.setDomain('cloudcoin.global');
    this.raidaJS.setDefaultNetworkNumber(1);
    this.raidaJS.setDefaultRAIDA(7);

    // const params: any = {
    //   sn: '2451121222',
    //   cardnumber: '454645687787864544',
    //   cvv: '487'
    // };
    // console.log(this.newpanGenerate(params));
  }

  set(keys, value): string {
    const key = CryptoJS.enc.Utf8.parse(keys);
    const iv = CryptoJS.enc.Utf8.parse(keys);
    const encrypted = CryptoJS.AES.encrypt(CryptoJS.enc.Utf8.parse(value.toString()), key,
      {
        keySize: 128 / 8,
        iv,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
      });

    /**/

    return encrypted.toString();
  }

  // The get method is use for decrypt the value.
  get(keys, value): string {
    const key = CryptoJS.enc.Utf8.parse(keys);
    const iv = CryptoJS.enc.Utf8.parse(keys);
    const decrypted = CryptoJS.AES.decrypt(value, key, {
      keySize: 128 / 8,
      iv,
      mode: CryptoJS.mode.CBC,
      padding: CryptoJS.pad.Pkcs7
    });

    return decrypted.toString(CryptoJS.enc.Utf8);
  }

  serverRequest(e) {
    return this.http.get(e + '/service/echo');
  }

  register(params): Promise<any> {
    /*
    params = {
      "name" : "mywallet",
      "coin" : {
        "sn" : "4343",
        "an" :  ["f9f2b05d74192e31478846f1b7bdd661","74025cf02053edb09b93ef532a37099d","c3518632d60f897d84ae62e75a7059a3",
        "66dfb17c08b6dbc2846fbe8938bece1a","2f0744735d8b124cc0e31a349770d1f4","cd13fcc1a2806a75322d5a9fda0feaa4",
        "f611a8eb968d4d4b0dd82d8a05b2d8eb","23f8f118f4e76e8cc1488514e6bc6881","d31849f975223a06e765d3433d3e6a9b",
        "4502d00825ccae4c3507cfe1749980d1","62925225e48a9b0fe497dcde66de9227","54688f1c40550d113b8f4f513bf6b8d4",
        "9c2b39d22d0b3e4012eb6e962e99b31b","1564dacd34ace94eb4abfe2f378abe87","1b890b7fa38069745c1b7c7729b242c1",
        "23a0120db1384da7fed62a9100c2f56f","07500e20b49fd14ea5880aa279061aea","72c35043e9a0ea06dc3a29e0409af6ed",
        "415110f4d85b09cf6618aa13164f6b87","8bcf9c8ca170528891bb9eb4ffcbaec0","506c76f5422e92297f4daa453a0d195b",
        "8608a6edb997d0abfec8f88782ff61bd","56d153108902aa4bfe5dab55d9298250","763ec57476e3923eb3f4d9309c5651d6",
        "6938b4aafd39bd136141a2ac31fc8141"]
      }
    }
    */
    return this.raidaJS.apiRegisterSkyWallet(params);
    // .then(response => {
    //   console.log(response);
    // })
  }
  ccRegister(params): Promise<any> {
    return this.raidaJS.apiCreateCCForRegistration(params);
  }

  loginWithCardImage(params): Promise<any> {
    return this.raidaJS.extractStack(params);
  }

  loginWithPassword(params): Promise<any> {
    return this.raidaJS.apiGetCCByUsernameAndPassword(params);
  }

  loginWithCard(params): Promise<any> {
    return this.raidaJS.apiGetCCByCardData(params);
  }

  apiDetect(params): Promise<any> {
    return this.raidaJS.apiDetect(params, raidaNumber => {
      console.log('RAIDA ' + raidaNumber + ' finished detecting');
    });
  }
  fixFracked(params): Promise<any> {
    return this.raidaJS.apiFixfracked(params, (raidaNumber, operation) => {
      console.log('RAIDA ' + raidaNumber + ' finished ' + operation);  // 'operation' is either 'multi_fix' or 'multi_get_ticket'
    });
  }
  deposit(params): void {
    /*
    {
      coinSN0 : {             // Coin serial number
        // Coin Info
        nn: Number,     // Coin network number
        sn: Number,     // Coin serial number

        pownstring: String // PownString
            },

      coinSN1 : { //  Next Coin
        ...
      }

    }

    */
    this.raidaJS.apiDetect(params).then(response => {
      console.log('Detect finished. Fixing Fracked coins');
      this.raidaJS.apiFixfracked(response.result).then(response => {
        console.log('Fixing finished. Total coins fixed: ' + response.fixedNotes);
      });
    });

  }

  transfer(params): Promise<any> {
    return this.raidaJS.apiTransfer(params);
  }


  payment(params): Promise<any> {
    return this.raidaJS.apiPay(params);
  }

  showBalance(params): Promise<any> {
    return this.raidaJS.apiShowBalance(params);
  }
  showCoins(params): Promise<any>{
    return this.raidaJS.apiShowCoins(params);
  }
  fixCoinsSync(params): Promise<any> {
    return this.raidaJS.apiFixTransferSync(params);
  }

  newpanGenerate(params): any[] {
    const sn = params.sn;
    const cardNumber = params.cardnumber;
    const cvv = params.cvv;
    const part = cardNumber.substring(3, cardNumber.length - 1);
    const ans = [];
    for (let i = 0; i < 25; i++) {
      const seed = '' + i + sn + part + cvv;
      ans[i] = {server: i, serverKey: '' + CryptoJS.MD5(seed)};
    }
    return ans;
  }
  generatePan(): any[] {
    return this.raidaJS._generatePan();
  }
  receive(params): any {
    return this.raidaJS.apiReceive(params);
  }

  getDenomination(serial): number{
    return this.raidaJS.getDenomination(serial);
  }
  depositLocal(params): Promise<any> {
    return this.raidaJS.apiSend(params);
  }
  getCardImage(): Promise<any> {
    const params = {
      // array of coins
      "coins" :[{
        "sn" : "4343",
        "an" :  ["f9f2b05d74192e31478846f1b7bdd661","74025cf02053edb09b93ef532a37099d","c3518632d60f897d84ae62e75a7059a3","66dfb17c08b6dbc2846fbe8938bece1a","2f0744735d8b124cc0e31a349770d1f4","cd13fcc1a2806a75322d5a9fda0feaa4","f611a8eb968d4d4b0dd82d8a05b2d8eb","23f8f118f4e76e8cc1488514e6bc6881","d31849f975223a06e765d3433d3e6a9b","4502d00825ccae4c3507cfe1749980d1","62925225e48a9b0fe497dcde66de9227","54688f1c40550d113b8f4f513bf6b8d4","9c2b39d22d0b3e4012eb6e962e99b31b","1564dacd34ace94eb4abfe2f378abe87","1b890b7fa38069745c1b7c7729b242c1","23a0120db1384da7fed62a9100c2f56f","07500e20b49fd14ea5880aa279061aea","72c35043e9a0ea06dc3a29e0409af6ed","415110f4d85b09cf6618aa13164f6b87","8bcf9c8ca170528891bb9eb4ffcbaec0","506c76f5422e92297f4daa453a0d195b","8608a6edb997d0abfec8f88782ff61bd","56d153108902aa4bfe5dab55d9298250","763ec57476e3923eb3f4d9309c5651d6","6938b4aafd39bd136141a2ac31fc8141"]
      }],

      // PNG URL. Must be the compatible with CORS policy
      // The URL can be specified in Base64 format if you prepend 'data:application/octet-binary;base64,' to it
      // "template" : "https://127.0.0.1/image.png"

      // Example (base64)
      "template" : "data:application/octet-binary;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAACqSURBVDhPYxhwwAilGa7duCkJZZZBaQEovRFEaGmobwDz0AATlCYbILtgPpS5HUrfh9LNUDoDRABd8gDMgwKKXYBsgAIUH4Di81C8A4pVoRgFUOwCFigNAhOhdCeUvggibj15Vw+i1598EAii0QH1YgEGgLEhAmWqgYjOtefYQPSNp+8XgWgg8AMRJ6dlXQDR1HcBLmCeNc0AyswHEUAXJIJoil0w0ICBAQCRzCfuSDyCswAAAABJRU5ErkJggg=="
    };
    return this.raidaJS.embedInImage(params);
  }
}
