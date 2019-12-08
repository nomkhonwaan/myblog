import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpEvent, HttpHandler, HttpRequest } from '@angular/common/http';
import { Store, select } from '@ngrx/store';
import { Observable, ObservableInput, of } from 'rxjs';
import { finalize, first, mergeMap } from 'rxjs/operators';

import { isFetching, isNotFetching } from './app.actions';

@Injectable()
export class AppHttpInterceptor implements HttpInterceptor {

  auth$: Observable<{ accessToken: string }>;

  constructor(private store: Store<{ app: AppState }>) {
    this.auth$ = store.pipe(select('app', 'auth'));
  }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    this.store.dispatch(isFetching());

    return this.withAuthorization(req).pipe(
      first(),
      mergeMap((req: HttpRequest<any>): ObservableInput<HttpEvent<any>> => {
        return next.handle(req);
      }),
      finalize(() => { this.store.dispatch(isNotFetching()); }),
    );
  }

  private withAuthorization(req: HttpRequest<any>): Observable<HttpRequest<any>> {
    return this.auth$.pipe(
      first(),
      mergeMap((auth?: { accessToken: string }): ObservableInput<any> => {
        if (auth && auth.accessToken !== '') {
          req = req.clone({
            setHeaders: {
              Authorization: `Bearer ${auth.accessToken}`,
            },
            withCredentials: true,
          });
        }

        return of<HttpRequest<any>>(req);
      }),
    );
  }

}
