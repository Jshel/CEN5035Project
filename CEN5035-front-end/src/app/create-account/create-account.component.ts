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

<<<<<<< HEAD
  isError:boolean = false
=======
  isError: boolean = false
>>>>>>> 2961a4c8b9a0388286cb481a5aa2dfc18f7526da
  isSuccessful: boolean = false
  
  constructor(private http: HttpClient) {}

  onSubmit(f: NgForm) {
    this.removeNotification()
    const body = {
      "name": f.value.name,
      "email": f.value.email,
      "username": f.value.username,
      "password": f.value.password
    };
<<<<<<< HEAD
    return this.http.post<UserRegistration>("/api/create-account", body).subscribe(response => {this.isError=false; this.isSuccessful=true},err =>{this.isError=true; this.isSuccessful=false});
  }

  removeNotification(){
    this.isError=false
    this.isSuccessful=false
=======
    return this.http.post<UserRegistration>("/api/create-account", body).subscribe(response => {this.isError=false; this.isSuccessful=true}, err => {this.isError=true; this.isSuccessful=false});
  }

  removeNotification(){
    this.isSuccessful = false
    this.isError = false
>>>>>>> 2961a4c8b9a0388286cb481a5aa2dfc18f7526da
  }


  ngOnInit(): void {
  }

}

