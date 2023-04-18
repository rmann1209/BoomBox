import { Component } from '@angular/core';

@Component({
  selector: 'app-review',
  templateUrl: './review.component.html',
  styleUrls: ['./review.component.css']
})
export class ReviewComponent {
  searchQuery: string;
  showReview: boolean;
  review: string;
  selectedStar: number;
  hoveredStar: number;
  stars: number[];

  constructor() {
    this.searchQuery = '';
    this.showReview = false;
    this.review = '';
    this.selectedStar = 0;
    this.hoveredStar = 0;
    this.stars = [1, 2, 3, 4, 5];
  }

  search(): void {
    // your search logic here
    this.showReview = true;
  }

  selectStar(star: number): void {
    this.selectedStar = star;
  }

  submitReview(): void {
    // your review submission logic here
    this.review = '';
    this.selectedStar = 0;
    this.hoveredStar = 0;
    this.showReview = false;
  }
}
