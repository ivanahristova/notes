import { Component } from '@angular/core';
import { NotesService } from 'src/app/services/notes.service';

@Component({
  selector: 'app-notes',
  templateUrl: './notes.component.html',
  styleUrls: ['./notes.component.css']
})
export class NotesComponent {
  form: any = {
    title: null,
    description: null
  };

  constructor(private notesService: NotesService) { }

  ngOnInit(): void {
    this.notesService.getNotes("1").subscribe({
      next: data => {
        console.log(data);
      },
      error: err => {
        // this.errorMessage = err.error.error.charAt(0).toUpperCase()
        //                   + err.error.error.slice(1);
      }
    });
  }

  reloadPage(): void {
    window.location.reload();
  }

}
