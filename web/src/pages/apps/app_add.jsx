import React from "react";
import {connect} from "react-redux";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Button, Divider, Grid, Paper, TextField, Typography} from "@material-ui/core";
import Progress from "../../components/progress";
import AppService from "../../services/apps";


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

class AppAdd extends React.Component {

  constructor(props) {
    super(props);
    this.state = ({name: "", nameError: false, nameHelp: "", description: ""})
    this.addHandler = this.add.bind(this)
    this.service = new AppService(props.dispatch)
  }

  validName() {
    if( this.state.name === "") {
      this.setState({nameError: true, nameHelp: "the name is required"})
      return false
    } else {
      this.setState({nameError: false, nameHelp: ""})
      return true
    }

  }

  add() {
    if (this.validName()) {
      this.service.save({name: this.state.name, description: this.state.description})
        .then(r => {
          if (this.props.onAdded !== null) {
            this.props.onAdded()
          }
        })
    }
  }


  render() {
    const {classes, loading} = this.props

    return (
      <Paper className={classes.root}>
        <div className={classes.card}>
          <div className={classes.title}>
            <Typography color="secondary" variant="h6">
              Add a new app
            </Typography>
          </div>
          <Progress loading={loading}/>
          <div className={classes.box}>
            <Grid container direction="column" spacing={3}>
              <Grid item>
                <TextField label="Name" fullWidth required
                           error={this.state.nameError} helperText={this.state.nameHelp}
                           onChange={e => this.setState({name: e.target.value})}/>
              </Grid>
              <Grid item>
                <TextField label="Description" fullWidth
                           onChange={e => this.setState({description: e.target.value})}/>
              </Grid>


            </Grid>
          </div>
          <Divider/>
          <div className={classes.actions}>
            <div className={classes.action}>
              <Button variant="contained" color="primary" onClick={this.addHandler} >
                Add
              </Button>
            </div>
            <div className={classes.action}>
              <Button variant="contained" color="primary" onClick={this.props.onCancel} >
                Cancel
              </Button>
            </div>
          </div>
        </div>
      </Paper>
    )
  }
}

export default connect(s => s.apps)(withStyles(styles)(withTheme(AppAdd)));
