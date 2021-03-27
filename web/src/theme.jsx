import { createMuiTheme } from '@material-ui/core/styles';

export const theme = createMuiTheme({
  typography: {
    fontFamily: [
      '-apple-system',
      'BlinkMacSystemFont',
      '"Segoe UI"',
      'Roboto',
      '"Helvetica Neue"',
      'Arial',
      'sans-serif',
      '"Apple Color Emoji"',
      '"Segoe UI Emoji"',
      '"Segoe UI Symbol"',
    ].join(','),
  },
  palette: {
    primary: {
      main: '#c2185b',
      dark: '#8c0032',
      light: '#fa5788',
      contrastText: '#fce4ec',
    },
    secondary: {
      main: '#00796b',
      dark: '#004c40',
      light: '#48a999',
      contrastText: '#ffffff',
    },
    text: {
      primary: '#212121',
      secondary: '#757575',
    },
    divider: '#BDBDBD',
    contrastThreshold: 3,
    tonalOffset: 0.2,
  },
});
