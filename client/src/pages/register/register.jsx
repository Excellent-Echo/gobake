import React from 'react'

function Register() {
    return (
        <div className="regis-wrapper">

            <div className="regis">
                <form id="regis-form">
                <h2>
                    Create new account
                </h2>

                <div className="input-container">
                    <label htmlFor="firstname">First name</label>
                    <input type="text" name="firstname" id="firstname" />
                </div>
                <div className="input-container">
                    <label htmlFor="lastname">Last Name</label>
                    <input type="text" name="lastname" id="lastname" />
                </div>
                <div className="input-container">
                    <label htmlFor="username">Username</label>
                    <input type="text" name="username" id="username" />
                </div>
                <div className="input-container">
                    <label htmlFor="email">Email</label>
                    <input type="text" name="name" id="name" />
                </div>
                <div className="input-container">
                    <label htmlFor="password">Password</label>
                    <input type="password" name="password" id="password" />
                </div>
                <div className="input-container">
                    <label htmlFor="phone">Phone</label>
                    <input type="number" name="phone" id="phone" />
                </div>
                <div className="btn-regis">
                    <button>Sign up</button>
                </div>
                </form>
            </div>
            
        </div>
    )
}

export default Register
