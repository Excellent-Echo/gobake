import { LOGIN_FAIL, LOGIN_SUCCESS, REGISTER_FAIL, REGISTER_SUCCESS, LOGOUT } from "../actions/actionTypes";

// narik data user dari local stroage
const user = JSON.parse(localStorage.getItem("user"));

//cek state jika user ada, login = true, simpan user di state
const initialState = 
    user ? {
        isLogin: true,
        user
    } :
    {
        isLogin: false,
    }

const authReducer = (state = initialState, action) => {
//action jika register dan login berhasil / tidak
    if (action.type === REGISTER_SUCCESS) {
        return {
            ...state,
            isLogin: false
        }
    }
    if (action.type === REGISTER_FAIL) {
        return {
            ...state,
            isLogin: false
        }
    }
    if (action.type === LOGIN_SUCCESS) {
        return {
            ...state,
            isLogin: true
        }
    }
    if (action.type === LOGIN_FAIL) {
        return {
            ...state,
            isLogin: false
        }
    }
//action for logout
    if (action.type === LOGOUT) {
        return {
            ...state,
            isLogin: false,
            user: null
        }
    }

    return state
}

export default authReducer;