import { Component, OnInit } from '@angular/core';
import { PosService, Dish } from './pos.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [PosService]
})
export class AppComponent implements OnInit {
  menu: Dish[];
  title = 'ui';
  view = 0;

  constructor(private posService: PosService) { }

  toggleFocus(event) {
    if (event.srcElement.classList.contains("transaction")) {
      event.srcElement.nextElementSibling.classList.remove("active");
      event.srcElement.classList.add("active");
      this.view = 0;
    } else if (event.srcElement.classList.contains("menu")) {
      event.srcElement.previousElementSibling.classList.remove("active");
      event.srcElement.classList.add("active");
      this.view = 1;
    }
  }

  ngOnInit() {
    this.menu = [];
    this.getAll();
    // this.interval = setInterval(() => {
    //   this.getAll();
    // }, 1000);
  }

  getAll() {
    this.posService.getMenu().subscribe((data: Dish[]) => {
      this.menu = data;
    })
  }
}
