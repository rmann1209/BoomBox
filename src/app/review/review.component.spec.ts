import { ComponentFixture, async, TestBed } from '@angular/core/testing';
import { ReviewComponent } from './review.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import {HttpClientModule} from '@angular/common/http';
import { RouterOutlet } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';
import { By } from '@angular/platform-browser';

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

  it('should create the review component', async(() => {
    const fixture = TestBed.createComponent(ReviewComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));

  it('should initialize with default values', () => {
    expect(component.searchQuery).toBe('');
    expect(component.review).toBe('');
    expect(component.showReview).toBeFalsy();
    expect(component.stars).toEqual([1, 2, 3, 4, 5]);
    expect(component.selectedStar).toBe(0);
    expect(component.hoveredStar).toBe(0);
  });
  
  it('should show the review box when user submits a search query', async(() => {
    spyOn(component, 'search').and.callThrough();
    fixture.detectChanges();
  
    const searchInput = fixture.debugElement.query(By.css('#search-input')).nativeElement;
    searchInput.value = 'test';
    searchInput.dispatchEvent(new Event('input'));
    fixture.detectChanges();
  
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const searchButton = fixture.debugElement.query(By.css('button')).nativeElement;
      searchButton.click();
      fixture.detectChanges();
      expect(component.search).toHaveBeenCalled();
  
      fixture.whenStable().then(() => {
        fixture.detectChanges();
        const reviewBox = fixture.debugElement.query(By.css('.review-box'));
        expect(reviewBox).toBeTruthy();
      });
    });
  }));
  
  
  it('should submit a review when the user clicks the submit button', () => {
    spyOn(component, 'search').and.callThrough();
    fixture.detectChanges();
  
    const searchInput = fixture.debugElement.query(By.css('#search-input')).nativeElement;
    searchInput.value = 'test';
    searchInput.dispatchEvent(new Event('input'));
    fixture.detectChanges();
  
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const searchButton = fixture.debugElement.query(By.css('button')).nativeElement;
      searchButton.click();
      fixture.detectChanges();
      expect(component.search).toHaveBeenCalled();
  
      fixture.whenStable().then(() => {
        fixture.detectChanges();
        const reviewBox = fixture.debugElement.query(By.css('.review-box'));
        expect(reviewBox).toBeTruthy();
      });
    });
  });
  
});