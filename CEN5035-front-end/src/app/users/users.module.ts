import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UsersRoutingModule } from './users-routing.module';
import { UserListComponent } from './user-list/user-list.component';
import { UserSingleComponent } from './user-single/user-single.component';
import { AnalyticsCardComponent } from './analytics-card/analytics-card.component';


@NgModule({
  declarations: [
    UserListComponent,
    UserSingleComponent,
    AnalyticsCardComponent
  ],
  imports: [
    CommonModule,
    UsersRoutingModule
  ]
})
export class UsersModule { }
