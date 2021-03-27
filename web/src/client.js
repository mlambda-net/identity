
import axios from 'axios';
import {store} from "./store/store";


axios.defaults.baseURL = process.env.REACT_APP_API_URL;

axios.interceptors.request.use((config) =>{
  const {auth} = store.getState()
  auth.data.manager.signinSilentCallback().then(c=>{})
  if(!auth.data.user.expired) {
    config.headers.Authorization = "bearer " + auth.data.user.access_token
  } else {

  }
  return config;
})

export default axios
