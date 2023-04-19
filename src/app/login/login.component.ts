import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { User, loginUser } from './login.service';
import { Router } from '@angular/router';


var UserName = null

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  title = 'Boombox';

  constructor(private formBuilder:FormBuilder, private LoginUser:loginUser, private router:Router){}

  accountForm = this.formBuilder.group({
    Username:[''],
    Password:['']
  })

  loginUser(username: string, password: string) : void {
    this.LoginUser.loginUser({username, password} as User)
    .subscribe((response: any) => {
      if (response == null){
        UserName = username
        console.log(UserName)
        this.router.navigate(['/review']);
       //TODO clear UserName variable when logging out?
      }
    });
  }

}