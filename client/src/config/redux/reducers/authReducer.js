import { LOGIN_FAIL, LOGIN_SUCCESS, REGISTER_FAIL, REGISTER_SUCCESS, LOGOUT } from "../actions/actionTypes";

const user = JSON.parse(localStorage.getItem("user"));

const initialState = 
    user ? {
        isLogin: true,
        user
    } :
    {
        isLogin: false
    }

const authReducer = (state = initialState, action) => {

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