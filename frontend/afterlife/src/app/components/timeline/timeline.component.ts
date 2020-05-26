import {Component, OnInit} from '@angular/core';
import {Router, ActivatedRoute} from "@angular/router";
import {MatDialog} from "@angular/material/dialog";

import {ApiService} from "../../services/api.service";
import {Event} from "../../services/timeline-resolver.service";
import {EventDialogComponent} from "../../dialogs/event-dialog/event-dialog.component";

@Component({
  selector: 'app-timeline',
  templateUrl: './timeline.component.html',
  styleUrls: ['./timeline.component.scss']
})
export class TimelineComponent implements OnInit {
  timeline: Event[] = [];

  constructor(private api: ApiService, private router: Router, private route: ActivatedRoute, private dialog: MatDialog) {}

  ngOnInit() {
    this.timeline = this.route.snapshot.data["timeline"];
    console.log("timeline", this.timeline);
  }

  logout = () => {
    this.api.logout().subscribe(() => {
      this.router.navigate(["/"]);
    }, err => {
      console.warn("unable to logout", err);
    })
  }

  newEvent = () => {
    const dialogRef = this.dialog.open(EventDialogComponent, {
      width: "66vw",
      data: {
        event: new Event()
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      window.location.reload();
    })
  }

  editEvent = (event: Event) => {
    const dialogRef = this.dialog.open(EventDialogComponent, {
      width: "66vw",
      data: {
        event: event,
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      window.location.reload();
    })
  }
}
