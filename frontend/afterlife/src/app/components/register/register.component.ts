import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";

import {ApiService} from "../../services/api.service";

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {
  hide = true;
  email = "";
  password = "";
  firstName = "";
  lastName = "";

  constructor(private api: ApiService, private router: Router) {}

  ngOnInit() {}

  register = () => {
    this.api.register(this.email, this.password, this.firstName, this.lastName).subscribe(() => {
      this.router.navigate(["/login"]);
    }, err => {
      console.warn("failed to login", err);
    });
  }
}
