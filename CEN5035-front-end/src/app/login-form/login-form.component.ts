import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import {HttpClient} from '@angular/common/http'
import {User} from './user'
import { Router } from '@angular/router';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})
export class LoginFormComponent implements OnInit {

  isError: boolean = false
  isSuccessful: boolean = false
  
  constructor(
    private http: HttpClient,
    private router: Router) { 
    
  }

  onSubmit(f: NgForm) {
    this.removeNotification()
    const body = {
      "email": f.value.email,
      "password": f.value.password
    };
    return this.http.post<User>("/api/login", body).subscribe(response => {this.isError=false; this.isSuccessful=true; this.router.navigateByUrl("/browse-contracts");}, err => {this.isError=true; this.isSuccessful=false});
  }

  
  removeNotification(){
    this.isSuccessful = false
    this.isError = false
  }
  

  ngOnInit(): void {
  }

}
