import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-field-search',
  templateUrl: './field-search.component.html',
  styleUrls: ['./field-search.component.css']
})
export class FieldSearchComponent implements OnInit {
  @Input('element')
  fieldElement!: { name: string; };

  constructor() { }

  ngOnInit(): void {
  }

}
