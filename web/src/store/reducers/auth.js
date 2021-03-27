import {Auth} from "../actions";


const data = {
  data: null,
  user: {
    name:"",
    email:""
  }
}

export default function auth (state = data, action) {
  switch (action.type) {
    case Auth.SetAuth:
      return  {...state, data: action.payload}
    case Auth.Profile:
      return {...state, user: action.payload}
    case Auth.Logout:
      state.data.manager.removeUser().then(c=> {
        window.location.href = "/"
      })
      return {...state}
    default:
      return state
  }
}
