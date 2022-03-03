import {MatSidenavModule} from '@angular/material/sidenav';
import {MatTableModule} from '@angular/material/table';
import { Component, OnInit } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { map } from 'rxjs/operators';

export interface ContractExample {
  contract_ID: number;
  contract_type: string;
  date_created: string;
  termination_date: string;
  valid_signiture: boolean;
  payment_type: string;
  amount_paid: number;
  amount_owed: number;
  attorney_name: string;
  attorney_ID: number;
  client_name: string;
  client_ID: number;
}

export interface MessageExample {
  recipients: string[];
  id: number;
  received: boolean[];
  contents: string;
  send_date: string; //Store as yyyy/mm/dd
}

export interface NotificationExample {
  required_action: string;
  id: number;
  deadline: string; //Store as yyyy/mm/dd
}

const CONTRACT_DATA: ContractExample[] = [
{
  contract_ID: -1,
  contract_type: "fake",
  date_created: "never",
  termination_date: "always",
  valid_signiture: false,
  payment_type: "none",
  amount_paid: -1,
  amount_owed: -1,
  attorney_name: "fake name",
  attorney_ID: -1,
  client_name: "fake name",
  client_ID: -1
}
];

const MESSAGE_DATA: MessageExample[] = [
{recipients: ["John", "Joe"], id: 0, received: [false,false] , contents: "Sign the contract", send_date: "2022/01/01"},
];

const NOTIFICATION_DATA: NotificationExample[] = [
{required_action: "Do contract stuff", id: 0, deadline: "2022/01/01"},
];

@Component({
  selector: 'app-browse-contracts',
  templateUrl: './browse-contracts.component.html',
  styleUrls: ['./browse-contracts.component.css']
})
export class BrowseContractsComponent implements OnInit {
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json').set('Access-Control-Allow-Origin', "http://localhost:8080");
  tab:number = 1;
  contractColumns: string[] = ['contract_ID', 'contract_type', 'client_name', 'termination_date'];
  contracts = CONTRACT_DATA;
  messageColumns: string[] = ['recipients', 'id', 'received', 'contents', 'send_date'];
  messages = MESSAGE_DATA;
  notificationColumns: string[] = ['required_action', 'id', 'deadline'];
  notifications = NOTIFICATION_DATA;
  constructor(private http: HttpClient){}

  ngOnInit(): void {
  }

  setTab(tabChoice:number): void{
    this.tab=tabChoice;
    if (tabChoice == 3){
      this.http.get<ContractExample>("http://localhost:4200/api/get-contract?attorneyID=00000001&contractID=00000000", {headers: this.headers})
      .subscribe(
         (response) => {
            this.contracts = [];
            this.contracts.push(response as ContractExample);
            this.contracts = [...this.contracts];
            console.log(this.contracts);
          },
         (error) => { console.log("Error happened" + error)}
      );
    }
  }

  getTab(): number{
    return this.tab;
  }
}
