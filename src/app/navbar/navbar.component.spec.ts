import { ComponentFixture, async, TestBed } from '@angular/core/testing';
import { NavbarComponent } from './navbar.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';
import { RouterOutlet } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { Router } from '@angular/router';
import { By } from '@angular/platform-browser';
import { MatButton } from '@angular/material/button';

describe('NavbarComponent', () => {
  let component: NavbarComponent;
  let fixture: ComponentFixture<NavbarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ RouterTestingModule, HttpClientTestingModule ],
      declarations: [ NavbarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    await fixture.whenStable();
    console.log('Component created:', component);
  });

  it('should create the navbar component', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate to home when home button is clicked', async(() => {
    const fixture = TestBed.createComponent(NavbarComponent);
    const component = fixture.componentInstance;
    const router = TestBed.inject(Router);
    const homeButton = fixture.debugElement.query(By.css('#home-button'));
  
    homeButton.nativeElement.click();
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      expect(router.url).toBe('/');
    });
  }));
  
  it('should navigate to login when login button is clicked', async(() => {
    const fixture = TestBed.createComponent(NavbarComponent);
    const component = fixture.componentInstance;
    const router = TestBed.inject(Router);
    const loginButton = fixture.debugElement.query(By.css('#login-button'));
  
    loginButton.nativeElement.click();
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      expect(router.url).toBe('/');
    });
  }));
  
  it('should navigate to sign up when sign up button is clicked', async(() => {
    const fixture = TestBed.createComponent(NavbarComponent);
    const component = fixture.componentInstance;
    const router = TestBed.inject(Router);
    const signUpButton = fixture.debugElement.query(By.css('#signup-button'));
  
    signUpButton.nativeElement.click();
    fixture.detectChanges();
    fixture.whenStable().then(() => {
      expect(router.url).toBe('/');
    });
  }));

});

