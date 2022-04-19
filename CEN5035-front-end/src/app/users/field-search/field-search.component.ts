import { Component, Input, OnInit } from '@angular/core';
import { FieldListComponent } from '../field-list/field-list.component';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import{ GlobalComponent } from '../../global-component';
export interface ContractExample {
  contract_ID: string;
  contract_type: string;
  date_created: string;
  termination_date: string;
  valid_signiture: boolean;
  payment_type: string;
  amount_paid: number;
  amount_owed: number;
  attorney_name: string;
  attorney_email: string;
  client_name: string;
  client_ID: number;
}

const CONTRACT_DATA: ContractExample[] = [
{
  contract_ID: "",
  contract_type: "",
  date_created: "",
  termination_date: "",
  valid_signiture: false,
  payment_type: "",
  amount_paid: -1,
  amount_owed: -1,
  attorney_name: "",
  attorney_email: "",
  client_name: "",
  client_ID: -1
}
];

export interface MessageExample {
  sender: string;
  receiver: string;
  message: string;
  time: string;
}

const MESSAGE_DATA: MessageExample[] = [
{
  sender: "",
  receiver: "",
  message: "",
  time: "",
}
];

@Component({
  selector: 'app-field-search',
  templateUrl: './field-search.component.html',
  styleUrls: ['./field-search.component.css']
})
export class FieldSearchComponent implements OnInit {
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json').set('Access-Control-Allow-Origin', "http://localhost:8080");
  @Input('element')
  fieldElement!: { name: string; };


  contractList = CONTRACT_DATA;
  messageList = MESSAGE_DATA;

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
  }

}
