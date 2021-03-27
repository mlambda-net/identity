import React from "react";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Button, Divider, Grid, Paper, Tab, TextField, Typography} from "@material-ui/core";
import validator from "validator";
import {UserService} from "../../services";
import {connect} from "react-redux";
import Progress from "../../components/progress";
import {TabContext, TabList, TabPanel} from "@material-ui/lab";
import UserRoles from "./user_role";


const styles = (theme) => ({

  root: {
    width: '100%'
  },
  card: {

  },
  title: {
    margin:  theme.spacing(2, 3)
  },
  box: {
    margin: theme.spacing(5, 10),
    minHeight: '450px'
  },
  actions: {
    margin:  theme.spacing(1),
    padding: theme.spacing(1,2),
    display: 'flex',
  },
  action: {
    padding: theme.spacing(1),
  }
})

class EditUser extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.user.name,
      nameError: false,
      nameHelp: "",
      lastName: props.user.lastName,
      lastNameError: false,
      lastNameHelp: "",
      tab: "1",
      roles: props.user.roles
    }
    this.service = new UserService(props.dispatch)
    this.editHandler = this.editUser.bind(this)
    this.rolesSelected = this.rolesSelected.bind(this)
  }


  hasName() {
    if (validator.isEmpty(this.state.name)) {
      this.setState({
        nameError: true,
        nameHelp: "the name is required"
      })
      return false;
    }
    this.setState({
      nameError: false,
      nameHelp: ""
    })

    return true
  }

  hasLastName() {
    if (validator.isEmpty(this.state.lastName)) {

      this.setState({
        lastNameError: true,
        lastNameHelp: "the name is required"
      })
      return false;
    }
    this.setState({
      lastNameError: false,
      lastNameHelp: ""
    })

    return true
  }

  isValid() {
    const hasName = this.hasName()
    const hasLastName = this.hasLastName()
    return hasName && hasLastName
  }

  rolesSelected(roles) {
    this.setState({roles: roles})
  }

  editUser() {
    if (this.props.onEdit != null) {
      if (this.isValid()) {
        let user = {
          name: this.state.name,
          lastName: this.state.lastName,
          id: this.props.user.id,
          roles: this.state.roles
        }
        this.service.update(user).then(() => {
          if (this.props.onEdit !== null) {
            this.props.onEdit()
          }
        }).catch(e => {
          console.log(e)
        })
      }
    }
  }

  render() {

    const {classes, loading, user} = this.props

    return (
      <Paper className={classes.root}>
        <div className={classes.card}>
          <div className={classes.title}>
            <Typography color="secondary" variant="h6">
              Edit User
            </Typography>
          </div>
          <Progress loading={loading}/>
          <div className={classes.box}>
            <TabContext value={this.state.tab}>
              <TabList onChange={(e, value) => this.setState({tab: value})} textColor="primary"
                       indicatorColor="secondary">
                <Tab label="General" value="1"/>
                <Tab label="Roles" value="2"/>
              </TabList>
              <TabPanel value="1">

                <Grid container direction="column" spacing={3}>
                  <Grid item>
                    <TextField value={this.state.name} error={this.state.nameError} label="Name"
                               helperText={this.state.nameHelp} fullWidth
                               onChange={e => this.setState({name: e.target.value})}>
                      {user.name}
                    </TextField>
                  </Grid>
                  <Grid item>
                    <TextField value={this.state.lastName} error={this.state.nameError} label="Last Name"
                               helperText={this.state.lastNameHelp} fullWidth
                               onChange={e => this.setState({lastName: e.target.value})}/>
                  </Grid>
                </Grid>
              </TabPanel>
              <TabPanel value="2">
                <UserRoles roles={this.state.roles} onSelect={this.rolesSelected}/>
              </TabPanel>
            </TabContext>
          </div>
          <Divider/>
          <div className={classes.actions}>
            <div className={classes.action}>
              <Button variant="contained" color="primary" onClick={this.editHandler}>
                Edit
              </Button>
            </div>
            <div className={classes.action}>
              <Button variant="contained" color="primary" onClick={this.props.onCancel}>
                Cancel
              </Button>
            </div>
          </div>
        </div>
      </Paper>

    )
  }
}

export default connect(s => s.users)(withStyles(styles)(withTheme(EditUser)));
