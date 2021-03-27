import React from 'react';
import {ThemeProvider} from '@material-ui/core/styles';
import { theme } from './theme';
import Admin from "./pages/admin";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";

import {Auth} from "./store/actions";
import settings from "./oauth";
import {connect} from "react-redux";
import {AuthProvider} from "@mlambda-net/core/packages/oauth/authprovider";

const styles = (themes) => ({
});

function Display(props) {
  const {isAuthenticate} = props
  if(isAuthenticate) {
    return <Admin/>
  }
  return <div/>
}


class App extends React.Component {

  constructor(props) {
    super(props);
    this.dispatch = props.dispatch
    this.state = {isAuthenticate: false}
    this.onLoginHandler = this.onLogin.bind(this)
  }

  onLogin(auth) {
    this.dispatch({type: Auth.SetAuth, payload: auth})
    this.setState({isAuthenticate: true})
  }

  render = () => {

    return (
        <AuthProvider settings={settings} onLogin={this.onLoginHandler}>
          <ThemeProvider theme={theme}>
            <Display isAuthenticate={this.state.isAuthenticate}/>
          </ThemeProvider>
        </AuthProvider>
    );
  }
}

export default connect(state => state.auth)(withUtils(styles)(App));
