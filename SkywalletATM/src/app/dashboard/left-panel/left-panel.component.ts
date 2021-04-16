import { EventService } from 'src/app/service/event.service';
import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import RaidaJS from 'raidajs';
import { RaidaServiceService } from 'src/app/raida-service.service';
import { LoginService } from 'src/app/service/login.service';
import { Router } from '@angular/router';
import Swal from 'sweetalert2'
@Component({
  selector: 'app-left-panel',
  templateUrl: './left-panel.component.html',
  styleUrls: ['./left-panel.component.scss']
})
export class LeftPanelComponent implements OnInit {
  public isLoggedIn = false;
  Options: any = {
    // The main domain for the RAIDA
    domain : 'cloudcoin.global',

    // Prefix used to construct the final url
    prefix : 'raida',

    // Protocol schema
    protocol: 'https',

    // Read Timeout for HTTP requests in milliseconds
    timeout: 10000,

    // Default Network Number
    defaultCoinNn: 1,

    // Debug. If set, an additional 'details' field will be set in the response data.
    // The field contains a raw response from the RAIDA server
    debug: false,

    // RAIDA to query when we create SkywWallets
    defaultRaidaForQuery: 7,

    // DDNS service for SkyWallets
    ddnsServer: 'ddns.cloudcoin.global'
  };
  serverList: any = [];
  serverNumber: any = 0;
  result: any;
  serverResponseList: any = [];
  public raidaleftJS: any = new RaidaJS(this.Options);
  constructor(private changeDetectorRef: ChangeDetectorRef, private raida: RaidaServiceService, public eventService: EventService, private auth: LoginService, private router: Router) {
    if (this.auth.getLoggedIn())
    {
      this.isLoggedIn = true;
    }
    this.auth.watch().subscribe((value) => {
      this.isLoggedIn = value;
    });
    this.raidaleftJS.apiEcho().then(response => {
      this.serverNumber = response.totalServers;
      let count: any = 0;
      for (let index = 0; index < response.totalServers; index++) {
        let onStatus = 'down';
        for (let online = 0; online < response.onlineServers; online++) {
          if (index === online){
            onStatus = 'ready';
          }
        }
        this.serverResponseList.push({
          server: count.toString(),
          status: onStatus,
          message: '',
          time: '',
          version: ''
        });
        count++;
      }
    });
    this.raidaleftJS.setTimeout(15000);
    this.raidaleftJS.setProtocol('https');
    this.raidaleftJS.setDomain('cloudcoin.global');
    this.raidaleftJS.setDefaultNetworkNumber(1);
    this.raidaleftJS.setDefaultRAIDA(7);
   }

  ngOnInit(): void {



    this.changeDetectorRef.detectChanges();
    this.serverList = this.raidaleftJS.getServers();
    this.serverList.forEach((e, i) => {
      setInterval(() => {
        this.serverRequest(e);
      }, 1000000);
    });

  }
  fireEvent(){
    this.eventService.emitSideBarFireEvent();
  }
  serverRequest(e){
    this.raida.serverRequest(e).subscribe(res => {
      this.result = res;
      this.serverResponseList.forEach((element, index) => {
        if (this.serverResponseList[index].server === this.result.server){
          this.serverResponseList[index] = this.result;
          // console.log(this.result);
        }
      });
    });
  }

  logout(): void
  {
    // if (confirm('Are you sure you want to log out ?'))
    // {
    //   localStorage.setItem('cc', 'null');
    //   localStorage.removeItem('cc');
    //   this.auth.checkLoginStatus();
    //   this.router.navigate(['/welcome']);
    // }
    Swal.fire({
      title: 'Are you sure you want to log out?',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonText: 'Yes',
      cancelButtonText: 'No'
    }).then((result) => {
      if (result.value) {
        {
          localStorage.setItem('cc', 'null');
          localStorage.removeItem('cc');
          this.auth.checkLoginStatus();
          this.router.navigate(['/welcome']);
        }
        this.fireEvent();
      }
    })
  }
}
