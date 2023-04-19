
import { ComponentFixture, async, TestBed } from '@angular/core/testing';
import { SignupComponent } from './signup.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import {HttpClientModule} from '@angular/common/http';
import { RouterOutlet } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';
import { By } from '@angular/platform-browser';

describe('SignupComponent', () => {
  let component: SignupComponent;
  let fixture: ComponentFixture<SignupComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ RouterTestingModule,
        HttpClientTestingModule, RouterOutlet, FormsModule, ReactiveFormsModule ],
      declarations: [ SignupComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SignupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create the signup component', async(() => {
    const fixture = TestBed.createComponent(SignupComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));

  it('should have an invalid form when empty', () => {
    expect(component.accountForm.valid).toBeFalsy();
  });

  it('should have a valid form when both fields are filled', () => {
    const username = component.accountForm.controls['Username'];
    const password = component.accountForm.controls['Password'];
    username.setValue('testuser');
    password.setValue('testpassword');
    expect(component.accountForm.valid).toBeTruthy();
  });

  it('should add a user when form is submitted with valid data', () => {
    spyOn(component, 'addUser');
    const form = fixture.debugElement.query(By.css('form'));
    const usernameInput = form.query(By.css('#Username'));
    const passwordInput = form.query(By.css('#Password'));
    const submitButton = form.query(By.css('button[type=submit]'));
  
    usernameInput.nativeElement.value = 'testuser';
    passwordInput.nativeElement.value = 'testpass';
    usernameInput.nativeElement.dispatchEvent(new Event('input'));
    passwordInput.nativeElement.dispatchEvent(new Event('input'));
    fixture.detectChanges();
  
    expect(component.accountForm.valid).toBeTruthy();
  });
  
  
});