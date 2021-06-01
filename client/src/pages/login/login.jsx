import React from 'react'

function Login() {
    return (
        <div className="login-wrapper">
            <div className="login">
                <form id="login-form">
                    <h2>
                        Login
                    </h2>

                    <div className="input-container">
                    <label htmlFor="email">Email</label>
                    <input type="text" name="email" id="email" />
                </div>
                <div className="input-container">
                    <label htmlFor="password">Password</label>
                    <input type="password" name="password" id="password" />
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
