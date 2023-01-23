import { Fragment } from "react";

import './assets/styles/app.css'

import Routing from "./routes";
import logo from './assets/images/infind-logo.png'
import { Link } from "@mui/material";

const App = () => {
  return (
    <Fragment>
      <Link href="/home">
        <img className='logo' src={logo}/>
      </Link>
      <Routing/>
    </Fragment>
  );
}

export default App;