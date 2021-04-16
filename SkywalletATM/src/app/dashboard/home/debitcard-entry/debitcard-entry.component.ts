import { Component, OnInit, ViewChild, ElementRef, AfterViewInit } from '@angular/core';
import {RaidaServiceService} from '../../../raida-service.service';
import {LoginService} from '../../../service/login.service';
import {Router} from '@angular/router';
import { AnimationItem } from 'lottie-web';
import { AnimationOptions } from 'ngx-lottie';


@Component({
  selector: 'app-debitcard-entry',
  templateUrl: './debitcard-entry.component.html',
  styleUrls: ['./debitcard-entry.component.scss']
})
export class DebitcardEntryComponent implements OnInit, AfterViewInit {

  walletName: string = null;
  cardNumber: string = null;
  expiryDate: string = null;
  cvv2: number = null;
  errorMessage: string = null;
  options: AnimationOptions = {
    path: '/assets/animations/cloud_login.json'
  };
  errorOptions: AnimationOptions = {
    path: '/assets/animations/error.json'
  };
  showLoader = false;
  showNormal = true;
  showError = false;



  constructor(private raida: RaidaServiceService, private auth: LoginService, private  router: Router) { }
  @ViewChild('autofocus') autoFocusField: ElementRef;
  ngOnInit(): void {

    if (this.auth.getLoggedIn())
    {
      this.router.navigate(['/balance']);
    }
  }
  ngAfterViewInit(): void {
    this.autoFocusField.nativeElement.focus();
  }

  login(): void {
    if (!this.walletName || this.walletName === '')
    {
      this.showErrorMessage('Invalid wallet name');
      return;
    }
    if (!this.cardNumber || this.cardNumber === '')
    {
      this.showErrorMessage('Invalid card number');
      return;
    }
    else
    {
      if (this.cardNumber.replace(' ', '').length !== 16)
      {
        this.showErrorMessage('Invalid card number');
        return;
      }
      else {
        this.cardNumber = this.cardNumber.replace(' ', '');
      }

      if (!this.cvv2 || this.cvv2.toString().length < 4)
      {
        this.showErrorMessage('Invalid PIN');
        return;
      }
      // Everything seems okay now so we can call login
     // this.walletName = this.walletName + '.skywallet.cc';
      const params = {
        // Username
        username : this.walletName,
        // Cardnumber
        cardnumber : this.cardNumber.toString(),

        // CVV
        cvv: this.cvv2.toString()
      };
      this.showLoading(true);
      this.raida.loginWithCard(params).then(response => {
        if (response.status === 'error')
        {
          if (response.errorText.indexOf('Failed to resolve DNS name') !== -1) {
            this.showErrorMessage('Invalid Debit Card Details');
          }
          else
          {
            this.showErrorMessage(response.errorText);
          }
        }
        else {
          if (response.status === 'done')
          {
            const cloudcoin = response.cc;
            // console.log(cloudcoin);
            if (!cloudcoin.an && cloudcoin.ans) {
              cloudcoin.an = cloudcoin.ans;
            }
            const coinParams = {
              sn: cloudcoin.sn,
              an: cloudcoin.an
            };
            this.raida.showBalance(coinParams).then(coinResponse => {
              let loggedIn = false;
              console.log(coinResponse);
              if (coinResponse.raidaStatuses.indexOf('p') === -1) {
                loggedIn = false;
              } else {
                loggedIn = true;
              }
              if (loggedIn === true) {
                localStorage.setItem('cc', JSON.stringify(cloudcoin));
                localStorage.setItem('skywallet', this.walletName);
                // console.log('inside');
                this.auth.checkLoginStatus();
                setTimeout(() => {
                  this.router.navigate(['/balance']);
                }, 2500);
              } else {
                this.showErrorMessage('Invalid SkyWallet name/password');
              }
            });
          }
        //  console.log(response);
        }
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
    }
    else
    {
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
