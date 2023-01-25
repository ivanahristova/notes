import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { Note } from '../dto/note';
import { NotesResponse } from '../dto/response';

const USER_API = 'http://localhost:8080/api/user';
const GET_USER_NOTES = USER_API + '/notes';
const ADD_NOTE = USER_API + '/notes';
const GET_NOTE = USER_API + '/notes/{noteId}';
const DELETE_NOTE = USER_API + '/notes/{noteId}';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
  })
};

@Injectable({
  providedIn: 'root'
})
export class NotesService {
  constructor(private http: HttpClient, private cookieService: CookieService) { }

  public getUserNotes(): Observable<NotesResponse> {
    return this.http.get<NotesResponse>(GET_USER_NOTES, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    });
  }

  public getNote(noteId: number): Observable<Note> {
    const url = GET_NOTE.replace('{noteId}', noteId.toString());
    return this.http.get<Note>(url, {
      headers: { 'Authorization': 'Bearer ' + this.cookieService.get('user-jwt') }
    });
  }

  public addNote(title: string, description: string): Observable<any> {
    return this.http.post(ADD_NOTE, {title, description}, httpOptions);
  }

  public deleteNote(noteId: number): Observable<any> {
    const url = DELETE_NOTE.replace('{noteId}', noteId.toString());
    return this.http.post(url, httpOptions);
  }

}
