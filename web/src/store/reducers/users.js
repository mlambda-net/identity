import {Users} from "../actions";


const data = {
  items: [],
  loading: false,
  error: {},
  user: {
    roles: null
  }
}


export default function users (state = data, action) {

  switch (action.type) {
    case Users.BeginLoad:
      return {...state, loading: true}
    case Users.EndLoad:
      return {...state, items: action.payload, loading: false}
    case Users.Edit:
      return {...state, user: action.payload}
    case Users.BeginGet:
      return {...state, loading: true}
    case Users.EndGet:
      return {...state, loading: false, user: action.payload}

    default:
      return state
  }
}
