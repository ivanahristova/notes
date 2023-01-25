import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import { Note } from '../dto/note';
import { NotesResponse } from '../dto/response';

const NOTES_API = 'http://localhost:8080/api/notes';
const GET_USER_NOTES = NOTES_API + '/';
const ADD_NOTE = NOTES_API + '/';
const GET_NOTE = NOTES_API + '/{noteId}';
const UPDATE_NOTE = NOTES_API + '/{noteId}';
const DELETE_NOTE = NOTES_API + '/notes/{noteId}';

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

  public updateNote(title: string, description: string): Observable<any> {
    return this.http.put(UPDATE_NOTE, {title, description}, httpOptions);
  }

  public deleteNote(noteId: number): Observable<any> {
    const url = DELETE_NOTE.replace('{noteId}', noteId.toString());
    return this.http.post(url, httpOptions);
  }

}
