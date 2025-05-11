import {Component, inject, OnInit} from '@angular/core';
import {RequestService} from '@onega-ui/core';

@Component({
  selector: 'app-auth',
  imports: [],
  templateUrl: './auth.component.html',
  standalone: true,
  styleUrl: './auth.component.scss'
})
export class AuthComponent implements OnInit {

  private readonly req = inject(RequestService)

  ngOnInit(): void {
    this.req.post('/auth//pam/sign-in', {hello: 1}).subscribe(value => console.log(value))
    // this.req.post('/auth/logout', {hello: 1}).subscribe(value => console.log(value))
  }


}
