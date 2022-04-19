import { Component, OnInit } from '@angular/core';
import{ GlobalComponent } from './../global-component';
import { FormBuilder, FormControl, ReactiveFormsModule } from '@angular/forms';
import {HttpClient, HttpHeaders} from '@angular/common/http';
@Component({
  selector: 'app-message-drafter',
  templateUrl: './message-drafter.component.html',
  styleUrls: ['./message-drafter.component.css']
})
export class MessageDrafterComponent implements OnInit {

  readonly headers = new HttpHeaders().set('Content-Type', 'undefined').set('Access-Control-Allow-Origin', "http://localhost:8080");
  formData = new FormData();
  messageForm = this.formBuilder.group({
    sender: '',
    receiver: '',
    message: '',
    time: 0
  });
  Message = {
    sender: '',
    receiver: '',
    message: '',
    time: ''
  }

  constructor(private http: HttpClient, private formBuilder: FormBuilder) { }

  ngOnInit(): void {
  }

  onSubmit(): void {

    this.Message.sender = "fakeaccount@fakeaccount.com";
    // const current = new Date();

    // current.setHours(0)

    // current.setMinutes(0)

    // current.setSeconds(0)

    // current.setMilliseconds(0)

    this.Message.time = "0";
    
    console.log(this.Message);
    var stringMessage = JSON.stringify(this.Message)
    console.log(stringMessage)
    console.log(this.Message)
    this.http.post<any>("http://localhost:4200/api/send-message",stringMessage)
    .subscribe();
    this.draftOpen("/users")
  }

  setMessage(event: Event): void{
    this.Message.message = (<HTMLInputElement>event.target).value;
  }

  setReceiver(event: Event): void{
    this.Message.receiver = (<HTMLInputElement>event.target).value;
  }


  draftOpen(url:string): void{
    window.open(url + "/" + GlobalComponent.email, "_self")
  }

}
