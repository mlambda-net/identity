import React from "react";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Button, Divider, Grid, Paper, Tab, TextField, Typography} from "@material-ui/core";
import validator from "validator";
import {UserService} from "../../services";
import {connect} from "react-redux";
import Progress from "../../components/progress";
import {TabContext, TabList, TabPanel} from "@material-ui/lab";
import UserRoles from "./user_role"

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

class CreateUser extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "",
      nameError: false,
      nameHelp: "",
      lastName: "",
      lastNameError: false,
      lastNameHelp: "",
      email: "",
      password: "",
      confirm: "",
      emailError: false,
      emailHelp: "",
      passwordError: false,
      passwordHelp: "",
      confirmError: false,
      confirmHelp: "",
      tab: "1",
      roles: [],
    }
    this.service = new UserService(props.dispatch)
    this.raiseHandler = this.raiseCreate.bind(this)
    this.rolesSelected = this.rolesSelected.bind(this)
  }

  validEmail() {
    if (!validator.isEmail(this.state.email)) {

      this.setState({
        emailError: true,
        emailHelp: "the email is invalid"
      })
      return false;
    }
    this.setState({
      emailError: false,
      emailHelp: ""
    })

    return true
  }

  strongPassword() {

    const options = {
      minLength: 8,
      minLowercase: 1,
      minUppercase: 1,
      minNumbers: 1
    }

    if (!validator.isStrongPassword(this.state.password, options)) {
      this.setState({
        passwordError: true,
        passwordHelp: "the password is not strong"
      })
      return false
    } else {
      this.setState({
        passwordError: false,
        passwordHelp: ""
      })
      return true
    }
  }

  matchPassword() {
    if (this.state.password !== this.state.confirm) {
      this.setState({
        passwordError: true,
        passwordHelp: "the password doesn't match with the confirm"
      })
      return false
    } else {
      this.setState({
        passwordError: false,
        passwordHelp: ""
      })
      return true
    }
  }

  hasPassword() {
    if (validator.isEmpty(this.state.password)) {
      this.setState({
        passwordError: true,
        passwordHelp: "the password cannot be empty"
      })
      return false
    } else {
      this.setState({
        passwordError: false,
        passwordHelp: ""
      })
      return true
    }
  }

  hasConfirm() {
    if (validator.isEmpty(this.state.confirm)) {
      this.setState({
        confirmError: true,
        confirmHelp: "the confirm password cannot be empty"
      })
      return false
    } else {
      this.setState({
        confirmError: false,
        confirmHelp: ""
      })
      return true
    }
  }

  validPassword() {

    const hasPassword = this.hasPassword()
    const hasConfirm = this.hasConfirm()

    if (hasPassword && hasConfirm) {
      if (this.strongPassword() && this.matchPassword()) {
        return true
      }
    }
    return false;
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

  validName() {
    const hasName = this.hasName()
    const hasLastName = this.hasLastName()
    return hasName && hasLastName
  }

  isValid() {
    const validName = this.validName()
    const validEmail = this.validEmail()
    const validPassword = this.validPassword()
    return validName && validEmail && validPassword
  }

  raiseCreate() {
    if (this.props.onCreate != null) {
      if (this.isValid()) {

        let user = {
          name: this.state.name,
          lastName: this.state.lastName,
          email: this.state.email,
          password: this.state.password,
          roles: this.state.roles
        }

        this.service.save(user).then(() => {
          if (this.props.onCreate !== null) {
            this.props.onCreate()
          }
        })
      }
    }
  }

  rolesSelected(roles) {
    this.setState({roles: roles})
  }

  render() {
    const {classes, loading} = this.props
    return (
      <Paper className={classes.root}>
        <div className={classes.card}>
          <div className={classes.title}>
            <Typography color="secondary" variant="h6">
              Create User
            </Typography>
          </div>
          <Progress loading={loading}/>
          <div className={classes.box}>

           <TabContext value={this.state.tab}>
             <TabList onChange={(e, value)=> this.setState({tab: value})} textColor="primary" indicatorColor="secondary">
               <Tab label="General" value="1"/>
               <Tab label="Roles" value="2"/>
             </TabList>
             <TabPanel value="1">
               <Grid container direction="column" spacing={3}>
                 <Grid item>
                   <TextField value={this.state.name} error={this.state.nameError} label="Name" helperText={this.state.nameHelp} fullWidth
                              onChange={e => this.setState({name: e.target.value})}/>
                 </Grid>
                 <Grid item>
                   <TextField value={this.state.lastName} error={this.state.lastNameError} label="Last Name" helperText={this.state.lastNameHelp}
                              fullWidth onChange={e => this.setState({lastName: e.target.value})}/>
                 </Grid>
                 <Grid item>
                   <TextField value={this.state.email} error={this.state.emailError} label="Email" helperText={this.state.emailHelp} fullWidth
                              onChange={e => this.setState({email: e.target.value})}/>
                 </Grid>
                 <Grid item>
                   <TextField value={this.state.password} error={this.state.passwordError} label="Password" helperText={this.state.passwordHelp}
                              type="password" fullWidth onChange={e => this.setState({password: e.target.value})}/>
                 </Grid>
                 <Grid item>
                   <TextField value={this.state.confirm} error={this.state.confirmError} label="Confirm Password"
                              helperText={this.state.confirmHelp} type="password" fullWidth
                              onChange={e => this.setState({confirm: e.target.value})}/>
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
              <Button variant="contained" color="primary" onClick={this.raiseHandler}>
                Create
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

export default connect(s => s.users)(withStyles(styles)(withTheme(CreateUser)));
