import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-analytics-card',
  templateUrl: './analytics-card.component.html',
  styleUrls: ['./analytics-card.component.css']
})
export class AnalyticsCardComponent implements OnInit {
  cardData = {notificationCount: 100, contractCount: 300, messageCount: 200}

  constructor() { }

  ngOnInit(): void {
  }

}
