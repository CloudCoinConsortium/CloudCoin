import {Component, OnInit} from '@angular/core';
import {RaidaServiceService} from '../../../raida-service.service';
import {LoginService} from '../../../service/login.service';
import {Router} from '@angular/router';
import {AnimationOptions} from 'ngx-lottie';
import {AnimationItem} from 'lottie-web';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-balance',
  templateUrl: './balance.component.html',
  styleUrls: ['./balance.component.scss']
})
export class BalanceComponent implements OnInit {
  public balance = 0;
  public balanceString = '0';
  public balances: any = {};
  errorMessage: string = null;
  options: AnimationOptions = {
    path: '/assets/animations/cloud_loading.json'
  };
  errorOptions: AnimationOptions = {
    path: '/assets/animations/error.json'
  };
  showLoader = false;
  showNormal = true;
  showError = false;
  skywallet: string = null;
  balanceBreakup: string = null;
  opinions = 0;
  loadingMessage = '';

  constructor(private raida: RaidaServiceService, private auth: LoginService, private router: Router) {
  }

  ngOnInit(): void {
    this.checkBalance();
    this.auth.watch().subscribe((value) => {
      if (!value) {
        this.router.navigate(['welcome']);
      }
    });

    if (localStorage.getItem('skywallet')) {
      this.skywallet = localStorage.getItem('skywallet');
    }

  }

  checkBalance(): void {
    this.showLoading(true);
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
              // console.log(key, response.balances[key]);
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
                {
                 this.syncAccount();
                }
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
        this.showLoading(false);

      });
    }

  }

  animationCreated(animationItem: AnimationItem): void {
    console.log(animationItem);
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
  showLoading(state): void {
    this.loadingMessage = '';
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


}
