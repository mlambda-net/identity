

import {Status} from "../actions";


const data = {
  message: "",
  type: "",
  open: false,
}

export default function status (state = data, action) {

  switch (action.type) {
    case Status.Info:
      return  {...state, message: action.payload, type: "info", open:true}
    case Status.Error:
      return  {...state, message: action.payload, type: "error", open:true}
    case Status.Warning:
      return  {...state, message: action.payload, type: "warning", open:true}
    case Status.Close:
      return {...state, open: false}
    default:
      return state
  }
}
