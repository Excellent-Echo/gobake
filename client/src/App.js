import React from "react";
import { Header, Footer } from "./components";
import "./App.scss";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { Home, Login, Menu, Register } from "./pages";

function App() {
  return (
    <Router>
      <div>
        <Header />
      <Switch>
        <Route path="/login">
          <Login />
        </Route>
        <Route path="/register">
          <Register />
        </Route>
        <Route path="/menu">
          <Menu />
        </Route>
        <Route path="/">
          <Home />
        </Route>
      </Switch>
        <Footer />
      </div>
    </Router>
  );
}

export default App;
