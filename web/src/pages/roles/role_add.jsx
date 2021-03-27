import React from "react";
import {connect} from "react-redux";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {
  Button,
  Divider,
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Paper,
  Select,
  TextField,
  Typography
} from "@material-ui/core";
import Progress from "../../components/progress";
import AppService from "../../services/apps";
import {RoleService} from "../../services";


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
    margin: theme.spacing(5, 10)
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

class RoleAdd extends React.Component {

  constructor(props) {
    super(props);
    this.state ={name:"",nameError:false, nameHelp:"",   app:"",appError:false, description:""}
    this.addHandler = this.add.bind(this)
    this.appService = new AppService(props.dispatch)
    this.roleService = new RoleService(props.dispatch)
  }

  componentDidMount() {
    this.appService.search()
  }

  add() {
    if (this.state.name === "") {
      this.setState({nameError: true, nameHelp: "the name is required"})
    }
    if (this.state.app === "") {
      this.setState({appError: true})
    }

    if (this.state.name !== "" && this.state.app !== "") {
      this.roleService.save({name: this.state.name, app: this.state.app, description: this.state.description})
        .then(r => {
          if (this.props.onAdded !== null) {
            this.props.onAdded()
          }
        })
    }
  }


  render() {
    const {apps, roles, classes} = this.props

    return (
      <Paper className={classes.root}>
        <div className={classes.card}>
          <div className={classes.title}>
            <Typography color="secondary" variant="h6">
              Add a new role
            </Typography>
          </div>
          <Progress loading={roles.loading}/>
          <div className={classes.box}>
            <Grid container direction="column" spacing={3}>
              <Grid item>
                <TextField label="Name"  fullWidth  required
                           error={this.state.nameError} helperText={this.state.nameHelp}
                           onChange={e => this.setState({name: e.target.value})}/>
              </Grid>
              <Grid item>
                <FormControl fullWidth>
                  <InputLabel id="app-select" error={this.state.appError} >App</InputLabel>
                  <Select error={this.state.nameError} fullWidth labelId="app-select" value={this.state.app}
                          onChange={e => this.setState({app: e.target.value})}>
                    {
                      apps.items.map((app) => <MenuItem value={app.id}><em>{app.name}</em></MenuItem>)
                    }
                  </Select>

                </FormControl>
              </Grid>
              <Grid item>
                <TextField label="Description" fullWidth value={this.state.description}
                           onChange={e => this.setState({description: e.target.value})}/>
              </Grid>

            </Grid>
          </div>
          <Divider/>
          <div className={classes.actions}>
            <div className={classes.action}>
              <Button variant="contained" color="primary" onClick={this.addHandler}>
                Add
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

export default connect(s => { return {roles: s.roles, apps: s.apps} })(withStyles(styles)(withTheme(RoleAdd)));
