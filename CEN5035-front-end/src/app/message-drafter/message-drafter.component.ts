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
  formData = new FormData;
  messageForm = this.formBuilder.group({
    sender: '',
    receiver: '',
    message: '',
    time: 0
  });


  constructor(private http: HttpClient, private formBuilder: FormBuilder) { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    this.formData.set('sender', 'bob');
    this.formData.set('time', '0')
    console.log(this.formData.get('recipient'))
    var Message = JSON.stringify(this.formData)

    var stringMessage = JSON.stringify(Message)
    console.log(stringMessage)
    console.log(Message)
    this.http.post<any>("http://localhost:4200/api/send-message",Message)
    .subscribe();
  }

  draftOpen(url:string): void{
    window.open(url + "/" + GlobalComponent.email)
  }

}
