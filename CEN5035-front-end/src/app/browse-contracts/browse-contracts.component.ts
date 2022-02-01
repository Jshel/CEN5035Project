import {MatSidenavModule} from '@angular/material/sidenav';
import {MatTableModule} from '@angular/material/table';
import { Component, OnInit } from '@angular/core';


export interface ContractExample {
  parties: string[];
  id: number;
  signed: boolean;
  deadline: string; //Store as yyyy/mm/dd
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
  {parties: ["John", "Joe"], id: 0, signed: true, deadline: "2023/01/01"},
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
  tab:number = 1;
  contractColumns: string[] = ['parties', 'id', 'signed', 'deadline'];
  contracts = CONTRACT_DATA;
  messageColumns: string[] = ['recipients', 'id', 'received', 'contents', 'send_date'];
  messages = MESSAGE_DATA;
  notificationColumns: string[] = ['required_action', 'id', 'deadline'];
  notifications = NOTIFICATION_DATA;
  constructor() { 
  }

  ngOnInit(): void {
  }

  setTab(tabChoice:number): void{
    this.tab=tabChoice;
  }

  getTab(): number{
    return this.tab;
  }
}
