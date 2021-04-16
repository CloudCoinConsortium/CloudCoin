import {AfterViewInit, ChangeDetectorRef, Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {RaidaServiceService} from 'src/app/raida-service.service';
import * as CryptoJS from 'crypto-js';
import {AnimationOptions} from 'ngx-lottie';
import {AnimationItem} from 'lottie-web';
import RaidaJS from 'raidajs';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit, AfterViewInit {
  selectedFile: File;
  selectedData: any = null;
  walletName: any;
  walletEmail: any;
  walletPin: string;
  agreed1 = false;
  agreed2 = false;
  registerTapped = false;
  complete = false;
  errorMessage: string = null;
  registerParams: any = null;
  completeMessage: string = null;
  cardImage: string = null;
  cardImageBlank: string = null;
  progressMessage: string = null;
  options: AnimationOptions = {
    path: '/assets/animations/cloud_login.json'
  };
  errorOptions: AnimationOptions = {
    path: '/assets/animations/error.json'
  };
  showLoader = false;
  showNormal = true;
  showError = false;
  buttonText = 'Select CloudCoin';
  public raidaJS = new RaidaJS();

  constructor(private changeDetectorRef: ChangeDetectorRef, private raida: RaidaServiceService) {
  }

  @ViewChild('autofocus') autoFocusField: ElementRef;
  ngAfterViewInit(): void {
    this.autoFocusField.nativeElement.focus();
  }

  ngOnInit(): void {
    // console.log(environment.card_image);
    /* const params = {
      // array of coins
      coins: [{
        sn: 4343,
        nn: 1,
        an: ['f9f2b05d74192e31478846f1b7bdd661', '74025cf02053edb09b93ef532a37099d', 'c3518632d60f897d84ae62e75a7059a3', '66dfb17c08b6dbc2846fbe8938bece1a', '2f0744735d8b124cc0e31a349770d1f4', 'cd13fcc1a2806a75322d5a9fda0feaa4', 'f611a8eb968d4d4b0dd82d8a05b2d8eb', '23f8f118f4e76e8cc1488514e6bc6881', 'd31849f975223a06e765d3433d3e6a9b', '4502d00825ccae4c3507cfe1749980d1', '62925225e48a9b0fe497dcde66de9227', '54688f1c40550d113b8f4f513bf6b8d4', '9c2b39d22d0b3e4012eb6e962e99b31b', '1564dacd34ace94eb4abfe2f378abe87', '1b890b7fa38069745c1b7c7729b242c1', '23a0120db1384da7fed62a9100c2f56f', '07500e20b49fd14ea5880aa279061aea', '72c35043e9a0ea06dc3a29e0409af6ed', '415110f4d85b09cf6618aa13164f6b87', '8bcf9c8ca170528891bb9eb4ffcbaec0', '506c76f5422e92297f4daa453a0d195b', '8608a6edb997d0abfec8f88782ff61bd', '56d153108902aa4bfe5dab55d9298250', '763ec57476e3923eb3f4d9309c5651d6', '6938b4aafd39bd136141a2ac31fc8141']
      }],
      template: 'assets/card.png',
    };
    try {
      this.raidaJS.embedInImage(params).then(coinResponse => {
        // console.log(coinResponse);
        // alert(coinResponse);
        this.showCompleteMessage('Registration Successful!');
        const cardData = {
          name: 'Partha Dasgupta',
          number: '1234 5678 9012 3456',
          expiry: '04-2026',
          pin: '519'
        }; // 'Partha Dasgupta', 50, 50, '#FFFFFF', '18px courier'
        this.addTextToImage('data:image/png;base64,' + coinResponse,cardData);
        // this.cardImage =
      });
    } catch (e) {
      console.log(e);
    }
    */
  }

  onFileChanged(event): void {
    this.buttonText = 'Select CloudCoin';
    this.selectedFile = event.target.files[0];
    const fileReader = new FileReader();
    const me = this;
    fileReader.readAsText(this.selectedFile, 'UTF-8');
    fileReader.onload = () => {
      this.selectedData = fileReader.result;
      // console.log(JSON.parse(this.selectedData));
      try {
        const coin = JSON.parse(this.selectedData);
        if (coin.cloudcoin) {
          const stack = coin.cloudcoin;
          if (stack.length > 1) {
            this.showErrorMessage('Stack for registration must contain only one CloudCoin');
            this.selectedData = {};
            return;
          }
          for (let i = 0; i < stack.length; i++) {
            const cc = stack[i];
            if (!cc.sn) {
              this.showErrorMessage('Please choose a valid CloudCoin');
              this.selectedData = {};
              return;
            }
            const total = me.raida.getDenomination(cc.sn);
            if (total < 1) {
              this.showErrorMessage('Please choose a valid CloudCoin');
              this.selectedData = {};
              return;
            } else {
              this.buttonText = '1 CloudCoin Selected';
            }
          }
        } else {
          this.showErrorMessage('Please choose a valid CloudCoin');
          this.selectedData = {};
          return;
        }

      } catch (e) {
        if (e.indexOf('Failed to resolve DNS name') === -1) {
          console.log(e);
          me.showErrorMessage('Failed to parse CloudCoin');
          return;
        }
      }

    };
    fileReader.onerror = (error) => {
      console.log(error);
      this.showErrorMessage('Cannot Parse CloudCoin');
    };
  }

  validateEmail(email): boolean {
    const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
  }

  register(): void {
    if (this.registerTapped) {
      return;
    }
    else {
      this.registerTapped = true;
    }

    // alert(this.agreed1 + 'and ' + this.agreed2);
    if (!this.agreed1 || !this.agreed2)
    {
      this.showErrorMessage('Please accept to the terms and conditions of registration to proceed');
      this.registerTapped = false;
      return;
    }

    if (!this.walletName || this.walletName === '') {
        this.showErrorMessage('Please choose a valid wallet name');
        this.registerTapped = false;
        return;

    }

    if (!this.walletEmail || this.walletEmail === '' || !this.validateEmail(this.walletEmail)) {
      this.showErrorMessage('Please type a valid email address');
      this.registerTapped = false;
      return;
    }


    if (!this.walletPin || this.walletPin.length < 8) {
      this.showErrorMessage('Please enter a valid Password (minimum 8 characters)');
      this.registerTapped = false;
      return;
    }
    if (!this.selectedData || this.selectedData === {}) {
      this.showErrorMessage('Please select the CloudCoin to be used for registration');
      this.registerTapped = false;
      return;
    }
    if (!(this.walletName.endsWith('.skywallet.cc'))) {
      this.showErrorMessage('Only wallet names ending with .skywallet.cc  are allowed, e.g. john.skywallet.cc');
      this.registerTapped = false;
      return;
    }

    const alias = this.walletName.replace(/\.skywallet\.cc$/, '');
    if (!alias.match(/[a-z0-9]{1,62}/)) {
      this.showErrorMessage('Only alphanumeric characters are allowed. Max length 62');
      this.registerTapped = false;
      return;
    }
    this.raidaJS._resolveDNS(this.walletName).then(data => {
        if (data != null) {
          this.showErrorMessage('SkyWallet already exists');
          this.registerTapped = false;
          return;
        } else {

          // console.log('name', this.walletName);
          if (this.selectedData != null) {
            // console.log("selectedData",);
            let snString = '';
            let anArray = [];
            let nnString = 1;
            //  alert('now going inside');
           // console.log(this.selectedData);
            try {
              JSON.parse(this.selectedData).cloudcoin.forEach((element, index) => {
                if (index === 0) {
                  snString = element.sn;
                  anArray = element.an;
                  nnString = element.nn;
                }
              });
              if (snString !== '' && anArray !== []) {
                // console.log('generating card');
                // generate a new card number
                const params = {
                  // Username
                  sn : snString,
                  // Password
                  password : this.walletPin,
                  // Recovery Email
                  email: this.walletEmail
                };
                console.log(params);
                const response = this.raida.ccRegister(params).then(ccCardData => {
                  console.log(ccCardData);
                  if ((ccCardData as any).status === 'done')
                  {
                    const cardGenerated = this.makeCard((ccCardData as any).rand);
                    const cardNumber = cardGenerated.number;
                    const cvv = (ccCardData as any).cvv;
                    const pans = (ccCardData as any).pans;
                    // pown the card with the new ans
                    this.showLoading(true);
                    this.progressMessage = 'Detecting Coin\'s authenticity..';
                    const pownParams = [{sn: parseInt(snString), an: anArray, pan: pans}];
                    console.log('pown params', pownParams);
                    const result = this.raida.apiDetect(pownParams).then(detectResponse => {
                      // console.log('Detect finished. Fixing Fracked coins');
                       console.log(detectResponse);
                      const validCoins = detectResponse.authenticNotes + detectResponse.frackedNotes;
                      if (validCoins === 0) {
                        this.showErrorMessage('Selected CloudCoin is not valid');
                        this.registerTapped = false;
                        return;
                      }
                       this.progressMessage = 'Fixing fracked coin..';
                       this.raida.fixFracked(detectResponse.result).then(frackedResponse => {
                        this.progressMessage = 'Generating debit card..';
                        // console.log(frackedResponse);
                        // console.log('Fixing finished. Total coins fixed: ' + frackedResponse.fixedNotes);
                        /*if (frackedResponse.fixedNotes < 1) {
                          this.showErrorMessage('Selected CloudCoin is not valid');
                          return;
                        }*/
                        const ip = '1.' + ((parseInt(snString) >> 16) & 0xff) + '.' + ((parseInt(snString) >> 8) & 0xff) + '.'
                          + ((parseInt(snString)) & 0xff);
                        const cardData = {
                          name: this.walletName,
                          number: cardNumber,
                          expiry: cardGenerated.expiry,
                          pin: cvv,
                          ip
                        }; // 'Partha Dasgupta', 50, 50, '#FFFFFF', '18px courier'
                        this.registerParams = {
                          name: this.walletName,
                          coin: {
                            sn: snString,
                            an: pans,
                            nn: nnString
                          }
                        };
                        // console.log(this.registerParams);
                        this.addTextToImage('assets/card.png', cardData, this.registerParams, this.doRegister);
                      });
                    });
                  }
                  else
                  {
                    this.showErrorMessage('Error generating card from info');
                    return;
                  }
                });
              } else {
                this.showErrorMessage('Selected CloudCoin is not valid');
                this.registerTapped = false;
                return;
              }
            } catch (e) {
              console.log(e);
              if (e.indexOf('Failed to resolve DNS name') === -1) {
                this.showErrorMessage(e.toString());
                this.registerTapped = false;
                return;
              }
            }
          } else {
            this.showErrorMessage('Please select the CloudCoin to be used for registration');
            this.registerTapped = false;
          }
        }
      }
    );
  }

  doRegister(me, registerParams): void {
   // console.log('over here');
     console.log(registerParams);
    me.progressMessage = 'Creating SkyWallet Account...';
    me.raida.register(registerParams).then(registerResponse => {
      // console.log(registerResponse);
      if (registerResponse.status === 'done') {
        // if correct coin
        me.progressMessage = 'SkyWallet account created..';
        const coinParam = {
          // array of coins
          coins: [{
            sn: parseInt(registerParams.coin.sn),
            nn: 1,
            an: registerParams.coin.pan
          }],
          // PNG URL. Must be the compatible with CORS policy
          // The URL can be specified in Base64 format if you prepend 'data:application/octet-binary;base64,' to it
          template: me.cardImageBlank
        };
        // console.log(coinParam);
        // alert('now embedding');
        try {
          me.raidaJS.embedInImage(coinParam).then(coinResponse => {
            // console.log(coinResponse);
             if (coinResponse.status && coinResponse.status === 'error')
             {
               me.showErrorMessage(coinResponse.errorText);
               this.registerTapped = false;
             }
             else
             {
               // alert(coinResponse);
               me.showCompleteMessage('Registration Successful!');
               me.cardImage = 'data:image/png;base64,' + coinResponse;
             }

            // this.cardImage =
          });
        } catch (e) {
          console.log(e);
        }
      }
    });
  }

  makeCard(rand): any {
    const precardNumber = '401' + rand;
    const reverse = precardNumber.split('').reverse().join('');
    let total = 0;
    for (let i = 0; i < reverse.length; i++) {
      let num = parseInt(reverse.charAt(i));
      if ((i + 3) % 2) {
        num *= 2;
        if (num > 9) {
          num -= 9;
        }
      }
      total += num;
    }

    let remainder = 10 - (total % 10);
    if (remainder === 10) {
      remainder = 0;
    }

    const cardNumber = precardNumber + remainder;

    const fiveYearsFromNow = new Date();
    fiveYearsFromNow.setFullYear(fiveYearsFromNow.getFullYear() + 5);

    const month = fiveYearsFromNow.getMonth() + 1;
    const year = fiveYearsFromNow.getFullYear().toString().substr(-2);
    return {number: cardNumber, expiry: month + '-' + year};
  }

  downloadImage(): void {
    const a = document.createElement('a'); // Create <a>
    a.href = this.cardImage; // Image Base64 Goes here
    a.download = this.walletName + '.png'; // File name Here
    a.click();

  }

  getPansFromPasswordV2(sn, email, password): any {

    const params = {
      // Username
      sn : sn,
      // Password
      password : password,
      // Recovery Email
      email: email
    };
    let grv = null;
    console.log(params);
    const response = this.raida.ccRegister(params).then(data => {
        console.log(data);
        if ((data as any).status === 'done')
        {
          grv = {pas: (data as any).pans, rand: (data as any).rand, pin: (data as any).cvv};
        }
        else
        {
          this.showErrorMessage('Error generating card from info');
          return;
        }
    });
    return grv;
  }


  getPansFromPassword(sn, email, password): any {
    let finalStr = '';
    for (let i = 0; i < password.length; i++) {
      const code = password.charCodeAt(i);
      finalStr += '' + code;
    }

    // Generating rand and pin from the password
    const rand = finalStr.slice(0, 12);
    const pin = finalStr.slice(12, 16);
    const pans = [];
    for (let i = 0; i < 25; i++) {
      const seed = '' + i + sn + rand + pin;
      const p = '' + CryptoJS.MD5(seed);

      const p0 = p.substring(0, 24);
      let component = '' + sn + '' + i + email;
      component = '' + CryptoJS.MD5(component);
      const p1 = component.substring(0, 8);
      pans[i] = p0 + p1;
    }

    const grv = {pans, rand, pin};

    return grv;
  }


  generate(n): string {
    const add = 1;
    let max = 12 - add;   // 12 is the min safe number Math.random() can generate without it starting to pad the end with zeros.

    if (n > max) {
      return this.generate(max) + this.generate(n - max);
    }

    max = Math.pow(10, n + add);
    const min = max / 10; // Math.pow(10, n) basically
    const cardNumber = Math.floor(Math.random() * (max - min + 1)) + min;

    return ('' + cardNumber).substring(add);
  }

  animationCreated(animationItem: AnimationItem): void {
    // console.log(animationItem);
  }

  showLoading(state): void {
    if (state) {
      this.showNormal = false;
      this.showLoader = true;
      this.showError = false;
    } else {
      this.showNormal = true;
      this.showLoader = false;
      this.showError = false;
    }

  }

  showErrorMessage(message): void {
    this.errorMessage = message;
    this.showNormal = false;
    this.showError = true;
    this.showLoader = false;


  }

  hideErrorMessage(): void {
    this.errorMessage = '';
    this.showNormal = true;
    this.showError = false;
    this.showLoader = false;
  }

  showCompleteMessage(message): void {
    this.errorMessage = '';
    this.showNormal = false;
    this.showError = false;
    this.showLoader = false;
    this.completeMessage = message;
    this.complete = true;
  }

  addTextToImage(imagePath, cardData, registerParams, callback): void {
    /*const cardData = {
      name: 'Partha Dasgupta',
      number: '1234 5678 9012 3456',
      expiry: '04-2026',
      pin: '519'
    }; // 'Partha Dasgupta', 50, 50, '#FFFFFF', '18px courier'
     */
    // console.log(registerParams);

    const card_canvas = document.createElement('canvas');
    card_canvas.setAttribute('width', '700');
    card_canvas.setAttribute('height', '906');
    const context = card_canvas.getContext('2d');
    const me = this;
    // Draw Image function
    const img = new Image();
    img.src = imagePath;
    img.onload = function() {
      context.drawImage(img, 0, 0);
      context.lineWidth = 0;
      // context.lineStyle = color;
      context.fillStyle = '#FFFFFF';
      context.font = '30px sans-serif';
      context.fillText(cardData.name, 50, 400);
      context.font = '40px sans-serif';
      context.fillText(cardData.number.replace(/(.{4})/g, '$1 '), 50, 300);
      context.font = '18px sans-serif';
      context.fillText('Keep these numbers secret, do not give to merchants', 50, 325);
      context.font = '18px sans-serif';
      context.fillText('Share this name with others for receiving payments', 50, 425);

      context.font = '25px sans-serif';
      context.fillText(cardData.expiry, 450, 355);
      context.fillStyle = '#000000';
      context.fillText('CVV (keep secret): ' + cardData.pin, 50, 675);
      context.fillStyle = '#ffffff';
      context.fillText( 'IP ' + String(cardData.ip), 150, 740);


      me.cardImageBlank = card_canvas.toDataURL(); // 'data:image/png;base64,' + coinResponse;
      callback(me, registerParams);
    };
  }

}
