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
  @Output() toggleModal = new EventEmitter<{isModalToggled: boolean, name: string, modalID: string, pdfURL: string}>();

  rowElements = CONTRACT_DATA;
  draftRowElements = CONTRACT_DATA;
  constructor(private http: HttpClient, private domSanitizer: DomSanitizer) { }

  ngOnInit(): void {
    console.log(GlobalComponent.username)
    this.http.get<ContractExample>("http://localhost:4200/api/get-contract?username=" + GlobalComponent.username + "&contractID=0000000a", {headers: this.headers})
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
    window.open("/" + str.toLowerCase().slice(0,-1) + "-draft", "_self")
  }

  showModal(thisname: string, attorneyEmail: string){
    this.isModalToggled = true
    var pdfsrc =  "http://" + window.location.host + "/api/download?attorney_email=" + attorneyEmail + "&contract_id=" + thisname
    // var sanitizedsrc = this.domSanitizer.sanitize(SecurityContext.URL, pdfsrc)
    window.open(pdfsrc)
  }
}
