import React from 'react';

class Login extends React.Component {
	
	
	constructor(props) {
		super(props);
		this.state = {
			weapon: 1,
			shield: 1,
			engine: 0,
		};
		this.handleChange = this.handleChange.bind(this);		
	}
	
	handleChange(e) {
		let newState = {};
		newState[e.target.attributes.name.value] = e.target.value;
		this.setState(newState);
	}

	render() {
		return (
			<div className="component-app">
			
			<h2>Login</h2>
			
			<div className="form__row">
			<div className="form__row__label">Weapon level</div>
			<div className="form__row__input"><input type="text" name="weapon" value={this.state.weapon} onChange={this.handleChange}></input></div>
			</div>
			
			<div className="form__row">
			<div className="form__row__label">Engine level</div>
			<div className="form__row__input"><input type="text" name="engine" value={this.state.engine} onChange={this.handleChange}></input></div>
			</div>
			
			<div className="form__row">
			<div className="form__row__label">Shield level</div>
			<div className="form__row__input"><input type="text" name="shield" value={this.state.shield} onChange={this.handleChange}></input></div>
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