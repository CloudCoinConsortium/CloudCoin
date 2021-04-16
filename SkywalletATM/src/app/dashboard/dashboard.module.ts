import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DashboardRoutingModule } from './dashboard-routing.module';
import { DashboardComponent } from './dashboard.component';
import { LeftPanelComponent } from './left-panel/left-panel.component';
import { HomeComponent } from './home/home.component';
import { TooltipModule } from 'ng2-tooltip-directive';

@NgModule({
  declarations: [DashboardComponent, LeftPanelComponent],
  imports: [
    CommonModule,
    DashboardRoutingModule,
    TooltipModule
  ]
})
export class DashboardModule { }
