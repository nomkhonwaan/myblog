import { Component, OnInit, Input, HostBinding, ChangeDetectionStrategy } from '@angular/core';

import { PostTitleComponent } from './post-title.component';

@Component({
  selector: 'app-post-featured-image',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <a [routerLink]="isDisabledLink() ? null : href" [attr.aria-label]="post.title">
      <img class="lazyload" *ngIf="src" [attr.alt]="post.title" [attr.src]="src" />
    </a>
  `,
  styles: [
    `
      :host {
        display: block;
      }
    `,
    `
      img {
        border-radius: .4rem;
        max-width: 100%;
      }
    `
  ],
})
export class PostFeaturedImageComponent extends PostTitleComponent implements OnInit {

  @HostBinding('class.-with-featured-image')
  withFeaturedImage = false;

  @HostBinding('style.background-image')
  src: string;

  ngOnInit(): void {
    super.ngOnInit();

    if (this.hasFeaturedImage()) {
      this.src = `/api/v2.1/storage/${this.post.featuredImage.slug}?width=${this.innerWidth}`;
      this.withFeaturedImage = true;
    }
  }

  hasFeaturedImage(): boolean {
    return !!this.post && this.post.featuredImage.slug.length > 0;
  }

}
