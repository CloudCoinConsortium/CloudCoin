import { EventEmitter, Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class EventService {

  sideBarFireEvent: EventEmitter<object> = new EventEmitter();

  constructor() { }

  getSideBarFireEvent(){
    return this.sideBarFireEvent;
  }

  emitSideBarFireEvent(data:any = ''){
    this.sideBarFireEvent.emit(data);
  }

}
