import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { DomSanitizer} from '@angular/platform-browser';
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
<<<<<<< HEAD

export interface MessageExample {
  sender: string;
  receiver: string;
  message: string;
  time: string;
}

export interface MessageSet {
  Messages: MessageExample[]
}

const MESSAGE_DATA:MessageExample[] = [
{
  sender: "",
  receiver: "",
  message: "",
  time: "",
}
];
=======
>>>>>>> qualityOfLifeFixes

@Component({
  selector: 'app-field-list',
  templateUrl: './field-list.component.html',
  styleUrls: ['./field-list.component.css']
})

<<<<<<< HEAD

=======
>>>>>>> qualityOfLifeFixes
export class FieldListComponent implements OnInit {
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json').set('Access-Control-Allow-Origin', "http://localhost:8080");
  contractColumns: string[] = ['contract_ID', 'contract_type', 'client_name', 'termination_date'];
  isModalToggled = false;
  rowName = "";
  rowElements = CONTRACT_DATA;
  messageElements = MESSAGE_DATA;
  @Input('element')
  elementAttributes!: { name: string; };
  @Output() toggleModal = new EventEmitter<{isModalToggled: boolean, name: string, modalID: string}>();
<<<<<<< HEAD
=======

  rowElements = CONTRACT_DATA;
  draftRowElements = CONTRACT_DATA;
>>>>>>> qualityOfLifeFixes
  constructor(private http: HttpClient, private domSanitizer: DomSanitizer) { }

  ngOnInit(): void {
    console.log(GlobalComponent.givenName)
    console.log(GlobalComponent.username)
     console.log(GlobalComponent.output)
      console.log(GlobalComponent.email)
      this.http.get<ContractExample>("http://localhost:8080/api/count-messages?attorney_email=bob", {headers: this.headers})
      .subscribe(
         (response) => {
            console.log("Counts " + response);
          },
         (error) => { console.log("Counting error" + error)}
      );
    this.http.get<ContractExample>("http://localhost:4200/api/get-contract?username=" + GlobalComponent.givenName + "&contractID=0000000a", {headers: this.headers})
      .subscribe(
         (response) => {
            this.rowElements = [];
            this.rowElements.push(response as ContractExample);
            this.rowElements = [...this.rowElements];
            console.log(this.rowElements);
          },
         (error) => { console.log("Error happened" + error)}
      );
  }

  draftOpen(): void{
    var str = this.elementAttributes.name
    window.open("/users/" + GlobalComponent.username + "/" + str.toLowerCase().slice(0,-1) + "-draft", "_self")
  }

<<<<<<< HEAD
  showPDF(thisname: string, attorneyEmail: string){
    var pdfsrc = "/api/download?attorney_email=fakeaccount@fakeaccount.com" + "&contract_id=" + thisname
    window.open(pdfsrc, '_self')
  }

  showModal(thisname: string){
=======
  showModal(thisname: string, attorneyEmail: string){
>>>>>>> qualityOfLifeFixes
    this.isModalToggled = true
    var pdfsrc =  "http://" + window.location.host + "/api/download?attorney_email=" + attorneyEmail + "&contract_id=" + thisname
    window.open(pdfsrc)
  }

   searchValsContract(event:Event){
    let element = CONTRACT_DATA
    this.rowElements = this.rowElements!.filter(item => item == element[0]);
    this.http.get<ContractExample>(encodeURI("http://localhost:4200/api/get-contract?username=fakeaccount@fakeaccount.com&contractID=" + (<HTMLTextAreaElement>event.target).value), {headers: this.headers})
      .subscribe(
         (response) => {
            this.rowElements = CONTRACT_DATA;
            let element = (response as ContractExample)
            this.rowElements.push(element);
            this.rowElements = this.rowElements!.filter(item => item == element);
            console.log(this.rowElements);
          },
         (error) => { console.log("Error Loading Contracts" + error)}
      );
  }

  searchValsMessage(event:Event){
      this.http.get<MessageSet>(encodeURI("/api/get-message?sender=" + (<HTMLTextAreaElement>event.target).value + "&receiver=fakeaccount@fakeaccount.com&n=10"), {headers: this.headers})
      .subscribe(
         (response) => {

            this.messageElements = response.Messages;
          },
         (error) => { console.log("Error Loading Messages" + error)}
      );
  }
}