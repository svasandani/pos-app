import { Component, OnInit } from '@angular/core';
import { PosService, Dish } from '../pos.service';

@Component({
  selector: 'app-transaction',
  templateUrl: './transaction.component.html',
  styleUrls: ['./transaction.component.css'],
  providers: [PosService],
  inputs: ['menu']
})
export class TransactionComponent implements OnInit {
  transactions: any[];
  dishlist: Dish[][] = [];
  menu = this.menu;

  constructor(private posService: PosService) { }

  ngOnInit() {
    this.transactions = [];
    this.getAll();
    // this.interval = setInterval(() => {
    //   this.getAll();
    // }, 1000);
  }

  getAll() {
    this.posService.getAllTransactions().subscribe((data: any[]) => {
      this.transactions = data;
      this.transactions.forEach((item, index) => {
        this.dishlist[index] = [];
        item.skulist.forEach((sku, index2) => {
          this.dishlist[index][index2] = this.menu[sku - 1];
        })
      })
    })
  }

}
