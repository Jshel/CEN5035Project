import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-field-list-modal',
  templateUrl: './field-list-modal.component.html',
  styleUrls: ['./field-list-modal.component.css']
})
export class FieldListModalComponent implements OnInit {
  isModalToggled = true;
  @Input()
  modalAttributes!: { name: string; modalID: string, pdfURL: string};
  @Output() closeModalEmitter = new EventEmitter<{isModalToggled: boolean}>();

  constructor() { }

  ngOnInit(): void {
  }

  closeModal(){
    this.isModalToggled = false
    this.closeModalEmitter.emit({isModalToggled: this.isModalToggled})
  }

}
