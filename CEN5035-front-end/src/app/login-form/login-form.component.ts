import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})
export class LoginFormComponent implements OnInit {
  onSubmit(f: NgForm) {
    console.log(f.value); 
    console.log(f.valid); 
    alert("Submitting Form with Username: " + f.value.username + " and Password: " + f.value.password)
    
    //should move this to a function that gets a response from backend saying true or false for login
  //   this.http.post('/api/login', {
  //     "user": f.value.username,
  //     "password": f.value.password
  //   })
}
  
  constructor(
    private httpClient: HttpClient
  ) { }

  ngOnInit(): void {
  }


}
