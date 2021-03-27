import {IconButton, InputAdornment} from "@material-ui/core";
import React from "react";
import {Search} from "@material-ui/icons";


export default function SearchAdorn(props) {
  return (
    <InputAdornment position="end">
      <IconButton color="primary" onClick={props.onClick}>
        <Search/>
      </IconButton>
    </InputAdornment>
  )
}
