import React from 'react';
import './Alert.scss'

class Alert extends React.Component {
  
    render() {
      return (
        <div className="alert">

            <div className="alert__close"><span>+</span></div>
            <div className="alert__description">Illegal move. Cannot move outside area.</div>

        </div>
      );
    }
  }
  export default Alert;