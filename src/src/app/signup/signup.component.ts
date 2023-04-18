import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms'
import { User, signupUser } from './signup.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent {
  constructor(private SignupUser:signupUser, private formBuilder:FormBuilder){}

  accountForm = this.formBuilder.group({
    Username:[''],
    Password:['']
  })

  addUser(username: string, password: string) : void {
    this.SignupUser.addUser({username, password} as User)
    .subscribe((response: any) => {
      console.log(response);
    });
  }
}
