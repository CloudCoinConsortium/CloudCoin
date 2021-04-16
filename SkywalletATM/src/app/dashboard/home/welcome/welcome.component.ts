import { Component, OnInit } from '@angular/core';
import {LoginService} from '../../../service/login.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.scss']
})
export class WelcomeComponent implements OnInit {

  constructor( private auth: LoginService, private router: Router) { }

  ngOnInit(): void {
    if (this.auth.getLoggedIn())
    {
      this.router.navigate(['/balance']);
    }
  }

}
