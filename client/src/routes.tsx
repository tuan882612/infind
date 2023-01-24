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

const Routing = () => (
  <Router>
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
        path="/dashboard"
        element={<Dashboard/>}
      />
      <Route 
        path="*" 
        element={<NotFound/>}
      />
    </Routes>
  </Router>
);

export default Routing;