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
  UNSAFE_componentWillReceiveProps(nextProps) {
    let newStyle = {...this.state.style}
    for (let key in nextProps.style) {
      newStyle[key] = nextProps.style[key]
    }
    this.setState({style: newStyle})
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