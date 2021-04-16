import { TestBed } from '@angular/core/testing';

import { RaidaServiceService } from './raida-service.service';

describe('RaidaServiceService', () => {
  let service: RaidaServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RaidaServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
