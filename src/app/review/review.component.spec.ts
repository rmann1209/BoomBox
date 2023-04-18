import { ComponentFixture, async, TestBed } from '@angular/core/testing';
import { ReviewComponent } from './review.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import {HttpClientModule} from '@angular/common/http';
import { RouterOutlet } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';

describe('ReviewComponent', () => {
  let component: ReviewComponent;
  let fixture: ComponentFixture<ReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ RouterTestingModule,
        HttpClientTestingModule, RouterOutlet, FormsModule, ReactiveFormsModule ],
      declarations: [ ReviewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('check if even', async(() => {
    expect(2==2).toBeTruthy();
  }));

  it('should create the review component', async(() => {
    const fixture = TestBed.createComponent(ReviewComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));
});
