import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NotificationDrafterComponent } from './notification-drafter.component';

describe('NotificationDrafterComponent', () => {
  let component: NotificationDrafterComponent;
  let fixture: ComponentFixture<NotificationDrafterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NotificationDrafterComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NotificationDrafterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
