import React from "react";
import {connect} from "react-redux";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Action, State} from "@mlambda-net/core/packages/common";
import AddApp from "./apps/app_add"
import AppSearch from "./apps/app_search"
import AppEdit from "./apps/app_edit"

const styles = (theme) => ({

})

class Apps extends React.Component {

  constructor(props) {
    super(props);
    this.state = {actual: "search"}
  }

  render() {
    return (
      <Action actual={this.state.actual}>
        <State name="search">
          <AppSearch
            onCreate={() => this.setState({actual: "add"})}
            onEdit={() => this.setState({actual: "edit"})}/>
        </State>
        <State name="add">
          <AddApp onAdded={() => this.setState({actual: "search"})} onCancel={() => this.setState({actual: "search"})}/>
        </State>
        <State name="edit">
          <AppEdit onCancel={() => this.setState({actual: "search"})} onEdit={() => this.setState({actual: "search"})}/>
        </State>
      </Action>
    )
  }
}

export default connect(s => s.apps)(withStyles(styles)(withTheme(Apps)));

