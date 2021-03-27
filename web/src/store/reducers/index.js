import { combineReducers} from "redux";
import auth from "./auth";
import users from "./users";
import status from "./status";
import apps from "./apps";
import roles from "./roles";

export default combineReducers({
  auth, users, status, apps, roles,
})
