import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ContractDrafterComponent } from '../contract-drafter/contract-drafter.component';
import { MessageDrafterComponent } from '../message-drafter/message-drafter.component';
import { NotificationDrafterComponent } from '../notification-drafter/notification-drafter.component';
import { UserListComponent } from './user-list/user-list.component';
import { UserSingleComponent } from './user-single/user-single.component';

const routes: Routes = [
  {
    path: '',
    component: UserListComponent
  },
  {
    path: ':username',
    component: UserSingleComponent
  },
  {
    path: ':username/message-draft',
    component: MessageDrafterComponent
  },
  {
    path: ':username/contract-draft',
    component: ContractDrafterComponent
  },
  {
    path: ':username/notification-draft',
    component: NotificationDrafterComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }
