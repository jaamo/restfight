import { applyMiddleware, createStore, } from 'redux';
import thunkMiddleware from 'redux-thunk';
import { reducers } from '../reducers/reducers.js';
  
export const store = createStore(reducers, applyMiddleware(thunkMiddleware));
  