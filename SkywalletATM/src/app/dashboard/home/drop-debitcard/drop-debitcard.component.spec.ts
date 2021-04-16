import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DropDebitcardComponent } from './drop-debitcard.component';

describe('DropDebitcardComponent', () => {
  let component: DropDebitcardComponent;
  let fixture: ComponentFixture<DropDebitcardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DropDebitcardComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DropDebitcardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
