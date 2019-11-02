import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { PostAuthorComponent } from './post-author.component';
import { PostCategoriesComponent } from './post-categories.component';
import { PostContentComponent, PostContentDirective } from './post-content.component';
import { PostExcerptComponent } from './post-excerpt.component';
import { PostMetadataComponent } from './post-metadata.component';
import { PostShareToComponent } from './post-share-to.component';
import { PostTagsComponent } from './post-tags.component';
import { PostTitleComponent } from './post-title.component';

import { TemplateModule } from '../template/template.module';

@NgModule({
  imports: [
    CommonModule,
    FontAwesomeModule,
    RouterModule,
    TemplateModule,
  ],
  declarations: [
    PostAuthorComponent,
    PostCategoriesComponent,
    PostContentComponent,
    PostExcerptComponent,
    PostMetadataComponent,
    PostShareToComponent,
    PostTagsComponent,
    PostTitleComponent,
    PostContentDirective,
  ],
  exports: [
    PostAuthorComponent,
    PostCategoriesComponent,
    PostContentComponent,
    PostExcerptComponent,
    PostMetadataComponent,
    PostShareToComponent,
    PostTagsComponent,
    PostTitleComponent,
  ],
})
export class PostModule { }
