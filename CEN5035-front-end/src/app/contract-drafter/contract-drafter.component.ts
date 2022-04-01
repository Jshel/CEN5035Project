import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-contract-drafter',
  templateUrl: './contract-drafter.component.html',
  styleUrls: ['./contract-drafter.component.css']
})
export class ContractDrafterComponent implements OnInit {

  constructor() { }
  clients: number[] = [0]
  attorneys: number[] = [0]
  ngOnInit(): void {
  }

  addPerson(isClient:boolean): void{
    if (isClient){
      this.clients.push(this.clients.length)
    }
    else{
      this.attorneys.push(this.attorneys.length)
    }
  }

  removePerson(personNumber: number, isClient:boolean): void{
    if (isClient){
      this.clients = this.clients.filter(item => item !== personNumber);
    }
    else{
      this.attorneys = this.attorneys.filter(item => item !== personNumber);
    }
  }

  draftOpen(url:string): void{
    window.open(url, "_self")
  }

}
