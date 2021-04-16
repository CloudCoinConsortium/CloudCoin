import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BalanceRoutingModule } from './balance-routing.module';
import { BalanceComponent } from './balance.component';
import { LottieModule } from 'ngx-lottie';
import {FontAwesomeModule} from '@fortawesome/angular-fontawesome';
import { TooltipModule } from 'ng2-tooltip-directive';
@NgModule({
  declarations: [BalanceComponent],
  imports: [
    CommonModule,
    BalanceRoutingModule,
    LottieModule,
    FontAwesomeModule,
    TooltipModule
  ]
})
export class BalanceModule { }
