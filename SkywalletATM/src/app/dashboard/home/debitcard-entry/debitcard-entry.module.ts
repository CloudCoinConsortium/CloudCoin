import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { DebitcardEntryRoutingModule } from './debitcard-entry-routing.module';
import { DebitcardEntryComponent } from './debitcard-entry.component';
import { NgxMaskModule, IConfig } from 'ngx-mask';
import { LottieModule } from 'ngx-lottie';
import {FontAwesomeModule} from '@fortawesome/angular-fontawesome';

// export const options: Partial<IConfig> | (() => Partial<IConfig>) = null;
const maskConfig: Partial<IConfig> = {
  validation: true,
};

@NgModule({
  declarations: [DebitcardEntryComponent],
  imports: [
    CommonModule,
    FormsModule,
    DebitcardEntryRoutingModule,
    NgxMaskModule.forRoot(),
    LottieModule,
    FontAwesomeModule
  ]
})
export class DebitcardEntryModule { }
