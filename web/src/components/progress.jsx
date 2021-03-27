import {Divider, LinearProgress} from "@material-ui/core";
import React from "react";


function Progress(props) {
  if (props.loading){
    return <LinearProgress color="primary"  />
  }
  return<Divider/>
}


export default Progress
