
import axios from 'axios';
import {store} from "./store/store";


axios.defaults.baseURL = process.env.REACT_APP_API_URL;

axios.interceptors.request.use((config) =>{
  const {auth} = store.getState()

  if(!auth.data.user.expired) {
    config.headers.Authorization = "bearer " + auth.data.user.access_token
  } else {
    auth.data.renew()
  }
  return config;
})

export default axios
