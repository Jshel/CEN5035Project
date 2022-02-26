import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import {HttpClient, HttpHeaders} from '@angular/common/http'
import {UserRegistration} from './UserRegistration'

@Component({
  selector: 'app-create-account',
  templateUrl: './create-account.component.html',
  styleUrls: ['./create-account.component.css']
})
export class CreateAccountComponent implements OnInit {

  readonly headers = new HttpHeaders().set('Content-Type', 'application/json');
  
  constructor(private http: HttpClient) { 

  }

  onSubmit(f: NgForm) {
    console.log(f.value); 
    console.log(f.valid); 
    alert(`Submitting Form with Name:${f.value.name}, Username: ${f.value.username} , Email: ${f.value.email} , Password: ${f.value.password}`)
    const body = {
      "name": f.value.name,
      "email": f.value.email,
      "username": f.value.username,
      "password": f.value.password
    };
    return this.http.post<UserRegistration>("/api/create-account", body, {headers: this.headers}).subscribe(response => console.log(response));
    //this.authenticate(f.value.username, f.value.password)
  }

  // authenticate(username: string, password: string){
  //   const url = '/api/login';
  //   const body = {
  //     "username": username,
  //     "password": password
  //   };
    
  //   return this.http.post<UserRegistration>(url, body, {headers: this.headers}).subscribe();
  // }


  ngOnInit(): void {
  }

}

