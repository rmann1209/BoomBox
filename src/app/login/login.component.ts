import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms'
import { User, signupUser } from '../signup/signup.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  title = 'Boombox';

  constructor(private formBuilder:FormBuilder, private SignupUser:signupUser){}

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
