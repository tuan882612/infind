import { 
  Route, 
  Routes,
  Navigate,
  BrowserRouter as Router
} from "react-router-dom";

import Home from "./pages/home";
import NotFound from "./pages/notFound";
import Login from "./pages/login";
import Register from "./pages/register";
import { Dashboard } from "@mui/icons-material";

const Routing = () => {
  return window.sessionStorage.getItem('login') === 'false'?
    <Routes>
      <Route 
        path="/" 
        element={<Navigate to="/home" replace/>}
      />
      <Route 
        path="/home" 
        element={<Home/>}
      />
      <Route 
        path="/login" 
        element={<Login/>}/>
      <Route 
        path="/register" 
        element={<Register/>}
      />
      <Route 
        path="*" 
        element={<NotFound/>}
      />
    </Routes>:
    <Routes>
      <Route 
        path="/dashboard"
        element={<Dashboard/>}
      />
    </Routes>
}

export default Routing;