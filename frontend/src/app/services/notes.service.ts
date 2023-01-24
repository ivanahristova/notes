import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { Note } from '../dto/note';
import { NotesResponse } from '../dto/response';

const NOTES_API = 'http://localhost:8080/api/auth/';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
  })
};

@Injectable({
  providedIn: 'root'
})
export class NotesService {
  private getNotesUrl: string = 'http://localhost:8080/api/{userId}/notes'

  constructor(private http: HttpClient, private cookieService: CookieService) { }

  public getUserNotes(userId: string): Observable<NotesResponse> {
    const url = this.getNotesUrl.replace('{userId}', userId);
    return this.http.get<NotesResponse>(url, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    });
  }

  public getNote(userId: number, noteId: number): Observable<Note> {
    const url = 'http://localhost:8080/api/' + userId + '/notes/' + noteId;
    return this.http.get<Note>(url, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    });
  }

  createNote(title: string, description: string): Observable<any> {
    return this.http.post(
      NOTES_API + 'new',
      {
        title,
        description
      },
      httpOptions
    );
  }

}
