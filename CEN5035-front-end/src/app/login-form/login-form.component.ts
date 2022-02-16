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
    const formData = new FormData();
    formData.append("username", f.value.username)
    formData.append("password", f.value.password)
    this.httpClient.get<any>('http://localhost:8080/api')
    alert("debug")
}
  
  constructor(
    private httpClient: HttpClient
  ) { }

  ngOnInit(): void {
  }


}
