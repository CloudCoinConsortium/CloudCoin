import { Component, OnInit } from '@angular/core';
import { EventService } from '../service/event.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  public leftBarShow: boolean;
  constructor(private eventService: EventService) {
    this.leftBarShow = true;
  }

  ngOnInit(): void {
    if (window.innerWidth < 993) {
      this.leftBarShow = false;
    }

    this.eventService.getSideBarFireEvent().subscribe(() => {
      this.leftBarShow = false;
    });
  }

  sideBarCaller(){
    if (window.innerWidth < 993) {
      this.leftBarShow = true;
    }
  }
  sideBarCollapse(){
    if (window.innerWidth < 993) {
      this.leftBarShow = false;
    }
  }

}
