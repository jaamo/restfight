import React from 'react';
import './Controls.scss'

class Controls extends React.Component {
  
    render() {
      return (
        <div className="controls">

            <h2>Move</h2>

            <div className="controls__move-row1">
                <button className="btn"><i>A</i><span>Left</span></button>
            </div>
            <div className="controls__move-row2">
                <button className="btn"><i>D</i><span>Right</span></button>
                <button className="btn"><i>W</i><span>Up</span></button>
                <button className="btn"><i>S</i><span>Down</span></button>
            </div>

            <h2>Shoot</h2>

            <div className="form__row">
              <div className="form__row__label">Target x</div>
              <div className="form__row__input"><input type="text" value="1" readOnly></input></div>
            </div>

            <div className="form__row">
              <div className="form__row__label">Target y</div>
              <div className="form__row__input"><input type="text" value="1" readOnly></input></div>
            </div>

            <div className="form__row">
              <div className="form__row__label"></div>
              <div className="form__row__input"><button className="btn"><span>Shoot</span></button></div>
            </div>

            <h2>End turn</h2>

            <div className="btn-row">
                <button className="btn"><i>T</i><span>End turn</span></button>
            </div>

        </div>
      );
    }
  }
  export default Controls;