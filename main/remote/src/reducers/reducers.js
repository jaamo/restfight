import { combineReducers, } from 'redux';

const initialState = { 
	view: 'login',
	title: 'pers',
};

const remote = (state = initialState, action) => {
	switch (action.type) {
		case 'SHOW_CONTROLS':
			return { view: 'controls' };
		default:
			return state;
	}
};

export const reducers = combineReducers({
	remote,
});