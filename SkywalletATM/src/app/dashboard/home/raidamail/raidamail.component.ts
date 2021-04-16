import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-raidamail',
  templateUrl: './raidamail.component.html',
  styleUrls: ['./raidamail.component.scss']
})
export class RaidamailComponent implements OnInit {

  public authCheck: boolean;

  constructor() {
    this.authCheck = false;
  }

  ngOnInit(): void {
  }

}
