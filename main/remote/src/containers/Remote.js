import React from 'react';
import './Remote.scss'
import Login from '../components/Login.js'
import Controls from '../components/Controls.js'
import Alert from '../components/Alert.js'
import Spinner from '../components/Spinner';

class App extends React.Component {

		constructor(props) {
			super(props);
			this.state = {
				view: 'controls'
			};
		}
	
		showControls = () => {
			this.setState({view:'controls'});
		};
	
		render() {
			return (
				<div className="component-app">

					{ this.state.view == 'login' &&
						<Login showControls={this.showControls}></Login>
					}
					{ this.state.view == 'controls' &&
						<Controls></Controls>
					}
					{/* <Alert></Alert>
					<Spinner></Spinner> */}

				</div>
			);
		}
	}
	export default App;