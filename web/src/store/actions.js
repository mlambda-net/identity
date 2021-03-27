
const Status = {
  Error: "STATUS_ERROR",
  Info : "STATUS_INFO",
  Warning: "STATUS_WARNING",
  Close: "STATUS_CLOSE",
}

const  Auth = {
  SetAuth: "SET_AUTH",
  Profile: "SET_PROFILE",
  Logout: "LOGOUT"
}

const Users = {
  BeginLoad: "USER_LOAD",
  EndLoad: "USER_END_LOAD",
  BeginSave: "USER_BEGIN_SAVE",
  EndSave: "USER_END_SAVE",
  Edit: "USER_EDIT",
  BeginGet: "USER_BEGIN_GET",
  EndGet: "USER_END_GET"

}

const Apps = {
  BeginSave: "APP_BEGIN_SAVE",
  EndSave: "APP_END_SAVE",
  BeginLoad: "APP_BEGIN_LOAD",
  EndLoad: "APP_END_LOAD",
  Edit: "APP_EDIT",
  BeginEdit: "APP_BEGIN_EDIT",
  EndEdit: "APP_EDIT_END",
  BeginGet: "APP_BEGIN_GET",
  EndGet: "APP_END_GET"
}

const Roles = {
  BeginSave: "ROLES_BEGIN_SAVE",
  EndSave: "ROLES_END_SAVE",
  BeginLoad: "ROLES_BEGIN_LOAD",
  EndLoad: "ROLES_END_LOAD",
  Edit: "ROLES_EDIT",
  BeginEdit: "ROLES_BEGIN_EDIT",
  EndEdit: "ROLES_EDIT_END",
  BeginGet: "ROLES_BEGIN_GET",
  EndGet: "ROLES_END_GET",
}

export {
  Auth, Users,Status, Apps, Roles
}
