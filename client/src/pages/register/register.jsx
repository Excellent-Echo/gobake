import React, { useState } from "react";
import { resetForm } from "../../utils/utils";
import {useHistory} from 'react-router-dom';
import {useDispatch, useSelector} from 'react-redux'
import { registerToAPI } from "../../config/redux/actions/authAction";

import "./register.scss"

function Register() {
  const [inputFirst, setInputFirst] = useState("");
  const [inputLast, setInputLast] = useState("");
  const [inputUsername, setInputUsername] = useState("");
  const [inputEmail, setInputEmail] = useState("");
  const [inputPassword, setInputPassword] = useState("");
  const [inputPhone, setInputPhone] = useState("");
  const {isLogin} = useSelector(state => state.authReducer)

  const history = useHistory();
//dispatch untuk memanggil action
  const dispatch = useDispatch()

  const handleSubmit = (e) => {
    e.preventDefault();

    const data = {
      first_name: inputFirst,
      last_name: inputLast,
      email: inputEmail,
      username: inputUsername,
      password: inputPassword,
      phone: parseInt(inputPhone),
      address: "-",
    };
    
    dispatch(registerToAPI(data, resetForm("regis-form"), history))

  };

  if (isLogin) {
    return history.push('/')
}
   
  return (
    <div className="regis-wrapper">
      <div className="regis">
        <form onSubmit={handleSubmit} id="regis-form">
          <h2 className="register-title">Create new account</h2>
          <div className="input-container">
            <label htmlFor="firstname">First name</label>
            <input
              onChange={(e) => setInputFirst(e.target.value)}
              type="text"
              name="firstname"
              id="firstname"
            />
          </div>

          <div className="input-container">
            <label htmlFor="lastname">Last Name</label>
            <input
              onChange={(e) => setInputLast(e.target.value)}
              type="text"
              name="lastname"
              id="lastname"
            />
          </div>

          <div className="input-container">
            <label htmlFor="username">Username</label>
            <input
              onChange={(e) => setInputUsername(e.target.value)}
              type="text"
              name="username"
              id="username"
            />
          </div>

          <div className="input-container">
            <label htmlFor="email">Email</label>
            <input
              onChange={(e) => setInputEmail(e.target.value)}
              type="text"
              name="email"
              id="email"
            />
          </div>

          <div className="input-container">
            <label htmlFor="password">Password</label>
            <input
              onChange={(e) => setInputPassword(e.target.value)}
              type="password"
              name="password"
              id="password"
            />
          </div>

          <div className="input-container">
            <label htmlFor="phone">Phone</label>
            <input
              onChange={(e) => setInputPhone(e.target.value)}
              type="number"
              name="phone"
              id="phone"
            />
          </div>

          <div className="btn-regis">
            <button>Sign up</button>
          </div>
        </form>
      </div>
    </div>
  );
}

export default Register;
