import {AfterViewInit, Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {LoginService} from '../../../service/login.service';
import {Router} from '@angular/router';
import {RaidaServiceService} from '../../../raida-service.service';
import {AnimationOptions} from 'ngx-lottie';
import {AnimationItem} from 'lottie-web';
import Swal from "sweetalert2";

@Component({
  selector: 'app-transfer',
  templateUrl: './transfer.component.html',
  styleUrls: ['./transfer.component.scss']
})
export class TransferComponent implements OnInit, AfterViewInit {
  public amount = 0;
  public memo: string = null;
  public to: string = null;
  errorMessage: string = null;
  loadingOptions: AnimationOptions = {
    path: '/assets/animations/cloud_loading.json'
  };
  options: AnimationOptions = {
    path: '/assets/animations/cloud_transfer.json'
  };
  errorOptions: AnimationOptions = {
    path: '/assets/animations/error.json'
  };
  showLoader = false;
  balanceLoader = false;
  showNormal = true;
  showError = false;
  complete = false;
  public balance = 0;
  public balanceString = '0';
  balanceBreakup: string = null;
  opinions = 0;
  public balances: any = {};
  loadingMessage = '';


  constructor(private auth: LoginService, private router: Router, private raida: RaidaServiceService) { }
  @ViewChild('autofocus') autoFocusField: ElementRef;
  ngAfterViewInit(): void {
    if(this.autoFocusField)
      this.autoFocusField.nativeElement.focus();
  }
  ngOnInit(): void {
    if (!this.auth.getLoggedIn())
    {
      this.auth.checkLoginStatus();
      this.router.navigate(['/welcome']);
    }
    this.auth.watch().subscribe((value) => {
      if (!value) {
        this.router.navigate(['/welcome']);
      }
    });
    this.checkBalance();
  }
  checkBalance(): void {
    this.showLoadingBalance(true);
    if (!this.auth.getLoggedIn()) {
      this.auth.checkLoginStatus();
      this.router.navigate(['welcome']);
    } else {
      const token = this.auth.getToken();

      const params = {
        sn: token.sn,
        an: token.an
      };

      this.raida.showBalance(params).then(response => {
        this.balance = 0;
        console.log(response);
        this.opinions = 0;
        let highestOpinion = 0;
        let highestOpinionBalance = 0;
        this.balanceBreakup = '<div style="text-align: left; display: table;">Your SkyWallet is out of sync. Balance returned by RAIDA servers' +
          'are as follows: <ul style="text-align: left!important;">';
        if (response.balances) {
          this.balances = response.balances;
          for (const key in response.balances) {
            // check if the property/key is defined in the object itself, not in parent
            if (response.balances.hasOwnProperty(key)) {
              console.log(key, response.balances[key]);
              this.balanceBreakup += '<li style="text-align: left!important">' + response.balances[key] +
                ' RAIDA(s) - ' + parseInt(key).toLocaleString() + ' cc </li>';
              if (response.balances[key] > highestOpinion)
              {
                highestOpinion = response.balances[key];
                highestOpinionBalance = parseInt(key);
                this.balance = highestOpinionBalance;
                this.balanceString = this.balance.toLocaleString();
              }
            }
            this.opinions++;
          }
          this.balanceBreakup += '</ul> It is advised that you sync your SkyWallet by clicking the sync icon</div>';
          if (this.opinions > 1) {
            Swal.fire({
              titleText: 'Your SkyWallet is out of sync',
              text: 'It is recommended that you sync your SkyWallet, to avoid losing coins during transactions. Do you want to sync your' +
                'SkyWallet now?',
              icon: 'warning',
              showCancelButton: true,
              confirmButtonText: 'Yes',
              cancelButtonText: 'No'
            }).then((result) => {
              if (result.value) {
                  this.syncAccount();
              }
            });
          }
        }

        // for (const [key, value] of response.balances.entries()) {
        //   console.log(key, value);
        // }

        /* for (const coin of Object.values(response.coins)) {
           // console.log('counting' + (coin as any).denomination);
           this.balance = this.balance + (coin as any).denomination;
           }
         this.balanceString = this.balance.toLocaleString();*/
        this.showLoadingBalance(false);

      });
    }

  }
  syncAccount(): void {
    this.showLoading(true);
    let completeCallBack = false;

    if (!this.auth.getLoggedIn()) {
      this.auth.checkLoginStatus();
      this.router.navigate(['welcome']);
    } else {
      const token = this.auth.getToken();
      this.loadingMessage = 'Synchronizing SkyWallet...';

      const params = {
        sn: token.sn,
        an: token.an
      };
      const syncObjs = [];

      this.raida.showCoins(params).then(async response => {
        // console.log(response.coinsPerRaida);
        let syncObj = {};
        let counter = 0;
        let realCounter = 0;
        for (const key in response.coinsPerRaida) {
          // console.log(key, response.coinsPerRaida[key]);
          syncObj[key] = response.coinsPerRaida[key];
          counter++;
          realCounter++;
          if (counter === 300)
          {
            syncObjs.push(syncObj);
            syncObj = {};
            counter = 0;
          }
        }
        // alert(realCounter + 'total coins');
        if (syncObj !== {})
        {
          syncObjs.push(syncObj);
        }
        // alert(syncObjs.length + 'total groups');
        // console.log(syncObjs);
        let fixCounter = 0;
        let returnCounter = 0;
        if (syncObjs.length === 0)
        {
          if(!completeCallBack)
          {
            this.checkBalance();
            completeCallBack = true;
          }
        }
        else {
          this.loadingMessage = 'Synchronizing SkyWallet...Group 0/' + syncObjs.length;
          for (const fixObj of syncObjs) {
            fixCounter++;
            this.loadingMessage = 'Synchronizing SkyWallet...Group ' + (fixCounter + 1) + '/' + syncObjs.length;
            await this.sleep(50);
            this.raida.fixCoinsSync(fixObj).then(fixResponse => {
              // console.log(fixResponse);
              returnCounter++;
              if (returnCounter >= fixCounter) {
                if(!completeCallBack)
                {
                  this.checkBalance();
                  completeCallBack = true;
                }
              }

            }).catch(error => {
              returnCounter++;

              if (returnCounter >= fixCounter) {
                if(!completeCallBack)
                {
                  this.checkBalance();
                  completeCallBack = true;
                }
              }


            });

          }
        }

      });
    }
  }
  sleep(ms): Promise<any> {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  transfer(): void {
    if (isNaN(this.amount) ||  this.amount < 1 || this.amount > this.balance)
    {
      this.showErrorMessage('Please enter a valid amount');
      return;
    }
    if (!this.to || this.to === '' )
    {
      this.showErrorMessage('Please enter a valid recipient SkyWallet account');
      return;
    }

    if (this.to.substr(0, 9).toLowerCase() === 'cloudcoin')
    {
      this.showErrorMessage('Transfer option cannot be used to send money to a Merchant SkyWallet Account, please use the "Payment" ' +
        'option instead from the menu');
      return;
    }
    if (!this.memo || this.memo === '') {
      this.memo = 'Transfer';
    }


    if (!this.auth.getLoggedIn())
    {
      this.auth.checkLoginStatus();
      this.router.navigate(['welcome']);
    }
    else {
      const token = this.auth.getToken();

      const coin = {
        sn: token.sn,
        an: token.an
      };
      const params = {...coin, amount: this.amount, to: this.to, memo : this.memo };
      this.showLoading(true);
      this.raida.transfer(params).then(response => {
        // console.log(response);
        if (response.status === 'error') {
          if ('errorText' in response) {
            if (response.errorText.indexOf('Failed to resolve') !== -1) {
              this.showErrorMessage('Invalid Recipient SkyWallet Address: ' + this.to);
            }
            else
            {
              this.showErrorMessage(response.errorText);
            }
          }
          else {
            this.showErrorMessage('Your login session is not valid, please logout and try again');
          }
          return;
        }

        if (!('result' in response)) {
          this.showErrorMessage('Invalid response received');
          return;
        }

        for (const sn in response.result) {
          const cc = response.result[sn];
          const ccadd = {
            sn : sn,
            nn : cc.nn,
            an : cc.an,
            pown : cc.pownstring
          };

          // console.log(ccadd);
        }

        this.complete = true;
        setTimeout(() => {
          this.router.navigate(['/balance']);
        }, 2000);
      });
    }

    }

  animationCreated(animationItem: AnimationItem): void {
    console.log(animationItem);
  }

  showLoading(state): void {
    if (state)
    {
      this.showNormal = false;
      this.showLoader = true;
      this.showError = false;
      this.balanceLoader = false;
    }
    else
    {
      this.showNormal = true;
      this.showLoader = false;
      this.showError = false;
      this.balanceLoader = false;
      const that = this;
      setTimeout(() => {
          if (this.autoFocusField)
          {
            this.autoFocusField.nativeElement.focus();
          }
        }, 100);
    }

  }

  showLoadingBalance(state): void {
    if (state)
    {
      this.showNormal = false;
      this.showLoader = false;
      this.showError = false;
      this.balanceLoader = true;
    }
    else
    {
      this.showNormal = true;
      this.showLoader = false;
      this.showError = false;
      this.balanceLoader = false;
      setTimeout(() => {
        if (this.autoFocusField)
        {
          this.autoFocusField.nativeElement.focus();
        }
      }, 100);
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

}
