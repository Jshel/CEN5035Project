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
    this.authenticate(f.value.username, f.value.password)
  }

  authenticate(username: string, password: string){
    const url = 'http://localhost:4200/api/login';
    const body = {
      "username": username,
      "password": password
    };

      const temp = this.http.post<any>(url, body).subscribe()
  }
  
  constructor(private http: HttpClient) { }

  ngOnInit(): void {
  }

}
