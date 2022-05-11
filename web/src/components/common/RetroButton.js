import React from 'react';
/**
 * Attributes: 
 * primary   boolean    use green style(default)
 * danger    boolean    use red style
 * warning   boolean    use orange style
 * style     object
 * onClick   funcion
 */
class RetroButton extends React.Component {
  constructor(props) {
    super(props);
    let color = "lightgreen"
    if (props.danger) {
      color = "red"
    } else if (props.warning) {
      color = "orange"
    }
    //set style
    this.state = {
      style: {
        color: color,
        backgroundColor: "black",
        border: "1px solid " + color,
        padding: "3px",
        display: "inline-block",
        cursor: "pointer",
        textAlign:"center",
      }
    };
    //copy the remaining style
    for (let key in props.style) {
      this.state.style[key] = props.style[key]
    }

    //bind this for methods
    this.onMouseEnterLeave = this.onMouseEnterLeave.bind(this)
    this.onClick = this.onClick.bind(this)
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
      <div style={this.state.style} onMouseEnter={this.onMouseEnterLeave} onMouseLeave={this.onMouseEnterLeave} onClick={this.onClick}>
        {this.props.children}
      </div>
    );
  }

  onClick() {
    if (this.props.onClick) {
      this.props.onClick()
    }
  }

  onMouseEnterLeave() {
    let newStyle = {...this.state.style}
    let tmp = newStyle.color
    newStyle.color = newStyle.backgroundColor
    newStyle.backgroundColor = tmp
    this.setState({style: newStyle})
  }

}
export default RetroButton;