import {getWeather} from 'lib/networking/api';

export default class Meteor {
  constructor() {
    this.q = 'London,uk';
    this.a = undefined;
    this.id = undefined;
    this.lat = undefined;
    this.lon = undefined;
    this.zip = undefined;

    this._myposition = undefined;
    this._watchID = navigator.geolocation.watchPosition(
      position => {
        this._myposition = {
          latitude: position.coords.latitude,
          longitude: position.coords.longitude,
          latitudeDelta: 0.00922 * 1.5,
          longitudeDelta: 0.00421 * 1.5,
        };
      },
      error => console.log(error),
    );
  }

  _clearPos() {
    this.q = undefined;
    this.a = undefined;
    this.id = undefined;
    this.lat = undefined;
    this.lon = undefined;
    this.zip = undefined;
  }

  getMyWeather() {
    this._clearPos();
    navigator.geolocation.getCurrentPosition(
      position => {
        this._myposition = {
          latitude: position.coords.latitude,
          longitude: position.coords.longitude,
          latitudeDelta: 0.00922 * 1.5,
          longitudeDelta: 0.00421 * 1.5,
        };

        this.lat = this._myposition.latitude;
        this.lon = this._myposition.longitude;

        this.getWeather();
      },
      error => console.log(error),
      {enableHighAccuracy: true, timeout: 20000, maximumAge: 0},
    );
  }

  async getWeather() {
    let data = await getWeather(this);
    console.log(data);
  }
}
