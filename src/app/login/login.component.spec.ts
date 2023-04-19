import { ComponentFixture, async, TestBed } from '@angular/core/testing';
import { LoginComponent } from './login.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import {HttpClientModule} from '@angular/common/http';
import { RouterOutlet } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ RouterTestingModule,
        HttpClientTestingModule, RouterOutlet, FormsModule, ReactiveFormsModule ],
      declarations: [ LoginComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create the login component', async(() => {
    const fixture = TestBed.createComponent(LoginComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));

  it('form should be invalid initially', () => {
    expect(component.accountForm.valid).toBeFalsy();
  });
  
  it('form should be valid when fields are filled out', () => {
    component.accountForm.controls['Username'].setValue('testuser');
    component.accountForm.controls['Password'].setValue('testpassword');
    expect(component.accountForm.valid).toBeTruthy();
  });
  
  it('loginUser function should be called on button click', () => {
    spyOn(component, 'loginUser');
    const button = fixture.debugElement.nativeElement.querySelector('#submit');
    button.click();
    expect(component.loginUser).toHaveBeenCalled();
  });
});

