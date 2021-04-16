import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RaidamailComponent } from './raidamail.component';

describe('RaidamailComponent', () => {
  let component: RaidamailComponent;
  let fixture: ComponentFixture<RaidamailComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RaidamailComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RaidamailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
