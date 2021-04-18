import React from "react";
import {withStyles, withTheme} from "@material-ui/core/styles";
import { Paper, TextField, Typography} from "@material-ui/core";
import {DataGrid} from '@material-ui/data-grid';
import {UserService} from "../../services";
import {connect} from "react-redux";
import Progress from "../../components/progress";
import ToolBar from "../../components/toolbar";
import Box from "@material-ui/core/Box";
import SearchAdorn from "../../components/search";
import {theme} from "../../theme";

const styles = (theme) => ({
  root: {
    width: '100%'
  },
  card: {

  },
  title: {
    margin: "20px 0"
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
    margin: '20px 10px',
    [theme.breakpoints.up('xs')]: {
      height: '300px'
    },

    [theme.breakpoints.up('md')]: {
      height: '565px'
    },

  },
  filter: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center'
  }

})



class UserList extends React.Component {
  constructor(props) {
    super(props);
    this.service = new UserService(props.dispatch)
    this.state = {loading: false, canEdit: false, edit: {}, filter: ""}
    this.columns = [
      {field: "name", headerName: "Name", width: 200},
      {field: "lastName", headerName: "Last Name", width: 200},
      {field: "email", headerName: "Email", width: 300},
      {field: "active", headerName: "Active", width: 100},
    ]
    this.searchChangeHandler =  this.searchChange.bind(this)
    this.selectionHandler = this.selection.bind(this)
    this.editHandler = this.edit.bind(this)
    this.searchHandler = this.search.bind(this)
    this.keyPress = this.keyPress.bind(this)
  }

  keyPress(e) {
    if(e.keyCode === 13){
      this.service.fetch(this.state.filter)
    }
  }

  selection(item) {
    this.setState({canEdit: true})
    this.service.get(item.selectionModel[0])
  }

  edit() {
    if (this.props.onEdit !== null) {
      this.props.onEdit()
    }
  }

  componentDidMount() {
    this.service.fetch("")
  }

  search() {
    this.service.fetch(this.state.filter)
  }

  searchChange(e) {
    this.setState({filter: e.target.value})
  }

  render() {
    const {classes, items, loading} = this.props
    return (
      <Paper className={classes.root} elevation={3}>
        <div className={classes.card}>
          <div className={classes.cardTitle}>

            <Box className={classes.filter}>
              <Typography color="secondary" variant="h6">
                User List
              </Typography>
              <TextField label="search" variant="outlined" size="small" onKeyDown={this.keyPress}
                         InputProps={{endAdornment: <SearchAdorn onClick={this.searchHandler}/>}}
                         onChange={this.searchChangeHandler} />
            </Box>

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
                        Toolbar: () => (<ToolBar create={this.props.onCreate}
                                                 canEdit={this.state.canEdit}
                                                 edit={this.editHandler}
                                                 search={this.state.filter}
                                                 onSearch={this.searchHandler}
                        />),
                      }}/>
          </div>
        </div>
      </Paper>
    )
  }
}



export default connect(s => s.users)(withStyles(styles)(withTheme(UserList)));
