import { FontAwesomeIcon, faBirthdayCake, faShoppingCart, faUser } from '../../../assets/index'
import React from 'react'
import "./header.scss"
import {useHistory} from "react-router-dom"

function Header() {

    const history = useHistory()
    
    return (
        <div className= "header-wrapper">
            <div className="header">

                <div className="logo">
                    <h2><FontAwesomeIcon icon={faBirthdayCake} /> GO-BAKE</h2>
                </div>

                <nav className="navbar">
                    <ul>
                        <li>Menu</li>
                        <li><FontAwesomeIcon icon={faShoppingCart}/> <span className="cart-qty">0</span> </li>
                        <li onClick={() => history.push("/login")}><FontAwesomeIcon icon={faUser}/></li>
                    </ul>
                </nav>

            </div>
        </div>
    )
}

export default Header
