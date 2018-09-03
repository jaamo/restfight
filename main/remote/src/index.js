import React from 'react';
import ReactDOM from 'react-dom';
import Remote from './containers/Remote.js';
import "./index.scss"

import { Provider } from 'react-redux';
import { store } from './store/store.js';

ReactDOM.render(
  <Provider store={store}>
    <Remote></Remote>
  </Provider>,
  document.getElementById('app')
);

module.hot.accept();