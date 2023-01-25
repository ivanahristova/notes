import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
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
  public notesResponse = new NotesResponse();
  public notes: Note[] = [];
  private userId: number = 0;

  selectedNote?: Note;

  constructor(private notesService: NotesService, private activatedRoute: ActivatedRoute,
              private router: Router, private cookieService: CookieService) {}

  ngOnInit(): void {
    this.getUserNotes();
  }

  getUserNotes() {
    // const jwtService = new JwtHelperService();
    // const userId: string = jwtService.decodeToken(this.cookieService.get("user-jwt"))['user_id'];
    this.notesService.getUserNotes().subscribe({
      next: (response: NotesResponse) => {
        this.notes = response.notes
      },
      error: (error: any) => {
        console.log(error);
      }
    });
  }

  routeToAddNote() {
    this.router.navigate(['/new-note']);
  }

  routeToUpdateNote(noteId: string) {
    // let  idString: string = noteId.toString;
    this.router.navigate(['/notes/{noteId}']);
  }

  routeToNotes() {
    this.router.navigate(['/notes']);
  }

  deleteCookie() {
    this.cookieService.delete('user-jwt');
  }
}
