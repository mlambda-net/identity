import React from "react";
import {connect} from "react-redux";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Action, State} from "@mlambda-net/core/packages/common";
import RoleAdd from "./roles/role_add"
import RoleEdit from "./roles/role_edit"
import RoleSearch from "./roles/role_search"

const styles = (theme) => ({

})


class Roles extends React.Component {

  constructor(props) {
    super(props);
    this.state = {actual: "search"}
  }


  render() {
    return (
      <Action actual={this.state.actual}>
        <State name="search">
          <RoleSearch
            onCreate={() => this.setState({actual: "add"})}
            onEdit={() => this.setState({actual: "edit"})}/>
        </State>
        <State name="add">
          <RoleAdd onAdded={() => this.setState({actual: "search"})} onCancel={() => this.setState({actual: "search"})}/>
        </State>
        <State name="edit">
          <RoleEdit onCancel={() => this.setState({actual: "search"})} onEdit={() => this.setState({actual: "search"})}/>
        </State>
      </Action>
    );
  }
}


export default connect(s => s.users)(withStyles(styles)(withTheme(Roles)));
