import { LOGIN_FAIL, LOGIN_SUCCESS, LOGOUT, REGISTER_FAIL, REGISTER_SUCCESS } from "../actions/actionTypes";
import axios from 'axios';
import swal from 'sweetalert'
import { API_URL } from "../../../utils/utils";


//fungsi untuk resgister
//param data =ambil data dri input, reset = reset form
//history = jika register berhasil lempar url ke login page
export const registerToAPI = (data, reset, history) => async (dispatch) => {
    await axios.post(`${API_URL}/users/register`, data)
    .then((res) => {
        swal("Resgister success!", "", "success");
        setTimeout(() => {
            history.push("/login")
        }, 1500)
        dispatch({
            type: REGISTER_SUCCESS
        })
        reset()
    })
    .catch((err) => {
      console.log(err.message)
      dispatch({
          type: REGISTER_FAIL
      })
    });
}

export const loginToAPI = (data, reset, history) => async (dispatch) => {
    await axios.post(`${API_URL}/users/login`, data)
    .then((res)=>{

        //set user dlm local storage jika token ada
        if (res.data.data.token) {
            localStorage.setItem("user", JSON.stringify(res.data.data))
        }
        swal("Login Success", "", "success");
        dispatch({
            type: LOGIN_SUCCESS
        })
        setTimeout( async () => {
            await history.push("/")
        }, 1500)
        reset()
    }).catch((err) => {
        console.log(err);
        dispatch({
            type: LOGIN_FAIL
        })
    })
}

export const logout = () => (dispatch) => {
    localStorage.removeItem("user");
    dispatch({
        type: LOGOUT
    })
}