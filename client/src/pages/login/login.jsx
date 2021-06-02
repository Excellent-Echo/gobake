import Axios from 'axios'
import React, { useState } from 'react'
import swal from 'sweetalert'
import { API_URL, resetForm } from '../../utils/utils'

function Login() {

    const [inputEmail, setEmail] = useState("")
    const [inputPw, setPw] = useState("")

    const handleSubmit = (e) => {
        
        e.preventDefault()

        const data = {
            email : inputEmail,
            password : inputPw
        }

        console.log(data);

        Axios.post(`${API_URL}/users/login`, data)
            .then((res)=>{

                if (res.data.Data.token) {
                    localStorage.setItem("user", JSON.stringify(res.data.Data))
                }
                swal("Resgister success!", "", "success");
                resetForm("login-form")
                console.log(res)
            })
    }

    return (
        <div className="login-wrapper">
            <div className="login">
                <form onSubmit={handleSubmit} id="login-form">
                    <h2>
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
                <div className="btn-login">
                    <button>Log in</button>
                </div>
                </form>
            </div>
            
        </div>
    )
}

export default Login
