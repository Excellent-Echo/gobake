import React from 'react'
import { Header, } from './components'
import './App.scss'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { Login, Register } from './pages'
import Footer from './components/molecules/footer';


function App() {
  return (
   
      
    <Router>
    <div>
       <Header />
   </div>
     <Switch>
       <Route path="/login"><Login /></Route>
         <Route path="/register"><Register/></Route>
     </Switch>
     <Footer/>
    </Router>
      
  
    
  

   
  )
}

export default App