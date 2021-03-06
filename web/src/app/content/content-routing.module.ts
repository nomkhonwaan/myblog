import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from '../auth';
import { ArchiveComponent } from './archive';
import { ContentComponent } from './content.component';
import { MyPostsComponent } from './my-posts';
import { PageNotFoundComponent } from './page-not-found';
import { RecentPostsComponent } from './recent-posts';
import { SingleComponent } from './single';

const routes: Routes = [
  {
    path: '', component: ContentComponent,
    children: [
      {
        path: '',
        pathMatch: 'full',
        component: RecentPostsComponent
      },
      {
        path: 'my-posts',
        pathMatch: 'full',
        canActivate: [AuthGuard],
        component: MyPostsComponent,
      },
      {
        path: ':page',
        pathMatch: 'full',
        component: ArchiveComponent,
        data: { from: 'all' },
      },
      {
        path: 'category/:slug',
        pathMatch: 'full',
        component: ArchiveComponent,
        data: { from: 'category' },
      },
      {
        path: 'category/:slug/:page',
        pathMatch: 'full',
        component: ArchiveComponent,
        data: { from: 'category' },
      },
      {
        path: 'tag/:slug',
        pathMatch: 'full',
        component: ArchiveComponent,
        data: { from: 'tag' },
      },
      {
        path: 'tag/:slug/:page',
        pathMatch: 'full',
        component: ArchiveComponent,
        data: { from: 'tag' },
      },
      {
        path: ':year/:month/:date/:slug',
        pathMatch: 'full',
        component: SingleComponent,
      },
      {
        path: '**',
        component: PageNotFoundComponent,
      },
    ],
  },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes),
  ],
  exports: [
    RouterModule,
  ],
})
export class ContentRoutingModule { }
