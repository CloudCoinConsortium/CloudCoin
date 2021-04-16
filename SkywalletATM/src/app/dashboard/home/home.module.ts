import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { HomeRoutingModule } from './home-routing.module';
import { HomeComponent } from './home.component';
import { WelcomeComponent } from './welcome/welcome.component';
import { NgxDropzoneModule } from 'ngx-dropzone';
import { RaidamailComponent } from './raidamail/raidamail.component';


@NgModule({
  declarations: [HomeComponent, WelcomeComponent, RaidamailComponent],
  imports: [
    CommonModule,
    HomeRoutingModule,
    NgxDropzoneModule
  ]
})
export class HomeModule { }
