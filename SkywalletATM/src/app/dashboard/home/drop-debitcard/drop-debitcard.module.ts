import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DropDebitcardRoutingModule } from './drop-debitcard-routing.module';
import { DropDebitcardComponent } from './drop-debitcard.component';
import { NgxDropzoneModule } from 'ngx-dropzone';
import { LottieModule } from 'ngx-lottie';
import {FontAwesomeModule} from '@fortawesome/angular-fontawesome';

@NgModule({
  declarations: [DropDebitcardComponent],
  imports: [
    CommonModule,
    DropDebitcardRoutingModule,
    NgxDropzoneModule,
    LottieModule,
    FontAwesomeModule
  ]
})
export class DropDebitcardModule { }
