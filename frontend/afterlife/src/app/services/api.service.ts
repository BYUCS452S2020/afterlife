import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';

import {Event} from './timeline-resolver.service';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(private http: HttpClient) {}

  getTimeline = () => {
    return this.http.get<Event[]>("/api/timeline");
  }

  login = (email: string, password: string) => {
    return this.http.post("/api/login", {
      email: email,
      password: password
    });
  }

  register = (email: string, password: string, first: string, last: string) => {
    return this.http.post("/api/register", {
      email: email,
      password: password,
      firstName: first,
      lastName: last
    });
  }

  logout = () => {
    return this.http.post("/api/logout", {});
  }

  createEvent = (event: Event) => {
    return this.http.post("/api/event", event);
  }

  updateEvent = (event: Event) => {
    return this.http.put("/api/event", event);
  }

  deleteEvent = (id: string) => {
    return this.http.delete("/api/event/" + id);
  }
}
