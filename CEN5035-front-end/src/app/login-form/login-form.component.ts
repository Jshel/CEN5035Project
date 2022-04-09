import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import {HttpClient, HttpHeaders} from '@angular/common/http'
import {User} from './user'
import { Router } from '@angular/router';

export class User2 {
  constructor(
    public name?: string,
    public email?: string,
    public username?: string,
    public password?: string,
  ) {}
}

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})

export class LoginFormComponent implements OnInit {

  isError: boolean = false
  isSuccessful: boolean = false
  
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json')
  constructor(
    private http: HttpClient,
    private router: Router) {}

  onSubmit(f: NgForm) {
    this.removeNotification()
    const body = {
      "email": f.value.email,
      "password": f.value.password,
    };


    this.http.get<User2>("/api/getuser").subscribe((response) => {console.log(response)}, err => {console.log("no session found")});

    this.http.post<User>("/api/login", body).subscribe(response => {this.isError=false; this.isSuccessful=true, this.router.navigateByUrl("/browse-contracts");}, err => {this.isError=true; this.isSuccessful=false});
  }
  
  removeNotification(){
    this.isSuccessful = false
    this.isError = false
  }
  

  ngOnInit(): void {
  }

}
