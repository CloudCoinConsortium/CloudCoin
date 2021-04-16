import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DebitcardEntryComponent } from './debitcard-entry.component';

describe('DebitcardEntryComponent', () => {
  let component: DebitcardEntryComponent;
  let fixture: ComponentFixture<DebitcardEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DebitcardEntryComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DebitcardEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
