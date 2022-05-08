import React from 'react';
import RetroCard from './common/RetroCard';
class RightBody extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      style: {
        width: "74%",
        verticalAlign: "top",
        marginLeft: "1%",
        padding: "10px",
      }
    }
    //copy the remaining style
    for (let key in props.style) {
      this.state.style[key] = props.style[key]
    }
  }
  render() {
    return (
      <RetroCard style={this.state.style}>
        {this.props.children}
      </RetroCard>
    )
  }
}
export default RightBody