import React from 'react';
import {
  AppBar, Avatar,
  Button, Divider, Drawer,
  IconButton,
  List, ListItem, Paper, Popover,
  Toolbar,
  Typography
} from "@material-ui/core";
import routes from "../routes";
import MenuIcon from "@material-ui/icons/Menu";
import {theme} from "../theme";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import {Display, RouteProvider} from "@mlambda-net/core/packages/routes";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";
import clsx from "clsx";
import {Apps, Done,  Person} from "@material-ui/icons";
import Error from "../components/error";
import {connect} from "react-redux";
import {ProfileService} from "../services";
import {Auth} from "../store/actions";

const styles = (themes) => ({

  app: {
    height: '100%'
  },

  title: {
    flexGrow: "1",
  },

  drawer: {
    width: '240px',
    flexShrink: 0,
  },
  paper: {
    width: '240px'
  },

  root: {
    flexGrow: "1",
    width: '100%',
    height: '100%'
  },

  region: {
    margin: 'auto',
    width: '70%',
    height: 'calc(100% - 65px)',
    alignItems: 'center',
  },

  component: {
    display: 'flex',
    justifyContent: 'center',
    alignItems:'center',
    width:'100%',
    height: '100%',
    background: 'rgba(220,220,220,0.5)'
  },

  appBar: {
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },

  appBarShift: {
    width: `calc(100% - 240px)`,
    marginLeft: 240,
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },

  menu: {
    display: 'flex',
    alignItems: 'center',
    margin: '8px',
    justifyContent:'space-between'
  },

  menuItem: {
    display: 'flex',
    alignItems: 'center',
  },

  menuIcon: {
    margin: "0 5px",
  },

  avatar: {
    backgroundColor: theme.palette.primary.dark
  },

  profile:{
    width: "250px",
    height: "200px",
    display: "flex",
    flexDirection: "column",
    padding: "10px",
    alignItems:"center"
  },
  profileBody: {
    flexGrow: "2",
    display:"flex",
    justifyContent: "center",
    flexDirection: "column",
    height:"100%",
    padding: "5px"
  },
  profileActions: {
    flexGrow: "1"
  }

});

class Admin extends React.Component {

  constructor(props) {
    super(props);
    this.service = new ProfileService(props.dispatch)
    this.state = {
      open: false,
      profileOpen: false,
      profileEl: null
    }

    this.logoutHandler = this.logout.bind(this)
    this.openDrawerHandler = this.openDrawer.bind(this)
    this.closeDrawerHandler = this.closeDrawer.bind(this)
  }


  componentDidMount() {
    this.service.load()
  }

  logout() {
    this.props.dispatch({type: Auth.Logout })
  }


  openDrawer() {
    this.setState({open: true})
  }

  closeDrawer() {
    this.setState({open: false})
  }

  render() {

    const {classes, route, user} = this.props;
    const {open, profileOpen, profileEl} = this.state

    return (
      <RouteProvider routes={routes}>
        <div className={classes.app}>
          <AppBar position="static" className={clsx(classes.appBar, {[classes.appBarShift]: open})}>
            <Toolbar>
              <IconButton edge="start" color="inherit" onClick={this.openDrawerHandler}>
                <MenuIcon/>
              </IconButton>
              <Typography variant="h6" className={classes.title}>
                Identity
              </Typography>
              <div>
                <Button aria-describedby="profile" color="inherit"
                        onClick={(e) => this.setState({profileOpen: true, profileEl: e.currentTarget})}>
                  <Avatar className={classes.avatar}>
                    <Person color="inherit"/>
                  </Avatar>
                  <Typography variant="caption" style={{padding: "10px"}}>
                    {user.email}
                  </Typography>
                </Button>
                <Popover
                  id="profile"
                  anchorEl={profileEl}
                  onClose={() => this.setState({ profileOpen: false})}
                  anchorOrigin={{vertical: 'bottom', horizontal: 'center'}}
                  open={profileOpen}>

                  <Paper className={classes.profile}>
                    <div className={classes.profileBody}>
                      <Typography variant="subtitle1" color="secondary" paragraph align="center">Hi {user.name}</Typography>
                      <Typography variant="subtitle2" color="secondary" paragraph align="center">{user.email}</Typography>
                    </div>
                    <div className={classes.profileActions}>
                      <Button variant="contained" color="secondary" onClick={this.logoutHandler}>
                        Log out
                      </Button>
                    </div>
                  </Paper>
                </Popover>
              </div>

            </Toolbar>
          </AppBar>
          <Drawer open={open} variant="persistent" anchor="left" className={classes.drawer}
                  classes={{paper: classes.paper}}>

            <div className={classes.menu}>
              <Typography variant="h6" color="secondary">Menu</Typography>
              <IconButton onClick={this.closeDrawerHandler}>
                {theme.direction === 'ltr' ? <ChevronLeftIcon/> : <ChevronRightIcon/>}
              </IconButton>
            </div>

            <Divider/>

            <List component="nav" aria-label="main mailbox folders">
              <ListItem button onClick={() => route.to('list_user')} >
                <div className={classes.menuItem}>
                  <Person className={classes.menuIcon} color="secondary"/>
                  <Typography variant="subtitle2" color="secondary">List Users</Typography>
                </div>
              </ListItem>
              <ListItem button onClick={() => route.to('list_app')}>
                <div className={classes.menuItem}>
                  <Apps className={classes.menuIcon} color="secondary"/>
                  <Typography className={classes.menuItem} color="secondary" variant="subtitle2">List Apps</Typography>
                </div>
              </ListItem>
              <ListItem button onClick={() => route.to('list_roles')}>
                <div className={classes.menuItem}>
                  <Done className={classes.menuIcon} color="secondary"/>
                  <Typography className={classes.menuItem} color="secondary" variant="subtitle2">List Roles</Typography>
                </div>
              </ListItem>
            </List>


          </Drawer>
          <div className={classes.region}>
            <div className={clsx(classes.appBar, classes.root, {[classes.appBarShift]: open})}>
              <div className={classes.component}>
                <Display name="global"/>
              </div>
            </div>
          </div>
          <Error/>
        </div>
      </RouteProvider>

    )
  }
}


export default  connect(state => state.auth)(withUtils(styles)(Admin));
