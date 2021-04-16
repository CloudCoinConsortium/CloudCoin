import {Component, OnInit} from '@angular/core';
import {LoginService} from '../../../service/login.service';
import {Router} from '@angular/router';
import {RaidaServiceService} from '../../../raida-service.service';
import {AnimationOptions} from 'ngx-lottie';
import {AnimationItem} from 'lottie-web';

@Component({
  selector: 'app-drop-debitcard',
  templateUrl: './drop-debitcard.component.html',
  styleUrls: ['./drop-debitcard.component.scss']
})
export class DropDebitcardComponent implements OnInit {

  files: File[] = [];
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

  constructor(private auth: LoginService, private router: Router, private  raida: RaidaServiceService) {
  }


  onSelect(event): void {
    // console.log(event);
    this.files.push(...event.addedFiles);
    this.getBase64(event);
  }

  getBase64(event): void {
    const me = this;
    const file = event.addedFiles[0];
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = function() {
      // me.modelvalue = reader.result;
     // console.log(reader.result);
      const params = {
        template: reader.result
      };
      me.showLoading(true);
      me.raida.loginWithCardImage(params).then(response => {
        if (response.status === 'error')
        {
          if(response.errorText.indexOf('PNG signature'))
          {
            me.showErrorMessage('The dropped debit card image is not valid. Please ensure you are using the ' +
              'original unmodified debit card image, which can be modified if shared via messaging/file sharing apps.');
          }
          else
          {
            me.showErrorMessage(response.errorText);
          }
        }
        else {
          if (response.status === 'done' && response.cloudcoin !== null && response.cloudcoin.length > 0)
          {
            const cloudcoin = response.cloudcoin[0];

            const coinParams = {
              sn: cloudcoin.sn,
              an: cloudcoin.an
            };
            me.raida.showBalance(coinParams).then(coinResponse => {
              let loggedIn = false;
              console.log(coinResponse);
              if (coinResponse.raidaStatuses.indexOf('p') === -1) {
                loggedIn = false;
              } else {
                loggedIn = true;
              }
              if (loggedIn === true) {
                localStorage.setItem('cc', JSON.stringify(cloudcoin));
                // console.log('inside');
                me.auth.checkLoginStatus();
                setTimeout(() => {
                  me.router.navigate(['/balance']);
                }, 2500);
              } else {
                me.showErrorMessage('Invalid SkyWallet name/password');
              }
            });

          }
          // console.log(response);
        }
      });

    };
    reader.onerror = function(error) {
      console.log('Error: ', error);
      me.showErrorMessage('Error reading Card Image');
    };
  }

  onRemove(event): void {
    console.log(event);
    this.files.splice(this.files.indexOf(event), 1);
  }


  ngOnInit(): void {
    if (this.auth.getLoggedIn()) {
      this.router.navigate(['/balance']);
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
