import { Routes, Route } from "react-router-dom";

import Login from "./pages/Login";
import Register from "./pages/Register";
import Products from "./pages/Products";
import Cart from "./pages/Cart";
import Orders from "./pages/Orders";
import Admin from "./pages/Admin";

import PrivateRoute from "./components/PrivateRoute";
import AdminRoute from "./components/AdminRoute";


function App() {

  return (

    <Routes>

      <Route 
        path="/" 
        element={<Login />} 
      />


      <Route 
        path="/register" 
        element={<Register />} 
      />


      <Route
        path="/products"
        element={
          <PrivateRoute>
            <Products />
          </PrivateRoute>
        }
      />


      <Route
        path="/cart"
        element={
          <PrivateRoute>
            <Cart />
          </PrivateRoute>
        }
      />


      <Route
        path="/orders"
        element={
          <PrivateRoute>
            <Orders />
          </PrivateRoute>
        }
      />


      {/* ADMIN ROUTE */}

      <Route
        path="/admin"
        element={
          <AdminRoute>
            <Admin />
          </AdminRoute>
        }
      />


    </Routes>

  );

}


export default App;