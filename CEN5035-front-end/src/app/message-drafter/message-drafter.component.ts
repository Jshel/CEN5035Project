import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-message-drafter',
  templateUrl: './message-drafter.component.html',
  styleUrls: ['./message-drafter.component.css']
})
export class MessageDrafterComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  draftOpen(url:string): void{
    window.open(url)
  }

}
