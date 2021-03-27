
import React from "react";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Alert} from "@material-ui/lab";
import {connect} from "react-redux";
import {Snackbar} from "@material-ui/core";
import {Status} from "../store/actions";



const styles = (theme) => ({
  alert: {
    position:"absolute",
    bottom: "1px",
    width: "100%",
    justifyContent: "center",
    display: "flex"
  }

})


class Error extends  React.Component {

  constructor(props) {
    super(props);
    this.closeHandler = this.close.bind(this)
  }

  close(reason) {
      if(reason === "clickaway"){
        return;
      }
      this.props.dispatch({type: Status.Close})
  }

  render() {

    const {classes, open, message, type} = this.props

    return (
      <div className={classes.alert}>
        <Snackbar open={open} autoHideDuration={4000} onClose={this.closeHandler}>
          <Alert onClose={this.closeHandler} variant="filled" severity={type}>
            {message}
          </Alert>
        </Snackbar>
      </div>
    )
  }
}

export default  connect(s => s.status)(withStyles(styles)(withTheme(Error)))
