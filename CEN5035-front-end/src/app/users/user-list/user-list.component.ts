import { Component, OnInit } from '@angular/core';
import { UserService } from 'src/app/user.service';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  showModal = false
  modalData = {name: "",modalID: ""}
  fieldListElements = [{name: 'Contracts'},{name: 'Messages'}]
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

  onModalClose(eventData: {isModalToggled: boolean}){
    this.showModal = eventData.isModalToggled
  }

}
