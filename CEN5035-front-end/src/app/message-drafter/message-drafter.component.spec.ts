import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MessageDrafterComponent } from './message-drafter.component';

describe('MessageDrafterComponent', () => {
  let component: MessageDrafterComponent;
  let fixture: ComponentFixture<MessageDrafterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MessageDrafterComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MessageDrafterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
