import {applyMiddleware, combineReducers, createStore} from 'redux';
import authReducer from './reducers/authReducer' 
import thunk from 'redux-thunk';

const store = createStore(combineReducers({
    authReducer
}), applyMiddleware(thunk));

export default store;