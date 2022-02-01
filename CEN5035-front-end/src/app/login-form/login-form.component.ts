import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';

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
  }
  
  constructor() { }

  ngOnInit(): void {
  }

}
