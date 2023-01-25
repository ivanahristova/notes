import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CookieService } from 'ngx-cookie-service';
import { Note } from 'src/app/dto/note';
import { NotesResponse } from 'src/app/dto/response';
import { NotesService } from 'src/app/services/notes.service';

@Component({
  selector: 'app-notes',
  templateUrl: './notes.component.html',
  styleUrls: ['./notes.component.css']
})
export class NotesComponent implements OnInit {
  public notes: NotesResponse = new NotesResponse();
  private userId: number = 0;

  selectedNote?: Note;

  constructor(private notesService: NotesService, private activatedRoute: ActivatedRoute,
              private router: Router, private cookieService: CookieService) {}

  ngOnInit(): void {
    // this.userId = Number(this.activatedRoute.snapshot.paramMap.get('userId'));
    this.notesService.getUserNotes().subscribe({
      next: (response: NotesResponse) => {
        console.log(response);
        this.notes = response;
      },
      error: error => {
        console.log(error);
      }
    });
  }

  getUserNotes() {
    const jwtService = new JwtHelperService();
    // const userId: string = jwtService.decodeToken(this.cookieService.get("user-jwt"))['user_id'];
    this.notesService.getUserNotes().subscribe({
      next: (response: NotesResponse) => {
        this.notes = response
      },
      error: (error: any) => {
        console.log(error);
      }
    });
  }

  updateNote(noteId: string) {
    // let  idString: string = noteId.toString;
    this.router.navigate([`/notes/${this.userId}/update/${noteId}`]);
  }

  deleteCookie() {
    this.cookieService.delete('user-jwt');
  }
}
