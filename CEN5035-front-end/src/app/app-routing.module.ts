import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginFormComponent } from './login-form/login-form.component';
import { BrowseContractsComponent } from './browse-contracts/browse-contracts.component';
import { CreateAccountComponent } from './create-account/create-account.component';
import { MessageDrafterComponent } from './message-drafter/message-drafter.component';
import { ContractDrafterComponent } from './contract-drafter/contract-drafter.component';
const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    pathMatch: 'full'
  },
  {
    path: 'login',
    component: LoginFormComponent
  },
  {
    path: 'message-draft',
    component: MessageDrafterComponent
  },
  {
    path: 'contract-draft',
    component: ContractDrafterComponent
  },
  {
    path: 'users',
    loadChildren: () => import('./users/users.module').then(mod => mod.UsersModule)
  },
  {
    path: 'browse-contracts',
    component: BrowseContractsComponent
  },
  {
    path: 'create-account',
    component: CreateAccountComponent
  },
  {
    path: '**',
    component: HomeComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
