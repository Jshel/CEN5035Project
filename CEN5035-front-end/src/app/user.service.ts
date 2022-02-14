import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  apiUrl: string = 'https://api.github.com/users'

  constructor(private http: HttpClient){}
  
  //get all users

  getUsers(){
    return this.http.get(`${this.apiUrl}?per_page=10`)
  }

  getUser(username: string){
    return this.http.get(`${this.apiUrl}/${username}`)
  }
}
