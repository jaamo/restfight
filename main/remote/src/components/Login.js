import React from 'react';

class Login extends React.Component {
  
    render() {
      return (
        <div className="component-app">

            <h2>Login</h2>

            <div className="form__row">
              <div className="form__row__label">Weapon level</div>
              <div className="form__row__input"><input type="text" value="1" readOnly></input></div>
            </div>

            <div className="form__row">
              <div className="form__row__label">Engine level</div>
              <div className="form__row__input"><input type="text" value="1" readOnly></input></div>
            </div>

            <div className="form__row">
              <div className="form__row__label">Shield level</div>
              <div className="form__row__input"><input type="text" value="0" readOnly></input></div>
            </div>

            <div className="form__row">
              <div className="form__row__label"><button className="btn"><span>Reset game</span></button></div>
              <div className="form__row__input"><button className="btn" onClick={this.props.showControls}><i>A</i><span>Join game</span></button></div>
            </div>

        </div>
      );
    }
  }
  export default Login;