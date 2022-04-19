import { Component, OnInit } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';

@Component({
  selector: 'app-analytics-card',
  templateUrl: './analytics-card.component.html',
  styleUrls: ['./analytics-card.component.css']
})
export class AnalyticsCardComponent implements OnInit {
  readonly headers = new HttpHeaders().set('Content-Type', 'undefined').set('Access-Control-Allow-Origin', "http://localhost:8080");
  cardData = {contractCount: 0, messageCount: 0}

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.http.get<number>(encodeURI("http://localhost:4200/api/count-contracts?attorney_email=fakeaccount@fakeaccount.com"), {headers: this.headers})
      .subscribe(
         (response) => {
            this.cardData.contractCount = response;
          },
         (error) => { console.log("Error Loading Contract Count" + error)}
      );

      this.http.get<number>(encodeURI("http://localhost:4200/api/count-messages?attorney_email=fakeaccount@fakeaccount.com"), {headers: this.headers})
      .subscribe(
         (response) => {
            this.cardData.messageCount = response;
          },
         (error) => { console.log("Error Loading Message Count" + error)}
      );
  }

}
