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

class AppEdit extends React.Component {

  constructor(props) {
    super(props);
    this.actual = this.props.actual
    this.state = ({  description: props.actual.description})
    this.editHandler = this.edit.bind(this)
    this.service = new AppService(props.dispatch)
  }

  edit() {
    this.service.edit({id: this.actual.id, name: this.actual.name, description: this.state.description})
      .then(r => {
        if (this.props.onEdit != null) {
          this.props.onEdit()
        }
      })
  }


  render() {
    const {classes, loading, actual} = this.props

    return (
      <Paper className={classes.root}>
        <div className={classes.card}>
          <div className={classes.title}>
            <Typography color="secondary" variant="h6">
              Edit the app {actual.name}
            </Typography>
          </div>
          <Progress loading={loading}/>
          <div className={classes.box}>
            <Grid container direction="column" spacing={3}>
              <Grid item>
                <TextField disabled label="Name" fullWidth value={actual.name} />
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
              <Button variant="contained" color="primary" onClick={this.editHandler} >
                Save
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

export default connect(s => s.apps)(withStyles(styles)(withTheme(AppEdit)));
