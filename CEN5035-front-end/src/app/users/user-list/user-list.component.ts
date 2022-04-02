import { Component, OnInit } from '@angular/core';
import { UserService } from 'src/app/user.service';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  showModal = false
  showMenuModal = false
  modalData = {name: "",modalID: ""}
  fieldListElements = [{name: 'Contracts'},{name: 'Messages'},{name: 'Notifications'}]
  fieldElements = [{name: 'Contract'}, {name: 'Attorney'}]

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.userService.getUsers().subscribe(data => {
      console.log(data)
    })
  }

  onModalToggle(eventData: {isModalToggled: boolean, name: string, modalID: string}){
    console.log("name: " + eventData.name)
    console.log("modalID: " + eventData.modalID)
    this.modalData.name = eventData.name
    this.modalData.modalID = eventData.modalID
    this.showModal = eventData.isModalToggled
  }

  onMenuModalToggle(eventData: {isModalToggled: boolean}){
    console.log("eventData.isModalToggled: " + eventData.isModalToggled)
    this.showMenuModal = eventData.isModalToggled
  }

  onModalClose(eventData: {isModalToggled: boolean}){
    this.showModal = eventData.isModalToggled
  }

}
