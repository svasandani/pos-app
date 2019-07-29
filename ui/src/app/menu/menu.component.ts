import { Component, OnInit } from '@angular/core';
import { PosService, Dish } from '../pos.service';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.css'],
  inputs: ['menu']
})
export class MenuComponent {
  menu = this.menu;
  interval: any;
}
