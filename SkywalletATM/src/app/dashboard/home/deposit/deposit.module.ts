import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DepositRoutingModule } from './deposit-routing.module';
import { DepositComponent } from './deposit.component';
import { NgxDropzoneModule } from 'ngx-dropzone';
import { LottieModule } from 'ngx-lottie';
import {FontAwesomeModule} from '@fortawesome/angular-fontawesome';

@NgModule({
  declarations: [DepositComponent],
  imports: [
    CommonModule,
    DepositRoutingModule,
    NgxDropzoneModule,
    LottieModule,
    FontAwesomeModule
  ]
})
export class DepositModule { }
