import React from 'react'
import { Header, } from './components'
import './App.scss'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { Login, Register } from './pages'
import Footer from './components/molecules/footer';


function App() {
  return (
   
<<<<<<< HEAD
   <Router>
   <div>
      <Header />
  </div>
    <Switch>
        <Route path="/login"><Login /></Route>
        <Route path="/register"><Register/></Route>
    </Switch>
   </Router>
=======
      
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
      
  
    
  

   
>>>>>>> 252d5d5ad890b820ec97c3a949c8c2ded8e9d0ac
  )
}

export default App