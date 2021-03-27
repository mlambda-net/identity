

const settings = {
  authority: process.env.REACT_APP_OAUTH_SERVER,
  client_id: process.env.REACT_APP_CLIENT_ID,
  client_secret: process.env.REACT_APP_SECRET,
  redirect_uri: process.env.REACT_APP_HOST,
  scope: 'all',
  response_type: 'code',
  metadata: {
    token_endpoint: process.env.REACT_APP_OAUTH_SERVER + "/token",
    authorization_endpoint: process.env.REACT_APP_OAUTH_SERVER + "/authorize"
  }
}

export default  settings;
