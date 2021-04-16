import {AfterViewInit, Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {LoginService} from '../../../service/login.service';
import {Router} from '@angular/router';
import {RaidaServiceService} from '../../../raida-service.service';
import {AnimationOptions} from 'ngx-lottie';
import {AnimationItem} from 'lottie-web';

@Component({
  selector: 'app-withdraw',
  templateUrl: './withdraw.component.html',
  styleUrls: ['./withdraw.component.scss']
})
export class WithdrawComponent implements OnInit, AfterViewInit {

  public amount = 0;
  errorMessage: string = null;
  options: AnimationOptions = {
    path: '/assets/animations/cloud_download.json'
  };
  errorOptions: AnimationOptions = {
    path: '/assets/animations/error.json'
  };
  showLoader = false;
  showNormal = true;
  showError = false;


  constructor(private auth: LoginService, private router: Router, private  raida: RaidaServiceService) { }
  @ViewChild('autofocus') autoFocusField: ElementRef;
  ngAfterViewInit(): void {
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
  }
  withdraw(): void {
    if (isNaN(this.amount) ||  this.amount < 1)
    {
      this.showErrorMessage('Please enter a valid amount');
      return;
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
      const params = { ...coin, amount : this.amount };
      this.showLoading(true);
      this.raida.receive(params).then(response => {
       // console.log(response);
        if (response.status === 'error') {
          if ('errorText' in response) {
            this.showErrorMessage(response.errorText);
          }
          else {
            this.showErrorMessage('Your login session is not valid. Please logout and try again');
          }
          return;
        }

        if (!('result' in response)) {
          this.showErrorMessage('Invalid response received');
          return;
        }

        const d = new Date();
        let month = d.getMonth();
        let year = d.getFullYear();
        month += 1;
        year += 1905;

        const dateStr = month + '-' + year;


        const stack = {cloudcoin : [] };
        for (const sn in response.result) {
          const cc = response.result[sn];

          for (let n = 0; n < cc.an.length; n++) {
            if (cc.an[n] === undefined) {
              cc.an[n] = this.raida.generatePan();
            }
          }
          cc.an = cc.an.map(v => v.toLowerCase());
          const ccadd = {
            sn,
            nn : cc.nn,
            an : cc.an,
            pown : cc.pownstring,
            aoid : [],
            ed : dateStr
          };

          // console.log(ccadd);
          stack.cloudcoin.push(ccadd);

        }
        // let cstr = JSON.stringify(stack, null, 2)
        const cstr = this._stringifyJSONCoin(stack);
        const filename = this.amount + '.CloudCoin.' + Date.now() + '.stack';
        this.download(filename, cstr);
        this.showLoading(false);
        this.router.navigate(['/balance']);
      });


    }


  }

   download(filename, text): void {
    const pom = document.createElement('a');
    pom.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    pom.setAttribute('download', filename);

    if (document.createEvent) {
      const event = document.createEvent('MouseEvents');
      event.initEvent('click', true, true);
      pom.dispatchEvent(event);
    } else {
      pom.click();
    }
  }

  _stringifyJSONCoin(stack): string {
    const coins = stack.cloudcoin;
    let first = true;
    let json = '{\n\t"cloudcoin": [{\n';

    for (const cc of coins) {
      if (!first) {
        json += ', {\n';
      }
      json += '\t\t"nn": "' + cc.nn + '",\n';
      json += '\t\t"sn": "' + cc.sn + '",\n';
      json += '\t\t"an": [\n\t\t\t"';
      let i = 0;
      for (const an of cc.an) {
        json +=  an;
        if (i !== 24) {
          json += '", ';
          if (((i + 1) % 5) === 0) {
            json += '\n\t\t\t';
          }
          json += '"';
        }

        i++;

      }
      json += '"\n\t\t],\n';
      if ('pown' in cc) {
        json += '\t\t"pown": "' + cc.pown + '"';
      }
      if ('ed' in cc) {
        json += ',\n\t\t"ed": "' + cc.ed + '"';
      }

      if ('aoid' in cc) {
        if (cc.aoid.length === 0)
        {
          json += ',\n\t\t"aoid": []\n';
        }
        else {
          json += ',\n\t\t"aoid": "' + cc.aoid + '"\n';
        }
      }
      else {
        json += ',\n\t\t"aoid": []\n';
      }

      json += '\t}';
      first = false;
    }
    json += ']\n}';

    return json;
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
