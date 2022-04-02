import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-menu-list-modal',
  templateUrl: './menu-list-modal.component.html',
  styleUrls: ['./menu-list-modal.component.css']
})
export class MenuListModalComponent implements OnInit {
  isModalToggled = true;
  @Input()
  modalAttributes!: { name: string; modalID: string};
  @Output() closeMenuListModalEmitter = new EventEmitter<{isModalToggled: boolean}>();

  constructor() { }

  ngOnInit(): void {
  }

  closeModal(){
    this.isModalToggled = false
    this.closeMenuListModalEmitter.emit({isModalToggled: this.isModalToggled})
  }

}
