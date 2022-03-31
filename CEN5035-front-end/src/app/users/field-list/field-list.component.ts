import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-field-list',
  templateUrl: './field-list.component.html',
  styleUrls: ['./field-list.component.css']
})
export class FieldListComponent implements OnInit {
  isModalToggled = false;
  rowName = "";
  @Input('element')
  elementAttributes!: { name: string; };
  @Output() toggleModal = new EventEmitter<{isModalToggled: boolean, name: string, modalID: string}>();

  rowElements = [{name: 'Contract 1'},{name: 'Contract 2'},{name: 'Contract 3'}]
  draftRowElements = [{name: 'Contract 1'},{name: 'Contract 2'},{name: 'Contract 3'}]
  constructor() { }

  ngOnInit(): void {
  }

  draftOpen(): void{
    var str = this.elementAttributes.name
    console.log(str)
    window.open("/" + str.toLowerCase().slice(0,-1) + "-draft","_blank")
  }

  showModal(thisname: string){
    this.isModalToggled = true
    this.toggleModal.emit({isModalToggled: this.isModalToggled, name: thisname, modalID: "1"})
  }
}
