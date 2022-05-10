import React from 'react';

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
      type:"text",
      content:""
    };

    if(props.password){
      this.state.type="password"
    }

    //copy the remaining style
    for (let key in props.style) {
      this.state.style[key] = props.style[key]
    }
    this.onChange=this.onChange.bind(this)
  }

  onChange(e){
    if(this.props.onChange){
      this.props.onChange(e)
    }
  }

  render(){
    return (
      <input type={this.state.type} style={this.state.style} onChange={this.onChange} />
    )
  }
  
}

export default RetroInput