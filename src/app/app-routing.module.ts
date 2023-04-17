import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { AppComponent } from './app.component';
import { Router, Routes } from '@angular/router';
import { RouterModule } from '@angular/router';
import { ProfileComponent } from './profile/profile.component';

const appRoutes: Routes = [
  {
  path : 'home',
  component: HomeComponent
  },

  {
  path : 'login',
  component: LoginComponent
  },

  {
  path : 'signup',
  component: SignupComponent
  },
  {
    path: 'profile',
    component: ProfileComponent
  }

];

export default appRoutes;

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule
  ],
  exports:[]
})
export class AppRoutingModule { }
