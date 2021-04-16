import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { WithdrawRoutingModule } from './withdraw-routing.module';
import { WithdrawComponent } from './withdraw.component';
import { LottieModule } from 'ngx-lottie';
import {FontAwesomeModule} from '@fortawesome/angular-fontawesome';


@NgModule({
  declarations: [WithdrawComponent],
  imports: [
    CommonModule,
    WithdrawRoutingModule,
    FormsModule,
    LottieModule,
    FontAwesomeModule
  ]
})
export class WithdrawModule { }
