import { NgModule } from '@angular/core';
import { flush } from '@angular/core/testing';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home.component';
import { RaidamailComponent } from './raidamail/raidamail.component';
import { WelcomeComponent } from './welcome/welcome.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    children: [
      {
        path: '',
        component: WelcomeComponent
      },
      {
        path: 'welcome',
        redirectTo: ''
      },
      {
        path: 'login',
        loadChildren: () => import('./login/login.module').then(m => m.LoginModule)
      },
      {
        path: 'register',
        loadChildren: () => import('./register/register.module').then(m => m.RegisterModule)
      },
      {
        path: 'card-entry',
        loadChildren: () => import('./debitcard-entry/debitcard-entry.module').then(m => m.DebitcardEntryModule)
      },
      {
        path: 'drop-debit',
        loadChildren: () => import('./drop-debitcard/drop-debitcard.module').then(m => m.DropDebitcardModule)
      },
      {
        path: 'balance',
        loadChildren: () => import('./balance/balance.module').then(m => m.BalanceModule)
      },
      {
        path: 'withdraw',
        loadChildren: () => import('./withdraw/withdraw.module').then(m => m.WithdrawModule)
      },
      {
        path: 'transfer',
        loadChildren: () => import('./transfer/transfer.module').then(m => m.TransferModule)
      },
      {
        path: 'payment',
        loadChildren: () => import('./payment/payment.module').then(m => m.PaymentModule)
      },
      {
        path: 'deposit',
        loadChildren: () => import('./deposit/deposit.module').then(m => m.DepositModule)
      },
      {
        path: 'raida-mail',
        component: RaidamailComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class HomeRoutingModule { }
