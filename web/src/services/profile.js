import axios from "../client";
import {Auth, Status} from "../store/actions";


export default class ProfileService {

  constructor(dispatch) {
    this.dispatch = dispatch
  }

  load() {
    axios.get("/profile").then(resp => {
        this.dispatch({type: Auth.Profile, payload: resp.data})
    }).catch(c => {
      this.dispatch({type: Status.Error, payload: "The profile cannot be loaded: " + c.response.data})
    })
  }

}
