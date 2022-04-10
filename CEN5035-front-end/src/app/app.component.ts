import { Component, OnInit} from '@angular/core';
import{ GlobalComponent } from './global-component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  title = 'Contract Management';

  ngOnInit(): void {
    this.title = GlobalComponent.username;
    GlobalComponent.output = "Hello";
    console.log(GlobalComponent.output)
  }
}
