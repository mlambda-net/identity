import {GridToolbarContainer} from "@material-ui/data-grid";
import {Button} from "@material-ui/core";
import {Add, Create, Delete} from "@material-ui/icons";
import React from "react";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";

const styles = (theme) => ({
  toolbar: {
    display: 'flex',
  },
})

class ToolBar extends React.Component {
  render() {
    const {classes} = this.props
    return (
      <GridToolbarContainer className={classes.toolbar}>
        <Button color="primary" onClick={this.props.create}> <Add/>Create</Button>
        <Button color="primary" onClick={this.props.edit} disabled={!this.props.canEdit}> <Create/>Edit</Button>
        <Button color="primary" onClick={this.props.delete} disabled={!this.props.canEdit}><Delete/> Delete</Button>
      </GridToolbarContainer>
    )
  }
}

export default withUtils(styles)(ToolBar)
