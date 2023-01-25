import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CookieService } from 'ngx-cookie-service';

import { Observable } from 'rxjs';

const AUTH_API = 'http://localhost:8080/api/auth/';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
  })
};

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient, private cookieService: CookieService) { }

  signup(email: string, username: string, password: string): Observable<any> {
    return this.http.post(AUTH_API + 'signup', { email, username, password }, httpOptions);
  }

  login(username: string, password: string): Observable<any> {
    return this.http.post(AUTH_API + 'login', { username, password, }, httpOptions);
  }

  public isAuthenticated(): boolean {
    const token = this.cookieService.get('user-jwt');
    if (token === null) {
      return false;
    }
    const jwtHelper = new JwtHelperService();
    return !jwtHelper.isTokenExpired(token);
  }
}
