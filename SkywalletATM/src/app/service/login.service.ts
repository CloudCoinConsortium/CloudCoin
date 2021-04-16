import {Injectable} from '@angular/core';
import {Observable, Subject} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  private isLoggedIn = false;
  loginStatus: Subject<any> = new Subject<any>();
  private cloudcoin: any = null;

  constructor() {
    this.getLoggedIn();
  }

  checkLoginStatus(): void {
    this.loginStatus.next(this.getLoggedIn());

  }

  watch(): Observable<any> {
    return this.loginStatus.asObservable();
  }

  getToken(): any {
    if (this.getLoggedIn())
    {
      return JSON.parse(localStorage.getItem('cc'));
    }
    else
    {
      this.checkLoginStatus();
    }
  }

  getLoggedIn(): boolean
  {
    if (localStorage.getItem('cc') !== null) {
     // console.log(localStorage.getItem('cc'));
      try {
        const cc = JSON.parse(localStorage.getItem('cc'));
        // console.log('parsed coin');
        // console.log(cc);
        if (cc.sn && cc.sn > 0 && cc.an !== null && cc.an.length === 25)
        {
          this.isLoggedIn = true;
        }
      }
      catch (e)
      {
        // console.log(e);
        // console.log('cannot parse coin');
        this.isLoggedIn = false;
        localStorage.setItem('cc', 'null');
        localStorage.removeItem('cc');
      }
    }
    else {
      this.isLoggedIn = false;
    }
    return this.isLoggedIn;
  }
}
