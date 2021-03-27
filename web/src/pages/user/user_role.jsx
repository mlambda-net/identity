import React from "react";
import {connect} from "react-redux";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {RoleService} from "../../services";
import {Checkbox, List, ListItem, ListItemIcon,  Typography} from "@material-ui/core";

const styles = (theme) => ({
  role : {
    display: "flex",
    flexDirection: "column"
  },
  name: {
  },
  app : {
  }
})

class UserRoles extends React.Component {
  constructor(props) {
    super(props);
    this.service = new RoleService(props.dispatch)
    if(this.props.roles != null) {
      this.state = {selected: props.roles}
    }else {
      this.state = {selected:[]}
    }
    this.select = this.select.bind(this)
  }

  select(value) {
    if(!this.exists(value)) {
      this.state.selected.push(value)
    } else {
      this.state.selected.splice(this.state.selected.indexOf(value), 1)
    }
    this.setState(this.state.selected)
    this.props.onSelect(this.state.selected)
  }

  componentDidMount() {
    this.service.search("")
  }

  exists(id) {
    return this.state.selected.indexOf(id) > -1
  }

  isChecked(id) {
    return this.exists(id)
  }

  render() {
    const {items, classes} = this.props
    return (
      <List>
        {items.map(role =>
          <ListItem key={role.id} dense button onClick={ _ => this.select(role.id)}>
            <ListItemIcon>
              <Checkbox edge="start" checked={this.isChecked(role.id)}/>
            </ListItemIcon>
            <div className={classes.role}>
              <Typography variant="subtitle2">{role.name}</Typography>
              <Typography variant="caption" color="textSecondary">{role.appName}</Typography>
            </div>
          </ListItem>
        )}
      </List>
    )
  }

}


export default connect(s=> s.roles)(withStyles(styles)(withTheme(UserRoles)));
