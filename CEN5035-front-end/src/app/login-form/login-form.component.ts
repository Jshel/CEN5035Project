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

<<<<<<< HEAD
  isError:boolean = false
  isSuccessful: boolean = false
=======
  isError: boolean = false
  isSuccessful: boolean = false
  
  constructor(private http: HttpClient) { 
>>>>>>> 2961a4c8b9a0388286cb481a5aa2dfc18f7526da

  readonly headers = new HttpHeaders().set('Content-Type', 'application/json')
  
  constructor(private http: HttpClient) {}

  onSubmit(f: NgForm) {
    this.removeNotification()
    const body = {
      "email": f.value.email,
      "password": f.value.password
    };
<<<<<<< HEAD
    return this.http.post<User>("/api/login", body).subscribe(response => {this.isError=false; this.isSuccessful=true}, err =>{this.isError=true; this.isSuccessful=false})
  }

  removeNotification(){
    this.isError=false
    this.isSuccessful=false
=======
    return this.http.post<User>("/api/login", body).subscribe(response => {this.isError=false; this.isSuccessful=true}, err => {this.isError=true; this.isSuccessful=false});
  }
  
  removeNotification(){
    this.isSuccessful = false
    this.isError = false
>>>>>>> 2961a4c8b9a0388286cb481a5aa2dfc18f7526da
  }
  

  ngOnInit(): void {
  }

}
