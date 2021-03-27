
import User from "./pages/user";
import Roles from "./pages/roles"
import Apps from "./pages/apps"

const routes = [
  {
    name: 'global',
    actual: 'list_user',
    routes: [
      {path:"list_user", name: "list_user", component: User},
      {path:"list_roles", name: "list_roles", component: Roles},
      {path: "list_app", name: "list_app", component: Apps}
    ]
  }

]

export default routes
