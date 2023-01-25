import { Component } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { AuthService } from 'src/app/services/auth.service';
import { NotesService } from 'src/app/services/notes.service';

@Component({
  selector: 'app-new-note',
  templateUrl: './new-note.component.html',
  styleUrls: ['./new-note.component.css']
})
export class NewNoteComponent {
  form: any = {
    title: null,
    description: null
  };
  isSuccessful = false;
  errorMessage = '';

  constructor(private notesService: NotesService, private authService: AuthService,
              private cookieService: CookieService) { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    const { title, description } = this.form;

    this.notesService.addNote(title, description).subscribe({
      next: note => {
        console.log(note);
        this.isSuccessful = true;
      },
      error: err => {
        this.errorMessage = err.error.message;
      }
    });
  }

  deleteCookie() {
    this.cookieService.delete('user-jwt');
  }

}
