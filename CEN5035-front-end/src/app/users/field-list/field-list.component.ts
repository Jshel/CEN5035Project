import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { GlobalComponent } from 'src/app/global-component';
import {HttpClient, HttpHeaders} from '@angular/common/http';

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
  attorney_ID: number;
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
  attorney_ID: -1,
  client_name: "",
  client_ID: -1
}
];

@Component({
  selector: 'app-field-list',
  templateUrl: './field-list.component.html',
  styleUrls: ['./field-list.component.css']
})

export class FieldListComponent implements OnInit {
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json').set('Access-Control-Allow-Origin', "http://localhost:8080");
  contractColumns: string[] = ['contract_ID', 'contract_type', 'client_name', 'termination_date'];
  isModalToggled = false;
  rowName = "";
  @Input('element')
  elementAttributes!: { name: string; };
  @Output() toggleModal = new EventEmitter<{isModalToggled: boolean, name: string, modalID: string}>();

  rowElements = CONTRACT_DATA;
  draftRowElements = CONTRACT_DATA;
  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.http.get<ContractExample>("http://localhost:4200/api/get-contract?attorneyID=00000001&contractID=00000000", {headers: this.headers})
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
    window.open("/users/" + GlobalComponent.email + "/" + str.toLowerCase().slice(0,-1) + "-draft", "_self")
  }

  showModal(thisname: string){
    this.isModalToggled = true
    this.toggleModal.emit({isModalToggled: this.isModalToggled, name: thisname, modalID: "1"})
  }
}
