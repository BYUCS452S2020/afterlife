import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";

import {ApiService} from "../../services/api.service";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  hide = true;
  email = "";
  password = "";

  constructor(private api: ApiService, private router: Router) {}

  ngOnInit() {}

  login = () => {
    this.api.login(this.email, this.password).subscribe(() => {
      console.log("redirecting to timeline");
      this.router.navigate(["/timeline"]);
    }, err => {
      console.warn("failed to login", err);
    });
  }

  register = () => {
    window.location.pathname = "/register"
  }
}
