import { Component, ComponentFactoryResolver, OnInit } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http'
import { User2 } from './login-form/login-form.component';
import{ GlobalComponent } from './global-component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  title = 'Contract Management';
  email: string | undefined = "undefined"

  readonly headers = new HttpHeaders().set('Content-Type', 'application/json')
  constructor(
    private http: HttpClient,
    private router: Router) {}

  setCookies(cookieData: User2 | undefined){
    GlobalComponent.email = cookieData?.email
    GlobalComponent.username = cookieData?.username
    GlobalComponent.givenName = cookieData?.name
    console.log(GlobalComponent);
  }

  async ngOnInit(): Promise<void> {
    const response = await this.http.get<User2>("/api/getuser").toPromise();
    response?.email != undefined ? this.router.navigateByUrl("/users/" + response.email) : undefined;
    this.setCookies(response)
  }
}
