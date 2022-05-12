import React from 'react';
/**
 * props:
 * style object
 * value (defaule value)
 * onChange function
 */
class RetroInput extends React.Component {
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
      },
      type: "text",
    };

    if (props.password) {
      this.state.type = "password"
    }

    //copy the remaining style
    for (let key in props.style) {
      this.state.style[key] = props.style[key]
    }
    this.onInputChange = this.onInputChange.bind(this)
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
      <input type={this.state.type} style={this.state.style} onChange={this.onInputChange} value={this.props.value} />
    )
  }

  onInputChange(e) {
    if (this.props.onChange) {
      this.props.onChange(e)
    }
  }

}

export default RetroInput