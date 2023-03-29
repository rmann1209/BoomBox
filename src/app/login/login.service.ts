import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { Inject, Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';


const httpPost = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  })
}


export interface User {
  username: string;
  password: string;
}

@Injectable({
  providedIn: 'root'
})
export class loginUser {

  baseUrl: string = 'http://localhost:8080';
  constructor(private http : HttpClient){}


  loginUser(user: User): Observable<User> {

    const loginURL:string = this.baseUrl + '/login';

    console.log("loginUser: " + user.username + ' ' + user.password + ' '+ loginURL);

    return this.http.post<User>(loginURL, user, httpPost);
  }


  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      console.error('An error occurred:', error.error);
    }
    else {
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error);
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error('Error Occurred. Try Again Later.'));
  }

}
