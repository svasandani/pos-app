import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class PosService {
  constructor(private httpClient: HttpClient) { }

  // login http
  login(user, id) {
    return this.httpClient.post(environment.gateway + '/login?user=' + user + '&id=' + id, "");
  }

  // menu http
  getMenu() {
    return this.httpClient.get(environment.gateway + '/menu');
  }

  addMenuDish(server, dish: Dish) {
    return this.httpClient.post(environment.gateway + '/menu?server=' + server, dish);
  }

  deleteMenuDish(server, sku) {
    return this.httpClient.delete(environment.gateway + '/menu?server=' + server + '&sku=' + sku);
  }

  editMenuDish(server, dish: Dish) {
    return this.httpClient.put(environment.gateway + '/menu?server=' + server, dish);
  }

  // transaction http
  getAllTransactions() {
    return this.httpClient.get(environment.gateway + '/transaction');
  }

  getNewTransaction(server) {
    return this.httpClient.get(environment.gateway + '/transaction/new?server=' + server);
  }

  addDishTransaction(id, sku) {
    return this.httpClient.post(environment.gateway + '/transaction?id=' + id + '&sku=' + sku, "");
  }

  deleteDishTransaction(id, index) {
    return this.httpClient.delete(environment.gateway + '/transaction?id=' + id + '&index=' + index);
  }

  toggleDishServe(id, index) {
    return this.httpClient.put(environment.gateway + '/transaction?id=' + id + '&index=' + index, "");
  }

  payTransaction(id, method) {
    return this.httpClient.post(environment.gateway + '/transaction/pay?id=' + id + '&method=' + method, "");
  }

  // server http
  addServer(adder, server: Server) {
    return this.httpClient.post(environment.gateway + '/server?server=' + adder, server);
  }

  deleteServer(deleter, deleted) {
    return this.httpClient.delete(environment.gateway + '/server?server=' + deleter + '&object=' + deleted);
  }
}

export class Dish {
  sku: number;
  name: string;
  price: string;
}

export class Server {
  id: string;
  name: string;
  date: string;
}
