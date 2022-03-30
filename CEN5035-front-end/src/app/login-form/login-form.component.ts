import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import {HttpClient, HttpHeaders} from '@angular/common/http'
import {User} from './user'

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})
export class LoginFormComponent implements OnInit {

  isError: boolean = false
  isSuccessful: boolean = false

  readonly headers = new HttpHeaders().set('Content-Type', 'application/json')
  
  constructor(private http: HttpClient) {}

  onSubmit(f: NgForm) {
    this.removeNotification()
    const body = {
      "email": f.value.email,
      "password": f.value.password
    };
    return this.http.post<User>("/api/login", body).subscribe(response => {this.isError=false; this.isSuccessful=true}, err => {this.isError=true; this.isSuccessful=false});
  }
  
  removeNotification(){
    this.isSuccessful = false
    this.isError = false
  }
  

  ngOnInit(): void {
  }

}
