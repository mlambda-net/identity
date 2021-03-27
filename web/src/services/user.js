import axios from "../client";
import {Status, Users} from "../store/actions";


export default class UserService {
  constructor(dispatch) {
    this.dispatch = dispatch
  }

  save(data) {
    return new Promise((resolve, reject) => {
      this.dispatch({type: Users.BeginSave})
      axios.post("/user", data).then(r => {
        this.dispatch({type: Status.Info, payload: "saving the user"})
        this.dispatch({type: Users.EndSave})
        resolve(r)
      }).catch(c => {
        reject(c)
        this.dispatch({type: Users.EndSave})
        this.dispatch({type: Status.Error, payload: "The user cannot be saved: " + c.response.data})
      })
    })
  }

  update(data) {
    return new Promise((resolve, reject) => {
      this.dispatch({type: Users.BeginSave})
      axios.put("/user", data).then(r => {
        this.dispatch({type: Status.Info, payload: "updating the user"})
        this.dispatch({type: Users.EndSave})
        resolve(r)
      }).catch(c => {
        reject(c)
        this.dispatch({type: Users.EndSave})
        this.dispatch({type: Status.Error, payload: "The user cannot be updated: " + c.response.data})
      })
    })
  }

  fetch(filter) {
    this.dispatch({type: Users.BeginLoad})
    const params = new URLSearchParams([['filter', filter]]);
    axios.get("/user", {params}).then(resp => {
      if (resp.data !== null) {
        this.dispatch({type: Users.EndLoad, payload: resp.data})
      } else {
        this.dispatch({type: Users.EndLoad, payload: []})
      }
    }).catch(c => {
      this.dispatch({type: Users.EndLoad })
      this.dispatch({type: Status.Error, payload: "The users cannot be loaded: " + c.response.data})
    })
  }

  get(id) {
    this.dispatch({type: Users.BeginGet})
    axios.get("/user/" + id).then(resp => {
      this.dispatch({type: Users.EndGet, payload: resp.data})
    }).catch(c => {
      this.dispatch({type: Users.EndGet})
      this.dispatch({type: Status.Error, payload: "The users cannot be loaded: " + c.response.data})
    })
  }
}


