import React, { useState } from 'react'
import {resetForm } from '../../utils/utils'
import {Redirect, useHistory} from 'react-router-dom';
import {useDispatch, useSelector} from 'react-redux'
import { loginToAPI } from "../../config/redux/actions/authAction";

import './login.scss'

function Login() {

    const [inputEmail, setEmail] = useState("")
    const [inputPw, setPw] = useState("")
    const {isLogin} = useSelector(state => state.authReducer)

    const history = useHistory();

    const dispatch = useDispatch()
    
    const handleSubmit = (e) => {
        
        e.preventDefault()

        const data = {
            email : inputEmail,
            password : inputPw
        }

        dispatch(loginToAPI(data, resetForm("login-form"), history)) 
    }

    if (isLogin) {
        return <Redirect to="/" />
    }

    return (
        <div className="login-wrapper">
            <div className="login">
                <form onSubmit={handleSubmit} id="login-form">
                    <h2 className="login-title">
                        Login
                    </h2>
                    <div className="input-container">
                    <label htmlFor="email">Email</label>
                    <input onChange={(e) => setEmail(e.target.value)} type="text" name="email" id="email" />
                </div>

                <div className="input-container">
                    <label htmlFor="password">Password</label>
                    <input onChange={(e) => setPw(e.target.value) } type="password" name="password" id="password" />
                </div>
                <div className="forgot-login">
                    <p>Forgot Password?</p>
                </div>
                <div className="btn-login">
                    <button>Login</button>
                </div>
                <div className="toRegister">
                    <p onClick={() => history.push("/register")}>Or Sign Up</p>
                </div>
                </form>
            </div>
            
        </div>
    )
}

export default Login
