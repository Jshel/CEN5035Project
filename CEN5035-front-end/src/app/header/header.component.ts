import { Component, OnInit } from '@angular/core';
import { GlobalComponent } from '../global-component';
import { User2 } from '../login-form/login-form.component';
import {HttpClient, HttpHeaders} from '@angular/common/http'

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  loginStatus : boolean = false

  readonly headers = new HttpHeaders().set('Content-Type', 'application/json')
  constructor(
    private http: HttpClient) {}

  async ngOnInit(): Promise<void> {
    const response = await this.http.get<User2>("/api/getuser").toPromise();
    this.loginStatus = (response != undefined) ? true : false
  }

  onLogout(){
    
  }

}
