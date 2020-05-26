import {Component, Inject} from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from "@angular/material/dialog";
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import {MatChipInputEvent} from '@angular/material/chips';

import {ApiService} from "../../services/api.service";
import {Event} from "../../services/timeline-resolver.service";

@Component({
  selector: 'app-event-dialog',
  templateUrl: './event-dialog.component.html',
  styleUrls: ['./event-dialog.component.scss']
})
export class EventDialogComponent {

  readonly today = new Date();
  readonly separatorKeysCodes: number[] = [ENTER, COMMA];

  constructor(private api: ApiService, public dialogRef: MatDialogRef<EventDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: {
      event: Event;
    }) {
  }

  cancel = () => {
    this.dialogRef.close();
  }

  save = () => {
    console.log("this.event", this.data.event);

    this.api.createEvent(this.data.event).subscribe(() => {
      this.dialogRef.close();
      window.location.reload();
    }, err => {
      console.warn("failed to save event", err);
    });
  }

  addTo = (event: MatChipInputEvent) => {
    const input = event.input;
    const value = event.value;

    if ((value || '').trim()) {
      this.data.event.email.to.push(value.trim())
    }

    if (input) {
      input.value = '';
    }
  }

  removeTo = (addr: string) => {
    const idx = this.data.event.email.to.indexOf(addr);

    if (idx >= 0) {
      this.data.event.email.to.splice(idx, 1);
    }
  }
}
