import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DropDebitcardComponent } from './drop-debitcard.component';

const routes: Routes = [
  {
    path: '',
    component: DropDebitcardComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DropDebitcardRoutingModule { }
