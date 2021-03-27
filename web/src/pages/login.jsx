import React from 'react';
import { withStyles, withTheme } from '@material-ui/core/styles';

import Box from '@material-ui/core/Box';
import SignIn from "@mlambda-net/core/packages/login/signin";

const styles = (themes) => ({
  root: {
    height: '100%',
    //backgroundImage: `url(${background})`,
    backgroundRepeat: 'no-repeat',
    backgroundSize: 'cover',
  },
  login: {
    width: '600px',
  },
});

class Login extends React.Component {
  render() {
    const { classes } = this.props;

    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        className={classes.root}>
        <SignIn className={classes.login} />
      </Box>
    );
  }
}

export default withStyles(styles)(withTheme(Login));
