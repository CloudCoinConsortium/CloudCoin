import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { LottieModule } from 'ngx-lottie';
import player from 'lottie-web/build/player/lottie_svg';
import { FontAwesomeModule, FaIconLibrary } from '@fortawesome/angular-fontawesome';
import {faArrowLeft, faCheck, faCloudDownloadAlt, faSync, faCreditCard, faExclamationCircle,
  faExclamationTriangle, faCoins, faBug} from '@fortawesome/free-solid-svg-icons';

export function playerFactory(): any {
  return import(/* webpackChunkName: 'lottie-web' */ 'lottie-web');
}


@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    LottieModule.forRoot({ player: playerFactory, useCache: true }),
    FontAwesomeModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { constructor(library: FaIconLibrary) {
  // Add an icon to the library for convenient access in other components
  library.addIcons(faArrowLeft, faSync, faCheck, faCloudDownloadAlt, faCreditCard, faExclamationCircle, faExclamationTriangle,
    faCoins, faBug);
}}
