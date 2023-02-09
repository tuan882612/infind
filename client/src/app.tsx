import { Link } from "@mui/material";
import { Fragment, useEffect } from "react";

import './assets/styles/app.css'
import logo from './assets/images/infind-logo.png'
import Routing from "./routes";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Register from "./pages/register";
import NotFound from "./pages/notFound";
import Dashboard from "./pages/dashboard";

const App = () => {
  useEffect(() => {
    console.log('reloaded - app')
  }, [])

  return (
    <>
      <BrowserRouter>
        <Routing/>
      </BrowserRouter>
    </>
  );
}

export default App;