import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BrowseContractsComponent } from './browse-contracts.component';

describe('BrowseContractsComponent', () => {
  let component: BrowseContractsComponent;
  let fixture: ComponentFixture<BrowseContractsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BrowseContractsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BrowseContractsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
