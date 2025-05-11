import {Routes} from '@angular/router';
import {PageNotFoundComponent} from './pages/page-not-found/page-not-found.component';

export const routes: Routes = [
  {path: 'auth', title: 'Auth', loadComponent: () => import('./pages/auth/auth.component').then(c => c.AuthComponent)},
  {path: 'home', title: 'Home', loadComponent: () => import('./pages/main/main.component').then(c => c.MainComponent)},
  {path: 'lk', title: 'Profile', loadComponent: () => import('./pages/lk/lk.component').then(c => c.LkComponent)},
  //@arwos:routes@
  {path: '', redirectTo: '/home', pathMatch: 'full'},
  {path: '**', component: PageNotFoundComponent},
];
