import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ContractDrafterComponent } from './contract-drafter.component';

describe('ContractDrafterComponent', () => {
  let component: ContractDrafterComponent;
  let fixture: ComponentFixture<ContractDrafterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ContractDrafterComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ContractDrafterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
