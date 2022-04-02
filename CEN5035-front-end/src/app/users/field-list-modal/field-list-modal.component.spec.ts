import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FieldListModalComponent } from './field-list-modal.component';

describe('FieldListModalComponent', () => {
  let component: FieldListModalComponent;
  let fixture: ComponentFixture<FieldListModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FieldListModalComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FieldListModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
