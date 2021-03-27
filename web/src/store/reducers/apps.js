import {Apps} from "../actions";

const initial = {
  items: [],
  actual: {},
  loading: false,
}


export default function apps (state = initial, action) {
  switch (action.type) {
    case Apps.BeginLoad:
      return {...state, loading: true}
    case Apps.EndLoad:
      return {...state, loading: false, items: action.payload}
    case Apps.BeginSave:
      return {...state, loading: true}
    case Apps.EndSave:
      return {...state, loading: false}
    case Apps.Edit:
      return {...state, actual: action.payload}
    case Apps.BeginEdit:
      return {...state, loading: true}
    case Apps.EndEdit:
      return {...state, loading: false}
    case Apps.BeginGet:
      return {...state, loading: true}
    case Apps.EndGet:
      return {...state, loading: false, actual: action.payload}

    default:
      return state
  }
}
