import React from "react";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Action, State} from "@mlambda-net/core/packages/common";
import CreateUser from "./user/create_user"
import UserList from "./user/user_list"
import EditUser from "./user/edit_user"
import {connect} from "react-redux";




const styles = (theme) => ({

})


class User extends React.Component {
  constructor(props) {
    super(props);
    this.state = {actual: "items"}
    this.createUserHandler = this.createUser.bind(this)
    this.cancelCreateHandler = this.cancelUser.bind(this)
    this.AddUserHandler = this.created.bind(this)
    this.editUserHandler = this.editUser.bind(this)
    this.editHandler = this.edited.bind(this)
  }

  cancelUser() {
    this.setState({actual: "items"})
  }

  createUser() {
    this.setState({actual: "add"})
  }

  created() {
    this.setState({actual: "items"})
  }

  editUser() {
    this.setState({actual: "edit"})
  }

  edited() {
    this.setState({actual: "items"})
  }

  render() {
    return (
      <Action actual={this.state.actual}>
        <State name="items">
          <UserList onCreate={this.createUserHandler} onEdit={this.editUserHandler}/>
        </State>
        <State name="add">
          <CreateUser onCreate={this.AddUserHandler} onCancel={this.cancelCreateHandler}/>
        </State>
        <State name="edit">
          <EditUser onCancel={this.cancelCreateHandler} onEdit={this.editHandler}/>
        </State>
      </Action>
    )
  }
}

export default connect(s => s.users)(withStyles(styles)(withTheme(User)));
