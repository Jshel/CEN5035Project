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

  readonly headers = new HttpHeaders().set('Content-Type', 'application/json');
  
  constructor(private http: HttpClient) { 

  }

  onSubmit(f: NgForm) {
    console.log(f.value); 
    console.log(f.valid); 
    alert("Submitting Form with Username: " + f.value.username + " and Password: " + f.value.password)
    const body = {
      "username": f.value.username,
      "password": f.value.password
    };
    //return this.http.post<User>("/api/login", body, {headers: this.headers}).subscribe(response => console.log(response));
    this.authenticate(f.value.username, f.value.password)
  }

  authenticate(username: string, password: string){
    const url = '/api/login';
    const body = {
      "username": username,
      "password": password
    };
    
    return this.http.post<User>(url, body, {headers: this.headers}).subscribe();
  }
  
  

  ngOnInit(): void {
  }

}
