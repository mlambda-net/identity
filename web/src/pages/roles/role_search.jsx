import React from "react";
import {connect} from "react-redux";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Paper, Typography} from "@material-ui/core";
import Progress from "../../components/progress";
import {DataGrid} from "@material-ui/data-grid";
import ToolBar from "../../components/toolbar";
import {RoleService} from "../../services";

const styles = (theme) => ({

  root: {
    width: '100%'
  },
  card: {

  },
  title: {
    margin: "20px 0",
    background: theme.primary
  },
  body: {
  },
  action: {
    display: 'flex',
    justifyContent: 'flex-end'
  },
  cardTitle: {
    margin: '20px'
  },
  cardBox: {
    width: 'calc(100% - 11px)',
    height: '50vh',
    margin: '20px 10px',
  }

})

class RoleSearch extends React.Component {
  constructor(props) {
    super(props);
    this.selectionHandler = this.selection.bind(this)
    this.editHandler = this.edit.bind(this)
    this.service = new RoleService(props.dispatch)
    this.state = {canEdit: false, edit: {}}
    this.columns = [
      {field: "name", headerName: "Name", width: 200},
      {field: "appName", headerName: "App", width: 200},
      {field: "description", headerName: "Description", width: 600},
    ]
  }



  componentDidMount() {
    this.service.search()
  }

  edit() {
    if(this.props.onEdit != null) {
      this.props.onEdit()
    }
  }

  selection(item) {
    this.service.get(item.selectionModel[0])
    this.setState({canEdit:true})
  }

  render() {
    const {classes, items, loading} = this.props
    return (
      <Paper className={classes.root}>
        <div className={classes.card}>
          <div className={classes.cardTitle}>
            <Typography color="secondary" variant="h6">
              Role List
            </Typography>
          </div>
          <Progress loading={loading}/>
          <div className={classes.cardBox}>
            <DataGrid getRowId={(r) => r.id}
                      density="true"
                      loading={loading}
                      nonce
                      columns={this.columns}
                      onSelectionModelChange={this.selectionHandler}
                      rows={items}
                      pageSize={20}
                      components={{
                        Toolbar: () => ( <ToolBar create={this.props.onCreate} canEdit={this.state.canEdit} edit={this.editHandler} delete={this.deleteHandler} />),
                      }} columnBuffer={20}/>
          </div>
        </div>
      </Paper>
    )}
}

export default connect(s => s.roles)(withStyles(styles)(withTheme(RoleSearch)));
