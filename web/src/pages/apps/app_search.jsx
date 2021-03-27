import React from "react";
import {withStyles, withTheme} from "@material-ui/core/styles";
import { Paper, Typography} from "@material-ui/core";
import {DataGrid} from '@material-ui/data-grid';
import {connect} from "react-redux";
import Progress from "../../components/progress";
import AppService from "../../services/apps";
import ToolBar from "../../components/toolbar";


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

class AppSearch extends React.Component {
  constructor(props) {
    super(props);
    this.service = new AppService(props.dispatch)
    this.state = {loading: false, canEdit: false, edit: {}}
    this.columns = [
      {field: "name", headerName: "Name", width: 200},
      {field: "description", headerName: "Description", width: 600},
    ]
    this.selectionHandler = this.selection.bind(this)
    this.editHandler = this.edit.bind(this)
    this.deleteHandler = this.delete.bind(this)
  }

  selection(item) {
    this.setState({canEdit: true})
    this.service.get(item.selectionModel[0])
  }

  edit() {
    if(this.props.onEdit != null) {
      this.props.onEdit()
    }
  }

  delete() {
  }

  componentDidMount() {
    this.service.search()
  }

  render() {

    const {classes, items, loading} = this.props
    return (
      <Paper className={classes.root}>
        <div className={classes.card}>
          <div className={classes.cardTitle}>
            <Typography color="secondary" variant="h6">
              App List
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
    )
  }
}



export default connect(s => s.apps)(withStyles(styles)(withTheme(AppSearch)));
