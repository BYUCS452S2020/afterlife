import {Injectable} from '@angular/core';
import {Resolve, ActivatedRouteSnapshot, RouterStateSnapshot, Router} from "@angular/router";
import {Observable, EMPTY} from 'rxjs';
import {catchError, map} from 'rxjs/operators';

import {ApiService} from "./api.service";

export class Event {
}

@Injectable({
  providedIn: 'root'
})
export class TimelineResolver implements Resolve<Event[]> {
  constructor(private api: ApiService, private router: Router) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<Event[]> | Observable<never> {
    return this.api.getTimeline().pipe(
      map(data => data),
      catchError(err => {
        console.warn("unable to get timeline", err);
        this.router.navigate(["/login"])
        return EMPTY;
      })
    )
  }
}
