import { Component, OnInit } from '@angular/core';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';
import {HttpClient, HttpHeaders} from '@angular/common/http';

@Component({
  selector: 'app-contract-drafter',
  templateUrl: './contract-drafter.component.html',
  styleUrls: ['./contract-drafter.component.css']
})
export class ContractDrafterComponent implements OnInit {
  readonly headers = new HttpHeaders().set('Content-Type', 'undefined').set('Access-Control-Allow-Origin', "http://localhost:8080");
  contractForm = this.formBuilder.group({
    contract_type: '',
    termination_date: '',
    payment_type: '',
    ammount_paid: '',
    ammount_owed: '',
    client_email: '',
    client_name: '',
    contract: ''
  });

  constructor(private http: HttpClient, private formBuilder: FormBuilder,) { }
  clients: number[] = [0]
  clientEmails: number[] = [0]
  ngOnInit(): void {
  }

  addPerson(isClient:boolean): void{
    if (isClient){
      this.clients.push(this.clients.length)
    }
    else{
      this.clientEmails.push(this.clientEmails.length)
    }
  }

  removePerson(personNumber: number, isClient:boolean): void{
    if (isClient){
      this.clients = this.clients.filter(item => item !== personNumber);
    }
    else{
      this.clientEmails = this.clientEmails.filter(item => item !== personNumber);
    }
  }

  draftOpen(url:string): void{
    window.open(url, "_self")
  }

  onSubmit(): void {
    const formData = new FormData;
    formData.append('contract', this.contractForm.get('contract')!.value, 'contract.pdf')
    console.log(this.headers)
    this.http.post<any>("http://localhost:4200/api/upload?contract_type=" + this.contractForm.get('contract_type')!.value + "&termination_date="  + this.contractForm.get('termination_date')!.value + "&payment_type=" + this.contractForm.get('payment_type')!.value  + "&ammount_paid=" + this.contractForm.get('ammount_paid')!.value + "&ammount_owed=" + this.contractForm.get('ammount_owed')!.value + "&client_email=" + this.contractForm.get('client_email')!.value + "&client_name=" + this.contractForm.get('client_name')!.value, formData)
    .subscribe();
  }
}
