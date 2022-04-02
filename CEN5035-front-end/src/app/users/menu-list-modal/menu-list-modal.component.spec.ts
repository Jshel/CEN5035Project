import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MenuListModalComponent } from './menu-list-modal.component';

describe('MenuListModalComponent', () => {
  let component: MenuListModalComponent;
  let fixture: ComponentFixture<MenuListModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MenuListModalComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MenuListModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
