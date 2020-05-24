import {Component, OnInit} from '@angular/core';
import {Router, ActivatedRoute} from "@angular/router";

import {ApiService} from "../../services/api.service";

@Component({
  selector: 'app-timeline',
  templateUrl: './timeline.component.html',
  styleUrls: ['./timeline.component.scss']
})
export class TimelineComponent implements OnInit {

  constructor(private api: ApiService, private router: Router, private route: ActivatedRoute) {}

  ngOnInit() {
    console.log("timeline", this.route.snapshot.data['timeline']);
  }

  logout = () => {
    this.api.logout().subscribe(() => {
      this.router.navigate(["/"]);
    }, err => {
      console.warn("unable to logout", err);
    })
  }
}
