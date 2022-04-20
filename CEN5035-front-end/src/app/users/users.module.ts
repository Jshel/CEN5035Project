import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UsersRoutingModule } from './users-routing.module';
import { UserListComponent } from './user-list/user-list.component';
import { UserSingleComponent } from './user-single/user-single.component';
import { AnalyticsCardComponent } from './analytics-card/analytics-card.component';
import { FieldSearchComponent } from './field-search/field-search.component';
import { FieldListComponent } from './field-list/field-list.component';
import { FieldListModalComponent } from './field-list-modal/field-list-modal.component';
import { MenuListComponent } from './menu-list/menu-list.component';
import { MenuListModalComponent } from './menu-list-modal/menu-list-modal.component';
import { SplashScreenComponent } from '../splash-screen/splash-screen.component';

@NgModule({
  declarations: [
    UserListComponent,
    UserSingleComponent,
    AnalyticsCardComponent,
    FieldSearchComponent,
    FieldListComponent,
    FieldListModalComponent,
    MenuListComponent,
    MenuListModalComponent,
    SplashScreenComponent
  ],
  imports: [
    CommonModule,
    UsersRoutingModule
  ]
})
export class UsersModule { }
