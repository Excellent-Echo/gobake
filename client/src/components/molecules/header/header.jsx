import { FontAwesomeIcon, faBirthdayCake, faShoppingCart, faUser } from '../../../assets/index'
import React from 'react'
import "./header.scss"
import {useHistory} from "react-router-dom"
import { useSelector, useDispatch } from 'react-redux'
import { logout } from '../../../config/redux/actions/authAction'


function Header() {

    const {isLogin} = useSelector(state => state.authReducer)
    const history = useHistory();

    const dispatch = useDispatch()
    
    const handleLogout = () => {
        dispatch(logout());
    }

    return (
        <div className= "header-wrapper">
            <div className="header">

                <div className="logo">
                    <h2 onClick={() => history.push("/")}><FontAwesomeIcon icon={faBirthdayCake} /> GO-BAKE</h2>
                </div>

                <nav className="navbar">
                    <ul>
                        <li className="menu-header" onClick={() => history.push("/menu")}>Menu</li>
                        <li><FontAwesomeIcon icon={faShoppingCart}/> <span className="cart-qty">0</span> </li>
                        <li className="login-header" onClick={() => history.push("/login")}>
                            {isLogin ? <div className="logout-header" onClick={handleLogout}>Logout</div> : <FontAwesomeIcon icon={faUser}/>}
                        </li>
                    </ul>
                </nav>

            </div>
        </div>
    )
}

export default Header
