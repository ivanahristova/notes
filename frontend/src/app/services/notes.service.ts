import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

const NOTES_API = 'http://localhost:8080/api/notes';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
  })
};

@Injectable({
  providedIn: 'root'
})
export class NotesService {

  constructor(private http: HttpClient) { }

  getNotes(userId: string): Observable<any> {
    return this.http.get(
      NOTES_API + '/all'
    );
  }

  addNote(userId: string): Observable<any> {
    return this.http.post(
      NOTES_API + '/new',
      {
        userId,
      },
      httpOptions
    );
  }
}
