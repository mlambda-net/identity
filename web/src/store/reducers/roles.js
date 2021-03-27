import {Roles} from "../actions";

const initial = {
  items: [],
  actual: {},
  loading: false,
}


export default function roles (state = initial, action) {
  switch (action.type) {
    case Roles.BeginLoad:
      return {...state, loading: true}
    case Roles.EndLoad:
      return {...state, loading: false, items: action.payload}
    case Roles.BeginSave:
      return {...state, loading: true}
    case Roles.EndSave:
      return {...state, loading: false}
    case Roles.Edit:
      return {...state, actual: action.payload}
    case Roles.BeginEdit:
      return {...state, loading: true}
    case Roles.EndEdit:
      return {...state, loading: false}
    case Roles.BeginGet:
      return {...state, loading: true}
    case Roles.EndGet:
      return {...state, loading: false, actual: action.payload}

    default:
      return state
  }
}
