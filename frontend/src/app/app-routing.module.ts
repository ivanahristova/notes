import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { LoginGuard } from './guards/login.guard';

import { LoginComponent } from './pages/login/login.component';
import { NewNoteComponent } from './pages/new-note/new-note.component';
import { NotesComponent } from './pages/notes/notes.component';
import { SignupComponent } from './pages/signup/signup.component';

const routes: Routes = [
  {
    path: 'login',
    component: LoginComponent,
    canActivate: [LoginGuard]
  },
  {
    path: 'signup',
    component: SignupComponent
  },
  {
    path: 'notes',
    component: NotesComponent,
    canActivate: [AuthGuard]
  },
  {
    path: 'new-note',
    component: NewNoteComponent,
    canActivate: [AuthGuard]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
