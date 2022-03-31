import { Component, EventEmitter, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-menu-list',
  templateUrl: './menu-list.component.html',
  styleUrls: ['./menu-list.component.css']
})
export class MenuListComponent implements OnInit {
  isModalToggled = true;
  @Output("toggleMenuModel") toggleModal = new EventEmitter<{isModalToggled: boolean}>();

  constructor() { }

  ngOnInit(): void {
  }

  showModal(){
    this.isModalToggled = true
    this.toggleModal.emit({isModalToggled: this.isModalToggled})
  }

}
