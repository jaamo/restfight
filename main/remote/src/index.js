import React from 'react';
import ReactDOM from 'react-dom';
import Remote from './containers/Remote.js';
import "./index.scss"
// https://blog.tylerbuchea.com/super-simple-react-redux-application-example/

ReactDOM.render(
  <Remote></Remote>,
  document.getElementById('app')
);

module.hot.accept();