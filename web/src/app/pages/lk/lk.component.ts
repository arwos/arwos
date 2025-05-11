import {Component, inject, OnInit} from '@angular/core';
import {RequestService} from '@onega-ui/core';

@Component({
  selector: 'app-lk',
  imports: [],
  templateUrl: './lk.component.html',
  standalone: true,
  styleUrl: './lk.component.scss'
})
export class LkComponent implements OnInit {

  private readonly req = inject(RequestService)

  ngOnInit(): void {
    // this.req.post('/auth//pam/sign-in', {hello: 1}).subscribe(value => console.log(value))
    this.req.post('/auth/logout', {hello: 1}).subscribe(value => console.log(value))
  }


}

