import React from 'react';
import "./RetroTextArea.css"
/**
 * props:
 * style
 * value (defaule value)
 * onChange
 */
class RetroTextArea extends React.Component {
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
        border: "2px solid " + color,
        padding: "3px",
        display: "inline-block",
        scrollBarBaseColor: "green"
      },
      value: ""
    };


    //copy the remaining style
    for (let key in props.style) {
      this.state.style[key] = props.style[key]
    }
    this.onTextAreaChange = this.onTextAreaChange.bind(this)
  }

  UNSAFE_componentWillReceiveProps(nextProps) {
    let newStyle = {...this.state.style}
    for (let key in nextProps.style) {
      newStyle[key] = nextProps.style[key]
    }
    this.setState({style: newStyle})
    if (nextProps.value) {
      this.setState({value: nextProps.value})
    }
  }

  onTextAreaChange(e) {
    this.setState({value: e.target.value})
    if (this.props.onChange) {
      this.props.onChange(e)
    }
  }

  render() {
    return (
      <textarea style={this.state.style} onChange={this.onTextAreaChange} value={this.state.value} className="retrotextarea" />
    )
  }

}

export default RetroTextArea