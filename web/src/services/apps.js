import axios from "../client";
import {Status, Apps} from "../store/actions";

export default class AppService {

  constructor(dispatch) {
    this.dispatch = dispatch
  }


  search() {
    this.dispatch({type: Apps.BeginLoad})
    axios.get("/app").then(resp => {
      if(resp.data == null) {
        this.dispatch({type: Apps.EndLoad, payload: []})
      }
      else {
        this.dispatch({type: Apps.EndLoad, payload: resp.data})
      }
    }).catch(c => {
      this.dispatch({type: Apps.EndLoad })
      this.dispatch({type: Status.Error, payload: "The app cannot be loaded: " + c.response.data})
    })
  }

  save(data) {
    return new Promise((resolve, reject) => {
      this.dispatch({type: Apps.BeginSave})
      axios.post("/app", data)
        .then(r => {
          this.dispatch({type: Apps.EndSave})
          this.dispatch({type: Status.Info, payload: "Saving the app " + data.name})
          resolve(r)
        })
        .catch(c => {
          this.dispatch({type: Apps.EndSave})
          this.dispatch({type: Status.Error, payload: "The app cannot be added: " + c.response.data})
          reject(c)
        })
    })
  }


  edit(data) {
    return new Promise((resolve, reject) => {
      this.dispatch({type: Apps.BeginEdit})
      axios.put("/app", data)
        .then(r => {
          this.dispatch({type: Apps.EndEdit})
          this.dispatch({type: Status.Info, payload: "Updating the app " + data.name})
          resolve(r)
        })
        .catch(c => {
          this.dispatch({type: Apps.EndEdit})
          this.dispatch({type: Status.Error, payload: "The app " + data.name + " cannot be edited: " + c.response.data})
          reject(c)
        })
    })
  }

  get(id) {
    this.dispatch({type: Apps.BeginGet})
    axios.get("/app/" + id).then(resp => {
            this.dispatch({type: Apps.EndGet, payload: resp.data})
    }).catch(c => {
      this.dispatch({type: Apps.EndGet})
      this.dispatch({type: Status.Error, payload: "The app cannot be loaded: " + c.response.data})
    })
  }
}
