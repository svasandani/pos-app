import { TestBed } from '@angular/core/testing';

import { PosService } from './pos.service';

describe('PosService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: PosService = TestBed.get(PosService);
    expect(service).toBeTruthy();
  });
});
