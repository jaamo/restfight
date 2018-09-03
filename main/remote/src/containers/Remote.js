import React from 'react';
import './Remote.scss'
import Login from '../components/Login.js'
import Controls from '../components/Controls.js'
import Alert from '../components/Alert.js'
import Spinner from '../components/Spinner';
import { connect } from 'react-redux';
import { showControls, changeTitle, getStatus, } from '../actions/actions.js';

class Remote extends React.Component {

	constructor(props) {
		super(props);
	}

	showControls = () => {
		this.props.showControls();
	};

	componentDidMount() {
		this.props.getStatus();
	}

	render() {
		return (
			<div className="component-app">

				{ this.props.view == 'login' &&
					<Login joinGame={this.showControls}></Login>
				}
				{ this.props.view == 'controls' &&
					<Controls></Controls>
				}
				{/* <Alert></Alert>
				<Spinner></Spinner> */}

			</div>
		);
	}
}

const mapStateToProps = (state, ownProps) => ({
	view: state.remote.view,
});
  
const mapDispatchToProps = {
	getStatus,
	showControls,
};
  
const RemoteContainer = connect(
	mapStateToProps,
	mapDispatchToProps
)(Remote);
  
export default RemoteContainer;	
