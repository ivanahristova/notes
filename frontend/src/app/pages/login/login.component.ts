import { Component } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  form: any = {
    username: null,
    password: null
  };
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  username = '';

  constructor(private authService: AuthService, private cookieService: CookieService) { }

  ngOnInit(): void {

  }

  onSubmit(): void {
    const { username, password } = this.form;

    this.authService.login(username, password).subscribe({
      next: data => {
        console.log(data);
        this.cookieService.set('user-jwt', data.data);

        this.isLoginFailed = false;
        this.isLoggedIn = true;
        this.username = username;
      },
      error: err => {
        this.errorMessage = err.error.error.charAt(0).toUpperCase() + err.error.error.slice(1);
        this.isLoginFailed = true;
      }
    });
  }

  reloadPage(): void {
    window.location.reload();
  }
}
