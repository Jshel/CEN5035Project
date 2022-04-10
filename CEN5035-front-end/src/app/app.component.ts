<<<<<<< HEAD
import { Component, OnInit } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http'
import { User2 } from './login-form/login-form.component';
=======
import { Component, OnInit} from '@angular/core';
import{ GlobalComponent } from './global-component';
>>>>>>> main

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  title = 'Contract Management';

<<<<<<< HEAD
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json')
  constructor(
    private http: HttpClient) {}

  ngOnInit(): void {
    this.http.get<User2>("/api/getuser").subscribe((response) => {console.log(response)}, err => {console.log("no session found")});
  }
}
=======
  ngOnInit(): void {
    this.title = GlobalComponent.username;
    GlobalComponent.output = "Hello";
    console.log(GlobalComponent.output)
  }
}
>>>>>>> main
