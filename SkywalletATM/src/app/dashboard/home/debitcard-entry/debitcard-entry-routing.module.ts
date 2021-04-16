import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DebitcardEntryComponent } from './debitcard-entry.component';

const routes: Routes = [
  {
    path: '',
    component: DebitcardEntryComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DebitcardEntryRoutingModule { }
