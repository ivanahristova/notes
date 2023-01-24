import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CookieService } from 'ngx-cookie-service';
import { NotesService } from 'src/app/services/notes.service';
import { Note } from '../../dto/note';

@Component({
  selector: 'app-note',
  templateUrl: './note.component.html',
  styleUrls: ['./note.component.css']
})
export class NoteComponent implements OnInit {
  private userId: number = 0;
  private noteId: number = 0;

  @Input() note?: Note;

  constructor(private notesService: NotesService, private actovatedRoute: ActivatedRoute,
    private router: Router, private cookieService: CookieService) {}

  ngOnInit(): void {
    this.userId = Number(this.actovatedRoute.snapshot.paramMap.get('userId'));
    this.noteId = Number(this.actovatedRoute.snapshot.paramMap.get('noteId'));

    this.notesService.getNote(this.userId, this.noteId).subscribe({
      next: (response: Note)  => {
        console.log(response);
        this.note = response;
      },
      error: (error: any) => {
        console.log(error);
      }
    });
  }

  deleteCookie() {
    this.cookieService.delete('user-jwt');
  }

}
