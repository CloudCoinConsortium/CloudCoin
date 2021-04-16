import { ChangeDetectorRef } from '@angular/core';
import { Component, OnInit } from '@angular/core';
import RaidaJS from 'raidajs';
import { RaidaServiceService } from 'src/app/raida-service.service';
import { EventService } from 'src/app/service/event.service';
import {LoginService} from '../../service/login.service';
import {Router} from '@angular/router';
import Swal from 'sweetalert2';
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})

export class HomeComponent implements OnInit {
  public isLoggedIn = false;
  public isDiscModalShow:boolean;
  public isSupportModalShow:boolean;
  constructor(private changeDetectorRef: ChangeDetectorRef, private raida: RaidaServiceService, private eventService: EventService,
              private auth: LoginService, private router: Router) {
    this.isDiscModalShow = false;
    this.isSupportModalShow = false;
    if (this.auth.getLoggedIn())
    {
      this.isLoggedIn = true;
    }
    this.auth.watch().subscribe((value) => {
      // alert("observed");
      //  console.log(value);
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

  files: File[] = [];

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
  public raidaleftJS = new RaidaJS(this.Options);

	onSelect(event): void {
		console.log(event);
		this.files.push(...event.addedFiles);
	}

	onRemove(event): void {
		console.log(event);
		this.files.splice(this.files.indexOf(event), 1);
	}

  ngOnInit(): void {


    // // console.log();
    // this.serverList = this.raidaJS.getServers();
    // // console.log(this.serverList);
    // this.serverList.forEach(e=>{this.http.get(e+'/service/echo').subscribe(res=>{
    //     this.serverResponseList.push(res);
    //   });
    // });

    // console.log(this.serverResponseList);
    // this.raidaJS.apiEcho().then(response => {
    //   console.log("Available servers: " + response.onlineServers);
    //   console.log("Response",response);
    // })
    this.changeDetectorRef.detectChanges();
    this.serverList = this.raidaleftJS.getServers();
    this.serverList.forEach((e, i) => {
      setInterval(() => {
        this.serverRequest(e);
      }, 1000000);
    });

  }


  fireEvent(): void{
    this.eventService.emitSideBarFireEvent();
  }
  serverRequest(e){
    this.raida.serverRequest(e).subscribe(res => {
      this.result = res;
      this.serverResponseList.forEach((element, index) => {
        if (this.serverResponseList[index].server === this.result.server){
          this.serverResponseList[index] = this.result;
          // console.log(this.result);/**/
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
          localStorage.setItem('skywallet', 'null');
          localStorage.removeItem('skywallet');
          this.auth.checkLoginStatus();
          this.router.navigate(['/welcome']);
        }
      }
    });
  }
  toggleDiscModal(){
    this.isDiscModalShow = !this.isDiscModalShow;
  }
  toggleSupportModal(){
    this.isSupportModalShow = !this.isSupportModalShow;
  }

  logOut(): void {
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
      }
    })
  }
}

