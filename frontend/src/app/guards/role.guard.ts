import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';
import { JwtHelperService } from '@auth0/angular-jwt';
import { AuthService } from '../services/auth.service';


@Injectable({
  providedIn: 'root'
})
export class RoleGuard implements CanActivate {

  constructor(private cookieService: CookieService, private router: Router, private auth: AuthService) {}

  canActivate(route: ActivatedRouteSnapshot): boolean {
    const expectedRole = route.data['expectedRole'];
    const token = this.cookieService.get('user-jwt');
    if (token === null) {
      return false;
    }
    const jwtService: JwtHelperService = new JwtHelperService();
    const role: boolean = jwtService.decodeToken(token)['admin'];
    if (expectedRole === 'ADMIN' && !role) {
      if (this.auth.isAuthenticated()) {
        this.router.navigate(['notes'])
      } else {
        this.router.navigate(['login']);
      }
      return false;
    }
    return true;
  }

}
