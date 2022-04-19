import { Component, OnInit } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http'
import { Router } from '@angular/router';
import { NgForm } from '@angular/forms';
import { ContractDraft } from './contract_draft';

@Component({
  selector: 'app-contract-drafter',
  templateUrl: './contract-drafter.component.html',
  styleUrls: ['./contract-drafter.component.css']
})

export class ContractDrafterComponent implements OnInit {
  isError: boolean = false
  isSuccessful: boolean = false
  
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json' )


  constructor(
    private http: HttpClient,
    private router: Router) {}

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

  onSubmit(f: NgForm) {
    this.isSuccessful = false
    this.isError = false
    const body = {
      "userID" : f.value.userID,
      "attorney_list" : f.value.attorney_list,
      "client_list" : f.value.client_list,
      "contract_title" : f.value.contract_title,
      "date" : f.value.date,
      "termination_date" : f.value.termination_date,
      "payment_type" : f.value.payment_type,
      "notes" : f.value.notes,
      "resume" : f.value.resume
    };
    return this.http.post<ContractDraft>("/api/contract-draft", body).subscribe(response => {this.isError=false; this.isSuccessful=true; console.log(body)}, err => {this.isError=true; this.isSuccessful=false});
  }

}
