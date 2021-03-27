import axios from "../client";
import {Status, Roles} from "../store/actions";

export default class RoleService {

  constructor(dispatch) {
    this.dispatch = dispatch
  }

  search() {
    this.dispatch({type: Roles.BeginLoad})
    axios.get("/role").then(resp => {
      if(resp.data == null) {
        this.dispatch({type: Roles.EndLoad, payload: []})
      }
      else {
        this.dispatch({type: Roles.EndLoad, payload: resp.data})
      }
    }).catch(c => {
      this.dispatch({type: Roles.EndLoad})
      this.dispatch({type: Status.Error, payload: "The app cannot be loaded: " + c.response.data})
    })
  }

  save(data) {
    return new Promise((resolve, reject) => {
      this.dispatch({type: Roles.BeginSave})
      axios.post("/role", data)
        .then(r => {
          this.dispatch({type: Roles.EndSave})
          this.dispatch({type: Status.Info, payload: "Saving the app " + data.name})
          resolve(r)
        })
        .catch(c => {
          this.dispatch({type: Roles.EndSave})
          this.dispatch({type: Status.Error, payload: "The app cannot be added: " + c.response.data})
          reject(c)
        })
    })
  }


  edit(data) {
    return new Promise((resolve, reject) => {
      this.dispatch({type: Roles.BeginEdit})
      axios.put("/role", data)
        .then(r => {
          this.dispatch({type: Roles.EndEdit})
          this.dispatch({type: Status.Info, payload: "Updating the app " + data.name})
          resolve(r)
        })
        .catch(c => {
          this.dispatch({type: Roles.EndEdit})
          this.dispatch({type: Status.Error, payload: "The app " + data.name + " cannot be edited: " + c.response.data})
          reject(c)
        })
    })
  }

  get(id) {

    this.dispatch({type: Roles.BeginGet})
    axios.get("/role/" + id).then(resp => {
        this.dispatch({type: Roles.EndGet, payload: resp.data})

    }).catch(c => {
      this.dispatch({type: Roles.EndLoad})
      this.dispatch({type: Status.Error, payload: "The app cannot be loaded: " + c.response.data})
    })
  }
}
